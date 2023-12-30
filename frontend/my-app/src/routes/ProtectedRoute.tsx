import { Navigate, Outlet } from "react-router-dom";
import {useAuth} from "../components/authProvider"

export default function ProtectedRoute(){
    const token = useAuth()
    if (token?.token==""){
        return <Navigate to="/login"/>
    }
    return <Outlet/>
}