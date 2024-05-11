export interface MessagesData {
    sender: number,
    receiver: number,
    content: string,
    create_time: Date,
    type: number
}

export interface ChatRecord{
    friendId: number
    MessagesData: MessagesData[]
}

