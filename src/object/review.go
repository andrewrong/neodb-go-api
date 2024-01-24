package object

// 评论
type Review struct {
	URL         string      `json:"url"`
	Visibility  int         `json:"visibility"`
	Item        interface{} `json:"item"` // 评论的对象
	CreatedTime string      `json:"created_time"`
	Title       string      `json:"title"`
	Body        string      `json:"body"`
	HtmlContent string      `json:"html_content"`
}
