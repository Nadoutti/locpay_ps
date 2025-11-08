package server

import (
	"github.com/gin-gonic/gin"
	"locpay/db"
	"locpay/routes"
)

func main() {

	// Initialize Supabase
	db.InitSupabase()

	r := gin.Default()

	// Register routes
	routes.SetupRouter(r)

	r.Run() // listen and serve on

}
