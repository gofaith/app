package app

import (
	"bytes"
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func quoteString(s string) string {
	sep := "'"
	if strings.Contains(s, sep) {
		return "\"" + s + "\""
	}
	return sep + s + sep
}

func subBefore(s string, sep, def string) string {
	for i := 0; i <= len(s)-len(sep); i++ {
		if s[i:i+len(sep)] == sep {
			return s[:i]
		}
	}
	return def
}
func subBeforeLast(s, sep, def string) string {
	for i := len(s) - len(sep); i > -1; i-- {
		if s[i:i+len(sep)] == sep {
			return s[:i]
		}
	}
	return def
}

func subAfter(s, sep, def string) string {
	for i := 0; i <= len(s)-len(sep); i++ {
		if s[i:i+len(sep)] == sep {
			return s[i+len(sep):]
		}
	}
	return def
}

func subAfterLast(s, sep, def string) string {
	for i := len(s) - len(sep); i > -1; i-- {
		if s[i:i+len(sep)] == sep {
			return s[i+len(sep):]
		}
	}
	return def
}

func trimStarts(s string, trim string) string {
	for {
		if strings.HasPrefix(s, trim) {
			s = s[len(trim):]
			continue
		}
		break
	}
	return s
}

func trimEnds(s string, trim string) string {
	for {
		if strings.HasSuffix(s, trim) {
			s = s[:len(s)-len(trim)]
			continue
		}
		break
	}
	return s
}

func trimStart(s string, trim string) string {
	if strings.HasPrefix(s, trim) {
		return s[len(trim):]
	}
	return s
}

func trimEnd(s string, trim string) string {
	if strings.HasSuffix(s, trim) {
		return s[:len(s)-len(trim)]
	}
	return s
}

func trimBoth(s string, trims string) string {
	return trimStart(trimEnd(s, trims), trims)
}

func subBetween(s string, start, end rune) (string, error) {
	var buf *bytes.Buffer
	for _, r := range s {
		if r == start && buf == nil {
			buf = bytes.NewBufferString("")
			continue
		}
		if buf != nil {
			if r == end {
				return buf.String(), nil
			}
			buf.WriteRune(r)
		}
	}
	return "", errors.New("no end " + string(end))
}

var letterRunes = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func init() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("bootstrap:", len(_bootstrapCSS))
}

func randomString(n int) string {
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
