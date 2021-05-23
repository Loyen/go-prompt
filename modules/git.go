package modules

import (
	"bufio"
	"os/exec"
	"strconv"
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

func (git *Git) GetOutput() string {
	statuses := []string{}

	status := git.GetStatus()
	if status.Addition > 0 {
		statuses = append(statuses, "+"+strconv.Itoa(status.Addition))
	}

	if status.Deletion > 0 {
		statuses = append(statuses, "-"+strconv.Itoa(status.Deletion))
	}

	if status.Untracked > 0 {
		statuses = append(statuses, "?"+strconv.Itoa(status.Untracked))
	}

	if status.Staged > 0 {
		statuses = append(statuses, "^"+strconv.Itoa(status.Staged))
	}

	if status.Stashed > 0 {
		statuses = append(statuses, "="+strconv.Itoa(status.Stashed))
	}

	if status.Conflict > 0 {
		statuses = append(statuses, "*"+strconv.Itoa(status.Conflict))
	}

	output := git.GetCurrentBranchName()

	if len(statuses) > 0 {
		output += " " + strings.Join(statuses, " ")
	}

	return output
}
