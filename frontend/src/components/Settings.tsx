// Copyright (c) 2026, Farzan Hajian
// All rights reserved.
//
// This source code is licensed under the BSD 3-Clause license found in the
// LICENSE file in the root directory of this source tree.

import { useContext, useEffect, useState } from "react";
import { Button, Label, Select } from "flowbite-react";
import { useForm, type SubmitHandler } from "react-hook-form"
import { useApiToken } from "../utils/useApiToken";
import type { User } from "../models/user";
import { getMyUser, editMyUser } from "../services/userService";
import ErrorDisplay from "./ErrorDisplay";


function toError(err: unknown): Error {
    return err instanceof Error ? err : new Error(String(err));
}

export default function Settings() {
    const { register, handleSubmit, reset, formState: { errors } } = useForm<User>({ defaultValues: { studyItemsPerSession: 10 } });
    const getToken = useApiToken();
    const [error, setError] = useState<Error | null>(null);

    // useEffect(() => {
    //     const {user, modifyUser} = useContext<User|null>(UserContext);

    //     let ignore = false;

    //     (async () => {
    //         try {
    //             const token = await getToken();
    //             const me = await getMyUser(token);
    //             if (!ignore) reset(me);
    //         } catch (err) {
    //             if (!ignore) setError(toError(err));
    //         }
    //     })();

    //     return () => {
    //         ignore = true;
    //     };
    // }, [getToken, reset]);

    const onSubmit: SubmitHandler<User> = async (data) => {
        setError(null);
        try {
            const token = await getToken();
            await editMyUser(token, { studyItemsPerSession: data.studyItemsPerSession });
        } catch (err) {
            setError(toError(err));
        }
    };

    return (
        <form className="flex flex-col sm:mx-9" onSubmit={handleSubmit(onSubmit)}>
            <div className="max-w-md">
                {error && <ErrorDisplay error={error} />}

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
