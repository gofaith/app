package app

import (
	"strings"

	"github.com/webview/webview"
)

type Window struct {
	title         string
	debug         bool
	width, height int
	hint          webview.Hint
	w             webview.WebView
	children      []Widget
	binders       []binder
}

var instance *Window

func NewWindow() *Window {
	if instance != nil {
		panic("You can only create one window")
	}
	instance = &Window{
		width:  600,
		height: 400,
		hint:   webview.HintNone,
	}
	return instance
}

func (w *Window) Debug(debug bool) *Window {
	w.debug = debug
	return w
}

func (w *Window) IsRunning() bool {
	return w.w != nil
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
	w.w = webview.New(w.debug)
	w.w.SetSize(w.width, w.height, w.hint)
	w.w.SetTitle(w.title)

	//render
	buf := new(strings.Builder)
	for _, v := range w.children {
		buf.WriteString(v.Render())
	}
	w.w.SetHtml(buf.String())

	//bind
	for _, v := range w.binders {
		e := w.w.Bind(v.key, v.value)
		if e != nil {
			panic(e)
		}
	}

	w.w.Run()
}

func (w *Window) Size(width, height int) *Window {
	w.width = width
	w.height = height
	return w
}

func (w *Window) SizeMax() *Window {
	w.hint = webview.HintMax
	w.Size(0, 0)
	return w
}
func (w *Window) SizeMin() *Window {
	w.hint = webview.HintMin
	w.Size(0, 0)
	return w
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

func (w *Window) Body(widgets ...Widget) *Window {
	w.children = widgets
	return w
}
