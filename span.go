package app

type SpanWidget struct {
	Element
}

func Span() *SpanWidget {
	v := &SpanWidget{}
	v.name = "span"
	return v
}

func Text(s string) *SpanWidget {
	v := Span()
	v.Text(s)
	return v
}
