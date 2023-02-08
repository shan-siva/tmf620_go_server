/*
 * Product Catalog Management
 *
 * ## TMF API Reference: TMF620 - Product Catalog Management  ### Release : 20.5 - March 2021  Product Catalog API is one of Catalog Management API Family. Product Catalog API goal is to provide a catalog of products.   ### Operations Product Catalog API performs the following operations on the resources : - Retrieve an entity or a collection of entities depending on filter criteria - Partial update of an entity (including updating rules) - Create an entity (including default values and creation rules) - Delete an entity - Manage notification of events
 *
 * API version: 4.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/tmf-api/productCatalogManagement/v4//",
		Index,
	},

	Route{
		"CreateCatalog",
		strings.ToUpper("Post"),
		"/tmf-api/productCatalogManagement/v4/catalog",
		CreateCatalog,
	},

	Route{
		"DeleteCatalog",
		strings.ToUpper("Delete"),
		"/tmf-api/productCatalogManagement/v4/catalog/{id}",
		DeleteCatalog,
	},

	Route{
		"ListCatalog",
		strings.ToUpper("Get"),
		"/tmf-api/productCatalogManagement/v4/catalog",
		ListCatalog,
	},

	Route{
		"PatchCatalog",
		strings.ToUpper("Patch"),
		"/tmf-api/productCatalogManagement/v4/catalog/{id}",
		PatchCatalog,
	},

	Route{
		"RetrieveCatalog",
		strings.ToUpper("Get"),
		"/tmf-api/productCatalogManagement/v4/catalog/{id}",
		RetrieveCatalog,
	},

	Route{
		"CreateCategory",
		strings.ToUpper("Post"),
		"/tmf-api/productCatalogManagement/v4/category",
		CreateCategory,
	},

	Route{
		"DeleteCategory",
		strings.ToUpper("Delete"),
		"/tmf-api/productCatalogManagement/v4/category/{id}",
		DeleteCategory,
	},

	Route{
		"ListCategory",
		strings.ToUpper("Get"),
		"/tmf-api/productCatalogManagement/v4/category",
		ListCategory,
	},

	Route{
		"PatchCategory",
		strings.ToUpper("Patch"),
		"/tmf-api/productCatalogManagement/v4/category/{id}",
		PatchCategory,
	},

	Route{
		"RetrieveCategory",
		strings.ToUpper("Get"),
		"/tmf-api/productCatalogManagement/v4/category/{id}",
		RetrieveCategory,
	},

	Route{
		"RegisterListener",
		strings.ToUpper("Post"),
		"/tmf-api/productCatalogManagement/v4/hub",
		RegisterListener,
	},

	Route{
		"UnregisterListener",
		strings.ToUpper("Delete"),
		"/tmf-api/productCatalogManagement/v4/hub/{id}",
		UnregisterListener,
	},

	Route{
		"CreateExportJob",
		strings.ToUpper("Post"),
		"/tmf-api/productCatalogManagement/v4/exportJob",
		CreateExportJob,
	},

	Route{
		"DeleteExportJob",
		strings.ToUpper("Delete"),
		"/tmf-api/productCatalogManagement/v4/exportJob/{id}",
		DeleteExportJob,
	},

	Route{
		"ListExportJob",
		strings.ToUpper("Get"),
		"/tmf-api/productCatalogManagement/v4/exportJob",
		ListExportJob,
	},

	Route{
		"RetrieveExportJob",
		strings.ToUpper("Get"),
		"/tmf-api/productCatalogManagement/v4/exportJob/{id}",
		RetrieveExportJob,
	},

	Route{
		"CreateImportJob",
		strings.ToUpper("Post"),
		"/tmf-api/productCatalogManagement/v4/importJob",
		CreateImportJob,
	},

	Route{
		"DeleteImportJob",
		strings.ToUpper("Delete"),
		"/tmf-api/productCatalogManagement/v4/importJob/{id}",
		DeleteImportJob,
	},

	Route{
		"ListImportJob",
		strings.ToUpper("Get"),
		"/tmf-api/productCatalogManagement/v4/importJob",
		ListImportJob,
	},

	Route{
		"RetrieveImportJob",
		strings.ToUpper("Get"),
		"/tmf-api/productCatalogManagement/v4/importJob/{id}",
		RetrieveImportJob,
	},

	Route{
		"ListenToCatalogAttributeValueChangeEvent",
		strings.ToUpper("Post"),
		"/tmf-api/productCatalogManagement/v4/listener/catalogAttributeValueChangeEvent",
		ListenToCatalogAttributeValueChangeEvent,
	},

	Route{
		"ListenToCatalogBatchEvent",
		strings.ToUpper("Post"),
		"/tmf-api/productCatalogManagement/v4/listener/catalogBatchEvent",
		ListenToCatalogBatchEvent,
	},

	Route{
		"ListenToCatalogCreateEvent",
		strings.ToUpper("Post"),
		"/tmf-api/productCatalogManagement/v4/listener/catalogCreateEvent",
		ListenToCatalogCreateEvent,
	},

	Route{
		"ListenToCatalogDeleteEvent",
		strings.ToUpper("Post"),
		"/tmf-api/productCatalogManagement/v4/listener/catalogDeleteEvent",
		ListenToCatalogDeleteEvent,
	},

	Route{
		"ListenToCatalogStateChangeEvent",
		strings.ToUpper("Post"),
		"/tmf-api/productCatalogManagement/v4/listener/catalogStateChangeEvent",
		ListenToCatalogStateChangeEvent,
	},

	Route{
		"ListenToCategoryAttributeValueChangeEvent",
		strings.ToUpper("Post"),
		"/tmf-api/productCatalogManagement/v4/listener/categoryAttributeValueChangeEvent",
		ListenToCategoryAttributeValueChangeEvent,
	},

	Route{
		"ListenToCategoryCreateEvent",
		strings.ToUpper("Post"),
		"/tmf-api/productCatalogManagement/v4/listener/categoryCreateEvent",
		ListenToCategoryCreateEvent,
	},

	Route{
		"ListenToCategoryDeleteEvent",
		strings.ToUpper("Post"),
		"/tmf-api/productCatalogManagement/v4/listener/categoryDeleteEvent",
		ListenToCategoryDeleteEvent,
	},

	Route{
		"ListenToCategoryStateChangeEvent",
		strings.ToUpper("Post"),
		"/tmf-api/productCatalogManagement/v4/listener/categoryStateChangeEvent",
		ListenToCategoryStateChangeEvent,
	},

	Route{
		"ListenToProductOfferingAttributeValueChangeEvent",
		strings.ToUpper("Post"),
		"/tmf-api/productCatalogManagement/v4/listener/productOfferingAttributeValueChangeEvent",
		ListenToProductOfferingAttributeValueChangeEvent,
	},

	Route{
		"ListenToProductOfferingCreateEvent",
		strings.ToUpper("Post"),
		"/tmf-api/productCatalogManagement/v4/listener/productOfferingCreateEvent",
		ListenToProductOfferingCreateEvent,
	},

	Route{
		"ListenToProductOfferingDeleteEvent",
		strings.ToUpper("Post"),
		"/tmf-api/productCatalogManagement/v4/listener/productOfferingDeleteEvent",
		ListenToProductOfferingDeleteEvent,
	},

	Route{
		"ListenToProductOfferingPriceAttributeValueChangeEvent",
		strings.ToUpper("Post"),
		"/tmf-api/productCatalogManagement/v4/listener/productOfferingPriceAttributeValueChangeEvent",
		ListenToProductOfferingPriceAttributeValueChangeEvent,
	},

	Route{
		"ListenToProductOfferingPriceCreateEvent",
		strings.ToUpper("Post"),
		"/tmf-api/productCatalogManagement/v4/listener/productOfferingPriceCreateEvent",
		ListenToProductOfferingPriceCreateEvent,
	},

	Route{
		"ListenToProductOfferingPriceDeleteEvent",
		strings.ToUpper("Post"),
		"/tmf-api/productCatalogManagement/v4/listener/productOfferingPriceDeleteEvent",
		ListenToProductOfferingPriceDeleteEvent,
	},

	Route{
		"ListenToProductOfferingPriceStateChangeEvent",
		strings.ToUpper("Post"),
		"/tmf-api/productCatalogManagement/v4/listener/productOfferingPriceStateChangeEvent",
		ListenToProductOfferingPriceStateChangeEvent,
	},

	Route{
		"ListenToProductOfferingStateChangeEvent",
		strings.ToUpper("Post"),
		"/tmf-api/productCatalogManagement/v4/listener/productOfferingStateChangeEvent",
		ListenToProductOfferingStateChangeEvent,
	},

	Route{
		"ListenToProductSpecificationAttributeValueChangeEvent",
		strings.ToUpper("Post"),
		"/tmf-api/productCatalogManagement/v4/listener/productSpecificationAttributeValueChangeEvent",
		ListenToProductSpecificationAttributeValueChangeEvent,
	},

	Route{
		"ListenToProductSpecificationCreateEvent",
		strings.ToUpper("Post"),
		"/tmf-api/productCatalogManagement/v4/listener/productSpecificationCreateEvent",
		ListenToProductSpecificationCreateEvent,
	},

	Route{
		"ListenToProductSpecificationDeleteEvent",
		strings.ToUpper("Post"),
		"/tmf-api/productCatalogManagement/v4/listener/productSpecificationDeleteEvent",
		ListenToProductSpecificationDeleteEvent,
	},

	Route{
		"ListenToProductSpecificationStateChangeEvent",
		strings.ToUpper("Post"),
		"/tmf-api/productCatalogManagement/v4/listener/productSpecificationStateChangeEvent",
		ListenToProductSpecificationStateChangeEvent,
	},

	Route{
		"CreateProductOffering",
		strings.ToUpper("Post"),
		"/tmf-api/productCatalogManagement/v4/productOffering",
		CreateProductOffering,
	},

	Route{
		"DeleteProductOffering",
		strings.ToUpper("Delete"),
		"/tmf-api/productCatalogManagement/v4/productOffering/{id}",
		DeleteProductOffering,
	},

	Route{
		"ListProductOffering",
		strings.ToUpper("Get"),
		"/tmf-api/productCatalogManagement/v4/productOffering",
		ListProductOffering,
	},

	Route{
		"PatchProductOffering",
		strings.ToUpper("Patch"),
		"/tmf-api/productCatalogManagement/v4/productOffering/{id}",
		PatchProductOffering,
	},

	Route{
		"RetrieveProductOffering",
		strings.ToUpper("Get"),
		"/tmf-api/productCatalogManagement/v4/productOffering/{id}",
		RetrieveProductOffering,
	},

	Route{
		"CreateProductOfferingPrice",
		strings.ToUpper("Post"),
		"/tmf-api/productCatalogManagement/v4/productOfferingPrice",
		CreateProductOfferingPrice,
	},

	Route{
		"DeleteProductOfferingPrice",
		strings.ToUpper("Delete"),
		"/tmf-api/productCatalogManagement/v4/productOfferingPrice/{id}",
		DeleteProductOfferingPrice,
	},

	Route{
		"ListProductOfferingPrice",
		strings.ToUpper("Get"),
		"/tmf-api/productCatalogManagement/v4/productOfferingPrice",
		ListProductOfferingPrice,
	},

	Route{
		"PatchProductOfferingPrice",
		strings.ToUpper("Patch"),
		"/tmf-api/productCatalogManagement/v4/productOfferingPrice/{id}",
		PatchProductOfferingPrice,
	},

	Route{
		"RetrieveProductOfferingPrice",
		strings.ToUpper("Get"),
		"/tmf-api/productCatalogManagement/v4/productOfferingPrice/{id}",
		RetrieveProductOfferingPrice,
	},

	Route{
		"CreateProductSpecification",
		strings.ToUpper("Post"),
		"/tmf-api/productCatalogManagement/v4/productSpecification",
		CreateProductSpecification,
	},

	Route{
		"DeleteProductSpecification",
		strings.ToUpper("Delete"),
		"/tmf-api/productCatalogManagement/v4/productSpecification/{id}",
		DeleteProductSpecification,
	},

	Route{
		"ListProductSpecification",
		strings.ToUpper("Get"),
		"/tmf-api/productCatalogManagement/v4/productSpecification",
		ListProductSpecification,
	},

	Route{
		"PatchProductSpecification",
		strings.ToUpper("Patch"),
		"/tmf-api/productCatalogManagement/v4/productSpecification/{id}",
		PatchProductSpecification,
	},

	Route{
		"RetrieveProductSpecification",
		strings.ToUpper("Get"),
		"/tmf-api/productCatalogManagement/v4/productSpecification/{id}",
		RetrieveProductSpecification,
	},
}
