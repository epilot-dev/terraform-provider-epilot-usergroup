// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/epilot-dev/terraform-provider-epilot-usergroup/internal/sdk/models/shared"
	"net/http"
)

type ActivateUserRequest struct {
	// Invite Token
	Token                 string                        `queryParam:"style=form,explode=true,name=token"`
	UserActivationPayload *shared.UserActivationPayload `request:"mediaType=application/json"`
}

func (o *ActivateUserRequest) GetToken() string {
	if o == nil {
		return ""
	}
	return o.Token
}

func (o *ActivateUserRequest) GetUserActivationPayload() *shared.UserActivationPayload {
	if o == nil {
		return nil
	}
	return o.UserActivationPayload
}

type ActivateUserResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ActivateUserResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ActivateUserResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ActivateUserResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
