import { useState } from "react";
import { useNavigate } from "react-router-dom";

// ALERT
import { useAlert } from '@blaumaus/react-alert'

// SERVICES
import RegisterService from "../services/RegisterService";



const RegisterPage = () => {

  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");

  const navigate = useNavigate();
  const alert = useAlert()

  const inputHandler = (e) => {
    let targetID = e.target.id;
    switch (targetID) {
      case "username":
        setUsername(e.target.value);
        break;
      case "password":
        setPassword(e.target.value);
        break;
      case "confirm-password":
        setConfirmPassword(e.target.value);
        break;
    }
  }

  const registerHandler = async (e) => {
    e.preventDefault();
    // Validate input
    if (username.length === 0 || password.length === 0) {
      alert.error("Please enter a username and password.")
      return;
    }
    if(password !== confirmPassword){
      alert.error("Passwords does not match.")
    }
    // Call register service
    let response = await RegisterService.register(username, password);
    if (response === null) {
      alert.error("Error on register service. Please try a few minites later.");
      return;
    }
    // Response actions
    if (200 <= response.Status && response.Status <= 299) {
      alert.success(response.Message);
      navigate("/login");
    } else {
      alert.error(response.Message);
    }
  }

  return (
    <div id="register" className="mb-5">
      <h2 className="text-center pt-5">Welcome To {process.env.REACT_APP_COMPANY_NAME}</h2>
      <div className="container">
        <div id="Register-row" className="row justify-content-center align-items-center">
          <div id="Register-column" className="col-md-6">
            <div id="Register-box" className="col-md-12 border border-success rounded">
              <form id="Register-form" className="form" onSubmit={registerHandler}>
                <h4 className="text-center mt-3">Register Now</h4>
                <div className="form-group">
                  <label htmlFor="username">Username:</label><br />
                  <input type="text" name="username" id="username" className="form-control" value={username} onChange={inputHandler} autoComplete="off"/>
                </div>
                <div className="form-group">
                  <label htmlFor="password">Password:</label><br />
                  <input type="password" name="password" id="password" className="form-control" value={password} onChange={inputHandler} />
                </div>
                <div className="form-group">
                  <label htmlFor="confirm-password">Confirm Password:</label><br />
                  <input type="password" name="confirm-password" id="confirm-password" className="form-control" value={confirmPassword} onChange={inputHandler} />
                </div>
                <div className="form-group text-center">
                  <input type="submit" name="submit" className="btn btn-success btn-md" value="Register" />
                </div>
              </form>
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}

export default RegisterPage;
