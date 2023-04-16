-- Create movies table inside the database
CREATE TABLE movies (
  id serial not null primary key,
  title varchar(255) not null unique,
  description text not null,
  release_date date,
  year integer not null,
  runtime integer not null,
  rating integer not null,
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

-- Delete movie_id foreign key
ALTER TABLE movies_genres DROP CONSTRAINT fk_movie_id;

-- Add movie foreign key with on delete cascade
ALTER TABLE movies_genres ADD CONSTRAINT fk_movie_id FOREIGN KEY (movie_id) REFERENCES movies (id) ON DELETE CASCADE;

-- Delete genre_id foreign key
ALTER TABLE movies_genres DROP CONSTRAINT fk_genre_id;

-- Add movie foreign key with on delete cascade
ALTER TABLE movies_genres ADD CONSTRAINT fk_genre_id FOREIGN KEY (genre_id) REFERENCES genres (id) ON DELETE CASCADE;