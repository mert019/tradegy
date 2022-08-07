export const setToken = (token) => {
    return {
        type: 'SET_TOKEN',
        payload: token
    };
};

export const deleteToken = () => {
    return {
        type: 'DELETE_TOKEN'
    };
};
