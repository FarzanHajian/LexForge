// Copyright (c) 2026, Farzan Hajian
// All rights reserved.
//
// This source code is licensed under the BSD 3-Clause license found in the
// LICENSE file in the root directory of this source tree.

export default function ErrorDisplay(props: { message: string }) {
    return(
        <div className="text-red-500">{props.message}</div>
    );
}