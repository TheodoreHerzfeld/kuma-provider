package provider

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/carlmjohnson/requests"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type data_usersAuth struct {
	Host  string
	Token string
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &data_usersAuth{}
	_ datasource.DataSourceWithConfigure = &data_usersAuth{}
)

// NewUserDataSource is a helper function to simplify the provider implementation.
func NewUsersDataSource() datasource.DataSource {
	return &data_usersAuth{}
}

// Configure adds the provider configured client to the data source.
func (d *data_usersAuth) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
func (d *data_usersAuth) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_users"
}

// Schema defines the schema for the data source.
func (d *data_usersAuth) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"users": schema.SetNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Computed: true,
						},
						"username": schema.StringAttribute{
							Computed: true,
						},
						"created_at": schema.StringAttribute{
							Computed: true,
						},
						"last_visit": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
		},
	}
}

type usersDataModel struct {
	Users []userDataModel `tfsdk:"users"`
}

// Read refreshes the Terraform state with the latest data.
func (d *data_usersAuth) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {

	var state usersDataModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var response string
	err := requests.
		URL(d.Host).
		Bearer(d.Token).
		Path("/users/").
		ToString(&response).
		Fetch(ctx)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error with response",
			"got "+err.Error(),
		)
	}

	var users []JSON_userDataModel
	err = json.Unmarshal([]byte(response), &users)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error unmarshalling response",
			err.Error(),
		)
	}

	var tout usersDataModel
	for _, user := range users {
		tout.Users = append(tout.Users, userDataModel{
			ID:         types.Int64Value(user.ID),
			Username:   types.StringValue(user.Username),
			Created_At: types.StringValue(user.Created_At),
			Last_Visit: types.StringValue(user.Last_Visit),
		})
	}

	diags = resp.State.Set(ctx, &tout)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

}
