package middlewares

import (
	"net/http"

	"git.garena.com/sea-labs-id/batch-02/andreas-timothy/entry-task-mo-be/helpers"
	"github.com/gin-gonic/gin"
)

func ErrorMiddleware(c *gin.Context) {
	c.Next()
	if len(c.Errors) > 0 {
		err := c.Errors[0].Err
		errRes, ok := err.(helpers.FailedResponses)
		if ok {
			c.JSON(errRes.Code, errRes)
			return
		}
		c.JSON(http.StatusInternalServerError, helpers.FailedResponse(http.StatusInternalServerError, err.Error()))
		return
	}
}

func NotFoundError(c *gin.Context) {
	response := helpers.FailedResponse(http.StatusNotFound, "Page Not Found")
	c.JSON(response.Code, response)
}