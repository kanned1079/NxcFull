// 項目的根組件

function App() {

    const clickHandler = (e) => {
        console.log('button被點擊了', e)
    }

    const clickHandler1 = (name) => {
        console.log('button被點擊了', name)
    }

    const clickHandler2 = (name, e) => {
        console.log('button被點擊了', name, e)
    }

  return (
      <div className="App">
          <p>綁定事件</p>
          <button onClick={clickHandler}>Click me</button>
          <button onClick={() => clickHandler1('kanna')}>Click me</button>
          <p>如果值和事件參數都要 但是要注意位置</p>
          <button onClick={(e) => clickHandler2('kanna', e)}>Click me</button>
      </div>
  );
}


export default App;
