import React from "react";

const TextComponent = (props: {
  content: string;
  style: React.CSSProperties;
}) => {
  return <p style={props.style}> {props.content}</p>;
};

export default TextComponent;
