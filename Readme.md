INSTALL
-------
go get -u github.com/msonawane/grpc-example/greeter_client
go get -u github.com/msonawane/grpc-example/greeter_server

TRY IT!
-------
greeter_server &
greeter_client Manoj

OPTIONAL - Rebuilding the generated code
----------------------------------------

1 First [install protoc](https://github.com/google/protobuf/blob/master/INSTALL.txt)
  - For now, this needs to be installed from source
  - This is will change once proto3 is officially released

2 Install the protoc Go plugin.

```
$ go get -a github.com/golang/protobuf/protoc-gen-go
$
$ # from this dir; invoke protoc
run protoc -I ./protos/ ./protos/helloworld.proto --go_out=plugins=grpc:protos
