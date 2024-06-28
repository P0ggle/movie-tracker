import React from "react";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import HomePage from "./pages/HomePage";
import WatchListPage from "./pages/WatchListPage";

const App: React.FC = () => {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<HomePage />} />
        <Route path="/watch-list" element={<WatchListPage />} />
      </Routes>
    </Router>
  );
};

export default App;
