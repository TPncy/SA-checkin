package main

import (
	"github.com/TPncy/sapro/controller"

	"github.com/TPncy/sapro/entity"
	"github.com/TPncy/sapro/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	entity.SetupDatabase()

	r := gin.Default()
	r.Use(CORSMiddleware())

	api := r.Group("")
	{
		protected := api.Use(middlewares.Authorizes())
		{
			// Room Routes
			protected.GET("/rooms", controller.ListRooms)
			protected.GET("/room/:id", controller.GetRoom)
			protected.POST("/rooms", controller.CreateRoom)
			protected.PATCH("/rooms", controller.UpdateRoom)
			protected.DELETE("/rooms/:id", controller.DeleteRoom)

			// Customer Routes
			protected.GET("/customers", controller.ListCustomers)
			protected.GET("/customer/:id", controller.GetCustomer)
			protected.POST("/customers", controller.CreateCustomer)
			protected.PATCH("/customers", controller.UpdateCustomer)
			protected.DELETE("/customers/:id", controller.DeleteCustomer)

			// RoomPayment Routes
			protected.GET("/roompayments", controller.ListRoomPayments)
			protected.GET("/roompayment/:id", controller.GetRoomPayment)
			protected.POST("/roompayments", controller.CreateRoomPayment)
			protected.PATCH("/roompayments", controller.UpdateRoomPayment)
			protected.DELETE("/roompayments/:id", controller.DeleteRoomPayment)

			// Employee Routes
			protected.GET("/employees", controller.ListEmployees)
			protected.GET("/employee/:id", controller.GetEmployee)
			//protected.POST("/employees", controller.CreateEmployee)
			protected.PATCH("/employees", controller.UpdateEmployee)
			protected.DELETE("/employees/:id", controller.DeleteEmployee)

			// CheckIn Routes
			protected.GET("/check_ins", controller.ListCheckIns)
			protected.GET("/checkins/:id", controller.GetCheckIn)
			protected.POST("/check_ins", controller.CreateCheckIn)
			protected.PATCH("/check_ins", controller.UpdateCheckIn)
			protected.DELETE("/checkins/:id", controller.DeleteCheckIn)

		}
	}

	// emp Routes
	r.POST("/employees", controller.CreateEmployee)

	// Authentication Routes
	r.POST("/login", controller.Login)

	// Run the server
	r.Run()
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
