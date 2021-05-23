package main

import (
	"fmt"
	"os"

	modules "github.com/loyen/go-prompt/modules"
)

func main() {
	path := os.Getenv("PWD")

	location := modules.Location{}
	location.SetPath(path)

	locationOutput := location.GetOutput()

	gitOutput := ""
	gitRepository := modules.Git{}
	gitRepository.SetPath(path)

	if gitRepository.IsGitRepository() {
		gitOutput = gitRepository.GetOutput()
	}

	fmt.Printf("%s %s\n%s",
		gitOutput,
		locationOutput,
		"$ ",
	)
}
