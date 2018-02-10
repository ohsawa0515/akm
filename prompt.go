package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/manifoldco/promptui"
)

func accessKeyIdValidate(input string) error {
	if !strings.HasPrefix(input, "AKI") {
		return errors.New("access key starts with `AKI`")
	}

	return nil
}

func secretAccessKeyValidate(input string) error {
	return nil
}

func (ac *AwsCredential) AccessKeyIdPrompt() error {
	var label string
	message := "AWS Access Key ID [%s]"
	if len(ac.AccessKeyId) == 0 {
		label = fmt.Sprintf(message, "nil")
	} else {
		label = fmt.Sprintf(message, ac.AccessKeyId)
	}

	prompt := promptui.Prompt{
		Label:    label,
		Validate: accessKeyIdValidate,
		Default:  ac.AccessKeyId,
	}

	result, err := prompt.Run()
	ac.AccessKeyId = result
	if err != nil {
		return err
	}

	return nil
}

func (ac *AwsCredential) SecretAccessKeyPrompt() error {
	var label string
	message := "AWS Secret Access Key [%s]"
	if len(ac.SecretAccessKey) == 0 {
		label = fmt.Sprintf(message, "nil")
	} else {
		label = fmt.Sprintf(message, "****")
	}

	prompt := promptui.Prompt{
		Label:    label,
		Validate: secretAccessKeyValidate,
		Default:  ac.SecretAccessKey,
		Mask:     '*',
	}

	result, err := prompt.Run()
	ac.SecretAccessKey = result
	if err != nil {
		return err
	}

	return nil
}

func (ac *AwsCredential) RegionPrompt() error {
	var label string
	message := "Default region name [%s]"
	if len(ac.Region) == 0 {
		label = fmt.Sprintf(message, "nil")
	} else {
		label = fmt.Sprintf(message, ac.Region)
	}

	prompt := promptui.Select{
		Label: label,
		Items: Regions(),
		Size:  len(Regions()),
	}
	_, result, err := prompt.Run()
	ac.Region = result
	if err != nil {
		return err
	}

	return nil
}
