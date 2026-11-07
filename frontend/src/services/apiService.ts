// Copyright (c) 2026, Farzan Hajian
// All rights reserved.
//
// This source code is licensed under the BSD 3-Clause license found in the
// LICENSE file in the root directory of this source tree.

import { ApiError } from "../utils/apiError";

export const API_BASE = import.meta.env.VITE_LEXFORGE_SERVER_ADDRESS ?? ''

export async function get<T>(path: string, token: string): Promise<T> {
    return apiFetch<T>(path, token, { method: 'GET' });
}

export async function post<T>(path: string, token: string, body: unknown): Promise<T> {
    return apiFetch<T>(path, token, { method: 'POST', body: JSON.stringify(body) });
}

export async function put<T>(path: string, token: string, body: unknown): Promise<T> {
    return apiFetch<T>(path, token, { method: 'PUT', body: JSON.stringify(body) });
}

export async function del<T>(path: string, token: string): Promise<T> {
    return apiFetch<T>(path, token, { method: 'DELETE' });
}

export async function apiFetch<T>(path: string, token: string, init?: RequestInit): Promise<T> {
    try {
        const res = await fetch(`${API_BASE}${path}`, {
            headers: {
                'Content-Type': 'application/json',
                Authorization: `Bearer ${token}`,
                ...init?.headers,
            },
            ...init,
        });

        if (!res.ok) {
            const body = await res.json().catch(() => null);
            throw new ApiError(res.status, body?.code, body?.message ?? res.statusText);
        }

        return res.json() as Promise<T>;
    } catch (ex: any) {
        throw ApiError.unreachableServer();
    }
}