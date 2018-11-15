package controller

import (
	"github.com/scalog/scalog-operator/pkg/controller/scalogservice"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, scalogservice.Add)
}
