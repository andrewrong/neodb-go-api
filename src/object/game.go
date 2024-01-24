package object

type Game struct {
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
	Genre             []string           `json:"genre"`
	Developer         []string           `json:"developer"`
	Publisher         []string           `json:"publisher"`
	Platform          []string           `json:"platform"`
	ReleaseDate       string             `json:"release_date"`
	OfficialSite      string             `json:"official_site"`
}
