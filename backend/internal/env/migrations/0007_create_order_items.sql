CREATE TABLE IF NOT EXISTS order_items(
    id VARCHAR(50) PRIMARY KEY,
    order_id VARCHAR(50) NOT NULL REFERENCES orders(id) ON DELETE CASCADE,

    book_id VARCHAR(50) NOT NULL REFERENCES books(id),
    title VARCHAR(255) NOT NULL,
    price NUMERIC(10,2) NOT NULL CHECK (price>=0),
    quantity INT NOT NULL CHECK (quantity>0),

    total_price NUMERIC(10,2) NOT NULL CHECK (total_price>0),

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
)