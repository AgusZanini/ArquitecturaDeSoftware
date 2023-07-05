import React from "react";
import { Page } from "./Componentes/Page";
import { Header } from "./Componentes/header";
import {BrowserRouter as Router} from "react-router-dom";

function App() {
  return (
    <div className="App">
      <Router>
        <Header/>
        <Page/>
      </Router>
    </div>
  );
}

export default App;
