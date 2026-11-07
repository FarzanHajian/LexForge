// Copyright (c) 2026, Farzan Hajian
// All rights reserved.
//
// This source code is licensed under the BSD 3-Clause license found in the
// LICENSE file in the root directory of this source tree.

import { useCallback } from "react";
import { useAuth0 } from "@auth0/auth0-react";

export function useApiToken() {
    const { getAccessTokenSilently } = useAuth0();

    return useCallback(async (): Promise<string> => {
        const token = await getAccessTokenSilently();
        if (!token) {
            throw new Error("No access token available");
        }
        return token;
    }, [getAccessTokenSilently]);
}
