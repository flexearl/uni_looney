import {RouterProvider, createBrowserRouter} from "react-router-dom"

import Logout from "../components/Logout/Logout"
import { useAuth } from "../components/authProvider"
import ProtectedRoute from "./ProtectedRoute"
import Login from "../components/Login/Login"
import AddPost from "../components/NewPost/NewPost"
import NewPost from "../components/NewPost/NewPost"
import Home from "../components/Home/Home"
import Account from "../components/Account/Account"

export default function Routes(){
    const token = useAuth()
    console.log("Token", token.token)
    const publicRoutes = [
        {
            path:"/",
            element: <Home/>
        },
        {
            path:"/all_posts",
            element: <div>All Posts</div>
        },


    ]

    const needAuthenticationRoutes = [
        {
            path:"/",
            element:<ProtectedRoute/>,
            children:[
                {
                    path: "/add_post",
                    element:<NewPost/>
                },
                {
                    path:"/change_post",
                    element:<div>Change Post</div>
                },
                {
                    path:"/logout",
                    element: <Logout/>
                },
                {
                    path:"/account",
                    element:<Account/>
                }
            ]
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