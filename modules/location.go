package modules

import (
	"os"
	"strings"
)

type LocationModule struct {
	Path string
}

func (locationModule *LocationModule) SetPath(path string) {
	locationModule.Path = path
}

func (locationModule *LocationModule) processPath() string {
	outputPath := locationModule.Path
	userHome := "/home/" + os.Getenv("USER") + "/"
	if strings.HasPrefix(outputPath, userHome) {
		outputPath = strings.Replace(outputPath, userHome, "~/", 1)
	}

	return formatValue("LOCATION", outputPath)
}

func (locationModule *LocationModule) GetOutput() string {
	return locationModule.processPath()
}
