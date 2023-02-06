package server

import (
	"time"

	"git.garena.com/sea-labs-id/batch-02/andreas-timothy/entry-task-mo-be/handlers"
	"git.garena.com/sea-labs-id/batch-02/andreas-timothy/entry-task-mo-be/middlewares"
	"git.garena.com/sea-labs-id/batch-02/andreas-timothy/entry-task-mo-be/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type RouterConfig struct {
	UserService     services.UserService
	EmployeeService  services.EmployeeService
}

func NewRouter(c *RouterConfig) *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Content-Length", "Authorization"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "*"
		},
		MaxAge: 12 * time.Hour,
	}))

	h := handlers.New(&handlers.HandlerConfig{
		UserService:     c.UserService,
		EmployeeService:  c.EmployeeService,
	})

	r.NoRoute(middlewares.NotFoundError)

	r.Use(middlewares.ErrorMiddleware)

	authorizedGroup := r.Group("/")
	authorizedGroup.Use(middlewares.AuthorizeJWT)

	adminGroup := authorizedGroup.Group("/")
	adminGroup.Use(middlewares.AuthorizeRole("admin"))
	
	r.POST("/login", h.Login)
	r.POST("/login-with-google", h.LoginWithGoogleHandler)
	r.GET("/employees", h.GetAllEmployee)

	adminGroup.POST("/employees", h.AddEmployee)
	adminGroup.PATCH("/employees/:id", h.UpdateEmployee)
	adminGroup.DELETE("/employees/:id", h.DeleteEmployee)

	r.Run()

	return r
}
