import React from 'react';
import ReactDOM from 'react-dom/client';

// REACT ROUTER
import { BrowserRouter, Routes, Route } from "react-router-dom";

// REDUX
import { Provider } from 'react-redux';

// REDUX PERSIST
import { PersistGate } from 'redux-persist/integration/react';
import { store, persistor } from './states/configureStore';

// PAGES
import LoginPage from './pages/LoginPage';
import MainPage from './pages/MainPage';
import RegisterPage from './pages/RegisterPage';
import ProfilePage from './pages/ProfilePage';
import CreateOrderPage from './pages/CreateOrderPage';
import LeaderboardPage from './pages/LeaderboardPage';

// ALERT
import { Provider as AlertProvider } from '@blaumaus/react-alert'

// COMPONENTS
import Alert, { alertOptions } from './components/Alert';
import AnonymousRoute from './components/authComponents/AnonymousRoute';
import ProtectedRoute from './components/authComponents/ProtectedRoute';

// CSS
import './index.css';
import 'devextreme/dist/css/dx.light.css';

import App from './App';



const root = ReactDOM.createRoot(document.getElementById('root'));

root.render(
  <React.StrictMode>
    <Provider store={store}>
      <PersistGate loading={null} persistor={persistor}>
        <AlertProvider template={Alert} {...alertOptions}>
          <BrowserRouter>
            <Routes>
              
              <Route path="/" element={<App />} >
                <Route path="" element={<MainPage />} />
                <Route path="leaderboard" element={<LeaderboardPage />} />
                <Route path="*" element={<p>404 PAGE</p>} />
              </Route>

              <Route path="/" element={<AnonymousRoute />} >
                <Route path="login" element={<LoginPage />} />
                <Route path="register" element={<RegisterPage />} />
                <Route path="*" element={<p>404 PAGE</p>} />
              </Route>

              <Route path="/" element={<ProtectedRoute />} >
                <Route path="profile" element={<ProfilePage />} />
                <Route path="createorder" element={<CreateOrderPage />} />
                <Route path="*" element={<p>404 PAGE</p>} />
              </Route>

            </Routes>
          </BrowserRouter>
        </AlertProvider>
      </PersistGate>
    </Provider>
  </React.StrictMode>
);
