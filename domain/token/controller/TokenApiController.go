package controller

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"manage-template/config"
	"manage-template/domain/token/service"
	"math/big"
	"net/http"
)

type TokenApiController struct {
}

var TokenApiService = new(service.TokenApiService)

func SetRouterToken(r *gin.RouterGroup) {

	tokenRouter := r.Group("/token")

	tokenApiController := new(TokenApiController)

	tokenRouter.GET("/transactionReceipt", tokenApiController.TransactionReceipt)
	tokenRouter.GET("/balanceToken", tokenApiController.BalanceToken)
	tokenRouter.POST("/transferToken", tokenApiController.TransferToken)
	tokenRouter.POST("/transferTokenPolling", tokenApiController.TransferTokenPolling)
	tokenRouter.POST("/transferTokenCancel", tokenApiController.TransferTokenCancel)

}

type RequestParam struct {
	From  string  `json:"From" example:"0x46C65D87bE47255882561bcc7CFf3bBA186F0848"`
	To    string  `json:"To" example:"0x5EFC0751759b01759BEc45a06930A2d338a5E234"`
	Value big.Int `json:"Value"`
}

// TransferTokenCancel godoc
// @Summary TransferTokenCancel
// @Description do TransferTokenCancel
// @name TransferTokenCancel
// @Accept  json
// @Produce  json
// @Param	RequestParam body RequestParam true "To Address, From Address, Value Token"
// @Router /token/transferTokenCancel [post]
// @Success 200 {object} config.ResponseObject
// @Failure 400
func (ctrl TokenApiController) TransferTokenCancel(c *gin.Context) {
	var requestJson RequestParam

	if err := json.NewDecoder(c.Request.Body).Decode(&requestJson); err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			config.Builder().
				IsSuccess_(false).
				Message_(err.Error()).
				Build(),
		)
		return
	}

	if requestJson.From == "" || requestJson.To == "" || requestJson.Value.String() == "0" {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			config.Builder().
				IsSuccess_(false).
				Message_("Request values do not match").
				Build(),
		)
		return
	}

	hash, err := TokenApiService.TransferTokenCancel(requestJson.From, requestJson.To, requestJson.Value)

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			config.Builder().
				IsSuccess_(false).
				Message_(err.Error()).
				Build(),
		)
		return
	}

	type Result struct {
		Hash common.Hash
	}

	c.JSON(http.StatusOK, config.Builder().
		IsSuccess_(true).
		Body_(Result{Hash: hash}).
		Build(),
	)
}

// TransferTokenPolling godoc
// @Summary TransferTokenPolling
// @Description do TransferTokenPolling
// @name TransferTokenPolling
// @Accept  json
// @Produce  json
// @Param	RequestParam body RequestParam true "To Address, From Address, Value Token"
// @Router /token/transferTokenPolling [post]
// @Success 200 {object} config.ResponseObject
// @Failure 400
func (ctrl TokenApiController) TransferTokenPolling(c *gin.Context) {
	var requestJson RequestParam

	if err := json.NewDecoder(c.Request.Body).Decode(&requestJson); err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			config.Builder().
				IsSuccess_(false).
				Message_(err.Error()).
				Build(),
		)
		return
	}

	if requestJson.From == "" || requestJson.To == "" || requestJson.Value.String() == "0" {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			config.Builder().
				IsSuccess_(false).
				Message_("Request values do not match").
				Build(),
		)
		return
	}

	hash, status, err := TokenApiService.TransferTokenPolling(requestJson.From, requestJson.To, requestJson.Value)

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			config.Builder().
				IsSuccess_(false).
				Message_(err.Error()).
				Build(),
		)
		return
	}

	type Result struct {
		Status uint64
		Hash   common.Hash
	}

	c.JSON(http.StatusOK, config.Builder().
		IsSuccess_(true).
		Body_(Result{Status: status, Hash: hash}).
		Build(),
	)
}

// TransferToken godoc
// @Summary TransferToken
// @Description do TransferToken
// @name TransferToken
// @Accept  json
// @Produce  json
// @Param	RequestParam body RequestParam true "To Address, From Address, Value Token"
// @Router /token/transferToken [post]
// @Success 200 {object} config.ResponseObject
// @Failure 400
func (ctrl TokenApiController) TransferToken(c *gin.Context) {
	var requestJson RequestParam

	if err := json.NewDecoder(c.Request.Body).Decode(&requestJson); err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			config.Builder().
				IsSuccess_(false).
				Message_(err.Error()).
				Build(),
		)
		return
	}

	if requestJson.From == "" || requestJson.To == "" || requestJson.Value.String() == "0" {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			config.Builder().
				IsSuccess_(false).
				Message_("Request values do not match").
				Build(),
		)
		return
	}

	hash, err := TokenApiService.TransferToken(requestJson.From, requestJson.To, requestJson.Value)

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			config.Builder().
				IsSuccess_(false).
				Message_(err.Error()).
				Build(),
		)
		return
	}

	type Result struct {
		Hash common.Hash
	}

	c.JSON(http.StatusOK, config.Builder().
		IsSuccess_(true).
		Body_(Result{Hash: hash}).
		Build(),
	)
}

// BalanceToken godoc
// @Summary BalanceToken
// @Description do BalanceToken
// @name BalanceToken
// @Accept  json
// @Produce  json
// @Param	address query string true "Account Address" example(0x46C65D87bE47255882561bcc7CFf3bBA186F0848)
// @Router /token/balanceToken [get]
// @Success 200 {object} config.ResponseObject
// @Example
// @Failure 400
func (ctrl TokenApiController) BalanceToken(c *gin.Context) {

	address := c.Query("address")

	if address == "" {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			config.Builder().
				IsSuccess_(false).
				Message_("Address values do not match").
				Build(),
		)
		return
	}

	balance, err := TokenApiService.BalanceToken(address)

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			config.Builder().
				IsSuccess_(false).
				Message_(err.Error()).
				Build(),
		)
		return
	}

	type Result struct {
		Balance *big.Int
	}

	fmt.Printf("Result{Balance: balance}: %v\n", Result{Balance: balance})

	c.JSON(http.StatusOK, config.Builder().
		IsSuccess_(true).
		Body_(Result{Balance: balance}).
		Build(),
	)
}

// TransactionReceipt godoc
// @Summary TransactionReceipt
// @Description do TransactionReceipt
// @name TransactionReceipt
// @Accept  json
// @Produce  json
// @Param	hash query string true "Hash Address" example(0x81f05c7087dc3fc1ff6576f6303694b803122366e5fba54afd1f526a2f567d56)
// @Router /token/transactionReceipt [get]
// @Success 200 {object} config.ResponseObject
// @Example
// @Failure 400
func (ctrl TokenApiController) TransactionReceipt(c *gin.Context) {

	hash := c.Query("hash")

	if hash == "" {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			config.Builder().
				IsSuccess_(false).
				Message_("Hash values do not match").
				Build(),
		)
		return
	}

	receipt, err := TokenApiService.TransactionReceipt(hash)

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			config.Builder().
				IsSuccess_(false).
				Message_(err.Error()).
				Build(),
		)
		return
	}

	type Result struct {
		Status uint64
	}

	c.JSON(http.StatusOK, config.Builder().
		IsSuccess_(true).
		Body_(Result{Status: receipt.Status}).
		Build(),
	)
}
