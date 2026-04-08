CREATE TABLE IF NOT EXISTS cart_items(
    id          VARCHAR(50)     PRIMARY KEY,
    
    cart_id     VARCHAR(50)     NOT NULL  REFERENCES carts(id) ON DELETE CASCADE,

    book_id     VARCHAR(50)     NOT NULL  REFERENCES books(id),
    quantity    INT             NOT NULL  DEFAULT 0 CHECK (quantity>=0),
    price       NUMERIC(10,2)   NOT NULL  CHECK (price>0),

    created_at  TIMESTAMP       NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP       NOT NULL DEFAULT NOW()
)