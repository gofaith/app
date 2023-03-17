package app

type StyleWidget struct {
	Element
}

func Style(s string) *StyleWidget {
	v := &StyleWidget{}
	v.name = "style"
	v.Text(s)
	return v
}
