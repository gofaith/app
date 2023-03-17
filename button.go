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
	b.AddAttr("onclick", id+"()")
	return b
}

func PrimaryButton() *ButtonWidget {
	v := Button()
	v.AddClass("btn").AddClass("btn-primary")
	return v
}
