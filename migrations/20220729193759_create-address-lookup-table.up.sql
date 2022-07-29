CREATE TABLE IF NOT EXISTS address_lookups
(
    id         serial PRIMARY KEY,
    address    text UNIQUE NOT NULL,
    address_id integer     NOT NULL,
    CONSTRAINT fk_address
        FOREIGN KEY (address_id)
            REFERENCES addresses (id)
);