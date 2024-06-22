package main

import (
	"log"
	"net/http"

	accSvc "com/txfer/bounded_contexts/account/service"
	txSvc "com/txfer/bounded_contexts/transfer/service"
	"com/txfer/configs"
	"com/txfer/db"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	server          *gin.Engine
	AccountHandler  accSvc.AccountHandler
	TransferHandler txSvc.TransferHandler
)

func main() {
	// Load configs
	config, err := configs.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	// Initialize DB
	db.InitDB(&config)

	// Initialize REST server
	initHTTPServer(config)
}

func initHTTPServer(config configs.Config) {
	AccountHandler = accSvc.NewAccountHandler(db.DB)
	TransferHandler = txSvc.NewTransferHandler(db.DB)
	gin.SetMode(gin.ReleaseMode)
	server = gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8000", config.ClientOrigin}
	corsConfig.AllowCredentials = true
	server.Use(cors.New(corsConfig))

	router := server.Group("/api/v1")
	initHealthCheckRoute(router)
	initAccountRoutes(router)
	initTransactionRoutes(router)
	log.Fatal(server.Run(":" + config.ServerPort))
}

func initHealthCheckRoute(router *gin.RouterGroup) {
	router.GET("/healthcheck", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "live", "message": "ok"})
	})
}
func initAccountRoutes(router *gin.RouterGroup) {
	accountGr := router.Group("accounts")
	accountGr.GET("/:account_id", AccountHandler.GetAccount)
	accountGr.POST("/", AccountHandler.CreateAccount)
}

func initTransactionRoutes(router *gin.RouterGroup) {
	transferGr := router.Group("transactions")
	transferGr.POST("/", TransferHandler.CreateTransfer)
}
