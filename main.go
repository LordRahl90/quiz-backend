package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/LordRahl90/little_quiz_backend/src/app"
	"github.com/LordRahl90/little_quiz_backend/src/controllers"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.Use(app.JwtAuthentication)

	router.HandleFunc("/api/user/register", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Login).Methods("POST")
	router.HandleFunc("/api/test/initiate", controllers.InitiateTest).Methods("POST")
	router.HandleFunc("/api/test/complete", controllers.MarkTest).Methods("POST")

	port := 5000
	log.Println("Server Started Successfully...")
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))
	// err := http.ListenAndServe(":"+strconv.Itoa(port), handlers.CORS(handlers.AllowedHeaders(), handlers.AllowedMethods(), handlers.AllowedOrigins()) )
	// if err != nil {
	// 	fmt.Print(err)
	// }

}
