import React from 'react';

import Home from './components/Home/Home';

import AuthProvider from './components/authProvider';
import Routes from './routes';

import './App.css';

function App() {
  return (
    <AuthProvider  >
      <Routes />
    </AuthProvider>
  );
}

export default App;
