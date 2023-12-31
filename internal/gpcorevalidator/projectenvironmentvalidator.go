package gpcorevalidator

import (
	cloudv1 "buf.build/gen/go/gportal/gpcore/protocolbuffers/go/gpcore/api/cloud/v1"
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"golang.org/x/exp/slices"
	"strings"
)

type ProjectEnvironmentValidator struct {
}

// Description returns a plain text description of the validator's behavior, suitable for a practitioner to understand its impact.
func (v ProjectEnvironmentValidator) Description(ctx context.Context) string {
	return "Validates the project environment."
}

// MarkdownDescription returns a markdown formatted description of the validator's behavior, suitable for a practitioner to understand its impact.
func (v ProjectEnvironmentValidator) MarkdownDescription(ctx context.Context) string {
	return "Ensures a valid project environment is provided"
}

// ValidateString runs the main validation logic of the validator, reading configuration data out of `req` and updating `resp` with diagnostics.
func (v ProjectEnvironmentValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	// If the value is unknown or null, there is nothing to validate.
	if req.ConfigValue.IsUnknown() || req.ConfigValue.IsNull() {
		return
	}
	if !slices.Contains(validEnvironments, req.ConfigValue.ValueString()) {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid Project Environment",
			fmt.Sprintf("Invalid project environment specified: %s\nValid environments: %v", req.ConfigValue.ValueString(), validEnvironments),
		)
	}
}

var validEnvironments []string

func init() {
	for _, s := range cloudv1.ProjectEnvironment_name {
		if strings.HasSuffix(s, "_UNSPECIFIED") {
			continue
		}
		validEnvironments = append(validEnvironments, s)
	}
}
