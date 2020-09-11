CREATE TABLE IF NOT EXISTS users (
  uid uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
  name VARCHAR,
  email VARCHAR NOT NULL,
  password VARCHAR NOT NULL
);