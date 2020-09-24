import React, { useState } from "react";
import { Link } from "react-router-dom";
import { useAuth, User } from "../store/auth";

type NavbarProps = {
  currentUser?: User;
};

const Navbar: React.FC<NavbarProps> = ({ currentUser }) => {
  const { signOut } = useAuth();
  const [isNavbarOpen, setNavbarOpen] = useState(false);
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
    <div className={`navbar-menu${isNavbarOpen ? " is-active" : ""}`}>
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
          <button onClick={() => signOut()} type="button" className="button">
            Sign Out
          </button>
        </div>
        <div className="navbar-item">
          <a href="/account" target="_blank">
            <figure className="image">
              {currentUser.imageUrl ? (
                <img
                  src={currentUser.imageUrl}
                  alt="Profile"
                  style={{ width: "auto" }}
                />
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
        {/* insert logo below */}
        <div className="navbar-item"></div>
        <div
          role="button"
          className={`navbar-burger burger${isNavbarOpen ? " is-active" : ""}`}
          aria-label="menu"
          aria-expanded="false"
          onClick={() => setNavbarOpen(!isNavbarOpen)}
        >
          <span aria-hidden="true"></span>
          <span aria-hidden="true"></span>
          <span aria-hidden="true"></span>
        </div>
      </div>

      {navigationMenu}
    </nav>
  );
};

export default Navbar;
