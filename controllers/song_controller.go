package controllers

import (
	"log"
	"music-library/models"
	"music-library/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Get all songs
// @Description Get list of songs with pagination and filtering by all fields
// @Tags songs
// @Accept  json
// @Produce  json
// @Param group query string false "Filter by group"
// @Param song query string false "Filter by song"
// @Param releaseDate query string false "Filter by release date"
// @Param page query int true "Page number"
// @Param size query int true "Page size"
// @Success 200 {array} models.Song
// @Failure 500 {object} gin.H{"error": "Internal Server Error"}
// @Router /songs [get]





func GetSongs(c *gin.Context) {
    page, _ := strconv.Atoi(c.Query("page"))
    size, _ := strconv.Atoi(c.Query("size"))
    songs, err := services.GetSongs(page, size)
    if err != nil {
        log.Println("Error fetching songs:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, songs)
}


// @Summary Get song text
// @Description Get paginated song text by song ID
// @Tags songs
// @Accept  json
// @Produce  json
// @Param id path int true "Song ID"
// @Param page query int true "Page number"
// @Param size query int true "Page size"
// @Success 200 {string} string
// @Failure 500 {object} gin.H{"error": "Internal Server Error"}
// @Router /songs/{id}/text [get]

func GetSongText(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    page, _ := strconv.Atoi(c.Query("page"))
    size, _ := strconv.Atoi(c.Query("size"))
    text, err := services.GetSongText(id, page, size)
    if err != nil {
        log.Println("Error fetching song text:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, text)
}

// @Summary Add a new song
// @Description Add a new song to the library, with enrichment from external API
// @Tags songs
// @Accept  json
// @Produce  json
// @Param song body models.Song true "New song data"
// @Success 201 {object} models.Song
// @Failure 400 {object} gin.H{"error": "Bad Request"}
// @Failure 500 {object} gin.H{"error": "Internal Server Error"}
// @Router /songs [post]

func AddSong(c *gin.Context) {
    var song models.Song
    if err := c.ShouldBindJSON(&song); err != nil {
        log.Println("Error binding JSON:", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    err := services.AddSong(&song)
    if err != nil {
        log.Println("Error adding song:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, song)
}

// @Summary Update an existing song
// @Description Update song data by song ID
// @Tags songs
// @Accept  json
// @Produce  json
// @Param id path int true "Song ID"
// @Param song body models.Song true "Updated song data"
// @Success 200 {object} models.Song
// @Failure 400 {object} gin.H{"error": "Bad Request"}
// @Failure 404 {object} gin.H{"error": "Song not found"}
// @Failure 500 {object} gin.H{"error": "Internal Server Error"}
// @Router /songs/{id} [put]

func UpdateSong(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var song models.Song
    if err := c.ShouldBindJSON(&song); err != nil {
        log.Println("Error binding JSON:", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    err := services.UpdateSong(id, &song)
    if err != nil {
        log.Println("Error updating song:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, song)
}

// @Summary Delete a song
// @Description Delete a song by song ID
// @Tags songs
// @Accept  json
// @Produce  json
// @Param id path int true "Song ID"
// @Success 204 "No Content"
// @Failure 404 {object} gin.H{"error": "Song not found"}
// @Failure 500 {object} gin.H{"error": "Internal Server Error"}
// @Router /songs/{id} [delete]

func DeleteSong(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    err := services.DeleteSong(id)
    if err != nil {
        log.Println("Error deleting song:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.Status(http.StatusNoContent)
}
