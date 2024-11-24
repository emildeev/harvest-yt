package configcmd

import (
	"fmt"
	"regexp"

	"github.com/manifoldco/promptui"
)

func HarvestAccountID(defaultVal string) (string, error) {
	validate := func(input string) error {
		validateRegexp := "[0-9]+"
		matched, err := regexp.MatchString(validateRegexp, input)
		if err != nil {
			return fmt.Errorf("validate error: %w", err)
		}
		if !matched {
			return fmt.Errorf("invalid account id format, shuld be: %s", validateRegexp)
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:     "Harvest account id",
		Default:   defaultVal,
		Validate:  validate,
		AllowEdit: true,
	}

	accountID, err := prompt.Run()

	return accountID, err
}

func HarvestGetToken(defaultVal string) (string, error) {
	validate := func(input string) error {
		validateRegexp := "[0-9a-zA-Z\\-_\\.]+"
		matched, err := regexp.MatchString(validateRegexp, input)
		if err != nil {
			return fmt.Errorf("validate error: %w", err)
		}
		if !matched {
			return fmt.Errorf("invalid token format, shuld be: %s", validateRegexp)
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Harvest token",
		Default:  defaultVal,
		Mask:     '*',
		Validate: validate,
	}

	token, err := prompt.Run()

	return token, err
}
