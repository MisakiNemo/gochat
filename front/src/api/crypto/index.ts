import request from "@/utils/request";
import {getPublicKeyResponseData} from "@/api/crypto/type";


export const getPublicKey = () => {
    return request<getPublicKeyResponseData>({
        url: '/user/publickey',
        method: 'post',
        data: {}
    })
}

export const updateAESKey = (AESkey: string) => {
    return request({
        url: '/user/aeskey',
        method: 'post',
        data: {
            aes: AESkey
        }
    })
}