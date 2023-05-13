import React from 'react';
import AliceCarousel from 'react-alice-carousel';

import Forrest_gump_img from '../../assets/images/forrest_gump.jpg';

import classes from './ImageSlider.module.css';
import 'react-alice-carousel/lib/alice-carousel.css';

const ImageSlider = () => {
  const items = [
    <div className={classes.slider_item} data-value="1">
      <div className={classes.slider_image}>
        <img src={Forrest_gump_img} alt="Movie1" />
      </div>
      <div className={classes.slider_title}>
        <a className={classes.slider_title_link} href="#">Forrest Gump (1994)</a>
      </div>
    </div>,
    <div className={classes.slider_item} data-value="2">
      <div className={classes.slider_image}>
        <img src={Forrest_gump_img} alt="Movie2" />
      </div>
      <div className={classes.slider_title}>
      <a className={classes.slider_title_link} href="#">Forrest Gump (1994)</a>
      </div>
    </div>,
    <div className={classes.slider_item} data-value="3">
      <div className={classes.slider_image}>
        <img src={Forrest_gump_img} alt="Movie3" />
      </div>
      <div className={classes.slider_title}>
      <a className={classes.slider_title_link} href="#">Forrest Gump (1994)</a>
      </div>
    </div>,
  ];
  return (
    <AliceCarousel
      items={items}
      autoPlay
      disableDotsControls
      autoPlayStrategy="none"
      autoPlayInterval={2000}
      animationDuration={2000}
      animationType="fadeout"
      infinite
      disableButtonsControls
    />
  );
};

export default ImageSlider;
