import React, { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import { getMoviesToWatch } from "../services/api";
import "./WatchListPage.css";

interface MovieToWatch {
  id: number;
  name: string;
  poster_path: string;
  time_added: string;
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
        <Link to="/" className="button-style back-button">
          Back to Home
        </Link>
      </div>
      <div className="watchlist-movies-grid">
        {movies.length > 0 ? (
          movies.map((movie) => (
            <div className="watchlist-movie-card" key={movie.id}>
              <h3>{movie.name}</h3>
              <img
                src={`https://image.tmdb.org/t/p/w500${movie.poster_path}`}
                alt={movie.name}
              />
              <p>Added on: {new Date(movie.time_added).toLocaleString()}</p>
            </div>
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
