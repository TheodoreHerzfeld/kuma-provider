package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/carlmjohnson/requests"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func cleanString(dirty string) string {

	tout := strings.Trim(dirty, "\"")
	tout = strings.TrimLeft(tout, "<")
	tout = strings.TrimRight(tout, ">")

	return tout
}

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
	// defaultCodes, diags := types.SetValue(types.StringType, []attr.Value{types.StringValue("200")})
	// resp.Diagnostics.Append(diags...)
	// if resp.Diagnostics.HasError() {
	// 	return
	// }

	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Computed: true,
			},
			"type": schema.StringAttribute{
				Required: true,
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"interval": schema.Int64Attribute{
				Required: false,
				Optional: true,
			},
			"retry_interval": schema.Int64Attribute{
				Optional: true,
			},
			"resend_interval": schema.Int64Attribute{
				Optional: true,
			},
			"max_retries": schema.Int64Attribute{
				Optional: true,
			},
			"upside_down": schema.BoolAttribute{
				Optional: true,
			},
			"notification_id_list": schema.SetAttribute{
				ElementType: types.Int64Type,
				Optional:    true,
			},
			"url": schema.StringAttribute{
				Required: true,
			},
			"expiry_notification": schema.BoolAttribute{
				Optional: true,
			},
			"ignore_tls": schema.BoolAttribute{
				Optional: true,
			},
			"max_redirects": schema.Int64Attribute{
				Optional: true,
			},
			"accepted_statuscodes": schema.SetAttribute{
				ElementType: types.StringType,
				Optional:    true,
			},
			"proxy_id": schema.Int64Attribute{
				Optional: true,
			},
			"method": schema.StringAttribute{
				Optional: true,
			},
			"body": schema.StringAttribute{
				Optional: true,
			},
			"headers": schema.StringAttribute{
				Optional: true,
			},
			"auth_method": schema.StringAttribute{
				Optional: true,
			},
			"basic_auth_user": schema.StringAttribute{
				Optional: true,
			},
			"basic_auth_pass": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
			},
			"auth_domain": schema.StringAttribute{
				Optional: true,
			},
			"auth_workstation": schema.StringAttribute{
				Optional: true,
			},
			"keyword": schema.StringAttribute{
				Optional: true,
			},
			"hostname": schema.StringAttribute{
				Optional: true,
			},
			"port": schema.Int64Attribute{
				Optional: true,
			},
			"dns_resolve_server": schema.StringAttribute{
				Optional: true,
			},
			"dns_resolve_type": schema.StringAttribute{
				Optional: true,
			},
			"mqtt_username": schema.StringAttribute{
				Optional: true,
			},
			"mqtt_password": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
			},
			"mqtt_topic": schema.StringAttribute{
				Optional: true,
			},
			"mqtt_success_message": schema.StringAttribute{
				Optional: true,
			},
			"database_connection_string": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
			},
			"database_query": schema.StringAttribute{
				Optional: true,
			},
			"docker_container": schema.StringAttribute{
				Optional: true,
			},
			"docker_host": schema.Int64Attribute{
				Optional: true,
			},
			"radius_username": schema.StringAttribute{
				Optional: true,
			},
			"radius_password": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
			},
			"radius_secret": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
			},
			"radius_called_station_id": schema.StringAttribute{
				Optional: true,
			},
			"radius_calling_station_id": schema.StringAttribute{
				Optional: true,
			},
			"active": schema.BoolAttribute{
				Optional: true,
			},
			"force_inactive": schema.BoolAttribute{
				Optional: true,
			},
			"game": schema.StringAttribute{
				Optional: true,
			},
			"gamedig_given_port_only": schema.BoolAttribute{
				Optional: true,
			},
			"grpc_body": schema.StringAttribute{
				Optional: true,
			},
			"grpc_enable_tls": schema.BoolAttribute{
				Optional: true,
			},
			"grpc_metadata": schema.StringAttribute{
				Optional: true,
			},
			"grpc_method": schema.StringAttribute{
				Optional: true,
			},
			"grpc_protobuf": schema.StringAttribute{
				Optional: true,
			},
			"grpc_service_name": schema.StringAttribute{
				Optional: true,
			},
			"grpc_url": schema.StringAttribute{
				Optional: true,
			},
			"http_body_encoding": schema.StringAttribute{
				Optional: true,
			},
			"include_sensitive_data": schema.BoolAttribute{
				Optional: true,
			},
			"invert_keyword": schema.BoolAttribute{
				Optional: true,
			},
			"json_path": schema.StringAttribute{
				Optional: true,
			},
			"kafka_producer_allow_auto_topic_creation": schema.BoolAttribute{
				Optional: true,
			},
			"kafka_producer_brokers": schema.SetAttribute{
				ElementType: types.StringType,
				Optional:    true,
			},
			"kafka_producer_message": schema.StringAttribute{
				Optional: true,
			},
			"kafka_producer_sasl_options": schema.StringAttribute{
				Optional: true,
			},
			"kafka_producer_ssl": schema.BoolAttribute{
				Optional: true,
			},
			"kafka_producer_topic": schema.StringAttribute{
				Optional: true,
			},
			"maintenance": schema.BoolAttribute{
				Optional: true,
			},
			"oauth_auth_method": schema.StringAttribute{
				Optional: true,
			},
			"oauth_client_id": schema.StringAttribute{
				Optional: true,
			},
			"oauth_client_secret": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
			},
			"oauth_scopes": schema.SetAttribute{
				ElementType: types.StringType,
				Optional:    true,
			},
			"oauth_token_url": schema.StringAttribute{
				Optional: true,
			},
			"packet_size": schema.Int64Attribute{
				Optional: true,
			},
			"parent": schema.StringAttribute{
				Optional: true,
			},
			"path_name": schema.StringAttribute{
				Optional: true,
			},
			"push_token": schema.StringAttribute{
				Optional: true,
			},
			"screenshot": schema.StringAttribute{
				Optional: true,
			},
			"tags": schema.SetAttribute{
				ElementType: types.StringType,
				Optional:    true,
			},
			"timeout": schema.Int64Attribute{
				Optional: true,
			},
			"tls_ca": schema.StringAttribute{
				Optional: true,
			},
			"tls_cert": schema.StringAttribute{
				Optional: true,
			},
			"tls_key": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
			},
			"weight": schema.Int64Attribute{
				Optional: true,
			},
		},
	}
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

	tflog.Debug(ctx, "STAGE: map json representation - NAME:"+plan.Name.String()+"|"+cleanString(plan.Name.String()))

	makeMon := JSON_monitorModel{
		Type:                     cleanString(plan.Type.String()),
		Name:                     cleanString(plan.Name.String()),
		Interval:                 plan.Interval.ValueInt64(),
		RetryInterval:            plan.RetryInterval.ValueInt64(),
		ResendInterval:           plan.ResendInterval.ValueInt64(),
		MaxRetries:               plan.MaxRetries.ValueInt64(),
		UpsideDown:               plan.UpsideDown.ValueBool(),
		NotificationIDList:       notificationIDs,
		URL:                      cleanString(plan.URL.String()),
		ExpiryNotification:       plan.ExpiryNotification.ValueBool(),
		IgnoreTls:                plan.IgnoreTls.ValueBool(),
		MaxRedirects:             plan.MaxRedirects.ValueInt64(),
		AcceptedStatusCodes:      acceptedStatusCodes,
		ProxyID:                  plan.ProxyID.ValueInt64(),
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
		Port:                     plan.Port.ValueInt64(),
		DNSResolveServer:         cleanString(plan.DNSResolveServer.String()),
		DNSResolveType:           cleanString(plan.DNSResolveType.String()),
		MQTTUsername:             cleanString(plan.MQTTUsername.String()),
		MQTTPassword:             cleanString(plan.MQTTPassword.String()),
		MQTTTopic:                cleanString(plan.MQTTTopic.String()),
		MQTTSucessMessage:        cleanString(plan.MQTTSucessMessage.String()),
		DatabaseConnectionString: cleanString(plan.DatabaseConnectionString.String()),
		DatabaseQuery:            cleanString(plan.DatabaseQuery.String()),
		DockerContainer:          cleanString(plan.DockerContainer.String()),
		DockerHost:               plan.DockerHost.ValueInt64(),
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

	tflog.Debug(ctx, "NEW MONITOR JSON: "+string(debugJSON))

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

	tflog.Debug(ctx, "STAGE: map new monitor onto schema")

	newNotificationIDList, diags := types.SetValueFrom(ctx, types.StringType, newMon.NotificationIDList)
	resp.Diagnostics.Append(diags...)
	newAcceptedStatusCodes, diags := types.SetValueFrom(ctx, types.StringType, newMon.AcceptedStatusCodes)
	resp.Diagnostics.Append(diags...)

	resultMon := monitorModel{
		ID:                       types.Int64Value(newMon.ID),
		Type:                     types.StringValue(newMon.Type),
		Name:                     types.StringValue(newMon.Name),
		Interval:                 types.Int64Value(newMon.Interval),
		RetryInterval:            types.Int64Value(newMon.RetryInterval),
		ResendInterval:           types.Int64Value(newMon.ResendInterval),
		MaxRetries:               types.Int64Value(newMon.MaxRetries),
		UpsideDown:               types.BoolValue(newMon.UpsideDown),
		NotificationIDList:       newNotificationIDList,
		URL:                      types.StringValue(newMon.URL),
		ExpiryNotification:       types.BoolValue(newMon.ExpiryNotification),
		IgnoreTls:                types.BoolValue(newMon.IgnoreTls),
		MaxRedirects:             types.Int64Value(newMon.MaxRedirects),
		AcceptedStatusCodes:      newAcceptedStatusCodes,
		ProxyID:                  types.Int64Value(newMon.ProxyID),
		Method:                   types.StringValue(newMon.Method),
		Body:                     types.StringValue(newMon.Body),
		Headers:                  types.StringValue(newMon.Headers),
		AuthMethod:               types.StringValue(newMon.AuthMethod),
		BasicAuthUser:            types.StringValue(newMon.BasicAuthUser),
		BasicAuthPass:            types.StringValue(newMon.BasicAuthPass),
		AuthDomain:               types.StringValue(newMon.AuthDomain),
		AuthWorkstation:          types.StringValue(newMon.AuthWorkstation),
		Keyword:                  types.StringValue(newMon.Keyword),
		Hostname:                 types.StringValue(newMon.Hostname),
		Port:                     types.Int64Value(newMon.Port),
		DNSResolveServer:         types.StringValue(newMon.DNSResolveServer),
		DNSResolveType:           types.StringValue(newMon.DNSResolveType),
		MQTTUsername:             types.StringValue(newMon.MQTTUsername),
		MQTTPassword:             types.StringValue(newMon.MQTTPassword),
		MQTTTopic:                types.StringValue(newMon.MQTTTopic),
		MQTTSucessMessage:        types.StringValue(newMon.MQTTSucessMessage),
		DatabaseConnectionString: types.StringValue(newMon.DatabaseConnectionString),
		DatabaseQuery:            types.StringValue(newMon.DatabaseQuery),
		DockerContainer:          types.StringValue(newMon.DockerContainer),
		DockerHost:               types.Int64Value(newMon.DockerHost),
		RadiusUsername:           types.StringValue(newMon.RadiusUsername),
		RadiusPassword:           types.StringValue(newMon.RadiusPassword),
		RadiusSecret:             types.StringValue(newMon.RadiusSecret),
		RadiusCalledStationId:    types.StringValue(newMon.RadiusCalledStationId),
		RadiusCallingStationId:   types.StringValue(newMon.RadiusCallingStationId),
	}

	diags = resp.State.Set(ctx, resultMon)
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
