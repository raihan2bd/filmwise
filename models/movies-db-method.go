package models

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"time"
)

func (m *DBModel) GetAllMovies(findByName string) ([]*Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var dbArgs []interface{}
	where := " WHERE title ILIKE $1 OR description ILIKE $2"
	dbArgs = append(dbArgs, "%"+findByName+"%", "%"+findByName+"%")

	q2 := ` LEFT JOIN ratings r ON (r.movie_id = m.id)`
	q3 := ` GROUP BY m.id, m.title, m.description, m.year, m.release_date, m.runtime, m.created_at, m.updated_at
	order by rating desc limit 2 offset 1`

	query := `SELECT m.id, m.title, m.description, m.year, m.release_date, COALESCE(trunc(AVG(r.rating)::numeric, 1), 1.0) AS rating, m.runtime, m.created_at, m.updated_at 
	FROM movies m`

	query += q2 + where + q3

	rows, err := m.DB.QueryContext(ctx, query, dbArgs...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []*Movie
	for rows.Next() {
		var movie Movie
		err = rows.Scan(
			&movie.ID,
			&movie.Title,
			&movie.Description,
			&movie.Year,
			&movie.ReleaseDate,
			&movie.Rating,
			&movie.Runtime,
			&movie.CreatedAt,
			&movie.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		// // get genres, if any
		// genreQuery := `select
		// 	mg.id, mg.movie_id, mg.genre_id, g.genre_name
		// from
		// 	movies_genres mg
		// 	left join genres g on (g.id = mg.genre_id)
		// where
		// 	mg.movie_id = $1
		// `

		// genreRows, _ := m.DB.QueryContext(ctx, genreQuery, movie.ID)

		// genres := make(map[int]string)
		// for genreRows.Next() {
		// 	var mg MovieGenre
		// 	err := genreRows.Scan(
		// 		&mg.ID,
		// 		&mg.MovieID,
		// 		&mg.GenreID,
		// 		&mg.Genre.GenreName,
		// 	)
		// 	if err != nil {
		// 		return nil, err
		// 	}
		// 	genres[mg.GenreID] = mg.Genre.GenreName
		// }
		// genreRows.Close()

		// movie.MovieGenre = genres
		movies = append(movies, &movie)
	}

	return movies, nil
}

// CheckGenre checks if genre exists
func (m *DBModel) CheckGenre(genreID int) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id from genres where id = $1`

	var id int
	err := m.DB.QueryRowContext(ctx, query, genreID).Scan(&id)
	if err != nil {
		return false, err
	}

	if id <= 0 {
		return false, errors.New("Genre not found")
	}

	return true, nil
}

// InsertGenre inserts a new genre into the database
func (m *DBModel) InsertGenre(genreName string) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `insert into genres (genre_name, created_at, updated_at) values ($1, $2, $3) returning id`

	var id int
	err := m.DB.QueryRowContext(ctx, query, genreName, time.Now(), time.Now()).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// UpdateGenre updates a genre in the database
func (m *DBModel) UpdateGenre(id int, genreName string) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `update genres set genre_name = $1, updated_at = $2 where id = $3`

	_, err := m.DB.ExecContext(ctx, query, genreName, time.Now(), id)
	if err != nil {
		return id, err
	}

	return id, nil
}

// DeleteGenre deletes a genre from the database
func (m *DBModel) DeleteGenre(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `delete from genres where id = $1`

	_, err := m.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

// Check rating
func (m *DBModel) CheckRating(movieID, userID int) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id from ratings where movie_id = $1 and user_id = $2`

	ratingID := 0

	err := m.DB.QueryRowContext(ctx, query, movieID, userID).Scan(&ratingID)
	if err != nil {
		return 0, errors.New("Rating not found")
	}

	return ratingID, nil
}

// InsertRating inserts a new rating into the database
func (m *DBModel) InsertRating(rating *Rating) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `insert into ratings (movie_id, user_id, rating, created_at, updated_at) values ($1, $2, $3, $4, $5) returning id`

	var id int
	err := m.DB.QueryRowContext(ctx, query, rating.MovieID, rating.UserID, rating.Rating, rating.CreatedAt, rating.UpdatedAt).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

// UpdateRating updates a rating in the database
func (m *DBModel) UpdateRating(rating *Rating) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `update ratings set rating = $1, updated_at = $2 where id = $3`

	_, err := m.DB.ExecContext(ctx, query, rating.Rating, rating.UpdatedAt, rating.ID)
	if err != nil {
		return rating.ID, errors.New("Rating not found")
	}

	return rating.ID, nil
}

