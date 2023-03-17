package app

type htmlWidget struct {
	Element
}

func html() *htmlWidget {
	v := &htmlWidget{}
	v.name = "html"
	v.AddAttr("lang", "en")
	v.children = append(v.children, Head())
	return v
}

func (h *htmlWidget) Head(head *HeadWidget) *htmlWidget {
	if len(h.children) == 0 {
		h.children = append(h.children, head)
		return h
	}
	h.children[0] = head
	return h
}

func (h *htmlWidget) Body(body *BodyWidget) *htmlWidget {
	if len(h.children) == 0 {
		h.children = append(h.children, Head())
		return h
	}
	if len(h.children) == 1 {
		h.children = append(h.children, body)
		return h
	}
	h.children[1] = body
	return h
}
