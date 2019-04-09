package main

import (
	"avnos/number1"
	"avnos/number2"
	"avnos/number3"
	"avnos/number4"
	"net/http"
	"avnos/api/service/generalservice"
	"github.com/julienschmidt/httprouter"
	"avnos/api/service/userservice"
	"encoding/json"
)
type Middleware struct {
	next http.Handler
	message string
}

// Our middleware handler
func (m *Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// We can modify the request here; for simplicity, we will just log a message
	if r.RequestURI == "/login" || r.RequestURI == "/register" {
		m.next.ServeHTTP(w, r)
	} else {
		if generalservice.TokenChecker(r.Header.Get("token")) {
			m.next.ServeHTTP(w, r)
		} else {
			json.NewEncoder(w).Encode(map[string]interface{}{"result":false, "msg": "Token Invalid"})
			return
		}
	}
}

func NewMiddleware(next http.Handler) *Middleware {
	return &Middleware{next: next}
}

func main() {
	number1.MaxBinaryGap()
	number2.Reverse()
	number3.Tree()
	number4.Pair()
	
	// Rest
	router := httprouter.New()
	router.POST("/login", generalservice.Login)
	router.POST("/register", generalservice.Register)
	router.POST("/logout", generalservice.Logout)
	router.GET("/user/list", userservice.ListUser)
	router.POST("/user/update/:key", userservice.UpdateUser)
	router.DELETE("/user/delete/:key", userservice.DeleteUser)
	m := NewMiddleware(router)
	http.ListenAndServe(":8181", m)
}