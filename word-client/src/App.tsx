import React, { useEffect } from "react";
import "./App.scss";
import Loader from "./components/ui/Loader";
import { useAuth } from "./store/auth";

const App: React.FC = () => {
  const { getUser, isLoading } = useAuth();

  useEffect(() => {
    getUser();
  }, [getUser]);

  return (
    <section className="section">
      <div className="container">
        <h1 className="title">Hello World</h1>
        {isLoading && <Loader />}
      </div>
    </section>
  );
};

export default App;
