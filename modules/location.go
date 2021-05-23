package modules

type Location struct {
	Path string
}

func (location *Location) SetPath(path string) {
	location.Path = path
}

func (location *Location) GetOutput() string {
	return location.Path
}
