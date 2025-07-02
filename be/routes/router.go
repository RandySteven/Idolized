package routes

import (
	"github.com/RandySteven/Idolized/enums"
	"github.com/RandySteven/Idolized/handlers/apis"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type (
	HandlerFunc func(w http.ResponseWriter, r *http.Request)

	Router struct {
		path        string
		handler     HandlerFunc
		method      string
		middlewares []enums.Middleware
	}

	RouterPrefix map[enums.RouterPrefix][]*Router
)

func NewEndpointRouters(api *apis.APIs) RouterPrefix {
	endpoint := make(RouterPrefix)

	return endpoint
}

func InitRouter(routers RouterPrefix, r *mux.Router) {
}

func (router *Router) RouterLog(prefix string) {
	log.Printf("%12s | %4s/ \n", router.method, prefix+router.path)
}
