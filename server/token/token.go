package token

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var sampleSecretKey = []byte("Secret")

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateJWT(username string) (string, error) {
	expirationTime := time.Now().Add(10 * time.Minute)
	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(sampleSecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := w.Header()
		header.Add("Access-Control-Allow-Origin", "*")
		header.Add("Access-Control-Allow-Methods", "DELETE, POST, GET, OPTIONS")
		header.Add("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
		fmt.Println("Request", r.Header)
		fmt.Println("Authorisation", r.Header.Get("Authorization"))
		reqToken := r.Header.Get("Authorization")
		fmt.Println("Checking verification")
		fmt.Println("Request token", reqToken)
		splitToken := strings.Split(reqToken, "Bearer ")
		fmt.Println(splitToken)
		if len(splitToken) > 1 {
			reqToken = splitToken[1]
		}

		claims := &Claims{}
		if reqToken != "" {
			fmt.Println("Checking")
			/*
				token, err := jwt.Parse(reqToken, func(token *jwt.Token) (interface{}, error) {
					_, ok := token.Method.(*jwt.SigningMethodECDSA)
					if !ok {
						w.WriteHeader(http.StatusUnauthorized)
						_, err := w.Write([]byte("You're Unauthorised!"))
						if err != nil {
							return nil, err
						}
					}
					fmt.Println("Sending blank")
					return []byte(""), nil
				})
			*/
			token, err := jwt.ParseWithClaims(reqToken, claims, func(token *jwt.Token) (any, error) {
				return []byte(sampleSecretKey), nil
			})
			//If error is returned then the signature verification has failed
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Println(err)
				_, err2 := w.Write([]byte("You're Unauthorised due to error parsing the JWT"))
				if err2 != nil {

					return
				}
			}
			if token.Valid {
				fmt.Println("Serving")
				next.ServeHTTP(w, r)
			} else {
				//w.WriteHeader(http.StatusUnauthorized)
				_, err := w.Write([]byte("You're unauthorised due to no token in the header"))
				if err != nil {
					return
				}
			}

		}
	})

}
