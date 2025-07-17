package zabbix

// MediaType represents a Zabbix media type returned from the Zabbix API.
//
// See: https://www.zabbix.com/documentation/7.4/en/manual/api/reference/mediatype/object
type MediaType struct {
	// ID of the media type.
	MediaTypeID string `json:"mediatypeid"`

	// Name of the media type.
	Name string `json:"name"`

	// Transport used by the media type.
	//
	// Possible values:
	// 0 - Email;
	// 1 - Script;
	// 2 - SMS;
	// 4 - Webhook.
	Type int `json:"type,string"`

	// Name of the script file (e.g., notification.sh) that is located in
	// the directory specified in the AlertScriptsPath server configuration parameter.
	ExecPath string `json:"exec_path"`

	// Serial device name of the GSM modem.
	GsmModem string `json:"gsm_modem"`

	// Authentication password.
	PassWd string `json:"passwd"`

	// Email provider.
	//
	// Possible values:
	// 0 - (default) Generic SMTP;
	// 1 - Gmail;
	// 2 - Gmail relay;
	// 3 - Office365;
	// 4 - Office365 relay.
	Provider int `json:"provider,string"`

	// Email address from which notifications will be sent.
	SmtpEmail string `json:"smtp_email"`

	// SMTP HELO.
	SmtpHelo string `json:"smtp_helo"`

	// SMTP server.
	SmtpServer string `json:"smtp_server"`

	// SMTP server port to connect to.
	SmtpPort int `json:"smtp_port,string"`

	// SMTP connection security level to use.
	//
	// Possible values:
	// 0 - (default) None;
	// 1 - STARTTLS;
	// 2 - SSL/TLS.
	SmtpSecurity int `json:"smtp_security,string"`

	// SSL verify host for SMTP.
	//
	// Possible values:
	// 0 - (default) No;
	// 1 - Yes.
	SmtpVerifyHost int `json:"smtp_verify_host,string"`

	// SSL verify peer for SMTP.
	//
	// Possible values:
	// 0 - (default) No;
	// 1 - Yes.
	SmtpVerifyPeer int `json:"smtp_verify_peer,string"`

	// SMTP authentication method to use.
	//
	// Possible values:
	// 0 - (default) None;
	// 1 - Normal password;
	// 2 - OAuth token.
	// OAuth authentication is not allowed for Office365 relay email provider.
	SmtpAuthentication int `json:"smtp_authentication,string"`

	// Zabbix frontend URL to redirect back OAuth authorization.
	//
	// Default:
	// Value of API settings property url with part zabbix.php?action=oauth.authorize
	RedirectionUrl string `json:"redirection_url"`

	// The client identifier registered within the OAuth authorization server.
	ClientID string `json:"client_id"`

	// The client secret registered within the OAuth authorization server.
	// Accessible only for user of type Super Admin.
	ClientSecret string `json:"client_secret"`

	// OAuth URL, with parameters, to get access and refresh tokens.
	AuthorizationUrl string `json:"authorization_url"`

	// OAuth URL to exchange authorization token to access and refresh tokens.
	// This URL also is used by server to refresh invalid access token.
	TokenUrl string `json:"token_url"`

	// Bit mask on tokens' status.
	//
	// Possible values:
	// 0 - (default) Both tokens contain invalid value
	// 1 - Access token contains valid value
	// 2 - Refresh token contains valid value
	// 3 - Both tokens contain valid value.
	TokensStatus int `json:"tokens_status,string"`

	// OAuth access token value.
	AccessToken string `json:"access_token"`

	// Timestamp of last modification of access_token done by server when refreshing
	// with refresh_token or API on token changes.
	AccessTokenUpdated int `json:"access_token_updated,string"`

	// Time in seconds when access_token will become outdated and will require to make
	// request to refresh_url.
	// Is set by Zabbix server on access_token refresh or by API on token changes.
	//
	// Timestamp is calculated by adding value of access_token_updated.
	AccessExpiresIn int `json:"access_expires_in,string"`

	// OAuth refresh token value.
	RefreshToken string `json:"refresh_token"`

	// Whether the media type is enabled.
	//
	// Possible values:
	// 0 - (default) Enabled;
	// 1 - Disabled.
	Status int `json:"status,string"`

	// User name.
	Username string `json:"username"`

	// The maximum number of alerts that can be processed in parallel.
	//
	// Possible values if type is set to 'SMS': 1.
	//
	// Possible values if type is set to 'Email', 'Script', or 'Webhook': 0-100.
	//
	// Default: 1.
	MaxSessions int `json:"maxsessions,string"`

	// The maximum number of attempts to send an alert.
	//
	// Possible values: 1-100.
	//
	// Default: 3.
	MaxAttempts int `json:"maxattempts,string"`

	// The interval between retry attempts.
	// Accepts seconds and time unit with suffix.
	//
	// Possible values: 0-1h.
	//
	// Default: 10s.
	AttemptInterval string `json:"attempt_interval"`

	// Message format.
	//
	// Possible values:
	// 0 - Plain text;
	// 1 - (default) HTML.
	MessageFormat int `json:"message_format,string"`

	// Webhook script body (JavaScript).
	Script string `json:"script"`

	// Webhook script timeout.
	// Accepts seconds and time unit with suffix.
	//
	// Possible values: 1-60s.
	//
	// Default: 30s.
	Timeout string `json:"timeout"`

	// Process JSON property values in Webhook script response as tags. These tags are
	// added to any existing problem tags.
	//
	// Possible values:
	// 0 - (default) Ignore webhook script response;
	// 1 - Process webhook script response as tags.
	ProcessTags int `json:"process_tags,string"`

	// Include an entry in the event menu that links to a custom URL. Also adds the urls
	// property to the output of problem.get and event.get.
	//
	// Possible values:
	// 0 - (default) Do not include event menu entry or urls property;
	// 1 - Include event menu entry and urls property.
	ShowEventMenu int `json:"show_event_menu,string"`

	// URL used in the event menu entry and in the urls property returned by problem.get
	// and event.get.
	EventMenuUrl string `json:"event_menu_url"`

	// Name used for the event menu entry and in the urls property returned by problem.get
	// and event.get.
	EventMenuName string `json:"event_menu_name"`

	// Webhook or script parameters.
	Parameters []MediaTypeParameters `json:"parameters"`

	//Media type description.
	Description string `json:"description"`

	// Actions is an array of actions that use the media type.
	//
	// Actions is only populated if MediaTypeGetParams.SelectActions is given in the
	// query parameters that returned this MediaType.
	Actions []Action `json:"actions"`

	// MessageTemplates is an array of media type messages.
	//
	// MessageTemplates is only populated if MediaTypeGetParams.SelectMessageTemplates is
	// given in the query parameters that returned this MediaType.
	MessageTemplates []MediaTypeMessageTemplate `json:"message_templates"`

	// Users is an array of media type messages.
	//
	// Users is only populated if MediaTypeGetParams.SelectUsers is given in the
	// query parameters that returned this MediaType.
	Users []User `json:"users"`
}

