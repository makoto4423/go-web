package filter

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var filterMap map[string]interface{}

func init() {
	filterMap = make(map[string]interface{})
	filterMap["/login"] = nil
}

func Filter() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		if _, ok := filterMap[path]; !ok {
			token := c.Request.Header.Values("Authorization")
			fmt.Println("token is ", token)
		}
		c.Next()
	}
}
