import { combineReducers } from "redux";

// IMPORT REDUCERS
import tokenReducer from "./token";



const allReducers = combineReducers({
    token: tokenReducer
});

export default allReducers;
