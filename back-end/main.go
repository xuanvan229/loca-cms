package main 


import (
	"github.com/gin-gonic/gin"
	"strconv"
)
type Users struct {
	Id	int  `form:"id" json:"id"`
	Firstname string `form:"firstname" json:"firstname"`
	Lastname string `form:"lastname" json:"lastname"`
}


func main () {
	r := gin.Default()
	v1 := r.Group("api/v1")
	{

		v1.POST("/users",PostUser)
		v1.GET("/users",GetUsers)
		v1.GET("/users/:id",GetUser)
		v1.Put("/user/:id",UpdateUser)
		v1.DELETE("/users/:id",DeleteUser)
	}
	r.Run()
}
