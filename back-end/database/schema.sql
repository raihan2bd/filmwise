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

-- Create users table inside the database
Create TABLE users (
  id serial not null primary key,
  name varchar(100) not null,
  email varchar(255) not null unique,
  password varchar(60) not null,
  user_type varchar(55) not null default 'user',
  created_at timestamp,
  updated_at timestamp
);

-- Create ratings table inside the database
CREATE TABLE ratings (
    id serial not null primary key,
    movie_id integer not null,
    user_id integer not null,
    rating float4 not null,
    created_at timestamp,
    updated_at timestamp,
    CONSTRAINT fk_movie_id
      FOREIGN KEY(movie_id)
      REFERENCES movies(id)
      ON DELETE CASCADE,
    CONSTRAINT fk_user_id
      FOREIGN KEY(user_id)
      REFERENCES users(id)
      ON DELETE CASCADE
);

-- Create favorites table inside the database
CREATE TABLE favorites (
    id serial not null primary key,
    movie_id integer not null,
    user_id integer not null,
    created_at timestamp,
    updated_at timestamp,
    CONSTRAINT fk_movie_id
      FOREIGN KEY(movie_id)
      REFERENCES movies(id)
      ON DELETE CASCADE,
    CONSTRAINT fk_user_id
      FOREIGN KEY(user_id)
      REFERENCES users(id)
      ON DELETE CASCADE
);

-- Create comments table inside the database
CREATE TABLE comments (
    id serial not null primary key,
    movie_id integer not null,
    user_id integer not null,
    comment text not null,
    created_at timestamp,
    updated_at timestamp,
    CONSTRAINT fk_movie_id
      FOREIGN KEY(movie_id)
      REFERENCES movies(id)
      ON DELETE CASCADE,
    CONSTRAINT fk_user_id
      FOREIGN KEY(user_id)
      REFERENCES users(id)
      ON DELETE CASCADE
);

-- Alter table movies drop column rating
ALTER TABLE movies DROP COLUMN rating;

-- Create table images
CREATE TABLE images (
    id serial not null primary key,
    user_id integer not null,
    image_path varchar(255) not null,
    image_name varchar(255) not null,
    is_used boolean not null default false,
    created_at timestamp,
    updated_at timestamp
);

-- Alter table movies add column image
ALTER TABLE movies ADD COLUMN image varchar(255);