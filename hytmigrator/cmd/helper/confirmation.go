package helper

import (
	"errors"

	"github.com/manifoldco/promptui"
)

func GetConfirmation() bool {
	validateNumberPrompt := func(input string) error {
		if input != "y" && input != "n" && input != "Y" && input != "N" {
			return errors.New("it should be y or n")
		}
		return nil
	}

	numberPrompt := promptui.Prompt{
		Label:    "Confirm? (y/n)",
		Validate: validateNumberPrompt,
	}

	confirmation, err := numberPrompt.Run()
	if err != nil {
		return false
	}

	if confirmation == "y" || confirmation == "Y" {
		return true
	}
	return false
}
