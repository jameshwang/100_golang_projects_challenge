-- `psql -f 001_create_users.sql <db_name>`

CREATE TABLE users (
  id serial primary key,
  name text,
  email text unique
)
