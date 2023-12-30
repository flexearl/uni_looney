package routes

import (
	"net/http"

	"github.com/flexearl/uni_looney.git/middleware"
	"github.com/flexearl/uni_looney.git/token"
)

func HandleRoutes() {

	http.Handle("/add_post", token.VerifyJWT(http.HandlerFunc(middleware.AddPost)))
	http.HandleFunc("/register", middleware.Register)
	http.HandleFunc("/login", middleware.Login)

}
