import React, { useEffect } from "react";
import "./App.scss";
import Loader from "./components/ui/Loader";
import { useAuth } from "./store/auth";

const App: React.FC = () => {
  const { getUser, isLoading, currentUser } = useAuth();

  useEffect(() => {
    getUser(true);
  }, [getUser]);

  return (
    <section className="section">
      <div className="container">
        <h1 className="title">Hello Mouth-breathing-homosapiens!</h1>
        {isLoading && <Loader />}
        {currentUser && <h1>{currentUser.name}</h1>}
      </div>
    </section>
  );
};

export default App;
