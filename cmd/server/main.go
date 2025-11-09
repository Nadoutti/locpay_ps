package main

import (
	"github.com/gin-gonic/gin"
	"locpay_api/pkg/utils"
	"locpay_api/internal/delivery/routes"
)

func main() {

	// Initialize Supabase
	utils.InitSupabase()

	r := gin.Default()

	// Register routes
	routes.SetupRouter(r)

	r.Run() // listen and serve on

}
