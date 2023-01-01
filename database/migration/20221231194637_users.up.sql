CREATE TABLE users (
  id serial PRIMARY KEY,
  username varchar(64) NOT NULL,
  email varchar(64) NOT NULL,
  auth_code int
);
