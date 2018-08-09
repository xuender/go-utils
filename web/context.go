package web

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

type Context struct {
	Writer  http.ResponseWriter
	Request *http.Request
	Render  *render.Render
	values  map[string]string
}

func (c *Context) Text(status int, v string) error {
	return c.Render.Text(c.Writer, status, v)
}
func (c *Context) String(status int, v string) error {
	return c.Render.Text(c.Writer, status, v)
}
func (c *Context) HTML(status int, name string, binding interface{}, htmlOpt ...render.HTMLOptions) error {
	return c.Render.HTML(c.Writer, status, name, binding, htmlOpt...)
}
func (c *Context) Data(status int, v []byte) error {
	return c.Render.Data(c.Writer, status, v)
}
func (c *Context) JSON(status int, v interface{}) error {
	return c.Render.JSON(c.Writer, status, v)
}
func (c *Context) JSONP(status int, callback string, v interface{}) error {
	return c.Render.JSONP(c.Writer, status, callback, v)
}
func (c *Context) XML(status int, v interface{}) error {
	return c.Render.XML(c.Writer, status, v)
}
func (c *Context) Get(key string) string {
	if c.values == nil {
		c.values = mux.Vars(c.Request)
	}
	return c.values[key]
}
