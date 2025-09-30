package link

import (
	"fmt"
	"net/http"
	"projects/GoLinkStat/configs"
	"projects/GoLinkStat/internal/stat"
	"projects/GoLinkStat/pkg/middleware"
	"projects/GoLinkStat/pkg/request"
	"projects/GoLinkStat/pkg/response"
	"strconv"

	"gorm.io/gorm"
)

type LinkHandlerDeps struct {
	*LinkRepository
	*stat.StatRepository
	*configs.Config
}
type LinkHandler struct {
	*LinkRepository
	*stat.StatRepository
}

func NewLinkHandler(router *http.ServeMux, deps LinkHandlerDeps) {
	handler := &LinkHandler{
		LinkRepository: deps.LinkRepository,
		StatRepository: deps.StatRepository,
	}
	router.HandleFunc("POST /link", handler.Create())
	router.Handle("GET /{hash}", middleware.IsAuthed(handler.GoTo(), deps.Config))
	router.Handle("GET /link/", middleware.IsAuthed(handler.GetAll(), deps.Config))
	router.Handle("PATCH /link/{id}", middleware.IsAuthed(handler.Update(), deps.Config))
	router.HandleFunc("DELETE /link/{id}", handler.Delete())
}

func (handler *LinkHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[linkCreateRequest](&w, r)
		if err != nil {
			return
		}
		link := NewLink(body.Url)
		for {
			if existedLink, _ := handler.LinkRepository.GetByHash(link.Hash); existedLink == nil {
				break
			}
			link.GenerateHash()
		}
		createdLink, err := handler.LinkRepository.Create(link)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		response.Json(createdLink, w, 201)
	}
}
func (handler *LinkHandler) GoTo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hash := r.PathValue("hash")
		link, err := handler.LinkRepository.GetByHash(hash)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		handler.StatRepository.AddClick(link.ID)
		http.Redirect(w, r, link.Url, http.StatusTemporaryRedirect)
	}
}

func (handler *LinkHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
		if err != nil {
			http.Error(w, "Invalid limit", http.StatusBadRequest)
		}
		offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
		if err != nil {
			http.Error(w, "Invalid offset", http.StatusBadRequest)
		}
		links := handler.GetAllLinks(limit, offset) 
		count := handler.GetCount()
		response.Json(GetAllLinksResponse{
			Links: links,
			Count: count,
		}, w, 200)
	}
}

func (handler *LinkHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email, ok := r.Context().Value(middleware.ContextEmailKey).(string)
		if ok {
			fmt.Println(email)
		}
		body, err := request.HandleBody[linkUpdateRequest](&w, r)
		if err != nil {
			return
		}
		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		link, err := handler.LinkRepository.Update(&Link{
			Model: gorm.Model{ID: uint(id)},
			Url:   body.Url,
			Hash:  body.Hash,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		response.Json(link, w, 201)
	}
}
func (handler *LinkHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		link, err := handler.LinkRepository.GetById(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		err = handler.LinkRepository.Delete(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		response.Json(link, w, 200)
	}
}
