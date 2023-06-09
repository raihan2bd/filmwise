-- Insert dummy data into the genres table
insert into genres(genre_name, created_at, updated_at)
  values
    ('Drama', '2023-04-06', '2023-04-06'),
    ('Crime', '2023-04-06', '2023-04-06'),
    ('Action', '2023-04-06', '2023-04-06'),
    ('Comic Book', '2023-04-06', '2023-04-06'),
    ('Sci-Fi', '2023-04-06', '2023-04-06'),
    ('Mystery', '2023-04-06', '2023-04-06'),
    ('Adventure', '2023-04-06', '2023-04-06'),
    ('Comedy', '2023-04-06', '2023-04-06'),
    ('Romance', '2023-04-06', '2023-04-06');

-- Insert dummy data into the movies table
insert into movies (title, description, release_date, "year", runtime, created_at, updated_at)
  values
    ('The Shawshank Redemption', 'Two imprisoned men bond over a number of years', '1994-10-14', 1994, 142, '2023-04-06', '2023-04-06'),
    ('The Pursuit of Happyness', 'Based on a true story about a man named Christopher Gardner.', '2006-12-15', 2006, 117, '2023-04-06', '2023-04-06'),
    ('The Dark Knight', 'The menace known as the Joker wreaks havoc on Gotham City.', '2008-07-18', 2008, 152, '2023-04-06', '2023-04-06'),
    ('Forrest Gump', 'Forrest Gump is a simple man with a low I.Q. but good intentions.', '1994-07-06', 1994, 142, '2023-04-06', '2023-04-06');

    -- Inset dummy data into movies_genres
    INSERT INTO public.movies_genres (movie_id,genre_id,created_at,updated_at) VALUES
	 (1,1,'2023-04-06 00:00:00.000','2023-04-06 00:00:00.000'),
	 (1,2,'2023-04-06 00:00:00.000','2023-04-06 00:00:00.000'),
	 (1,6,'2023-04-06 00:00:00.000','2023-04-06 00:00:00.000'),
	 (2,1,'2023-04-06 00:00:00.000','2023-04-06 00:00:00.000'),
	 (3,1,'2023-04-06 00:00:00.000','2023-04-06 00:00:00.000'),
	 (3,2,'2023-04-06 00:00:00.000','2023-04-06 00:00:00.000'),
	 (3,3,'2023-04-06 00:00:00.000','2023-04-06 00:00:00.000'),
	 (3,6,'2023-04-06 00:00:00.000','2023-04-06 00:00:00.000'),
	 (4,1,'2023-04-06 00:00:00.000','2023-04-06 00:00:00.000'),
	 (4,9,'2023-04-06 00:00:00.000','2023-04-06 00:00:00.000');