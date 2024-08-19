package handlers

import (
	"fmt"
	"log"
	"net/http"
	"github.com/sammatcha/urlshortener/pkg/services"
	"github.com/sammatcha/urlshortener/pkg/utils"
)

type Handler struct {
	URLService *service.URLService
}
	//shorten url
func(h *Handler) ShortenURLHandler(w http.ResponseWriter, r *http.Request) {
	// log.Println("shorten url handler triggered")

		// Extract long URL from http request
	longURL:= r.FormValue("url")
	//validate url not empty
	if longURL == "" {
		log.Println("received empty url")
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

		//generate shortened key for long url
	shortURL, err :=utils.GenerateShortURL()
	if err != nil {
		// log.Println("error generating short url", err)
		http.Error(w ,"Error generating short URL", http.StatusInternalServerError)
		return 
	}

		//save url to db
	err = h.URLService.SaveURL(longURL, shortURL)
	if err!= nil {
		log.Println("error saving to db", http.StatusInternalServerError)
		return
	}
	// log.Println("URL stored:", shortURL, "->", longURL)
    fmt.Fprintf(w, "Generated Short URL: %s", shortURL)
}

	//redirect
func (h *Handler) RedirectHandler(w http.ResponseWriter, r *http.Request) {
	shortURL := r.URL.Path[1:]

	// log.Println("redirect triggered with short url", shortURL)

		//retrieve long url from db
	url, err := h.URLService.GetURL(shortURL)
	if err != nil{
		http.Error(w, "Error retrieving URL", http.StatusInternalServerError)
		return
	}
	if url == nil {
		log.Println("short url not found:", shortURL)
		http.NotFound(w,r)
		return
	} 
		
	// log.Println("Redirecting to:", url.LongURL)
	http.Redirect(w,r,url.LongURL, http.StatusFound)
}