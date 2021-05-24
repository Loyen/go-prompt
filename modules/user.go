package modules

import (
	"os"
)

type User struct {
}

func (user *User) GetOutput() string {
	userType := ""

	if os.Getenv("USER") == "root" {
		userType = "ROOT"
	} else {
		userType = "DEFAULT"
	}

	return formatValue("USER", os.Getenv("GO_PROMPT_USER_"+userType))
}
