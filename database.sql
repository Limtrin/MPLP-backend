create TABLE products(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    price MONEY,
    instock BOOLEAN
)