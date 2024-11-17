package provider

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/carlmjohnson/requests"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

type data_serverInfoAuth struct {
	Host  string
	Token string
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &data_serverInfoAuth{}
	_ datasource.DataSourceWithConfigure = &data_serverInfoAuth{}
)

// NewUserDataSource is a helper function to simplify the provider implementation.
func NewServerInfoDataSource() datasource.DataSource {
	return &data_serverInfoAuth{}
}

// Configure adds the provider configured client to the data source.
func (d *data_serverInfoAuth) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Add a nil check when handling ProviderData because Terraform
	// sets that data after it calls the ConfigureProvider RPC.
	if req.ProviderData == nil {
		return
	}

	authbytes, ok := req.ProviderData.([]byte)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected []byte, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}
	json.Unmarshal(authbytes, &d)

}

// Metadata returns the data source type name.
func (d *data_serverInfoAuth) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_server_info"
}

// Schema defines the schema for the data source.
func (d *data_serverInfoAuth) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"primary_base_url": schema.StringAttribute{
				Computed: true,
			},
			"server_timezone": schema.StringAttribute{
				Computed: true,
			},
			"server_timezone_offset": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

type serverInfoModel struct {
	PrimaryBaseUrl       types.String `tfsdk:"primary_base_url"`
	ServerTimezone       types.String `tfsdk:"server_timezone"`
	ServerTimezoneOffset types.String `tfsdk:"server_timezone_offset"`
}

type JSON_serverInfoModel struct {
	PrimaryBaseUrl       string `json:"primaryBaseUrl"`
	ServerTimezone       string `json:"serverTimezone"`
	ServerTimezoneOffset string `json:"serverTimezoneOffset"`
}

// Read refreshes the Terraform state with the latest data.
func (d *data_serverInfoAuth) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {

	var state serverInfoModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Requesting "+d.Host+"/info/")

	var responseString string
	err := requests.
		URL(d.Host).
		Bearer(d.Token).
		Path("/info/").
		ToString(&responseString).
		Fetch(ctx)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error with response",
			"got "+err.Error(),
		)
	}

	var response JSON_serverInfoModel
	err = json.Unmarshal([]byte(responseString), &response)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error unmarshalling response",
			err.Error(),
		)
	}

	tflog.Debug(ctx, "Got tag data: "+responseString)

	state.PrimaryBaseUrl = types.StringValue(response.PrimaryBaseUrl)
	state.ServerTimezone = types.StringValue(response.ServerTimezone)
	state.ServerTimezoneOffset = types.StringValue(response.ServerTimezoneOffset)

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

}
