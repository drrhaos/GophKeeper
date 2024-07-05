// Package crypt предназначен для шифрования данных.
package crypt

import (
	"reflect"
	"testing"

	"gophkeeper/pkg/proto"

	"github.com/stretchr/testify/assert"
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
					Password: "",
				},
				secretKey: "akjdsfakjLKJNaskdflkjnJnlasjkdnf",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotEnc := EncField(tt.args.dec, tt.args.secretKey); !reflect.DeepEqual(gotEnc, tt.wantEnc) {
				decr := DecField(gotEnc, tt.args.secretKey)
				assert.Equal(t, tt.args.dec.Name, decr.Name)

			}
		})
	}
}

func TestDecField(t *testing.T) {
	type args struct {
		enc       *proto.FieldKeep
		secretKey string
	}
	tests := []struct {
		name    string
		args    args
		wantDec *proto.FieldKeep
	}{
		{
			name: "positive",
			args: args{
				enc: &proto.FieldKeep{
					Name:     "OXxQERV0nvMhHmW8GmQV37CBQsXCI6ux928mlw==",
					Login:    "OXxQERV0nvMhHmW8GmQV37CBQsXCI6ux928mlw==",
					Password: "OXxQERV0nvMhHmW8GmQV37CBQsXCI6ux928mlw==",
				},
				secretKey: "akjdsfakjLKJNaskdflkjnJnlasjkdnf",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDec := DecField(tt.args.enc, tt.args.secretKey); !reflect.DeepEqual(gotDec, tt.wantDec) {
				assert.Equal(t, "", gotDec.Name)
			}
		})
	}
}
