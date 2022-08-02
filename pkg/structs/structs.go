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

//////////////////////////////////////
