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

	var t string
	if c.Request.Method == "GET" {
		t = c.Query("type")
	} else if c.Request.Method == "POST" {
		t = c.PostForm("type")
	}
	if len(t) == 0 {
		return "json"
	}
	return t

}
