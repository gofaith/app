package app

type DivWidget struct {
	Element
}

func Div() *DivWidget {
	v := &DivWidget{}
	v.name = "div"
	return v
}

func Column(widgets ...Widget) *DivWidget {
	v := Div()
	v.ColumnCenter().Body(widgets...)
	return v
}

func Row(widgets ...Widget) *DivWidget {
	v := Div()
	v.RowCenter().Body(widgets...)
	return v
}

