// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type OrganizationType string

const (
	OrganizationTypeVendor  OrganizationType = "Vendor"
	OrganizationTypePartner OrganizationType = "Partner"
)

func (e OrganizationType) ToPointer() *OrganizationType {
	return &e
}
func (e *OrganizationType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Vendor":
		fallthrough
	case "Partner":
		*e = OrganizationType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OrganizationType: %v", v)
	}
}

type Address struct {
	Country      *string `json:"country,omitempty"`
	City         *string `json:"city,omitempty"`
	PostalCode   *string `json:"postal_code,omitempty"`
	Street       *string `json:"street,omitempty"`
	StreetNumber *string `json:"street_number,omitempty"`
}

func (o *Address) GetCountry() *string {
	if o == nil {
		return nil
	}
	return o.Country
}

func (o *Address) GetCity() *string {
	if o == nil {
		return nil
	}
	return o.City
}

func (o *Address) GetPostalCode() *string {
	if o == nil {
		return nil
	}
	return o.PostalCode
}

func (o *Address) GetStreet() *string {
	if o == nil {
		return nil
	}
	return o.Street
}

func (o *Address) GetStreetNumber() *string {
	if o == nil {
		return nil
	}
	return o.StreetNumber
}

type Organization struct {
	ID               *string           `json:"id,omitempty"`
	Type             *OrganizationType `json:"type,omitempty"`
	Name             *string           `json:"name,omitempty"`
	Signature        *string           `json:"signature,omitempty"`
	Symbol           *string           `json:"symbol,omitempty"`
	PricingTier      *string           `json:"pricing_tier,omitempty"`
	Email            *string           `json:"email,omitempty"`
	Phone            *string           `json:"phone,omitempty"`
	Website          *string           `json:"website,omitempty"`
	Address          *Address          `json:"address,omitempty"`
	LogoURL          *string           `json:"logo_url,omitempty"`
	LogoThumbnailURL *string           `json:"logo_thumbnail_url,omitempty"`
	IsUnlicensedOrg  *bool             `json:"is_unlicensed_org,omitempty"`
	CognitoDetails   *CognitoDetails   `json:"cognito_details,omitempty"`
}

func (o *Organization) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Organization) GetType() *OrganizationType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *Organization) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Organization) GetSignature() *string {
	if o == nil {
		return nil
	}
	return o.Signature
}

func (o *Organization) GetSymbol() *string {
	if o == nil {
		return nil
	}
	return o.Symbol
}

func (o *Organization) GetPricingTier() *string {
	if o == nil {
		return nil
	}
	return o.PricingTier
}

func (o *Organization) GetEmail() *string {
	if o == nil {
		return nil
	}
	return o.Email
}

func (o *Organization) GetPhone() *string {
	if o == nil {
		return nil
	}
	return o.Phone
}

func (o *Organization) GetWebsite() *string {
	if o == nil {
		return nil
	}
	return o.Website
}

func (o *Organization) GetAddress() *Address {
	if o == nil {
		return nil
	}
	return o.Address
}

func (o *Organization) GetLogoURL() *string {
	if o == nil {
		return nil
	}
	return o.LogoURL
}

func (o *Organization) GetLogoThumbnailURL() *string {
	if o == nil {
		return nil
	}
	return o.LogoThumbnailURL
}

func (o *Organization) GetIsUnlicensedOrg() *bool {
	if o == nil {
		return nil
	}
	return o.IsUnlicensedOrg
}

func (o *Organization) GetCognitoDetails() *CognitoDetails {
	if o == nil {
		return nil
	}
	return o.CognitoDetails
}
