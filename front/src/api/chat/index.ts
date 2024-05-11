import request from "@/utils/request";



export const handleChat = () => {
    return request({
        url: '/user/test',
        method: 'post',
        data: {
        }
    })
}
