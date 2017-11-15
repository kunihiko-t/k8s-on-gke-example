
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  name character varying(200), 
  created_at timestamp without time zone,
  updated_at timestamp without time zone
);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE users;