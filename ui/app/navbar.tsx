'use client'
import React from "react";
import { usePathname } from 'next/navigation';
import { Container, Nav, NavDropdown, Navbar } from "react-bootstrap";
import Image from 'react-bootstrap/Image';


export default function AppNavBar() {
  const pathname = usePathname();
  return (

    <Navbar expand="lg" className="bg-body-tertiary">
      <Container style={{ width: "54%" }}>
        <Navbar.Brand href="/">
          <Image src="/logo.png" style={{ width: 50 }} alt="Elasticsearch Data Collector" title="Elasticsearch Data Collector"/>
        </Navbar.Brand>
        <Navbar.Toggle aria-controls="basic-navbar-nav" />
        <Navbar.Collapse id="basic-navbar-nav">
          <Nav className="me-auto">
            <Nav.Link href="/datasources" active={pathname.startsWith("/datasources")}>Datasources</Nav.Link>
            <Nav.Link href="/indices" active={pathname.startsWith("/indices")}>Indices</Nav.Link>
            <Nav.Link href="/about" active={pathname.startsWith("/about")}>About & Help</Nav.Link>
          </Nav>
        </Navbar.Collapse>
      </Container>
    </Navbar>
  );
}
