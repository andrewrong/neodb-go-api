package object

// 关于用户的结构体
type User struct {
	URL          string `json:"url"`
	ExternalAcct string `json:"external_acct"`
	DisplayName  string `json:"display_name"`
	Avatar       string `json:"avatar"`
	Username     string `json:"username"`
}
