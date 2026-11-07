// Copyright (c) 2026, Farzan Hajian
// All rights reserved.
//
// This source code is licensed under the BSD 3-Clause license found in the
// LICENSE file in the root directory of this source tree.

const ERROR_CODE_UNKOWN: string = "UNKNOWN";
const ERROR_CODE_UNREACHABLE_SERVER: string = "UNREACHABLE_SERVER";

export class ApiError extends Error {


    readonly status: number;
    readonly code: string;

    constructor(status: number, code: string | null, message: string) {
        super(message);
        this.name = "ApiError";
        this.status = status;
        this.code = code ?? ERROR_CODE_UNKOWN;
    }

    static unreachableServer(): ApiError {
        return new ApiError(500, ERROR_CODE_UNREACHABLE_SERVER, "Server is unreachable now. Please try again later.");
    }
}