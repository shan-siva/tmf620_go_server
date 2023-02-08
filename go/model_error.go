/*
 * Product Catalog Management
 *
 * ## TMF API Reference: TMF620 - Product Catalog Management  ### Release : 20.5 - March 2021  Product Catalog API is one of Catalog Management API Family. Product Catalog API goal is to provide a catalog of products.   ### Operations Product Catalog API performs the following operations on the resources : - Retrieve an entity or a collection of entities depending on filter criteria - Partial update of an entity (including updating rules) - Create an entity (including default values and creation rules) - Delete an entity - Manage notification of events
 *
 * API version: 4.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Used when an API throws an Error, typically with a HTTP error response-code (3xx, 4xx, 5xx)
type ModelError struct {

	// Application relevant detail, defined in the API or a common list.
	Code string `json:"code"`

	// Explanation of the reason for the error which can be shown to a client user.
	Reason string `json:"reason"`

	// More details and corrective actions related to the error which can be shown to a client user.
	Message string `json:"message,omitempty"`

	// HTTP Error code extension
	Status string `json:"status,omitempty"`

	// URI of documentation describing the error.
	ReferenceError string `json:"referenceError,omitempty"`

	// When sub-classing, this defines the super-class.
	BaseType string `json:"@baseType,omitempty"`

	// A URI to a JSON-Schema file that defines additional attributes and relationships
	SchemaLocation string `json:"@schemaLocation,omitempty"`

	// When sub-classing, this defines the sub-class entity name.
	Type_ string `json:"@type,omitempty"`
}
