package provider

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int32default"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource              = &monitorResource{}
	_ resource.ResourceWithConfigure = &monitorResource{}
)

// NewOrderResource is a helper function to simplify the provider implementation.
func NewMonitorResource() resource.Resource {
	return &monitorResource{}
}

// monitorResource is the resource implementation.
type monitorResource struct {
	Host  string
	Token string
}

type resourceAuthData struct {
	Host  string
	Token string
}

// Configure adds the provider configured client to the data source.
func (d *monitorResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

	tflog.Info(ctx, "Received a good token! @monitor_resource.configure")
}

// Metadata returns the resource type name.
func (r *monitorResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_monitor"
}

// {
// 	"type": "http",
// 	"name": "string",
// 	"interval": 60,
// 	"retryInterval": 60,
// 	"resendInterval": 0,
// 	"maxretries": 0,
// 	"upsideDown": false,
// 	"notificationIDList": [
// 	  "string"
// 	],
// 	"url": "string",
// 	"expiryNotification": false,
// 	"ignoreTls": false,
// 	"maxredirects": 10,
// 	"accepted_statuscodes": [
// 	  "string"
// 	],
// 	"proxyId": 0,
// 	"method": "GET",
// 	"body": "string",
// 	"headers": "string",
// 	"authMethod": "",
// 	"basic_auth_user": "string",
// 	"basic_auth_pass": "string",
// 	"authDomain": "string",
// 	"authWorkstation": "string",
// 	"keyword": "string",
// 	"hostname": "string",
// 	"port": 53,
// 	"dns_resolve_server": "1.1.1.1",
// 	"dns_resolve_type": "A",
// 	"mqttUsername": "string",
// 	"mqttPassword": "string",
// 	"mqttTopic": "string",
// 	"mqttSuccessMessage": "string",
// 	"databaseConnectionString": "string",
// 	"databaseQuery": "string",
// 	"docker_container": "",
// 	"docker_host": 0,
// 	"radiusUsername": "string",
// 	"radiusPassword": "string",
// 	"radiusSecret": "string",
// 	"radiusCalledStationId": "string",
// 	"radiusCallingStationId": "string"
//   }

// Schema defines the schema for the resource.
func (r *monitorResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"type": schema.StringAttribute{
				Required: true,
			},
			"url": schema.StringAttribute{
				Required: true,
			},
			"port": schema.Int32Attribute{
				Required: false,
				Default:  int32default.StaticInt32(80),
				// TODO: validate that this is reasonable
			},
			"interval": schema.Int32Attribute{
				Required: false,
				Default:  int32default.StaticInt32(60),
			},
			"maxRetries": schema.Int32Attribute{
				Required: false,
				Default:  int32default.StaticInt32(0),
			},
			"retryInterval": schema.Int32Attribute{
				Required: false,
				Default:  int32default.StaticInt32(60),
			},
			"upsideDown": schema.BoolAttribute{
				Required: false,
				Default:  booldefault.StaticBool(false),
			},
			"notificationIDList": 
			},
		},
	}
}

// Create creates the resource and sets the initial Terraform state.
func (r *monitorResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
}

// Read refreshes the Terraform state with the latest data.
func (r *monitorResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *monitorResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *monitorResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}
