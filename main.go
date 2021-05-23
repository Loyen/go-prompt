package main

import (
	"fmt"
	"os"

	modules "github.com/loyen/go-prompt/modules"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("No Path given.")
		return
	}

	path := string(os.Args[1])

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
