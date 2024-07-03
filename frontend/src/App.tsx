import React from "react";
import { BrowserRouter as Router, Route, Routes, Navigate } from "react-router-dom";
import HomePage from "./pages/HomePage";
import WatchListPage from "./pages/WatchListPage";
import LoginPage from "./pages/LoginPage";
import SignupPage from "./pages/SignupPage";
import { getToken } from "./services/api"; 
import "./global.css";

// checks whether there is a token available in the local storage. If there is, it renders the children, otherwise, it redirects to the login page.
const PrivateRoute = ({ children }: { children: JSX.Element }) => {
  const token = getToken();
  return token ? children : <Navigate to="/login" />;
};

const App: React.FC = () => {
  return (
    <Router>
      <Routes>
        <Route path="/login" element={<LoginPage />} />
        <Route path="/signup" element={<SignupPage />} />
        <Route
          path="/"
          element={
              <HomePage />
          }
        />
        <Route
          path="/watch-list"
          element={
            <PrivateRoute>
              <WatchListPage />
            </PrivateRoute>
          }
        />
      </Routes>
    </Router>
  );
};

export default App;
