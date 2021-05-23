package modules

import "os"

type User struct {
}

func formatUserValue(userType string, value string) string {
	return os.Getenv("GO_PROMPT_USER_"+userType) + value + os.Getenv("GO_PROMPT_DEFAULT")
}

func (user *User) GetOutput() string {
	userName := os.Getenv("USER")

	prompt := ""

	if userName == "root" {
		prompt = formatUserValue("ROOT", "# ")
	} else {
		prompt = formatUserValue("DEFAULT", "$ ")
	}

	return prompt
}
