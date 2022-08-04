package structs

type Config struct {
	Auth auth
}

type auth struct {
	Token string `mapstructure:"token"`
	Email string `mapstructure:"email"`
	Key   string `mapstructure:"key"`
}

///////// CLOUDFLARE API RESPONSE DATA /////////////////////
type Message struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Response struct {
	Success  bool      `json:"success"`
	Messages []Message `json:"messages"`
	Errors   []Error   `json:"errors"`
}

type Meta struct {
	Step                      int  `json:"step"`
	Custom_certificate_quota  int  `json:"custom_certificate_quota"`
	Page_rule_quota           int  `json:"page_rule_quota"`
	Phishing_detected         bool `json:"phishing_detected"`
	Multiple_railguns_allowed bool `json:"multiple_railguns_allowed"`
}

type Owner struct {
	Id   string `json:"id"`
	Type string `json:"type"`
	Name string `json:"name"`
}

type Account struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Tenant struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Tenamt_unit struct {
	Id string `json:"id"`
}

type Plan struct {
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

type Result struct {
	Id                    string      `json:"id"`
	Name                  string      `json:"name"`
	Status                string      `json:"status"`
	Paused                bool        `json:"paused"`
	Type                  string      `json:"type"`
	Development_mode      int         `json:"development_mode"`
	Verification_key      string      `json:"verification_key"`
	Cname_suffix          string      `json:"cname_suffix"`
	Original_name_servers []string    `json:"original_name_servers"`
	Original_registrar    string      `json:"original_registrar"`
	Original_dnshost      string      `json:"original_dnshost"`
	Modified_on           string      `json:"modified_on"`
	Created_on            string      `json:"created_on"`
	Activated_on          string      `json:"activated_on"`
	Meta                  Meta        `json:"meta"`
	Owner                 Owner       `json:"owner"`
	Account               Account     `json:"account"`
	Tenant                Tenant      `json:"tenant"`
	Tenant_unit           Tenamt_unit `json:"tenant_unit"`
	Permissions           []string    `json:"permissions"`
	Plan                  Plan        `json:"plan"`
}

type Result_Info struct {
	Page        int `json:"page"`
	Per_page    int `json:"per_page"`
	Total_pages int `json:"total_pages"`
	Count       int `json:"count"`
	Total_count int `json:"total_count"`
}

//////////////////////////////////////
type ZoneListResponse struct {
	Success     bool        `json:"success"`
	Messages    []Message   `json:"messages"`
	Errors      []Error     `json:"errors"`
	Result      []Result    `json:"result"`
	Result_Info Result_Info `json:"result_info"`
}
