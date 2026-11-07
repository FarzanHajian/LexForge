// Copyright (c) 2026, Farzan Hajian
// All rights reserved.
//
// This source code is licensed under the BSD 3-Clause license found in the
// LICENSE file in the root directory of this source tree.

// import { createContext, useState } from 'react';
// import type { User } from '../models/user';


// export function UserProvider(props: { children: React.ReactNode }) {
//     const [user, setUser] = useState<User | null>(null);

//     const modifyUser = (studyItemsPerSession: number) => {
//         setUser((prev) => prev?.studyItemsPerSession === studyItemsPerSession ? prev : { ...prev, studyItemsPerSession });
//     };

//     return (
//         <UserContext value= {{ user, modifyUser }}>
//             { children }
//         </UserContext>
//     );
// }