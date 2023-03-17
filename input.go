package app

type InputWidget struct {
	Element
}

func Input() *InputWidget {
	v := &InputWidget{}
	v.name = "input"
	v.nobody = true
	v.addAttr("type", "text")
	return v
}

func (i *InputWidget) Assign(v **InputWidget) *InputWidget {
	*v = i
	i.GetID()
	if i.getAttr("onkeyup") == "" {
		i.OnKeyUp(nil)
	}
	return i
}

func (i *InputWidget) TypeFile() *InputWidget {
	i.addAttr("type", "file")
	return i
}

func (i *InputWidget) Value(s string) *InputWidget {
	if instance.IsRunning() {
		instance.w.Eval(i.evalGetElementById() + ".value=" + quoteString(s))
		return i
	}
	i.addAttr("value", s)
	return i
}

func (i *InputWidget) GetValue() string {
	return i.getAttr("value")
}

func (i *InputWidget) OnChange(fn func(s string)) *InputWidget {
	value := i.getAttr("onchange")
	if value != "" {
		value = subBefore(value, "(", value)
		for idx, v := range instance.binders {
			if v.key == value {
				instance.binders[idx].value = fn
				return i
			}
		}
	}

	value = GenerateID()
	instance.addBinder(value, fn)
	i.addAttr("onchange", value+"(this.value)")
	return i
}

func (i *InputWidget) OnKeyUp(fn func(keyCode string)) *InputWidget {
	value := i.getAttr("onkeyup")
	if value != "" {
		value = subBefore(value, "(", value)
		for idx, v := range instance.binders {
			if v.key == value {
				instance.binders[idx].value = func(s string, code string) {
					i.addAttr("value", s)
					if fn != nil {
						fn(code)
					}
				}
				return i
			}
		}
	}
	value = GenerateID()
	instance.addBinder(value, func(s string, code string) {
		i.addAttr("value", s)
		if fn != nil {
			fn(code)
		}
	})
	i.addAttr("onkeyup", value+"(this.value,event.code)")
	return i
}
