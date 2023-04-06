-- Create movies table inside the database
CREATE TABLE movies (
  id serial not null primary key,
  title varchar(255) not null,
  description text,
  release_date date,
  runtime integer,
  rating integer,
  created_at timestamp,
  updated_at timestamp
);
