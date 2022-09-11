package server

import (
	"go-blog/router"
	"log"
	"net/http"
)
var App = &Server{}
type Server struct {

}

func (*Server) Start(ip, port string) {
	server := http.Server{
		Addr: ip+":"+port,
	}

	router.Router()

	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}