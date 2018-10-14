package gin
 
import (
//	"fmt"
//	"net/http"
//	"net/http/httptest"
	"testing"
//"github.com/gin-gonic/gin"
)


func TestLoginFormPass(t *testing.T) {
        passnya := "123"

        if passnya != "456" {
                t.Errorf("Password was incorrect ",passnya)
        }
}
