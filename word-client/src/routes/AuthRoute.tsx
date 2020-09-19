import React from "react";
import { Route, Redirect, RouteProps } from "react-router-dom";
import { User } from "../store/auth";

// Properties to add to standard route props
type AuthRouteProps = RouteProps & {
  user?: User;
  redirectPath: string;
};

const AuthRoute: React.FC<AuthRouteProps> = ({
  children,
  user,
  redirectPath,
  ...rest
}) => {
  return (
    <Route
      {...rest}
      render={({ location }) =>
        user ? (
          children
        ) : (
          <Redirect
            to={{
              pathname: redirectPath,
              state: {
                from: location,
              },
            }}
          />
        )
      }
    />
  );
};

export default AuthRoute;
