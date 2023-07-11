package core


func GetReaderFromProviderString(provider string, initArgs any) Reader{
	initArgsMap, ok := initArgs.(map[string]interface{})
	
	if !ok {
		return nil
	}

	switch provider {
    case "mock":
        return &MockReader{initArgsMap}
	default:
		return nil
    }
}
