package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
	"warehouse-service/internal/domain/inventory"
	"warehouse-service/internal/service/warehouse"
	"warehouse-service/pkg/server/response"
	"warehouse-service/pkg/storage"
)

type inventoryHandler struct {
	InventoryService *warehouse.Service
}

func NewInventoryHandler(s *warehouse.Service) *inventoryHandler {
	return &inventoryHandler{InventoryService: s}
}

func (h *inventoryHandler) Routes() chi.Router {
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

// List of inventories from the database
//
//	@Summary	List of inventories from the database
//	@Tags		inventories
//	@Accept		json
//	@Produce	json
//	@Success	200				{array}		response.Object
//	@Failure	500				{object}	response.Object
//	@Router		/inventories 	[get]
func (h *inventoryHandler) list(w http.ResponseWriter, r *http.Request) {
	res, err := h.InventoryService.ListInventory(r.Context())
	if err != nil {
		response.InternalServerError(w, r, err)
		return
	}

	response.OK(w, r, res)
}

// Add a new inventory to the database
//
//	@Summary	Add a new inventory to the database
//	@Tags		inventories
//	@Accept		json
//	@Produce	json
//	@Param		request	body		inventory.Request	true	"body param"
//	@Success	200		{object}	response.Object
//	@Failure	400		{object}	response.Object
//	@Failure	500		{object}	response.Object
//	@Router		/inventories [post]
func (h *inventoryHandler) add(w http.ResponseWriter, r *http.Request) {
	req := inventory.Request{}
	if err := render.Bind(r, &req); err != nil {
		response.BadRequest(w, r, err, req)
		return
	}

	res, err := h.InventoryService.AddInventory(r.Context(), req)
	if err != nil {
		response.InternalServerError(w, r, err)
		return
	}

	response.OK(w, r, res)
}

// Read the inventory from the database
//
//	@Summary	Read the inventory from the database
//	@Tags		inventories
//	@Accept		json
//	@Produce	json
//	@Param		id	path		string	true	"path param"
//	@Success	200	{object}	response.Object
//	@Failure	404	{object}	response.Object
//	@Failure	500	{object}	response.Object
//	@Router		/inventories/{id} [get]
func (h *inventoryHandler) get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	res, err := h.InventoryService.GetInventory(r.Context(), id)
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

// Update the inventory in the database
//
//	@Summary	Update the inventory in the database
//	@Tags		inventories
//	@Accept		json
//	@Produce	json
//	@Param		id		path	string				true	"path param"
//	@Param		request	body	inventory.Request	true	"body param"
//	@Success	200
//	@Failure	400	{object}	response.Object
//	@Failure	404	{object}	response.Object
//	@Failure	500	{object}	response.Object
//	@Router		/inventories/{id} [put]
func (h *inventoryHandler) update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	req := inventory.Request{}
	if err := render.Bind(r, &req); err != nil {
		response.BadRequest(w, r, err, req)
		return
	}

	err := h.InventoryService.UpdateInventory(r.Context(), id, req)
	if err != nil && err != storage.ErrorNotFound {
		response.InternalServerError(w, r, err)
		return
	}

	if err == storage.ErrorNotFound {
		response.NotFound(w, r, err)
		return
	}
}

// Delete the inventory from the database
//
//	@Summary	Delete the inventory from the database
//	@Tags		inventories
//	@Accept		json
//	@Produce	json
//	@Param		id	path	string	true	"path param"
//	@Success	200
//	@Failure	404	{object}	response.Object
//	@Failure	500	{object}	response.Object
//	@Router		/inventories/{id} [delete]
func (h *inventoryHandler) delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := h.InventoryService.DeleteInventory(r.Context(), id)
	if err != nil && err != storage.ErrorNotFound {
		response.InternalServerError(w, r, err)
		return
	}

	if err == storage.ErrorNotFound {
		response.NotFound(w, r, err)
		return
	}
}
