import { Outlet } from "react-router-dom";

// COMPONENTS
import Navbar from "./components/Navbar";
import Footer from "./components/Footer";



function App() {
  return (
    <div className="App">
      <Navbar />
      <Outlet /> 
      <Footer />
    </div>
  );
}

export default App;
