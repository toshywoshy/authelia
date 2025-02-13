package validator

import (
	"fmt"

	"github.com/authelia/authelia/v4/internal/configuration/schema"
)

// ValidateNotifier validates and update notifier configuration.
func ValidateNotifier(config *schema.NotifierConfiguration, validator *schema.StructValidator) {
	if config == nil || (config.SMTP == nil && config.FileSystem == nil) {
		validator.Push(fmt.Errorf(errFmtNotifierNotConfigured))

		return
	} else if config.SMTP != nil && config.FileSystem != nil {
		validator.Push(fmt.Errorf(errFmtNotifierMultipleConfigured))

		return
	}

	if config.FileSystem != nil {
		if config.FileSystem.Filename == "" {
			validator.Push(fmt.Errorf(errFmtNotifierFileSystemFileNameNotConfigured))
		}

		return
	}

	validateSMTPNotifier(config.SMTP, validator)
}

func validateSMTPNotifier(config *schema.SMTPNotifierConfiguration, validator *schema.StructValidator) {
	if config.StartupCheckAddress == "" {
		config.StartupCheckAddress = schema.DefaultSMTPNotifierConfiguration.StartupCheckAddress
	}

	if config.Host == "" {
		validator.Push(fmt.Errorf(errFmtNotifierSMTPNotConfigured, "host"))
	}

	if config.Port == 0 {
		validator.Push(fmt.Errorf(errFmtNotifierSMTPNotConfigured, "port"))
	}

	if config.Timeout == 0 {
		config.Timeout = schema.DefaultSMTPNotifierConfiguration.Timeout
	}

	if config.Sender.Address == "" {
		validator.Push(fmt.Errorf(errFmtNotifierSMTPNotConfigured, "sender"))
	}

	if config.Subject == "" {
		config.Subject = schema.DefaultSMTPNotifierConfiguration.Subject
	}

	if config.Identifier == "" {
		config.Identifier = schema.DefaultSMTPNotifierConfiguration.Identifier
	}

	if config.TLS == nil {
		config.TLS = schema.DefaultSMTPNotifierConfiguration.TLS
	}

	if config.TLS.ServerName == "" {
		config.TLS.ServerName = config.Host
	}
}
