import {useRef, useEffect} from "react";


function App() {
    const inputDomRef = useRef<HTMLInputElement>(null)
    const intervalId = useRef<ReturnType<typeof setInterval> | undefined>(undefined)
    useEffect(() => {
        inputDomRef.current?.focus()
        intervalId.current = setInterval(() => {
            console.log('定时器')
        }, 1000)

        return () => {
            clearInterval(intervalId.current)
        }

    }, []);

    return (
        <>
            <h2>This is app.</h2>
            <input type="text" ref={inputDomRef} />
            <button onClick={() => clearInterval(intervalId.current)}>stop interval</button>
        </>
    )
}



export default App
