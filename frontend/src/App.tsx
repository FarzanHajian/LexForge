// Copyright (c) 2026, Farzan Hajian
// All rights reserved.
//
// This source code is licensed under the BSD 3-Clause license found in the
// LICENSE file in the root directory of this source tree.

import MainMenu from "./components/MainMenu";
import Landing from "./components/Landing";
import { useAuth0 } from "@auth0/auth0-react";
import ErrorDisplay from "./components/ErrorDisplay";
import Home from "./components/Home";
import { Spinner } from "flowbite-react/components/Spinner";

export default function App() {
    const { isAuthenticated, isLoading, error } = useAuth0();

    return (
        <>
            <MainMenu />
            <section role="main" className="px-4 lg:px-6 grow">
                {isLoading
                    ? <div className="relative flex flex-col items-center justify-center h-full space-y-4"><Spinner size="xl" />Loading...</div>
                    : error
                        ? <ErrorDisplay message={error.message} />
                        : (isAuthenticated ? <Home /> : <Landing />)
                }
            </section>
        </>
    )
}