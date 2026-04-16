
CREATE TABLE IF NOT EXISTS orders(
    id VARCHAR(50) PRIMARY KEY,
    user_id VARCHAR(50) NOT NULL REFERENCES users(id),

    status VARCHAR(20) NOT NULL DEFAULT 'pending'  CHECK (status IN ('pending', 'confirmed', 'shipped', 'delivered', 'cancelled')),

    discount_value NUMERIC(10,2) NOT NULL  DEFAULT 0 CHECK (discount_value>=0),
    discount_type VARCHAR(20) NOT NULL DEFAULT 'fixed' CHECK(discount_type IN ('fixed','percentage')),

    gross_amount NUMERIC(10,2) NOT NULL CHECK (gross_amount>=0),  --before discount
    net_amount NUMERIC(10,2) NOT NULL CHECK (net_amount>=0),   -- after discount
        CHECK (net_amount <= gross_amount),

    ---shipping address
    shipping_address JSONB NOT NULL,
    shipping_city VARCHAR(255) NOT NULL,
    shipping_state VARCHAR(255) NOT NULL,
    shipping_pincode VARCHAR(10) NOT NULL,

    --payments
    payment_status VARCHAR(20) NOT NULL DEFAULT 'pending'  CHECK (payment_status IN ('pending', 'paid', 'failed')),
    payment_method VARCHAR(20) NOT NULL CHECK (payment_method IN ('cod', 'online')),

    currency VARCHAR(10) NOT NULL DEFAULT 'INR',

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
)