package main

import (
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


func main() {
	//router := gin.Default()
        router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/users", GetUsers)

	// Listen and serve on 0.0.0.0:8080
	router.Run(":8082")
}
