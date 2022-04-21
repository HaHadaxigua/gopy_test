package factory

import "fmt"

type A struct {
	NameA string
}

func (a A) Creator() CreateFunction {
	return func() (XInterface, error) {
		return A{
			NameA: "helloA",
		}, nil
	}
}

func (a A) DoA() {
	fmt.Println("DoA: im A")
}

func (a A) DoB() error {
	fmt.Println("DoB: im A")
	return nil
}

// -----------------------------------------------------------------------------

func init() {
	AddFuncToFactory("A", A{})
}
