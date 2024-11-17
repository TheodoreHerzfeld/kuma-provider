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

type data_tagAuth struct {
	Host  string
	Token string
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &data_tagAuth{}
	_ datasource.DataSourceWithConfigure = &data_tagAuth{}
)

// NewUserDataSource is a helper function to simplify the provider implementation.
func NewTagDataSource() datasource.DataSource {
	return &data_tagAuth{}
}

// Configure adds the provider configured client to the data source.
func (d *data_tagAuth) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
func (d *data_tagAuth) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tag"
}

// Schema defines the schema for the data source.
func (d *data_tagAuth) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Required: true,
			},
			"name": schema.StringAttribute{
				Computed: true,
			},
			"color": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

type tagDataModel struct {
	ID    types.Int64  `tfsdk:"id"`
	Name  types.String `tfsdk:"name"`
	Color types.String `tfsdk:"color"`
}

type JSON_tagDataModel struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

type tagResponse struct {
	Tag JSON_tagDataModel `json:"tag"`
}

// Read refreshes the Terraform state with the latest data.
func (d *data_tagAuth) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {

	var state tagDataModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	id := cleanString(state.ID.String())

	tflog.Debug(ctx, "Requesting "+d.Host+"/tags/"+id)

	var responseString string
	err := requests.
		URL(d.Host).
		Bearer(d.Token).
		Path("/tags/" + id).
		ToString(&responseString).
		Fetch(ctx)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error with response",
			"got "+err.Error(),
		)
	}

	var response tagResponse
	err = json.Unmarshal([]byte(responseString), &response)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error unmarshalling response",
			err.Error(),
		)
	}

	tflog.Debug(ctx, "Got tag data: "+responseString)

	state.ID = types.Int64Value(response.Tag.ID)
	state.Name = types.StringValue(response.Tag.Name)
	state.Color = types.StringValue(response.Tag.Color)

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

}
