import axios from "axios";

const createChannelStore = (set) => {
    return {
        channelList: [],
        handleFetchChannels: async () => {
            let {data} = await axios.get('https://geek.itheima.net/v1_0/channels')
            console.log(data)
            if (data.message === 'OK')
                set({
                    channelList: data.data.channels
                })
        },
        handleClearChannels:  () => {
            set({
                channelList: [],
            })
        }
    }
}

export default createChannelStore