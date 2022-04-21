package factory

import (
	"fmt"
)

// XInterface need to implement the XInterface
type XInterface interface {
	DoA()
	DoB() error
}

// -----------------------------------------------------------------------------

func BuildXInterface(typeName string) (XInterface, error) {
	fn, ok := GetFuncFromFactory(typeName)
	if !ok {
		return nil, fmt.Errorf("cannot build instance of %s", typeName)
	}
	return fn()
}

// -----------------------------------------------------------------------------

type (
	XInterfaceCreator interface {
		Creator() CreateFunction
	}

	CreateFunction func() (XInterface, error)
)

var factory = map[string]CreateFunction{}

func AddFuncToFactory(typeName string, pc XInterfaceCreator) {
	factory[typeName] = pc.Creator()
}

func GetFuncFromFactory(typeName string) (CreateFunction, bool) {
	fn, ok := factory[typeName]
	return fn, ok
}

// -----------------------------------------------------------------------------

func IsA(x XInterface) bool {
	_, ok := x.(A)
	return ok
}

func IsB(x XInterface) bool {
	_, ok := x.(B)
	return ok
}
