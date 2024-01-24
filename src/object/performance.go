package object

type Actor struct {
	Name string `json:"name"`
	Role string `json:"role"`
}

type CRew struct {
	Name string `json:"name"`
	Role string `json:"role"`
}

// 演出
type Performance struct {
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
	OrigTitle         string             `json:"orig_title"`
	OtherTitle        []string           `json:"other_title"`
	Genre             []string           `json:"genre"`
	Language          []string           `json:"language"`
	OpeningDate       string             `json:"opening_date"`
	ClosingDate       string             `json:"closing_date"`
	Director          []string           `json:"director"`
	Playwright        []string           `json:"playwright"`
	OrigCreator       []string           `json:"orig_creator"`
	Composer          []string           `json:"composer"`
	Choreographer     []string           `json:"choreographer"`
	Performer         []string           `json:"performer"`
	Actor             []Actor            `json:"actor"`
	Crew              []CRew             `json:"crew"`
	OfficialSite      string             `json:"official_site"`
}
