package router

import (
	"mini/app/gin-server/handler"

	"github.com/gin-gonic/gin"
)

func Router(app *gin.Engine, walletHandler *handler.WalletHandler) {
	app.GET("/ping", func(h *gin.Context) {
		h.JSON(200, "PONG")
	})

	api := app.Group("/api")

	wallet := api.Group("/wallet")
	wallet.POST("/topup", walletHandler.TopUp)
	wallet.POST("/transfer", walletHandler.Transfer)
	wallet.GET("/history/:user_id", walletHandler.GetHistory)
}
