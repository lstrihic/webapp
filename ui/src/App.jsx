import React from 'react';
import { Helmet } from 'react-helmet-async';
import { BrowserRouter as Router } from 'react-router-dom';

import { Routing } from './Routing';

export const App = () => {
  return (
    <Router>
      <Helmet defaultTitle="Web App" titleTemplate="%s - Web App">
        <title>Web App</title>
      </Helmet>
      <Routing />
    </Router>
  );
};
