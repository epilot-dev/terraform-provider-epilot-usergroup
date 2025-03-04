// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"net/http"
)

type CheckInviteTokenRequest struct {
	// Invite Token
	Token string `queryParam:"style=form,explode=true,name=token"`
}

func (o *CheckInviteTokenRequest) GetToken() string {
	if o == nil {
		return ""
	}
	return o.Token
}

// CheckInviteTokenResponseBody - Invite token found
type CheckInviteTokenResponseBody struct {
	// Organization ID of the organization that invited the user
	InvitationOrgID string `json:"invitation_org_id"`
	// Name of the organization that invited the user
	InvitationOrgName string `json:"invitation_org_name"`
	// Logo URL of the organization that invited the user
	InvitationOrgLogoURL *string `json:"invitation_org_logo_url,omitempty"`
	// Logo Thumbnail URL of the organization that invited the user
	InvitationOrgLogoThumbnailURL *string `json:"invitation_org_logo_thumbnail_url,omitempty"`
	// User ID of the invited user
	InviteeUserID string `json:"invitee_user_id"`
	// Organization ID of the primary organization of the user (when inviting an existing epilot user)
	InviteePrimaryOrgID *string `json:"invitee_primary_org_id,omitempty"`
}

func (o *CheckInviteTokenResponseBody) GetInvitationOrgID() string {
	if o == nil {
		return ""
	}
	return o.InvitationOrgID
}

func (o *CheckInviteTokenResponseBody) GetInvitationOrgName() string {
	if o == nil {
		return ""
	}
	return o.InvitationOrgName
}

func (o *CheckInviteTokenResponseBody) GetInvitationOrgLogoURL() *string {
	if o == nil {
		return nil
	}
	return o.InvitationOrgLogoURL
}

func (o *CheckInviteTokenResponseBody) GetInvitationOrgLogoThumbnailURL() *string {
	if o == nil {
		return nil
	}
	return o.InvitationOrgLogoThumbnailURL
}

func (o *CheckInviteTokenResponseBody) GetInviteeUserID() string {
	if o == nil {
		return ""
	}
	return o.InviteeUserID
}

func (o *CheckInviteTokenResponseBody) GetInviteePrimaryOrgID() *string {
	if o == nil {
		return nil
	}
	return o.InviteePrimaryOrgID
}

type CheckInviteTokenResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Invite token found
	Object *CheckInviteTokenResponseBody
}

func (o *CheckInviteTokenResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CheckInviteTokenResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CheckInviteTokenResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CheckInviteTokenResponse) GetObject() *CheckInviteTokenResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
