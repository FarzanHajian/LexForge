
import { Navbar, NavbarBrand, NavbarCollapse, NavbarLink, NavbarToggle } from "flowbite-react";

export function MainMenu() {
    return (
        <Navbar fluid className="bg-primary-100 p-1">
            <NavbarBrand href="/">
                <img src="/favicon.png" className="mr-3 h-6 sm:h-9" alt="Flowbite React Logo" />
                <span className="self-center whitespace-nowrap text-xl font-semibold">
                    <span className="text-brand">Lex</span>
                    <span className="text-accent">Forge</span>
                </span>
            </NavbarBrand>
        </Navbar>
    );
}
