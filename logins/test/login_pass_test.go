package gin
 
import (
//	"fmt"
//	"net/http"
//	"net/http/httptest"
	"testing"
//"github.com/gin-gonic/gin"
)

// Binding from JSON
type Login struct {
	User     string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func TestLoginPass(t *testing.T) {

        usernya := "cbn"

        passwordnya := "123"

	if usernya != "cbn" || passwordnya != "123" {
       		t.Errorf("login succesfully")
    	}
}
