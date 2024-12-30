// 項目的入口 整個項目從這裡開始運行

// 運行必須要的兩個包
import React from 'react';
import ReactDOM from 'react-dom/client';

import App from './App';

// 將根組件渲染到root節點上
const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  // <React.StrictMode>
    <App />
  // </React.StrictMode>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
