package controller

import (
	"github.com/hfuss/spring-boot-app-operator/pkg/controller/springbootapp"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, springbootapp.Add)
}
