import React from "react";
import "./MovieCard.css";

interface MovieCardProps {
  original_title?: string;
  original_name?: string;
  poster_path: string;
  overview?: string;
  addedDate?: string;
  className?: string;
  onClick?: () => void;
}

const MovieCard: React.FC<MovieCardProps> = ({
  original_title,
  original_name,
  poster_path,
  addedDate,
  className,
  onClick,
}) => {
  const placeholderImageUrl =
    "https://via.placeholder.com/500x750?text=No+Image+Available";
  const posterUrl = poster_path
    ? `https://image.tmdb.org/t/p/w500${poster_path}`
    : placeholderImageUrl;

  return (
    <div className={`movie-card ${className}`} onClick={onClick}>
      <img
        src={posterUrl}
        alt={original_title || original_name || "Movie Poster"}
      />
      {className?.includes("movie-card-watched") && (
        <div className="watched-badge">Watched</div>
      )}
      <div className="movie-card-content">
        <h3>{original_title || original_name}</h3>
        {addedDate && (
          <p className="added-date">
            Added on: {new Date(addedDate).toLocaleString()}
          </p>
        )}
      </div>
    </div>
  );
};

export default MovieCard;

