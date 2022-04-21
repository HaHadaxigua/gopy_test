package factory

import "fmt"

type B struct {
	NameB string
}

func (b B) Creator() CreateFunction {
	return func() (XInterface, error) {
		return B{
			NameB: "hello B",
		}, nil
	}
}

func (b B) DoA() {
	fmt.Println("DoA: im b")
}

func (b B) DoB() error {
	fmt.Println("DoB: im b")
	return nil
}

// -----------------------------------------------------------------------------

func init() {
	AddFuncToFactory("B", B{})
}
