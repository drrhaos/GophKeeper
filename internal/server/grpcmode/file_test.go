package grpcmode

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFile_SetFile(t *testing.T) {
	fileTmp := "/tmp/agent.json"
	defer os.Remove(fileTmp)
	dirPath, nameFile := filepath.Split(fileTmp)

	type args struct {
		fileName string
		path     string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Positive test #1",
			args: args{
				fileName: nameFile,
				path:     dirPath,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := NewFile()
			if err := f.SetFile(tt.args.fileName, tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("File.SetFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFile_Write(t *testing.T) {
	fileTmp := "/tmp/agent.json"
	defer os.Remove(fileTmp)
	dirPath, nameFile := filepath.Split(fileTmp)

	type args struct {
		fileName string
		path     string
		chunk    []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Positive test #1",
			args: args{
				fileName: nameFile,
				path:     dirPath,
				chunk:    []byte("test"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := NewFile()
			err := f.SetFile(tt.args.fileName, tt.args.path)
			if err != nil {
				t.Error(err)
			}

			if err := f.Write(tt.args.chunk); (err != nil) != tt.wantErr {
				t.Errorf("File.Write() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFile_Close(t *testing.T) {
	fileTmp := "/tmp/agent.json"
	defer os.Remove(fileTmp)
	dirPath, nameFile := filepath.Split(fileTmp)

	type args struct {
		fileName string
		path     string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Positive test #1",
			args: args{
				fileName: nameFile,
				path:     dirPath,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := NewFile()
			err := f.SetFile(tt.args.fileName, tt.args.path)
			if err != nil {
				t.Error(err)
			}
			if err := f.Close(); (err != nil) != tt.wantErr {
				t.Errorf("File.Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
