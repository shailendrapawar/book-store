CREATE TABLE IF NOT EXISTS books(

    id          VARCHAR(50)     PRIMARY KEY,
    isbn        VARCHAR(20)     NOT NULL UNIQUE,
    title       VARCHAR(255)    NOT NULL,
    author      VARCHAR(255)    NOT NULL,
    description TEXT,

    price       NUMERIC(10,2)   NOT NULL CHECK (price>0),
    
    stock       INT NOT NULL    DEFAULT 0 CHECK(stock >= 0),
    reserved    INT NOT NULL    DEFAULT 0 CHECK (reserved>=0),

    is_active   BOOLEAN         NOT NULL DEFAULT TRUE,

    created_at  TIMESTAMP       NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP       NOT NULL DEFAULT NOW(),

 CONSTRAINT stock_reserved_check CHECK (reserved <= stock)  
)