package response

type ZoneDetailResponse struct {
	Success  bool      `json:"success"`
	Messages []message `json:"messages"`
	Errors   []error   `json:"errors"`
	Result   result    `json:"result"`
}
