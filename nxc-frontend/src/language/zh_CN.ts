export default {
    commonHeader: {
        menuTxt: '菜單',
        userData: '用戶資料',
        editUserData: '編輯用戶資料',
        logout: '退出登陸',
    },
    commonAside: {
        admin: {},
        user: {
            dashboard: '儀表板',
            document: '使用文檔',
            subscribe: '訂閱',
            purchase: '購買訂閱',
            surplus: '我的密鑰',
            fiance: '財務',
            myOrder: '我的訂單',
            myInvite: '我的邀請',
            user: '用戶',
            profile: '個人中心',
            support: '我的工單',
            activateLog: '激活紀錄'
        }
    },
    userLogin: {
        loginToContinue: '登入以繼續',
        email: '郵件地址',
        password: '密碼',
        haveNoAccount: '還沒有您的帳戶？',
        login: '登入',
        reg: '立即註冊',
        otherMethods: '或使用其他方式繼續',
        github: '以Github帳戶繼續',
        apple: '以Apple帳戶繼續',
        google: '以Google帳戶繼續',
        backHomePage: '回到首頁',

    },
    userSummary: {
        myPlan: '我的訂閱',
        shortcut: '捷徑',
        timeLeft: '订阅有效，将在 {msg} 过期。',
        toPurchase: '購買訂閱',
        tutorial: {
            title: '查看教程',
            content: '學習如何使用 {name}',
        },
        checkKey: {
            title: '查看密鑰',
            content: '快速將密鑰導入對應客戶端進行使用',
        },
        renewPlan: {
            title: '續費訂閱',
            content: '對您當前的訂閱進行續費',
        },
        support: {
            title: '遇到問題',
            content: '遇到問題可以通過工單與我們的人機溝通',
        }
    },
    userDocument: {
        searchPlaceholder: '請輸入要搜索的內容（模糊搜索）',
        searchBtn: '搜索',
    },
    newPurchase: {
        headerPlaceholder: '選擇最適合您的計畫',
        purchaseBtn: '訂購',
        monthPay: '月付',
        moreMethod: '更多選擇',
    },
    newSettlement: {
        err: '錯誤',
        monthPay: '月付',
        quarterPay: '季付',
        halfYearPay: '半年付',
        yearPay: '年付',
        payCycle: '付款週期',
        couponPlaceholder: '有優惠券 ？',
        verifyBtn: '驗證',
        orderTotalTitle: '訂單總額',
        total: '總計',
        submitOrder: '提交訂單',
        coupon: '優惠券',
        notify: {
            passTitle: '驗證通過',
            couponVerified: '優惠券有效',
            couponInvalid: "優惠券無效",
            couponIsNull: '輸入的優惠券碼不能為空',
        }
    },
    // pass
    userProfile: {
        myWallet: '我的錢包',
        walletSub: '帳戶餘額（僅用於消費）',
        alertPwd: '修改密鑰',
        oldPwd: '舊密鑰',
        oldPwdSub: '請輸入舊密碼',
        newPwd: '新密鑰',
        newPwdSub: '請輸入新的密鑰',
        newPwdAgain: '確認密鑰',
        newPwdAgainSub: '請再輸入一遍新的密鑰',
        saveBtn: '保存',
        notify: '通知',
        enableNotify: '啟用到期通知',
        deleteAccount: '註銷帳戶',
        deleteAccountSub: '您的帳戶將被標記為刪除，如果需要重新使用我們的服務，請重新註冊',
        deleteBtn: '註銷我的帳戶',
        oldPwdVerifiedFailure: '舊密碼驗證失敗',
        alertFailure: '密鑰修改失敗',
        alertSuccess: '修改成功',
        alertSuccessSub: '請使用新密碼登錄',
        errorPwdFormat: '密碼格式錯誤',
        pwdNotMatch: '兩次密碼輸入不一致',
        oldPwdNotNull: '舊密碼不能為空',
    },
    tickets: {
        userTickets: '歷史工單',
        addTicket: '新建工單',
    },
    userKeys: {
        myKeys: '我的密鑰'
    },
    welcome: {
        A: {
            aboutUs: '關於我們',
            pricing: '定價',
            login: '登錄',
            register: '註冊帳號',
            welcomeTo: '歡迎來到',
            welcomeToSub: '“穿过县境上长长的隧道，便是雪国。夜空下，大地一片莹白，火车在信号所前停下来。”在这里川端康成用几近吝啬的简洁文字，拉开了《雪国》的序幕。',
            startingUse: '開始使用',
            whyUs: '為什麼選擇我們',
            whyUsSub: '“穿过县境上长长的隧道，便是雪国。夜空下，大地一片莹白，火车在信号所前停下来。”在这里川端康成用几近吝啬的简洁文字，拉开了《雪国》的序幕。',
            browseSafe:' 瀏覽安全',
            browseSafeSub: '優秀的防火牆過濾系統能有效防禦網路釣魚和惡意網站',
            encrypt: '端到端加密',
            encryptSub: '雙向 SSL 和端對端加密保護您的隱私安全，即使是服務器也無法讀取您的信息',
            mgr: '高效管理',
            mgrSub: '一個用戶介面管理所有密鑰，管理功能完善豐富，無須擔心訂閱洩露問題',
            fast: '方便快捷',
            fastSub: '提供完整的 API 文檔供 WebApp 或是嵌入到第三方軟件中',
            fastLink: '快速連結',
            subscribeUs: '關注我們',
            filingsCode: '備案號 {code}',
        }
    }
}