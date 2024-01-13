CREATE TABLE users
(
    id BIGINT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE KEY,
    password VARCHAR(100)
);

CREATE TABLE products
(
    id BIGINT PRIMARY KEY,
    product_name VARCHAR(100) NOT NULL,
    detail VARCHAR(100) NOT NULL,
    quantity INT NOT NULL
);

