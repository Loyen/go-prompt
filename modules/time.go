package modules

import (
	"time"
)

type TimeModule struct {
}

func (timeModule *TimeModule) GetOutput() string {
	return formatValue("TIME", time.Now().Format("15:04:05"))
}
