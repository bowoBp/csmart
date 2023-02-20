package httpserver

type HttpServer interface {
	Run(addr ...string) error
	GroupHandler
}

type GroupHandler interface {
	Group(relativePath string, handles ...HandlerFunc) RouteHandler
}

type RouteHandler interface {
	GroupHandler
	BasePath() string
	Use(middleware ...HandlerFunc) RouteHandler
	Handle(httpMethod string, relativePath string, handles ...HandlerFunc) RouteHandler
	GET(relativePath string, handles ...HandlerFunc) RouteHandler
	POST(relativePath string, handles ...HandlerFunc) RouteHandler
	DELETE(relativePath string, handles ...HandlerFunc) RouteHandler
	PATCH(relativePath string, handles ...HandlerFunc) RouteHandler
	PUT(relativePath string, handles ...HandlerFunc) RouteHandler
	OPTIONS(relativePath string, handles ...HandlerFunc) RouteHandler
	HEAD(relativePath string, handles ...HandlerFunc) RouteHandler
}

type HandlerFunc func(ctx Context)

type Context interface {
	Next()
	Abort()
	AbortWithStatus(code int)
	AbortWithStatusJSON(code int, jsonObj any)
	Set(key string, value any)
	Get(key string) (value any, exists bool)
	Param(key string) string
	Query(key string) (value string)
	DefaultQuery(key, defaultValue string) string
	GetQuery(key string) (string, bool)
	Bind(obj any) error
	BindJSON(obj any) error
	BindQuery(obj any) error
	Status(code int)
	Header(key, value string)
	GetHeader(key string) string
	JSON(code int, obj any)
}
