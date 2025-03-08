package cmd

import (
    "log"
    "datos/src/core"
    "datos/src/server"
    cardDeps "datos/src/internal/infrastructure"
)

func Init() {
    db, err := core.ConnectDB()
    if err != nil {
        log.Fatal("Error al conectar a la base de datos:", err)
    }

    rabbitMQ, err := core.NewRabbitMQConnection()
    if err != nil {
        log.Fatal("Error al conectar a RabbitMQ:", err)
    }
    defer rabbitMQ.Close()

    cardDependencies := cardDeps.NewCardDependencies(db, rabbitMQ)
    cardRoutes := cardDependencies.GetRoutes()

    server.Run(cardRoutes)
}