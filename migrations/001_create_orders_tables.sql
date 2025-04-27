-- Create orders table
CREATE TABLE IF NOT EXISTS orders (
    id VARCHAR(36) PRIMARY KEY,
    status VARCHAR(50) NOT NULL,
    currency VARCHAR(3) NOT NULL,
    total_amount BIGINT NOT NULL,
    paid_amount BIGINT NOT NULL DEFAULT 0,
    description TEXT,
    created BIGINT NOT NULL,
    updated BIGINT NOT NULL,
    INDEX idx_created (created),
    INDEX idx_status (status)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Create order_products table
CREATE TABLE IF NOT EXISTS order_products (
    order_id VARCHAR(36) NOT NULL,
    product_id VARCHAR(36) NOT NULL,
    PRIMARY KEY (order_id, product_id),
    FOREIGN KEY (order_id) REFERENCES orders(id) ON DELETE CASCADE,
    INDEX idx_product_id (product_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci; 