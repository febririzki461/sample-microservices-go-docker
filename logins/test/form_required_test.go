package gin
 
import (
//	"fmt"
//	"net/http"
//	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"
//"github.com/gin-gonic/gin"
)


func TestLoginRequired(t *testing.T){

	usernya := ""
        passnya := ""

        if usernya != "" {
                assert := assert.New(t)
    		assert.Equal(usernya, false, "seharusnya false")

        }

        if passnya != "" {
		assert := assert.New(t)
                assert.Equal(usernya, false, "seharusnya false")
        }
}


