package configcmd

import (
	"fmt"
	"time"

	"github.com/manifoldco/promptui"
)

func GetStartTime(defaultVal string) (string, error) {
	validate := func(input string) error {
		_, err := time.Parse(time.TimeOnly, input)
		if err != nil {
			return fmt.Errorf("validate error: %w", err)
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Time from start spend time (" + time.TimeOnly + ")",
		Default:  defaultVal,
		Validate: validate,
	}

	timeStr, err := prompt.Run()
	if err != nil {
		return "", err
	}

	return timeStr, nil
}
