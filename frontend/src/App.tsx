import React from "react";
import "./App.css";
import LandingPage from "./Pages/LandingPage.tsx";

function App() {
  return (
    <>
      <div style={{ backgroundColor: "#f6e5d6", display: "flex", flex: 1 }}>
        {" "}
        <LandingPage />
      </div>
    </>
  );
}

export default App;
