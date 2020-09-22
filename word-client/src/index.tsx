import React from "react";
import ReactDOM from "react-dom";
import { QueryCache, ReactQueryCacheProvider } from "react-query";
import App from "./App";

const queryCache = new QueryCache();

ReactDOM.render(
  <React.StrictMode>
    <ReactQueryCacheProvider queryCache={queryCache}>
      <App />
    </ReactQueryCacheProvider>
  </React.StrictMode>,
  document.getElementById("root")
);
