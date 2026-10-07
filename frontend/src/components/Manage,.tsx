// Copyright (c) 2026, Farzan Hajian
// All rights reserved.
//
// This source code is licensed under the BSD 3-Clause license found in the
// LICENSE file in the root directory of this source tree.

import { TabItem, Tabs } from "flowbite-react/components/Tabs";
import { HiWrenchScrewdriver, HiBookOpen } from "react-icons/hi2";
import Settings from "./Settings";

const theme = {
    tablist: {
        base: "justify-center",
        tabitem: {
            base: "lf-text-size"
        },
    }
};

export default function Manage() {
    return (
        <Tabs theme={theme} variant="underline">
            <TabItem active title="Notebooks" icon={HiBookOpen}>
                This is <span className="font-medium text-gray-800 dark:text-white">Profile tab's associated content</span>.
                Clicking another tab will toggle the visibility of this one for the next. The tab JavaScript swaps classes to
                control the content visibility and styling.
            </TabItem>
            <TabItem title="Settings" icon={HiWrenchScrewdriver}>
                <Settings />
            </TabItem>
        </Tabs>
    );
}
