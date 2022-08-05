package response

type userResult struct {
	Id                                string `json:"id"`
	Email                             string `json:"email"`
	First_name                        string `json:"first_name"`
	Last_name                         string `json:"last_name"`
	Username                          string `json:"username"`
	Telephone                         string `json:"telephone"`
	Country                           string `json:"contriy"`
	Zipcode                           string `json:"zipcode"`
	Created_on                        string `json:"created_on"`
	Modified_on                       string `json:"modified_on"`
	Two_factor_authentication_enabled bool   `json:"two_factor_authentication_enabled"`
	Suspended                         bool   `json:"suspended"`
}

type UserReponse struct {
	Success  bool       `json:"success"`
	Messages []message  `json:"messages"`
	Errors   []error    `json:"errors"`
	Result   userResult `json:"result"`
}
