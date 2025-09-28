package link

type linkCreateRequest struct {
	Url string `json:"url" validate:"required,url"`
}

type linkUpdateRequest struct {
	Url  string `json:"url" validate:"required,url"`
	Hash string
}
type GetAllLinksResponse struct {
	Links []Link `json:"links"`
	Count int64  `json:"count"`
}
