CREATE TABLE IF NOT EXISTS addresses(
    id serial PRIMARY KEY,
    street_number text NOT NULL,
    street_name text NOT NULL,
    city text NOT NULL,
    state text NOT NULL,
    zipcode integer NOT NULL
);