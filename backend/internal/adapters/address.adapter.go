package adapters

type Address struct {
	Id          string `json:"id"`
	UserID      string `json:"user_id" binding:"required"`
	AddressType string `json:"address_type" binding:"required"`
	Line1       string `json:"line1" binding:"required"`
	Line2       string `json:"line2"`
	Landmark    string `json:"landmark"`
	City        string `json:"city" binding:"required"`
	Pincode     string `json:"pincode" binding:"required"`
	District    string `json:"district" binding:"required"`
	State       string `json:"state" binding:"required"`
	Country     string `json:"country" binding:"required"`
}

type CreateAddressRequest struct {
	AddressType string `json:"address_type" binding:"required,oneof=home work other"`
	Line1       string `json:"line1" binding:"required"`
	Line2       string `json:"line2"`
	Landmark    string `json:"landmark"`
	City        string `json:"city" binding:"required"`
	Pincode     string `json:"pincode" binding:"required"`
	District    string `json:"district" binding:"required"`
	State       string `json:"state" binding:"required"`
	Country     string `json:"country" binding:"required"`
}

type SearchAddressFilters struct {
	UserID    *string `json:"user_id"`
	State     *string `json:"state"`
	City      *string `json:"city"`
	District  *string `json:"district"`
	IsDeleted *bool   `json:"is_deleted"`
}

type UpdateAddressRequest struct {
	AddressType *string `json:"address_type" binding:"required,omitnull"`
	Line1       *string `json:"line1" binding:"required,omitnull"`
	Line2       *string `json:"line2" binding:"omitnull"`
	Landmark    *string `json:"landmark" binding:"omitnull"`
	City        *string `json:"city" binding:"required,omitnull"`
	Pincode     *string `json:"pincode" binding:"required,omitnull"`
	District    *string `json:"district" binding:"required,omitnull"`
	State       *string `json:"state" binding:"required,omitnull"`
	Country     *string `json:"country" binding:"required,omitnull"`
}
