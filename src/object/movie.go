package object

type Movie struct {
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
	Director          []string           `json:"director"`
	Playwright        []string           `json:"playwright"`
	Actor             []string           `json:"actor"`
	Genre             []string           `json:"genre"`
	Language          []string           `json:"language"`
	Area              []string           `json:"area"`
	Year              int                `json:"year"`
	Site              string             `json:"site"`
	Duration          string             `json:"duration"`
	IMDB              string             `json:"imdb"`
}
