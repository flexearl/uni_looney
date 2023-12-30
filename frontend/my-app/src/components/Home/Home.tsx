import React, { useState } from "react";
import Login from "../Login/Login";

export default function Home(){
    const [token, setToken] = useState("")
    return(
        <div className="Home-Container">
            <header>
                <h1>UNILOONEY</h1>
                <div className="Tag-Box">
                    <h3>Student to student support</h3>
                </div>
            </header>
            <div className="Post-Container">
                <h3>Posts</h3>
    
            </div>
        </div>
    )
}