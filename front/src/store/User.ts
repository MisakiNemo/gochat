import {defineStore} from "pinia";
import {Ref, ref} from "vue";
import {
    GetUserInfoResponseData,
    getUserInfoService,
    UserData,
    getFriendListService,
    GetFriendListResponseData
} from "@/api/user";

export const useUserStore = defineStore("User",
    () => {
        const userInfoData: Ref<UserData> = ref({
            id: -1,
            name: "",
            avatar: "https://i.loli.net/2021/04/14/nNly8EdXJ2aHYTe.jpg",
            status: false
        });

        const setUserInfo = (userInfo: UserData) => {
            console.log("set")
            userInfoData.value = userInfo;
        };

        const updateUserInfo = async () => {
            await getUserInfoService().then((res) => {
                let data: GetUserInfoResponseData = res.data;
                if (data.code === 1) {
                    console.log("获取用户信息失败");
                    console.log(data.msg)
                } else {
                    console.log(data)
                    setUserInfo(data.data);
                }

            }).catch((err) => {
                console.log("获取用户信息失败");
                console.log(err);
            })
        }

        const isUser =(id:number)=>{
            return userInfoData.value.id === id;
        }



        //好友信息板块
        const friendList:Ref<UserData[]> =ref([])

        const setFriendList = (list: UserData[]) => {
            friendList.value.splice(0, friendList.value.length);
            list.sort((a, b) => {
                if (a.status === b.status) {
                    return 0;
                } else if (a.status) {
                    return -1;
                } else {
                    return 1;
                }
            })
            list.forEach((item) => {
                friendList.value.push(item);
            })
        }

        const updateFriendList = async () => {
            await getFriendListService().then((res) => {
                let data: GetFriendListResponseData = res.data;
                if (data.code == 0) {
                    setFriendList(data.Rows);
                } else {
                    console.log("获取好友列表失败");
                }
                console.log(data);
            }).catch((err) => {
                console.log("获取好友列表失败");
                console.log(err);
            })
        }

        const getFriendData = (id: number) => {
            return friendList.value.find((item) => {
                return item.id === id;
            })
        }



        return {
            userInfoData,
            updateUserInfo,
            isUser,
            friendList,
            updateFriendList,
            getFriendData,


        }

    }
)
