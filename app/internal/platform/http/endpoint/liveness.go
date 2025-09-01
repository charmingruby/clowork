package endpoint

import (
	"github.com/charmingruby/clowork/pkg/delivery/http/rest"
	"github.com/gin-gonic/gin"
)

func makeLiveness() gin.HandlerFunc {
	return func(c *gin.Context) {
		rest.SendOKResponse(c, "", nil)
	}
}
