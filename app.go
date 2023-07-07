package app

import (
	_ "embed"
	"sync"
)

type binder struct {
	key   string
	value interface{}
}

//go:embed bootstrap.min.js
var _bootstrapJS string

//go:embed bootstrap.min.css
var _bootstrapCSS string

var (
	idmap    = make(map[string]struct{})
	idwidth  = 5
	idLocker = sync.Mutex{}
)

func GenerateID() string {
	id := randomString(idwidth)
	idLocker.Lock()
	defer idLocker.Unlock()
	if _, ok := idmap[id]; ok {
		idwidth++
		return GenerateID()
	}
	idmap[id] = struct{}{}
	return id
}
