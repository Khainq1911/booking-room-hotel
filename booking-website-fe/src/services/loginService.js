import api from './api';

export const checkLogin = async (data) => {
    const response = await api.post('/login', data);
    return response;
};

export const logout = async () => {
    const response = await api.post('/logout');
    return response;
};
