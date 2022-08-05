package response

type data struct {
}

type dnsResultMeta struct {
	Auto_added bool   `json:"auto_added"`
	Source     string `json:"source"`
}

type dnsDetailResult struct {
	Id          string        `json:"id"`
	Type        string        `json:"type"`
	Name        string        `json:"name"`
	Content     string        `json:"content"`
	Proxiable   bool          `json:"proxiable"`
	Proxied     bool          `json:"proxied"`
	Ttl         int           `json:"ttl"`
	Locked      bool          `json:"locked"`
	Zone_id     string        `json:"zone_id"`
	Zone_name   string        `json:"zone_name"`
	Created_on  string        `json:"create_on"`
	Modified_on string        `json:"modified_on"`
	Data        data          `json:"data"`
	Meta        dnsResultMeta `json:"meta"`
}

type DnsDetailResponse struct {
	Success  bool            `json:"success"`
	Messages []message       `json:"messages"`
	Errors   []error         `json:"errors"`
	Result   dnsDetailResult `json:"result"`
}
