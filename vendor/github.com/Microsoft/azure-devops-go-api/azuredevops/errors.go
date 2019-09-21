package azuredevops

type ArgumentNilError struct {
	ArgumentName string
}

func (e ArgumentNilError) Error() string {
	return "Argument " + e.ArgumentName + " can not be nil"
}

type ArgumentNilOrEmptyError struct {
	ArgumentName string
}

func (e ArgumentNilOrEmptyError) Error() string {
	return "Argument " + e.ArgumentName + " can not be nil or empty"
}
