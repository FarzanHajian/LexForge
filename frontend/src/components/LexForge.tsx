// Copyright (c) 2026, Farzan Hajian
// All rights reserved.
//
// This source code is licensed under the BSD 3-Clause license found in the
// LICENSE file in the root directory of this source tree.

import clsx from 'clsx';

export default function LexForge(props: { textSize?: string }) {
    return (
        <span className={clsx('self-center whitespace-nowrap font-semibold', props.textSize && `text-${props.textSize}`)}>
            <span className="text-brand">Lex</span>
            <span className="text-accent">Forge</span>
        </span>
    );
}