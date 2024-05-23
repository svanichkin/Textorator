package rest

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func checkInt(s string) int64 {

	if len(s) == 0 {
		return 0
	}
	u, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}
	return u

}

func checkType(c *gin.Context) string {

	if len(c.Query("type")) == 0 {
		return "json"
	}
	return c.Query("type")

}
