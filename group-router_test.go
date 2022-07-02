package go_gin_utils

import (
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGroupRouter_DELETE(t *testing.T) {
	type fields struct {
		groups []*gin.RouterGroup
	}
	type args struct {
		path    string
		handler []gin.HandlerFunc
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GroupRouter{
				groups: tt.fields.groups,
			}
			g.DELETE(tt.args.path, tt.args.handler...)
		})
	}
}

func TestGroupRouter_GET(t *testing.T) {
	type fields struct {
		groups []*gin.RouterGroup
	}
	type args struct {
		path    string
		handler []gin.HandlerFunc
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GroupRouter{
				groups: tt.fields.groups,
			}
			g.GET(tt.args.path, tt.args.handler...)
		})
	}
}

func TestGroupRouter_Handle(t *testing.T) {
	type fields struct {
		groups []*gin.RouterGroup
	}
	type args struct {
		method  string
		path    string
		handler []gin.HandlerFunc
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GroupRouter{
				groups: tt.fields.groups,
			}
			g.Handle(tt.args.method, tt.args.path, tt.args.handler...)
		})
	}
}

func TestGroupRouter_PATCH(t *testing.T) {
	type fields struct {
		groups []*gin.RouterGroup
	}
	type args struct {
		path    string
		handler []gin.HandlerFunc
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GroupRouter{
				groups: tt.fields.groups,
			}
			g.PATCH(tt.args.path, tt.args.handler...)
		})
	}
}

func TestGroupRouter_POST(t *testing.T) {
	type fields struct {
		groups []*gin.RouterGroup
	}
	type args struct {
		path    string
		handler []gin.HandlerFunc
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GroupRouter{
				groups: tt.fields.groups,
			}
			g.POST(tt.args.path, tt.args.handler...)
		})
	}
}

func TestGroupRouter_PUT(t *testing.T) {
	type fields struct {
		groups []*gin.RouterGroup
	}
	type args struct {
		path    string
		handler []gin.HandlerFunc
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GroupRouter{
				groups: tt.fields.groups,
			}
			g.PUT(tt.args.path, tt.args.handler...)
		})
	}
}

func TestNewGroupRouter(t *testing.T) {
	type args struct {
		groups []*gin.RouterGroup
	}
	tests := []struct {
		name string
		args args
		want GroupRouter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGroupRouter(tt.args.groups...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGroupRouter() = %v, want %v", got, tt.want)
			}
		})
	}
}
