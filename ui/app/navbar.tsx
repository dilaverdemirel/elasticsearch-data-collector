'use client'
import React from "react";
import { usePathname} from 'next/navigation';
import { Container, Nav, NavDropdown, Navbar } from "react-bootstrap";


export default function AppNavBar() {
  const pathname = usePathname();
  return (

    <Navbar expand="lg" className="bg-body-tertiary">
      <Container style={{width:"54%"}}>
        <Navbar.Brand href="/">Elasticsearch Data Collector</Navbar.Brand>
        <Navbar.Toggle aria-controls="basic-navbar-nav" />
        <Navbar.Collapse id="basic-navbar-nav">
          <Nav className="me-auto">
            <Nav.Link href="/datasources" active={pathname.startsWith("/datasources")}>Datasources</Nav.Link>
            <Nav.Link href="/indices" active={pathname.startsWith("/indices")}>Indices</Nav.Link>
          </Nav>
        </Navbar.Collapse>
      </Container>
    </Navbar>

    // <Navbar isBordered>
    //   <NavbarBrand>
    //     <AcmeLogo />
    //     <p className="font-bold text-inherit">
    //       <Link color="foreground" href="/">
    //         ES-DataCollector
    //       </Link>
    //     </p>
    //   </NavbarBrand>
    //   <NavbarContent className="sm:flex gap-4" justify="center">
    //     <NavbarItem isActive = {"/datasources"=== pathname}>
    //       <Link color="foreground" href="/datasources">
    //         Datasources
    //       </Link>
    //     </NavbarItem>
    //     <NavbarItem isActive = {"/indices"=== pathname}>
    //       <Link color="foreground" href="/indices">
    //         Indices
    //       </Link>
    //     </NavbarItem>
    //   </NavbarContent>
    //   <NavbarContent justify="end">
    //     <NavbarItem>
    //       <Button as={Link} color="primary" href="/about" variant="flat">
    //         About
    //       </Button>
    //     </NavbarItem>
    //   </NavbarContent>
    // </Navbar>
  );
}
