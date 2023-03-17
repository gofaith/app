package app

type MetaWidget struct {
	Element
}

func Meta() *MetaWidget {
	v := &MetaWidget{}
	v.name = "meta"
	v.enclosure = true
	return v
}
