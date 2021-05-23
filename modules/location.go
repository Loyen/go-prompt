package modules

import (
	"os"
	"strings"
)

type Location struct {
	Path string
}

func (location *Location) SetPath(path string) {
	location.Path = path
}

func (location *Location) processPath() string {
	outputPath := location.Path
	userHome := "/home/" + os.Getenv("USER") + "/"
	if strings.HasPrefix(outputPath, userHome) {
		outputPath = strings.Replace(outputPath, userHome, "~/", 1)
	}

	return outputPath
}

func (location *Location) GetOutput() string {
	return location.processPath()
}
