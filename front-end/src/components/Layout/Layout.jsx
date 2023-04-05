import Header from '../Header/Header';

const Layout = ({ children }) => {
  return (
    <>
      <Header />
      <main className="main-container">{children}</main>
      <footer className="footer">
        <p>&copy; 2023 FilmWise, Inc. All Rights Reserved.</p>
      </footer>
    </>
  );
};

export default Layout;
