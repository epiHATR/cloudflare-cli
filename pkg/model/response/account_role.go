package response

type AccountRolesResponse struct {
	Success     bool        `json:"success"`
	Messages    []message   `json:"messages"`
	Errors      []error     `json:"errors"`
	Result_Info result_Info `json:"result_info"`
	Result      []Role      `json:"result"`
}
