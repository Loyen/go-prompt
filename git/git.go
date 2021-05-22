package git

import (
	"errors"
	"os/exec"
	"strings"
)

type GitInfo struct {
	Path   string
	Branch string
}

func IsGitDir(path string) bool {
	gitPathCmd := exec.Command("git", "-C", path, "rev-parse", "--absolute-git-dir")
	_, err := gitPathCmd.Output()

	return err == nil
}

func GetInfoByPath(path string) (GitInfo, error) {
	gitPathCmd := exec.Command("git", "-C", path, "rev-parse", "--absolute-git-dir")
	gitPathOutput, err := gitPathCmd.Output()

	if err != nil {
		return GitInfo{}, errors.New(err.Error())
	}

	gitPath := strings.Trim(string(gitPathOutput), "\n ")

	return GitInfo{
		Path:   gitPath,
		Branch: GetCurrentBranchName(gitPath),
	}, nil
}

func GetCurrentBranchName(path string) string {
	gitBranchCmd := exec.Command("git", "-C", path, "rev-parse", "--abbrev-ref", "HEAD")
	gitBranchOutput, _ := gitBranchCmd.Output()

	return string(gitBranchOutput)
}
