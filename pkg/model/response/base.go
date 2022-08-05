package response

type message struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type result_Info struct {
	Page        int `json:"page"`
	Per_page    int `json:"per_page"`
	Total_pages int `json:"total_pages"`
	Count       int `json:"count"`
	Total_count int `json:"total_count"`
}
