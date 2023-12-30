import axios from "axios"
import{
    ReactNode,
    createContext,
    useContext,
    useEffect,
    useMemo,
    useState,
} from "react"

type Context =  {
    token: string | null;
    changeToken:(c: string) => void ; 
}

const AuthContext = createContext<Context>({
    token:"",
    changeToken:() => {},
})

interface AuthProviderProps{
    children : ReactNode
}

export default function AuthProvider(props: AuthProviderProps){
    const[token, setToken] = useState(localStorage.getItem("token"))
    function changeToken(newToken: string){
        setToken(newToken)
    }
    useEffect(() => {
        if (token){
            axios.defaults.headers.common['Authorization'] = `Bearer ${token}`
            localStorage.setItem("token", token)
        }else{
            delete axios.defaults.headers.common["Authorization"]
            localStorage.removeItem('token')
        }
    },[token])
    const contextValue = useMemo(
        () => ({
            token,
            changeToken,
        }),
        [token]
    )
    return(
        <AuthContext.Provider value={contextValue}>
            {props.children}
        </AuthContext.Provider>
    )
}
export const useAuth = () => {
    return useContext(AuthContext);
  };



