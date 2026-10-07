// Copyright (c) 2026, Farzan Hajian
// All rights reserved.
//
// This source code is licensed under the BSD 3-Clause license found in the
// LICENSE file in the root directory of this source tree.

import { Button, Label, Select } from "flowbite-react";

export default function Settings() {
    return (
        <form className="flex flex-col sm:mx-9" onSubmit={(e) => { e.preventDefault();/* todo */ }}>
            <div className="max-w-md">
                <div className="mb-2 block">
                    <Label className="lf-text-size" htmlFor="study-item-count">Study items per session</Label>
                </div>
                <Select id="study-item-count" required>
                    <option>5</option>
                    <option>10</option>
                    <option>15</option>
                    <option>20</option>
                </Select>

                <div className="flex justify-center mt-9">
                    <Button className="lf-primary-button" type="submit">Save</Button>
                </div>
            </div>
        </form>
    );
}