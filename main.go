package main

import (
	"fmt"
	"os"

	git "github.com/loyen/go-prompt/modules"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("No Path given.")
		return
	}

	path := string(os.Args[1])

	gitRepository := git.Git{}
	gitRepository.SetPath(path)
	gitInfo := ""

	if gitRepository.IsGitRepository() {
		gitStatus := gitRepository.GetStatus()
		gitInfo = fmt.Sprintf("%s +%d -%d ?%d ^%d !%d =%d",
			gitRepository.GetCurrentBranchName(),
			gitStatus.Addition,
			gitStatus.Deletion,
			gitStatus.Untracked,
			gitStatus.Staged,
			gitStatus.Conflict,
			gitStatus.Stashed,
		)
	}

	fmt.Printf("%s\n%s", gitInfo, "$ ")
}
