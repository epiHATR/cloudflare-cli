package response

type setting struct {
	Enforce_twofactor                bool `json:"enforce_twofactor"`
	Use_account_custom_ns_by_default bool `json:"use_account_custom_ns_by_default"`
}

type AccountDetails struct {
	Id         string  `json:"id"`
	Name       string  `json:"name"`
	Settings   setting `json:"setting"`
	Created_on string  `json:"created_on"`
}

type AccountListResponse struct {
	Success     bool             `json:"success"`
	Messages    []message        `json:"messages"`
	Errors      []error          `json:"errors"`
	Result_Info result_Info      `json:"result_info"`
	Result      []AccountDetails `json:"result"`
}

type AccountDetailsResponse struct {
	Success  bool           `json:"success"`
	Messages []message      `json:"messages"`
	Errors   []error        `json:"errors"`
	Result   AccountDetails `json:"result"`
}
