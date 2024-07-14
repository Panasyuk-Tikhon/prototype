package handlers

import (
	"net/http"
)

func HandleFiles(res http.ResponseWriter, req *http.Request) {
	filePath := "statics" + req.URL.Path
	if req.URL.Path == "/" {
		filePath = "statics/index.html"
	}
	http.ServeFile(res, req, filePath)
}
