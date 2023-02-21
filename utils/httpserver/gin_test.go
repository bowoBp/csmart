package httpserver

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGin_Group(t *testing.T) {
	handles := []HandlerFunc{
		func(ctx Context) {},
	}

	t.Run("success", func(t *testing.T) {
		g := Default()
		if got := g.Group("/", handles...); got == nil {
			t.Errorf("Group() = %v", got)
		}
	})
}

func TestGinRouteHandler_Group(t *testing.T) {
	handles := []HandlerFunc{
		func(ctx Context) {},
	}

	t.Run("success", func(t *testing.T) {
		grh := &GinRouteHandler{routerGroup: &gin.RouterGroup{}}
		if got := grh.Group("/", handles...); got == nil {
			t.Errorf("Group() = %v", got)
		}
	})
}

func TestGinRouteHandler_BasePath(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		grh := &GinRouteHandler{routerGroup: &gin.RouterGroup{}}
		if got := grh.BasePath(); got != "" {
			t.Errorf("BasePath() = %v", got)
		}
	})
}

func TestGinRouteHandler_Use(t *testing.T) {
	handles := []HandlerFunc{
		func(ctx Context) {},
	}

	t.Run("success", func(t *testing.T) {
		grh := &GinRouteHandler{routerGroup: &gin.RouterGroup{}}
		if got := grh.Use(handles...); got == nil {
			t.Errorf("Use() = %v", got)
		}
	})
}

func TestGin_Run(t *testing.T) {
	t.Run("run", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("not panic")
			}
		}()
		g := Gin{}
		_ = g.Run()
	})
}

func Test_handleHandles(t *testing.T) {
	t.Run("handle", func(t *testing.T) {
		handle := handleHandles([]HandlerFunc{
			func(ctx Context) {},
		})
		handle(nil)
	})
}

func TestGinRouteHandler_Handle(t *testing.T) {
	t.Run("Handle", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("not panic")
			}
		}()
		grh := GinRouteHandler{}
		grh.Handle("", "", nil)
	})
}

func TestGinRouteHandler_GET(t *testing.T) {
	t.Run("Handle", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("not panic")
			}
		}()
		grh := GinRouteHandler{}
		grh.GET("", nil)
	})
}

func TestGinRouteHandler_POST(t *testing.T) {
	t.Run("Handle", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("not panic")
			}
		}()
		grh := GinRouteHandler{}
		grh.POST("", nil)
	})
}

func TestGinRouteHandler_DELETE(t *testing.T) {
	t.Run("Handle", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("not panic")
			}
		}()
		grh := GinRouteHandler{}
		grh.DELETE("", nil)
	})
}

func TestGinRouteHandler_PATCH(t *testing.T) {
	t.Run("Handle", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("not panic")
			}
		}()
		grh := GinRouteHandler{}
		grh.PATCH("", nil)
	})
}

func TestGinRouteHandler_PUT(t *testing.T) {
	t.Run("Handle", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("not panic")
			}
		}()
		grh := GinRouteHandler{}
		grh.PUT("", nil)
	})
}

func TestGinRouteHandler_OPTIONS(t *testing.T) {
	t.Run("Handle", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("not panic")
			}
		}()
		grh := GinRouteHandler{}
		grh.OPTIONS("", nil)
	})
}

func TestGinRouteHandler_HEAD(t *testing.T) {
	t.Run("Handle", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("not panic")
			}
		}()
		grh := GinRouteHandler{}
		grh.HEAD("", nil)
	})
}
