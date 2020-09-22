import React, { useEffect, useState } from "react";
import { BrowserRouter, Link, Route, Switch } from "react-router-dom";
import "./App.scss";
import Loader from "./components/ui/Loader";
// import Loader from "./components/ui/Loader";
import AuthRoute from "./routes/AuthRoute";
import { useAuth } from "./store/auth";
import Edit from "./views/Edit";
import Overview from "./views/Overview";
import Welcome from "./views/Welcome";

const App: React.FC = () => {
  const getUser = useAuth((state) => state.getUser);
  const [beginUserLoad, setBeginUserLoad] = useState(false);
  const isLoading = useAuth((state) => state.isLoading);
  const currentUser = useAuth((state) => state.currentUser);

  useEffect(() => {
    getUser(true);
    setBeginUserLoad(true);
  }, [getUser]);

  const placeHolderImage = (
    <div
      style={{
        height: 48,
        width: 48,
        backgroundColor: "hsl(0, 0%, 86%)",
        borderRadius: 24,
        display: "flex",
        flexDirection: "column",
        justifyContent: "center",
      }}
    >
      <div
        style={{
          textAlign: "center",
          fontSize: 20,
          fontWeight: "bold",
        }}
      >
        {currentUser?.name ? currentUser.name[0].toUpperCase() : "U"}
      </div>
    </div>
  );

  const navigationMenu = currentUser ? (
    <div className="navbar-menu">
      <div className="navbar-start">
        <Link to="/" className="navbar-item">
          Overview
        </Link>
        <Link to="/edit" className="navbar-item">
          Edit
        </Link>
      </div>
      <div className="navbar-end">
        <div className="navbar-item">
          <a href="/account" target="_blank">
            <figure className="image">
              {currentUser.imageUrl ? (
                <img src={currentUser.imageUrl} alt="Profile" />
              ) : (
                placeHolderImage
              )}
            </figure>
          </a>
        </div>
      </div>
    </div>
  ) : undefined;

  // since the auth state's isLoading is initially false, we need to make
  // sure we also initiating the auth state check (getUser) before loading routes
  const routes =
    beginUserLoad && !isLoading ? (
      <Switch>
        <Route exact path="/welcome">
          <Welcome />
        </Route>
        <AuthRoute
          user={currentUser}
          exact
          path="/edit"
          redirectPath="/welcome"
        >
          <Edit />
        </AuthRoute>
        <AuthRoute user={currentUser} exact path="/" redirectPath="/welcome">
          <Overview />
        </AuthRoute>
      </Switch>
    ) : undefined;

  return (
    <BrowserRouter>
      <nav className="navbar is-info" role="navigation">
        <div className="navbar-brand">
          <div className="navbar-item"></div>
        </div>
        {navigationMenu}
      </nav>
      <section className="section">
        <div className="container">
          {isLoading || (!beginUserLoad && <Loader radius={200} />)}
          {routes}
        </div>
      </section>
    </BrowserRouter>
  );
};

export default App;
