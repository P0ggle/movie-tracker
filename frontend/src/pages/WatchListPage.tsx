import React, { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import { getMoviesToWatch } from "../services/api";
import MovieCard from "../components/MovieCard";
import "./WatchListPage.css";

interface MovieToWatch {
  id: number;
  name: string;
  poster_path: string;
  time_added?: string;
}

const WatchListPage: React.FC = () => {
  const [movies, setMovies] = useState<MovieToWatch[]>([]);

  useEffect(() => {
    const fetchMovies = async () => {
      try {
        const data = await getMoviesToWatch();
        setMovies(data);
      } catch (error) {
        console.error("Error fetching movies to watch:", error);
      }
    };

    fetchMovies();
  }, []);

  return (
    <div className="watchlist-page">
      <div className="watchlist-header">
        <h1>My Watch List</h1>
        <Link to="/" className="button-style back-button">Back to Home</Link>
      </div>
      <div className="watchlist-movies-grid">
        {movies.length > 0 ? (
          movies.map((movie) => (
            <MovieCard
              key={movie.id}
              original_title={movie.name}
              poster_path={movie.poster_path}
              addedDate={movie.time_added}
              onClick={() => console.log(`Clicked on movie ID: ${movie.id}`)}
            />
          ))
        ) : (
          <p>No movies in your watch list.</p>
        )}
      </div>
      <Link to="/" className="button-style">
        Back to Home
      </Link>
    </div>
  );
};

export default WatchListPage;

