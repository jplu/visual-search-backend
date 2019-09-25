# visual-search-backend
BFF over a Tensorflow and FAISS index gRPC server for visual search.

## Requirements

* Go 1.11 and later
* gRPC server providing the [Tensorflow model](https://github.com/jplu/visual-search-scripts)
* gRPC server providing the [FAISS index](https://github.com/jplu/faiss-grpc-server)

And then install Tensorflow and Tensorflow serving Go packages:
```
git clone -b r1.15 https://github.com/tensorflow/tensorflow.git
git clone -b r1.15 https://github.com/tensorflow/serving.git
GO111MODULE=off go run protoc.go
cd proto/tensorflow/core && go mod init github.com/tensorflow/tensorflow/tensorflow/go/core && cd -
go build ./proto/tensorflow/serving
```

## Compiling and creating the Docker image

Compiling:
```
make build
```

Create Docker image:
```
make docker
```

## Run
To run the visual search backend, create a `conf.yaml` from the example file to put your values, and finally run:
```
./visual-search-backend
```

Or you can your values through the corresponding parameters:
```
visual-search-backend

Usage:
  visual-search-backend [flags]
  visual-search-backend [command]

Available Commands:
  help        Help about any command
  version     Show build and version

Flags:
      --ann.host string               host on which the ANN gRPC server should listen (default "localhost")
      --ann.port int                  port on which the ANN gRPC server should listen (default 8080)
      --conf string                   configuration file to use
  -h, --help                          help for visual-search-backend
      --img.host string               host on which the Image gRPC server should listen (default "localhost")
      --img.port int                  port on which the ANN gRPC server should listen (default 8080)
      --log.format string             one of text or json (default "text")
      --log.level string              one of debug, info, warn, error or fatal (default "debug")
      --log.line                      enable filename and line in logs (default true)
      --server.cors.all               defines that all origins are allowed (default true)
      --server.cors.disabled          disable CORS completely
      --server.cors.expose strings    array of exposed headers
      --server.cors.headers strings   array of allowed headers (default [Origin,Authorization,Content-Type])
      --server.cors.methods strings   array of allowed method when cors is enabled (default [GET,PUT,POST,DELETE,OPTION,PATCH])
      --server.cors.origins strings   array of allowed origins (overwritten if all is active)
      --server.host string            host on which the server should listen (default "localhost")
      --server.mode string            server mode can be either 'debug', 'test' or 'release' (default "debug")
      --server.port int               port on which the server should listen (default 8081)

Use "visual-search-backend [command] --help" for more information about a command.
```
