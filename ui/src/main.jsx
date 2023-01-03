import { Global, css } from '@emotion/react';
import '@fontsource/roboto';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import 'bulma/css/bulma.css';
import React from 'react';
import ReactDOM from 'react-dom/client';
import { HelmetProvider } from 'react-helmet-async';

import { App } from './App';

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      refetchOnWindowFocus: false,
    },
  },
});

ReactDOM.createRoot(document.getElementById('app')).render(
  <React.StrictMode>
    <Global
      styles={css`
        #app {
          overflow: auto;
          display: block;
          width: 100%;
          height: 100vh;
          background: #f8f9fa;
        }
      `}
    />
    <QueryClientProvider client={queryClient}>
      <HelmetProvider>
        <App />
      </HelmetProvider>
    </QueryClientProvider>
  </React.StrictMode>
);
