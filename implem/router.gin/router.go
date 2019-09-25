package router

import (
	"github.com/jplu/visual-search-backend/uc"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Router is a struct embedding an interactor and a gin engine
type Router struct {
	interactor uc.LogicHandler
	router     *gin.Engine
}

// NewRouter will return a new router
func NewRouter(i uc.LogicHandler, r *gin.Engine) Router {
	return Router {
		interactor: i,
		router:     r,
	}
}

// SetRoutes will create the necessary routes
func (r Router) SetRoutes() {
	r.router.Use(gin.Logger())

	// Simple health check returning 200 OK
	r.router.GET("/health", func(c *gin.Context) { c.Status(http.StatusOK) })

	// API Group
	api := r.router.Group("/api/v1")
	{
		api.POST("/similar", r.Similar)
	}
}
