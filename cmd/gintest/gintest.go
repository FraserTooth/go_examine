package gintest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelloHandler(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "Hello %s", name)
}

func WelcomeHandler(c *gin.Context) {
	c.String(http.StatusOK, "Hello World from Go")
}
