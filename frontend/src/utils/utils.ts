// Copyright (c) 2026, Farzan Hajian
// All rights reserved.
//
// This source code is licensed under the BSD 3-Clause license found in the
// LICENSE file in the root directory of this source tree.

import toast from 'react-hot-toast';

export type InfoTypes = "data-saved" | "data-deleted";
const INFO_MESSAGES = {
    "data-saved": "Data saved successfully.",
    "data-deleted": "Data deleted successfully.",
} as const;

export function toError(error: unknown): Error {
    return error instanceof Error ? error : new Error(String(error));
}

export function showError(error: Error) {
    toast.error(error.message);
}

export function showInfo(infoType: InfoTypes) {
    toast.success(INFO_MESSAGES[infoType]);
}