package validator

import (
	"fmt"

	"github.com/concourse/concourse-pipeline-resource/concourse"
)

func ValidateCheck(input concourse.CheckRequest) error {
	if input.Source.Target == "" {
		return fmt.Errorf("%s must be provided in source", "target")
	}

	return ValidateTeams(input.Source.Teams)
}
