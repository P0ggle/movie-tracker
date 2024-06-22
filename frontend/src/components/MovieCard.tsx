import React from "react";

interface MovieCardProps {
    title: string;
    poster_path: string;
}

const MovieCard: React.FC<MovieCardProps> = ({ title, poster_path }) => {
    const posterUrl = poster_path
        ? `https://image.tmdb.org/t/p/w500${poster_path}`
        : "placeholder_image_url";

    return (
        <div>
            <h3>{title}</h3>
            <img src={posterUrl} alt={title} />
        </div>
    );
};

export default MovieCard;
