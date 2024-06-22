import React, { useState, useEffect } from "react";
import { fetchMovie } from "../services/api";
import MovieCard from "../components/MovieCard";

const HomePage: React.FC = () => {
    const [movie, setMovie] = useState<{ title: string; poster: string } | null>(
        null,
    );

    useEffect(() => {
        const loadMovie = async () => {
            const movieData = await fetchMovie("27205"); // Replace '123' with a valid movie ID
            setMovie(movieData);
        };
        loadMovie();
    }, []);

    return (
        <div>
            {movie ? (
                <MovieCard title={movie.title} poster={movie.poster} />
            ) : (
                <p>Loading...</p>
            )}
        </div>
    );
};

export default HomePage;
