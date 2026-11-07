// Copyright (c) 2026, Farzan Hajian
// All rights reserved.
//
// This source code is licensed under the BSD 3-Clause license found in the
// LICENSE file in the root directory of this source tree.

import ToastX from "./ToastX";

export default function ErrorDisplay(props: { error: Error; onDismiss: () => void; timeout?: number }) {
    const { error, onDismiss, timeout = 5000 } = props;

    return (
        <ToastX
            message={error.message}
            toastType="error"
            onDismiss={() => onDismiss()}
            timeout={timeout}
        />
    );
}
