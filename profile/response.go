package profile

type ProfileResponse struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	ImageUrl string `json:"image_url"`
}
