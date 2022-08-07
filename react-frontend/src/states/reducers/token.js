const tokenReducer = (state = "", action) => {
    switch (action.type) {
        case 'SET_TOKEN':
            return action.payload;
        case 'DELETE_TOKEN':
            return "";
        default:
            return state
    }
}

export default tokenReducer;
