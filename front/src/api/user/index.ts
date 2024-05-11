import request from "@/utils/request";
import type * as User from "./type/user";
import {
    GetFriendListResponseData,
    GetUserInfoResponseData,
    LoginResponseData,
    RegisterResponseData,
    UserData
} from "./type/user";

export const registerService = (data: User.LoginOrRegisterRequestData) => {
  return request<RegisterResponseData>({
    url: "/user/register",
    method: "POST",
    data,
  })
}

export const loginService = (data: User.LoginOrRegisterRequestData) => {
    return request<LoginResponseData>({
        url: "/user/login",
        method: "POST",
        data,
    })

}


export const getUserInfoService = () => {
  return request<GetUserInfoResponseData>({
    url:"/user/info",
    method:"GET",
  })
}


export const getFriendListService = () => {
    return request<GetFriendListResponseData>({
        url:"/user/friends",
        method:"GET",
    })
}




export * from './type/user';
