package modules

import (
	"bufio"
	"os/exec"
	"strings"
)

type Git struct {
	Path string
}

type GitStatus struct {
	Addition  int
	Deletion  int
	Untracked int
	Staged    int
	Stashed   int
	Conflict  int
}

func (git *Git) SetPath(path string) {
	gitPathCmd := exec.Command("git", "-C", path, "rev-parse", "--show-toplevel")
	gitPathOutput, err := gitPathCmd.Output()

	if err == nil {
		git.Path = strings.Trim(string(gitPathOutput), "\n ")
	}
}

func (git *Git) IsGitRepository() bool {
	return git.Path != ""
}

func (git *Git) GetCurrentBranchName() string {
	gitBranchCmd := exec.Command("git", "-C", git.Path, "rev-parse", "--abbrev-ref", "HEAD")
	gitBranchOutput, _ := gitBranchCmd.Output()

	return strings.Trim(string(gitBranchOutput), "\n ")
}

func (git *Git) GetStatus() GitStatus {
	gitStatusCmd := exec.Command("git", "-C", git.Path, "status", "--porcelain")
	gitStatusOutput, _ := gitStatusCmd.Output()

	scanner := bufio.NewScanner(strings.NewReader(string(gitStatusOutput)))

	status := GitStatus{}

	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), " M") {
			status.Addition++
		} else if strings.HasPrefix(scanner.Text(), " D") {
			status.Deletion++
		} else if strings.HasPrefix(scanner.Text(), " U") {
			status.Conflict++
		} else if strings.HasPrefix(scanner.Text(), "??") {
			status.Untracked++
		} else {
			status.Staged++
		}
	}

	return status
}
