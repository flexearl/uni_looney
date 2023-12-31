package middleware

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/flexearl/uni_looney.git/helper"
	"github.com/flexearl/uni_looney.git/post"
	"github.com/flexearl/uni_looney.git/token"
	"github.com/flexearl/uni_looney.git/user"
)

type Message struct {
	Status string `json:"status"`
	Info   string `json:"info"`
}

func GetPosts(w http.ResponseWriter, r *http.Request) {
	header := w.Header()
	fmt.Println("Getting posts")
	header.Add("Access-Control-Allow-Methods", "DELETE, POST, GET, OPTIONS")

	header.Add("Content-Type", "application/json")
	header.Add("Access-Control-Allow-Headers", "*")
	posts := []post.Post{{"dd", "fff", "fff", 5}, {"dd", "fff", "fff", 6}}
	a, _ := json.Marshal(posts)
	w.Write(a)
}

func AddPost(w http.ResponseWriter, r *http.Request) {
	header := w.Header()

	header.Add("Access-Control-Allow-Methods", "DELETE, POST, GET, OPTIONS")
	header.Add("Content-Type", "application/json")
	header.Add("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")

	fmt.Println("Adding post")
	var post post.Post
	helper.DecodeJSON(w, r, &post)
	fmt.Println(w.Header())
	w.Write([]byte("Successful!"))
}

func Register(w http.ResponseWriter, r *http.Request) {
	var user user.User
	helper.DecodeJSON(w, r, &user)
	//password, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
}

func Login(w http.ResponseWriter, r *http.Request) {

	header := w.Header()
	header.Add("Access-Control-Allow-Origin", "*")
	header.Add("Access-Control-Allow-Methods", "DELETE, POST, GET, OPTIONS")
	header.Add("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)

	}
	validLogin := true
	fmt.Println("login")
	var user user.User
	helper.DecodeJSON(w, r, &user)
	if validLogin {
		token, err := token.GenerateJWT(user.Username)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(token)
		fmt.Fprint(w, token)
		w.Header().Set("token", token)

		http.SetCookie(w, &http.Cookie{
			Name:     "token",
			Value:    token,
			Path:     "/",
			HttpOnly: true,
		})

	}

}
