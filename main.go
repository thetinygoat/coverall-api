package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func main() {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	r.Get("/top", func(w http.ResponseWriter, r *http.Request) {
		res, err := http.Get("https://newsapi.org/v2/top-headlines?sources=the-hindu&sortBy=popularity&apiKey=" + os.Getenv("API_KEY"))
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		w.Write(body)
	})
	r.Get("/search", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("q")
		res, err := http.Get("https://newsapi.org/v2/everything?q=" + query + "&sources=the-hindu&sortBy=popularity&apiKey=" + os.Getenv("API_KEY"))
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		w.Write(body)
	})
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), r))
}
