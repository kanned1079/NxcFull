// 項目的根組件
function App() {

 
  return (
      <div className="App">
          {/*使用組件*/}
        <MyButton></MyButton>
          <MyButton2></MyButton2>
      </div>
  );
}

// 定義組件
function MyButton() {
    return (
        <button>click me</button>
    )
}

const MyButton2 = () => {
    return (
        <button>click me</button>
    )
}

export default App;
