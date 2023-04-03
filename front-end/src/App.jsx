import {Routes, Route} from 'react-router-dom';

import Layout from './components/Layout/Layout';
import HomePage from './pages/Homepage';
import AboutPage from './pages/AboutPage';
import ContactPage from './pages/ContactPage';
import PageNotFound from './pages/PageNotFound';

const App = () => {
  return (
  <Layout>
    <Routes>
      <Route path="/" element={<HomePage />} />
      <Route path="/about" element={<AboutPage />} />
      <Route path="/contact" element={<ContactPage />} />
      <Route path="*" element={<PageNotFound />} />
    </Routes>
  </Layout>
  )
}

export default App
