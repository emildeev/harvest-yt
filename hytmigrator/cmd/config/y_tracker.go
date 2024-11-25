package configcmd

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/manifoldco/promptui"
)

func YTrackerGetOrgID(defaultVal int) (int, error) {
	fmt.Println()
	fmt.Println("for get organization ID: https://tracker.yandex.ru/settings")
	fmt.Println()

	validate := func(input string) error {
		if input == "" {
			return fmt.Errorf("organization ID is required")
		}
		val, err := strconv.Atoi(input)
		if err != nil {
			return fmt.Errorf("invalid value, shuld be number: %w", err)
		}
		if val == 0 {
			return fmt.Errorf("invalid value, shuld be more than 0")
		}
		return nil
	}

	defaultValSrt := ""
	if defaultVal != 0 {
		defaultValSrt = strconv.Itoa(defaultVal)
	}

	prompt := promptui.Prompt{
		Label:     "Yandex tracker organization ID",
		Default:   defaultValSrt,
		Validate:  validate,
		AllowEdit: true,
	}

	IDStr, err := prompt.Run()
	if err != nil {
		return 0, err
	}

	ID, err := strconv.Atoi(IDStr)
	if err != nil {
		return 0, fmt.Errorf("convert error: %w", err)
	}

	return ID, err
}

func YTrackerGetToken(defaultVal string) (string, error) {
	fmt.Println()
	fmt.Println(
		"for generate yandex tracker  token\n" +
			"https://oauth.yandex.ru/authorize?response_type=token&client_id=711865fe0ef3478ea09e895878cd275b",
	)
	fmt.Println()

	validate := func(input string) error {
		if input == "" {
			return fmt.Errorf("token is required")
		}
		validateRegexp := "[0-9a-zA-Z\\-_]{60}"
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
		Label:    "Yandex tracker token",
		Default:  defaultVal,
		Mask:     '*',
		Validate: validate,
	}

	token, err := prompt.Run()

	return token, err
}
