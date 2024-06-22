import React, { useState, useEffect } from "react";
import { fetchMovie } from "../services/api";
import MovieCard from "../components/MovieCard";

const HomePage: React.FC = () => {
    const [movie, setMovie] = useState<{
        title: string;
        poster_path: string;
    } | null>(null);

    useEffect(() => {
        const loadMovie = async () => {
            try {
                const movieData = await fetchMovie("27205");
                setMovie(movieData);
            } catch (error) {
                console.error("Error fetching movie:", error);
            }
        };
        loadMovie();
    }, []);

    return (
        <div>
            {movie ? (
                <MovieCard title={movie.title} poster_path={movie.poster_path} />
            ) : (
                <p>Loading...</p>
            )}
        </div>
    );
};

export default HomePage;
