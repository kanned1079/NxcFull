const { app, BrowserWindow } = require('electron')
const path = require('path')

function createWindow() {
    const mainWindow = new BrowserWindow({
        width: 1200,
        height: 800,
        webPreferences: {
            nodeIntegration: true, // 允许渲染进程使用 Node.js API
            contextIsolation: false, // 关闭上下文隔离，允许渲染进程直接访问 Electron API
        },
    })

    // 加载 Vue 打包后的 index.html 文件
    mainWindow.loadFile(path.resolve(__dirname, './dist/index.html'))
    mainWindow.webContents.openDevTools()
}


app.whenReady().then(() => {
    createWindow()

    app.on('activate', function () {
        if (BrowserWindow.getAllWindows().length === 0) createWindow()
    })
})

app.on('window-all-closed', function () {
    if (process.platform !== 'darwin') app.quit()
})