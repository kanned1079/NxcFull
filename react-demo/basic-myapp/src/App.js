// 項目的根組件
import "./App.css"
import {useState} from "react";

const list = [
    {
        rpid: 3,    // 評論id
        user: { // 用戶信息
            uid: '298236035623',
            avatar: '', // 用戶的頭像，這裡暫時為空
            uname: 'kanna', // 用戶名
        },
        content: '哎唷真不錯',   // 評論內容
        ctime: '2024-10-23 21:09:52',   // 評論時間
        like: 81,   // 點贊數
    },
    {
        rpid: 4,
        user: {
            uid: '398123047536',
            avatar: '',
            uname: 'miku',
        },
        content: '這個視頻好有趣！',
        ctime: '2024-10-23 22:10:15',
        like: 150,
    },
    {
        rpid: 5,
        user: {
            uid: '482019374651',
            avatar: '',
            uname: 'taro',
        },
        content: '不錯啦，不過還有提升空間',
        ctime: '2024-10-23 22:45:30',
        like: 55,
    },
    {
        rpid: 6,
        user: {
            uid: '592304869457',
            avatar: '',
            uname: 'ayumi',
        },
        content: '完全不懂為什麼大家這麼喜歡',
        ctime: '2024-10-23 23:05:42',
        like: 32,
    },
    {
        rpid: 7,
        user: {
            uid: '672840235862',
            avatar: '',
            uname: 'yuki',
        },
        content: '看完覺得很感動',
        ctime: '2024-10-23 23:50:11',
        like: 120,
    },
];

const user = {
    uid: '92364592674902',
    avatar: '',
    uname: 'sasa',
}


function App() {
    // 1.渲染評論列表
    const [commentList, setCommentList] = useState(list);
    const [commentText, setCommentText] = useState('')

    const sendMyComment = () => {
        console.log('send')
        setCommentList([...commentList, {
            rpid: commentList[commentList.length - 1].rpid + 100,
            user,
            content: commentText,
            ctime: new Date().toString(),
            like: 0,
        }])
    }

    const deleteMyComment = (rpid) => {
        console.log('刪除', rpid)
        setCommentList(commentList.filter(item => item.rpid !== rpid))
    }

    const sortAsLatest = () => {
        console.log('sortAsLatest')
    }

    const sortAsLike = () => {
        console.log('sortAsLike');
        const sortedList = [...commentList].sort((a, b) => b.like - a.like);
        setCommentList(sortedList);
    };

    return (
        <div className="App">
            <div className="comment-part">
                <div className="top-header">
                    <h4 style={{marginRight: 40}}>所有評論</h4>
                    <button className="transparent-btn" onClick={sortAsLatest}>按照時間</button>
                    <button className="transparent-btn" onClick={sortAsLike}>按照熱度</button>
                </div>
                {commentList.map(item => (
                    <div key={item.rpid} className="commentRoot">
                        <div className="l-content">
                            <div className="avatar"></div>
                        </div>
                        <div className="r-content">
                            <p style={{fontSize: '0.9rem', opacity: 0.7}}>{item.user.uname}</p>
                            <p style={{fontSize: '1rem', margin: '10px 0'}}>{item.content}</p>
                            <div className="func-part">
                                <p className="right-gap" style={{fontSize: '0.9rem', opacity: 0.7, marginRight: '10px'}}>{item.ctime}</p>
                                <p className="right-gap" style={{fontSize: '0.9rem', opacity: 0.7}}>點讚數:{item.like}</p>
                                {(user.uid === item.user.uid) && <button className="delete-btn" onClick={() => deleteMyComment(item.rpid)}>刪除</button>}
                            </div>
                        </div>
                    </div>
                ))}
                <hr style={{margin: '20px 0', opacity: 0.4}}/>
                <h4>添加評論</h4>
                <div className="my-push-comment-root">
                    <input
                        value={commentText}
                        onChange={(e) => setCommentText(e.target.value)}
                        className="input-text"
                        type="text"
                    />
                    <button className="submit-btn" onClick={sendMyComment}>評論</button>
                </div>
            </div>
        </div>
    );
}

export default App;
