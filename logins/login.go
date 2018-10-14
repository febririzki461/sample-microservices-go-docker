package main




import (
	"net/http"

	"github.com/gin-gonic/gin"
)


// Binding from JSON
type Login struct {
	User     string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}



func PostLogin(c *gin.Context) {
	var json Login
                if err := c.ShouldBindJSON(&json); err != nil {
                        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
                        return
                }

                if json.User != "manu" || json.Password != "123" {
                        c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
                        return
                } 

                c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})

}


func main() {
	//router := gin.Default()
        router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Example for binding JSON ({"user": "manu", "password": "123"})


	router.POST("/login", PostLogin)

	// Listen and serve on 0.0.0.0:8080
	router.Run(":9080")
}
