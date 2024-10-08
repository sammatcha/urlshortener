package main

import (
	"log"
	"net/http"
	"github.com/sammatcha/urlshortener/db"
	"github.com/sammatcha/urlshortener/pkg/handlers"
	"github.com/sammatcha/urlshortener/pkg/services"
)
//Cors for frontend-backend
func EnableCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173") // Allow only React frontend
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions { // Handle preflight requests
            w.WriteHeader(http.StatusOK)
            return
        }

        next.ServeHTTP(w, r)
    })
}

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
	http.Handle("/shorten", EnableCors(http.HandlerFunc(handler.ShortenURLHandler)))
	http.Handle("/",EnableCors(http.HandlerFunc(handler.RedirectHandler)))

	// log.Println("Server is running at http://localhost:8080")

		//starting server and listening 
	log.Fatal(http.ListenAndServe(":8080", nil))
}

