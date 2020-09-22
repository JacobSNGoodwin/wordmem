import React from "react";
import { Link } from "react-router-dom";
import { User } from "../store/auth";

type NavbarProps = {
  currentUser?: User;
};

const Navbar: React.FC<NavbarProps> = ({ currentUser }) => {
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
          Your Day
        </Link>
        <Link to="/edit" className="navbar-item">
          Your List
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

  return (
    <nav className="navbar is-info" role="navigation">
      <div className="navbar-brand">
        <div className="navbar-item"></div>
      </div>
      {navigationMenu}
    </nav>
  );
};

export default Navbar;
