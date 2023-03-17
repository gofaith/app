package app

type TitleWidget struct {
	Element
}

func Title(s string) *TitleWidget {
	v := &TitleWidget{}
	v.name = "title"
	v.Text(s)
	return v
}
