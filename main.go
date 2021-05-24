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

	gitOutput := ""
	gitRepository := modules.Git{}
	gitRepository.SetPath(path)

	if gitRepository.IsGitRepository() {
		gitOutput = gitRepository.GetOutput()
	}

	user := modules.User{}

	fmt.Printf("%s%s%s",
		gitOutput,
		location.GetOutput(),
		user.GetOutput(),
	)
}
