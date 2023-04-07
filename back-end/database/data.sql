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
insert into movies (title, description, release_date, runtime, rating, created_at, updated_at)
  values
    ('The Shawshank Redemption', 'Two imprisoned men bond over a number of years', '1994-10-14', 142, 9.3, '2023-04-06', '2023-04-06'),
    ('The Pursuit of Happyness', 'Based on a true story about a man named Christopher Gardner.', '2006-12-15', 117, 8.0, '2023-04-06', '2023-04-06'),
    ('The Dark Knight', 'The menace known as the Joker wreaks havoc on Gotham City.', '2008-07-18', 152, 9.0, '2023-04-06', '2023-04-06'),
    ('Forrest Gump', 'Forrest Gump is a simple man with a low I.Q. but good intentions.', '1994-07-06', 142, 8.8, '2023-04-06', '2023-04-06');