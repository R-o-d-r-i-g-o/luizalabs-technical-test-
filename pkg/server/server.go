package server

import (
	"github.com/gin-gonic/gin"
)

// HandlerImp interface defines a contract for registering routes.
type HandlerImp interface {
	Register(g *gin.RouterGroup)
}

// GinServerImp interface defines the methods required for a server.
type GinServerImp interface {
	Run(addr string) error
	SetupHandlers(version string, handlers ...func(*gin.RouterGroup))
	SetupMiddleware(middleware ...gin.HandlerFunc)
}

// ginServer struct implements the Server interface.
type ginServer struct {
	router *gin.Engine
}

// NewGinServer creates a new instance of the Gin server.
func NewGinServer() GinServerImp {
	return &ginServer{router: gin.Default()}
}

// Run starts the server on the specified address.
func (s *ginServer) Run(addr string) error {
	return s.router.Run(addr)
}

// SetupMiddleware sets up middleware for the router.
func (s *ginServer) SetupMiddleware(middleware ...gin.HandlerFunc) {
	for _, mw := range middleware {
		s.router.Use(mw)
	}
}

// SetupHandlers registers multiple hander setup functions in the server.
func (s *ginServer) SetupHandlers(version string, handlers ...func(*gin.RouterGroup)) {
	apiVersioning := s.router.Group("/" + version)

	for _, route := range handlers {
		route(apiVersioning)
	}
}
