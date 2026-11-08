// Copyright (c) 2026, Farzan Hajian
// All rights reserved.
//
// This source code is licensed under the BSD 3-Clause license found in the
// LICENSE file in the root directory of this source tree.

import { useContext } from "react";
import { Button, Label, Select } from "flowbite-react";
import { useForm, type SubmitHandler } from "react-hook-form"
import { useApiToken } from "../utils/useApiToken";
import type { User } from "../models/user";
import { editMyUser } from "../services/userService";
import { UserContext } from "../utils/userContext";
import { toError, showError, showInfo } from "../utils/utils";

export default function Settings() {
    const { user, setUser } = useContext(UserContext);
    const { register, handleSubmit, formState: { errors } } = useForm<User>({ defaultValues: { studyItemsPerSession: 10 }, values: user ?? undefined, });
    const getToken = useApiToken();

    const onSubmit: SubmitHandler<User> = async (data) => {
        try {
            const token = await getToken();
            const updated = await editMyUser(token, { studyItemsPerSession: data.studyItemsPerSession });
            setUser(updated);
            showInfo("data-saved");
            showError(new Error("An unexpected error occured."));
        } catch (err) {
            showError(toError(err));
        }
    };

    return (
        <form className="flex flex-col sm:mx-9" onSubmit={handleSubmit(onSubmit)}>
            <div className="max-w-md">
                <div className="mb-2 block">
                    <Label className="lf-text-size" htmlFor="study-item-count">Study items per session</Label>
                </div>
                <Select id="study-item-count"
                    aria-invalid={errors.studyItemsPerSession ? "true" : "false"}
                    {...register("studyItemsPerSession", { required: true, valueAsNumber: true })} >
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
