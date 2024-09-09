import api from './api';

export const checkLogin = async (data) => {
    const response = await api.post('/login', data);
    return response;
};
