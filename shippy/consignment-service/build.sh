 protoc -I . --go_out=plugins=micro:./ ./proto/consignment/consignment.proto
 go build -o ./bin/consignment-service.exe

