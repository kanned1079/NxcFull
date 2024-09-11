import {type NotificationType, useNotification} from 'naive-ui'

let makeNotify = (type: NotificationType, title: string, msg: string) => {
    let notification = useNotification()
    notification[type]({
        content: title,
        meta: msg,
        duration: 2500,
        keepAliveOnHover: true
    })
}

export {
    makeNotify
}