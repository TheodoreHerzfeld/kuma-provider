package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/carlmjohnson/requests"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &authData{}
	_ datasource.DataSourceWithConfigure = &authData{}
)

// NewUserDataSource is a helper function to simplify the provider implementation.
func NewUserDataSource() datasource.DataSource {
	return &authData{}
}

// Configure adds the provider configured client to the data source.
func (d *authData) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
func (d *authData) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_user"
}

// Schema defines the schema for the data source.
func (d *authData) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Computed: true,
			},
			"username": schema.StringAttribute{
				Computed: false,
				Required: true,
			},
			"created_at": schema.StringAttribute{
				Computed: true,
			},
			"last_visit": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

type userDataModel struct {
	ID         types.Int64  `tfsdk:"id"`
	Username   types.String `tfsdk:"username"`
	Created_At types.String `tfsdk:"created_at"`
	Last_Visit types.String `tfsdk:"last_visit"`
}

type JSON_userDataModel struct {
	ID         int64  `json:"id"`
	Username   string `json:"username"`
	Created_At string `json:"created_at"`
	Last_Visit string `json:"last_visit"`
}

// Read refreshes the Terraform state with the latest data.
func (d *authData) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {

	var state userDataModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	username := strings.ReplaceAll(state.Username.ValueString(), "\"", "")

	var tout JSON_userDataModel
	err := requests.
		URL(d.Host).
		Bearer(d.Token).
		Pathf("/users/" + username).
		ToJSON(&tout).
		Fetch(ctx)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error with response",
			"got "+err.Error(),
		)
	}

	state.ID = types.Int64Value(tout.ID)
	state.Created_At = types.StringValue(tout.Created_At)
	state.Last_Visit = types.StringValue(tout.Last_Visit)

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

}
