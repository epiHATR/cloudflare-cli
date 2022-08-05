package response

///////// CLOUDFLARE API RESPONSE DATA /////////////////////
type message struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type meta struct {
	Step                      int  `json:"step"`
	Custom_certificate_quota  int  `json:"custom_certificate_quota"`
	Page_rule_quota           int  `json:"page_rule_quota"`
	Phishing_detected         bool `json:"phishing_detected"`
	Multiple_railguns_allowed bool `json:"multiple_railguns_allowed"`
}

type owner struct {
	Id   string `json:"id"`
	Type string `json:"type"`
	Name string `json:"name"`
}

type account struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type tenant struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type tenamt_unit struct {
	Id string `json:"id"`
}

type plan struct {
	Id                 string  `json:"id"`
	Name               string  `json:"name"`
	Price              float32 `json:"Price"`
	Currency           string  `json:"Currency"`
	Frequency          string  `json:"Frequency"`
	Is_subscribed      bool    `json:"is_subscribed"`
	Can_subscribe      bool    `json:"can_subscribe"`
	Legacy_id          string  `json:"legacy_id"`
	Legacy_discount    bool    `json:"legacy_discount"`
	Externally_managed bool    `json:"externally_managed"`
}

type result struct {
	Id                    string       `json:"id"`
	Name                  string       `json:"name"`
	Status                string       `json:"status"`
	Paused                bool         `json:"paused"`
	Type                  string       `json:"type"`
	Development_mode      int          `json:"development_mode"`
	Verification_key      string       `json:"verification_key"`
	Cname_suffix          string       `json:"cname_suffix"`
	Original_name_servers []string     `json:"original_name_servers"`
	Original_registrar    string       `json:"original_registrar"`
	Original_dnshost      string       `json:"original_dnshost"`
	Modified_on           string       `json:"modified_on"`
	Created_on            string       `json:"created_on"`
	Activated_on          string       `json:"activated_on"`
	Meta                  meta         `json:"meta"`
	Owner                 owner        `json:"owner"`
	Account               account      `json:"account"`
	Tenant                tenant       `json:"tenant"`
	Tenant_unit           tenamt_unit  `json:"tenant_unit"`
	Permissions           []string     `json:"permissions"`
	Plan                  plan         `json:"plan"`
	Plan_pending          plan_pending `json:"plan_pending"`
	Name_servers          []string     `json:"name_servers"`
}

type dnsListMeta struct {
	Auto_added             bool   `json:"auto_added"`
	Managed_by_apps        bool   `json:"managed_by_apps"`
	Managed_by_argo_tunnel bool   `json:"managed_by_argo_tunnel"`
	Source                 string `json:"source"`
}

type dnslistResult struct {
	Id          string      `json:"id"`
	Zone_id     string      `json:"zone_id"`
	Zone_name   string      `json:"zone_name"`
	Name        string      `json:"name"`
	Type        string      `json:"type"`
	Content     string      `json:"content"`
	Created_on  string      `json:"created_on"`
	Modified_on string      `json:"modified_on"`
	Meta        dnsListMeta `json:"meta"`
}

type result_Info struct {
	Page        int `json:"page"`
	Per_page    int `json:"per_page"`
	Total_pages int `json:"total_pages"`
	Count       int `json:"count"`
	Total_count int `json:"total_count"`
}

type plan_pending struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	Price         int    `json:"price"`
	Currency      string `json:"currency"`
	Frequency     string `json:"frequency"`
	Legacy_id     string `json:"legacy_id"`
	Is_subscribed bool   `json:"is_subscribed"`
	Can_subscribe bool   `json:"can_subscribe"`
}

type Response struct {
	Success  bool      `json:"success"`
	Messages []message `json:"messages"`
	Errors   []error   `json:"errors"`
}

type ZoneListResponse struct {
	Success     bool        `json:"success"`
	Messages    []message   `json:"messages"`
	Errors      []error     `json:"errors"`
	Result      []result    `json:"result"`
	Result_Info result_Info `json:"result_info"`
}

type ZoneDetailResponse struct {
	Success  bool      `json:"success"`
	Messages []message `json:"messages"`
	Errors   []error   `json:"errors"`
	Result   result    `json:"result"`
}

type DnsListResponse struct {
	Success  bool            `json:"success"`
	Messages []message       `json:"messages"`
	Errors   []error         `json:"errors"`
	Result   []dnslistResult `json:"result"`
}
