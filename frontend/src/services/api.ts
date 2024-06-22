import axios from "axios";

const API_URL = process.env.REACT_APP_API_URL || "http://localhost:8080";

export const fetchMovie = async (id: string) => {
    const response = await axios.get(`${API_URL}/movies/${id}`);
    return response.data;
};

export const searchMovie = async (name: string) => {
    const response = await axios.get(`${API_URL}/search`, { params: { name } });
    return response.data;
};
