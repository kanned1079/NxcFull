import "./CommentItem.css"

function CommentItem(props) {

    return (
        <div
            className="bg"
        >
            <p className="id">评论Id: {props.id}</p>
            <p className="comment">评论内容: {props.comment}</p>
            <hr/>
        </div>
    )
}

export default CommentItem