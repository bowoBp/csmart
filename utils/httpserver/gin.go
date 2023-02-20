package httpserver

import "github.com/gin-gonic/gin"

type Gin struct {
	engine *gin.Engine
}

// New creates a new instance of the Gin http server with the provided gin.Engine.
func New(engine *gin.Engine) *Gin {
	return &Gin{
		engine: engine,
	}

}

// Default function creates a Gin engine with the default middleware
// and returns a pointer to a Gin type.
func Default() *Gin {
	engine := gin.Default()
	return New(engine)
}

// Run runs the Gin engine with the provided addresses.
func (g Gin) Run(addr ...string) error {
	return g.engine.Run(addr...)
}

// handleHandles takes a list of HandlerFuncs and returns a single
// func(c *gin.Context) that calls each HandlerFunc in the list in sequence.
func handleHandles(handles []HandlerFunc) func(c *gin.Context) {
	return func(c *gin.Context) {
		for _, handle := range handles {
			handle(c)
		}
	}
}

// Group creates a new router group with the provided relativePath
// and handles as middlewares. It returns a RouteHandler type.
func (g Gin) Group(relativePath string, handles ...HandlerFunc) RouteHandler {
	routerGroup := g.engine.Group(relativePath, handleHandles(handles))
	return &GinRouteHandler{routerGroup: routerGroup}
}

// GinRouteHandler represents a Gin Router Group, and
// implements the RouteHandler interface.
type GinRouteHandler struct {
	routerGroup *gin.RouterGroup
}

// Group creates a new child router group with the provided relativePath and
// handles as middlewares. It returns a RouteHandler type.
func (grh GinRouteHandler) Group(relativePath string, handles ...HandlerFunc) RouteHandler {
	routerGroup := grh.routerGroup.Group(relativePath, handleHandles(handles))
	return &GinRouteHandler{routerGroup: routerGroup}
}

// BasePath returns the base path of the Gin Router Group.
func (grh GinRouteHandler) BasePath() string {
	return grh.routerGroup.BasePath()
}

// Use adds the provided list of HandlerFunc as middlewares to the Gin Router Group.
// It returns the RouteHandler type.
func (grh GinRouteHandler) Use(middlewares ...HandlerFunc) RouteHandler {
	grh.routerGroup.Use(handleHandles(middlewares))
	return grh
}

// Handle creates a new route with the provided httpMethod, relativePath, and
// handles. It returns the RouteHandler type.
func (grh GinRouteHandler) Handle(
	httpMethod string,
	relativePath string,
	handles ...HandlerFunc,
) RouteHandler {
	grh.routerGroup.Handle(httpMethod, relativePath, handleHandles(handles))
	return grh
}

// GET creates a new GET route with the provided relativePath and
// handles. It returns the RouteHandler type.
func (grh GinRouteHandler) GET(relativePath string, handles ...HandlerFunc) RouteHandler {
	grh.routerGroup.GET(relativePath, handleHandles(handles))
	return grh
}

// POST creates a new POST route with the provided relativePath and
// handles. It returns the RouteHandler type.
func (grh GinRouteHandler) POST(relativePath string, handles ...HandlerFunc) RouteHandler {
	grh.routerGroup.POST(relativePath, handleHandles(handles))
	return grh
}

// DELETE creates a new DELETE route with the provided relativePath and
// handles. It returns the RouteHandler type.
func (grh GinRouteHandler) DELETE(relativePath string, handles ...HandlerFunc) RouteHandler {
	grh.routerGroup.DELETE(relativePath, handleHandles(handles))
	return grh
}

// PATCH creates a new PATCH route with the provided relativePath and
// handles. It returns the RouteHandler type.
func (grh GinRouteHandler) PATCH(relativePath string, handles ...HandlerFunc) RouteHandler {
	grh.routerGroup.PATCH(relativePath, handleHandles(handles))
	return grh
}

// PUT creates a new PUT route with the provided relativePath and
// handles. It returns the RouteHandler type.
func (grh GinRouteHandler) PUT(relativePath string, handles ...HandlerFunc) RouteHandler {
	grh.routerGroup.PUT(relativePath, handleHandles(handles))
	return grh
}

// OPTIONS creates a new OPTIONS route with the provided relativePath and
// handles. It returns the RouteHandler type.
func (grh GinRouteHandler) OPTIONS(relativePath string, handles ...HandlerFunc) RouteHandler {
	grh.routerGroup.OPTIONS(relativePath, handleHandles(handles))
	return grh
}

// HEAD creates a new HEAD route with the provided relativePath and
// handles. It returns the RouteHandler type.
func (grh GinRouteHandler) HEAD(relativePath string, handles ...HandlerFunc) RouteHandler {
	grh.routerGroup.HEAD(relativePath, handleHandles(handles))
	return grh
}
