// Copyright (c) 2026, Farzan Hajian
// All rights reserved.
//
// This source code is licensed under the BSD 3-Clause license found in the
// LICENSE file in the root directory of this source tree.

import { useEffect, useState } from "react";
import { Toast, ToastToggle } from "flowbite-react";
import { HiCheckCircle, HiXCircle } from "react-icons/hi2";

const CLOSE_DURATION_MS = 300;

export type ToastTypes = "success" | "error";

export default function ToastX(props: { message: string; toastType: ToastTypes; onDismiss: () => void; timeout?: number }) {
    const { message, onDismiss, timeout = 5000 } = props;
    const [closing, setClosing] = useState(false);

    const toastClasses = `group fixed top-5 right-5 z-50 overflow-hidden animate-toast-in transition-all ${closing ? "translate-x-4 opacity-0" : ""}`;
    const iconClasses = `inline-flex h-8 w-8 shrink-0 items-center justify-center rounded-lg ${props.toastType === "success" ? "bg-green-100 text-green-500" : "bg-red-100 text-red-500"}`;
    const progressClasses = `absolute bottom-0 left-0 h-1 w-full origin-left ${props.toastType === "success" ? "bg-green-500" : "bg-red-500"} animate-toast-progress group-hover:[animation-play-state:paused]`;

    // Let the fade-out transition finish before the parent unmounts the toast.
    useEffect(() => {
        if (!closing) return;
        const timer = setTimeout(onDismiss, CLOSE_DURATION_MS);
        return () => clearTimeout(timer);
    }, [closing, onDismiss]);

    return (
        <Toast
            duration={CLOSE_DURATION_MS}
            className={toastClasses}>
            <div className={iconClasses}>
                {
                    props.toastType === "success"
                        ? <HiCheckCircle className="h-5 w-5" />
                        : <HiXCircle className="h-5 w-5" />
                }
            </div>
            <div className="ml-3 text-sm font-normal">{message}</div>
            <ToastToggle onDismiss={() => setClosing(true)} />
            {/* The bar's animation doubles as the auto-close timer; hovering pauses both. */}
            <div
                className={progressClasses}
                style={{ animationDuration: `${timeout}ms` }}
                onAnimationEnd={() => setClosing(true)}
            />
        </Toast>
    );
}