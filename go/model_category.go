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
	"time"
)

// The category resource is used to group product offerings, service and resource candidates in logical containers. Categories can contain other categories and/or product offerings, resource or service candidates.
type Category struct {

	// Unique identifier of the category
	Id string `json:"id,omitempty"`

	// Reference of the category
	Href string `json:"href,omitempty"`

	// Description of the category
	Description string `json:"description,omitempty"`

	// If true, this Boolean indicates that the category is a root of categories
	IsRoot bool `json:"isRoot,omitempty"`

	// Date and time of the last update
	LastUpdate time.Time `json:"lastUpdate,omitempty"`

	// Used to indicate the current lifecycle status
	LifecycleStatus string `json:"lifecycleStatus,omitempty"`

	// Name of the category
	Name string `json:"name,omitempty"`

	// Unique identifier of the parent category
	ParentId string `json:"parentId,omitempty"`

	// Category version
	Version string `json:"version,omitempty"`

	// A product offering represents entities that are orderable from the provider of the catalog, this resource includes pricing information.
	ProductOffering []ProductOfferingRef `json:"productOffering,omitempty"`

	// The category resource is used to group product offerings, service and resource candidates in logical containers. Categories can contain other (sub-)categories and/or product offerings.
	SubCategory []CategoryRef `json:"subCategory,omitempty"`

	// The period for which the category is valid
	ValidFor *TimePeriod `json:"validFor,omitempty"`

	// When sub-classing, this defines the super-class
	BaseType string `json:"@baseType,omitempty"`

	// A URI to a JSON-Schema file that defines additional attributes and relationships
	SchemaLocation string `json:"@schemaLocation,omitempty"`

	// When sub-classing, this defines the sub-class Extensible name
	Type_ string `json:"@type,omitempty"`
}
