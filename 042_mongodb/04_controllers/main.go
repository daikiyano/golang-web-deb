
package main

import (
"github.com/julienschmidt/httprouter"
"golang-web-dev/042_mongodb/04_controllers/controllers"
"net/http"
)

func main() {
r := httprouter.New()
uc := controllers.NewUserController()
r.GET("/user/:id", uc.GetUser)
r.POST("/user", uc.CreateUser)
r.DELETE("/user/:id", uc.DeleteUser)
http.ListenAndServe("localhost:8080", r)
}