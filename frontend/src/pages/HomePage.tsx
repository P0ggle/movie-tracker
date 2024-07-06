import React, { useState, useEffect } from "react";
import {
  searchMovies,
  addMovieToList,
  logout,
  getToken,
  getUser,
} from "../services/api";
import MovieCard from "../components/MovieCard";
import { Link, useNavigate } from "react-router-dom";
import "./HomePage.css";
import "./Popup.css"; // Import Popup styles
import LogoutConfirmationPopup from "../components/LogoutConfirmationPopUp";

const HomePage: React.FC = () => {
  const [movies, setMovies] = useState<Array<{
    id: number;
    original_title?: string;
    original_name?: string;
    poster_path: string;
    overview: string;
  }> | null>(null);
  const [searchTerm, setSearchTerm] = useState<string>("");
  const [selectedMovie, setSelectedMovie] = useState<{
    id: number;
    original_title?: string;
    original_name?: string;
    poster_path: string;
    overview: string;
  } | null>(null);
  const [added, setAdded] = useState<boolean>(false);
  const [isLoggedIn, setIsLoggedIn] = useState<boolean>(false);
  const [username, setUsername] = useState<string | null>(null);
  const [hasSearched, setHasSearched] = useState<boolean>(false);
  const [showLogoutPopup, setShowLogoutPopup] = useState<boolean>(false);

  const navigate = useNavigate();

  useEffect(() => {
    const token = getToken();
    const user = getUser();
    if (token && user) {
      setIsLoggedIn(true);
      setUsername(user);
    }
  }, []);

  const handleSearchChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setSearchTerm(event.target.value);
  };

  const handleSearchSubmit = async (event: React.FormEvent) => {
    event.preventDefault();
    if (searchTerm) {
      try {
        const movieData = await searchMovies(searchTerm);
        setMovies(movieData);
      } catch (error) {
        console.error("Error searching movies:", error);
        setMovies(null); // Clear movies if there's an error
      }
      setHasSearched(true);
    }
  };

  const handleCardClick = (movie: any) => {
    setSelectedMovie(movie);
    setAdded(false); // Reset the added state
  };

  const handleAddToList = async () => {
    if (selectedMovie) {
      try {
        await addMovieToList(
          selectedMovie.original_title || selectedMovie.original_name || "",
          selectedMovie.poster_path,
        );
        setAdded(true); // Update the state to reflect the addition
      } catch (error) {
        console.error("Error adding movie to list:", error);
        alert("Failed to add movie to list.");
      }
    }
  };

  const handleLogout = () => {
    setShowLogoutPopup(true);
  };

  const confirmLogout = () => {
    logout();
    setIsLoggedIn(false);
    setUsername(null);
    setShowLogoutPopup(false);
  };

  const cancelLogout = () => {
    setShowLogoutPopup(false);
  };

  return (
    <div className="homepage">
      <div className="auth-buttons">
        {isLoggedIn ? (
          <>
            <button className="username-button">
              <i className="fas fa-user"></i> User: {username}!
            </button>
            <button className="button-style" onClick={handleLogout}>
              Logout
            </button>
          </>
        ) : (
          <>
            <button className="button-style" onClick={() => navigate("/login")}>
              Login
            </button>
            <button
              className="button-style"
              onClick={() => navigate("/signup")}
            >
              Register
            </button>
          </>
        )}
      </div>
      <h1>Movie Search</h1>
      <form className="search-form" onSubmit={handleSearchSubmit}>
        <input
          type="text"
          value={searchTerm}
          onChange={handleSearchChange}
          placeholder="Search for a Movie or TV show..."
        />
        <button type="submit">Search</button>
      </form>
      <div className="movie-results">
        {movies
          ? movies.map((movie) => (
            <MovieCard
              key={movie.id}
              original_title={movie.original_title}
              original_name={movie.original_name}
              poster_path={movie.poster_path}
              overview={movie.overview}
              onClick={() => handleCardClick(movie)}
            />
          ))
          : hasSearched && <p>No Movie or TV show found</p>}
      </div>
      <Link to="/watch-list" className="button-style watch-list-link">
        Go to Watch List
      </Link>

      {selectedMovie && (
        <div className="popup">
          <div className="popup-content">
            <h3>
              {selectedMovie.original_title || selectedMovie.original_name}
            </h3>
            <img
              src={`https://image.tmdb.org/t/p/w500${selectedMovie.poster_path}`}
              alt={selectedMovie.original_title || selectedMovie.original_name}
            />
            <p>{selectedMovie.overview}</p>
            <button onClick={handleAddToList} disabled={added}>
              {added ? "Added!" : "Add to List"}
            </button>
            <button onClick={() => setSelectedMovie(null)}>Close</button>
          </div>
        </div>
      )}

      {showLogoutPopup && (
        <LogoutConfirmationPopup
          username={username}
          onConfirm={confirmLogout}
          onCancel={cancelLogout}
        />
      )}
    </div>
  );
};

export default HomePage;
