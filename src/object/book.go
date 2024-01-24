package object

type ExternalResource struct {
	URL string `json:"url"`
}

type ShelfType string

var (
	Reading  ShelfType = "wishlist"
	Read     ShelfType = "progress"
	Complete ShelfType = "complete"
	Dropped  ShelfType = "dropped"
)

type Tag string

type Visibility int

var (
	Private  Visibility = 0
	Public   Visibility = 1
	Unlisted Visibility = 2
)

type Book struct {
	ID                string             `json:"id"`
	Type              string             `json:"type"`
	UUID              string             `json:"uuid"`
	URL               string             `json:"url"`
	APIURL            string             `json:"api_url"`
	Category          string             `json:"category"`
	ParentUUID        string             `json:"parent_uuid"`
	DisplayTitle      string             `json:"display_title"`
	ExternalResources []ExternalResource `json:"external_resources"`
	Title             string             `json:"title"`
	Brief             string             `json:"brief"`
	CoverImageURL     string             `json:"cover_image_url"`
	Rating            int                `json:"rating"`
	RatingCount       int                `json:"rating_count"`
	Subtitle          string             `json:"subtitle"`
	OrigTitle         string             `json:"orig_title"`
	Author            []string           `json:"author"`
	Translator        []string           `json:"translator"`
	Language          string             `json:"language"`
	PubHouse          string             `json:"pub_house"`
	PubYear           int                `json:"pub_year"`
	PubMonth          int                `json:"pub_month"`
	Binding           string             `json:"binding"`
	Price             string             `json:"price"`
	Pages             int                `json:"pages"`
	Series            string             `json:"series"`
	Imprint           string             `json:"imprint"`
	ISBN              string             `json:"isbn"`
}

type ShelfItem struct {
	ShelfT      ShelfType  `json:"shelf_type"`
	Visibility  Visibility `json:"visibility"`
	Item        Book       `json:"item"`
	CreatedTime string     `json:"created_time"`
	CommentText string     `json:"comment_text"`
	RatingGrade int        `json:"rating_grade"`
	Tags        []string   `json:"tags"`
}
