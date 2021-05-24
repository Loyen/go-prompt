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

	gitRepository := modules.Git{}
	gitRepository.SetPath(path)

	if gitRepository.IsGitRepository() {
		parts = append(parts, gitRepository.GetOutput())
	}

	location := modules.Location{}
	location.SetPath(path)

	parts = append(parts, location.GetOutput())

	user := modules.User{}
	parts = append(parts, user.GetOutput())

	fmt.Print(strings.Join(parts, os.Getenv("GO_PROMPT_SEPARATOR")))
}
