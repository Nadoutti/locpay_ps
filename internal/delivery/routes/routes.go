package routes

import (
	"locpay_api/internal/delivery/controllers"

	"github.com/gin-gonic/gin"
)


func SetupRouter(r *gin.Engine) {


	api := r.Group("/api/v1")
	{

		// rotas operations
		operations := api.Group("/operations")
		{
			operations.POST("", controllers.CreateOperation)
			operations.GET("/:id", controllers.GetOperationById)
			operations.POST("/:id/confirm", controllers.ConfirmOperation)
		}

		// rotas receivers
		receivers := api.Group("/receivers")
		{
			receivers.GET("/:id", controllers.GetRecieverById)
		}

	}

}
