// Copyright (c) 2026, Farzan Hajian
// All rights reserved.
//
// This source code is licensed under the BSD 3-Clause license found in the
// LICENSE file in the root directory of this source tree.

import { Avatar, Dropdown, DropdownDivider, DropdownHeader, DropdownItem, Navbar, NavbarBrand, NavbarCollapse, NavbarLink, NavbarToggle } from "flowbite-react";
import LexForge from "./LexForge";
import { useAuth0 } from "@auth0/auth0-react";

export default function MainMenu() {
    const { user, isAuthenticated, isLoading, logout } = useAuth0();
    const generalAvatar = `data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' width='110' height='110' viewBox='0 0 110 110'%3E%3Ccircle cx='55' cy='55' r='55' fill='%2363b3ed'/%3E%3Cpath d='M55 50c8.28 0 15-6.72 15-15s-6.72-15-15-15-15 6.72-15 15 6.72 15 15 15zm0 7.5c-10 0-30 5.02-30 15v3.75c0 2.07 1.68 3.75 3.75 3.75h52.5c2.07 0 3.75-1.68 3.75-3.75V72.5c0-9.98-20-15-30-15z' fill='%23fff'/%3E%3C/svg%3E`;

    return (
        <Navbar fluid className="bg-primary-50 p-1 border-b border-b-primary-200">
            <NavbarBrand href="/">
                <img src="/favicon.png" className="mr-3 h-6 sm:h-9" alt="LexForge Logo" />
                <LexForge textSize="xl" />
            </NavbarBrand>


            {!isLoading && isAuthenticated && user
                ?
                <>
                    <div className="flex md:order-2">
                        <Dropdown arrowIcon={false} inline label={<Avatar img={user.picture || generalAvatar} alt="User's avatar" rounded />}>
                            <DropdownHeader>
                                <span className="block">{user.name}</span>
                            </DropdownHeader>
                            <DropdownDivider />
                            <DropdownItem>Settings</DropdownItem>
                            <DropdownItem onClick={() => logout({ logoutParams: { returnTo: window.location.origin } })}>Sign out</DropdownItem>
                        </Dropdown>

                        <NavbarToggle />
                    </div>

                    <NavbarCollapse>
                        <NavbarLink href="#" active>Home</NavbarLink>
                        <NavbarLink href="#">Items of the Day</NavbarLink>
                        <NavbarLink href="#">Notebooks</NavbarLink>
                    </NavbarCollapse>
                </>
                : null
            }

        </Navbar>
    );
}