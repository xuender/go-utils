package web

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

type Web struct {
	Router  *mux.Router      // 路由
	Render  *render.Render   // 输出
	Negroni *negroni.Negroni // 服务
}

func Classic() *Web {
	return &Web{
		Router:  mux.NewRouter(),
		Render:  render.New(),
		Negroni: negroni.Classic(),
	}
}
func (w *Web) Handle(path string, f func(c *Context)) *mux.Route {
	return w.Router.HandleFunc(path, func(writer http.ResponseWriter, req *http.Request) {
		f(&Context{
			Writer:  writer,
			Request: req,
			Render:  w.Render,
		})
	})
}
func (w *Web) Get(path string, f func(c *Context)) *mux.Route {
	return w.Handle(path, f).Methods("GET")
}
func (w *Web) Post(path string, f func(c *Context)) *mux.Route {
	return w.Handle(path, f).Methods("POST")
}
func (w *Web) Put(path string, f func(c *Context)) *mux.Route {
	return w.Handle(path, f).Methods("PUT")
}
