// 項目的入口 整個項目從這裡開始運行

// 運行必須要的兩個包
import React from 'react';
import ReactDOM from 'react-dom/client';

import store from "./store" // 导入根store
import {Provider} from 'react-redux'    // 导入Provider

import {RouterProvider} from "react-router-dom";

import router from './router'

// 將根組件渲染到root節點上
const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
      <Provider store={store}>
          {/*<App />*/}
          <RouterProvider router={router} />
      </Provider>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
