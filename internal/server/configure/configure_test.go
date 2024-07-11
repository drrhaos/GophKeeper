// Package configure предназначен для настройки программы.
package configure

import (
	"os"
	"testing"

	"gotest.tools/v3/assert"
)

func TestConfig_ReadConfig(t *testing.T) {
	os.Args = append(os.Args, "--s=/home/user")
	os.Args = append(os.Args, "--d=test")

	type want struct {
		cfg Config
		ok  bool
	}
	tests := []struct {
		name  string
		want  want
		isEnv bool
	}{
		{
			name:  "positive test #1",
			isEnv: false,
			want: want{
				ok: true,
				cfg: Config{
					Port:        "8080",
					PortRest:    "8081",
					WorkPath:    "./home",
					StaticPath:  "/home/user",
					DatabaseDsn: "test",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var cfg Config
			ok := cfg.ReadConfig()
			assert.Equal(t, tt.want.ok, ok)
			assert.Equal(t, tt.want.cfg.Port, cfg.Port)
			assert.Equal(t, tt.want.cfg.PortRest, cfg.PortRest)
			assert.Equal(t, tt.want.cfg.StaticPath, cfg.StaticPath)
			assert.Equal(t, tt.want.cfg.DatabaseDsn, cfg.DatabaseDsn)
		})
	}
}
