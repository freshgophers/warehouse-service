package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
	"warehouse-service/internal/domain/store"
	"warehouse-service/internal/service/warehouse"
	"warehouse-service/pkg/server/response"
	"warehouse-service/pkg/storage"
)

type storeHandler struct {
	StoreService *warehouse.Service
}

func NewStoreHandler(s *warehouse.Service) *storeHandler {
	return &storeHandler{StoreService: s}
}

func (h *storeHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", h.list)
	r.Post("/", h.add)

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", h.get)
		r.Put("/", h.update)
		r.Delete("/", h.delete)
	})

	return r
}

// List of stores from the database
//
//	@Summary	List of stores from the database
//	@Tags		stores
//	@Accept		json
//	@Produce	json
//	@Success	200			{array}		response.Object
//	@Failure	500			{object}	response.Object
//	@Router		/stores 	[get]
func (h *storeHandler) list(w http.ResponseWriter, r *http.Request) {
	res, err := h.StoreService.ListStores(r.Context())
	if err != nil {
		response.InternalServerError(w, r, err)
		return
	}

	response.OK(w, r, res)
}

// Add a new store to the database
//
//	@Summary	Add a new store to the database
//	@Tags		stores
//	@Accept		json
//	@Produce	json
//	@Param		request	body		store.Request	true	"body param"
//	@Success	200		{object}	response.Object
//	@Failure	400		{object}	response.Object
//	@Failure	500		{object}	response.Object
//	@Router		/stores [post]
func (h *storeHandler) add(w http.ResponseWriter, r *http.Request) {
	req := store.Request{}
	if err := render.Bind(r, &req); err != nil {
		response.BadRequest(w, r, err, req)
		return
	}

	res, err := h.StoreService.AddStore(r.Context(), req)
	if err != nil {
		response.InternalServerError(w, r, err)
		return
	}

	response.OK(w, r, res)
}

// Read the store from the database
//
//	@Summary	Read the store from the database
//	@Tags		stores
//	@Accept		json
//	@Produce	json
//	@Param		id	path		string	true	"path param"
//	@Success	200	{object}	response.Object
//	@Failure	404	{object}	response.Object
//	@Failure	500	{object}	response.Object
//	@Router		/stores/{id} [get]
func (h *storeHandler) get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	res, err := h.StoreService.GetStore(r.Context(), id)
	if err != nil && err != storage.ErrorNotFound {
		response.InternalServerError(w, r, err)
		return
	}

	if err == storage.ErrorNotFound {
		response.NotFound(w, r, err)
		return
	}

	response.OK(w, r, res)
}

// Update the store in the database
//
//	@Summary	Update the store in the database
//	@Tags		stores
//	@Accept		json
//	@Produce	json
//	@Param		id		path	string			true	"path param"
//	@Param		request	body	store.Request	true	"body param"
//	@Success	200
//	@Failure	400	{object}	response.Object
//	@Failure	404	{object}	response.Object
//	@Failure	500	{object}	response.Object
//	@Router		/stores/{id} [put]
func (h *storeHandler) update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	req := store.Request{}
	if err := render.Bind(r, &req); err != nil {
		response.BadRequest(w, r, err, req)
		return
	}

	err := h.StoreService.UpdateStore(r.Context(), id, req)
	if err != nil && err != storage.ErrorNotFound {
		response.InternalServerError(w, r, err)
		return
	}

	if err == storage.ErrorNotFound {
		response.NotFound(w, r, err)
		return
	}
}

// Delete the store from the database
//
//	@Summary	Delete the store from the database
//	@Tags		stores
//	@Accept		json
//	@Produce	json
//	@Param		id	path	string	true	"path param"
//	@Success	200
//	@Failure	404	{object}	response.Object
//	@Failure	500	{object}	response.Object
//	@Router		/stores/{id} [delete]
func (h *storeHandler) delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := h.StoreService.DeleteStore(r.Context(), id)
	if err != nil && err != storage.ErrorNotFound {
		response.InternalServerError(w, r, err)
		return
	}

	if err == storage.ErrorNotFound {
		response.NotFound(w, r, err)
		return
	}
}
