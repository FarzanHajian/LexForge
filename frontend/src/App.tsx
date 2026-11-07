// Copyright (c) 2026, Farzan Hajian
// All rights reserved.
//
// This source code is licensed under the BSD 3-Clause license found in the
// LICENSE file in the root directory of this source tree.

import MainMenu from "./components/MainMenu";
import Landing from "./components/Landing";
import { useAuth0 } from "@auth0/auth0-react";
import { createContext, useEffect, useState } from "react";
import { getMyUser } from "./services/userService";
import ErrorDisplay from "./components/ErrorDisplay";
import Home from "./components/Home";
import { Spinner } from "flowbite-react/components/Spinner";
import { Routes, Route, Navigate } from "react-router";
import TodayItems from "./components/TodayItems";
import StudySessions from "./components/StudySessions";
import Manage from "./components/Manage";
import type { User } from "./models/user";
import { useApiToken } from "./utils/useApiToken";

export const UserContext = createContext<User | null>(null);

export default function App() {
    const { isAuthenticated, isLoading, error } = useAuth0();
    const [user, setUser] = useState<User | null>(null);
    const [error1, setError] = useState<Error | null>(null);
    const getToken = useApiToken();

    useEffect(() => {
        let ignore = false;

        if (isAuthenticated) {
            (async () => {
                try {
                    const token = await getToken();
                    const me = await getMyUser(token);
                    if (!ignore) setUser(me);
                } catch (err) {
                    if (!ignore) setError(toError(err));
                }
            })();
        }

        return () => {
            ignore = true;
        };
    }, [isAuthenticated, user]);


    return (
        <UserContext value={user}>
            <MainMenu />
            <section role="main" className="px-4 lg:px-6 grow">
                {error1 && <ErrorDisplay error={error1} onDismiss={() => setError(null)} />}
                {isLoading
                    ? <div className="relative flex flex-col items-center justify-center h-full space-y-4"><Spinner size="xl" />Loading...</div>
                    : (
                        isAuthenticated
                            ? (
                                <Routes>
                                    <Route path="/" element={<Home />} />
                                    <Route path="/today-items" element={<TodayItems />} />
                                    <Route path="/study-sessions" element={<StudySessions />} />
                                    <Route path="/manage" element={<Manage />} />
                                    <Route path="*" element={<Navigate to="/" replace />} />
                                </Routes>
                            )
                            : <Landing />
                    )
                }
            </section>
        </UserContext>
    );
}

function toError(err: unknown): Error {
    return err instanceof Error ? err : new Error(String(err));
}
