// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package int32validators

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ validator.Int32 = Int32NotNullValidator{}

// Int32NotNullValidator validates that an attribute is not null. Most
// attributes should set Required: true instead, however in certain scenarios,
// such as a computed nested attribute, all underlying attributes must also be
// computed for planning to not show unexpected differences.
type Int32NotNullValidator struct{}

// Description describes the validation in plain text formatting.
func (v Int32NotNullValidator) Description(_ context.Context) string {
	return "value must be configured"
}

// MarkdownDescription describes the validation in Markdown formatting.
func (v Int32NotNullValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

// Validate performs the validation.
func (v Int32NotNullValidator) ValidateInt32(ctx context.Context, req validator.Int32Request, resp *validator.Int32Response) {
	if !req.ConfigValue.IsNull() {
		return
	}

	resp.Diagnostics.AddAttributeError(
		req.Path,
		"Missing Attribute Value",
		req.Path.String()+": "+v.Description(ctx),
	)
}

// NotNull returns an validator which ensures that the attribute is
// configured. Most attributes should set Required: true instead, however in
// certain scenarios, such as a computed nested attribute, all underlying
// attributes must also be computed for planning to not show unexpected
// differences.
func NotNull() validator.Int32 {
	return Int32NotNullValidator{}
}
