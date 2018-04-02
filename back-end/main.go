package main 


import (
	"github.com/gin-gonic/gin"
)
type Users struct {
	Id	int  `form:"id" json:"id"`
	Firstname string `form:"firstname" json:"firstname"`
	Lastname string `form:"lastname" json:"lastname"`
}
func PostUser(c *gin.Context){
	
}
func GetUsers(c *gin.Context){
	var users = []Users{
		Users{Id: 1, Firstname: "Xuan", Lastname: "Hong"},
		Users{Id: 2, Firstname: "Loan", Lastname: "Minh"},
	}
	c.JSON(200,users)
}
func GetUser(c *gin.Context){
	
	
}
func UpdateUser( c *gin.Context){
	
}
func DeleteUser( c *gin.Context){
}



func CORSMiddleware() gin.HandlerFunc{
	return func(c *gin.Context){
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		 c.Next()
	}
} 
func main () {
	r := gin.Default()
	r.Use(CORSMiddleware())
	v1 := r.Group("api/v1")
	{

		v1.POST("/users",PostUser)
		v1.GET("/users",GetUsers)
		v1.GET("/users/:id",GetUser)
		v1.PUT("/user/:id",UpdateUser)
		v1.DELETE("/users/:id",DeleteUser)
	}
	r.Run(":8000")
}
