import {Routes, Route} from 'react-router-dom'
import HomePage from "./HomePage";
import ProjectsForm from "./projects";
import UsersForm from "./users";
const Routing = () => {
    return (
        <Routes>
            <Route exact path="/" element={<HomePage />} />
            <Route path="/projects" element={<ProjectsForm />} />
            <Route path="/users" element={<UsersForm />} />

        </Routes>
    )
}

export default Routing