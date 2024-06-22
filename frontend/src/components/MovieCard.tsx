import React from "react";

interface MovieCardProps {
    title: string;
    poster: string;
}

const MovieCard: React.FC<MovieCardProps> = ({ title, poster }) => {
    return (
        <div>
            <h3>{title}</h3>
            <h2>"aaaaa"</h2>
            <img src={`https://image.tmdb.org/t/p/${poster}`} alt={title} />
        </div>
    );
};

export default MovieCard;
