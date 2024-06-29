import React, { useState } from "react";
import { searchMovies, addMovieToList } from "../services/api";
import MovieCard from "../components/MovieCard";
import { Link } from "react-router-dom";
import "./HomePage.css";
import "./Popup.css"; // Import Popup styles

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

  return (
    <div className="homepage">
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
        {movies ? (
          movies.map((movie) => (
            <MovieCard
              key={movie.id}
              original_title={movie.original_title}
              original_name={movie.original_name}
              poster_path={movie.poster_path}
              overview={movie.overview}
              onClick={() => handleCardClick(movie)}
            />
          ))
        ) : (
          <p>No Movie or TV show found</p>
        )}
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
    </div>
  );
};

export default HomePage;
