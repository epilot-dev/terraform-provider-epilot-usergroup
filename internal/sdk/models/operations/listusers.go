// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/epilot-dev/terraform-provider-epilot-usergroup/internal/sdk/models/shared"
	"net/http"
)

type ListUsersRequest struct {
	// Comma-separated list of organization ids to filter by
	OrgIds []string `queryParam:"style=form,explode=false,name=org_ids"`
	// Query text to filter by
	Query *string `queryParam:"style=form,explode=true,name=query"`
	// Limit the results size
	Limit *float64 `queryParam:"style=form,explode=true,name=limit"`
	// Specify the offset
	Offset *float64 `queryParam:"style=form,explode=true,name=offset"`
}

func (o *ListUsersRequest) GetOrgIds() []string {
	if o == nil {
		return nil
	}
	return o.OrgIds
}

func (o *ListUsersRequest) GetQuery() *string {
	if o == nil {
		return nil
	}
	return o.Query
}

func (o *ListUsersRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListUsersRequest) GetOffset() *float64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

// ListUsersResponseBody - List of users
type ListUsersResponseBody struct {
	Users []shared.User `json:"users,omitempty"`
}

func (o *ListUsersResponseBody) GetUsers() []shared.User {
	if o == nil {
		return nil
	}
	return o.Users
}

type ListUsersResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// List of users
	Object *ListUsersResponseBody
}

func (o *ListUsersResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListUsersResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListUsersResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListUsersResponse) GetObject() *ListUsersResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
