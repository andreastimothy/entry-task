package handlers

import "git.garena.com/sea-labs-id/batch-02/andreas-timothy/entry-task-mo-be/services"

type Handler struct {
	userService    services.UserService
	employeeService services.EmployeeService
}

type HandlerConfig struct {
	UserService    services.UserService
	EmployeeService services.EmployeeService
}

func New(c *HandlerConfig) *Handler {
	return &Handler{
		userService:    c.UserService,
		employeeService: c.EmployeeService,
	}
}
