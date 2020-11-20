package main

import (
	"github.com/julienschmidt/httprouter"
	"golang-web-dev/042_mongodb/05_mongodb/01_update-user-controller/controllers"
	"gopkg.in/mgo.v2"
	"net/http"
)

func main() {
	//ルーターインスタンス生成
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)

}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	return s
}