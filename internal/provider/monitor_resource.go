package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/carlmjohnson/requests"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int32default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
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

// Schema defines the schema for the resource.
func (r *monitorResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	defaultCodes, diags := types.SetValue(types.StringType, []attr.Value{types.StringValue("200")})
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

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
			"method": schema.StringAttribute{
				Required: false,
				Computed: true,
				Default:  stringdefault.StaticString("GET"),
			},
			"port": schema.Int32Attribute{
				Required: false,
				Computed: true,
				Default:  int32default.StaticInt32(80),
				// TODO: validate that this is reasonable
			},
			"interval": schema.Int32Attribute{
				Required: false,
				Computed: true,
				Default:  int32default.StaticInt32(60),
			},
			"max_retries": schema.Int32Attribute{
				Required: false,
				Computed: true,
				Default:  int32default.StaticInt32(0),
			},
			"retry_interval": schema.Int32Attribute{
				Required: false,
				Computed: true,
				Default:  int32default.StaticInt32(60),
			},
			"resend_interval": schema.Int32Attribute{
				Required: false,
				Computed: true,
				Default:  int32default.StaticInt32(0),
			},
			"upside_down": schema.BoolAttribute{
				Required: false,
				Computed: true,
				Default:  booldefault.StaticBool(false),
			},
			"notification_id_list": schema.SetAttribute{
				ElementType: types.StringType,
				Required:    false,
				Optional:    true,
			},
			"expiry_notification": schema.BoolAttribute{
				Required: false,
				Computed: true,
				Default:  booldefault.StaticBool(false),
			},
			"ignore_tls": schema.BoolAttribute{
				Required: false,
				Computed: true,
				Default:  booldefault.StaticBool(false),
			},
			"max_redirects": schema.Int32Attribute{
				Required: false,
				Computed: true,
				Default:  int32default.StaticInt32(10),
			},
			"accepted_statuscodes": schema.SetAttribute{
				ElementType: types.StringType,
				Required:    false,
				Computed:    true,
				Default:     setdefault.StaticValue(defaultCodes),
			},
			"proxy_id": schema.Int32Attribute{
				Required: false,
				Optional: true,
			},
			"body": schema.StringAttribute{
				Required: false,
				Optional: true,
			},
			"headers": schema.StringAttribute{
				Required: false,
				Optional: true,
			},
			"auth_method": schema.StringAttribute{
				Required: false,
				Optional: true,
			},
			"basic_auth_user": schema.StringAttribute{
				Required: false,
				Optional: true,
			},
			"basic_auth_pass": schema.StringAttribute{
				Required:  false,
				Optional:  true,
				Sensitive: true,
			},
			"auth_domain": schema.StringAttribute{
				Required: false,
				Optional: true,
			},
			"auth_workstation": schema.StringAttribute{
				Required: false,
				Optional: true,
			},
			"keyword": schema.StringAttribute{
				Required: false,
				Optional: true,
			},
			"hostname": schema.StringAttribute{
				Required: false,
				Optional: true,
			},
			"dns_resolve_server": schema.StringAttribute{
				Required: false,
				Optional: true,
			},
			"dns_resolve_type": schema.StringAttribute{
				Required: false,
				Optional: true,
			},
			"mqtt_username": schema.StringAttribute{
				Required: false,
				Optional: true,
			},
			"mqtt_password": schema.StringAttribute{
				Required:  false,
				Optional:  true,
				Sensitive: true,
			},
			"mqtt_topic": schema.StringAttribute{
				Required: false,
				Optional: true,
			},
			"mqtt_success_message": schema.StringAttribute{
				Required: false,
				Optional: true,
			},
			"database_connection_string": schema.StringAttribute{
				Required:  false,
				Optional:  true,
				Sensitive: true,
			},
			"database_query": schema.StringAttribute{
				Required: false,
				Optional: true,
			},
			"docker_container": schema.StringAttribute{
				Required: false,
				Optional: true,
			},
			"docker_host": schema.Int32Attribute{
				Required: false,
				Optional: true,
			},
			"radius_username": schema.StringAttribute{
				Required: false,
				Optional: true,
			},
			"radius_password": schema.StringAttribute{
				Required:  false,
				Optional:  true,
				Sensitive: true,
			},
			"radius_secret": schema.StringAttribute{
				Required:  false,
				Optional:  true,
				Sensitive: true,
			},
			"radius_called_station_id": schema.StringAttribute{
				Required: false,
				Optional: true,
			},
			"radius_calling_station_id": schema.StringAttribute{
				Required: false,
				Optional: true,
			},
		},
	}
}

