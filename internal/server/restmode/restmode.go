// Package restmode REST сервер.
package restmode

import (
	"context"
	"fmt"
	"mime"
	"net/http"
	"strings"

	"gophkeeper/internal/server/configure"

	pb "gophkeeper/pkg/proto"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func serveSwagger(mux *http.ServeMux, cfg configure.Config) {
	mime.AddExtensionType(".svg", "image/svg+xml")

	fileServer := http.FileServer(http.Dir(cfg.StaticPath))
	prefix := "/swagger-ui/"
	mux.Handle(prefix, http.StripPrefix(prefix, fileServer))
}

var allowedHeaders = map[string]struct{}{
	"Authorization": {},
}

func isHeaderAllowed(s string) (string, bool) {
	// check if allowedHeaders contain the header
	if _, isAllowed := allowedHeaders[s]; isAllowed {
		// send uppercase header
		return strings.ToUpper(s), true
	}
	// if not in the allowed header, don't send the header
	return s, false
}

// Run запуск REST сервера.
func Run(cfg configure.Config) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := http.NewServeMux()

	gwmux := runtime.NewServeMux(
		runtime.WithOutgoingHeaderMatcher(isHeaderAllowed),
		runtime.WithMetadata(func(_ context.Context, request *http.Request) metadata.MD {
			md := map[string]string{
				"Authorization": request.Header.Get("Authorization"),
			}

			return metadata.New(md)
		}))
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := pb.RegisterGophKeeperHandlerFromEndpoint(ctx, gwmux, fmt.Sprintf("127.0.0.1:%s", cfg.Port), opts)
	if err != nil {
		panic(err)
	}

	mux.Handle("/", gwmux)
	serveSwagger(mux, cfg)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", cfg.PortRest), mux); err != nil {
		panic(err)
	}
}
