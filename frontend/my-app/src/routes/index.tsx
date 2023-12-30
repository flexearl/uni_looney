import {RouterProvider, createBrowserRouter} from "react-router-dom"

import Logout from "../components/Logout/Logout"
import { useAuth } from "../components/authProvider"
import ProtectedRoute from "./ProtectedRoute"
import Login from "../components/Login/Login"
import AddPost from "../components/AddPost/AddPost"

export default function Routes(){
    const token = useAuth()
    console.log("Token", token.token)
    const publicRoutes = [
        {
            path:"/home",
            element: <div>Home</div>
        },
        {
            path:"/all_posts",
            element: <div>All Posts</div>
        },
        {
            path:"/",
            element:<div>Root</div>
        }

    ]

    const needAuthenticationRoutes = [
        {
            path: "/add_post",
            element:<AddPost/>
        },
        {
            path:"/change_post",
            element:<div>Change Post</div>
        },
        {
            path:"/logout",
            element: <Logout/>
        }
    ]

    const unauthenticatedRoutesOnly = [
        {
            path:"/login",
            element: <Login/>
        }
    ]

    const router = createBrowserRouter([
        ...publicRoutes,
        ...(token.token == "" || token.token == null ? unauthenticatedRoutesOnly: []),
        ... needAuthenticationRoutes,

    ])

    return <RouterProvider router={router}/>

}