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

// A new product offering price being created by value or a reference to an existing product offering price that alreasy created. The polymorphic attributes @type, @schemaLocation & @referredType are related to the product offering price and not to this ReforValue structure
type ProductOfferingPriceRefOrValue struct {

	// unique identifier
	Id string `json:"id,omitempty"`

	// Hyperlink reference
	Href string `json:"href,omitempty"`

	// Description of the productOfferingPrice
	Description string `json:"description,omitempty"`

	// the last update time of this ProductOfferingPrice
	LastUpdate time.Time `json:"lastUpdate,omitempty"`

	// the lifecycle status of this ProductOfferingPrice
	LifecycleStatus string `json:"lifecycleStatus,omitempty"`

	// Name of the productOfferingPrice
	Name string `json:"name,omitempty"`

	// A category that describes the price charge, such as recurring, penalty, One time fee and so forth.
	PriceType string `json:"priceType,omitempty"`

	// The period type to repeat the application of the price Could be month, week...
	RecurringChargePeriod string `json:"recurringChargePeriod,omitempty"`

	// the period of the recurring charge:  1, 2, ... .It sets to zero if it is not applicable
	RecurringChargePeriodLength int32 `json:"recurringChargePeriodLength,omitempty"`

	// ProductOffering version
	Version string `json:"version,omitempty"`

	// The Constraint resource represents a policy/rule applied to ProductOfferingPrice.
	Constraint []ConstraintRef `json:"constraint,omitempty"`

	Price *ProductPriceValue `json:"price,omitempty"`

	PriceAlteration []PopAlteration `json:"priceAlteration,omitempty"`

	// A number and unit representing denominator of a rate. For example, for a data charge rate of $2 per 5 GB usage, the amount of unitOfMeasure will be 5 with units as GB.
	UnitOfMeasure *Quantity `json:"unitOfMeasure,omitempty"`

	// The period for which the productOfferingPrice is valid
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
