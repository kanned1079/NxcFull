import {configureStore} from "@reduxjs/toolkit";

// 导入子模块
import counterReducer from "./modules/counterStore.js";

const store =  configureStore({
    reducer: {
        counter: counterReducer
    }
})

export default store