// Copyright (c) 2026, Farzan Hajian
// All rights reserved.
//
// This source code is licensed under the BSD 3-Clause license found in the
// LICENSE file in the root directory of this source tree.

import MainMenu from "./components/MainMenu";
import Landing from "./components/Landing";
import { useAuth0 } from "@auth0/auth0-react";
import { Button } from "flowbite-react";
import ErrorDisplay from "./components/ErrorDisplay";


export default function App() {
    const { isAuthenticated, isLoading, error } = useAuth0();

    return (
        <>
            <MainMenu />
            <section role="main" className="px-4 lg:px-6 grow">
                {isLoading ? <div>Loading...</div> : null}

                {error ? <ErrorDisplay message={error.message} /> : null}

                {isAuthenticated ? (
                    <>
                        <h1 className="text-brand">Hello World!</h1>
                        <p className="text-text">
                            Contrary to popular belief, Lorem Ipsum is not simply random text. It has roots in a piece of classical Latin literature from 45 BC,
                            making it over 2000 years old. Richard McClintock, a Latin professor at Hampden-Sydney College in Virginia, looked up one of the more
                            obscure Latin words, consectetur, from a Lorem Ipsum passage, and going through the cites of the word in classical literature, discovered
                            the undoubtable source. Lorem Ipsum comes from sections 1.10.32 and 1.10.33 of "de Finibus Bonorum et Malorum" (The Extremes of Good and Evil)
                            by Cicero, written in 45 BC. This book is a treatise on the theory of ethics, very popular during the Renaissance. The first line of Lorem
                            Ipsum, "Lorem ipsum dolor sit amet..", comes from a line in section 1.10.32.

                            The standard chunk of Lorem Ipsum used since 1966 is reproduced below for those interested. Sections 1.10.32 and 1.10.33 from "de Finibus
                            Bonorum et Malorum" by Cicero are also reproduced in their exact original form, accompanied by English versions from the 1914 translation
                            by H. Rackham.
                        </p>
                        <Button className="bg-brand hover:bg-brand-hover">Custom Button</Button>
                    </>
                ) : (
                    <Landing />

                )}
            </section>
        </>
    )
}