import React from "react";
import { useNavigate } from "react-router-dom";
import { useAuth } from "../authProvider";

export default function Logout(): React.JSX.Element{
    const token = useAuth()
    const navigate = useNavigate()
    
    function handleLogout(){
        token?.changeToken("")
        navigate("/", {replace: true})
    }

    setTimeout(() => {
        handleLogout()
    }, 3* 1000)
    return <>Logout Page</>
}