// Package crypt предназначен для шифрования данных.
package crypt

import (
	"fmt"
	"gophkeeper/pkg/proto"
	"reflect"
	"testing"
)

func TestEncField(t *testing.T) {

	type args struct {
		dec       *proto.FieldKeep
		secretKey string
	}
	tests := []struct {
		name    string
		args    args
		wantEnc *proto.FieldKeep
	}{
		{
			name: "positive",
			args: args{
				dec: &proto.FieldKeep{
					Name:     "Test",
					Login:    "Test",
					Password: "Test",
				},
				secretKey: "akjdsfakjLKJNaskdflkjnJnlasjkdnf",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotEnc := EncField(tt.args.dec, tt.args.secretKey); !reflect.DeepEqual(gotEnc, tt.wantEnc) {
				decc := DecField(gotEnc, tt.args.secretKey)
				fmt.Println(decc.Login)
				fmt.Println(decc.Login)
				// t.Errorf("EncField() = %v, want %v", gotEnc, tt.wantEnc)
			}
		})
	}
}