type monitorModel struct {
	ID                       types.Int64  `tfsdk:"id"`
	Type                     types.String `tfsdk:"type"`
	Name                     types.String `tfsdk:"name"`
	Interval                 types.Int32  `tfsdk:"interval"`
	RetryInterval            types.Int32  `tfsdk:"retry_interval"`
	ResendInterval           types.Int32  `tfsdk:"resend_interval"`
	MaxRetries               types.Int32  `tfsdk:"max_retries"`
	UpsideDown               types.Bool   `tfsdk:"upside_down"`
	NotificationIDList       types.Set    `tfsdk:"notification_id_list"`
	URL                      types.String `tfsdk:"url"`
	ExpiryNotification       types.Bool   `tfsdk:"expiry_notification"`
	IgnoreTls                types.Bool   `tfsdk:"ignore_tls"`
	MaxRedirects             types.Int32  `tfsdk:"max_redirects"`
	AcceptedStatusCodes      types.Set    `tfsdk:"accepted_statuscodes"`
	ProxyID                  types.Int32  `tfsdk:"proxy_id"`
	Method                   types.String `tfsdk:"method"`
	Body                     types.String `tfsdk:"body"`
	Headers                  types.String `tfsdk:"headers"`
	AuthMethod               types.String `tfsdk:"auth_method"`
	BasicAuthUser            types.String `tfsdk:"basic_auth_user"`
	BasicAuthPass            types.String `tfsdk:"basic_auth_pass"`
	AuthDomain               types.String `tfsdk:"auth_domain"`
	AuthWorkstation          types.String `tfsdk:"auth_workstation"`
	Keyword                  types.String `tfsdk:"keyword"`
	Hostname                 types.String `tfsdk:"hostname"`
	Port                     types.Int32  `tfsdk:"port"`
	DNSResolveServer         types.String `tfsdk:"dns_resolve_server"`
	DNSResolveType           types.String `tfsdk:"dns_resolve_type"`
	MQTTUsername             types.String `tfsdk:"mqtt_username"`
	MQTTPassword             types.String `tfsdk:"mqtt_password"`
	MQTTTopic                types.String `tfsdk:"mqtt_topic"`
	MQTTSucessMessage        types.String `tfsdk:"mqtt_success_message"`
	DatabaseConnectionString types.String `tfsdk:"database_connection_string"`
	DatabaseQuery            types.String `tfsdk:"database_query"`
	DockerContainer          types.String `tfsdk:"docker_container"`
	DockerHost               types.Int32  `tfsdk:"docker_host"`
	RadiusUsername           types.String `tfsdk:"radius_username"`
	RadiusPassword           types.String `tfsdk:"radius_password"`
	RadiusSecret             types.String `tfsdk:"radius_secret"`
	RadiusCalledStationId    types.String `tfsdk:"radius_called_station_id"`
	RadiusCallingStationId   types.String `tfsdk:"radius_calling_station_id"`
}

type JSON_monitorModel struct {
	ID                       int      `json:"id"`
	Type                     string   `json:"type"`
	Name                     string   `json:"name"`
	Interval                 int32    `json:"interval"`
	RetryInterval            int32    `json:"retry_interval"`
	ResendInterval           int32    `json:"resend_interval"`
	MaxRetries               int32    `json:"max_retries"`
	UpsideDown               bool     `json:"upside_down"`
	NotificationIDList       []string `json:"notification_id_list"`
	URL                      string   `json:"url"`
	ExpiryNotification       bool     `json:"expiry_notification"`
	IgnoreTls                bool     `json:"ignore_tls"`
	MaxRedirects             int32    `json:"max_redirects"`
	AcceptedStatusCodes      []string `json:"accepted_statuscodes"`
	ProxyID                  int32    `json:"proxy_id"`
	Method                   string   `json:"method"`
	Body                     string   `json:"body"`
	Headers                  string   `json:"headers"`
	AuthMethod               string   `json:"auth_method"`
	BasicAuthUser            string   `json:"basic_auth_user"`
	BasicAuthPass            string   `json:"basic_auth_pass"`
	AuthDomain               string   `json:"auth_domain"`
	AuthWorkstation          string   `json:"auth_workstation"`
	Keyword                  string   `json:"keyword"`
	Hostname                 string   `json:"hostname"`
	Port                     int32    `json:"port"`
	DNSResolveServer         string   `json:"dns_resolve_server"`
	DNSResolveType           string   `json:"dns_resolve_type"`
	MQTTUsername             string   `json:"mqtt_username"`
	MQTTPassword             string   `json:"mqtt_password"`
	MQTTTopic                string   `json:"mqtt_topic"`
	MQTTSucessMessage        string   `json:"mqtt_success_message"`
	DatabaseConnectionString string   `json:"database_connection_string"`
	DatabaseQuery            string   `json:"database_query"`
	DockerContainer          string   `json:"docker_container"`
	DockerHost               int32    `json:"docker_host"`
	RadiusUsername           string   `json:"radius_username"`
	RadiusPassword           string   `json:"radius_password"`
	RadiusSecret             string   `json:"radius_secret"`
	RadiusCalledStationId    string   `json:"radius_called_station_id"`
	RadiusCallingStationId   string   `json:"radius_calling_station_id"`
}

