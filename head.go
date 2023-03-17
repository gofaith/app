package app

type HeadWidget struct {
	Element
}

func Head(widgets ...Widget) *HeadWidget {
	v := &HeadWidget{}
	v.name = "head"
	v.children = append(v.children, Meta().AddAttr("charset", "UTF-8"))
	v.children = append(v.children, Meta().AddAttr("http-equiv", "X-UA-Compatible").AddAttr("content", "IE=edge"))
	v.children = append(v.children, Meta().AddAttr("name", "viewport").AddAttr("content", "width=device-width, initial-scale=1.0"))
	v.children = append(v.children, widgets...)
	return v
}
