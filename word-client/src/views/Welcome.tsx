import React, { useEffect, useState } from "react";
import { useHistory } from "react-router-dom";
import { useAuth } from "../store/auth";

const Welcome: React.FC = () => {
  const [loginWindow, setLoginWindow] = useState<Window | undefined>(undefined);
  const history = useHistory();
  const getUser = useAuth((state) => state.getUser);

  if (loginWindow) {
    loginWindow.onbeforeunload = async () => {
      // try to navigate if other window is closed. If login window closed without successful login
      // we should be returned here
      await getUser(false); // if we don't await, user may not be in state in time for history.push("/")
      history.push("/");
    };
  }

  // not sure if this is necessary, but we'll do it!
  useEffect(() => {
    return () => {
      if (loginWindow) {
        loginWindow.close();
      }
    };
  });

  const openLoginWindow = () => {
    const popUp = window.open(
      "http://wordmem.test/account/authenticate?loginOnly",
      "_blank"
    );
    setLoginWindow(popUp ?? undefined);
  };

  return (
    <React.Fragment>
      <h1 className="title has-text-centered">Welcome to Wordmem!</h1>
      <div
        onClick={() => openLoginWindow()}
        className="buttons is-centered mt-6"
      >
        <button className="button is-info">Login</button>
      </div>
      <p className="has-text-centered mt-6">
        Insert marketing and buzzwords for management that happily employed
        engineers at the company don't even understand!
      </p>
    </React.Fragment>
  );
};

export default Welcome;
