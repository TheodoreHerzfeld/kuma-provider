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

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &data_monitorAuth{}
	_ datasource.DataSourceWithConfigure = &data_monitorAuth{}
)

// NewUserDataSource is a helper function to simplify the provider implementation.
func NewMonitorDataSource() datasource.DataSource {
	return &data_monitorAuth{}
}

type data_monitorAuth struct {
	Host  string
	Token string
}

type tagInstanceDataModel struct {
	ID    types.Int64  `tfsdk:"id"`
	Name  types.String `tfsdk:"name"`
	Color types.String `tfsdk:"color"`
	Value types.String `tfsdk:"value"`
}

type JSON_tagInstanceDataModel struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
	Value string `json:"value"`
}

type monitorModel struct {
	ID                                  types.Int64            `tfsdk:"id"`
	Type                                types.String           `tfsdk:"type"`
	Name                                types.String           `tfsdk:"name"`
	Interval                            types.Int64            `tfsdk:"interval"`
	RetryInterval                       types.Int64            `tfsdk:"retry_interval"`
	ResendInterval                      types.Int64            `tfsdk:"resend_interval"`
	MaxRetries                          types.Int64            `tfsdk:"max_retries"`
	UpsideDown                          types.Bool             `tfsdk:"upside_down"`
	NotificationIDList                  types.Set              `tfsdk:"notification_id_list"`
	URL                                 types.String           `tfsdk:"url"`
	ExpiryNotification                  types.Bool             `tfsdk:"expiry_notification"`
	IgnoreTls                           types.Bool             `tfsdk:"ignore_tls"`
	MaxRedirects                        types.Int64            `tfsdk:"max_redirects"`
	AcceptedStatusCodes                 types.Set              `tfsdk:"accepted_statuscodes"`
	ProxyID                             types.Int64            `tfsdk:"proxy_id"`
	Method                              types.String           `tfsdk:"method"`
	Body                                types.String           `tfsdk:"body"`
	Headers                             types.String           `tfsdk:"headers"`
	AuthMethod                          types.String           `tfsdk:"auth_method"`
	BasicAuthUser                       types.String           `tfsdk:"basic_auth_user"`
	BasicAuthPass                       types.String           `tfsdk:"basic_auth_pass"`
	AuthDomain                          types.String           `tfsdk:"auth_domain"`
	AuthWorkstation                     types.String           `tfsdk:"auth_workstation"`
	Keyword                             types.String           `tfsdk:"keyword"`
	Hostname                            types.String           `tfsdk:"hostname"`
	Port                                types.Int64            `tfsdk:"port"`
	DNSResolveServer                    types.String           `tfsdk:"dns_resolve_server"`
	DNSResolveType                      types.String           `tfsdk:"dns_resolve_type"`
	MQTTUsername                        types.String           `tfsdk:"mqtt_username"`
	MQTTPassword                        types.String           `tfsdk:"mqtt_password"`
	MQTTTopic                           types.String           `tfsdk:"mqtt_topic"`
	MQTTSucessMessage                   types.String           `tfsdk:"mqtt_success_message"`
	DatabaseConnectionString            types.String           `tfsdk:"database_connection_string"`
	DatabaseQuery                       types.String           `tfsdk:"database_query"`
	DockerContainer                     types.String           `tfsdk:"docker_container"`
	DockerHost                          types.Int64            `tfsdk:"docker_host"`
	RadiusUsername                      types.String           `tfsdk:"radius_username"`
	RadiusPassword                      types.String           `tfsdk:"radius_password"`
	RadiusSecret                        types.String           `tfsdk:"radius_secret"`
	RadiusCalledStationId               types.String           `tfsdk:"radius_called_station_id"`
	RadiusCallingStationId              types.String           `tfsdk:"radius_calling_station_id"`
	Active                              types.Bool             `tfsdk:"active"`
	ForceInactive                       types.Bool             `tfsdk:"force_inactive"`
	Game                                types.String           `tfsdk:"game"`
	GamedigGivenPortOnly                types.Bool             `tfsdk:"gamedig_given_port_only"`
	GrpcBody                            types.String           `tfsdk:"grpc_body"`
	GrpcEnableTls                       types.Bool             `tfsdk:"grpc_enable_tls"`
	GrpcMetadata                        types.String           `tfsdk:"grpc_metadata"`
	GrpcMethod                          types.String           `tfsdk:"grpc_method"`
	GrpcProtobuf                        types.String           `tfsdk:"grpc_protobuf"`
	GrpcServiceName                     types.String           `tfsdk:"grpc_service_name"`
	GrpcUrl                             types.String           `tfsdk:"grpc_url"`
	HttpBodyEncoding                    types.String           `tfsdk:"http_body_encoding"`
	IncludeSensitiveData                types.Bool             `tfsdk:"include_sensitive_data"`
	InvertKeyword                       types.Bool             `tfsdk:"invert_keyword"`
	JsonPath                            types.String           `tfsdk:"json_path"`
	KafkaProducerAllowAutoTopicCreation types.Bool             `tfsdk:"kafka_producer_allow_auto_topic_creation"`
	KafkaProducerBrokers                types.Set              `tfsdk:"kafka_producer_brokers"`
	KafkaProducerMessage                types.String           `tfsdk:"kafka_producer_message"`
	KafkaProducerSaslOptions            types.String           `tfsdk:"kafka_producer_sasl_options"`
	KafkaProducerSsl                    types.Bool             `tfsdk:"kafka_producer_ssl"`
	KafkaProducerTopic                  types.String           `tfsdk:"kafka_producer_topic"`
	Maintenance                         types.Bool             `tfsdk:"maintenance"`
	OAuthAuthMethod                     types.String           `tfsdk:"oauth_auth_method"`
	OAuthClientID                       types.String           `tfsdk:"oauth_client_id"`
	OAuthClientSecret                   types.String           `tfsdk:"oauth_client_secret"`
	OAuthScopes                         types.Set              `tfsdk:"oauth_scopes"`
	OAuthTokenURL                       types.String           `tfsdk:"oauth_token_url"`
	PacketSize                          types.Int64            `tfsdk:"packet_size"`
	Parent                              types.String           `tfsdk:"parent"`
	PathName                            types.String           `tfsdk:"path_name"`
	PushToken                           types.String           `tfsdk:"push_token"`
	Screenshot                          types.String           `tfsdk:"screenshot"`
	Tags                                []tagInstanceDataModel `tfsdk:"tags"`
	Timeout                             types.Int64            `tfsdk:"timeout"`
	TlsCa                               types.String           `tfsdk:"tls_ca"`
	TlsCert                             types.String           `tfsdk:"tls_cert"`
	TlsKey                              types.String           `tfsdk:"tls_key"`
	Weight                              types.Int64            `tfsdk:"weight"`
}

