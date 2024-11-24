package configtaskscmd

import (
	"context"
	"fmt"
	"regexp"

	"github.com/manifoldco/promptui"

	ytrackercore "github.com/emildeev/harvest-yt/internal/core/y_tracker"
	"github.com/emildeev/harvest-yt/internal/usecase"
)

func getTrackerTaskKey(ctx context.Context, defaultVal string, provider *usecase.Provider) (string, error) {
	newTaskKeyPromptValidate := func(input string) error {
		if input == "" {
			return nil
		}
		validateRegexp := ytrackercore.TicketKeyRegexp + "$"
		matched, err := regexp.MatchString(validateRegexp, input)
		if err != nil {
			return fmt.Errorf("validate error: %w", err)
		}
		if !matched {
			return fmt.Errorf("invalid account id format, shuld be: %s, like FULLSTACK-1111", validateRegexp)
		}

		err = provider.Validator.ValidateYTrackerTicket(ctx, input)
		return err
	}

	newTaskKeyPrompt := promptui.Prompt{
		Label:     "Yandex tracker task key (empty for delete)",
		Default:   defaultVal,
		Validate:  newTaskKeyPromptValidate,
		AllowEdit: true,
	}

	taskKey, err := newTaskKeyPrompt.Run()
	if err != nil {
		return defaultVal, err
	}

	return taskKey, nil
}
