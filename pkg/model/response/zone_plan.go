package response

type PlanDetail struct {
	Id                 string `json:"id"`
	Name               string `json:"name"`
	Currency           string `json:"currency"`
	Frequency          string `json:"frequency"`
	Price              int    `json:"price"`
	Is_subscribed      bool   `json:"is_subscribed"`
	Can_subscribe      bool   `json:"can_subscribe"`
	Legacy_id          string `json:"legacy_id"`
	Legacy_discount    bool   `json:"legacy_discount"`
	Externally_managed bool   `json:"externally_managed"`
}
type ZonePlanResponse struct {
	Success  bool         `json:"success"`
	Messages []message    `json:"messages"`
	Errors   []error      `json:"errors"`
	Result   []PlanDetail `json:"result"`
}
