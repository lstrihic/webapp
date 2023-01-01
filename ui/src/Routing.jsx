import React from 'react';
import { Route, Routes } from 'react-router-dom';

import { HomePage } from '@pages';

export const Routing = () => {
  return (
    <Routes>
      <Route path="/" element={<HomePage />} />
    </Routes>
  );
};
