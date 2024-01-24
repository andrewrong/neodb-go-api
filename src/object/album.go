package object

type Album struct {
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
	OtherTitle        []string           `json:"other_title"`
	Genre             []string           `json:"genre"`
	Artist            []string           `json:"artist"`
	Company           []string           `json:"company"`
	Duration          int                `json:"duration"`
	ReleaseDate       string             `json:"release_date"`
	TrackList         string             `json:"track_list"`
	Barcode           string             `json:"barcode"`
}
