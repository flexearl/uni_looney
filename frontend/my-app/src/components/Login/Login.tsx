import React, { useContext, useState } from "react";
import axios from 'axios'
import { useAuth } from "../authProvider";
import { useNavigate } from "react-router-dom";

export interface LoginProps{
    setToken: Function,
    token :string
}

interface User {
    username: string,
    password: string,
}



export default function Login(): React.JSX.Element{
    const navigate = useNavigate()
    const[username, setUsername] = useState("")
    const [password, setPassword] = useState("")
    const {token, changeToken} = useAuth()

    function handleUsernameChange(event: React.FormEvent<HTMLInputElement>){
        let newUsername = event.currentTarget.value
        setUsername(newUsername)
    }
    
    function handlePasswordChange(event: React.FormEvent<HTMLInputElement>){
        let newPassword = event.currentTarget.value 
        setPassword(newPassword)
    }

    async function LoginUser(user:User){
        return fetch("http://localhost:8000/login", {
            method:"POST",
            headers : {
                'Content-Type':'application/json'
            },
            body: JSON.stringify(user)
        }).then(data => console.log(data))
    }
    
    async function handleSubmit(event: React.MouseEvent<HTMLElement>){

        event.preventDefault()
        axios.post("http://localhost:8000/login", {
            username,
            password
        }).then((response) => {
            console.log(response.data)
            changeToken(response.data)
        })
        navigate("/", {replace: true})
    }

    



    return(
        <div className="Login-Container">
            <form>
                <input placeholder="username" value={username} onChange={handleUsernameChange}/>
                <input placeholder="password" value={password} onChange={handlePasswordChange}/>
                <button onClick={handleSubmit}>Submit</button>
                
            </form>
        </div>
    )
}
