package gpcorevalidator

import (
	cloudv1 "buf.build/gen/go/gportal/gpcore/protocolbuffers/go/gpcore/api/cloud/v1"
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"golang.org/x/exp/slices"
	"strings"
)

type BillingPeriodValidator struct {
}

// Description returns a plain text description of the validator's behavior, suitable for a practitioner to understand its impact.
func (v BillingPeriodValidator) Description(ctx context.Context) string {
	return "Validates the billing period."
}

// MarkdownDescription returns a markdown formatted description of the validator's behavior, suitable for a practitioner to understand its impact.
func (v BillingPeriodValidator) MarkdownDescription(ctx context.Context) string {
	return "Ensures a valid billing period is provided"
}

// ValidateString runs the main validation logic of the validator, reading configuration data out of `req` and updating `resp` with diagnostics.
func (v BillingPeriodValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	// If the value is unknown or null, there is nothing to validate.
	if req.ConfigValue.IsUnknown() || req.ConfigValue.IsNull() {
		return
	}
	if !slices.Contains(validBillingPeriods, req.ConfigValue.ValueString()) {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid Billing Period",
			fmt.Sprintf("Invalid billing period specified: %s\nValid billing periods: %v", req.ConfigValue.ValueString(), validBillingPeriods),
		)
	}
}

var validBillingPeriods []string

func init() {
	for _, s := range cloudv1.BillingPeriod_name {
		if strings.HasSuffix(s, "_UNSPECIFIED") {
			continue
		}
		validBillingPeriods = append(validBillingPeriods, s)
	}
}
