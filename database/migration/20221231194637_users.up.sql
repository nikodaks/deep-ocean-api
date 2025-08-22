CREATE TABLE users (
  id serial PRIMARY KEY,
  username varchar(64) UNIQUE NOT NULL,
  email varchar(64) UNIQUE NOT NULL,
  password varchar(128) NOT NULL,
  auth_code int
);
