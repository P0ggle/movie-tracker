import axios from "axios";

const API_URL = process.env.REACT_APP_API_URL || "http://localhost:8080";

const getToken = () => localStorage.getItem("token");
const setToken = (token: string) => localStorage.setItem("token", token);
const clearToken = () => localStorage.removeItem("token");

const setUser = (username: string) => localStorage.setItem("username", username);
const getUser = () => localStorage.getItem("username");
const clearUser = () => localStorage.removeItem("username");

export const login = async (username: string, password: string) => {
  const response = await axios.post(`${API_URL}/login`, { username, password });
  if (response.data.token) {
    setToken(response.data.token);
    setUser(username); // Store the username
  }
  return response.data;
};

export const signup = async (username: string, password: string, email: string) => {
  const response = await axios.post(`${API_URL}/signup`, { username, password, email });
  return response.data;
};

export const fetchMovie = async (id: string) => {
  const response = await axios.get(`${API_URL}/movies/${id}`, {
    headers: { Authorization: `Bearer ${getToken()}` },
  });
  return response.data;
};

export const searchMovies = async (name: string) => {
  const response = await axios.get(`${API_URL}/search`, {
    params: { name },
    headers: { Authorization: `Bearer ${getToken()}` },
  });
  return response.data;
};

export const addMovieToList = async (name: string, posterPath: string) => {
  const response = await axios.post(
    `${API_URL}/add-to-list`,
    { name, poster_path: posterPath },
    {
      headers: {
        Authorization: `Bearer ${getToken()}`,
      },
    }
  );
  return response.data;
};

export const getMoviesToWatch = async () => {
  const response = await axios.get(`${API_URL}/movies-to-watch`, {
    headers: { Authorization: `Bearer ${getToken()}` },
  });
  return response.data;
};

export const updateWatchedStatus = async (id: number, watched: boolean) => {
  const payload = { watched };
  const response = await axios.put(`${API_URL}/movies/${id}/watched`, payload, {
    headers: { Authorization: `Bearer ${getToken()}` },
  });
  return response.data;
};

export const logout = () => {
  clearToken();
  clearUser();
};

export { getToken, getUser };
