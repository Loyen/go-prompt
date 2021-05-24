package modules

import (
	"os"
)

type UserModule struct {
}

func (userModule *UserModule) GetOutput() string {
	userType := ""

	if os.Getenv("USER") == "root" {
		userType = "ROOT"
	} else {
		userType = "DEFAULT"
	}

	return formatValue("USER", os.Getenv("GO_PROMPT_USER_"+userType))
}
