package utils

import (
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
)

var promptTemplate = &promptui.PromptTemplates{
	Prompt:          "{{ . | cyan }}️ ▶ ",
	Valid:           "{{ . | cyan }}️ ▶ ",
	Invalid:         "{{ . | red }}️ ▶ ",
	Success:         "{{ . | green }}️ ▶ ",
	ValidationError: "{{ . | red }}",
}

func GetInput(label string, defaultValue ...string) string {
	prompt := promptui.Prompt{
		Label:     label,
		Templates: promptTemplate,
		Validate: func(input string) error {
			if len(input) < 1 {
				return fmt.Errorf("❌ %s cannot be empty", label)
			}
			return nil
		},
	}

	if len(defaultValue) > 0 {
		prompt.Default = defaultValue[0]
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("❌ Prompt failed: %v\n", err)
		os.Exit(1)
	}

	return result
}
