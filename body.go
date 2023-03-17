package app

type BodyWidget struct {
	Element
}

func Body(widgets ...Widget) *BodyWidget {
	v := &BodyWidget{}
	v.name = "body"
	v.children = widgets
	return v
}
