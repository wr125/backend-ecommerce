package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wr125/backend-ecommerce/controllers"
)

func UserRoutes(incomeRoutes *gin.Engine) {
	incomeRoutes.POST("/users/signup", controllers.SignUp())
	incomeRoutes.POST("/users/login", controllers.Login())
	incomeRoutes.POST("/admin/addproduct", controllers.ProductViewerAdmin())
	incomeRoutes.GET("/users/productview", controllers.SearchProduct())
	incomeRoutes.GET("/users/search", controllers.SearchProductByQuery())

}
