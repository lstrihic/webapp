import React from 'react';
import { BrowserRouter as Router } from 'react-router-dom';
import { Helmet } from 'react-helmet-async';

import { Routing } from './Routing';

function App() {
  return (
    <Router>
      <Helmet>
        <title>Web App</title>
      </Helmet>
      <Routing />
    </Router>
  );
}

export default App;
