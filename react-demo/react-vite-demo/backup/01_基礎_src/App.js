// 項目的根組件

const message = 'this is message'

function getName() {
    return 'kanna'
}

const list = [
    {id: 1001, name: 'Vue'},
    {id: 1002, name: 'React'},
    {id: 1003, name: 'Angular'},
]

const isLogin = true

function App() {
  return (
      <div className="App">
          {'基礎'}
          <h2>Message</h2>
          <p>{message}</p>
          <p>{'str'}</p>
          <p>{getName()}</p>
          <p>{new Date().getDate()}</p>
          <div style={{
              color: 'red'
          }}>this is red.</div>

          { '列表渲染' }
          <ul>
              {list.map(item => <li key={item.id}>{ item.name }</li>)}
          </ul>

          {'條件渲染'}
          { isLogin && <p>yes</p> }
          { isLogin ? <p>yes</p>:<p>no</p> }

          {'複雜條件渲染'}
          {getArticleTem(articleType)}
      </div>
  );
}

// 定義文章類型
const articleType = 1

// 核心函數
function getArticleTem() {
    if (articleType === 0) {
        return <div>這是0圖文章</div>
    } else if (articleType === 1) {
        return <div>這是1圖文章</div>
    } else {
        return <div>這是2圖文章</div>
    }
}


export default App;
