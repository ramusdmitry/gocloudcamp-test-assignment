package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ErrorMessage struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, logMessage, errMessage string) {
	logrus.Errorf("%s, cause: %s", logMessage, errMessage)
	c.AbortWithStatusJSON(statusCode, ErrorMessage{errMessage})
}
