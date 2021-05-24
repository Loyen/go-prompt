package modules

import (
	"os"
)

func formatValue(envKey string, value string) string {
	return os.Getenv("GO_PROMPT_"+envKey) + value + os.Getenv("GO_PROMPT_DEFAULT")
}
