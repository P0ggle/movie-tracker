import React, { useState } from "react";
import { searchMovie } from "../services/api";
import MovieCard from "../components/MovieCard";

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
        setMovie(null);
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
          placeholder="Search for a Movie or TV show..."
        />
        <button type="submit">Search</button>
      </form>
      <div>
        {movie ? (
          <MovieCard
            original_title={movie.original_title}
            original_name={movie.original_name}
            poster_path={movie.poster_path}
            overview={movie.overview}
          />
        ) : (
          <p>No Movie or TV show found</p>
        )}
      </div>
    </div>
  );
};

export default HomePage;
