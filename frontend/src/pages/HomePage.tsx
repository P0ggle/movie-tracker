import React, { useState } from "react";
import { searchMovie } from "../services/api";
import MovieCard from "../components/MovieCard";

const HomePage: React.FC = () => {
  const [movie, setMovie] = useState<{
    id: number;
    title: string;
    poster_path: string;
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

  return (
    <div>
      <form onSubmit={handleSearchSubmit}>
        <input
          type="text"
          value={searchTerm}
          onChange={handleSearchChange}
          placeholder="Search for a movie..."
        />
        <button type="submit">Search</button>
      </form>
      <div>
        {movie ? (
          <MovieCard
            key={movie.id}
            title={movie.title}
            poster_path={movie.poster_path}
          />
        ) : (
          <p>No movie found</p>
        )}
      </div>
    </div>
  );
};

export default HomePage;
