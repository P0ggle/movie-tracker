import React from "react";
import "./MovieCard.css";

interface MovieCardProps {
    original_title?: string;
    original_name?: string;
    poster_path: string;
    overview: string;
    onClick: () => void;
}

const MovieCard: React.FC<MovieCardProps> = ({
    original_title,
    original_name,
    poster_path,
    onClick,
}) => {
    const placeholderImageUrl =
        "https://via.placeholder.com/500x750?text=No+Image+Available";
    const posterUrl = poster_path
        ? `https://image.tmdb.org/t/p/w500${poster_path}`
        : placeholderImageUrl;

    return (
        <div className="movie-card" onClick={onClick}>
            <img
                src={posterUrl}
                alt={original_title || original_name || "Movie Poster"}
            />
            <h3>{original_title || original_name}</h3>
        </div>
    );
};

export default MovieCard;
