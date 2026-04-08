package adapters

type Cart struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CreateCartRequest struct {
}
