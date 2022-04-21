package factory

func BuildXInterfaces(typeNames []string) ([]XInterface, error) {
	var ret []XInterface
	for _, tn := range typeNames {
		xInterface, err := BuildXInterface(tn)
		if err != nil {
			return nil, err
		}
		ret = append(ret, xInterface)
	}
	return ret, nil
}