// MediaTypeParameters is params for Webhook and Script types
type MediaTypeParameters struct {
	Name      string `json:"name"`
	SortOrder int    `json:"sortorder,string"`
	Value     string `json:"value"`
}

// MediaTypeMessageTemplate is media type templates
type MediaTypeMessageTemplate struct {
	// Event source.
	//
	// Possible values:
	// 0 - Triggers;
	// 1 - Discovery;
	// 2 - Autoregistration;
	// 3 - Internal;
	// 4 - Services.
	EventSource int `json:"eventsource,string"`

	// Operation mode.
	//
	// Possible values:
	// 0 - Operations;
	// 1 - Recovery operations;
	// 2 - Update operations.
	Recovery int `json:"recovery,string"`

	// Message subject.
	Subject string `json:"subject"`

	// Message text.
	Message string `json:"message"`
}

// Media is user media
type Media struct {
	// ID of the user's media.
	MediaId string `json:"mediaid"`

	// ID of the media type used by the user's media.
	MediaTypeId string `json:"mediatypeid"`

	// Address, user name or other identifier of the recipient.
	//
	// If type of Media type is set to 'Email', values are represented as array.
	// For other types of Media types, value is represented as a string.
	SendTo []string `json:"sendto"`

	// Whether the media is enabled.
	//
	// Possible values:
	// 0 - (default) enabled;
	// 1 - disabled.
	Active int `json:"active,string"`

	// Trigger severities to send notifications about.
	//
	// Possible bitmap values:
	// 1 - Not classified;
	// 2 - Information;
	// 4 - Warning;
	// 8 - Average;
	// 16 - High;
	// 32 - Disaster.
	//
	// This is a bitmask field; any sum of possible bitmap values is acceptable (for example, 48 for Average, High, and Disaster).
	//
	// Default: 63.
	Severity int `json:"severity,string"`

	// Time when the notifications can be sent as a time period or user macros separated by a semicolon.
	//
	// Default: 1-7,00:00-24:00.
	Period string `json:"period"`

	// Whether the user has been provisioned.
	//
	// Possible values:
	// 0 - not provisioned;
	// 1 - provisioned.
	Provisioned int `json:"provisioned,string"`

	// User directory media mapping ID for provisioned media.
	UserDirectoryMediaId string `json:"userdirectory_mediaid"`
}

// MediaTypeGetParams is params for mediatype.get query
type MediaTypeGetParams struct {
	GetParameters

	// Return only media types with the given IDs.
	MediaTypeIDs []string `json:"mediatypeids,omitempty"`

	// Return only media types used by the given media.
	MediaIDs []string `json:"mediaids,omitempty"`

	// Return only media types used by the given users.
	UserIDs []string `json:"userids,omitempty"`

	// Return an actions property with the actions that use the media type.
	SelectActions SelectQuery `json:"selectActions,omitempty"`

	// Return a message_templates property with an array of media type messages.
	SelectMessageTemplates SelectQuery `json:"selectMessageTemplates,omitempty"`

	// Return a users property with the users that use the media type.
	SelectUsers SelectQuery `json:"selectUsers,omitempty"`
}

// GetMedias queries the Zabbix API for Medias matching the given search
// parameters.
//
// ErrMediaNotFound is returned if the search result set is empty.
// An error is returned if a transport, parsing or API error occurs.
func (c *Session) GetMediaTypes(params MediaTypeGetParams) ([]MediaType, error) {
	Mediatypes := make([]MediaType, 0)
	err := c.Get("mediatype.get", params, &Mediatypes)
	if err != nil {
		return nil, err
	}

	if len(Mediatypes) == 0 {
		return nil, ErrNotFound
	}

	return Mediatypes, nil
}
