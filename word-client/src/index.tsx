import React from "react";
import ReactDOM from "react-dom";
import { QueryCache, ReactQueryCacheProvider } from "react-query";
import App from "./App";

const queryCache = new QueryCache({
  defaultConfig: {
    queries: {
      staleTime: 1000,
      retry: 1,
      retryDelay: 500,
    },
  },
});

ReactDOM.render(
  <React.StrictMode>
    <ReactQueryCacheProvider queryCache={queryCache}>
      <App />
    </ReactQueryCacheProvider>
  </React.StrictMode>,
  document.getElementById("root")
);
