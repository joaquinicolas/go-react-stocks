import React from "react";
import { Navbar, NavbarBrand } from "reactstrap";

export const Appbar = () => {
  return (
    <Navbar color="light" light>
      <NavbarBrand href="/">
        <svg className="align-middle" viewBox="0 0 5 5" width="28" height="28">
          <title>IEX home</title>
          <path d="M0,0v5h5V0z M1,1h3v3H2V3H1z" />
        </svg>
      </NavbarBrand>
    </Navbar>
  );
};
