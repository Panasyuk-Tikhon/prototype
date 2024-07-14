package router

import (
	"net/http"
	"prototype/handlers"
)

type ServerHandler struct{}

func (m *ServerHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	handlers.HandleFiles(res, req)

}
