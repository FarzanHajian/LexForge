// Copyright (c) 2026, Farzan Hajian
// All rights reserved.
//
// This source code is licensed under the BSD 3-Clause license found in the
// LICENSE file in the root directory of this source tree.

import { createContext } from "react";
import type { User } from "../models/user";

export type UserContextValue = {
    user: User | null;
    setUser: (user: User | null) => void;
};

// Created once and shared by the provider (App) and all consumers.
// The default value is only a no-op fallback for components rendered without a provider.
export const UserContext = createContext<UserContextValue>({
    user: null,
    setUser: () => { },
});
