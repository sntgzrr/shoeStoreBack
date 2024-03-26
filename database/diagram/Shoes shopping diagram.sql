CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    user_full_name VARCHAR(100) NOT NULL,
    user_email VARCHAR(50) NOT NULL,
    user_password VARCHAR(50) NOT NULL,
    user_created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    user_updated_at TIMESTAMP
);

CREATE TABLE products (
    product_id SERIAL PRIMARY KEY,
    product_name VARCHAR(50) NOT NULL,
    product_price INTEGER NOT NULL,
    product_amount INTEGER NOT NULL,
    product_created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    product_updated_at TIMESTAMP
);

CREATE TABLE orders (
    order_id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    order_address VARCHAR(50) NOT NULL,
    order_created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY (user_id) REFERENCES users(user_id)
);

CREATE TABLE order_products (
    order_id INT NOT NULL,
    product_id INT NOT NULL,
    quantity INT NOT NULL DEFAULT 1,
    PRIMARY KEY (order_id, product_id),
    FOREIGN KEY (order_id) REFERENCES orders(order_id),
    FOREIGN KEY (product_id) REFERENCES products(product_id)
);
