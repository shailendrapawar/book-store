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
	AddressType string `json:"address_type" binding:"required,oneof= home office roaming"`
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
	UserID   *string `json:"user_id"`
	State    *string `json:"state"`
	City     *string `json:"city"`
	District *string `json:"district"`
	// IsDeleted *bool   `json:"is_deleted"`
}

type SearchAddressResponse struct {
	// pagination Pagination
	Addresses []Address `json:"addresses"`
}
type UpdateAddressRequest struct {
	AddressType *string `json:"address_type" binding:"omitempty,oneof=home work other"`
	Line1       *string `json:"line1" binding:"omitempty"`
	Line2       *string `json:"line2" binding:"omitempty"`
	Landmark    *string `json:"landmark" binding:"omitempty"`
	City        *string `json:"city" binding:"omitempty"`
	Pincode     *string `json:"pincode" binding:"omitempty"`
	District    *string `json:"district" binding:"omitempty"`
	State       *string `json:"state" binding:"omitempty"`
	Country     *string `json:"country" binding:"omitempty"`
}
