import Movie from "./Movie";

import classes from './Movies.module.css';

const Movies = ({movies}) => {

  let moviesContent = <p>No Movie Found!</p>;

  if (movies.length > 0) {
    moviesContent = movies.map(movie => {
      return (<Movie key={movie.id} movie={movie} />)
    });
  }

  return (
    <ul className={classes.movies_list}>
      {moviesContent}
    </ul>
  );
};

export default Movies;
