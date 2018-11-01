#!/usr/bin/env bash
protoc -I . --go_out=plugins=micro:./ ./proto/user/user.proto
go build -o ./bin/user-service.exe
