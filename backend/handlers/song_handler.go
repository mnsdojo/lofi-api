package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mnsdojo/lofi-api/backend/repository"
)

type SongHandler struct {
	songRepo repository.SongRepository
}

func NewSongHandler(songRepo repository.SongRepository) *SongHandler {
	return &SongHandler{songRepo: songRepo}
}

func (h *SongHandler) GetSongs(w http.ResponseWriter, r *http.Request) {
	songs, err := h.songRepo.GetSongs(r.Context())
	if err != nil {
		http.Error(w, "Failed to fetch songs", http.StatusInternalServerError)
		return
	}
	toJson(w, http.StatusOK, songs)

}

func toJson(w http.ResponseWriter, statuscode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statuscode)
	json.NewEncoder(w).Encode(data)
}
