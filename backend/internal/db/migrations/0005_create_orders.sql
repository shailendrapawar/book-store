CREATE TABLE IF NOT EXISTS orders(
    id VARCHAR(50) PRIMARY KEY,
    user_id VARCHAR(50) NOT NULL REFERENCES users(id),
    cart_id VARCHAR(50) NOT NULL REFERENCES carts(id),
    amount NUMERIC(10, 2) NOT NULL CHECK(amount > 0),
    delievery_address VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
)