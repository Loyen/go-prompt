package modules

import (
	"bufio"
	"os/exec"
	"strconv"
	"strings"
)

type GitModule struct {
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

func (gitModule *GitModule) SetPath(path string) {
	gitPathCmd := exec.Command("git", "-C", path, "rev-parse", "--show-toplevel")
	gitPathOutput, err := gitPathCmd.Output()

	if err == nil {
		gitModule.Path = strings.Trim(string(gitPathOutput), "\n ")
	}
}

func (gitModule *GitModule) IsGitRepository() bool {
	return gitModule.Path != ""
}

func (gitModule *GitModule) GetCurrentBranchName() string {
	gitBranchCmd := exec.Command("git", "-C", gitModule.Path, "rev-parse", "--abbrev-ref", "HEAD")
	gitBranchOutput, _ := gitBranchCmd.Output()

	return strings.Trim(string(gitBranchOutput), "\n ")
}

func (gitModule *GitModule) GetStatus() GitStatus {
	gitStatusCmd := exec.Command("git", "-C", gitModule.Path, "status", "--porcelain")
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

func (gitModule *GitModule) GetOutput() string {
	statuses := []string{}

	status := gitModule.GetStatus()
	if status.Addition > 0 {
		statuses = append(
			statuses,
			formatValue("GIT_ADDITION", strconv.Itoa(status.Addition)),
		)
	}

	if status.Deletion > 0 {
		statuses = append(
			statuses,
			formatValue("GIT_DELETION", strconv.Itoa(status.Deletion)),
		)
	}

	if status.Untracked > 0 {
		statuses = append(
			statuses,
			formatValue("GIT_UNTRACKED", strconv.Itoa(status.Untracked)),
		)
	}

	if status.Staged > 0 {
		statuses = append(
			statuses,
			formatValue("GIT_STAGED", strconv.Itoa(status.Staged)),
		)
	}

	if status.Stashed > 0 {
		statuses = append(
			statuses,
			formatValue("GIT_STASHED", strconv.Itoa(status.Stashed)),
		)
	}

	if status.Conflict > 0 {
		statuses = append(
			statuses,
			formatValue("GIT_CONFLICT", strconv.Itoa(status.Conflict)),
		)
	}

	output := formatValue("GIT_BRANCH", gitModule.GetCurrentBranchName())

	if len(statuses) > 0 {
		output += " " + strings.Join(statuses, " ")
	}

	return output
}
