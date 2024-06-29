import React, { useEffect, useState } from "react";
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
      <h1>My Watch List</h1>
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
    </div>
  );
};

export default WatchListPage;
