package main 


import (
	"strconv"
	"github.com/gin-gonic/gin"
)
type Users struct {
	Id	int  `form:"id" json:"id"`
	Firstname string `form:"firstname" json:"firstname"`
	Lastname string `form:"lastname" json:"lastname"`
}
func PostUsers(c *gin.Context){
}

func UpdateUser( c *gin.Context){

}
func DeleteUser( c *gin.Context){
}


func GetUsers(c *gin.Context){
	var users = []Users{
		Users{Id: 1, Firstname: "Xuan", Lastname:"Van"},
		Users{Id: 2, Firstname: "Loan", Lastname:"Van"},
	}
	c.JSON(200,users)
}
func GetUser(c *gin.Context){
	id := c.Params.ByName("id")
	user_id, _ := strconv.ParseInt(id,0,64)

	if user_id == 1 {
		content := gin.H{"id":user_id,"firstname":"Hong", "lastname":"Van"}
		c.JSON(200,content)
	} else if  user_id == 2 {
		content := gin.H{"id":user_id,"firstname":"Loan", "lastname":"Van"}
		c.JSON(200,content)
	} else {
		content := gin.H{"error":"user with id #" +id +"not found "}
		c.JSON(404, content)
	}
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

		v1.POST("/users",PostUsers)
		v1.GET("/users",GetUsers)
		v1.GET("/users/:id",GetUser)
		v1.PUT("/user/:id",UpdateUser)
		v1.DELETE("/users/:id",DeleteUser)
	}
	r.Run(":8000")
}
