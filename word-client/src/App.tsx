import React, { useEffect } from "react";
import "./App.scss";
import Loader from "./components/ui/Loader";
import { useAuth } from "./store/auth";

const App: React.FC = () => {
  const getUser = useAuth((state) => state.getUser);
  const isLoading = useAuth((state) => state.isLoading);
  const currentUser = useAuth((state) => state.currentUser);

  useEffect(() => {
    getUser(true);
  }, [getUser]);

  return (
    <section className="section">
      <div className="container">
        <h1 className="title">Welcome to WordMem</h1>
        {isLoading && <Loader />}
        {currentUser && <h1>{currentUser.name}</h1>}
      </div>
    </section>
  );
};

export default App;
