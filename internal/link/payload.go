package link

type linkCreateRequest struct {
	Url string `json:"url" validate:"required,url"`
}

type linkUpdateRequest struct {
	Url  string `json:"url" validate:"required,url"`
	Hash string
}
