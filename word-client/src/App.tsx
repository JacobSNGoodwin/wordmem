import React, { useEffect, useState } from "react";
import "./App.scss";
import Loader from "./components/ui/Loader";
import { useAuth } from "./store/auth";

const App: React.FC = () => {
  const [loginWindow, setLoginWindow] = useState<Window | undefined>(undefined);
  const getUser = useAuth((state) => state.getUser);
  const isLoading = useAuth((state) => state.isLoading);
  const currentUser = useAuth((state) => state.currentUser);

  if (loginWindow) {
    loginWindow.onbeforeunload = () => {
      getUser(false);
    };
  }

  useEffect(() => {
    getUser(true);
  }, [getUser]);

  // not sure if this is necessary, but we'll do it!
  useEffect(() => {
    return () => {
      if (loginWindow) {
        loginWindow.close();
      }
    };
  });

  const loginButton =
    !isLoading && !currentUser ? (
      <div
        onClick={() => openLoginWindow()}
        className="buttons is-centered mt-6"
      >
        <button className="button is-info">Login</button>
      </div>
    ) : null;

  const openLoginWindow = () => {
    const popUp = window.open(
      "http://wordmem.test/account/authenticate?loginOnly",
      "_blank"
    );
    setLoginWindow(popUp ?? undefined);
  };

  return (
    <section className="section">
      <div className="container">
        <h1 className="title has-text-centered">Welcome to WordMem</h1>
        {isLoading && <Loader />}
        {loginButton}
        {currentUser && currentUser.name}
      </div>
    </section>
  );
};

export default App;
