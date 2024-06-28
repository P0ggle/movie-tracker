import React from "react";

interface MovieCardProps {
    original_title?: string;
    original_name?: string;
    poster_path: string;
    overview: string;
}

const MovieCard: React.FC<MovieCardProps> = ({
    original_title,
    original_name,
    poster_path,
    overview,
}) => {
    const posterUrl = poster_path
        ? `https://image.tmdb.org/t/p/w500${poster_path}`
        : "placeholder_image_url";

    return (
        <div>
            <h3>{original_title || original_name}</h3>
            <img src={posterUrl} alt={original_title || original_name} />
            <p>{overview}</p>
        </div>
    );
};

export default MovieCard;
