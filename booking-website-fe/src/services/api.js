import axios from 'axios';

const URL = process.env.URL;

const api = axios.create({
    baseURL: URL,
    timeout: 1000,
    headers: { 'Content-Type': 'application/json' },
});

export default api;
