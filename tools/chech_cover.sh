#!/bin/bash

go test ./... -coverprofile=coverage.txt
grep -v "gophkeeper/pkg/proto" coverage.txt  > cover.txt
mv cover.txt coverage.txt
go tool cover -func coverage.txt
go tool cover -html coverage.txt -o index.html