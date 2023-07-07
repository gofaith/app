package app

import (
	"strings"

	"github.com/StevenZack/openurl"
	"github.com/gofaith/webview"
)

type Window struct {
	title   string
	debug   bool
	server  *bridgeServer
	head    []Widget
	body    []Widget
	binders []binder
}

var instance *Window

func NewWindow() *Window {
	if instance != nil {
		panic("You can only create one window")
	}
	instance = &Window{}
	return instance
}

func (w *Window) Debug(debug bool) *Window {
	w.debug = debug
	return w
}
func (w *Window) IsRunning() bool {
	return w.server != nil
}
func (w *Window) addBinder(key string, value interface{}) {
	for i, v := range w.binders {
		if v.key == key {
			w.binders[i].value = value
			return
		}
	}
	w.binders = append(w.binders, binder{key: key, value: value})
}

func (w *Window) Run() {
	//compose with bootstrap
	doc := html().Head(
		Head(
			append(w.head, Style(_bootstrapCSS))...,
		),
	).Body(
		Body(append(w.body, Script().Text(_bootstrapJS))...),
	)

	// webview

	//render
	buf := new(strings.Builder)
	buf.WriteString("<!DOCTYPE html>")
	buf.WriteString(doc.Render())

	openurl.OpenApp("data:text/html," + buf.String())

	//bind
	for _, v := range w.binders {
		e := w.w.Bind(v.key, v.value)
		if e != nil {
			panic(e)
		}
	}

}

func (w *Window) SizeFixed() *Window {
	w.hint = webview.HintFixed
	return w
}
func (w *Window) Title(s string) *Window {
	w.title = s
	if w.w != nil {
		w.w.SetTitle(s)
	}
	return w
}

func (w *Window) Assign(v **Window) *Window {
	*v = w
	return w
}
func (w *Window) Head(widgets ...Widget) *Window {
	w.head = widgets
	return w
}
func (w *Window) Body(widgets ...Widget) *Window {
	w.body = widgets
	return w
}
