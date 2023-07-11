package core


func GetReaderFromProviderString(provider string) Reader{
	switch provider {
    case "mock":
        return &MockReader{}
	default:
		return nil
    }
}
