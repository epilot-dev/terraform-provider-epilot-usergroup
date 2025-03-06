// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"github.com/epilot-dev/terraform-provider-epilot-usergroup/internal/sdk"
	"github.com/epilot-dev/terraform-provider-epilot-usergroup/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"net/http"
)

var _ provider.Provider = &EpilotUsergroupProvider{}

type EpilotUsergroupProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// EpilotUsergroupProviderModel describes the provider data model.
type EpilotUsergroupProviderModel struct {
	EpilotAuth types.String `tfsdk:"epilot_auth"`
	ServerURL  types.String `tfsdk:"server_url"`
}

func (p *EpilotUsergroupProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "epilot-usergroup"
	resp.Version = p.version
}

func (p *EpilotUsergroupProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"epilot_auth": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
			},
			"server_url": schema.StringAttribute{
				Description: `Server URL (defaults to https://user.sls.epilot.io)`,
				Optional:    true,
			},
		},
		MarkdownDescription: `User API: Manage users in epilot organization(s)`,
	}
}

func (p *EpilotUsergroupProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data EpilotUsergroupProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	ServerURL := data.ServerURL.ValueString()

	if ServerURL == "" {
		ServerURL = "https://user.sls.epilot.io"
	}

	epilotAuth := new(string)
	if !data.EpilotAuth.IsUnknown() && !data.EpilotAuth.IsNull() {
		*epilotAuth = data.EpilotAuth.ValueString()
	} else {
		epilotAuth = nil
	}
	security := shared.Security{
		EpilotAuth: epilotAuth,
	}

	providerHTTPTransportOpts := ProviderHTTPTransportOpts{
		SetHeaders: make(map[string]string),
		Transport:  http.DefaultTransport,
	}

	httpClient := http.DefaultClient
	httpClient.Transport = NewProviderHTTPTransport(providerHTTPTransportOpts)

	opts := []sdk.SDKOption{
		sdk.WithServerURL(ServerURL),
		sdk.WithSecurity(security),
		sdk.WithClient(httpClient),
	}
	client := sdk.New(opts...)

	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *EpilotUsergroupProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewUserGroupResource,
	}
}

func (p *EpilotUsergroupProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewUserGroupDataSource,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &EpilotUsergroupProvider{
			version: version,
		}
	}
}
