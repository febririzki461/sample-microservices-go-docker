package gin
 
import (
//	"fmt"
//	"net/http"
//	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"
//"github.com/gin-gonic/gin"
)


func TestLoginFormUser(t *testing.T) {
        usernya := "manu"

        if usernya != "123" {
                assert := assert.New(t)
                assert.NotEqual(usernya, false, "seharusnya manu")

        }
}
