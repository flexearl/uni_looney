import React, { useState } from 'react'
import axios from 'axios'
import { useAuth } from '../authProvider'

interface Post {
    title: string,
    body: string,
    description:string,
}

export default function AddPost(){
    
    const [post, setPost] = useState<Post>({title:"", body:"", description:""})

    const {token} = useAuth()

    function AddPost(event: React.MouseEvent<HTMLElement>){
        event.preventDefault()
        let newPost: Post = {body:"dadjd", description:"fjfjfj", title:"ffkfkn"}
        
        const headers = {
            "Content-Type": "application/json",
            "Authorization": `Bearer ${token}`
        }
        
        
        axios.post("http://localhost:8000/add_post", newPost,{
            headers: headers
        }).then((response) => {
            console.log(response.data)
            
        })
    }

    function handleInputChange(event: React.FormEvent<HTMLInputElement>){
        const {name, value} = event.currentTarget
        let newTitle = event.currentTarget.value
        setPost((prevState:Post) => {
            
            return {
            ...prevState,
            [name] : value    
            }
        })
    }

    


    return(
        <div>
            <input placeholder='Title' value={post?.title} name="title" onChange={handleInputChange}/>
            <input placeholder='Description' value={post.description} name='description' onChange={handleInputChange}/>
            <input placeholder='Body' value={post.body} name='body' onChange={handleInputChange}/>
            
        </div>    
    )
}