// GenreByID returns a single genre based on the ID provided
func (m *DBModel) GenreByID(id int) (Genre, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, genre_name, created_at, updated_at from genres where id = $1`

	var g Genre
	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&g.ID,
		&g.GenreName,
		&g.CreatedAt,
		&g.UpdatedAt,
	)
	if err != nil {
		return g, err
	}

	return g, nil
}

// Get all genres from db
func (m *DBModel) GenresAll() ([]*Genre, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, genre_name, created_at, updated_at from genres order by genre_name`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var genres []*Genre

	for rows.Next() {
		var g Genre
		err := rows.Scan(
			&g.ID,
			&g.GenreName,
			&g.CreatedAt,
			&g.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		genres = append(genres, &g)
	}

	return genres, nil
}

// GetFeatureMovies fetches the latest 5 featured movies from the database ordered by their update time
func (m *DBModel) GetFeatureMovies(userID ...int) ([]*Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Retrieve latest 5 featured movies ordered by update time
	query := `
		SELECT m.id, m.title, m.image, m.description, m.year, m.release_date,
		COALESCE(TRUNC(AVG(r.rating)::numeric, 1), 1.0) AS rating,
		m.runtime, m.created_at, m.updated_at
		FROM movies m
		LEFT JOIN ratings r ON r.movie_id = m.id
		GROUP BY m.id
		ORDER BY m.updated_at DESC
		LIMIT 5
	`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []*Movie
	var image sql.NullString
	for rows.Next() {
		var movie Movie
		err = rows.Scan(
			&movie.ID,
			&movie.Title,
			&image,
			&movie.Description,
			&movie.Year,
			&movie.ReleaseDate,
			&movie.Rating,
			&movie.Runtime,
			&movie.CreatedAt,
			&movie.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		// Check if the Image value is NULL or empty, and if it is, assign a default value
		if !image.Valid || image.String == "" {
			movie.Image = fmt.Sprintf("https://res.cloudinary.com/%s/image/upload/no-thumb.jpg", os.Getenv("CLOUD_NAME"))
		} else {
			movie.Image = fmt.Sprintf("https://res.cloudinary.com/%s/image/upload/%s", os.Getenv("CLOUD_NAME"), image.String)
		}

		// get genres, if any
		genreQuery := `select
			mg.id, mg.movie_id, mg.genre_id, g.genre_name
		from
			movies_genres mg
			left join genres g on (g.id = mg.genre_id)
		where
			mg.movie_id = $1
		`

		genreRows, _ := m.DB.QueryContext(ctx, genreQuery, movie.ID)

		genres := make(map[int]string)
		for genreRows.Next() {
			var mg MovieGenre
			err := genreRows.Scan(
				&mg.ID,
				&mg.MovieID,
				&mg.GenreID,
				&mg.Genre.GenreName,
			)
			if err != nil {
				return nil, err
			}
			genres[mg.GenreID] = mg.Genre.GenreName
		}
		genreRows.Close()

		if len(userID) > 0 {
			// check if movie is favorite
			favoriteQuery := `select id from favorites where movie_id = $1 and user_id = $2`
			_ = m.DB.QueryRowContext(ctx, favoriteQuery, movie.ID, userID[0]).Scan(&movie.IsFavorite)
		}

		movie.MovieGenre = genres
		movies = append(movies, &movie)
	}

	return movies, nil
}

// Get all movies by filter
func (m *DBModel) GetAllMoviesByFilter(page, perPage int, filter *MovieFilter, userID ...int) (*PaginatedMovies, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	offset := (page - 1) * perPage

	// add where query according to filter condition
	var dbArgs []interface{}
	where := " WHERE (title ILIKE $1 OR description ILIKE $2)"
	dbArgs = append(dbArgs, "%"+filter.FindByName+"%", "%"+filter.FindByName+"%")

	if filter.FilterByYear > 0 {
		where += " and year = $3"
		dbArgs = append(dbArgs, filter.FilterByYear)
	}

	if filter.FilterByGenre > 0 && filter.FilterByYear <= 0 {
		where += " and m.id in (select movie_id from movies_genres where genre_id = $3)"
		dbArgs = append(dbArgs, filter.FilterByGenre)
	}

	if (filter.FilterByGenre > 0) && (filter.FilterByYear > 0) {
		where += " and m.id in (select movie_id from movies_genres where genre_id = $4)"
		dbArgs = append(dbArgs, filter.FilterByGenre)
	}

	//	add order by query
	orderByQuery := ""
	switch filter.OrderBy {
	case "rating":
		orderByQuery = " order by rating desc"
	case "runtime":
		orderByQuery = " order by runtime desc"
	case "old":
		orderByQuery = " order by updated_at asc"
	case "name":
		orderByQuery = " order by title asc"
	default:
		orderByQuery = " order by updated_at desc"
	}

	// main query
	query := `SELECT
		m.id, 
		m.title,
		m.image, 
		m.description, 
		m.year, 
		m.release_date, 
		COALESCE(TRUNC(AVG(r.rating)::numeric, 1), 1.0) AS rating, 
		m.runtime, 
		m.created_at, 
		m.updated_at,
		COUNT(DISTINCT c.id) AS comments_count,
		COUNT(DISTINCT f.id) AS favorites_count
	FROM 
		movies m
		LEFT JOIN ratings r ON r.movie_id = m.id
		LEFT JOIN comments c ON c.movie_id = m.id
		LEFT JOIN favorites f ON f.movie_id = m.id`

	groupBYQuery := ` GROUP BY
			m.id`

	// pagination query
	paginationQuery := fmt.Sprintf(" limit %d offset %d", perPage, offset)

	countQuery := "SELECT COUNT(DISTINCT m.id) FROM movies m" + where
	var totalCount int
	err := m.DB.QueryRowContext(ctx, countQuery, dbArgs...).Scan(&totalCount)
	if err != nil {
		return nil, err
	}

	// join all the query
	query += where + groupBYQuery + orderByQuery + paginationQuery

	// execute query with context
	rows, err := m.DB.QueryContext(ctx, query, dbArgs...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []*Movie
	var image sql.NullString
	for rows.Next() {
		var movie Movie
		err = rows.Scan(
			&movie.ID,
			&movie.Title,
			&image,
			&movie.Description,
			&movie.Year,
			&movie.ReleaseDate,
			&movie.Rating,
			&movie.Runtime,
			&movie.CreatedAt,
			&movie.UpdatedAt,
			&movie.TotalComments,
			&movie.TotalFavorites,
		)

		if err != nil {
			return nil, err
		}

		// Check if the Image value is NULL or empty, and if it is, assign a default value
		if !image.Valid || image.String == "" {
			movie.Image = fmt.Sprintf("https://res.cloudinary.com/%s/image/upload/no-thumb.jpg", os.Getenv("CLOUD_NAME"))
		} else {
			movie.Image = fmt.Sprintf("https://res.cloudinary.com/%s/image/upload/%s", os.Getenv("CLOUD_NAME"), image.String)
		}

		// get genres, if any
		genreQuery := `select
			mg.id, mg.movie_id, mg.genre_id, g.genre_name
		from
			movies_genres mg
			left join genres g on (g.id = mg.genre_id)
		where
			mg.movie_id = $1
		`

		genreRows, _ := m.DB.QueryContext(ctx, genreQuery, movie.ID)

		genres := make(map[int]string)
		for genreRows.Next() {
			var mg MovieGenre
			err := genreRows.Scan(
				&mg.ID,
				&mg.MovieID,
				&mg.GenreID,
				&mg.Genre.GenreName,
			)
			if err != nil {
				return nil, err
			}
			genres[mg.GenreID] = mg.Genre.GenreName
		}
		genreRows.Close()

		if len(userID) > 0 {
			// check if movie is favorite
			favoriteQuery := `select id from favorites where movie_id = $1 and user_id = $2`
			var favID int
			_ = m.DB.QueryRowContext(ctx, favoriteQuery, movie.ID, userID[0]).Scan(&favID)

			if favID > 0 {
				movie.IsFavorite = true
			}
		}

		movie.MovieGenre = genres
		movies = append(movies, &movie)
	}

	// Create and return the PaginatedMovies struct
	paginatedMovies := &PaginatedMovies{
		TotalCount:  totalCount,
		PerPage:     perPage,
		CurrentPage: page,
		Movies:      movies,
	}

	return paginatedMovies, nil
}

// InsertMovie is help to insert new movie to the database
func (m *DBModel) InsertMovie(movie *Movie) (int, map[int]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	movieID := 0
	var movieGenres = make(map[int]string)

	// return if movie title is already exist
	q := `select id from movies where title = $1`
	_ = m.DB.QueryRowContext(ctx, q, movie.Title).Scan(&movieID)
	if movieID > 0 {
		return movieID, movieGenres, errors.New("the movie is already exist")
	}

	// verify genres
	for _, val := range movie.MovieGenre {
		genreID := 0
		query := `select id from genres where genre_name = $1`
		err := m.DB.QueryRowContext(ctx, query, val).Scan(&genreID)
		if err != nil {
			return movieID, movieGenres, errors.New("invalid genre name")
		}
		if genreID > 0 {
			if _, exists := movieGenres[genreID]; !exists {
				movieGenres[genreID] = val
			}
		}
	}

	if len(movieGenres) <= 0 {
		movieGenres[10] = "Unknown"
	}

	stmt := `insert into movies (title, description, year, release_date, runtime, image,
		created_at, updated_at) values ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id`

	err := m.DB.QueryRowContext(ctx, stmt,
		movie.Title,
		movie.Description,
		movie.Year,
		movie.ReleaseDate,
		movie.Runtime,
		movie.Image,
		time.Now(),
		time.Now(),
	).Scan(&movieID)
	if err != nil {
		log.Println(err)
		return movieID, movieGenres, err
		// return movieID, movieGenres, errors.New("invalid movie data! failed to save the movie")
	}

	for key := range movieGenres {
		stmt = `insert into movies_genres ( genre_id, movie_id, created_at, updated_at)
						values($1, $2, $3, $4)`
		_, err = m.DB.ExecContext(ctx, stmt, key, movieID, time.Now(), time.Now())
		if err != nil {
			return movieID, movieGenres, err
		}
	}
	return movieID, movieGenres, nil
}

// UpdateMovie is help to update a movie from the database
func (m *DBModel) UpdateMovie(movie *Movie) (int, map[int]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	movieID := 0
	var movieGenres = make(map[int]string)

	// verify genres
	for _, val := range movie.MovieGenre {
		genreID := 0
		query := `select id from genres where genre_name = $1`
		err := m.DB.QueryRowContext(ctx, query, val).Scan(&genreID)
		if err != nil {
			return movieID, movieGenres, errors.New("invalid genre name")
		}
		if genreID > 0 {
			if _, exists := movieGenres[genreID]; !exists {
				movieGenres[genreID] = val
			}
		}
	}

	if len(movieGenres) <= 0 {
		movieGenres[10] = "Unknown"
	}

	stmt := `update movies set title = $1, description = $2, year = $3, release_date = $4, 
	runtime = $5,
	image = $6,
	updated_at = $7 where id = $8
	RETURNING id`

	err := m.DB.QueryRowContext(ctx, stmt,
		movie.Title,
		movie.Description,
		movie.Year,
		movie.ReleaseDate,
		movie.Runtime,
		movie.Image,
		time.Now(),
		movie.ID,
	).Scan(&movieID)
	if err != nil {
		log.Println(err)
		return movieID, movieGenres, errors.New("invalid movie data! failed to update the movie")
	}

	q := `delete from movies_genres where movie_id = $1`

	_, err = m.DB.ExecContext(ctx, q, movieID)
	if err != nil {
		return movieID, movieGenres, errors.New("something went wrong")
	}

	for key := range movieGenres {
		stmt = `insert into movies_genres ( genre_id, movie_id, created_at, updated_at)
						values($1, $2, $3, $4)`
		_, err = m.DB.ExecContext(ctx, stmt, key, movieID, time.Now(), time.Now())
		if err != nil {
			return movieID, movieGenres, errors.New("failed to save the movie genres")
		}
	}
	return movieID, movieGenres, nil
}

// Get returns one movie and error, if any
func (m *DBModel) Get(id int) (*Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT m.id, m.title, m.description, m.year, m.release_date, m.runtime, m.image, m.created_at, m.updated_at,
    COALESCE(TRUNC(AVG(r.rating)::numeric, 1), 1.0) AS rating
FROM movies m
LEFT JOIN ratings r ON r.movie_id = m.id
WHERE m.id = $1
GROUP BY m.id;
`

	row := m.DB.QueryRowContext(ctx, query, id)

	var movie Movie
	var image sql.NullString

	err := row.Scan(
		&movie.ID,
		&movie.Title,
		&movie.Description,
		&movie.Year,
		&movie.ReleaseDate,
		&movie.Runtime,
		&image,
		&movie.CreatedAt,
		&movie.UpdatedAt,
		&movie.Rating,
	)
	if err != nil {
		return nil, err
	}

	// Check if the Image value is NULL or empty, and if it is, assign a default value
	if !image.Valid || image.String == "" {
		movie.Image = fmt.Sprintf("https://res.cloudinary.com/%s/image/upload/no-thumb.jpg", os.Getenv("CLOUD_NAME"))
	} else {
		movie.Image = fmt.Sprintf("https://res.cloudinary.com/%s/image/upload/%s", os.Getenv("CLOUD_NAME"), image.String)
	}

	// get genres, if any
	genreQuery := `select
	mg.id, mg.movie_id, mg.genre_id, g.genre_name
from
	movies_genres mg
	left join genres g on (g.id = mg.genre_id)
where
	mg.movie_id = $1
`

	genreRows, _ := m.DB.QueryContext(ctx, genreQuery, movie.ID)

	genres := make(map[int]string)
	for genreRows.Next() {
		var mg MovieGenre
		err := genreRows.Scan(
			&mg.ID,
			&mg.MovieID,
			&mg.GenreID,
			&mg.Genre.GenreName,
		)
		if err != nil {
			return nil, err
		}
		genres[mg.GenreID] = mg.Genre.GenreName
	}
	defer genreRows.Close()

	movie.MovieGenre = genres

	// Get comments ordered by recent update
	query = `SELECT
    c.id, c.user_id, c.comment, c.created_at, c.updated_at, u.name
    FROM
    	comments c
    LEFT JOIN users u ON (u.id = c.user_id)
    WHERE
    	c.movie_id = $1
  	ORDER BY c.created_at DESC
    `

	rows, _ := m.DB.QueryContext(ctx, query, id)
	defer rows.Close()

	var comments []Comment
	for rows.Next() {
		var comment Comment
		err := rows.Scan(
			&comment.ID,
			&comment.UserID,
			&comment.Comment,
			&comment.CreatedAt,
			&comment.UpdatedAt,
			&comment.UserName,
		)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	movie.Comments = comments
	movie.TotalComments = len(comments)

	return &movie, nil
}

func (m *DBModel) DeleteMovie(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := "delete from movies where id = $1"

	_, err := m.DB.ExecContext(ctx, stmt, id)
	if err != nil {
		return errors.New("failed to delete the movie")
	}

	return nil
}

// Get returns one movie and error, if any
func (m *DBModel) GetOneMovie(id, userID int) (*Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT m.id, m.title, m.description, m.year, m.release_date, m.runtime, m.image, m.created_at, m.updated_at,
    COALESCE(TRUNC(AVG(r.rating)::numeric, 1), 1.0) AS rating,
		COUNT(DISTINCT f.id) AS favorites_count
FROM movies m
LEFT JOIN ratings r ON r.movie_id = m.id
LEFT JOIN favorites f ON f.movie_id = m.id
WHERE m.id = $1
GROUP BY m.id;
`

	row := m.DB.QueryRowContext(ctx, query, id)

	var movie Movie
	var image sql.NullString

	err := row.Scan(
		&movie.ID,
		&movie.Title,
		&movie.Description,
		&movie.Year,
		&movie.ReleaseDate,
		&movie.Runtime,
		&image,
		&movie.CreatedAt,
		&movie.UpdatedAt,
		&movie.Rating,
		&movie.TotalFavorites,
	)
	if err != nil {
		return nil, err
	}

	// Check if the Image value is NULL or empty, and if it is, assign a default value
	if !image.Valid || image.String == "" {
		movie.Image = fmt.Sprintf("https://res.cloudinary.com/%s/image/upload/no-thumb.jpg", os.Getenv("CLOUD_NAME"))
	} else {
		movie.Image = fmt.Sprintf("https://res.cloudinary.com/%s/image/upload/%s", os.Getenv("CLOUD_NAME"), image.String)
	}

	// get genres, if any
	genreQuery := `select
	mg.id, mg.movie_id, mg.genre_id, g.genre_name
from
	movies_genres mg
	left join genres g on (g.id = mg.genre_id)
where
	mg.movie_id = $1
`

	genreRows, _ := m.DB.QueryContext(ctx, genreQuery, movie.ID)

	genres := make(map[int]string)
	for genreRows.Next() {
		var mg MovieGenre
		err := genreRows.Scan(
			&mg.ID,
			&mg.MovieID,
			&mg.GenreID,
			&mg.Genre.GenreName,
		)
		if err != nil {
			return nil, err
		}
		genres[mg.GenreID] = mg.Genre.GenreName
	}
	defer genreRows.Close()

	movie.MovieGenre = genres

	// Get comments ordered by recent update
	query = `SELECT
    c.id, c.user_id, c.comment, c.created_at, c.updated_at, u.name
    FROM
    	comments c
    LEFT JOIN users u ON (u.id = c.user_id)
    WHERE
    	c.movie_id = $1
  	ORDER BY c.created_at DESC
    `

	rows, _ := m.DB.QueryContext(ctx, query, id)
	defer rows.Close()

	var comments []Comment
	for rows.Next() {
		var comment Comment
		err := rows.Scan(
			&comment.ID,
			&comment.UserID,
			&comment.Comment,
			&comment.CreatedAt,
			&comment.UpdatedAt,
			&comment.UserName,
		)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	movie.Comments = comments
	movie.TotalComments = len(comments)

	if userID > 0 {
		// check if movie is favorite
		favoriteQuery := `select id from favorites where movie_id = $1 and user_id = $2`
		var favID int
		_ = m.DB.QueryRowContext(ctx, favoriteQuery, id, userID).Scan(&favID)

		if favID > 0 {
			movie.IsFavorite = true
		}
	}

	return &movie, nil
}

// CheckComment returns comment_id and error, if any
func (m *DBModel) CheckComment(commentID int) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var id int
	query := `select id from comments where id = $1`

	err := m.DB.QueryRowContext(ctx, query, commentID).Scan(&id)
	if err != nil {
		return id, errors.New("invalid comment id")
	}

	return id, nil
}

// Get Comment returns one comment and error, if any
func (m *DBModel) GetComment(id int) (*Comment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, movie_id, user_id, comment, created_at, updated_at from comments where id = $1`

	row := m.DB.QueryRowContext(ctx, query, id)

	var comment Comment

	err := row.Scan(
		&comment.ID,
		&comment.MovieID,
		&comment.UserID,
		&comment.Comment,
		&comment.CreatedAt,
		&comment.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &comment, nil
}

// InsertComment is help to add a comment to the database
func (m *DBModel) InsertComment(comment *Comment) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var commentID int
	stmt := `insert into comments (movie_id, user_id, comment, created_at, updated_at)
						values($1, $2, $3, $4, $5)
						RETURNING id`

	err := m.DB.QueryRowContext(ctx, stmt,
		comment.MovieID,
		comment.UserID,
		comment.Comment,
		time.Now(),
		time.Now(),
	).Scan(&commentID)
	if err != nil {
		return commentID, errors.New("failed to add the comment")
	}

	return commentID, nil
}

// UpdateComment is help to edit a comment
func (m *DBModel) UpdateComment(comment *Comment) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var commentID int
	stmt := `update comments set comment = $1, updated_at = $2 where id = $3
	RETURNING id`

	err := m.DB.QueryRowContext(ctx, stmt,
		comment.Comment,
		time.Now(),
		comment.ID,
	).Scan(&commentID)
	if err != nil {
		return commentID, errors.New("failed to update the comment")
	}

	return commentID, nil
}

// DeleteComment is help to delete a comment
func (m *DBModel) DeleteComment(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := "delete from comments where id = $1"

	_, err := m.DB.ExecContext(ctx, stmt, id)
	if err != nil {
		return errors.New("failed to delete the comment")
	}

	return nil
}

// FindFavorites is helps to find any favorite is exist base on movie_id and user_id
func (m *DBModel) FindFavorites(userID, movieID int) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id from favorites where movie_id = $1 and user_id = $2`

	favID := 0

	err := m.DB.QueryRowContext(ctx, query, movieID, userID).Scan(&favID)
	if err != nil {
		return 0, errors.New("Favorite does not found")
	}

	return favID, nil
}

// AddFavorite is help to add a favorite movie to the database
func (m *DBModel) AddFavorite(favorite *Favorite) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var favoriteID int
	stmt := `insert into favorites (user_id, movie_id, created_at, updated_at)
						values($1, $2, $3, $4)
						RETURNING id`

	err := m.DB.QueryRowContext(ctx, stmt,
		favorite.UserID,
		favorite.MovieID,
		time.Now(),
		time.Now(),
	).Scan(&favoriteID)
	if err != nil {
		return favoriteID, errors.New("failed to add the favorite")
	}

	return favoriteID, nil
}

// DeleteComment is help to delete a comment
func (m *DBModel) RemoveFavorite(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := "delete from favorites where id = $1"

	_, err := m.DB.ExecContext(ctx, stmt, id)
	if err != nil {
		return errors.New("failed to remove favorite")
	}

	return nil
}

// Insert image info to the database
func (m *DBModel) InsertImageInfo(image *Image) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var imageID int
	stmt := `insert into images (user_id, image_path, image_name, is_used, created_at, updated_at)
						values($1, $2, $3, $4, $5, $6)
						RETURNING id`

	err := m.DB.QueryRowContext(ctx, stmt,
		image.UserID,
		image.ImagePath,
		image.ImageName,
		image.IsUsed,
		time.Now(),
		time.Now(),
	).Scan(&imageID)
	if err != nil {
		return imageID, errors.New("failed to upload the image")
	}

	return imageID, nil
}

// Get Image returns one image name and error, if any
func (m *DBModel) GetImage(id int) (*Image, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, user_id, image_path, image_name, is_used, created_at, updated_at from images where id = $1`

	row := m.DB.QueryRowContext(ctx, query, id)

	var image Image

	err := row.Scan(
		&image.ID,
		&image.UserID,
		&image.ImagePath,
		&image.ImageName,
		&image.IsUsed,
		&image.CreatedAt,
		&image.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &image, nil
}

// get image movieID and error, if any
func (m *DBModel) GetImageByMovieID(movieID int) (*Image, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var imageName string
	query := `select image from movies where id = $1`

	row := m.DB.QueryRowContext(ctx, query, movieID)
	err := row.Scan(&imageName)

	if err != nil {
		return nil, err
	}

	var image Image

	query = `select id, user_id, image_path, image_name, is_used, created_at, updated_at from images where image_name = $1`

	row = m.DB.QueryRowContext(ctx, query, imageName)

	err = row.Scan(
		&image.ID,
		&image.UserID,
		&image.ImagePath,
		&image.ImageName,
		&image.IsUsed,
		&image.CreatedAt,
		&image.UpdatedAt,
	)

	if err != nil {
		return nil, errors.New("failed to get the image")
	}

	return &image, nil
}

// delete image info from the database and file from the server
func (m *DBModel) DeleteImage(image *Image) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// delete image info from the database
	stmt := "delete from images where id = $1"
	_, err := m.DB.ExecContext(ctx, stmt, image.ID)
	if err != nil {
		return errors.New("failed to delete the image from the database")
	}

	return nil
}

// // delete image info from the database and file from the server
// func (m *DBModel) DeleteImage(image *Image) error {
// 	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
// 	defer cancel()

// 	// delete image from the server
// 	err := os.Remove(fmt.Sprintf(".%s/%s", image.ImagePath, image.ImageName))
// 	if err != nil {
// 		return errors.New("failed to delete the image from the server")
// 	}

// 	// delete image info from the database
// 	stmt := "delete from images where id = $1"

// 	_, err = m.DB.ExecContext(ctx, stmt, image.ID)
// 	if err != nil {
// 		return errors.New("failed to delete the image from the database")
// 	}

// 	return nil
// }
