package app

type ButtonWidget struct {
	Element
}

func Button() *ButtonWidget {
	v := &ButtonWidget{}
	v.name = "button"
	return v
}

func (b *ButtonWidget) OnClick(fn func()) *ButtonWidget {
	id := GenerateID()
	instance.addBinder(id, fn)
	b.addAttr("onclick", id+"()")
	return b
}
