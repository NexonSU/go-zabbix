package zabbix

// User represents a Zabbix User returned from the Zabbix API.
//
// See: https://www.zabbix.com/documentation/7.4/en/manual/api/reference/user/object
type User struct {
	// ID of the user.
	UserId string `json:"userid"`

	// User's name.
	Username string `json:"username"`

	// User's password.
	//
	// The value of this parameter can be an empty string if the user is linked to a user directory.
	PassWd string `json:"passwd"`

	// ID of the role of the user.
	//
	// Note that users without a role can log into Zabbix only using LDAP or SAML authentication,
	// provided their LDAP/SAML information matches the user group mappings configured in Zabbix.
	RoleId string `json:"roleid"`

	// Time of the last unsuccessful login attempt.
	AttemptClock int `json:"attempt_clock,string"`

	// Recent failed login attempt count.
	AttemptFailed int `json:"attempt_failed,string"`

	// IP address from where the last unsuccessful login attempt came from.
	AttemptIp string `json:"attempt_ip"`

	// Whether to enable auto-login.
	//
	// Possible values:
	// 0 - (default) auto-login disabled;
	// 1 - auto-login enabled.
	AutoLogin int `json:"autologin,string"`

	// User session life time. Accepts seconds and time unit with suffix.
	// If set to 0s, the session will never expire.
	//
	//Default: 15m.
	AutoLogout string `json:"autologout"`

	// Language code of the user's language, for example, en_US.
	//
	//Default: default - system default.
	Lang string `json:"lang"`

	// Name of the user.
	Name string `json:"name"`

	// Whether the user has been provisioned.
	//
	// Possible values:
	// 0 - not provisioned;
	// 1 - provisioned.
	Provisioned int `json:"provisioned,string"`

	// Automatic refresh period. Accepts seconds or time unit with suffix
	// (e.g., 30s, 90s, 1m, 1h).
	//
	// Default: 30s.
	Refresh string `json:"refresh"`

	// Amount of object rows to show per page.
	//
	// Default: 50.
	RowsPerPage int `json:"rows_per_page,string"`

	// Surname of the user.
	Surname string `json:"surname"`

	// User's theme.
	//
	// Possible values:
	// default - (default) system default;
	// blue-theme - Blue;
	// dark-theme - Dark.
	Theme string `json:"theme"`

	// Time when the latest provisioning operation was made.
	TsProvisioned int `json:"ts_provisioned,string"`

	// URL of the page to redirect the user to after logging in.
	Url string `json:"url"`

	// ID of the user directory that the user in linked to.
	//
	// Used for provisioning (creating or updating), as well as to login a user
	// that is linked to a user directory.
	//
	// For login operations the value of this property will have priority over
	// the userdirectoryid property of user groups that the user belongs to.
	//
	// Default: 0.
	UserDirectoryId string `json:"userdirectoryid"`

	// User's time zone, for example, Europe/London, UTC.
	//
	// Default: default - system default.
	//
	// For the full list of supported time zones please refer to PHP documentation.
	Timezone string `json:"timezone"`

	// Medias is an array of user medias.
	//
	// Medias is only populated if UserGetParams.SelectMedias is given in the
	// query parameters that returned this User.
	Medias []Media `json:"medias"`

	// MediaTypes is an array of user media types.
	//
	// MediaTypes is only populated if UserGetParams.SelectMediatypes is given in the
	// query parameters that returned this User.
	MediaTypes []MediaType `json:"mediatypes"`
}

// UserGetParams is params for user.get query
type UserGetParams struct {
	GetParameters

	// Return only users that use the given media.
	MediaIDs []string `json:"mediaids,omitempty"`

	// Return only users that use the given media types.
	MediaTypeIDs []string `json:"mediatypeids,omitempty"`

	// Return only users with the given IDs.
	UserIDs []string `json:"userids,omitempty"`

	// Return only users that belong to the given user groups.
	UsrGrpIds []string `json:"usrgrpids,omitempty"`

	// Adds additional information about user permissions.
	//
	// Adds the following properties for each user:
	// gui_access - (integer) user's frontend authentication method.
	// Refer to the gui_access property of the user group object for a list of possible values.
	//
	// debug_mode - (integer) indicates whether debug is enabled for the user.
	// Possible values: 0 - debug disabled, 1 - debug enabled.
	//
	// users_status - (integer) indicates whether the user is disabled.
	// Possible values: 0 - user enabled, 1 - user disabled.
	GetAccess bool `json:"getAccess,omitempty"`

	// Return media used by the user in the medias property.
	SelectMedias SelectQuery `json:"selectMedias,omitempty"`

	// Return media types used by the user in the mediatypes property.
	//
	// See mediatype.get for restrictions based on user type.
	SelectMediatypes SelectQuery `json:"selectMediatypes,omitempty"`

	// Return user groups that the user belongs to in the usrgrps property.
	//
	// See usergroup.get for restrictions based on user type.
	//SelectUsrgrps SelectQuery `json:"selectUsrgrps,omitempty"`

	// Return user role in the role property.
	//SelectRole SelectQuery `json:"selectRole,omitempty"`
}

// GetUsers queries the Zabbix API for Users matching the given search
// parameters.
//
// ErrUserNotFound is returned if the search result set is empty.
// An error is returned if a transport, parsing or API error occurs.
func (c *Session) GetUsers(params UserGetParams) ([]User, error) {
	Users := make([]User, 0)
	err := c.Get("User.get", params, &Users)
	if err != nil {
		return nil, err
	}

	if len(Users) == 0 {
		return nil, ErrNotFound
	}

	return Users, nil
}
