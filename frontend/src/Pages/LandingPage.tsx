import React from "react";
import ImagePlaceHolder from "../Components/ImagePlaceHolder.tsx";
import TextComponent from "../Components/TextComponent.tsx";
const LandingPage = () => {
  return (
    <>
      <div style={{ display: "flex", flex: 0.5 }}>
        <TextComponent
          style={{
            marginTop: "35vh",
            marginLeft: "10vh",
            fontSize: "5vh",
            textAlign: "left",
            fontFamily: "sans-serif",
          }}
          content={"TRANSFORMING SALES FORECASTING WITH ADVANCED ANALYTICS"}
        />
      </div>
      <div style={{ display: "flex", flex: 0.5 }}>
        <ImagePlaceHolder
          style={{
            height: "50vh",
            width: "40vh",
            borderRadius: "2vh",
            marginTop: "25vh",
            marginLeft: "30vh",
            marginRight:"10vh"
          }}
          src={
            "https://logopond.com/logos/550d2e4cdf54206d8765f5e447e2b368.png"
          }
          alt={""}
        ></ImagePlaceHolder>
      </div>
    </>
  );
};

export default LandingPage;