func cleanString(dirty string) string {

	tout := strings.Trim(dirty, "\"")
	tout = strings.TrimLeft(tout, "<")
	tout = strings.TrimRight(tout, ">")

	return tout
}

// Create creates the resource and sets the initial Terraform state.
func (r *monitorResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan monitorModel

	tflog.Debug(ctx, "STAGE: map plan")
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var notificationIDs, acceptedStatusCodes []string

	diags = plan.NotificationIDList.ElementsAs(ctx, &notificationIDs, false)
	resp.Diagnostics.Append(diags...)
	diags = plan.AcceptedStatusCodes.ElementsAs(ctx, &notificationIDs, false)
	resp.Diagnostics.Append(diags...)

	tflog.Debug(ctx, "STAGE: map json representation")

	makeMon := JSON_monitorModel{
		Type:                     cleanString(plan.Type.String()),
		Name:                     cleanString(plan.Name.String()),
		Interval:                 plan.Interval.ValueInt32(),
		RetryInterval:            plan.RetryInterval.ValueInt32(),
		ResendInterval:           plan.ResendInterval.ValueInt32(),
		MaxRetries:               plan.MaxRetries.ValueInt32(),
		UpsideDown:               plan.UpsideDown.ValueBool(),
		NotificationIDList:       notificationIDs,
		URL:                      cleanString(plan.URL.String()),
		ExpiryNotification:       plan.ExpiryNotification.ValueBool(),
		IgnoreTls:                plan.IgnoreTls.ValueBool(),
		MaxRedirects:             plan.MaxRedirects.ValueInt32(),
		AcceptedStatusCodes:      acceptedStatusCodes,
		ProxyID:                  plan.ProxyID.ValueInt32(),
		Method:                   cleanString(plan.Method.String()),
		Body:                     cleanString(plan.Body.String()),
		Headers:                  cleanString(plan.Headers.String()),
		AuthMethod:               cleanString(plan.AuthMethod.String()),
		BasicAuthUser:            cleanString(plan.BasicAuthUser.String()),
		BasicAuthPass:            cleanString(plan.BasicAuthPass.String()),
		AuthDomain:               cleanString(plan.AuthDomain.String()),
		AuthWorkstation:          cleanString(plan.AuthWorkstation.String()),
		Keyword:                  cleanString(plan.Keyword.String()),
		Hostname:                 cleanString(plan.Hostname.String()),
		Port:                     plan.Port.ValueInt32(),
		DNSResolveServer:         cleanString(plan.DNSResolveServer.String()),
		DNSResolveType:           cleanString(plan.DNSResolveType.String()),
		MQTTUsername:             cleanString(plan.MQTTUsername.String()),
		MQTTPassword:             cleanString(plan.MQTTPassword.String()),
		MQTTTopic:                cleanString(plan.MQTTTopic.String()),
		MQTTSucessMessage:        cleanString(plan.MQTTSucessMessage.String()),
		DatabaseConnectionString: cleanString(plan.DatabaseConnectionString.String()),
		DatabaseQuery:            cleanString(plan.DatabaseQuery.String()),
		DockerContainer:          cleanString(plan.DockerContainer.String()),
		DockerHost:               plan.DockerHost.ValueInt32(),
		RadiusUsername:           cleanString(plan.RadiusUsername.String()),
		RadiusPassword:           cleanString(plan.RadiusPassword.String()),
		RadiusSecret:             cleanString(plan.RadiusSecret.String()),
		RadiusCalledStationId:    cleanString(plan.RadiusCalledStationId.String()),
		RadiusCallingStationId:   cleanString(plan.RadiusCallingStationId.String()),
	}

	debugJSON, err := json.Marshal(makeMon)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating new monitor (json encode)",
			"what we know: "+err.Error(),
		)
		return
	}

	tflog.Debug(ctx, "MONITOR JSON: "+string(debugJSON))

	var newMon JSON_monitorModel
	err = requests.
		URL(r.Host).
		Bearer(r.Token).
		Pathf("/monitors").
		BodyJSON(&makeMon).
		ToJSON(&newMon).
		Fetch(ctx)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating new monitor (api call)",
			"what we know: "+err.Error(),
		)
		return
	}

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

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
