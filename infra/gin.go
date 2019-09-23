package infra

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// GinServer is the struct gathering all the server details
type GinServer struct {
	host   string
	port   int
	Router *gin.Engine
}

// NewServer creates the gin Server
func NewServer(host string, port int, mode string, corsc *cors.Config) GinServer {
	s := GinServer{}
	s.port = port
	s.host = host

	setMode(mode)

	s.Router = gin.New()

	s.Router.Use(gin.Recovery())

	if corsc != nil {
		s.Router.Use(cors.New(*corsc))
	}

	return s
}

func setMode(mode string) {
	switch mode {
	case "debug":
		gin.SetMode(gin.DebugMode)
	case "test":
		gin.SetMode(gin.TestMode)
	case "release":
		gin.SetMode(gin.ReleaseMode)
	default:
		logrus.WithField("mode", mode).Warn("Unknown gin mode, fallback to release")
		gin.SetMode(gin.ReleaseMode)
	}
}

// Start tells the router to start listening
func (s GinServer) Start() {
	logrus.WithFields(logrus.Fields{"port": s.port, "host": s.host}).Info("Listening")
	if err := s.Router.Run(fmt.Sprintf("%s:%d", s.host, s.port)); err != nil {
		logrus.WithError(err).Fatal("Couldn't start router")
	}
}
