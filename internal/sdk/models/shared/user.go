// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/epilot-dev/terraform-provider-epilot-usergroup/internal/sdk/internal/utils"
)

type ImageURI struct {
	Original             *string `json:"original,omitempty"`
	Thumbnail32          *string `json:"thumbnail_32,omitempty"`
	AdditionalProperties any     `additionalProperties:"true" json:"-"`
}

func (i ImageURI) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *ImageURI) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ImageURI) GetOriginal() *string {
	if o == nil {
		return nil
	}
	return o.Original
}

func (o *ImageURI) GetThumbnail32() *string {
	if o == nil {
		return nil
	}
	return o.Thumbnail32
}

func (o *ImageURI) GetAdditionalProperties() any {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

type Properties struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (o *Properties) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Properties) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

type User struct {
	// User's unique identifier
	ID             string `json:"id"`
	OrganizationID string `json:"organization_id"`
	Email          string `json:"email"`
	// Deprecated: This will be removed in a future release, please migrate away from it as soon as possible.
	DisplayName *string `json:"display_name,omitempty"`
	// Deprecated: This will be removed in a future release, please migrate away from it as soon as possible.
	Name              string  `json:"name"`
	PreferredLanguage string  `json:"preferred_language"`
	Signature         *string `json:"signature,omitempty"`
	// Deprecated! Please use Permissions API instead
	//
	// Deprecated: This will be removed in a future release, please migrate away from it as soon as possible.
	Roles      []string     `json:"roles"`
	ImageURI   *ImageURI    `json:"image_uri,omitempty"`
	Properties []Properties `json:"properties"`
}

func (o *User) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *User) GetOrganizationID() string {
	if o == nil {
		return ""
	}
	return o.OrganizationID
}

func (o *User) GetEmail() string {
	if o == nil {
		return ""
	}
	return o.Email
}

func (o *User) GetDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.DisplayName
}

func (o *User) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *User) GetPreferredLanguage() string {
	if o == nil {
		return ""
	}
	return o.PreferredLanguage
}

func (o *User) GetSignature() *string {
	if o == nil {
		return nil
	}
	return o.Signature
}

func (o *User) GetRoles() []string {
	if o == nil {
		return []string{}
	}
	return o.Roles
}

func (o *User) GetImageURI() *ImageURI {
	if o == nil {
		return nil
	}
	return o.ImageURI
}

func (o *User) GetProperties() []Properties {
	if o == nil {
		return []Properties{}
	}
	return o.Properties
}
