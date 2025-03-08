package infrastructure

import (
    "github.com/gin-gonic/gin"
)

type CardRoutes struct {
    ValidateCardController *CardController
}

func NewCardRoutes(validateCardController *CardController) *CardRoutes {
    return &CardRoutes{ValidateCardController: validateCardController}
}

func (r *CardRoutes) AttachRoutes(router *gin.Engine) {
    cardsGroup := router.Group("/cards")
    {
        cardsGroup.POST("/validate", func(c *gin.Context) {
            r.ValidateCardController.ValidateCard(c.Writer, c.Request)
        })
    }
}