package main

import (
	"log"
	"net/http"
	"github.com/sammatcha/urlshortener/db"
	"github.com/sammatcha/urlshortener/pkg/handlers"
	"github.com/sammatcha/urlshortener/pkg/services"
)

func main(){
	database,err := db.ConnectDB()
	if err != nil {
		log.Fatal("failed to connect to db", err)
	}
		//initialize services
	urlService := service.NewURL(database)

		//initialize handler wih services
	handler := &handlers.Handler{
		URLService: urlService,
	}
	//routes
	http.HandleFunc("/shorten", handler.ShortenURLHandler)
	http.HandleFunc("/", handler.RedirectHandler)

	// log.Println("Server is running at http://localhost:8080")

		//starting server and listening 
	log.Fatal(http.ListenAndServe(":8080", nil))
}

