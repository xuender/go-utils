package web

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

type Web struct {
	Router  *mux.Router      // 路由
	Render  *render.Render   // 输出
	Negroni *negroni.Negroni // 服务
	Name    string           // 名称
	logger  *log.Logger      // 日志

}

func Classic(name string) *Web {
	m := mux.NewRouter()
	r := render.New()
	n := negroni.New(newRecovery(name), newLogger(name), negroni.NewStatic(http.Dir("www")))
	n.UseHandler(m)
	return &Web{
		Name:    name,
		Router:  m,
		Render:  r,
		Negroni: n,
	}
}
func newRecovery(name string) *negroni.Recovery {
	return &negroni.Recovery{
		Logger:     log.New(os.Stdout, "["+name+"] ", 0),
		PrintStack: true,
		StackAll:   false,
		StackSize:  1024 * 8,
		Formatter:  &negroni.TextPanicFormatter{},
	}
}
func newLogger(name string) *negroni.Logger {
	logger := &negroni.Logger{ALogger: log.New(os.Stdout, "["+name+"] ", 0)}
	logger.SetDateFormat(negroni.LoggerDefaultDateFormat)
	logger.SetFormat(negroni.LoggerDefaultFormat)
	return logger
}

// ListenAndServe run
func (w *Web) Run(addr string) {
	w.logger = log.New(os.Stdout, "["+w.Name+"] ", 0)
	w.logger.Printf("listening on %s", addr)
	w.logger.Fatal(http.ListenAndServe(addr, w.Negroni))
}
func (w *Web) Handle(path string, f func(c *Context) error) *mux.Route {
	return w.Router.HandleFunc(path, func(writer http.ResponseWriter, req *http.Request) {
		err := f(&Context{
			Writer:  writer,
			Request: req,
			Render:  w.Render,
		})
		if err != nil {
			w.logger.Println(err.Error())
		}
	})
}
func (w *Web) GET(path string, f func(c *Context) error) *mux.Route {
	return w.Handle(path, f).Methods("GET")
}
func (w *Web) POST(path string, f func(c *Context) error) *mux.Route {
	return w.Handle(path, f).Methods("POST")
}
func (w *Web) PUT(path string, f func(c *Context) error) *mux.Route {
	return w.Handle(path, f).Methods("PUT")
}
