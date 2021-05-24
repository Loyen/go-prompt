package main

import (
	"fmt"
	"os"
	"strings"

	modules "github.com/loyen/go-prompt/modules"
)

func main() {
	path := os.Getenv("PWD")
	parts := []string{}

	timeModule := modules.TimeModule{}
	parts = append(parts, timeModule.GetOutput())

	gitModule := modules.GitModule{}
	gitModule.SetPath(path)

	if gitModule.IsGitRepository() {
		parts = append(parts, gitModule.GetOutput())
	}

	locationModule := modules.LocationModule{}
	locationModule.SetPath(path)

	parts = append(parts, locationModule.GetOutput())

	userModule := modules.UserModule{}
	parts = append(parts, userModule.GetOutput())

	fmt.Print(strings.Join(parts, os.Getenv("GO_PROMPT_SEPARATOR")))
}
