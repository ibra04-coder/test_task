package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"music-library/db"
	"music-library/models"
	"net/http"
	"strings"
)

type SongDetail struct {
    ReleaseDate string `json:"releaseDate"`
    Text        string `json:"text"`
    Link        string `json:"link"`
}

func GetSongs(page, size int) ([]models.Song, error) {
    var songs []models.Song
    if err := db.DB.Offset((page - 1) * size).Limit(size).Find(&songs).Error; err != nil {
        return nil, err
    }
    return songs, nil
}

func GetSongText(id, page, size int) (string, error) {
    var song models.Song
    if err := db.DB.First(&song, id).Error; err != nil {
        return "", errors.New("song not found")
    }

    verses := splitSongIntoVerses(song.Text)
    if page < 1 || page > len(verses)/size {
        return "", errors.New("invalid page number")
    }

    start := (page - 1) * size
    end := start + size
    if end > len(verses) {
        end = len(verses)
    }
    return joinVerses(verses[start:end]), nil
}

func AddSong(song *models.Song) error {
    details, err := fetchSongDetails(song.Group, song.Song)
    if err != nil {
        return err
    }

    song.ReleaseDate = details.ReleaseDate
    song.Text = details.Text
    song.Link = details.Link

    return db.DB.Create(song).Error
}

func UpdateSong(id int, song *models.Song) error {
    var existingSong models.Song
    if err := db.DB.First(&existingSong, id).Error; err != nil {
        return errors.New("song not found")
    }

    existingSong.Group = song.Group
    existingSong.Song = song.Song
    return db.DB.Save(&existingSong).Error
}

func DeleteSong(id int) error {
    if err := db.DB.Delete(&models.Song{}, id).Error; err != nil {
        return errors.New("song not found")
    }
    return nil
}

func fetchSongDetails(group, song string) (*SongDetail, error) {
    url := fmt.Sprintf("https://api.example.com/info?group=%s&song=%s", group, song)
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("failed to fetch song details")
    }

    var details SongDetail
    if err := json.NewDecoder(resp.Body).Decode(&details); err != nil {
        return nil, err
    }

    return &details, nil
}

func splitSongIntoVerses(text string) []string {
    return strings.Split(text, "\n\n") 
}

func joinVerses(verses []string) string {
    return strings.Join(verses, "\n\n") 
}
