import React from "react";

interface MovieCardProps {
    title?: string;
    original_name?: string;
    poster_path: string;
}

const MovieCard: React.FC<MovieCardProps> = ({ title, original_name, poster_path }) => {
    const posterUrl = poster_path
        ? `https://image.tmdb.org/t/p/w500${poster_path}`
        : "placeholder_image_url";

    return (
        <div>
            <h3>{title || original_name}</h3>
            <img src={posterUrl} alt={title || original_name} />
        </div>
    );
};

export default MovieCard;

