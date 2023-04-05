import {Routes, Route} from 'react-router-dom';

import Layout from './components/Layout/Layout';
import HomePage from './pages/Homepage';
import AuthPage from './pages/AuthPage';
import MoviesPage from './pages/MoviesPage';
import SeriesPage from './pages/SeriesPage';
import MediaPage from './pages/MediaPage';
import PageNotFound from './pages/PageNotFound';
import SingleMediaPage from './pages/SingleMediaPage';

const App = () => {
  return (
  <Layout>
    <Routes>
      <Route path="/" element={<HomePage />} />
      <Route path="/auth" element={<AuthPage />} />
      <Route path="/movies" element={<MoviesPage />} />
      <Route path="/series" element={<SeriesPage />} />
      <Route path="/media/all" element={<MediaPage />} />
      <Route path="/media/title" element={<SingleMediaPage />} />
      <Route path="*" element={<PageNotFound />} />
    </Routes>
  </Layout>
  )
}

export default App
