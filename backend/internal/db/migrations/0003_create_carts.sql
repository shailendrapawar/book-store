CREATE TABLE IF NOT EXISTS carts(
     id          VARCHAR(50)     PRIMARY KEY,
     user_id     VARCHAR(50)     NOT NULL  REFERENCES users(id),

     created_at  TIMESTAMP       NOT NULL DEFAULT NOW(),
     updated_at  TIMESTAMP       NOT NULL DEFAULT NOW()
)

-- CREATE UNIQUE INDEX one_active_cart_per_user
-- ON carts(user_id)
-- WHERE status = 'active';