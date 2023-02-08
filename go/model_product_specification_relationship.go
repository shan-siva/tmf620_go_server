/*
 * Product Catalog Management
 *
 * ## TMF API Reference: TMF620 - Product Catalog Management  ### Release : 20.5 - March 2021  Product Catalog API is one of Catalog Management API Family. Product Catalog API goal is to provide a catalog of products.   ### Operations Product Catalog API performs the following operations on the resources : - Retrieve an entity or a collection of entities depending on filter criteria - Partial update of an entity (including updating rules) - Create an entity (including default values and creation rules) - Delete an entity - Manage notification of events
 *
 * API version: 4.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// A migration, substitution, dependency or exclusivity relationship between/among product specifications.
type ProductSpecificationRelationship struct {

	// unique identifier
	Id string `json:"id,omitempty"`

	// Hyperlink reference
	Href string `json:"href,omitempty"`

	// Name of the related entity.
	Name string `json:"name,omitempty"`

	// Type of relationship such as migration, substitution, dependency, exclusivity
	RelationshipType string `json:"relationshipType,omitempty"`

	// The period for which the productSpecificationRelationship is valid
	ValidFor *TimePeriod `json:"validFor,omitempty"`

	// When sub-classing, this defines the super-class
	BaseType string `json:"@baseType,omitempty"`

	// A URI to a JSON-Schema file that defines additional attributes and relationships
	SchemaLocation string `json:"@schemaLocation,omitempty"`

	// When sub-classing, this defines the sub-class Extensible name
	Type_ string `json:"@type,omitempty"`

	// The actual type of the target instance when needed for disambiguation.
	ReferredType string `json:"@referredType,omitempty"`
}
