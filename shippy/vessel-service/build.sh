 protoc -I . --go_out=plugins=micro:./ ./proto/vessel/vessel.proto
go build -o ./bin/vessel-service.exe
