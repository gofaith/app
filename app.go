package app

import (
	"math/rand"
	"sync"
	"time"
)

type binder struct {
	key   string
	value interface{}
}

var (
	idmap    = make(map[string]struct{})
	idwidth  = 5
	idLocker = sync.Mutex{}
)

var letterRunes = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomString(n int) string {
	if n < 1 {
		return ""
	}
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	if b[0] >= '0' && b[0] <= '9' {
		return "a" + string(b)
	}
	return string(b)
}

func GenerateID() string {
	id := RandomString(idwidth)
	idLocker.Lock()
	defer idLocker.Unlock()
	if _, ok := idmap[id]; ok {
		idwidth++
		return GenerateID()
	}
	idmap[id] = struct{}{}
	return id
}
