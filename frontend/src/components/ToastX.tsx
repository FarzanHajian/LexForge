// Copyright (c) 2026, Farzan Hajian
// All rights reserved.
//
// This source code is licensed under the BSD 3-Clause license found in the
// LICENSE file in the root directory of this source tree.

import toast, { ToastBar, Toaster } from 'react-hot-toast';
import { HiOutlineXCircle } from 'react-icons/hi2';

export default function ToastX() {
    return (
        <Toaster position="top-right" toastOptions={{ className: 'group', duration: 5000, }}>
            {(t) => (
                <ToastBar toast={t}>
                    {({ icon, message }) => (
                        <>
                            {icon}
                            {message}
                            {t.type !== 'loading' && (
                                <>
                                    <button onClick={() => toast.dismiss(t.id)}><HiOutlineXCircle /></button>
                                    <div
                                        className={`absolute bottom-0 left-0 h-1 w-full origin-left ${t.type === "error" ? "bg-red-500" : "bg-green-500"} animate-toast-progress group-hover:[animation-play-state:paused]`}
                                        style={{ animationDuration: `${5000}ms` }}
                                        onAnimationEnd={() => toast.dismiss(t.id)}
                                    />
                                </>
                            )}
                        </>
                    )}
                </ToastBar>
            )}
        </Toaster>
    );
}