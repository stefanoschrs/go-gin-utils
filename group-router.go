package go_gin_utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GroupRouter struct {
	groups []*gin.RouterGroup
}

func (g *GroupRouter) GET(path string, handler ...gin.HandlerFunc) {
	g.Handle(http.MethodGet, path, handler...)
}

func (g *GroupRouter) POST(path string, handler ...gin.HandlerFunc) {
	g.Handle(http.MethodPost, path, handler...)
}

func (g *GroupRouter) PUT(path string, handler ...gin.HandlerFunc) {
	g.Handle(http.MethodPut, path, handler...)
}

func (g *GroupRouter) PATCH(path string, handler ...gin.HandlerFunc) {
	g.Handle(http.MethodPatch, path, handler...)
}

func (g *GroupRouter) DELETE(path string, handler ...gin.HandlerFunc) {
	g.Handle(http.MethodDelete, path, handler...)
}

// Handle
func (g *GroupRouter) Handle(method string, path string, handler ...gin.HandlerFunc) {
	for _, group := range g.groups {
		group.Handle(method, path, handler...)
	}
}

// NewGroupRouter New Group Router
func NewGroupRouter(groups ...*gin.RouterGroup) GroupRouter {
	return GroupRouter{
		groups,
	}
}
