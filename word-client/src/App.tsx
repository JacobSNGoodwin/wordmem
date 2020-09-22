import React, { useEffect, useState } from "react";
import { ReactQueryDevtools } from "react-query-devtools";
import { BrowserRouter, Route, Switch } from "react-router-dom";
import "./App.scss";
import Navbar from "./components/Navbar";
import Loader from "./components/ui/Loader";
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

  // since the auth state's isLoading is initially false, we need to make
  // sure we also initiating the auth state check (getUser) before loading routes
  // we could also create a config array of routs that could be shared betwen this and
  // the navbar
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
    <>
      <BrowserRouter>
        <Navbar currentUser={currentUser} />
        <section className="section">
          <div className="container">
            {isLoading || (!beginUserLoad && <Loader radius={200} />)}
            {routes}
          </div>
        </section>
      </BrowserRouter>
      <ReactQueryDevtools />
    </>
  );
};

export default App;
