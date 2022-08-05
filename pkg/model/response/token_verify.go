package response

type tokenResult struct {
	Id          string `json:"id"`
	Status      string `json:"status"`
	Not_before  string `json:"not_before"`
	Expires_one string `json:"expires_one"`
}

type TokenResponse struct {
	Success  bool        `json:"success"`
	Messages []message   `json:"messages"`
	Errors   []error     `json:"errors"`
	Result   tokenResult `json:"result"`
}
