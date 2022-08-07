import { useNavigate } from "react-router-dom";
import { useSelector, useDispatch } from 'react-redux';
import { Link } from "react-router-dom";

// ALERT
import { useAlert } from '@blaumaus/react-alert'

import { deleteToken } from '../states/actions/token';


const Navbar = () => {

  const alert = useAlert();
  const navigate = useNavigate();
  const dispatch = useDispatch();
  const token = useSelector(state => state.token);

  const logoutHandler = (e) => {
    dispatch(deleteToken());
    alert.success("Logged out.")
    navigate("/");
  }
  
  return (
    <nav className="navbar navbar-expand-lg navbar-dark bg-dark">
      <Link to="/" className="navbar-brand">{process.env.REACT_APP_COMPANY_NAME.toUpperCase()}</Link>
      <button className="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
        <span className="navbar-toggler-icon"></span>
      </button>
      <div className="collapse navbar-collapse" id="navbarSupportedContent">
        <ul className="navbar-nav mr-auto">
          <li className="nav-item">
            <Link to="/" className="nav-link">Home</Link>
          </li>
          <li className="nav-item">
            <Link to="/protected/leaderboard" className="nav-link">Leaderboard</Link>
          </li>
        </ul>
        <div className="form-inline my-2 my-lg-0">
          {token.length > 0 ?
            <>
              <Link to="/profile" className="btn btn-outline-success mr-sm-2">Profile</Link>
              <button className="btn btn-outline-success mr-sm-2" type="button" onClick={logoutHandler}>Logout</button>
            </>
            :
            <>
              <Link to="/login" className="btn btn-outline-success mr-sm-2">Login</Link>
              <Link to="/register" className="btn btn-outline-success mr-sm-2">Register</Link>
            </>
          }
        </div>
      </div>
    </nav>
  )
}

export default Navbar
