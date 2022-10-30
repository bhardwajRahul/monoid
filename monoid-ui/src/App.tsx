import React from 'react';
import './App.css';

import { ApolloClient, InMemoryCache, ApolloProvider } from '@apollo/client';
import MonoidRouter from './pages/router';

const client = new ApolloClient({
  uri: `${process.env.REACT_APP_API_URL}/query`,
  cache: new InMemoryCache(),
});

function App() {
  return (
    <ApolloProvider client={client}>
      <MonoidRouter />
    </ApolloProvider>
  );
}

export default App;