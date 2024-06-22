package helpers

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func BindData(c *gin.Context, data interface{}) error {
	if err := c.ShouldBindJSON(data); err != nil {
		return errors.New("invalid input")
	}
	return nil
}
