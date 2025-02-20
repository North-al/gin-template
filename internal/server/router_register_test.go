package server

import (
	"testing"

	"github.com/gin-gonic/gin"
)

type TestHandler struct {
}

// @Auth()
func (t *TestHandler) TestRegister() int {
	return 1
}

func (t *TestHandler) TestLogin(c *gin.Context) string {
	return "login"
}

func NewTestRouter() *TestHandler {
	return &TestHandler{}
}

func TestRegister(t *testing.T) {
	r := gin.Default()

	AutoRegisterRoutes(r, NewTestRouter())
}