type JSON_monitorModel struct {
	ID                                  int64                       `json:"id"`
	Type                                string                      `json:"type"`
	Name                                string                      `json:"name"`
	Interval                            int64                       `json:"interval"`
	RetryInterval                       int64                       `json:"retry_interval"`
	ResendInterval                      int64                       `json:"resend_interval"`
	MaxRetries                          int64                       `json:"max_retries"`
	UpsideDown                          bool                        `json:"upside_down"`
	NotificationIDList                  []string                    `json:"notification_id_list"`
	URL                                 string                      `json:"url"`
	ExpiryNotification                  bool                        `json:"expiry_notification"`
	IgnoreTls                           bool                        `json:"ignore_tls"`
	MaxRedirects                        int64                       `json:"max_redirects"`
	AcceptedStatusCodes                 []string                    `json:"accepted_statuscodes"`
	ProxyID                             int64                       `json:"proxy_id"`
	Method                              string                      `json:"method"`
	Body                                string                      `json:"body"`
	Headers                             string                      `json:"headers"`
	AuthMethod                          string                      `json:"auth_method"`
	BasicAuthUser                       string                      `json:"basic_auth_user"`
	BasicAuthPass                       string                      `json:"basic_auth_pass"`
	AuthDomain                          string                      `json:"auth_domain"`
	AuthWorkstation                     string                      `json:"auth_workstation"`
	Keyword                             string                      `json:"keyword"`
	Hostname                            string                      `json:"hostname"`
	Port                                int64                       `json:"port"`
	DNSResolveServer                    string                      `json:"dns_resolve_server"`
	DNSResolveType                      string                      `json:"dns_resolve_type"`
	MQTTUsername                        string                      `json:"mqtt_username"`
	MQTTPassword                        string                      `json:"mqtt_password"`
	MQTTTopic                           string                      `json:"mqtt_topic"`
	MQTTSucessMessage                   string                      `json:"mqtt_success_message"`
	DatabaseConnectionString            string                      `json:"database_connection_string"`
	DatabaseQuery                       string                      `json:"database_query"`
	DockerContainer                     string                      `json:"docker_container"`
	DockerHost                          int64                       `json:"docker_host"`
	RadiusUsername                      string                      `json:"radius_username"`
	RadiusPassword                      string                      `json:"radius_password"`
	RadiusSecret                        string                      `json:"radius_secret"`
	RadiusCalledStationId               string                      `json:"radius_called_station_id"`
	RadiusCallingStationId              string                      `json:"radius_calling_station_id"`
	Active                              bool                        `json:"active"`
	ForceInactive                       bool                        `json:"force_inactive"`
	Game                                string                      `json:"game"`
	GamedigGivenPortOnly                bool                        `json:"gamedig_given_port_only"`
	GrpcBody                            string                      `json:"grpc_body"`
	GrpcEnableTls                       bool                        `json:"grpc_enable_tls"`
	GrpcMetadata                        string                      `json:"grpc_metadata"`
	GrpcMethod                          string                      `json:"grpc_method"`
	GrpcProtobuf                        string                      `json:"grpc_protobuf"`
	GrpcServiceName                     string                      `json:"grpc_service_name"`
	GrpcUrl                             string                      `json:"grpc_url"`
	HttpBodyEncoding                    string                      `json:"http_body_encoding"`
	IncludeSensitiveData                bool                        `json:"include_sensitive_data"`
	InvertKeyword                       bool                        `json:"invert_keyword"`
	JsonPath                            string                      `json:"json_path"`
	KafkaProducerAllowAutoTopicCreation bool                        `json:"kafka_producer_allow_auto_topic_creation"`
	KafkaProducerBrokers                []string                    `json:"kafka_producer_brokers"`
	KafkaProducerMessage                string                      `json:"kafka_producer_message"`
	KafkaProducerSaslOptions            string                      `json:"kafka_producer_sasl_options"`
	KafkaProducerSsl                    bool                        `json:"kafka_producer_ssl"`
	KafkaProducerTopic                  string                      `json:"kafka_producer_topic"`
	Maintenance                         bool                        `json:"maintenance"`
	OAuthAuthMethod                     string                      `json:"oauth_auth_method"`
	OAuthClientID                       string                      `json:"oauth_client_id"`
	OAuthClientSecret                   string                      `json:"oauth_client_secret"`
	OAuthScopes                         []string                    `json:"oauth_scopes"`
	OAuthTokenURL                       string                      `json:"oauth_token_url"`
	PacketSize                          int64                       `json:"packet_size"`
	Parent                              string                      `json:"parent"`
	PathName                            string                      `json:"path_name"`
	PushToken                           string                      `json:"push_token"`
	Screenshot                          string                      `json:"screenshot"`
	Tags                                []JSON_tagInstanceDataModel `json:"tags"`
	Timeout                             int64                       `json:"timeout"`
	TlsCa                               string                      `json:"tls_ca"`
	TlsCert                             string                      `json:"tls_cert"`
	TlsKey                              string                      `json:"tls_key"`
	Weight                              int64                       `json:"weight"`
}

