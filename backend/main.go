package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Service struct {
	SongRepo SongRepo
}

type Song struct {
	Author string `json:"author"`
	URL    string `json:"url"`
	Image  string `json:"image"`
}
type SongRepo interface {
	GetSongs() ([]Song, error)
}

type MockSongRepo struct{}

func (m *MockSongRepo) GetSongs() ([]Song, error) {
	songs := []Song{
		{Author: "Author1", URL: "url1", Image: "image1"},
		{Author: "Author2", URL: "url2", Image: "image2"},
		{Author: "Author3", URL: "url3", Image: "image3"},
	}
	return songs, nil
}

func getSongsHandler(repo SongRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		songs, err := repo.GetSongs()
		if err != nil {
			http.Error(w, "Failed to get songs", http.StatusInternalServerError)
			return
		}

		// Convert the songs slice to JSON
		songsJSON, err := json.Marshal(songs)
		if err != nil {
			http.Error(w, "Failed to marshal songs to JSON", http.StatusInternalServerError)
			return
		}

		// Set the response Content-Type header
		w.Header().Set("Content-Type", "application/json")

		// Write the JSON response
		w.WriteHeader(http.StatusOK)
		w.Write(songsJSON)
	}
}

func newRest() *mux.Router {
	songRepo := &MockSongRepo{}
	r := mux.NewRouter()
	r.HandleFunc("/", getSongsHandler(songRepo)).Methods("GET")
	return r
}

func main() {

	router := newRest()

	if err := http.ListenAndServe(":8080", router); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
