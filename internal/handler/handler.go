package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/swaggo/http-swagger/v2"
	"net/url"
	"warehouse-service/docs"
	_ "warehouse-service/docs"
	"warehouse-service/internal/config"
	"warehouse-service/internal/handler/http"
	"warehouse-service/internal/service/catalogue"
	"warehouse-service/internal/service/warehouse"
	"warehouse-service/pkg/server/router"
)

type Dependencies struct {
	Configs          config.Configs
	CatalogueService *catalogue.Service
	WarehouseService *warehouse.Service
}

// Configuration is an alias for a function that will take in a pointer to a Handler and modify it
type Configuration func(h *Handler) error

// Handler is an implementation of the Handler
type Handler struct {
	dependencies Dependencies

	HTTP *chi.Mux
}

// New takes a variable amount of Configuration functions and returns a new Handler
// Each Configuration will be called in the order they are passed in
func New(d Dependencies, configs ...Configuration) (h *Handler, err error) {
	// Create the handler
	h = &Handler{
		dependencies: d,
	}

	// Apply all Configurations passed in
	for _, cfg := range configs {
		// Pass the service into the configuration function
		if err = cfg(h); err != nil {
			return
		}
	}

	return
}

// WithHTTPHandler applies a http handler to the Handler
func WithHTTPHandler() Configuration {
	return func(h *Handler) (err error) {
		// Create the http handler, if we needed parameters, such as connection strings they could be inputted here
		h.HTTP = router.New()

		docs.SwaggerInfo.BasePath = "/api/v1"
		docs.SwaggerInfo.Host = h.dependencies.Configs.HTTP.Host
		docs.SwaggerInfo.Schemes = []string{h.dependencies.Configs.HTTP.Schema}

		swaggerURL := url.URL{
			Scheme: h.dependencies.Configs.HTTP.Schema,
			Host:   h.dependencies.Configs.HTTP.Host,
			Path:   "swagger/doc.json",
		}

		h.HTTP.Get("/swagger/*", httpSwagger.Handler(
			httpSwagger.URL(swaggerURL.String()),
		))

		productHandler := http.NewProductHandler(h.dependencies.CatalogueService)
		categoryHandler := http.NewCategory(h.dependencies.CatalogueService)
		storeHandler := http.NewStore(h.dependencies.WarehouseService)

		h.HTTP.Route("/api/v1", func(r chi.Router) {
			r.Mount("/products", productHandler.Routes())
			r.Mount("/categories", categoryHandler.Routes())
			r.Mount("/stores", storeHandler.Routes())
		})

		return
	}
}
