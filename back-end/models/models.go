package models

import (
	"database/sql"
	"time"
)

type DBModel struct {
	DB *sql.DB
}

// Models is the wrapper for database
type Models struct {
	DB DBModel
}

// NewModels returns models with db pool
func NewModels(db *sql.DB) Models {
	return Models{
		DB: DBModel{DB: db},
	}
}

// Movie is the type for movies
type Movie struct {
	ID             int            `json:"id"`
	Title          string         `json:"title"`
	Description    string         `json:"description"`
	Year           int            `json:"year"`
	ReleaseDate    time.Time      `json:"release_date"`
	Runtime        int            `json:"runtime"`
	Rating         float64        `json:"rating"`
	Ratings        []Rating       `json:"ratings"`         // this is for movie details
	TotalFavorites int            `json:"total_favorites"` // this is for movie details
	Favorites      []Favorite     `json:"favorites"`
	TotalComments  int            `json:"total_comments"`
	Comments       []Comment      `json:"comments"` // this is for movie details
	MovieGenre     map[int]string `json:"genres"`   // this is for movie details
	CreatedAt      time.Time      `json:"-"`
	UpdatedAt      time.Time      `json:"-"`
}

// Genre is the type for genre
type Genre struct {
	ID        int       `json:"id"`
	GenreName string    `json:"genre_name"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

// MovieGenre is the type for movie genre
type MovieGenre struct {
	ID        int       `json:"-"`
	MovieID   int       `json:"-"`
	GenreID   int       `json:"-"`
	Genre     Genre     `json:"genre"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

// MovieFilter will help to organize query
type MovieFilter struct {
	FindByName    string
	FilterByGenre int
	FilterByYear  int
	OrderBy       string
}

// query params helps to organize query parameters
// struct for query parameters
type QueryParam struct {
	Key   string
	Value interface{}
}

// model for rating
type Rating struct {
	ID        int       `json:"id"`
	MovieID   int       `json:"movie_id"`
	UserID    int       `json:"user_id"`
	Rating    float32   `json:"rating"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

// model for comment
type Comment struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	MovieID   int       `json:"movie_id"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"commented_at"`
}

// model for favorite
type Favorite struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	MovieID   int       `json:"movie_id"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"fav_at"`
}
