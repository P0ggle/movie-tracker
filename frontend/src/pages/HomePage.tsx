import React, { useState } from "react";
import { searchMovie } from "../services/api";
import MovieCard from "../components/MovieCard";

const HomePage: React.FC = () => {
  const [movie, setMovie] = useState<{
    id: number;
    title?: string;
    original_name?: string;
    poster_path: string;
  } | null>(null);
  const [searchTerm, setSearchTerm] = useState<string>("");
  const [searchType, setSearchType] = useState<string>("movie");

  const handleSearchChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setSearchTerm(event.target.value);
  };

  const handleTypeChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setSearchType(event.target.value);
  };

  const handleSearchSubmit = async (event: React.FormEvent) => {
    event.preventDefault();
    if (searchTerm) {
      try {
        const movieData = await searchMovie(searchTerm, searchType);
        setMovie(movieData);
      } catch (error) {
        console.error("Error searching movie:", error);
        setMovie(null); // Clear movie if there's an error
      }
    }
  };

  return (
    <div>
      <form onSubmit={handleSearchSubmit}>
        <input
          type="text"
          value={searchTerm}
          onChange={handleSearchChange}
          placeholder="Search for a movie or TV show..."
        />
        <div>
          <label>
            <input
              type="radio"
              value="movie"
              checked={searchType === "movie"}
              onChange={handleTypeChange}
            />
            Movie
          </label>
          <label>
            <input
              type="radio"
              value="tv"
              checked={searchType === "tv"}
              onChange={handleTypeChange}
            />
            TV Show
          </label>
        </div>
        <button type="submit">Search</button>
      </form>
      <div>
        {movie ? (
          <MovieCard
            key={movie.id}
            title={movie.title}
            original_name={movie.original_name}
            poster_path={movie.poster_path}
          />
        ) : (
          <p>No movie or TV show found</p>
        )}
      </div>
    </div>
  );
};

export default HomePage;


