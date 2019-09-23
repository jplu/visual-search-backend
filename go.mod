module github.mpi-internal.com/leboncoin-lab/visual-search-backend

require (
	github.com/gin-contrib/cors v0.0.0-20181008113111-488de3ec974f
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/gin-gonic/gin v1.4.0
	github.com/golang/protobuf v1.3.2
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/json-iterator/go v1.1.7 // indirect
	github.com/konsorten/go-windows-terminal-sequences v1.0.1 // indirect
	github.com/mattn/go-isatty v0.0.9 // indirect
	github.com/mitchellh/mapstructure v1.1.2 // indirect
	github.com/onrik/logrus v0.0.0-20181009124311-c9849815bc7c
	github.com/sirupsen/logrus v1.1.1
	github.com/spf13/cobra v0.0.3
	github.com/spf13/pflag v1.0.3 // indirect
	github.com/spf13/viper v1.2.1
	github.com/tensorflow/tensorflow/tensorflow/go/core v0.0.0-00010101000000-000000000000
	github.com/ugorji/go v1.1.7 // indirect
	golang.org/x/net v0.0.0-20190813141303-74dc4d7220e7 // indirect
	google.golang.org/grpc v1.23.0
)

replace github.com/tensorflow/tensorflow/tensorflow/go/core => ./proto/tensorflow/core
