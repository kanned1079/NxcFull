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
            subscription: '訂閱',
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
    adminViews: {
        summary: {
            systemConfig: '系統設置',
            paymentConfig: '支付設置',
            planMgr: '訂閱管理',
            userMgr: '用戶管理',
            orderMgr: '訂單管理',
            keyMgr: '密鑰管理',
            incomeText: '昨日收入 / 當月收入',
        },
        docMgr: {
            docId: '#',
            isShow: '是否顯示',
            sortAs: '排序',
            lang: '語言',
            category: '分類',
            title: '標題',
            createdAt: '創建時間',
            updatedAt: '更新時間',
            operate: '操作',
            edit: '編輯',
            delete: '刪除',

        },
        userMgr: {
            userManager: "用戶管理",
            enterEmail: "請輸入郵箱",
            enterValidEmail: "請輸入正確的郵箱格式",
            enterPassword: "請輸入密碼",
            passwordMinLength: "密碼長度不能少於 6 位",
            passwordMaxLength: "密碼長度不能超過 20 位",
            passwordStrength: "密碼必須包含字母、數字和特殊字符",
            validationSuccess: "驗證通過",
            validationFailed: "表單驗證失敗，請檢查輸入項",
            editProfile: "編輯資料",
            banUser: "封禁用戶",
            unbanUser: "解封用戶",
            noRecord: "無記錄",
            normal: "正常",
            banned: "封禁",
            nullContent: "NULL",
            balance: "餘額",
            orderCount: "訂單數量",
            planCount: "計劃數",
            actions: "操作",
            updateSuccess: "更新成功",
            addUserSuccess: "添加新用戶成功",
            unknownError: "未知錯誤",
            dataCountOptions10: "10條數據/頁",
            dataCountOptions20: "20條數據/頁",
            dataCountOptions50: "50條數據/頁",
            dataCountOptions100: "100條數據/頁",
            email: "郵箱",
            registerDate: "註冊日期",
            isAdmin: "管理員",
            isStaff: "員工",
            accountStatus: "帳戶狀態",
            inviteCode: "邀請碼",
            query: "查詢",
            reset: "重置",
            addNewUser: "新增用戶",
            searchUser: "查詢用戶",
            cancel: "取消",
            submit: "提交",
            userEmail: "用戶郵箱",
            inputUserEmail: "輸入用戶郵箱（模糊搜索）",
            inputEmail: "輸入郵箱",
            password: "密碼",
            inputPassword: "輸入密碼",
            // inviteCode: "邀請碼",
            // isAdmin: "管理員",
            // isStaff: "員工",
            // balance: "餘額",
            // email: "郵箱",
            editUser: "編輯用戶",
            no: "否",
            yes: "是"
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
    userRegister: {
        backHomePage: '回到首頁',
        newAccount: '準備您的新帳戶',
        email: '郵箱地址',
        verifyCode: '郵箱驗證碼',
        invalidEmailFormat: '郵箱格式不合法',
        sendVerifyCode: '發送郵箱驗證碼',
        sendSuccess: '郵件發送成功，請查看郵箱。',
        pwd: '密碼',
        pwdAgain: '確認密碼',
        inviteCode: '邀請碼（選填）',
        agreement: '我已閱讀並同意',
        terminalUserAgreement: '用戶條款',
        reg: '註冊',
        // regSuccess: '註冊成功',
        infoIncomplete: '信息不完整',
        pwdIncorrect: '密碼不一致',
        regSuccess: '註冊成功返回登錄',
        regFailure: '註冊失敗',
        success: '成功',
        failure: '失敗',
        unknowErr: '未知錯誤',
        verifyCodeErr: '驗證碼錯誤',

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
        toTopUp: '去充值',
        deleteMyTitle: '警告',
        deleteMyContent: '确定删除您的账户吗？如需要使用服务请重新注册。',
        deleteMyPositiveText: '确认删除',
        deleteMyNegativeText: '取消',
        deletedSuccessMsg: '账户已删除，后会有期。',
        deleteErrOccur: '遇到错误',

    },
    // tickets: {
    //     userTickets: '歷史工單',
    //     addTicket: '新建工單',
    // },
    userKeys: {
        myKeys: '我的密鑰',
        noKeys: '您還沒有有效的購買紀錄',
        keyDetail: '密鑰細節',
        keyId: '密鑰ID',
        orderId: '链接订单ID',
        clientId: '激活客戶端ID',
        active: '活躍',
        inActive: '不可用',
        valid: '密鑰有效',
        invalid: '密鑰不可用',
        isUsed: '已激活使用',
        noUsed: '待使用',
        releaseData: '密钥生成日期',
        expirationData: '到期日期',
        none: '無',
        authorizationFor: '密鑰授權給',
        hoverClickMention: '點擊可以複製到剪貼板',
        copiedSuccessMessage: '密鑰複製成功，請參照文檔說明繼續操作。',
        copyFailure: '複製失敗',
        hoverCopiedSuccessMention: '複製成功'
    },
    userOrders: {
        myOrders: '我的訂單',
        orderId: '# 訂單號',
        planCycle: '訂閱週期',
        orderPrice: '訂單金額',
        orderStatus: '訂單狀態',
        createdAt: '創建時間',
        operate: '操作',
        showDetail: '訂單細節',
        cancelOrder: '取消訂單',
        canceled: '已取消',
        period: {
            monthPay: '月付',
            quarterPay: '季付',
            halfYearPay: '半年付',
            yearPay: '年付',
        }
    },
    userConfirmOrder: {
        switchPlan: '切換訂閱',
        cancelOrder: '取消訂單',
        yourPrice: '您的價格',
        couponOffset: '優惠券抵折',
        price: '價格',
        submit: '提交',
        monthPay: '月付',
        quarterPay: '季付',
        halfYearPay: '半年付',
        yearPay: '年付',
        goodInfo: '商品信息',
        cycleOrType: '週期/類型',
        orderNumber: '訂單號',
        createdAt: '創建日期',
        orderExpired: '訂單已超時',
        paySuccessfully: '支付成功，感謝您的支持。',
        balanceNotEnough: '您的餘額不足請先充值，該訂單保留五分鐘。',
        orderErrorOccur: '訂單遇到錯誤',
        orderCancelled: '訂單已取消',
    },
    paymentResultParts: {
        goodInfoView: {
            goodInfo: '商品信息',
        },
        orderInfoView: {
            orderInfo: '訂單信息',
        },
    },
    orderPartUniversal: {
        period: {
            monthPay: '月付',
            quarterPay: '季付',
            halfYearPay: '半年付',
            yearPay: '年付',
        },
        orderDataHex: {
            goodInfo: '商品信息',
            orderInfo: '訂單信息',
            cycleOrType: '週期/類型',
            orderNumber: '訂單號',
            createdAt: '創建日期',
            amount: '支付金額',
            paidAt: '支付時間',
        }
    },
    orderDetail: {
        // goodInfo: '商品信息',
        // cycleOrType: '週期/類型',
        // monthPay: '月付',
        // quarterPay: '季付',
        // halfYearPay: '半年付',
        // yearPay: '年付',
        // orderNumber: '訂單號',
        // createdAt: '創建日期',
        // amount: '支付金額',
        // paidAt: '支付時間',

        finished: '已完成',
        finishedAndSuccessDescription: '訂單已成功支付並開通',
        useManual: '查看使用教程',

        payPending: '尚未支付',
        pendingDescription: '訂單暫時保留，可以點擊下面的按鈕以繼續支付。',
        toPay: '去支付',

        outDate: '已失效',
        outDateDescription: '由於您取消了訂單或未在指定時間內完成支付，因此該訂單已被取消，您可以重新選取訂閱。',
        chooseNewPlan: '選擇新的訂閱計畫',

    },
    userInvite: {
        myInvite: '我的邀請',
        unit: '人數',
        inviteCodeMgr: '您的邀請碼',
        generateInviteCode: '生成隨機邀請碼',
    },
    userTickets: {
        ticketId: '#',
        ticketSubject: '工單主題',
        ticketUrgency: '工單級別',
        ticketContent: '工單內容',
        ticketUrgencyLevel: {
            low: '低',
            med: '中',
            high: '高',
        },
        ticketStatus: '工單狀態',
        ticketCreatedAt: '創建時間',
        lastResponse: '最後回覆',
        operate: '操作',
        checkTicket: '查看工單',
        closeTicket: '關閉工單',
        userTickets: '歷史工單',
        addTicket: '新建工單',
        ticketActive: '未關閉',
        ticketInActive: '已關閉',
        form: {
            ticketSubjectDescription: '請輸入工單主題',
            ticketUrgencyDescription: '請選擇工單緊急程度',
            ticketBody: '請輸入你想要解決的問題，盡量全面。',
            ticketNotFinished: '請補全工單的信息再試'
        },
        cancel: '取消',
        submit: '提交',
        commitNewTicketSuccess: '提交新的工單成功',
        commitNewTicketFailure: '提交新的工單錯誤',
        noReply: '還沒有回覆',
        noTickets: '您還沒有提交過工單',
        ticketCloseSuccess: '工單關閉成功',
        ticketCloseFailure: '工單關閉失敗'
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
            browseSafe: ' 瀏覽安全',
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
    },
    adminTicket: {
        ticketMgr: '工單管理',

    }
}