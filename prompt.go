package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-ini/ini"
	"github.com/manifoldco/promptui"
	cli "gopkg.in/urfave/cli.v1"
)

func accessKeyIdValidate(input string) error {
	if len(input) > 0 && !strings.HasPrefix(input, "AKI") {
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
		label = fmt.Sprintf(message, "None")
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
		label = fmt.Sprintf(message, "None")
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
		label = fmt.Sprintf(message, "None")
	} else {
		label = fmt.Sprintf(message, ac.Region)
	}
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "-> {{ .Name }} [{{ .Code }}]",
		Inactive: "   {{ .Name }} [{{ .Code }}]",
		Selected: `{{ "\u2714" | green | bold }} {{ "Region:" | bold }} {{ .Name | bold }} {{ "[" | bold }}{{ .Code | bold }}{{ "]" | bold }}`,
	}
	prompt := promptui.Select{
		Label:     label,
		Items:     regions,
		Size:      len(regions),
		Templates: templates,
	}
	i, _, err := prompt.Run()
	ac.Region = regions[i].Code
	if err != nil {
		return err
	}

	return nil
}

func (ac *AwsCredential) SaveToCredentialsFilePrompt(profile, file string) error {
	cfg, err := ini.Load(file)
	if err != nil {
		return err
	}

	sec, err := cfg.NewSection(profile)
	if err != nil {
		return err
	}

	if len(ac.AccessKeyId) > 0 {
		if _, err := sec.NewKey("aws_access_key_id", ac.AccessKeyId); err != nil {
			return err
		}
	} else {
		sec.DeleteKey("aws_access_key_id")
	}

	if len(ac.SecretAccessKey) > 0 {
		if _, err := sec.NewKey("aws_secret_access_key", ac.SecretAccessKey); err != nil {
			return err
		}
	} else {
		sec.DeleteKey("aws_secret_access_key")
	}

	if len(ac.Region) > 0 {
		if _, err := sec.NewKey("region", ac.Region); err != nil {
			return err
		}
	} else {
		sec.DeleteKey("region")
	}

	// Save to credential file
	prompt := promptui.Prompt{
		Label:     fmt.Sprintf("Overwrite %s", file),
		IsConfirm: true,
	}
	if _, err := prompt.Run(); err != nil {
		return fmt.Errorf("not overwritten")
	}
	if err := cfg.SaveTo(file); err != nil {
		return cli.NewExitError(err, 1)
	}

	return nil
}

func (ac *AwsCredential) DeleteProfilePrompt(profile, file string) error {
	cfg, err := ini.Load(file)
	if err != nil {
		return err
	}

	cfg.DeleteSection(profile)

	// Save to credential file
	prompt := promptui.Prompt{
		Label:     fmt.Sprintf("Remove profile: %s, overwrite %s", profile, file),
		IsConfirm: true,
	}
	if _, err := prompt.Run(); err != nil {
		return fmt.Errorf("not deleted")
	}
	if err := cfg.SaveTo(file); err != nil {
		return err
	}

	return nil
}
