package server

import (
    "github.com/gin-gonic/gin"
    "datos/src/internal/infrastructure"
    "github.com/gin-contrib/cors"
)

func Run(cardRoutes *infrastructure.CardRoutes) {
    router := gin.Default()

    router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"*"}, 
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Content-Type", "Authorization"},
        AllowCredentials: true,
    }))

    cardRoutes.AttachRoutes(router) 

    router.Run(":8080")
}
