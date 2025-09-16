package link

type linkCreateRequest struct {
	Url string `json:"url" validate:"required,url"`
}
