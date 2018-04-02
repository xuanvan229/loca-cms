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
func PostUsers(c *gin.Context){

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

func UpdateUser( c *gin.Context){

}

func DeleteUser(c *gin.Context){

}
func main () {
	r := gin.Default()
	v1 := r.Group("api/v1")
	{

		v1.POST("/users",PostUsers)
		v1.GET("/users",GetUsers)
		v1.GET("/users/:id",GetUser)
		v1.PUT("/user/:id",UpdateUser)
		v1.DELETE("/users/:id",DeleteUser)
	}
	r.Run()
}
