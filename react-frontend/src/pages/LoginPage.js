import { useState } from "react";
import { useDispatch } from "react-redux";
import { useNavigate } from "react-router-dom";

// SERVICES
import AuthService from "../services/AuthService";

// REDUX ACTIONS
import { setToken } from "../states/actions/token"

// ALERT
import { useAlert } from '@blaumaus/react-alert'



const LoginPage = () => {

  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");

  const dispatch = useDispatch();
  const navigate = useNavigate();
  const alert = useAlert()

  const loginHandler = async (e) => {
    e.preventDefault();
    // Validate input
    if (username.length === 0 || password.length === 0) {
      alert.error("Please enter a username and password.")
      return;
    }
    // Call login service
    let response = await AuthService.login(username, password);
    if (response === null) {
      alert.error("Error on login service. Please try a few minites later.");
      return;
    }
    // Response actions
    if (200 <= response.Status && response.Status <= 299) {
      dispatch(setToken(response.Payload.Token));
      alert.success(response.Message);
      navigate("/");
    } else {
      alert.error(response.Message);
    }
  }

  const inputHandler = (e) => {
    let targetID = e.target.id;
    switch (targetID) {
      case "username":
        setUsername(e.target.value);
        break;
      case "password":
        setPassword(e.target.value);
        break;
    }
  }

  return (
    <div id="login" className="mb-5">
      <h2 className="text-center pt-5">Welcome To {process.env.REACT_APP_COMPANY_NAME}</h2>
      <div className="container">
        <div id="login-row" className="row justify-content-center align-items-center">
          <div id="login-column" className="col-md-6">
            <div id="login-box" className="col-md-12 border border-success rounded">
              <form id="login-form" className="form" onSubmit={loginHandler}>
                <h4 className="text-center mt-3">Login</h4>
                <div className="form-group">
                  <label htmlFor="username">Username:</label><br />
                  <input type="text" name="username" id="username" className="form-control" value={username} onChange={inputHandler} />
                </div>
                <div className="form-group">
                  <label htmlFor="password">Password:</label><br />
                  <input type="password" name="password" id="password" className="form-control" value={password} onChange={inputHandler} />
                </div>
                <div className="form-group text-center">
                  <input type="submit" name="submit" className="btn btn-success btn-md" value="Login" />
                </div>
              </form>
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}

export default LoginPage;
