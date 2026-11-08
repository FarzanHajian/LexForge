// Copyright (c) 2026, Farzan Hajian
// All rights reserved.
//
// This source code is licensed under the BSD 3-Clause license found in the
// LICENSE file in the root directory of this source tree.

import ToastX from "./ToastX";

export type InfoTypes = "data-saved" | "data-deleted";
const INFO_MESSAGES = {
    "data-saved": "Data saved successfully.",
    "data-deleted": "Data deleted successfully.",
} as const;

export default function InfoDisplay(props: { infoType: InfoTypes; onDismiss: () => void }) {
    const { infoType, onDismiss } = props;

    return (
        <ToastX
            message={INFO_MESSAGES[infoType]}
            toastType="success"
            onDismiss={() => onDismiss()}
        />
    );
}
