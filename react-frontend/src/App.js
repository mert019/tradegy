import { Outlet } from "react-router-dom";

// COMPONENTS
import Navbar from "./components/Navbar";
import Footer from "./components/Footer";



function App() {
  return (
    <div className="App">
      <Navbar />
      <div className="main-body">
        <Outlet />
      </div>
      <Footer />
    </div>
  );
}

export default App;
