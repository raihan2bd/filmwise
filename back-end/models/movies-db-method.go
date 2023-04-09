package models

import (
	"context"
	"fmt"
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

// func (m *DBModel) GetAllMoviesByFilter(page, perPage int, findByName, filterByGenre, filterByYear, orderBy string) ([]*Movie, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
// 	defer cancel()

// 	offset := (page - 1) * perPage

// 	where := fmt.Sprintf("where %s %s %s", findByName, filterByGenre, filterByYear)

// 	query := fmt.Sprintf("select id, title, description, year, release_date, rating, runtime, created_at, updated_at from movies %s %s limit $1 offset $2", where, orderBy)

// 	rows, err := m.DB.QueryContext(ctx, query, perPage, offset)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var movies []*Movie
// 	for rows.Next() {
// 		var movie Movie
// 		err = rows.Scan(
// 			&movie.ID,
// 			&movie.Title,
// 			&movie.Description,
// 			&movie.Year,
// 			&movie.ReleaseDate,
// 			&movie.Rating,
// 			&movie.Runtime,
// 			&movie.CreatedAt,
// 			&movie.UpdatedAt,
// 		)

// 		if err != nil {
// 			return nil, err
// 		}

// 		// get genres, if any
// 		genreQuery := `select
// 			mg.id, mg.movie_id, mg.genre_id, g.genre_name
// 		from
// 			movies_genres mg
// 			left join genres g on (g.id = mg.genre_id)
// 		where
// 			mg.movie_id = $1
// 		`

// 		genreRows, _ := m.DB.QueryContext(ctx, genreQuery, movie.ID)

// 		genres := make(map[int]string)
// 		for genreRows.Next() {
// 			var mg MovieGenre
// 			err := genreRows.Scan(
// 				&mg.ID,
// 				&mg.MovieID,
// 				&mg.GenreID,
// 				&mg.Genre.GenreName,
// 			)
// 			if err != nil {
// 				return nil, err
// 			}
// 			genres[mg.GenreID] = mg.Genre.GenreName
// 		}
// 		genreRows.Close()

// 		movie.MovieGenre = genres
// 		movies = append(movies, &movie)
// 	}

// 	return movies, nil
// }

func (m *DBModel) GetAllMoviesByFilter(page, perPage int, filter *MovieFilter) ([]*Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// offset := (page - 1) * perPage
	// fmt.Println(offset, perPage)

	var dbArgs []interface{}
	where := " WHERE (title LIKE $1 OR description LIKE $2)"
	dbArgs = append(dbArgs, "%"+filter.FindByName+"%", "%"+filter.FindByName+"%")
	// if filter.FilterByYear > 0 {
	// 	where += " and year = $3"
	// 	dbArgs = append(dbArgs, filter.FilterByYear)
	// }
	// if filter.FilterByGenre > 0 && filter.FilterByYear <= 0 {
	// 	where += " and m.id in (select movie_id from movies_genres where genre_id = $4)"
	// 	dbArgs = append(dbArgs, filter.FilterByGenre)
	// }

	// if (filter.FilterByGenre > 0) && (filter.FilterByYear > 0) {
	// 	where += " and m.id in (select movie_id from movies_genres where genre_id = $5)"
	// 	dbArgs = append(dbArgs, filter.FilterByGenre)
	// }

	// add order by query
	// orderByQuery := ""
	// switch filter.OrderBy {
	// case "rating":
	// 	orderByQuery = " order by rating desc"
	// case "runtime":
	// 	orderByQuery = " order by runtime desc"
	// case "old":
	// 	orderByQuery = " order by release_date asc"
	// case "name":
	// 	orderByQuery = " order by title asc"
	// default:
	// 	orderByQuery = " order by release_date desc"
	// }

	limitQuery := " limit $3"
	dbArgs = append(dbArgs, perPage)
	// offsetQuery := fmt.Sprintf(" OFFSET %d", offset)
	// dbArgs = append(dbArgs, perPage, offset)

	// paginationQuery := " limit $3 offset $4 ;"

	query := `select m.id, m.title, m.description, m.year, m.release_date, m.rating, m.runtime, m.created_at, m.updated_at,
	mg.id, mg.movie_id, mg.genre_id, g.genre_name
	from movies m
	left join movies_genres mg on (m.id = mg.movie_id)
	left join genres g on (g.id = mg.genre_id)`
	query += where + limitQuery + ";"

	rows, err := m.DB.QueryContext(ctx, query, dbArgs...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []*Movie

	for rows.Next() {
		var movie Movie
		var mg MovieGenre

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

			&mg.ID,
			&mg.MovieID,
			&mg.GenreID,
			&mg.Genre.GenreName,
		)
		if err != nil {
			return nil, err
		}

		// check if the movie already exists in the slice
		found := false
		for i := range movies {
			if movies[i].ID == movie.ID {
				// add the genre to the existing movie
				movies[i].MovieGenre[mg.GenreID] = mg.Genre.GenreName
				found = true
				break
			}
		}

		if !found {
			// create a new map for genres
			movie.MovieGenre = make(map[int]string)
			// add the genre to the new movie
			movie.MovieGenre[mg.GenreID] = mg.Genre.GenreName
			// append the new movie to the slice
			movies = append(movies, &movie)
		}
	}

	return movies, nil
}
