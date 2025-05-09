export default {
    "commonHeader": {
        "menuTxt": "菜單",
        "userData": "用戶資料",
        "editUserData": "編輯用戶資料",
        "logout": "退出登陸"
    },
    "commonAside": {
        "admin": {
            "dashboard": "儀表板",
            "queueMonit": "服務端監控",
            "settings": "設置",
            "systemConfig": "系統設置",
            "paymentConfig": "支付設置",
            "themeConfig": "主題設置",
            "server": "服務器",
            "privilege": "權限組管理",
            "finance": "財務",
            "subscription": "訂閱管理",
            "coupon": "優惠券管理",
            "order": "訂單管理",
            "activate": "激活紀錄",
            "key": "密鑰管理",
            "user": "用戶",
            "userMgr": "用戶管理",
            "notice": "公告管理",
            "ticket": "工單管理",
            "doc": "知識庫管理"
        },
        "user": {
            "dashboard": "儀表板",
            "document": "使用文檔",
            "app": "APP下載",
            "subscription": "訂閱",
            "purchase": "購買訂閱",
            "surplus": "我的密鑰",
            "fiance": "財務",
            "topUp": "儲值",
            "myOrder": "我的訂單",
            "myInvite": "我的邀請",
            "user": "用戶",
            "profile": "個人中心",
            "support": "我的工單",
            "activateLog": "激活紀錄"
        }
    },
    "adminViews": {
        "common": {
            "fetchDataSuccess": "獲取數據成功",
            "fetchDataFailure": "獲取數據失敗請稍後再試",
            "addSuccess": "添加成功",
            "addFailure": "添加失敗請稍後嘗試",
            "updateSuccess": "修改成功",
            "updateFailure": "修改失敗請稍後再試",
            "deleteSuccess": "刪除成功",
            "deleteFailure": "刪除失敗請稍後再試",
            "confirm": "確認",
            "cancel": "取消",
            "success": "操作成功",
            "failure": "操作失敗",
            "NotAllowed": "非法格式",
            "checkForm": "請檢查表單",
            "unknownErr": "未知錯誤",
            "dialog": { "delete": "您確認要刪除嗎" },
            "yes": "是",
            "no": "否"
        },
        "login": {
            "secureCard": {
                "title": "安全檢查",
                "securePath": "安全路徑",
                "hint": "為保障系統安全，您需輸入安全路徑方可進入管理員登錄頁面。 請在下方輸入框中輸入安全路徑，成功驗證安全路徑後您可選擇選擇保存，以便後續快捷登錄。",
                "placeholder": "請輸入安全路徑",
                "checkBtn": "檢查",
                "rememberPath": "記住安全路徑"
            },
            "card": {
                "back": "返回首頁",
                "toAdminCenter": "登錄到管理中心",
                "emailAddr": "管理員郵箱",
                "emailInputArea": { "title": "管理員郵箱", "placeholder": "郵箱地址" },
                "passwordInputArea": { "title": "管理員密碼", "placeholder": "密碼" },
                "login": "登入",
                "forgetPassword": "忘記密碼",
                "formNotPassed": "表單驗證不通過"
            },
            "mention": {
                "title": "提示",
                "description": "此頁面是管理員頁面，只有管理員才能訪問，如果您沒有管理員權限或意外來到此處，請點擊下面的鏈接返回主頁。"
            },
            "message": {
                "passwordErr": "密碼不正確",
                "adminNotExist": "管理員不存在",
                "noPrivilege": "無權限訪問",
                "authPassed": "驗證通過",
                "authFailure": "驗證失敗",
                "otherErr": "其他錯誤",
                "pathCheckPassed": "安全路徑檢查通過",
                "pathCheckFailure": "安全路徑不正確",
                "rememberSecureMention": "為了保證後台管理的安全性，如果這不是您的私人電腦請不要勾選。"
            }
        },
        "summary": {
            "cockpit": "儀表板",
            "systemConfig": "系統設置",
            "paymentConfig": "支付設置",
            "planMgr": "訂閱管理",
            "userMgr": "用戶管理",
            "orderMgr": "訂單管理",
            "keyMgr": "密鑰管理",
            "incomeText": "昨日收入 / 當月收入",
            "pendingTicket": "您有 {nums} 條待處理工單",
            "toHandle": "去查看",
            "apiAccessCard": "一週內API接口訪問次數",
            "apiAccessCardHint": "此數據僅供您了解當前的API訪問情況，並不代表您的服務器性能。",
            "incomeWeek": "一週內收入金額",
            "incomeCardHint": "這裡展示了一週內的收入金額圖表，如果緩存被清空則會導致顯示不準確。",
            "core": "核心",
            "reqErr": "遇到錯誤",
            "reqErrHint": "在獲取概覽信息時遇到錯誤，導致本次請求無法完成，因此圖表暫不能顯示，請您稍後再試。",
            "userCard": {
                "title": "用戶概覽",
                "allRegisteredUsers": "總註冊用戶",
                "activeUsers": "活躍用戶",
                "inactiveUsers": "非活躍用戶",
                "blockedUsers": "封禁或註銷"
            },
            "general": {
                "title": "一般",
                "localTime": "瀏覽器時間",
                "osType": "操作系統類型",
                "appName": "應用名稱",
                "appUrl": "應用網址",
                "currency": "貨幣單位",
                "allowRegister": "允許註冊"
            },
            "system": {
                "title": "系統配置",
                "axiosAddr": "HTTP後端地址",
                "wsAddr": "Websocket後端地址",
                "serverTime": "服務器時間",
                "uptime": "運行時長",
                "gatewayStatus": "API網關狀態",
                "dbStatus": "數據庫狀態",
                "redisStatus": "Redis狀態",
                "serverOsType": "服務器操作系統類型",
                "serverOsArch": "服務器操作系統架構",
                "runMode": "運行模式",
                "cpuNums": "網關服務器CPU核心數",
                "numCgoCall": "垃圾回收次數",
                "time": "次",
                "paymentMethods": "啟用的支付方式",
                "runOK": "運行正常",
                "runErr": "運行異常",
                "checkServer": "請檢查您後端服務器的環境配置",
                "stopRegisterHint": "您似乎禁用了新用戶註冊",
                "toSetting": "轉到設置"
            }
        },
        "queueMonit": {
            "title": "服務端監控",
            "headerHint": "請勿長時間停留在此頁面，在網絡高峰時段頻繁查詢將些許影響網關性能和數據庫吞吐量。",
            "latency": {
                "title": "服務器延遲",
                "retry": "再次測試",
                "hint": "*這裡的請求指的是客戶端發送請求到服務器後，服務器給予成功響應到客戶端所需要的時間。",
                "unit": "毫秒",
                "level": {
                    "l1": {
                        "title": "優秀",
                        "description": "這是很優秀的網絡情況，幾乎不會感覺到延遲。"
                    },
                    "l2": {
                        "title": "正常",
                        "description": "這是大多數時候的網絡情況，用戶幾乎感覺不到延遲。"
                    },
                    "l3": {
                        "title": "較差",
                        "description": "在在此情況下用戶可能感覺到輕微卡頓或延遲。"
                    },
                    "l4": {
                        "title": "差",
                        "description": "可以感受到明顯的延遲，有影響用戶體驗。"
                    },
                    "l5": {
                        "title": "非常差",
                        "description": "出現明顯的延遲並且加載速度變慢甚至無法刷新，嚴重影響用戶互動和體驗。"
                    },
                    "offline": {
                        "title": "服務器離線",
                        "description": "無法連接到服務器或處理請求錯誤，請檢查是否配置正確。"
                    }
                }
            },
            "api": {
                "title": "最近7天API請求狀態",
                "items": {
                    "ok": { "title": "成功 (StatusOK)", "unit": "次" },
                    "notFound": { "title": "狀態路徑錯誤 (404)", "unit": "次" },
                    "unAuthorized": { "title": "未授權訪問 (401)", "unit": "次" },
                    "login2reg": { "title": "登錄 / 註冊", "unit": "次" }
                }
            },
            "log": {
                "title": "日誌記錄",
                "deleteLogMsg": "刪除日誌 {nums} 條",
                "deleteLogErr": "刪除日誌失敗",
                "logTableRows": "日誌表總記錄數",
                "logTableSize": "日誌表佔用空間",
                "unit": { "lines": "行", "mb": "兆字節" },
                "deleteLog": "刪除日誌",
                "exportCsv": "導出CSV",
                "deleteLogHint": "這將執行刪除前一周前所有的日誌，刪除方式為硬刪除，所有被刪除的日誌將無法恢復，建議先保存備份。",
                "warn": {
                    "title": "您確認要刪除日誌嗎？",
                    "description": "將立即刪除條目，您不能撤銷此操作。"
                },
                "export": {
                    "title": "導出csv文件",
                    "description": "這將會導出下面的表格為csv文件並下載到本地，如果您沒有開啟下載權限，下載有可能會失敗，點擊確認按鈕以執行導出。"
                },
                "table": {
                    "id": "#",
                    "method": "請求方式",
                    "path": "請求路徑",
                    "code": "狀態碼",
                    "clientIp": "客戶端IP",
                    "responseTime": "處理時間",
                    "requestAt": "請求時間"
                }
            }
        },
        "systemConfig": {
            "title": "系統設置",
            "common": { "success": "更新配置成功", "err": "更新配置失敗" },
            "site": {
                "common": { "title": "站點" },
                "appName": {
                    "title": "站點名稱",
                    "shallow": "用於顯示需要站點名稱的地方。",
                    "placeholder": "站點名稱"
                },
                "appSubName": {
                    "title": "站點副標題",
                    "shallow": "一般顯示在主要標題的下面。",
                    "placeholder": "副標題"
                },
                "appDescription": {
                    "title": "站點描述",
                    "shallow": "用於顯示需要站點描述的地方。",
                    "placeholder": "站點描述"
                },
                "appUrl": {
                    "title": "站點網址",
                    "shallow": "當前網站最新網址，將會在郵件等需要用於網址處體現。",
                    "placeholder": "站點網址"
                },
                "forceHttps": {
                    "title": "強制 HTTPS",
                    "shallow": "當站點沒有使用 HTTPS，CDN 或反代開啟強制 HTTPS 時需要開啟。"
                },
                "logoUrl": {
                    "title": "LOGO",
                    "shallow": "用於顯示需要 LOGO 的地方。",
                    "placeholder": "logo 的 url 地址"
                },
                "subscribeUrl": {
                    "title": "訂閱 URL",
                    "shallow": "用於訂閱所使用，留空則為站點 URL。 如需多個訂閱 URL 隨機獲取請使用逗號進行分割。",
                    "placeholder": "訂閱 url"
                },
                "tosUrl": {
                    "title": "用戶條款(TOS)URL",
                    "shallow": "用於跳轉到用戶條款(TOS)",
                    "placeholder": "用戶條款地址"
                },
                "stopRegister": {
                    "title": "停止新用戶註冊",
                    "shallow": "開啟後任何人都將無法進行註冊。"
                },
                "inviteRequire": {
                    "title": "強制邀請",
                    "shallow": "開啟後當新用戶註冊時必需填寫邀請碼。"
                },
                "trialSubscribe": {
                    "title": "註冊試用",
                    "shallow": "選擇需要試用的訂閱，如果沒有選項請先前往訂閱管理添加。"
                },
                "trialTime": {
                    "title": "試用時間(小時)",
                    "shallow": "新用戶註冊時訂閱試用時間。"
                },
                "currency": {
                    "title": "貨幣單位",
                    "shallow": "僅用於展示使用，更改後系統中所有的貨幣單位都將發生變更。",
                    "placeholder": "CNY"
                },
                "currencySymbol": {
                    "title": "貨幣符號",
                    "shallow": "僅用於展示使用，更改後系統中所有的貨幣單位都將發生變更。",
                    "placeholder": "¥"
                }
            },
            "security": {
                "common": { "title": "安全設置" },
                "emailVerify": {
                    "title": "郵箱驗證",
                    "description": "開啟後將會強制要求用戶進行郵箱驗證。"
                },
                "gmailAlias": {
                    "title": "禁止使用Gmail多別名",
                    "description": "開啟後Gmail多別名將無法註冊。"
                },
                "safeMode": {
                    "title": "安全模式",
                    "description": "開啟後除了站點URL以外的綁定本站點的域名訪問都將會被403。"
                },
                "adminPath": {
                    "title": "後台路徑",
                    "description": "後台管理路徑，修改後將會改變原有的admin路徑。",
                    "placeholder": "https://x.com/logo.jpeg"
                },
                "emailWhitelist": {
                    "title": "郵箱後綴白名單",
                    "description": "開啟後在名單中的郵箱後綴才允許進行註冊。"
                },
                "recaptcha": {
                    "title": "防機器人",
                    "description": "開啟後將會啟用hCaptcha防止機器人。"
                },
                "hCaptchaSiteKey": {
                    "title": "hCaptcha SiteKey",
                    "description": "該SiteKey用於請求hCaptcha服務器來標識網站編號",
                    "placeholder": "a3ca066c-0ea0-42fe-bcd2-23f4ab48d528"
                },
                "ipRegisterLimit": {
                    "title": "IP註冊限制",
                    "description": "開啟後如果IP註冊賬戶達到規則要求將會被限制註冊，請注意IP判斷可能因為CDN或前置代理導致問題。"
                },
                "registerTimes": {
                    "title": "次數",
                    "description": "達到註冊次數後開啟懲罰。",
                    "placeholder": "請輸入"
                },
                "lockTime": {
                    "title": "懲罰時間(分鐘)",
                    "description": "需要等待懲罰時間過後才可以再次註冊。",
                    "placeholder": "請輸入"
                }
            },
            "frontend": {
                "common": { "title": "個性化" },
                "themePropColor": {
                    "default": "預設",
                    "darkBlueDay": "深藍色",
                    "milkGreenDay": "奶綠色",
                    "bambooGreen": "若竹",
                    "mistyPine": "霧松",
                    "glacierBlue": "冰川藍",
                    "grayTheme": "灰色"
                },
                "sidebarStyle": {
                    "title": "邊欄風格",
                    "shallow": "設置側邊欄的顏色。"
                },
                "headerStyle": { "title": "頭部風格", "shallow": "設置頂部的顏色。" },
                "themeColor": {
                    "title": "主題色",
                    "shallow": "設置整個網頁的主題色。"
                },
                "background": {
                    "title": "背景",
                    "shallow": "將會在後台登錄頁面進行展示。",
                    "placeholder": "https://x.com/logo.jpeg"
                }
            },
            "inviteAndRebate": {
                "common": { "title": "支付和返利" },
                "inviteRebateEnable": {
                    "title": "啟用用戶返利",
                    "description": "如果啟用則在被邀請的用戶充值時，按照下面設置的充值比例給予用戶返利。"
                },
                "inviteRebateRate": {
                    "title": "返利比例",
                    "description": "設置返利的金額比例。",
                    "placeholder": "請輸入返利比例"
                },
                "discountInfo": {
                    "title": "優惠信息",
                    "description": "設置優惠信息，它將會展示在充值頁面頂部。",
                    "placeholder": "設置優惠信息"
                },
                "inviteInfo": {
                    "title": "邀請信息",
                    "description": "設置邀請信息，它將會展示在用戶邀請頁面，用於顯示返利比例。",
                    "placeholder": "設置返利信息"
                }
            },
            "welcome": {
                "common": { "title": "首頁信息" },
                "homeDescription": {
                    "title": "首頁描述",
                    "description": "設置首頁的簡要描述內容。",
                    "placeholder": "請輸入首頁描述內容"
                },
                "whyChooseUs": {
                    "title": "為什麼選擇我們",
                    "description": "設置關於為什麼選擇我們的描述。",
                    "placeholder": "請輸入詳細描述"
                },
                "bilibiliLink": {
                    "title": "Bilibili 官方鏈接",
                    "description": "設置 Bilibili 官方賬號的鏈接地址。",
                    "placeholder": "https://space.bilibili.com/xxxx"
                },
                "youtubeLink": {
                    "title": "YouTube 官方鏈接",
                    "description": "設置 YouTube 官方賬號的鏈接地址。",
                    "placeholder": "https://youtube.com/channel/xxxx"
                },
                "instagramLink": {
                    "title": "Instagram 官方鏈接",
                    "description": "設置 Instagram 官方賬號的鏈接地址。",
                    "placeholder": "https://instagram.com/xxxx"
                },
                "wechatLink": {
                    "title": "微信公眾賬號鏈接",
                    "description": "設置微信公眾賬號的鏈接地址。",
                    "placeholder": "請輸入微信公眾鏈接"
                },
                "filingNumber": {
                    "title": "備案號",
                    "description": "設置站點的備案號。",
                    "placeholder": "如：粵ICP備12345678號"
                },
                "pageSuffix": {
                    "title": "站點後綴",
                    "description": "設置站點名稱後綴，用於標題顯示。",
                    "placeholder": "如：- 你的站點名稱"
                }
            },
            "sendMail": {
                "common": {
                    "title": "郵件設置",
                    "warning": "如果您更改了本頁面的配置，需要對隊列服務及逆行重啟。 另外本頁配置優先級高於.env中的郵件配置。",
                    "sendTestMailTitle": "發送測試郵件",
                    "sendTestMailShallow": "郵件將會發送到當前登錄管理員的郵箱",
                    "sendTestMailTo": "發送測試郵件到",
                    "sending": "發送郵件中",
                    "success": "成功",
                    "receiverAddr": "收信地址",
                    "senderHost": "發信服務器",
                    "senderPort": "發信端口",
                    "senderEncrypt": "發信加密方式",
                    "senderUsername": "發信用戶名",
                    "sendErr": "發送郵件失敗"
                },
                "smtpServerAddress": {
                    "title": "SMTP 服務器地址",
                    "shallow": "由郵件服務商提供的服務地址",
                    "placeholder": "請輸入服務器地址"
                },
                "smtpServerPort": {
                    "title": "SMTP 服務端口",
                    "shallow": "常見的端口有 25, 465, 587",
                    "placeholder": "請輸入端口號"
                },
                "smtpEncryption": {
                    "title": "SMTP 加密方式",
                    "shallow": "465 端口加密方式一般為 SSL，587 端口加密方式一般為 TLS",
                    "placeholder": "請輸入加密方式"
                },
                "smtpAccount": {
                    "title": "SMTP 賬號",
                    "shallow": "由郵件服務商提供的賬號",
                    "placeholder": "請輸入賬號"
                },
                "smtpPassword": {
                    "title": "SMTP 密碼",
                    "shallow": "由郵件服務商提供的密碼",
                    "placeholder": "請輸入密碼"
                },
                "senderAddress": {
                    "title": "發件地址",
                    "shallow": "由郵件服務商提供的發件地址",
                    "placeholder": "請輸入發件地址"
                },
                "emailTemplate": {
                    "title": "郵件模版",
                    "shallow": "你可以在文檔查看如何自定義郵件模板",
                    "placeholder": "請選擇郵件模板"
                }
            },
            "notice": {
                "common": { "title": "通知設置" },
                "displayName": {
                    "title": "顯示名稱",
                    "shallow": "僅用於前端頁面顯示",
                    "placeholder": "通用通知接口1"
                },
                "barkEndpoint": {
                    "title": "Bark 接入點",
                    "shallow": "Bark 服務器後端 API 地址",
                    "placeholder": "https://<ip>:<port>/<secret-key>"
                },
                "barkGroup": {
                    "title": "Bark 群組",
                    "shallow": "客戶端顯示的群組名稱",
                    "placeholder": "web"
                }
            },
            "appDownload": {
                "common": {
                    "title": "APP",
                    "hint": "用於自有客戶端(APP)的版本管理及更新"
                },
                "enabled": {
                    "title": "開放下載",
                    "shallow": "如果啟用則允許用戶訪問下載頁面"
                },
                "windows": {
                    "title": "Windows",
                    "shallow": "Windows 端版本號及下載地址",
                    "placeholder": "https://xxxx.com/xxx.exe"
                },
                "macos": {
                    "title": "macOS",
                    "shallow": "macOS 端版本號及下載地址",
                    "placeholder": "https://xxxx.com/xxx.dmg"
                },
                "linux": {
                    "title": "Linux",
                    "shallow": "Linux 端版本號及下載地址",
                    "placeholder": "https://xxxx.com/xxx.deb"
                },
                "android": {
                    "title": "Android",
                    "shallow": "Android 端版本號及下載地址",
                    "placeholder": "https://xxxx.com/xxx.apk"
                },
                "ios": {
                    "title": "IOS",
                    "shallow": "IOS 端版本號及下載地址",
                    "placeholder": "https://xxxx.com/xxx.ipk"
                }
            }
        },
        "payConfig": {
            "title": "支付設置",
            "description": "在這裡可以管理所有支持的付款方式，目前僅支持支付寶支付，但是您也可以先行配置其他支付方式，如沒有符合您的支付流程，可以等待項目進一步完善後在更新日誌中查看是否支持。",
            "attention": {
                "title": "注意事項",
                "point1": "務必先配置支付方式的信息再進行啟用，這真的很重要。",
                "point2": "修改付款方式配置時，如果顯示為\"---\"則代表該選項已經被設置且非空，如果需要重新修改直接輸入新值保存即可。"
            },
            "common": {
                "detail": "{method} 配置",
                "fillAttention": "為確保安全，不顯示詳細信息，重新填寫以創建或覆蓋已有配置。",
                "discountAmount": "優惠金額（可以在系統設置的 \"支付和返利\" 中設置用戶提示信息）",
                "saveConfigBtnHint": "保存",
                "cancelBtnHint": "取消",
                "saveSuccess": "配置保存成功",
                "alterSuccess": "配置修改成功",
                "discountPlaceholder": "優惠金額（充值金額大於優惠金額時應用優惠）",
                "saveOrAlterFailure": "未知錯誤"
            },
            "alipay": {
                "title": "支付寶",
                "config": {
                    "appId": "應用Id",
                    "appPublicKeyCertContent": "應用公鑰證書內容",
                    "appPrivateKey": "應用私鑰",
                    "alipayRootCert": "支付寶根證書",
                    "alipayPublicKeyCertContent": "支付寶公鑰證書內容"
                }
            },
            "wechat": {
                "title": "微信支付",
                "config": {
                    "mchId": "商戶Id",
                    "mchCertSerial": "商戶證書序列號",
                    "apiV3Key": "API v3 Key",
                    "privateKey": "私鑰"
                }
            },
            "apple": {
                "title": "Apple Pay",
                "config": {
                    "issuerId": "Issuer Id",
                    "bundleId": "Bundle Id",
                    "privateKeyId": "Private Key Id",
                    "privateKeyContent": "私鑰內容"
                }
            },
            "addPaymentMethod": "添加支付方式",
            "enableBtnHint": "啟用",
            "disableBtnHint": "禁用",
            "setupPaymentMethod": "配置"
        },
        "themeConfig": {
            "title": "主題配置",
            "using": "使用，",
            "setAsCurrent": ""
        },
        "groupMgr": {
            "title": "權限組管理",
            "description": "權限組是用來標示不同的訂閱等級，您可以將同等級別但是額度等有細微區別的訂閱計畫歸納在一個權限組中方便管理。",
            "common": {
                "addNewGroup": "新建權限組",
                "alterGroupName": "修改權限組名稱",
                "addConfirmed": "確認添加",
                "alterConfirmed": "確認修改",
                "cancel": "取消",
                "addSuccess": "添加成功",
                "addFailure": "添加失敗",
                "alterSuccess": "修改權限組成功",
                "alterFailure": "修改權限組失敗",
                "delSuccess": "刪除成功",
                "delFailure": "刪除失敗",
                "delMention": "將立即刪除條目，您不能撤銷此操作。 已經關聯訂閱計畫的權限組謹慎刪除。",
                "delNotAllowed": "該權限組有關聯資源，不可以刪除。"
            },
            "groupId": "權限組ID",
            "groupName": "權限組名稱",
            "groupPlaceholder": "輸入權限組名稱",
            "userCount": "用戶數量",
            "planCount": "訂閱數量",
            "operate": "操作",
            "editBtnHint": "編輯",
            "deleteBtnHint": "刪除"
        },
        "docMgr": {
            "title": "知識庫管理",
            "description": "在這裡您可以編寫給用戶的說明文檔，它可以是軟件的設計說明書、使用教程或註意事項等，其文檔的內容支持Markdown格式。",
            "addDoc": "添加新文檔",
            "addSuccess": "添加新文檔成功",
            "addFailure": "添加文檔失敗",
            "titleNotEmpty": "文檔的標題不能為空",
            "table": {
                "docId": "#",
                "isShow": "是否顯示",
                "sortAs": "排序",
                "lang": "語言",
                "category": "分類",
                "title": "標題",
                "createdAt": "創建時間",
                "updatedAt": "更新時間",
                "operate": "操作",
                "edit": "編輯",
                "delete": "刪除"
            },
            "form": {
                "add": "添加文檔",
                "edit": "編輯文檔",
                "cancel": "取消",
                "confirm": "確認",
                "addBtn": "添加",
                "editBtn": "修改",
                "title": { "title": "文檔標題", "placeholder": "輸入文檔的標題" },
                "sort": {
                    "title": "排序",
                    "placeholder": "輸入文檔的排序級別 越高的數值代表優先級越高"
                },
                "category": {
                    "title": "分類",
                    "placeholder": "文檔將按照該字段進行分類展示"
                },
                "lang": { "title": "文檔語言", "placeholder": "選擇文檔語言" }
            }
        },
        "planMgr": {
            "title": "訂閱管理",
            "description": "在這裡可以添加新的訂閱計畫、修改已有訂閱計畫的描述、價格、餘糧、其所屬的權限組等。",
            "addNewPlan": "添加新訂閱",
            "table": {
                "id": "#",
                "isSale": "啟用銷售",
                "isRenew": "允許續費",
                "sort": "排序等級",
                "group": "所屬權限組",
                "name": "名稱",
                "nums": "數量",
                "residue": "餘量",
                "operate": "操作",
                "operateBtn": { "update": "修改", "delete": "刪除" }
            },
            "mention": {
                "common": {
                    "success": "成功",
                    "failure": "失敗",
                    "delMention": "如果訂閱計畫已經處於銷售中，謹慎刪除。"
                }
            },
            "form": {
                "title": "添加訂閱",
                "items": {
                    "name": { "title": "訂閱名稱", "placeholder": "輸入訂閱計畫的名稱" },
                    "describe": {
                        "title": "訂閱描述",
                        "placeholder": "輸入訂閱的描述（支持Markdown）"
                    },
                    "isSale": { "title": "啟用銷售" },
                    "isRenew": { "title": "啟用續費" },
                    "group": { "title": "所屬權限組", "placeholder": "選擇所屬權限組" },
                    "capacityLimit": {
                        "title": "最大允許用戶數目",
                        "placeholder": "最大允許用戶數目"
                    },
                    "planResidue": {
                        "title": "剩餘數量",
                        "placeholder": "目前剩餘的數量"
                    },
                    "sort": { "title": "排序", "placeholder": "用於前端排序" },
                    "periodPlaceholder": {
                        "month": "輸入月付價格",
                        "quarter": "輸入季付價格",
                        "halfYear": "輸入半年付價格",
                        "year": "輸入年付價格"
                    }
                }
            }
        },
        "couponMgr": {
            "title": "優惠券管理",
            "description": "在這裡您可以為一些特定的節日等添加一些優惠券，它允許用戶在下單的時候使用並按照您設置的比例進行抵折優惠。",
            "fetchErr": "獲取數據失敗請重試",
            "fetchPlanKvFailurese": "獲取訂閱計畫列表失敗",
            "table": {
                "id": "#",
                "enabled": "是否啟用",
                "name": "名稱",
                "code": "券碼",
                "percentOff": "優惠信息",
                "capacity": "總數量",
                "residue": "剩餘數量",
                "startTime": "啟用時間",
                "endTime": "結束時間",
                "actions": "操作",
                "edit": "編輯",
                "delete": "刪除"
            },
            "modal": {
                "newCoupon": "添加新優惠券",
                "editCoupon": "修改優惠券信息",
                "confirmAdd": "確認添加",
                "confirmEdit": "確認修改",
                "emptyNotAllow": "該項目是必填的",
                "delMention": "條目刪除後優惠券將會立即失效，您不能撤銷此操作。",
                "cancel": "取消",
                "name": { "title": "優惠券名稱", "placeholder": "請輸入優惠券名稱" },
                "code": {
                    "title": "券碼",
                    "placeholder": "自定義優惠券碼（留空隨即生成）"
                },
                "percentOff": {
                    "title": "優惠信息（百分比）",
                    "placeholder": "輸入優惠百分比（如-12%OFF）"
                },
                "activeRange": { "title": "優惠券可使用的有效期範圍" },
                "capacity": {
                    "title": "優惠券最大使用次數",
                    "placeholder": "限制最大使用限制（為空則不限制）"
                },
                "residue": {
                    "title": "優惠券剩餘使用次數",
                    "placeholder": "設置優惠券剩餘使用的次數"
                },
                "perUserLimit": {
                    "title": "每個用戶可使用優惠券的次數",
                    "placeholder": "限制每個用戶可使用的次數（為空則不限制）"
                },
                "planLimit": {
                    "title": "指定訂閱計畫",
                    "placeholder": "限制指定訂閱計畫可以使用優惠（為空則不限制）"
                }
            }
        },
        "orderMgr": {
            "title": "訂單管理",
            "description": "在這裡您可以檢視所有訂閱計畫的訂單，篩選不同用戶的訂單、手動對用戶的訂單處理通過等。",
            "table": {
                "id": "#",
                "orderId": "訂單號",
                "email": "用戶郵箱",
                "status": { "title": "類型", "t1": "新購", "t2": "續費", "t3": "編輯" },
                "planName": "計畫名稱",
                "period": "週期",
                "group": "權限組",
                "amount": "實付金額",
                "price": "原始價格",
                "isSuccess": {
                    "title": "訂單狀態",
                    "cancelOrder": "取消訂單",
                    "passOrder": "通過訂單"
                },
                "createdAt": "訂單創建時間",
                "action": { "title": "操作", "showDetail": "顯示細節" }
            },
            "search": "查詢訂單",
            "resetSearch": "重置查詢",
            "failureReason": "失敗原因",
            "couponId": "優惠券ID",
            "couponName": "優惠券名稱",
            "noEntry": "無",
            "orderDetail": "訂單詳情",
            "searchModal": {
                "email": {
                    "title": "用戶郵箱",
                    "placeholder": "輸入用戶郵箱（模糊搜索）"
                },
                "sort": {
                    "title": "排序算法",
                    "placeholder": "選擇排序算法",
                    "ASC": "升序",
                    "DESC": "降序"
                }
            },
            "tradeWaiting": "未支付",
            "tradeFailure": "交易失敗",
            "tradeSuccess": "成功"
        },
        "userMgr": {
            "userManager": "用戶管理",
            "description": "你可以在這裡管理所有的用戶，包括員工和管理員，授予或取消管理權限、設定用戶餘額、重置密碼、手動添加新用戶等操作。",
            "enterEmail": "請輸入郵箱",
            "enterValidEmail": "請輸入正確的郵箱格式",
            "enterPassword": "請輸入密碼",
            "passwordMinLength": "密碼長度不能少於 6 位",
            "passwordMaxLength": "密碼長度不能超過 20 位",
            "passwordStrength": "密碼必須包含字母、數字和特殊字符",
            "validationSuccess": "驗證通過",
            "validationFailed": "表單驗證失敗，請檢查輸入項",
            "editProfile": "編輯資料",
            "banUser": "封禁用戶",
            "unbanUser": "解封用戶",
            "noRecord": "無記錄",
            "normal": "正常",
            "banned": "封禁",
            "deleted": "註銷",
            "nullContent": "NULL",
            "balance": "餘額",
            "orderCount": "訂單數量",
            "planCount": "計劃數",
            "actions": "操作",
            "updateSuccess": "更新成功",
            "addUserSuccess": "添加新用戶成功",
            "unknownError": "未知錯誤",
            "email": "郵箱",
            "registerDate": "註冊日期",
            "isAdmin": "管理員",
            "isStaff": "員工",
            "accountStatus": "帳戶狀態",
            "inviteCode": "邀請碼",
            "query": "查詢",
            "reset": "重置",
            "addNewUser": "新增用戶",
            "searchUser": "查詢用戶",
            "cancel": "取消",
            "submit": "提交",
            "userEmail": "用戶郵箱",
            "inputUserEmail": "輸入用戶郵箱（模糊搜索）",
            "inputEmail": "輸入郵箱",
            "password": "密碼",
            "inputPassword": "輸入密碼",
            "editUser": "編輯用戶",
            "no": "否",
            "yes": "是",
            "mention": {
                "title": "您確定要封禁用戶嗎？",
                "content": "如果封禁該用戶，那麼該用戶將無法登錄到{appName}，並且與其相關的所有資源將變為不可用。"
            }
        },
        "activation": {
            "activateLog": "激活紀錄",
            "description": "您可以檢視所有售出密鑰的具體激活情況，查看客戶端的識別碼、激活時間等。",
            "click2getKey": "點擊以獲取密鑰內容",
            "createdAt": "創建時間",
            "turn2keyPage": "轉到密鑰",
            "showKey": "顯示密鑰",
            "email": "郵箱",
            "key": "密鑰",
            "search": "搜尋",
            "filter": "篩選",
            "filterAll": "全部",
            "filterActive": "僅有效",
            "sortAlgorithm": "排序算法",
            "sortASC": "升序",
            "sortDESC": "降序"
        },
        "keysMgr": {
            "keyMgr": "密鑰管理",
            "description": "您可以查詢到所有生成的密鑰內容和他們的使用情況、有效期等等，您也可以執行禁用某一個密鑰。",
            "enableKey": "啟用密鑰",
            "disableKey": "禁用密鑰",
            "table": {
                "email": "郵箱地址",
                "key": "密鑰內容",
                "isValid": "是否有效",
                "isUsed": "是否使用",
                "isExpired": "是否到期",
                "active": "活躍",
                "inactive": "已過期",
                "yes": "有效",
                "no": "過期",
                "ok": "正常",
                "blocked": "禁用狀態",
                "unUsed": "未使用",
                "used": "已使用",
                "showDetail": "顯示細節",
                "blockKey": "禁用密鑰",
                "unblockKey": "解除禁用"
            },
            "searchModal": {
                "searchMethod": "查詢方式",
                "byId": "通過ID進行查詢",
                "byContent": "通過密鑰進行查詢（模糊）",
                "keyId": "密鑰ID",
                "idPlaceholder": "在這裡輸入密鑰的ID",
                "keyContent": "密鑰碼",
                "contentPlaceholder": "在這裡輸入密鑰的部分內容"
            },
            "mention": {
                "blockOk": "禁用密鑰成功 ID:{id}",
                "title": "您確定要禁用密鑰嗎",
                "content": "為保證用戶體驗，請再一次確認密鑰內容，密鑰一旦禁用後，在客戶端將不能再繼續使用。"
            },
            "detailModal": {
                "title": "密鑰細節",
                "userId": "用戶ID",
                "planName": "訂閱計畫名稱",
                "expiredAt": "到期日期",
                "keyGeneratedAt": "密鑰生成日期",
                "clientId": "客戶端ID"
            }
        },
        "noticeMgr": {
            "title": "公告管理",
            "description": "在這裡可以對公告進行管理，啟用的公告將展示在用戶首頁的輪播圖中，可以設置的公告有優惠活動、節日通知、注意事項等。",
            "addNotice": "添加公告",
            "table": {
                "id": "#",
                "show": "是否顯示",
                "title": "標題",
                "createdAt": "創建時間",
                "action": { "title": "操作", "edit": "編輯", "delete": "刪除" }
            },
            "mention": {
                "title": "您確認要刪除嗎？",
                "content": "將立即刪除該公告，您不能撤銷此操作。"
            },
            "modal": {
                "addNew": "新建公告",
                "title": {
                    "title": "公告標題",
                    "placeholder": "作為大標題顯示在輪播圖中"
                },
                "content": { "title": "公告內容", "placeholder": "編寫公告的主要內容" },
                "tag": { "title": "公告標籤", "placeholder": "輸入公告的標籤" },
                "img": { "title": "背景圖片URL", "placeholder": "不填寫則使用默認背景" }
            }
        },
        "adminTicket": {
            "ticketMgr": "工單管理",
            "description": "您可以在此處檢視到所有用戶提交的工單，工單有三個緊急程度，推薦您從緊急工單開始處理。",
            "pendingTicket": "待處理工單",
            "finishedTicket": "已完成工單",
            "type": {
                "pendingDescription": "活躍的工單，這是您應該先進行處理的工單，如果工單確認已經完成請關閉，如不關閉則該工單將始終置於此處。",
                "finishedDescription": "已經完成的工單，您可以在這裡檢視它們。"
            },
            "chooseOneNecessary": "您應該至少選擇一個項目",
            "mention": {
                "title": "您確定要關閉該工單嗎？",
                "content": "該工單在關閉後將會歸檔至已完成工單中，您可以再次檢視他們的內容，但是再也無法回复此工單。"
            }
        }
    },
    "userLogin": {
        "loginToContinue": "登入以繼續",
        "email": "郵件地址",
        "password": "密碼",
        "haveNoAccount": "還沒有您的帳戶？",
        "login": "登入",
        "reg": "立即註冊",
        "otherMethods": "或使用其他方式繼續",
        "github": "以Github帳戶繼續",
        "microsoft": "以Microsoft帳戶繼續",
        "apple": "以Apple帳戶繼續",
        "google": "以Google帳戶繼續",
        "backHomePage": "回到首頁",
        "form": {
            "emailRequire": "郵箱地址是必填的",
            "passwordRequire": "你還沒有輸入密碼"
        },
        "authPass": "驗證通過",
        "loginFailure": "登錄失敗",
        "checkForm": "請檢查表單",
        "if2FaEnabledHint": "如果您啟用了兩步驗證（非必填）",
        "reqErr": "請求出現錯誤請稍後再試",
        "accountLocked": "您的帳戶可能被註銷或封禁，暫時無法繼續使用，如你仍然認為這是一個錯誤，請聯繫我們的技術支持。",
        "tokenNotExist": "請提供Token"
    },
    "userRegister": {
        "backHomePage": "回到首頁",
        "newAccount": "準備您的新帳戶",
        "email": "郵箱地址",
        "verifyCode": "郵箱驗證碼",
        "invalidEmailFormat": "郵箱格式不合法",
        "sendVerifyCode": "發送郵箱驗證碼",
        "sendSuccess": "郵件發送成功，請查看郵箱。",
        "pwd": "密碼",
        "pwdAgain": "確認密碼",
        "inviteCode": "邀請碼（選填）",
        "agreement": "我已閱讀並同意",
        "terminalUserAgreement": "用戶條款",
        "reg": "註冊",
        "infoIncomplete": "信息不完整",
        "pwdIncorrect": "密碼不一致",
        "regSuccess": "註冊成功返回登錄",
        "regFailure": "註冊失敗",
        "success": "成功",
        "failure": "失敗",
        "unknownErr": "未知錯誤",
        "verifyCodeErr": "驗證碼錯誤",
        "verifyCodeExpireErr": "驗證碼錯誤或已過期，請重試或獲取新驗證碼。",
        "thisMailAlreadyExist": "該郵箱已經被註冊",
        "pageConfigFetchFailure": "配置獲取失敗請刷新重試",
        "stopRegisterTitle": "已停止註冊",
        "stopRegisterHint": "很抱歉，目前註冊功能已暫停。 如閣下有需要，請稍後再試或聯繫我們的支持團隊獲取更多信息。 感謝您的理解與支持。",
        "passwordComplexRequirePart1": "* 密碼需要符合",
        "passwordComplexRequirePart2": "複雜度要求",
        "passwordComplexHint1": "1. 需要使用大寫字母、小寫字母、數字、特殊符號中的三種及以上",
        "passwordComplexHint2": "2. 長度再10位以上",
        "form": {
            "checkForm": "請檢查表單內容",
            "emailFormatErr": "郵箱地址格式不正確",
            "gmailLimitErr": "不允許使用Gmail多別名",
            "gmailDotNotAllowed": "Gmail地址中不允許有\".\"",
            "gmailPartLowerForced": "Gmail地址的本地部分必須全部小寫",
            "googlemailNotAllowed": "不支持使用googlemail地址",
            "verifyCodeRequire": "請輸入驗證碼",
            "verifyCodeFormatErr": "驗證碼格式不正確",
            "passwordRequire": "請輸入密碼",
            "passwordLengthRequire": "密碼長度必須大於或等於10位",
            "passwordComplexRequire": "密碼必須包含大寫字母、小寫字母、數字、特殊符號中的至少三種",
            "passwordAgainNotMatch": "兩次輸入的密碼不匹配",
            "passwordAgainRequire": "請輸入確認密碼",
            "inviteCodeRequire": "邀請碼是必填的"
        },
        "hCaptcha": {
            "passed": "人機驗證通過",
            "expired": "已過期請重試",
            "challengeExpired": "驗證超時",
            "err": "其他錯誤"
        },
        "allRightsReserved": "版權所有",
        "securityAndLaws": "該網站受hCaptcha保護和驗證，請遵守當地法律。"
    },
    "userSummary": {
        "title": "儀表盤",
        "myPlan": "我的訂閱",
        "shortcut": "捷徑",
        "timeLeft": "訂閱有效，將在 {msg} 過期。",
        "toPurchase": "購買訂閱",
        "tutorial": { "title": "查看教程", "content": "學習如何使用 {name}" },
        "checkKey": {
            "title": "查看密鑰",
            "content": "快速將密鑰導入對應客戶端進行使用"
        },
        "renewPlan": { "title": "續費訂閱", "content": "對您當前的訂閱進行續費" },
        "appDown": {
            "title": "應用下載",
            "content": "下載我們不同平台的應用程序以獲得更好的體驗"
        },
        "support": {
            "title": "遇到問題",
            "content": "遇到問題可以通過工單與我們的人機溝通"
        },
        "haveTicket": "您有 {count} 條待處理的工單",
        "toCheckTicket": "去查看",
        "showAllKeys": "查看所有密鑰"
    },
    "userDocument": {
        "title": "使用文檔",
        "description": "您可以在這裡查閱不同的文檔，包括但不限於網站的使用方法、注意事項、操作流程等，如果您發現文章中有錯誤，請提交工單。",
        "searchPlaceholder": "請輸入要搜索的內容（模糊搜索）",
        "searchBtn": "搜尋",
        "noContentTitle": "無結果",
        "noContentTitleHint": "沒有符合您搜索結果或語言的文檔，嘗試換一個關鍵詞吧。"
    },
    "newPurchase": {
        "title": "購買訂閱",
        "description": "您可以在這裡選擇最適合您的訂閱計畫，如果您的餘額不足請先充值，訂單將為您保留5分鐘。",
        "headerPlaceholder": "選擇最適合您的計畫",
        "purchaseBtn": "訂購",
        "noLeft": "剩餘數量不足",
        "monthPay": "月付",
        "moreMethod": "更多選擇"
    },
    "newSettlement": {
        "err": "錯誤",
        "monthPay": "月付",
        "quarterPay": "季付",
        "halfYearPay": "半年付",
        "yearPay": "年付",
        "payCycle": "付款週期",
        "couponPlaceholder": "有優惠券 ？",
        "verifyBtn": "驗證",
        "orderTotalTitle": "訂單總額",
        "total": "總計",
        "submitOrder": "提交訂單",
        "coupon": "優惠券",
        "notify": {
            "passTitle": "驗證通過",
            "couponVerified": "優惠券有效",
            "couponInvalid": "優惠券無效",
            "couponIsNull": "輸入的優惠券碼不能為空"
        }
    },
    "userProfile": {
        "title": "個人中心",
        "myWallet": "我的錢包",
        "walletSub": "帳戶餘額（僅用於消費）",
        "alertPwd": "修改密鑰",
        "oldPwd": "舊密鑰",
        "oldPwdSub": "請輸入舊密碼",
        "newPwd": "新密鑰",
        "newPwdSub": "請輸入新的密鑰",
        "newPwdAgain": "確認密鑰",
        "newPwdAgainSub": "請再輸入一遍新的密鑰",
        "saveBtn": "保存",
        "notify": "通知",
        "enableNotify": "啟用到期通知",
        "deleteAccount": "註銷帳戶",
        "deleteAccountSub": "您的帳戶將被標記為刪除，如果需要重新使用我們的服務，請重新註冊",
        "deleteBtn": "註銷我的帳戶",
        "oldPwdVerifiedFailure": "舊密碼驗證失敗",
        "alertFailure": "密鑰修改失敗",
        "alertSuccess": "修改成功",
        "alertSuccessSub": "請使用新密碼登錄",
        "errorPwdFormat": "密碼格式錯誤",
        "pwdNotMatch": "兩次密碼輸入不一致",
        "oldPwdNotNull": "舊密碼不能為空",
        "toTopUp": "去充值",
        "deleteMyTitle": "警告",
        "deleteMyContent": "確定刪除您的賬戶嗎？ 如需要使用服務請重新註冊。",
        "deleteMyPositiveText": "確認刪除",
        "deleteMyNegativeText": "取消",
        "deletedSuccessMsg": "賬戶已刪除，後會有期。",
        "deleteErrOccur": "遇到錯誤",
        "faAuth": "兩步驗證2FA",
        "faAuthHint": "兩步驗證是一種安全機制，增加了登入帳戶的保護層。 用戶在輸入密碼後，還需完成第二步身份驗證，如輸入發送到手機的驗證碼、使用身份驗證應用程序生成的動態碼，或通過指紋等生物特徵確認。",
        "faAuthStatus": "兩步驗證狀態",
        "faEnabled": "已啟用",
        "faNotEnabled": "未啟用",
        "setup2Fa": "設置兩步驗證",
        "disable2Fa": "取消兩步驗證",
        "unknownErr": "未知錯誤",
        "disable2FaCancelled": "已取消",
        "closed2FaHint": "已經關閉兩步驗證，如有需要請重新添加。",
        "setup2FaModal": {
            "followStep": "根據提示在您的驗證器上加入",
            "step1Title": "按照以下步驟以啟2FA驗證",
            "step1Context1": "1. 您的移動設備上需要有一個通用的驗證器",
            "step1Context2": "2. 點擊驗證器上的Scan按鈕來掃描此處的QR碼",
            "step1Context3": "3. 該QR Code中包含有您的驗證信息和唯一密鑰，請妥善保存",
            "step2Context1": "為了確保您的驗證器能夠正常使用，我們需要進行測試。",
            "test2Fa": "測試",
            "cancel": "取消"
        },
        "deleteMyAccountModal": {
            "title": "註銷帳戶",
            "contentLine1": "帳號註銷是一個不可逆的操作。 一旦您確認刪除，您將永久失去該帳號的訪問權限，這意味著您將無法再次登入，且與此帳號相關的所有數據，包括但不限於您的個人資訊、歷史記錄、收藏內容、購買記錄等，都將無法再訪問。",
            "contentLine2": "如果您在我們的平台上有正在進行的業務，例如未完成的訂單、正在參與的活動、訂閱服務等，這些都將隨著帳號刪除而終止或取消，可能會給您帶來相應的損失。 同時，您與其他用戶之間通過本平台建立的聯繫、互動資訊等也都將不復存在。",
            "contentLine3": "請再次確認您的決定，如果您還有任何疑慮或問題，歡迎聯繫我們的客服，我們將竭誠為您解答。 若您仍然希望刪除帳號，請點擊「確認刪除」按鈕。",
            "inputHint1": "輸入 \"",
            "inputHint2": "\" 以繼續。",
            "confirmDelete": "確認刪除"
        },
        "failure": "遇到錯誤",
        "notLatestHint": "個人信息更新失敗，當前的數據可能不是最新的。",
        "resetPassword": {
            "previousPasswordRequire": "請輸入舊密碼",
            "previousPasswordVerifiedFailure": "舊密碼驗證失敗",
            "passwordRequire": "請輸入密碼",
            "passwordConflict": "新密碼不能與先前的密碼相同",
            "passwordLengthRequire": "密碼長度必須大於或等於10位",
            "passwordComplexRequire": "密碼必須包含大寫字母、小寫字母、數字、特殊符號中的至少三種",
            "passwordAgainNotMatch": "兩次輸入的密碼不匹配",
            "passwordAgainRequire": "請輸入確認密碼",
            "updatePasswordFailure": "更新密碼出錯"
        },
        "secondary": {
            "title": "基礎資料",
            "card": {
                "avatar": "用戶頭像",
                "name": "用戶名",
                "lastLoginAt": "上一次登錄時間",
                "accountStatus": "帳戶狀態"
            },
            "modify": {
                "uploadIconHint": "上傳",
                "alterAvatar": "新增/修改頭像",
                "alterShallow": "* 點擊頭像或用戶名以添加或修改（設置後均不允許清除）",
                "alterName": "修改用戶名",
                "input": {
                    "label": "使用者名稱",
                    "placeholder": "輸入用戶名",
                    "spaceIsNotAllowed": "用戶名驗證不通過",
                    "require": {
                        "p1": "不允許純數字且數字開頭",
                        "p2": "不允許有空格",
                        "p3": "長度大於三位"
                    }
                }
            },
            "mention": {
                "alterNameSuccess": "修改用戶名成功",
                "alterNameErr": "修改用戶名失敗請稍後再試",
                "newNameIsNotValid": "用戶名不合法",
                "click2SetName": "點擊以設置用戶名",
                "fetchAvatarErr": "獲取頭像數據失敗請稍後再試",
                "alterAvatarErr": "修改頭像失敗請稍後再試",
                "success": "成功",
                "alterAvatarSuccess": "修改頭像成功，",
                "uploadImageHint": "您可以上傳圖像作為您的新頭像",
                "imageRequire": {
                    "title": "注意",
                    "p1": "您可以上傳 *.jpg(jpeg), *.png, *.webp 等主流格式。",
                    "p2": "圖像的長寬比為1:1（正方形），如果為其他比例，將會被居中裁切為正方形，多餘的部分將會被刪除。",
                    "p3": "您上傳的圖像的大小將會被設置為160px。"
                },
                "click2Upload": "點擊或拖動文件到該區域來上傳",
                "uploadWarn": "*請不要上傳敏感數據，比如您的銀行卡、信用卡、密碼和安全碼等。"
            }
        }
    },
    "userKeys": {
        "myKeys": "我的密鑰",
        "description": "您可以檢視您的所有密鑰的活動情況、到期日期等，如需要為密鑰設置備注，請前往激活紀錄頁面。",
        "noKeys": "您還沒有有效的購買紀錄",
        "keyDetail": "密鑰細節",
        "keyId": "密鑰ID",
        "orderId": "鏈接訂單ID",
        "clientId": "激活客戶端ID",
        "active": "在有效期內",
        "inActive": "已過期",
        "valid": "密鑰狀態正常",
        "invalid": "密鑰已被禁用",
        "isUsed": "已激活使用",
        "noUsed": "還未使用",
        "releaseData": "密鑰生成日期",
        "expirationData": "到期日期",
        "none": "無",
        "authorizationFor": "密鑰授權給",
        "hoverClickMention": "點擊可以複製到剪貼板",
        "copiedSuccessMessage": "密鑰複製成功，請參照文檔說明繼續操作。",
        "copyFailure": "複製失敗",
        "hoverCopiedSuccessMention": "複製成功"
    },
    "userOrders": {
        "myOrders": "我的訂單",
        "description": "您所有的訂單將在此處展示，如果您有未支付的訂單將會展示在頂部，您可以點擊繼續支付或取消訂單，已經完成的訂單您可以在此處檢視訂單細節。",
        "orderId": "#",
        "planName": "訂閱名稱",
        "planCycle": "訂閱週期",
        "orderPrice": "訂單金額",
        "orderStatus": "訂單狀態",
        "createdAt": "創建時間",
        "operate": "操作",
        "showDetail": "訂單細節",
        "cancelOrder": "取消訂單",
        "canceled": "已取消",
        "period": {
            "monthPay": "月付",
            "quarterPay": "季付",
            "halfYearPay": "半年付",
            "yearPay": "年付"
        },
        "orderStatusTags": {
            "success": "成功",
            "cancelled": "失敗",
            "notPay": "未支付"
        },
        "orderCancelled": "訂單已取消",
        "unknownErr": "未知錯誤"
    },
    "userTopUp": {
        "topUp": "帳戶充值",
        "description": "您可以在這裡進行帳戶充值，支持自定義充值金額，您也可以關注下方是否有優惠信息展示，使用提及的支付方式以獲得優惠。",
        "chooseTopUpAmount": "選擇充值金額",
        "quickSelect": "快速選擇",
        "customAmount": "自定義金額",
        "maxAmount": "最大金額: 10,000,000",
        "amountInputPlaceHolder": "輸入要充值的金額",
        "otherAmount": "其他金額",
        "payMethod": "支付方式",
        "wechat": "微信支付",
        "alipay": "支付寶",
        "apple": "Apple Pay",
        "yourAmount": "您的金額",
        "discount": "優惠",
        "accountBalance": "帳戶餘額",
        "balanceResult": "餘額合計",
        "commitTopUp": "提交",
        "payMethodNotAllow": "支付方式不可用請選擇其他",
        "topUpIssueOccur": "充值遇到問題？",
        "payIssueOccur": "支付遇到問題？",
        "chatWithUs": "聯繫客服",
        "pay": "支付",
        "qrCodeScannedSuccess": "QR碼掃描成功",
        "orClickToApp": "或點擊條轉到App繼續",
        "topUpSuccess": "充值成功",
        "thankU": "感謝您的支持"
    },
    "userConfirmOrder": {
        "switchPlan": "切換訂閱",
        "cancelOrder": "取消訂單",
        "yourPrice": "您的價格",
        "couponOffset": "優惠券抵折",
        "price": "價格",
        "submit": "提交",
        "monthPay": "月付",
        "quarterPay": "季付",
        "halfYearPay": "半年付",
        "yearPay": "年付",
        "goodInfo": "商品信息",
        "cycleOrType": "週期/類型",
        "orderNumber": "訂單號",
        "createdAt": "創建日期",
        "orderExpired": "訂單已超時",
        "paySuccessfully": "支付成功，感謝您的支持。",
        "balanceNotEnough": "您的餘額不足請先充值，該訂單保留五分鐘。",
        "orderErrorOccur": "訂單遇到錯誤",
        "orderCancelled": "訂單已取消"
    },
    "paymentResultParts": {
        "goodInfoView": { "goodInfo": "商品信息" },
        "orderInfoView": { "orderInfo": "訂單信息" }
    },
    "orderPartUniversal": {
        "period": {
            "monthPay": "月付",
            "quarterPay": "季付",
            "halfYearPay": "半年付",
            "yearPay": "年付"
        },
        "orderDataHex": {
            "goodInfo": "商品信息",
            "orderInfo": "訂單信息",
            "cycleOrType": "週期/類型",
            "orderNumber": "訂單號",
            "createdAt": "創建日期",
            "amount": "支付金額",
            "paidAt": "支付時間"
        }
    },
    "orderDetail": {
        "finished": "已完成",
        "finishedAndSuccessDescription": "訂單已成功支付並開通",
        "useManual": "查看使用教程",
        "payPending": "尚未支付",
        "pendingDescription": "訂單暫時保留，可以點擊下面的按鈕以繼續支付。",
        "toPay": "去支付",
        "outDate": "已失效",
        "outDateDescription": "由於您取消了訂單或未在指定時間內完成支付，因此該訂單已被取消，您可以重新選取訂閱。",
        "chooseNewPlan": "選擇新的訂閱計畫"
    },
    "userInvite": {
        "myInvite": "我的邀請",
        "description": "如果管理員開啟了邀請返利，您可以在這裡看到具體的返利規則，生成邀請連結後分享給其他人註冊後即可獲取到返利。",
        "unit": "人數",
        "inviteCodeMgr": "您的邀請碼",
        "generateInviteCode": "生成隨機邀請碼",
        "faCodeManage": "邀請碼管理",
        "email": "郵箱",
        "createdAt": "註冊時間",
        "createFaCodeFailure": "創建失敗",
        "linkCopiedSuccess": "鏈接複製成功",
        "generateFaCode": "生成邀請碼",
        "flushFaCode": "刷新邀請碼",
        "faCode": "邀請碼",
        "noFaCode": "你還沒有邀請碼，請先生成。",
        "faLink": "邀請連結",
        "generateFaCodePlease": "請先生成邀請碼",
        "usersMyInvited": "我邀請的用戶",
        "generateCodeConfirm": "確認生成/刷新",
        "generateCodeHint": "請注意，邀請碼創建後不可關閉。",
        "positiveClick": "確認",
        "negativeClick": "取消"
    },
    "userTickets": {
        "description": "如果您遇到使用上的問題，請在此處提交工單，我們的技術支持和客服看到後將會進行回覆並在下方表格處標出具體的回覆時間，您可以點擊查看工單和我們交流。",
        "ticketId": "#",
        "ticketSubject": "工單主題",
        "ticketUrgency": "工單級別",
        "ticketContent": "工單內容",
        "ticketUrgencyLevel": { "low": "低", "med": "中", "high": "高" },
        "ticketStatus": "工單狀態",
        "ticketCreatedAt": "創建時間",
        "lastResponse": "最後回覆",
        "operate": "操作",
        "checkTicket": "查看工單",
        "closeTicket": "關閉工單",
        "userTickets": "歷史工單",
        "addTicket": "新建工單",
        "ticketActive": "未關閉",
        "ticketInActive": "已關閉",
        "form": {
            "ticketSubjectDescription": "請輸入工單主題",
            "ticketUrgencyDescription": "請選擇工單緊急程度",
            "ticketBody": "請輸入你想要解決的問題，盡量全面。",
            "ticketNotFinished": "請補全工單的信息再試"
        },
        "checkForm": "請檢查表單是否完整",
        "cancel": "取消",
        "submit": "提交",
        "commitNewTicketSuccess": "提交新的工單成功",
        "commitNewTicketFailure": "提交新的工單錯誤",
        "noReply": "還沒有回覆",
        "noTickets": "您還沒有提交過工單",
        "ticketCloseSuccess": "工單關閉成功",
        "ticketCloseFailure": "工單關閉失敗",
        "chatDialog": {
            "input": {
                "finished": "該工單已經結束",
                "inputHere": "輸入要發送的消息"
            },
            "send": "發送",
            "missingToken": "不可以創建Websocket連接，因為缺少Token。",
            "msgEmptyNotAllowed": "消息的內容不可以為空",
            "accessNotPermit": "非法訪問",
            "sendSuccess": "發送消息成功"
        }
    },
    "userActivation": {
        "activateLog": "激活紀錄",
        "description": "在此處您可以查看您的激活紀錄並設置解除綁定設備，也可以對每一個密鑰和激活紀錄設置備注信息。",
        "id": "#",
        "orderId": "訂單編號",
        "orderStatus": "訂單",
        "createdAt": "創建時間",
        "operate": "操作",
        "userId": "用戶Id",
        "email": "郵箱",
        "keyId": "密鑰Id",
        "isBind": "是否激活",
        "active": "有效",
        "inactive": "無效",
        "requestAt": "請求時間",
        "clientVersion": "客戶端",
        "osType": "操作系統",
        "remark": "備注",
        "noRemark": "無備注",
        "showDetail": "查看詳情",
        "actions": "操作",
        "details": "細節",
        "keyContent": "密鑰內容",
        "keyGeneratedAt": "密鑰生成時間",
        "activateRequestAt": "激活請求時間",
        "useIssueOccur": "使用遇到問題？",
        "chatWithUs": "聯繫我們",
        "cancelBind": "取消綁定",
        "alterRemark": "修改備注",
        "commitRemark": "提交",
        "updateSuccess": "更新成功",
        "setRemark": "在這裡設置備注信息"
    },
    "userAppDownload": {
        "title": "APP下載",
        "description": "",
        "common": {
            "title": "下載我們的應用程序",
            "shallow2": "為不同的客戶端獲取我們的應用程序",
            "shallow": "使用我們的應用程序您可以更便捷地訪問我們的服務，免去每次在瀏覽器繁瑣的操作；您可在在文檔中找到詳盡的安裝需求和使用教程，包括運行環境、報錯解決等，如果您遇到了其他問題請聯繫我們的技術支持。",
            "noDownload": "很抱歉暫時還沒有提供下載，請您稍後再試，如有問題請提交工單以聯繫我們的支持服務。"
        },
        "suffix": {
            "p1": "*macOS平台的應用程序請使用macOS14及以上並配合Apple晶片（M1或更高）。",
            "p2": "該App所展示的信息和使用請遵守當地的使用法規，是否允許使用該App也應該取決於您公司的IT規範。"
        },
        "card": {
            "common": { "welcome": "歡迎來到" },
            "mobile": {
                "designFor": "為移動端設計",
                "designShallow": "您可以在這裡獲取我們的移動端應用程序",
                "iosDownloadShallow": "下載 IOS 客戶端",
                "androidDownloadShallow": "下載 Android 客戶端"
            },
            "desktop": {
                "designFor": "為桌面端設計",
                "designShallow": "您可以在這裡獲取我們的桌面動端應用程序",
                "windowsDownloadShallow": "下載 Windows 客戶端",
                "osxDownloadShallow": "下載 macOS 客戶端",
                "linuxDownloadShallow": "下載 Linux 客戶端"
            }
        },
        "downloadDisabledHint": "很抱歉，App下載暫時不可用或被管理員禁用，如您有需要請在工單處聯繫我們的技術已獲取支持。",
        "windows": {
            "title": "Windows NT",
            "shallow": "該客戶端適用於NT內核的Windows操作系統，請查看文檔頁面以獲取兼容性支持。"
        },
        "osx": {
            "title": "macOS (OSX)",
            "shallow": "該客戶端適用於Darwin內核的macOS(OSX)操作系統，請查看文檔頁面以獲取兼容性支持。"
        },
        "linux": {
            "title": "Linux",
            "shallow": "該客戶端適用於Linux內核的各類發行版，由於發行版系列不同請查看文檔頁面以獲取兼容性支持。"
        },
        "android": {
            "title": "Android",
            "shallow": "該客戶端適用於搭載了Google Android操作系統的移動設備，請查看文檔頁面以獲取兼容性支持。"
        },
        "ios": {
            "title": "IOS",
            "shallow": "該客戶端適用於搭載了Apple IOS操作系統的移動設備，請查看文檔頁面以獲取兼容性支持。"
        }
    },
    "welcome": {
        "A": {
            "aboutUs": "關於我們",
            "pricing": "定價",
            "login": "登錄",
            "register": "註冊帳號",
            "welcomeTo": "歡迎來到",
            "welcomeToSub": "穿過縣境上長長的隧道，便是雪國。 夜空下，大地一片瑩白，火車在信號所前停下來。 \"在這裡川端康成用幾近吝嗇的簡潔文字，拉開了《雪國》的序幕。",
            "startingUse": "開始使用",
            "whyUs": "為什麼選擇我們",
            "whyUsSub": "\"穿過縣境上長長的隧道，便是雪國。夜空下，大地一片瑩白，火車在信號所前停下來。\"在這裡川端康成用幾近吝嗇的簡潔文字，拉開了《雪國》的序幕。",
            "browseSafe": "瀏覽安全",
            "browseSafeSub": "優秀的防火牆過濾系統能有效防禦網路釣魚和惡意網站",
            "encrypt": "端到端加密",
            "encryptSub": "雙向 SSL 和端對端加密保護您的隱私安全，即使是服務器也無法讀取您的信息",
            "mgr": "高效管理",
            "mgrSub": "一個用戶介面管理所有密鑰，管理功能完善豐富，無須擔心訂閱洩露問題",
            "fast": "方便快捷",
            "fastSub": "提供完整的 API 文檔供 WebApp 或是嵌入到第三方軟件中",
            "fastLink": "快速連結",
            "subscribeUs": "關注我們",
            "filingsCode": "備案號 {code}"
        }
    },
    "pagination": {
        "perPage10": "10 條數據/頁",
        "perPage20": "20 條數據/頁",
        "perPage50": "50 條數據/頁",
        "perPage100": "100 條數據/頁"
    },
    "modal": { "cancel": "取消", "confirm": "確認" },
    "week": {
        "monday": "週一",
        "tuesday": "週二",
        "wednesday": "週三",
        "thursday": "週四",
        "friday": "週五",
        "saturday": "週六",
        "sunday": "週日"
    },
    "period": {
        "month": "月付",
        "quarter": "季付",
        "halfYear": "半年付",
        "year": "年付"
    },
    "operate": {
        "cancel": "取消",
        "confirm": "確認",
        "update": "更新",
        "add": "添加"
    },
    "notFound": {
        "title": "404 頁面不存在",
        "description": "我們無法找到您請求的頁面，它可能已經被刪除或連結有誤。 如果您認為這是一個錯誤，請提交工單以聯繫我們。",
        "p1": "將在 {sec}s 後返回主頁，如果您的瀏覽器沒有響應，請點擊以下按鈕。",
        "manualBack": "返回主頁"
    },
    "forbidden": {
        "title": "403 無權限",
        "description": "您可能沒有足夠的權限來訪問本頁面。 如果您認為這是一個錯誤，請提交工單以聯繫我們。"
    }
}
