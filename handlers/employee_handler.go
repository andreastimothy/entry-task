package handlers

import (
	"net/http"
	"strconv"

	"git.garena.com/sea-labs-id/batch-02/andreas-timothy/entry-task-mo-be/dtos"
	"git.garena.com/sea-labs-id/batch-02/andreas-timothy/entry-task-mo-be/helpers"
	"git.garena.com/sea-labs-id/batch-02/andreas-timothy/entry-task-mo-be/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllEmployee(c *gin.Context) {
	allEmployee, err := h.employeeService.GetAllEmployeees()

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, helpers.SuccessResponse(http.StatusOK, "SUCCESS", allEmployee))
}

func (h *Handler) AddEmployee(c *gin.Context) {
	var input *dtos.Employee
	err := c.ShouldBindJSON(&input)

	inputEmployee := models.Employee {
		Name: input.Name,
		JobDescription: input.JobDescription,
		EntryDate: input.EntryDate,
	}

	if err != nil {
		err := helpers.FailedResponse(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		c.Error(err)
		return
	}

	newEmployee, err := h.employeeService.AddEmployee(&inputEmployee)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, helpers.SuccessResponse(http.StatusOK, "SUCCESS", newEmployee))
}

func (h *Handler) UpdateEmployee(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var input *dtos.EditEmployee
	err := c.ShouldBindJSON(&input)

	inputEmployee := models.Employee {
		Name: input.Name,
	}

	if err != nil {
		err:= helpers.FailedResponse(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		c.Error(err)
		return
	}

	newEmployee, err := h.employeeService.UpdateEmployee(id, &inputEmployee)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, helpers.SuccessResponse(http.StatusOK, "SUCCESS", newEmployee))
}

func (h *Handler) DeleteEmployee(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	deletedEmployee, err := h.employeeService.DeleteEmployee(id)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, helpers.SuccessResponse(http.StatusOK, "SUCCESS", deletedEmployee))
}