package main

import (
	"music-library/controllers"
	_ "music-library/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
    r := gin.Default()

    // Роут для отображения документации Swagger
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    // Остальные роуты
    r.GET("/songs", controllers.GetSongs)
    r.GET("/songs/:id/text", controllers.GetSongText)
    r.POST("/songs", controllers.AddSong)
    r.PUT("/songs/:id", controllers.UpdateSong)
    r.DELETE("/songs/:id", controllers.DeleteSong)

    r.Run()
}
