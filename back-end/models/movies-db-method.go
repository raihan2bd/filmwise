package models

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"
)

func (m *DBModel) GetAllMovies(genre ...int) ([]*Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	where := ""
	if len(genre) > 0 {
		where = fmt.Sprintf("where id in (select movie_id from movies_genres where genre_id = %d)", genre[0])
	}

	query := fmt.Sprintf("select id, title, description, year, release_date, rating, runtime, created_at, updated_at from movies  %s order by title", where)

	rows, err := m.DB.QueryContext(ctx, query)
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

		movie.MovieGenre = genres
		movies = append(movies, &movie)
	}

	return movies, nil
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

// Get all movies by filter
func (m *DBModel) GetAllMoviesByFilter(page, perPage int, filter *MovieFilter) ([]*Movie, error) {
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
		where += " and id in (select movie_id from movies_genres where genre_id = $3)"
		dbArgs = append(dbArgs, filter.FilterByGenre)
	}

	if (filter.FilterByGenre > 0) && (filter.FilterByYear > 0) {
		where += " and id in (select movie_id from movies_genres where genre_id = $4)"
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
		orderByQuery = " order by release_date asc"
	case "name":
		orderByQuery = " order by title asc"
	default:
		orderByQuery = " order by release_date desc"
	}

	// main query
	query := `select id, title, description, year, release_date, rating, runtime, created_at, updated_at from movies`

	// pagination query
	paginationQuery := fmt.Sprintf(" limit %d offset %d", perPage, offset)

	// join all the query
	query += where + orderByQuery + paginationQuery

	// execute query with context
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

		movie.MovieGenre = genres
		movies = append(movies, &movie)
	}

	return movies, nil
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
	fmt.Println(movieID)
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
		fmt.Println(genreID)
		if genreID > 0 {
			if _, exists := movieGenres[genreID]; !exists {
				movieGenres[genreID] = val
			}
		}
	}

	if len(movieGenres) <= 0 {
		movieGenres[10] = "Unknown"
	}

	stmt := `insert into movies (title, description, year, release_date, runtime, rating,
		created_at, updated_at) values ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id`

	err := m.DB.QueryRowContext(ctx, stmt,
		movie.Title,
		movie.Description,
		movie.Year,
		movie.ReleaseDate,
		movie.Runtime,
		movie.Rating,
		time.Now(),
		time.Now(),
	).Scan(&movieID)
	if err != nil {
		log.Println(err)
		return movieID, movieGenres, errors.New("invalid movie data! failed to save the movie")
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

	// // return if movie title is already exist
	// q := `select id from movies where title = $1`
	// _ = m.DB.QueryRowContext(ctx, q, movie.Title).Scan(&movieID)
	// fmt.Println(movieID)
	// if movieID > 0 {
	// 	return movieID, movieGenres, errors.New("the movie is already exist")
	// }

	// verify genres
	for _, val := range movie.MovieGenre {
		genreID := 0
		query := `select id from genres where genre_name = $1`
		err := m.DB.QueryRowContext(ctx, query, val).Scan(&genreID)
		if err != nil {
			return movieID, movieGenres, errors.New("invalid genre name")
		}
		fmt.Println(genreID)
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
	runtime = $5, rating = $6,
	updated_at = $7 where id = $8
	RETURNING id`

	err := m.DB.QueryRowContext(ctx, stmt,
		movie.Title,
		movie.Description,
		movie.Year,
		movie.ReleaseDate,
		movie.Runtime,
		movie.Rating,
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

	query := `select id, title, description, year, release_date, rating, runtime,
				created_at, updated_at from movies where id = $1
	`

	row := m.DB.QueryRowContext(ctx, query, id)

	var movie Movie

	err := row.Scan(
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

	// get genres, if any
	query = `select
				mg.id, mg.movie_id, mg.genre_id, g.genre_name
			from
				movies_genres mg
				left join genres g on (g.id = mg.genre_id)
			where
				mg.movie_id = $1
	`

	rows, _ := m.DB.QueryContext(ctx, query, id)
	defer rows.Close()

	genres := make(map[int]string)
	for rows.Next() {
		var mg MovieGenre
		err := rows.Scan(
			&mg.ID,
			&mg.MovieID,
			&mg.GenreID,
			&mg.Genre.GenreName,
		)
		if err != nil {
			return nil, err
		}
		genres[mg.ID] = mg.Genre.GenreName
	}

	movie.MovieGenre = genres

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
