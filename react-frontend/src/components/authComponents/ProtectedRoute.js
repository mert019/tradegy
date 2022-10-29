import { Navigate, Outlet } from "react-router-dom";
import { useSelector } from 'react-redux';

// COMPONENTS
import Navbar from "../Navbar";
import Footer from "../Footer";



const ProtectedRoute = () => {

  const token = useSelector(state => state.token);

  return (token.length > 0 ?
    <>
      <Navbar />
      <div className="main-body">
        <Outlet />
      </div>
      <Footer />
    </> :
    <Navigate to="/" replace />)
}

export default ProtectedRoute;
