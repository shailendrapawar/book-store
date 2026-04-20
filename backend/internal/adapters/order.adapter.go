package adapters

import "time"

type Order struct {
	Id     string `json:"id"`
	UserId string `json:"user_id"`

	Status string `json:"status"`

	DiscountValue float64 `json:"discount_value"`
	DiscountType  string  `json:"discount_type"`

	GrossAmount float64 `json:"gross_amount"`
	NetAmount   float64 `json:"net_amount"`

	ShippingAddress string `json:"shipping_address"`
	ShippingCity    string `json:"shipping_city"`
	ShippingState   string `json:"shipping_state"`
	ShippingPincode string `json:"shipping_pincode"`

	PaymentMethod string `json:"payment_method"`
	PaymentStatus string `json:"payment_status"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Items []*OrderItem `json:"items"`
}

type CreateOrderRequest struct {
	AddressId     string `json:"address_id" binding:"required"`
	CouponId      string `json:"coupon_id" binding:"omitempty"`
	Notes         string `json:"notes" binding:"omitempty"`
	PaymentMethod string `json:"payment_method" binding:"required"`
}
type CreateOrderPayload struct {
	User          *Claims  `json:"user"`
	Address       *Address `json:"address"`
	Coupon        string   `json:"coupon"`
	Books         []*Book  `json:"books"`
	GrossAmount   float64  `json:"gross_amount"`
	NetAmount     float64  `json:"net_amount"`
	PaymentMethod string   `json:"payment_method"`
	Currency      string   `json:"currency"`
}

type SearchOrderFilters struct {
	UserID *string `json:"user_id"`
}
