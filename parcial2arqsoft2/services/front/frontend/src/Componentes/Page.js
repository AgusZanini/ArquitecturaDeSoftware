import React from "react";
import { Routes, Route } from "react-router-dom";
import { Main } from "./Main/index";

import { InsertProperty} from "./InsertProperty/funtion";
import { InsertMessage } from "./message/index";
import Login from "./Login/login";


export const Page = () => {
    return (
        <section>
            <Routes>
                <Route path="/" exact element={<Login/>} />
                <Route path="/home" exact element={<Main/>} />
                <Route path="/properties" exact element={<InsertProperty/>} />
                <Route path="/message" exact element={<InsertMessage/>} />
            </Routes>
        </section>
    )
}