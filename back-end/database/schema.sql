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

-- Create genres table inside the database
CREATE TABLE genres (
    id serial not null primary key,
    genre_name varchar(100) not null,
    created_at timestamp,
    updated_at timestamp
);

-- Create Join table for movies and genres
CREATE TABLE movies_genres (
    id serial not null primary key,
    movie_id integer not null,
    genre_id integer not null,
    created_at timestamp,
    updated_at timestamp,
    CONSTRAINT fk_movie_id
      FOREIGN KEY(movie_id)
      REFERENCES movies(id),
    CONSTRAINT fk_genre_id
      FOREIGN KEY(genre_id)
      REFERENCES genres(id)
);

-- Alter table movies rating column type int to float
alter table movies
  alter column rating type float4;