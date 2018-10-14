package gin
 
import (
//	"fmt"
//	"net/http"
//	"net/http/httptest"
	"testing"
//"github.com/gin-gonic/gin"
)


func TestLoginFormUser(t *testing.T) {
        usernya := "manu"

        if usernya != "123" {
                t.Errorf("User was incorrect ",usernya)
        }
}
