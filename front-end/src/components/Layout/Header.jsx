import { Link, NavLink } from 'react-router-dom';

const Header = () => {
  return (
    <header className='header'>
      <div className="brand">
        <Link className="brand-name">FilmWise</Link>
      </div>
      <nav className="nav-bar">
        <ul className="nav-group">
          <li>
            <NavLink className="nav-link" to="#">Home</NavLink>
          </li>
          <li>
            <NavLink className="nav-link" to="/movies">Movies</NavLink>
          </li>
          <li>
            <NavLink className="nav-link" to="/series">Series</NavLink>
          </li>
          <li>
            <NavLink className="nav-link" to="/about">About</NavLink>
          </li>
          <li>
            <NavLink className="nav-link" to="/contact">Contact</NavLink>
          </li>
          <li>
            <NavLink className="nav-link" to="/login">Login</NavLink>
          </li>
        </ul>
        <form action="#" className='search-item'>
          <input type="text" placeholder="Search movies..." />
          <button type="submit">Search</button>
        </form>
      </nav>
    </header>
  );
};

export default Header;
