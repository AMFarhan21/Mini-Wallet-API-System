package handler

import (
	"context"
	"log"
	"mini/model"
	"net/http"
	"strconv"
	"strings"

	"github.com/AMFarhan21/fres"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type (
	WalletService interface {
		TopUp(ctx context.Context, userID int, amount int64) error
		Transfer(ctx context.Context, fromUserID, toUserID int, amount int64) error
		GetHistory(ctx context.Context, userID, page, limit int) ([]model.Transactions, error)
	}

	WalletHandler struct {
		walletService WalletService
		validate      *validator.Validate
	}

	TopUpRequest struct {
		UserID int   `json:"user_id" validate:"required"`
		Amount int64 `json:"amount" validate:"required"`
	}

	TransferRequest struct {
		FromUserID int   `json:"from_user_id" validate:"required"`
		ToUserID   int   `json:"to_user_id" validate:"required"`
		Amount     int64 `json:"amount" validate:"required"`
	}
)

func NewWalletHandler(walletService WalletService) *WalletHandler {
	return &WalletHandler{
		walletService: walletService,
		validate:      validator.New(),
	}
}

func (h *WalletHandler) TopUp(g *gin.Context) {
	var request TopUpRequest

	err := g.Bind(&request)
	if err != nil {
		log.Printf("Error on TopUp request body: %v", err.Error())
		g.JSON(http.StatusBadRequest, fres.Response.StatusBadRequest(err.Error()))
		return
	}

	err = h.validate.Struct(request)
	if err != nil {
		log.Printf("Error on TopUp request body: %v", err.Error())
		g.JSON(http.StatusBadRequest, fres.Response.StatusBadRequest(err.Error()))
		return
	}

	err = h.walletService.TopUp(g, request.UserID, request.Amount)
	if err != nil {
		if strings.Contains(err.Error(), "doesn't exists") {
			log.Printf("Error on TopUp: %v", err.Error())
			g.JSON(http.StatusNotFound, fres.Response.StatusNotFound(err.Error()))
			return
		}

		log.Printf("Error on TopUp internal: %v", err.Error())
		g.JSON(http.StatusInternalServerError, fres.Response.StatusInternalServerError("Error on TopUp internal"))
		return
	}

	g.JSON(http.StatusOK, fres.Response.StatusOK("top up success"))
}

func (h *WalletHandler) Transfer(g *gin.Context) {
	var request TransferRequest

	err := g.Bind(&request)
	if err != nil {
		log.Printf("Error on Transfer request body: %v", err.Error())
		g.JSON(http.StatusBadRequest, fres.Response.StatusBadRequest(err.Error()))
		return
	}

	err = h.validate.Struct(request)
	if err != nil {
		log.Printf("Error on Transfer request body: %v", err.Error())
		g.JSON(http.StatusBadRequest, fres.Response.StatusBadRequest(err.Error()))
		return
	}

	err = h.walletService.Transfer(g, request.FromUserID, request.ToUserID, request.Amount)
	if err != nil {
		if strings.Contains(err.Error(), "user not found") {
			log.Printf("Error on Transfer: %v", err.Error())
			g.JSON(http.StatusNotFound, fres.Response.StatusNotFound(err.Error()))
			return
		}

		if strings.Contains(err.Error(), "invalid") || strings.Contains(err.Error(), "insufficient") {
			log.Printf("Error on Transfer: %v", err.Error())
			g.JSON(http.StatusBadRequest, fres.Response.StatusBadRequest(err.Error()))
			return
		}

		log.Printf("Error on Transfer internal: %v", err.Error())
		g.JSON(http.StatusInternalServerError, fres.Response.StatusInternalServerError("Error on TopUp internal"))
		return
	}

	g.JSON(http.StatusOK, fres.Response.StatusOK("transfer success"))
}

func (h *WalletHandler) GetHistory(g *gin.Context) {
	userID, _ := strconv.Atoi(g.Param("user_id"))
	page, _ := strconv.Atoi(g.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(g.DefaultQuery("limit", "5"))

	history, err := h.walletService.GetHistory(g, userID, page, limit)
	if err != nil {
		log.Printf("Error on GetHistory internal: %v", err.Error())
		g.JSON(http.StatusInternalServerError, fres.Response.StatusInternalServerError("Error on GetHistory internal"))
		return
	}

	g.JSON(http.StatusOK, fres.Response.StatusOK(history))
}
