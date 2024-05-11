import {MessagesData} from "@/api/message";


const sortMessageByTime = (messageList: MessagesData[]) => {
    messageList.sort((a, b) => {
        return a.create_time.getTime() - b.create_time.getTime();
    })
}

const handleTime = (time: Date) => {
    const date = new Date(time);
    const hour = date.getHours();
    const minute = date.getMinutes();
    return `${hour}:${minute}`
}

export {sortMessageByTime, handleTime}


