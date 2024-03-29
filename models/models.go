package models

import (
	"database/sql"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
)

type DBModel struct {
	DB *sql.DB
}

// Models is the wrapper for database
type Models struct {
	DB  DBModel
	CLD *cloudinary.Cloudinary
}

// NewModels returns models with db pool
func NewModels(db *sql.DB, cld *cloudinary.Cloudinary) Models {
	return Models{
		DB:  DBModel{DB: db},
		CLD: cld,
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
	Ratings        []Rating       `json:"ratings,omitempty"` // this is for movie details
	TotalFavorites int            `json:"total_favorites"`   // this is for movie details
	IsFavorite     bool           `json:"is_favorite"`
	Favorites      []Favorite     `json:"favorites,omitempty"`
	TotalComments  int            `json:"total_comments"`
	Comments       []Comment      `json:"comments,omitempty"` // this is for movie details
	MovieGenre     map[int]string `json:"genres"`             // this is for movie details
	Image          string         `json:"image"`
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
	UserName  string    `json:"user_name"`
	MovieID   int       `json:"movie_id,omitempty"`
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

// model for Image
type Image struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	ImagePath string    `json:"image_path"`
	ImageName string    `json:"image_name"`
	IsUsed    bool      `json:"is_used"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

// model for User
type User struct {
	ID        int       `json:"id"`
	FullName  string    `json:"full_name,omitempty"`
	Email     string    `json:"email"`
	UserType  string    `json:"user_type"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

// Model for movies response
type PaginatedMovies struct {
	TotalCount  int      `json:"total_count"`
	PerPage     int      `json:"per_page"`
	CurrentPage int      `json:"current_page"`
	Movies      []*Movie `json:"movies"`
}
