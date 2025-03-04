// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type UserActivationPayload struct {
	// User's display name (default: email address)
	DisplayName *string `json:"display_name,omitempty"`
	// User's password
	Password *string `json:"password,omitempty"`
}

func (o *UserActivationPayload) GetDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.DisplayName
}

func (o *UserActivationPayload) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}
