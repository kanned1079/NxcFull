import {configureStore} from "@reduxjs/toolkit";

// 导入子模块
import counterReducer from "./modules/counterStore.js";
import channelReducer from "./modules/channelStore.js";

const store =  configureStore({
    reducer: {
        counter: counterReducer,
        channel: channelReducer,
    }
})

export default store