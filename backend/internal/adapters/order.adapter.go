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
}

type CreateOrderPayload struct {
	UserId    string `json:"user_id" binding:"required omitempty"`
	AddressId string `json:"address_id" binding:"required"`
	CouponId  string `json:"coupon_id" binding:"omitempty"`
}
