import { Link } from 'react-router-dom';

const PageNotFound = () => (
  <section>
    <h2>404 Page Not Found!</h2>
    <p>Page is not found in our server please Go Back to <Link to="/">Home</Link></p>
  </section>
);

export default PageNotFound;