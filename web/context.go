package web

import (
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

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

// 二进制输出
func (c *Context) Blob(status int, contextType string, v []byte) error {
	head := render.Head{
		ContentType: contextType,
		Status:      status,
	}
	d := render.Data{
		Head: head,
	}
	return c.Render.Render(c.Writer, d, v)
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

// 获取参数
func (c *Context) Param(key string) string {
	if c.values == nil {
		c.values = mux.Vars(c.Request)
	}
	return c.values[key]
}
func (c *Context) Get(key string) string {
	return c.Param(key)
}
func (c *Context) FormFile(name string) (*multipart.FileHeader, error) {
	_, fh, err := c.Request.FormFile(name)
	return fh, err
}

// 读取文件
func (c *Context) File(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	fi, _ := f.Stat()
	if fi.IsDir() {
		file = filepath.Join(file, "index.html")
		f, err = os.Open(file)
		if err != nil {
			return err
		}
		defer f.Close()
		if fi, err = f.Stat(); err != nil {
			return err
		}
	}
	http.ServeContent(c.Writer, c.Request, fi.Name(), fi.ModTime(), f)
	return nil
}
