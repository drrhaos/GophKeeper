// Package configure предназначен для настройки программы.
package configure

import (
	"os"
	"testing"

	"gotest.tools/v3/assert"
)

func TestConfig_ReadConfig(t *testing.T) {
	os.Args = append(os.Args, "--g=127.0.0.1:9080")
	os.Args = append(os.Args, "--c=ca.crt")

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
					Address: "127.0.0.1:9080",
					Secret:  "test",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var cfg Config
			ok := cfg.ReadConfig()
			assert.Equal(t, tt.want.ok, ok)
			// assert.Equal(t, tt.want.cfg.Address, cfg.Address)
			// assert.Equal(t, tt.want.cfg.CAFile, cfg.CAFile)
			// assert.Equal(t, tt.want.cfg.Secret, cfg.Secret)
		})
	}
}
