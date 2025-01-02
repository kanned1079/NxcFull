// 項目的根組件
import "./App.css"

import React, {useEffect, useState} from "react";
import axios from "axios";

// 封装自定义Hook封装数据请求
const useGetComment = () => {
    const [comments, setComments] = useState([]);
    const getComments = async () => {
        let {data} = await axios.get("http://localhost:8080/api/v1/comments17");
        if (data.code === 200) setComments(data.comments)
        else console.log("err fetching comments")
    }
    return [comments, getComments]
}


function App() {
    let [comments, getComments] = useGetComment();
    useEffect(() => {
        getComments()
    }, []);
    return (
        <div className="App">
            {comments.map((comment) => (
                <div
                    key={comment.id}
                    style={{marginTop: "5px"}}
                >
                    <p>{comment.comment}</p>
                </div>
            ))}
        </div>
    );
}


export default App;
