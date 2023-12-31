import { Navigate, Outlet } from "react-router-dom";
import {useAuth} from "../components/authProvider"

export default function ProtectedRoute(){
    const token = useAuth()
    console.log("In protected routes")
    if (token?.token=="" || token.token == null){
        console.log("Wrong place")
        return <Navigate to="/login"/>
    }
    return <Outlet/>
}