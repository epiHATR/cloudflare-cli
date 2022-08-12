package response

type actions struct {
	Read  bool `json:"read"`
	Write bool `json:"write"`
}

type permission struct {
	Analytics     actions `json:"analytics"`
	Billing       actions `json:"billing"`
	Cache_purge   actions `json:"cache_purge"`
	Dns           actions `json:"dns"`
	Dns_records   actions `json:"dns_records"`
	Lb            actions `json:"lb"`
	Logs          actions `json:"logs"`
	Organization  actions `json:"organization"`
	Ssl           actions `json:"ssl"`
	Waf           actions `json:"waf"`
	Zones         actions `json:"zones"`
	Zone_settings actions `json:"zone_settings"`
}
type Role struct {
	Id          string     `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Permissions permission `json:"permissions"`
}

type basicInfo struct {
	Id                                string `json:"id"`
	First_name                        string `json:"first_name"`
	Last_name                         string `json:"last_name"`
	Email                             string `json:"email"`
	Two_factor_authentication_enabled bool   `json:"two_factor_authentication_enabled"`
}

type UserDetails struct {
	Id     string    `json:"id"`
	Code   string    `json:"code"`
	User   basicInfo `json:"user"`
	Status string    `json:"status"`
	Roles  []Role    `json:"roles"`
}

type AccountUsersResponse struct {
	Success     bool          `json:"success"`
	Messages    []message     `json:"messages"`
	Errors      []error       `json:"errors"`
	Result      []UserDetails `json:"result"`
	Result_Info result_Info   `json:"result_info"`
}

type AccountUserDetailResponse struct {
	Success  bool        `json:"success"`
	Messages []message   `json:"messages"`
	Errors   []error     `json:"errors"`
	Result   UserDetails `json:"result"`
}
