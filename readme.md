# Repositório do curso 'gRPC [Golang] Master Class: Build Modern API & Microservices'
https://www.udemy.com/course/grpc-golang/

---

# Commands
protoc -Igreet/proto --go_out=. --go_opt=module=github.com/rodrigueslg/golang-grpc-udemy --go-grpc_out=. --go-grpc_opt=module=github.com/rodrigueslg/golang-grpc-udemy  greet/proto/dummy.proto
