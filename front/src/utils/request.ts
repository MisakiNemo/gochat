import axios from 'axios'
import {Token} from "@/api/user";




export const setToken = (token: Token) => {
    localStorage.setItem('token', token.access_token);
}

export const getToken = () => {
    return localStorage.getItem('token')||'';
}

const instance = axios.create({
    baseURL: 'http://localhost:8080/v1',
    timeout: 1000,
    headers: {
        'Content-Type': 'application/json',
        'Accept': 'application/json',
         'Authorization': getToken(),
    },
})


export default instance;