package gitlab

import (
	"context"
	"errors"
	"fmt"
	"strings"

	gitlabcore "github.com/emildeev/gitlab_helper/internal/core/gitlab"
	"github.com/manifoldco/promptui"
)

const (
	defaultMRTemplateName = "default"

	ticketField = "**Тикет:** XXX"

	tickerBaseURL = "https://tracker.yandex.ru/"
)

func (service *Service) GetDefaultMRTemplateDescription(
	ctx context.Context, projectID gitlabcore.ProjectID,
) (description string, err error) {
	mrTemplate, err := service.projectTemplateAdapter.GetMRTemplate(ctx, int(projectID), defaultMRTemplateName)
	if err != nil {
		return description, fmt.Errorf("GetDefaultMRTemplateDescription: %w", err)
	}
	return mrTemplate.Description, nil
}

func (service *Service) FillMRTemplateDescription(_ context.Context, description, tickerKey string) string {
	fields := service.mrTemplateFieldRegExp.FindAllString(description, -1)
	for _, field := range fields {
		if field == ticketField {
			ticket := strings.Replace(ticketField, "XXX", tickerBaseURL+tickerKey, 1)
			description = strings.Replace(description, ticketField, ticket, 1)
		} else {
			description = fillMRTemplateField(description, field)
		}
	}
	return description
}

func fillMRTemplateField(description, field string) string {
	fieldText := strings.Replace(field, "**", "", -1)
	fieldText = strings.Replace(fieldText, "XXX", "", -1)
	fieldText = strings.Replace(fieldText, ":", "", -1)
	fieldText = strings.TrimSpace(fieldText)
	fieldVal, err := getField(fieldText)
	if err != nil {
		return description
	}
	fieldNew := strings.Replace(field, "XXX", fieldVal, 1)
	description = strings.Replace(description, field, fieldNew, 1)
	return description
}

func getField(fieldText string) (string, error) {
	validate := func(input string) error {
		if input == "" {
			return errors.New("empty brunch")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:     fieldText,
		Validate:  validate,
		AllowEdit: true,
	}

	if fieldText == "План отката" {
		prompt.Default = "реверт"
	}

	fieldVal, err := prompt.Run()

	return fieldVal, err
}
