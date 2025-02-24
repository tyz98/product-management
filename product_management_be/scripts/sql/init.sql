CREATE TABLE IF NOT EXISTS products (
      id SERIAL PRIMARY KEY,                -- auto-incrementing primary key
      name VARCHAR(255) NOT NULL,           -- product name (NOT NULL)
      type VARCHAR(255) NOT NULL,           -- product type (NOT NULL)
      price DECIMAL(10, 2) NOT NULL,        -- product price (NOT NULL)
      description TEXT,                     -- product description
      image_path VARCHAR(1024)              -- path to the uploaded image (with a limit of 1024 characters)
);
-- Create single-column indexes
CREATE INDEX IF NOT EXISTS idx_name ON products (name);
CREATE INDEX IF NOT EXISTS idx_type ON products (type);
CREATE INDEX IF NOT EXISTS idx_price ON products (price);