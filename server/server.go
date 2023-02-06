package server

import (
	"fmt"

	"git.garena.com/sea-labs-id/batch-02/andreas-timothy/entry-task-mo-be/database"
	"git.garena.com/sea-labs-id/batch-02/andreas-timothy/entry-task-mo-be/repositories"
	"git.garena.com/sea-labs-id/batch-02/andreas-timothy/entry-task-mo-be/services"
)

func Init() {
	ur := repositories.NewUserRepository(&repositories.URConfig{
		DB: database.Get(),
	})

	us := services.NewUserService(&services.USConfig{
		UserRepository: ur,
	})

	er := repositories.NewEmployeeRepository(&repositories.ERConfig{
		DB: database.Get(),
	})

	es := services.NewEmployeeService(&services.ESConfig{
		EmployeeRepository: er,
	})

	router := NewRouter(&RouterConfig{
		UserService: us,
		EmployeeService: es,
	})
	err := router.Run()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}