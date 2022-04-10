import {useNavigate} from "react-router-dom";

export const RouteNavigation = () => {
    const navigate = useNavigate();

    const navigateTo = (url) => {
        navigate(url)
    }

    return(navigateTo)
}