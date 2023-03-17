package app

import (
	"fmt"
	"strings"
)

type Element struct {
	name     string
	class    []string
	style    []attr
	attrs    []attr
	children []Widget
	nobody   bool // 没有body的元素，如<input>
}

type attr struct {
	key, value string
}

type InnerText string

func (t InnerText) Render() string {
	return string(t)
}

// EqualString key='value'
func (a *attr) EqualString() string {
	return a.key + "=" + quoteString(a.value)
}

// ColonString key:value
func (a *attr) ColonString() string {
	return a.key + ":" + a.value
}

func (e *Element) Render() string {
	buf := new(strings.Builder)
	buf.WriteString("<" + e.name + " ")
	if len(e.class) > 0 {
		buf.WriteString("class='" + strings.Join(e.class, " ") + "' ")
	}

	//style
	if len(e.style) > 0 {
		ss := []string{}
		for _, v := range e.style {
			ss = append(ss, v.ColonString())
		}
		buf.WriteString("style='" + strings.Join(ss, ";") + "' ")
	}

	//attr
	for _, v := range e.attrs {
		buf.WriteString(v.EqualString() + " ")
	}

	buf.WriteString(">")

	if !e.nobody {
		//body
		for _, v := range e.children {
			buf.WriteString(v.Render())
		}
		buf.WriteString("</" + e.name + ">")
	}
	return buf.String()
}

func (e *Element) getAttr(key string) string {
	for _, v := range e.attrs {
		if v.key == key {
			return v.value
		}
	}
	return ""
}

func (e *Element) evalGetElementById() string {
	return "document.getElementById('" + e.GetID() + "')"
}

func (e *Element) Text(s string) *Element {
	if instance.IsRunning() {
		instance.w.Eval(e.evalGetElementById() + ".innerText=" + quoteString(s))
		return e
	}

	e.children = []Widget{InnerText(s)}
	return e
}

func (e *Element) HTML(s string) *Element {
	if instance.IsRunning() {
		instance.w.Eval(e.evalGetElementById() + ".innerHTML=" + quoteString(s))
		return e
	}

	e.children = []Widget{InnerText(s)}
	return e
}

func (e *Element) Body(widgets ...Widget) *Element {
	e.children = widgets
	return e
}

func (e *Element) addStyle(key, value string) {
	for i, v := range e.style {
		if v.key == key {
			e.style[i].value = value
			return
		}
	}
	e.style = append(e.style, attr{key: key, value: value})
}

func (e *Element) addClass(classname string) {
	for _, v := range e.class {
		if v == classname {
			return
		}
	}
	e.class = append(e.class, classname)
}

func (e *Element) addAttr(key, value string) {
	for i, v := range e.attrs {
		if v.key == key {
			e.attrs[i].value = value
			return
		}
	}
	e.attrs = append(e.attrs, attr{key: key, value: value})
}

func (e *Element) Column() *Element {
	e.addStyle("display", "flex")
	e.addStyle("flex-direction", "column")
	return e
}
func (e *Element) ColumnCenter() *Element {
	e.addStyle("align-items", "center")
	return e.Column()
}

func (e *Element) WideColumn() *Element {
	e.addStyle("width", "100%")
	return e.Column()
}
func (e *Element) MaxWidthPixel(width int) *Element {
	e.addStyle("max-width", fmt.Sprintf("%dpx", width))
	return e
}
func (e *Element) Row() *Element {
	e.addStyle("display", "flex")
	e.addStyle("flex-direction", "row")
	return e
}
func (e *Element) RowCenter() *Element {
	e.addStyle("align-items", "center")
	return e.Row()
}

func (e *Element) GetID() string {
	for _, v := range e.attrs {
		if v.key == "id" {
			return v.value
		}
	}
	id := GenerateID()
	e.addAttr("id", id)
	return id
}

func (e *Element) AssignElem(elem **Element) *Element {
	e.GetID()
	*elem = e
	return e
}
