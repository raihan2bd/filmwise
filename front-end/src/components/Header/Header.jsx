import { useState } from 'react';
import { Link, NavLink } from 'react-router-dom';
import { FiMenu, FiX, FiSearch } from 'react-icons/fi'

import './Header.css'

const Header = () => {
  const [showNav, setShowNav] = useState(false)
  const [showSearch, setShowSearch] = useState(false)

  const toggleNavHandler = () => {
    setShowNav((prevState) => !prevState);
  }
  
  const showSearchHandler = () => {
    setShowSearch(true);
  }

  const hideSearchHandler = () => {
    setShowSearch(false);
  }

  return (
    <header className='header bg-dark color-white d-flex jc-space-between align-items-center'>
      <div className="brand d-flex align-items-center">
        <button className='mob-menu sm-content' onClick={toggleNavHandler}>
          {!showNav ? <FiMenu /> : <FiX />}
        </button>
        <Link className="brand-name">FilmWise</Link>
      </div>
      <nav className="nav-bar d-flex gap-1">
        <ul className={showNav? "nav-group d-flex-sm-column d-flex-md-row gap-1 list-style-none": "nav-group d-flex-md-row gap-1 list-style-none"}>
          <li>
            <NavLink className="nav-link" to="/">Home</NavLink>
          </li>
          <li>
            <NavLink className="nav-link" to="/movies">Movies</NavLink>
          </li>
          <li>
            <NavLink className="nav-link" to="/series">Series</NavLink>
          </li>
          <li>
            <NavLink className="nav-link" to="/login">Login</NavLink>
          </li>

        </ul>

          <div className="sm-content">
            <button className='btn-search-toggle' onClick={showSearchHandler}>
              <FiSearch />
            </button>
            {showSearch && (<div className='sm-search-form d-flex jc-space-between'>
              <button onClick={hideSearchHandler} className='btn-search-toggle color-orange'>X</button>
              <form action="#" className='search-item d-flex'>
                <input className="search-input" type="search" placeholder="Search movies..." />
                <button className='btn-search-toggle' type="submit"><FiSearch /></button>
              </form>
            </div>)}
            
            </div>
          <form action="#" className='search-item d-flex md-content'>
            <input className="search-input" type="search" placeholder="Search movies..." />
            <button type="submit" className='btn-search-toggle'><FiSearch /></button>
          </form>
      </nav>
    </header>
  );
};

export default Header;
