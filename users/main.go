package main

import (
	"strconv"
	"net/http"
	"github.com/gin-gonic/gin"
)

type Users struct {
    Id        int    `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
    Firstname string `gorm:"not null" form:"firstname" json:"firstname"`
    Lastname  string `gorm:"not null" form:"lastname" json:"lastname"`
}


func GetUsers(c *gin.Context){
	var users = []Users{
        	Users{Id: 1, Firstname: "Oliver", Lastname: "Queen"},
        	Users{Id: 2, Firstname: "Malcom", Lastname: "Merlyn"},
    }

    c.JSON(http.StatusOK, gin.H{"status": "success","result": users})
}

func GetUser(c *gin.Context) {
    id := c.Params.ByName("id")
    user_id, _ := strconv.ParseInt(id, 0, 64)

    if user_id == 1 {
        content := gin.H{"id": user_id, "firstname": "Oliver", "lastname": "Queen"}
 	c.JSON(200, gin.H{"code":200,"status": "success","result": content})
    } else {
        content := gin.H{"error": "user with id#" + id + " not found"}
        c.JSON(404, gin.H{"code":404,"status": "failed","result": content})

    }

    // curl -i http://localhost:8080/api/v1/users/1
}



func main() {
	//router := gin.Default()
        router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/users", GetUsers)
	router.GET("/users/:id", GetUser)
	// Listen and serve on 0.0.0.0:8080
	router.Run(":8082")
}
