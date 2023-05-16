import Movies from '../components/Movies/Movies';
import classes from './MoviesPage.module.css'

const dummyMovies = [
    {
      id: 3,
      title: "The Dark Knight",
      description: "The menace known as the Joker wreaks havoc on Gotham City.",
      year: 2008,
      release_date: "2008-07-18T00:00:00Z",
      runtime: 152,
      rating: 1,
      total_favorites: 0,
      is_favorite: false,
      total_comments: 0,
      genres: {
        1: "Drama",
        2: "Crime",
        3: "Action",
        6: "Mystery"
      },
      image: ""
    },
    {
      id: 2,
      title: "The Pursuit of Happyness",
      description: "Based on a true story about a man named Christopher Gardner.",
      year: 2006,
      release_date: "2006-12-15T00:00:00Z",
      runtime: 117,
      rating: 1,
      total_favorites: 0,
      is_favorite: false,
      total_comments: 0,
      genres: {
        1: "Drama"
      },
      image: ""
    },
    {
      id: 1,
      title: "The Shawshank Redemption",
      description: "Two imprisoned men bond over a number of years",
      year: 1994,
      release_date: "1994-10-14T00:00:00Z",
      runtime: 142,
      rating: 1,
      total_favorites: 0,
      is_favorite: false,
      total_comments: 0,
      genres: {
        1: "Drama",
        2: "Crime",
        6: "Mystery"
      },
      image: ""
    }
  ];

const MoviesPage = () => (
  <section>
    <div className={classes.movies}>
    <div className={classes.movies_header}>
      <h2>Movies</h2>
      <button button>Filter</button>
    </div>
      <Movies movies={dummyMovies} />
    </div>
  </section>
);

export default MoviesPage;