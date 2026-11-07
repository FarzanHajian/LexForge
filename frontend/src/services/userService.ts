// Copyright (c) 2026, Farzan Hajian
// All rights reserved.
//
// This source code is licensed under the BSD 3-Clause license found in the
// LICENSE file in the root directory of this source tree.

import type { User } from "../models/user";
import { get, put } from "./apiService";

export async function getMyUser(token: string): Promise<User> {
    return get<User>('/users/me', token);
}

export async function editMyUser(token: string, user: Partial<User>): Promise<User> {
    return put<User>('/users/me', token, user);
}