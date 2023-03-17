package app

type ScriptWidget struct {
	Element
}

func Script() *ScriptWidget {
	v := &ScriptWidget{}
	v.name = "script"
	return v
}

