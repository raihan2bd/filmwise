import classes from './Movie.module.css'

const Movie = () => (
  <li className={classes.movie}>
    <div className={classes.movie_img}>
      <img src="" alt="movie"></img>
    </div>
    <a href="#" className={classes.movie_title}></a>
    <div className={classes.movie_short_info}>
      <button className={classes.favorite}>Fav</button>
      <div className={classes.rating}>
        <span>8.5</span>
      </div>
    </div>
  </li>
);

export default Movie;
