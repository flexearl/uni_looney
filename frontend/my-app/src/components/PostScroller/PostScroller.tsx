import React, { useEffect, useState } from "react";
import axios from "axios";
import { Post } from "../NewPost/NewPost";

export default function PostScroller(){
    const[allPosts, setAllPosts] = useState()
    useEffect(() => {
        axios.get("http://localhost:8000/all_posts",  {  
        data:"Hi",
        headers : {
            'Content-Type':'application/json',
        }}).then(response => {
            setAllPosts(response.data)
        })
    }, [])

    
    return(
        <div className="PostScroller-Container">
            
            <button>

            </button>
        </div>
    )
}