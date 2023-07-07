module github.com/gofaith/app

go 1.20

require (
	github.com/StevenZack/openurl v0.0.0-20190430065139-b25363f65ff8
	github.com/gofaith/webview v0.1.0
	github.com/gorilla/websocket v1.5.0
)

require github.com/StevenZack/tools v1.14.1

replace github.com/gofaith/webview => ../webview
