export interface LoginOrRegisterRequestData {
    name: string,
    password: string,
}


export interface UserData {
    id: number,
    name: string,
    avatar: string,
    status: boolean
}

export interface Token {
    access_token: string
    access_expire : number
}


export type GetUserInfoResponseData = Api<UserData>

export type RegisterResponseData = Api<Token>

export type LoginResponseData = Api<Token>

export type GetFriendListResponseData = ApiList<UserData>






