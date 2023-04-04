import ImageSlider from '../components/ImageSlider/ImageSlider';

import MovieImg from '../assets/images/forrest_gump.jpg';

const HomePage = () => {
  const images = [
    {
      id: 'slider1',
      src: MovieImg,
      alt: 'Image 1',
      title: 'Forrest Gump 1',
      description:
        'Lorem ipsum dolor sit amet consectetur adipisicing elit. Maiores repudiandae, soluta expedita repellat laborum dolore recusandae illum, itaque eveniet laboriosam quas ad impedit aliquid, ab dolores reprehenderit quod perspiciatis. Debitis?',
    },
    {
      id: 'slider2',
      src: MovieImg,
      alt: 'Image 2',
      title: 'Forrest Gump 2',
      description:
        'Lorem ipsum dolor sit amet consectetur adipisicing elit. Maiores repudiandae, soluta expedita repellat laborum dolore recusandae illum, itaque eveniet laboriosam quas ad impedit aliquid, ab dolores reprehenderit quod perspiciatis. Debitis?',
    },
    {
      id: 'slider3',
      src: MovieImg,
      alt: 'Image 3',
      title: 'Forrest Gump 3',
      description:
        'Lorem ipsum dolor sit amet consectetur adipisicing elit. Maiores repudiandae, soluta expedita repellat laborum dolore recusandae illum, itaque eveniet laboriosam quas ad impedit aliquid, ab dolores reprehenderit quod perspiciatis. Debitis?',
    },
  ];

  return (
    <>
      <h2>Hello From Home Page!</h2>
      <div>
        <ImageSlider images={images} />
      </div>
    </>
  );
};

export default HomePage;
