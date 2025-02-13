package validator

import (
	"fmt"

	"github.com/authelia/authelia/v4/internal/configuration/schema"
	"github.com/authelia/authelia/v4/internal/utils"
)

// ValidateNTP validates and update NTP configuration.
func ValidateNTP(config *schema.Configuration, validator *schema.StructValidator) {
	if config.NTP == nil {
		config.NTP = &schema.DefaultNTPConfiguration

		return
	}

	if config.NTP.Address == "" {
		config.NTP.Address = schema.DefaultNTPConfiguration.Address
	}

	if config.NTP.Version == 0 {
		config.NTP.Version = schema.DefaultNTPConfiguration.Version
	} else if config.NTP.Version < 3 || config.NTP.Version > 4 {
		validator.Push(fmt.Errorf(errFmtNTPVersion, config.NTP.Version))
	}

	if config.NTP.MaximumDesync == "" {
		config.NTP.MaximumDesync = schema.DefaultNTPConfiguration.MaximumDesync
	}

	_, err := utils.ParseDurationString(config.NTP.MaximumDesync)
	if err != nil {
		validator.Push(fmt.Errorf(errFmtNTPMaxDesync, err))
	}
}
