#!/bin/bash


git_hash=`git tag | head -n 1`

current_time=`date +"%Y-%m-%d:T%H:%M:%S"`

go build -ldflags "-X main.buildDate=${current_time} -X main.buildVersion=${git_hash}" -o ./cmd/client/client ./cmd/client/