import { createContext, useState, useEffect} from 'react'
import {fetchhome, fetchuser} from '..api/Api'
import Cookies from "universal-cookie";
import { Link } from 'react-router-dom';
import swal from "sweetalert2";


const Cookie = new Cookies();

export const Context = createContext()

export function ContextProvider(props) {
    const [username, setUsername] = useState("")
    const [password, setPassword] = useState("")

    


}