package routers

import (
	"gin-dasar/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine{
  router := gin.Default()
  router.POST("/cars", controllers.CreateCar)
  router.PUT("/cars/:carId", controllers.UpdateCar)
  router.GET("cars/:carId", controllers.GetCar)
  router.DELETE("cars/:carId",controllers.DeleteCar)
  return router
}
