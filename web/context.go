package web

import (
	"net/http"

	"github.com/unrolled/render"
)

type Context struct {
	Writer  http.ResponseWriter
	Request *http.Request
	Render  *render.Render
}
