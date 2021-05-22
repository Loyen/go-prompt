package main

import (
	"fmt"
	"os"

	"github.com/loyen/go-prompt/git"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("No Path given.")
		return
	}

	path := string(os.Args[1])

	isGitDir := git.IsGitDir(path)

	if isGitDir {
		info, _ := git.GetInfoByPath(path)
		fmt.Printf("Git info\n%s", info.Branch)
		return
	}

	fmt.Println("Empty")
}
