package main

import (
	"log"
	"mini/app/gin-server/handler"
	"mini/app/gin-server/router"
	"mini/repository"
	"mini/service"
	"mini/utils/config"
	"mini/utils/database"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	db, err := database.GetDBConnection(cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresHost, cfg.PostgresDb)
	if err != nil {
		log.Fatalf("Error trying to connect to db: %v", err.Error())
	}

	transactionRepo := repository.NewTransactionRepository(db)
	walletRepo := repository.NewWalletRepository(db)
	walletSvc := service.NewWalletService(db, walletRepo, transactionRepo)
	walletHandler := handler.NewWalletHandler(walletSvc)

	app := gin.Default()

	router.Router(app, walletHandler)

	log.Println("Connected to the server")
	app.Run(":8000")
}
