import React, { useState } from "react";
import { searchMovie, addMovieToList } from "../services/api";
import MovieCard from "../components/MovieCard";
import { Link } from "react-router-dom";
import "./HomePage.css";

const HomePage: React.FC = () => {
  const [movie, setMovie] = useState<{
    id: number;
    original_title?: string;
    original_name?: string;
    poster_path: string;
    overview: string;
  } | null>(null);
  const [searchTerm, setSearchTerm] = useState<string>("");

  const handleSearchChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setSearchTerm(event.target.value);
  };

  const handleSearchSubmit = async (event: React.FormEvent) => {
    event.preventDefault();
    if (searchTerm) {
      try {
        const movieData = await searchMovie(searchTerm);
        setMovie(movieData);
      } catch (error) {
        console.error("Error searching movie:", error);
        setMovie(null); // Clear movie if there's an error
      }
    }
  };

  const handleAddToList = async () => {
    if (movie) {
      try {
        await addMovieToList(
          movie.original_title || movie.original_name || "",
          movie.poster_path,
        );
        alert(
          `${movie.original_title || movie.original_name} added to your list!`,
        );
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
        {movie ? (
          <div>
            <MovieCard
              original_title={movie.original_title}
              original_name={movie.original_name}
              poster_path={movie.poster_path}
              overview={movie.overview}
            />
            <button onClick={handleAddToList}>Add to List</button>
          </div>
        ) : (
          <p>No Movie or TV show found</p>
        )}
      </div>
      <Link to="/watch-list" className="watch-list-link">
        Go to Watch List
      </Link>
    </div>
  );
};

export default HomePage;
