import React, { useState, useEffect } from "react";
import { Link } from "react-router-dom";
import { getMoviesToWatch, updateWatchedStatus } from "../services/api";
import MovieCard from "../components/MovieCard";
import "./WatchListPage.css";
import "./Popup.css"; // Import Popup styles

interface MovieToWatch {
  id: number;
  name: string;
  poster_path: string;
  time_added?: string;
  watched?: boolean;
}

const WatchListPage: React.FC = () => {
  const [movies, setMovies] = useState<MovieToWatch[]>([]);
  const [selectedMovie, setSelectedMovie] = useState<MovieToWatch | null>(null);
  const [sortOrder, setSortOrder] = useState<"all" | "watched" | "unwatched">("all");

  useEffect(() => {
    const fetchMovies = async () => {
      try {
        const data = await getMoviesToWatch();
        const sortedData = data.sort((a: MovieToWatch, b: MovieToWatch) => a.id - b.id);
        setMovies(sortedData);
      } catch (error) {
        console.error("Error fetching movies to watch:", error);
      }
    };

    fetchMovies();
  }, []);

  const handleWatchedToggle = async (movie: MovieToWatch) => {
    try {
      console.log("Toggling watched status for movie:", movie);
      await updateWatchedStatus(movie.id, !movie.watched);
      const updatedMovies = movies.map((m) =>
        m.id === movie.id ? { ...m, watched: !m.watched } : m
      );
      setMovies(updatedMovies);
      setSelectedMovie(null);
    } catch (error) {
      console.error("Error updating watched status:", error);
    }
  };

  const filteredMovies = movies.filter((movie) => {
    if (sortOrder === "watched") return movie.watched;
    if (sortOrder === "unwatched") return !movie.watched;
    return true;
  });

  return (
    <div className="watchlist-page">
      <div className="watchlist-header">
        <h1>My Watch List</h1>
        <div className="sort-container">
          <div className="sort-buttons">
            <button
              className={`button-style ${sortOrder === "all" ? "active" : ""}`}
              onClick={() => setSortOrder("all")}
            >
              All
            </button>
            <button
              className={`button-style ${sortOrder === "watched" ? "active" : ""}`}
              onClick={() => setSortOrder("watched")}
            >
              Watched
            </button>
            <button
              className={`button-style ${sortOrder === "unwatched" ? "active" : ""}`}
              onClick={() => setSortOrder("unwatched")}
            >
              Unwatched
            </button>
          </div>
        </div>
        <Link to="/" className="button-style link-to-home">
          Back to Home
        </Link>
      </div>
      <div className="watchlist-movies-grid">
        {filteredMovies.length > 0 ? (
          filteredMovies.map((movie) => (
            <MovieCard
              key={movie.id}
              original_title={movie.name}
              poster_path={movie.poster_path}
              addedDate={movie.time_added}
              className={movie.watched ? "movie-card-watched" : ""}
              onClick={() => setSelectedMovie(movie)}
            />
          ))
        ) : (
          <p>No movies in your watch list.</p>
        )}
      </div>
      {selectedMovie && (
        <div className="popup" onClick={() => setSelectedMovie(null)}>
          <div className="popup-content" onClick={(e) => e.stopPropagation()}>
            <h3>{selectedMovie.name}</h3>
            <img
              src={`https://image.tmdb.org/t/p/w500${selectedMovie.poster_path}`}
              alt={selectedMovie.name}
            />
            <button onClick={() => handleWatchedToggle(selectedMovie)}>
              {selectedMovie.watched ? "Mark as Unwatched" : "Mark as Watched"}
            </button>
            <button onClick={() => setSelectedMovie(null)}>Close</button>
          </div>
        </div>
      )}
    </div>
  );
};

export default WatchListPage;

