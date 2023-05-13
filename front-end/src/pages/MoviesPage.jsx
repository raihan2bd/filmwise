import classes from './MoviesPage.module.css'

const MoviesPage = () => (
  <section>
    <div className={classes.movies}>
    <div className={classes.movies_header}>
      <h2>Movies</h2>
      <button button>Filter</button>
    </div>
      <ul className={classes.movies_list}>
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
      </ul>
    </div>
  </section>
);

export default MoviesPage;