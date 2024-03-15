package repository

import (
	"context"

	"github.com/mnsdojo/lofi-api/backend/models"
)

type SongRepository interface {
	GetSongs(ctx context.Context) ([]models.Song, error)
}

type InMemorySongRepo struct {
}

func (r *InMemorySongRepo) GetSongs(ctx context.Context) ([]models.Song, error) {
	songs := []models.Song{
		{URL: "song1.mp3", Author: "Author 1", Image: "image1.jpg"},
		{URL: "song2.mp3", Author: "Author 2", Image: "image2.jpg"},
	}
	return songs, nil
}
