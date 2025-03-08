package server

import (
    "github.com/gin-gonic/gin"
    "datos/src/internal/infrastructure"
    "github.com/gin-contrib/cors"
)

func Run(cardRoutes *infrastructure.CardRoutes) {
    router := gin.Default()

    cardRoutes.AttachRoutes(router)

    router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, 
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	
	}))

    router.Run(":8080")
}