type monitorResponse struct {
	Monitor JSON_monitorModel `json:"monitor"`
}

// Configure adds the provider configured client to the data source.
func (d *data_monitorAuth) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
	tflog.Debug(ctx, "CONFIG received good token")

}

// Metadata returns the data source type name.
func (d *data_monitorAuth) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_monitor"
}

// Schema defines the schema for the data source.
func (d *data_monitorAuth) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id":              schema.Int64Attribute{Required: true},
			"type":            schema.StringAttribute{Computed: true},
			"name":            schema.StringAttribute{Computed: true},
			"interval":        schema.Int64Attribute{Computed: true},
			"retry_interval":  schema.Int64Attribute{Computed: true},
			"resend_interval": schema.Int64Attribute{Computed: true},
			"max_retries":     schema.Int64Attribute{Computed: true},
			"upside_down":     schema.BoolAttribute{Computed: true},
			"notification_id_list": schema.SetAttribute{
				ElementType: types.NumberType,
				Computed:    true,
			},
			"url":                 schema.StringAttribute{Computed: true},
			"expiry_notification": schema.BoolAttribute{Computed: true},
			"ignore_tls":          schema.BoolAttribute{Computed: true},
			"max_redirects":       schema.Int64Attribute{Computed: true},
			"accepted_statuscodes": schema.SetAttribute{
				ElementType: types.StringType,
				Computed:    true,
			},
			"proxy_id":        schema.Int64Attribute{Computed: true},
			"method":          schema.StringAttribute{Computed: true},
			"body":            schema.StringAttribute{Computed: true},
			"headers":         schema.StringAttribute{Computed: true},
			"auth_method":     schema.StringAttribute{Computed: true},
			"basic_auth_user": schema.StringAttribute{Computed: true},
			"basic_auth_pass": schema.StringAttribute{
				Sensitive: true,
				Computed:  true,
			},
			"auth_domain":        schema.StringAttribute{Computed: true},
			"auth_workstation":   schema.StringAttribute{Computed: true},
			"keyword":            schema.StringAttribute{Computed: true},
			"hostname":           schema.StringAttribute{Computed: true},
			"port":               schema.Int64Attribute{Computed: true},
			"dns_resolve_server": schema.StringAttribute{Computed: true},
			"dns_resolve_type":   schema.StringAttribute{Computed: true},
			"mqtt_username":      schema.StringAttribute{Computed: true},
			"mqtt_password": schema.StringAttribute{
				Sensitive: true,
				Computed:  true,
			},
			"mqtt_topic":           schema.StringAttribute{Computed: true},
			"mqtt_success_message": schema.StringAttribute{Computed: true},
			"database_connection_string": schema.StringAttribute{
				Sensitive: true,
				Computed:  true,
			},
			"database_query":   schema.StringAttribute{Computed: true},
			"docker_container": schema.StringAttribute{Computed: true},
			"docker_host":      schema.Int64Attribute{Computed: true},
			"radius_username":  schema.StringAttribute{Computed: true},
			"radius_password": schema.StringAttribute{
				Sensitive: true,
				Computed:  true,
			},
			"radius_secret": schema.StringAttribute{
				Sensitive: true,
				Computed:  true,
			},
			"radius_called_station_id":  schema.StringAttribute{Computed: true},
			"radius_calling_station_id": schema.StringAttribute{Computed: true},
			"active":                    schema.BoolAttribute{Computed: true},
			"force_inactive":            schema.BoolAttribute{Computed: true},
			"game":                      schema.StringAttribute{Computed: true},
			"gamedig_given_port_only":   schema.BoolAttribute{Computed: true},
			"grpc_body":                 schema.StringAttribute{Computed: true},
			"grpc_enable_tls":           schema.BoolAttribute{Computed: true},
			"grpc_metadata":             schema.StringAttribute{Computed: true},
			"grpc_method":               schema.StringAttribute{Computed: true},
			"grpc_protobuf":             schema.StringAttribute{Computed: true},
			"grpc_service_name":         schema.StringAttribute{Computed: true},
			"grpc_url":                  schema.StringAttribute{Computed: true},
			"http_body_encoding":        schema.StringAttribute{Computed: true},
			"include_sensitive_data":    schema.BoolAttribute{Computed: true},
			"invert_keyword":            schema.BoolAttribute{Computed: true},
			"json_path":                 schema.StringAttribute{Computed: true},
			"kafka_producer_allow_auto_topic_creation": schema.BoolAttribute{Computed: true},
			"kafka_producer_brokers": schema.SetAttribute{
				ElementType: types.StringType,
				Computed:    true,
			},
			"kafka_producer_message":      schema.StringAttribute{Computed: true},
			"kafka_producer_sasl_options": schema.StringAttribute{Computed: true},
			"kafka_producer_ssl":          schema.BoolAttribute{Computed: true},
			"kafka_producer_topic":        schema.StringAttribute{Computed: true},
			"maintenance":                 schema.BoolAttribute{Computed: true},
			"oauth_auth_method":           schema.StringAttribute{Computed: true},
			"oauth_client_id":             schema.StringAttribute{Computed: true},
			"oauth_client_secret": schema.StringAttribute{
				Sensitive: true,
				Computed:  true,
			},
			"oauth_scopes": schema.SetAttribute{
				ElementType: types.StringType,
				Computed:    true,
			},
			"oauth_token_url": schema.StringAttribute{Computed: true},
			"packet_size":     schema.Int64Attribute{Computed: true},
			"parent":          schema.StringAttribute{Computed: true},
			"path_name":       schema.StringAttribute{Computed: true},
			"push_token":      schema.StringAttribute{Computed: true},
			"screenshot":      schema.StringAttribute{Computed: true},
			"tags": schema.SetNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Computed: true,
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"color": schema.StringAttribute{
							Computed: true,
						},
						"value": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
			"timeout":  schema.Int64Attribute{Computed: true},
			"tls_ca":   schema.StringAttribute{Computed: true},
			"tls_cert": schema.StringAttribute{Computed: true},
			"tls_key": schema.StringAttribute{
				Sensitive: true,
				Computed:  true,
			},
			"weight": schema.Int64Attribute{Computed: true},
		},
	}
}

