// Copyright (c) 2026, Farzan Hajian
// All rights reserved.
//
// This source code is licensed under the BSD 3-Clause license found in the
// LICENSE file in the root directory of this source tree.

import { Button } from "flowbite-react";
import LexForge from "./LexForge";
import { useAuth0 } from "@auth0/auth0-react";

export default function Landing() {
    const { loginWithRedirect } = useAuth0();

    return (
        <div className="relative flex flex-col items-center justify-center h-full space-y-4">
            <img
                src="/logo-512.png"
                className="absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 opacity-5"
                alt="LexForge Logo"
            />

            <div className="relative z-10 text-4xl lg:text-6xl font-bold mb-4 flex">
                <span>Welcome to&nbsp;</span>
                <LexForge />
            </div>
            <div className="relative z-10 text-2xl">To continue, please login</div>

            <Button className="relative z-10 primary-button" onClick={() => loginWithRedirect()} >
                Login
            </Button>
        </div>
    );
}
