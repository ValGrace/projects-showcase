import {useState} from "react";

const UsersForm = () => {
  const [title, setTitle] = useState('')
 
  const [livelink, setLivelink] = useState('')
  const [gitlink, setGitlink] = useState('')
 
 
 
    return (
        <form method="post" action="/users/" encType="multipart/form-data" id="project-form">
          
          <input type="text" name="username" value={title} onChange={e => setTitle(e.target.value)} placeholder="username"/> 
          <input type="text" name="password" value={livelink} onChange={e => setLivelink(e.target.value)} placeholder="password"/>
          <input type="text" name="level" value={gitlink} onChange={e => setGitlink(e.target.value)} placeholder="level"/>  
          
          
          
          <input type="submit" value="register" id="register-btn"/>
        </form>
    )
}

export default UsersForm
