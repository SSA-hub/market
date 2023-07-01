package product_model

import "github.com/lib/pq"

type InsertProductBody struct {
	Name         *string        `json:"name"`
	Price        *float64       `json:"price"`
	Count        *int64         `json:"count"`
	Manufacturer *string        `json:"manufacturer"`
	Sex          *string        `json:"sex"`
	Country      *string        `json:"country"`
	Description  *string        `json:"description"`
	Subcategory  *string        `json:"subcategory"`
	Pictures     pq.StringArray `json:"pictures"`
	Show         *bool          `json:"show"`
}

type InsertManufacturerBody struct {
	Name *string `json:"name"`
}

type InsertCategoryBody struct {
	Name *string `json:"name"`
}

type FetchProductsInput struct {
	UserId        *int64   `query:"user_id"`
	Subcategory   *string  `query:"subcategory"`
	Manufacturers []string `query:"manufacturers"`
	MinPrice      *float64 `query:"min_price"`
	MaxPrice      *float64 `query:"max_price"`
	Show          *bool    `query:"show"`
	Like          *bool    `query:"like"`
	Sort          *string  `query:"sort"`
	Sexes         []string `query:"sexes"`
	Countries     []string `query:"countries"`
	Limit         *int64   `query:"limit"`
	Offset        *int64   `query:"offset"`
}

type FetchProductsResponse struct {
	Products []ProductInfo
	Count    *int64
}

type UpdateProductCountBody struct {
	Count int64 `json:"count"`
}
