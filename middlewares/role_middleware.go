package middlewares

import (
	"net/http"
	"os"

	"git.garena.com/sea-labs-id/batch-02/andreas-timothy/entry-task-mo-be/helpers"
	"git.garena.com/sea-labs-id/batch-02/andreas-timothy/entry-task-mo-be/models"
	"github.com/gin-gonic/gin"
)

func AuthorizeRole(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userContext, _ := c.Get("user")
		uContext := userContext.(models.User)

		if os.Getenv("STAGE") == "testing" {
			user := models.User{}
			c.Set("user", user)
			c.Next()
			return
		}

		if uContext.Role != role {
			c.AbortWithStatusJSON(http.StatusForbidden, helpers.FailedResponse(http.StatusForbidden, "Forbidden Access"))
			return
		}
		c.Next()
	}
}
