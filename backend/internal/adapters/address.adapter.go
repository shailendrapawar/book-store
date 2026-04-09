package adapters

type Address struct {
	AddressType string `json:"address_type" binding:"required"`
	Line1       string `json:"line1" binding:"required"`
	Line2       string `json:"line2"`
	Landmark    string `json:"landmark"`
	City        string `json:"city" binding:"required"`
	Pincode     string `json:"pincode" binding:"required"`
	District    string `json:"district" binding:"required"`
	State       string `json:"state" binding:"required"`
	Country     string `json:"country" binding:"required"`
	IsDefault   bool   `json:"is_default"`
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
	IsDefault   bool   `json:"is_default"`
}
