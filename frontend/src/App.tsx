// Copyright (c) 2026, Farzan Hajian
// All rights reserved.
//
// This source code is licensed under the BSD 3-Clause license found in the
// LICENSE file in the root directory of this source tree.

import MainMenu from "./components/MainMenu";
import Landing from "./components/Landing";
import { useAuth0 } from "@auth0/auth0-react";
import { useEffect, useMemo, useState } from "react";
import { getMyUser } from "./services/userService";
import Home from "./components/Home";
import { Spinner } from "flowbite-react/components/Spinner";
import { Routes, Route, Navigate } from "react-router";
import TodayItems from "./components/TodayItems";
import StudySessions from "./components/StudySessions";
import Manage from "./components/Manage";
import type { User } from "./models/user";
import { useApiToken } from "./utils/useApiToken";
import { UserContext } from "./utils/userContext";
import { toError, showError } from "./utils/utils";
import ToastX from "./components/ToastX";

export default function App() {
    const { isAuthenticated, isLoading, error: authError } = useAuth0();
    const [user, setUser] = useState<User | null>(null);
    const getToken = useApiToken();

    // Load the user on login. `user` is deliberately not a dependency: this
    // effect sets it, so depending on it would loop forever. Components that
    // change the user (e.g. Settings) update it via context instead.
    useEffect(() => {
        let ignore = false;

        if (!isAuthenticated) {
            if (authError && !ignore) showError(toError(authError));
            return;
        }

        (async () => {
            try {
                const token = await getToken();
                const me = await getMyUser(token);
                if (!ignore) setUser(me);
            } catch (err) {
                if (!ignore) showError(toError(err));
            }
        })();

        return () => {
            ignore = true;
        };
    }, [isAuthenticated, getToken, authError]);

    // Never expose a stale user after logout; derived here rather than cleared in the effect.
    const currentUser = isAuthenticated ? user : null;
    const userContextValue = useMemo(() => ({ user: currentUser, setUser }), [currentUser]);

    return (
        <UserContext value={userContextValue}>
            <MainMenu />
            <section role="main" className="px-4 lg:px-6 grow">
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
            <ToastX />
        </UserContext>
    );
}
