package gin
 
import (
//	"fmt"
//	"net/http"
//	"net/http/httptest"
	"testing"
//"github.com/gin-gonic/gin"
)

func TestLoginPassed(t *testing.T) {

        usernya := "cbn"

        passwordnya := "123"

	if usernya != "cbn" || passwordnya != "123" {
       		t.Errorf("login succesfully")
    	}
}
