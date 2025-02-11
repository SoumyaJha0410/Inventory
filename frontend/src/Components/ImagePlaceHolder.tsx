import React from "react";

const ImagePlaceHolder = (props: {
  src: string;
  alt: string;
  style?: React.CSSProperties;
}) => {
  const { src, alt, style } = props;
  return (
    <img src={src} alt={alt} style={style}>
    </img>
  );
};

export default ImagePlaceHolder;
