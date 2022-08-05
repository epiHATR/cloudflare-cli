package response

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

type DnsListResponse struct {
	Success  bool            `json:"success"`
	Messages []message       `json:"messages"`
	Errors   []error         `json:"errors"`
	Result   []dnslistResult `json:"result"`
}
