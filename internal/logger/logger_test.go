// Модуль logger предназначен для логирования в агенте и сервере.

package logger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequestLogger(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{
			name: "positive test #1",
			want: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			Initialize("info")

			assert.Equal(t, test.want, true)
		})
	}
}
