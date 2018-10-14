package gin
 
import (
//	"fmt"
//	"net/http"
//	"net/http/httptest"
	"testing"
//"github.com/gin-gonic/gin"
)


func TestLoginRequired(t *testing.T){

	usernya := "cbn"
        passnya := "admin"

        if usernya != "" {
                t.Errorf("required ","data user tidak boleh kosong")
        }

        if passnya != "" {
                t.Errorf("required ","data pass tidak boleh kosong")
        }



}