// Read refreshes the Terraform state with the latest data.
func (d *data_monitorAuth) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {

	var state monitorModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	id := cleanString(state.ID.String())

	var respDebug string
	tflog.Debug(ctx, "Requesting "+d.Host+"/monitors/"+id)
	err := requests.
		URL(d.Host).
		Bearer(d.Token).
		Path("/monitors/" + id).
		ToString(&respDebug).
		Fetch(ctx)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error with response",
			"got "+err.Error()+"\n"+
				"token was "+d.Token,
		)
	}

	var response monitorResponse
	err = json.Unmarshal([]byte(respDebug), &response)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error unmarshalling response",
			err.Error(),
		)
	}

	tflog.Debug(ctx, "Request got "+respDebug)

	notificationIDs, diags := types.SetValueFrom(ctx, types.NumberType, response.Monitor.NotificationIDList)
	resp.Diagnostics.Append(diags...)
	acceptedStatusCodes, diags := types.SetValueFrom(ctx, types.StringType, response.Monitor.AcceptedStatusCodes)
	resp.Diagnostics.Append(diags...)
	kafakaProducerBrokers, diags := types.SetValueFrom(ctx, types.StringType, response.Monitor.KafkaProducerBrokers)
	resp.Diagnostics.Append(diags...)
	OAuthScopes, diags := types.SetValueFrom(ctx, types.StringType, response.Monitor.OAuthScopes)
	resp.Diagnostics.Append(diags...)

	var tags []tagInstanceDataModel
	for _, tag := range response.Monitor.Tags {
		tags = append(tags, tagInstanceDataModel{
			ID:    types.Int64Value(tag.ID),
			Name:  types.StringValue(tag.Name),
			Value: types.StringValue(tag.Value),
			Color: types.StringValue(tag.Color),
		})
	}

	tout := monitorModel{
		ID:                                  types.Int64Value(response.Monitor.ID),
		Type:                                types.StringValue(response.Monitor.Type),
		Name:                                types.StringValue(response.Monitor.Name),
		Interval:                            types.Int64Value(response.Monitor.Interval),
		RetryInterval:                       types.Int64Value(response.Monitor.RetryInterval),
		ResendInterval:                      types.Int64Value(response.Monitor.ResendInterval),
		MaxRetries:                          types.Int64Value(response.Monitor.MaxRetries),
		UpsideDown:                          types.BoolValue(response.Monitor.UpsideDown),
		NotificationIDList:                  notificationIDs,
		URL:                                 types.StringValue(response.Monitor.URL),
		ExpiryNotification:                  types.BoolValue(response.Monitor.ExpiryNotification),
		IgnoreTls:                           types.BoolValue(response.Monitor.IgnoreTls),
		MaxRedirects:                        types.Int64Value(response.Monitor.MaxRedirects),
		AcceptedStatusCodes:                 acceptedStatusCodes,
		ProxyID:                             types.Int64Value(response.Monitor.ProxyID),
		Method:                              types.StringValue(response.Monitor.Method),
		Body:                                types.StringValue(response.Monitor.Body),
		Headers:                             types.StringValue(response.Monitor.Headers),
		AuthMethod:                          types.StringValue(response.Monitor.AuthMethod),
		BasicAuthUser:                       types.StringValue(response.Monitor.BasicAuthUser),
		BasicAuthPass:                       types.StringValue(response.Monitor.BasicAuthPass),
		AuthDomain:                          types.StringValue(response.Monitor.AuthDomain),
		AuthWorkstation:                     types.StringValue(response.Monitor.AuthWorkstation),
		Keyword:                             types.StringValue(response.Monitor.Keyword),
		Hostname:                            types.StringValue(response.Monitor.Hostname),
		Port:                                types.Int64Value(response.Monitor.Port),
		DNSResolveServer:                    types.StringValue(response.Monitor.DNSResolveServer),
		DNSResolveType:                      types.StringValue(response.Monitor.DNSResolveType),
		MQTTUsername:                        types.StringValue(response.Monitor.MQTTUsername),
		MQTTPassword:                        types.StringValue(response.Monitor.MQTTPassword),
		MQTTTopic:                           types.StringValue(response.Monitor.MQTTTopic),
		MQTTSucessMessage:                   types.StringValue(response.Monitor.MQTTSucessMessage),
		DatabaseConnectionString:            types.StringValue(response.Monitor.DatabaseConnectionString),
		DatabaseQuery:                       types.StringValue(response.Monitor.DatabaseQuery),
		DockerContainer:                     types.StringValue(response.Monitor.DockerContainer),
		DockerHost:                          types.Int64Value(response.Monitor.DockerHost),
		RadiusUsername:                      types.StringValue(response.Monitor.RadiusUsername),
		RadiusPassword:                      types.StringValue(response.Monitor.RadiusPassword),
		RadiusSecret:                        types.StringValue(response.Monitor.RadiusSecret),
		RadiusCalledStationId:               types.StringValue(response.Monitor.RadiusCalledStationId),
		RadiusCallingStationId:              types.StringValue(response.Monitor.RadiusCallingStationId),
		Active:                              types.BoolValue(response.Monitor.Active),
		ForceInactive:                       types.BoolValue(response.Monitor.ForceInactive),
		Game:                                types.StringValue(response.Monitor.Game),
		GamedigGivenPortOnly:                types.BoolValue(response.Monitor.GamedigGivenPortOnly),
		GrpcBody:                            types.StringValue(response.Monitor.GrpcBody),
		GrpcEnableTls:                       types.BoolValue(response.Monitor.GrpcEnableTls),
		GrpcMetadata:                        types.StringValue(response.Monitor.GrpcMetadata),
		GrpcMethod:                          types.StringValue(response.Monitor.GrpcMethod),
		GrpcProtobuf:                        types.StringValue(response.Monitor.GrpcProtobuf),
		GrpcServiceName:                     types.StringValue(response.Monitor.GrpcServiceName),
		GrpcUrl:                             types.StringValue(response.Monitor.GrpcUrl),
		HttpBodyEncoding:                    types.StringValue(response.Monitor.HttpBodyEncoding),
		IncludeSensitiveData:                types.BoolValue(response.Monitor.IncludeSensitiveData),
		InvertKeyword:                       types.BoolValue(response.Monitor.InvertKeyword),
		JsonPath:                            types.StringValue(response.Monitor.JsonPath),
		KafkaProducerAllowAutoTopicCreation: types.BoolValue(response.Monitor.KafkaProducerAllowAutoTopicCreation),
		KafkaProducerBrokers:                kafakaProducerBrokers,
		KafkaProducerMessage:                types.StringValue(response.Monitor.KafkaProducerMessage),
		KafkaProducerSaslOptions:            types.StringValue(response.Monitor.KafkaProducerSaslOptions),
		KafkaProducerSsl:                    types.BoolValue(response.Monitor.KafkaProducerSsl),
		KafkaProducerTopic:                  types.StringValue(response.Monitor.KafkaProducerTopic),
		Maintenance:                         types.BoolValue(response.Monitor.Maintenance),
		OAuthAuthMethod:                     types.StringValue(response.Monitor.OAuthAuthMethod),
		OAuthClientID:                       types.StringValue(response.Monitor.OAuthClientID),
		OAuthClientSecret:                   types.StringValue(response.Monitor.OAuthClientSecret),
		OAuthScopes:                         OAuthScopes,
		OAuthTokenURL:                       types.StringValue(response.Monitor.OAuthTokenURL),
		PacketSize:                          types.Int64Value(response.Monitor.PacketSize),
		Parent:                              types.StringValue(response.Monitor.Parent),
		PathName:                            types.StringValue(response.Monitor.PathName),
		PushToken:                           types.StringValue(response.Monitor.PushToken),
		Screenshot:                          types.StringValue(response.Monitor.Screenshot),
		Tags:                                tags,
		Timeout:                             types.Int64Value(response.Monitor.Timeout),
		TlsCa:                               types.StringValue(response.Monitor.TlsCa),
		TlsCert:                             types.StringValue(response.Monitor.TlsCert),
		TlsKey:                              types.StringValue(response.Monitor.TlsKey),
		Weight:                              types.Int64Value(response.Monitor.Weight),
	}

	diags = resp.State.Set(ctx, &tout)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

}
