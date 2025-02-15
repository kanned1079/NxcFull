export default {
    commonHeader: {
        menuTxt: '菜單',
        userData: '用戶資料',
        editUserData: '編輯用戶資料',
        logout: '退出登陸',
    },
    commonAside: {
        admin: {
            dashboard: '儀表板',
            queueMonit: '服務端監控',
            settings: '設置',
            systemConfig: '系統設置',
            paymentConfig: '支付設置',
            themeConfig: '主題設置',
            server: '服務器',
            privilege: '權限組管理',
            finance: '財務',
            subscription: '訂閱管理',
            coupon: '優惠券管理',
            order: '訂單管理',
            activate: '激活紀錄',
            key: '密鑰管理',
            user: '用戶',
            userMgr: '用戶管理',
            notice: '公告管理',
            ticket: '工單管理',
            doc: '知識庫管理',

        },
        user: {
            dashboard: '儀表板',
            document: '使用文檔',
            app: 'APP下载',
            subscription: '訂閱',
            purchase: '購買訂閱',
            surplus: '我的密鑰',
            fiance: '財務',
            topUp: '充值',
            myOrder: '我的訂單',
            myInvite: '我的邀請',
            user: '用戶',
            profile: '個人中心',
            support: '我的工單',
            activateLog: '激活紀錄'
        }
    },
    adminViews: {
        common: {
            fetchDataSuccess: '獲取數據成功',
            fetchDataFailure: '獲取數據失敗請稍後再試',
            addSuccess: '添加成功',
            addFailure: '添加失敗請稍後嘗試',
            updateSuccess: '修改成功',
            updateFailure: '修改失敗請稍後再試',
            deleteSuccess: '刪除成功',
            deleteFailure: '刪除失敗請稍後再試',
        },
        login: {
            secureCard: {
                title: '安全检查',
                securePath: '安全路径',
                hint: '为保障系统安全，您需输入安全路径方可进入管理员登录页面。请在下方输入框中输入安全路径，成功验证安全路径后您可选择选择保存，以便后续快捷登录。',
                placeholder: '请输入安全路径',
                checkBtn: '检查',
                rememberPath: '记住安全路径',
            },
            card: {
                back: '返回首页',
                toAdminCenter: '登录到管理中心',
                emailAddr: '管理员邮箱',
                emailInputArea: {
                    title: '管理员邮箱',
                    placeholder: '邮箱地址',
                },
                passwordInputArea: {
                    title: '管理员密码',
                    placeholder: '密码',
                },
                login: '登入',
                forgetPassword: '忘记密码',
                formNotPassed: '表单验证不通过'
            },
            mention: {
                title: '提示',
                description: '此页面是管理员页面，只有管理员才能访问，如果您没有管理员权限或意外来到此处，请点击下面的链接返回主页。',
            },
            message: {
                passwordErr: '密码不正确',
                adminNotExist: '管理员不存在',
                authPassed: '验证通过',
                authFailure: '验证失败',
                otherErr: '其他錯誤',
                pathCheckPassed: '安全路径检查通过',
                pathCheckFailure: '安全路径不正确',
                rememberSecureMention: '为了保证后台管理的安全性，如果这不是您的私人电脑请不要勾选。'
            }
        },
        summary: {
            cockpit: '儀表板',
            systemConfig: '系統設置',
            paymentConfig: '支付設置',
            planMgr: '訂閱管理',
            userMgr: '用戶管理',
            orderMgr: '訂單管理',
            keyMgr: '密鑰管理',
            incomeText: '昨日收入 / 當月收入',
            pendingTicket: '您有 {nums} 條待處理工單',
            toHandle: '去查看',
            userCard: {
                title: '用戶概覽',
                allRegisteredUsers: '總註冊用戶',
                activeUsers: '活躍用戶',
                inactiveUsers: '非活躍用戶',
                blockedUsers: '封禁或註銷'
            },
            general: {
                title: '一般',
                localTime: '瀏覽器時間',
                osType: '操作系統類型',
                appName: '應用名稱',
                appUrl: '應用網址',
                currency: '貨幣單位',
                allowRegister: '允許註冊'
            },
            system: {
                title: '系統配置',
                serverTime: '',
                gatewayStatus: '網關狀態',
                dbStatus: '數據庫狀態',
                redisStatus: 'Redis狀態',
                etcdStatus: 'Etcd狀態',
                serverOsType: '服務器操作系統類型',

            }
        },
        queueMonit: {
            title: '服務端監控',
            headerHint: '請勿長時間停留在此頁面，在網絡高峰時段頻繁查詢將些許影響網關性能和數據庫吞吐量。',
            latency: {
                title: '服務器延遲',
                retry: '再次測試',
                hint: '*這裡的請求指的是客戶端發送請求到服務器後，服務器給予成功響應到客戶端所需要的時間。',
                unit: '毫秒',
                level: {
                    l1: {
                        title: '優秀',
                        description: '這是很優秀的網絡情況，幾乎不會感覺到延遲。',
                    },
                    l2: {
                        title: '正常',
                        description: '這是大多數時候的網絡情況，用戶幾乎感覺不到延遲。',
                    },
                    l3: {
                        title: '較差',
                        description: '在在此情况下用户可能感觉到轻微卡顿或延迟。',
                    },
                    l4: {
                        title: '差',
                        description: '可以感受到明显的延迟，有影响用户体验。',
                    },
                    l5: {
                        title: '非常差',
                        description: '出现明显的延迟并且加载速度变慢甚至无法刷新，严重影响用户互动和体验。',
                    },
                    offline: {
                        title: '服務器離線',
                        description: '無法連接到服務器或處理請求錯誤，請檢查是否配置正確。',
                    }
                },
            },
            api: {
                title: '最近7天API请求状态',
                items: {
                    ok: {
                        title: '成功 (StatusOK)',
                        unit: '次',
                    },
                    notFound: {
                        title: '状态路径错误 (404)',
                        unit: '次',
                    },
                    internalErr: {
                        title: '服务器内部错误 (500)',
                        unit: '次',
                    },
                    login2reg: {
                        title: '登录 / 注册',
                        unit: '次'
                    }
                },
            },
            log: {
                title: '日誌記錄',
                deleteLogMsg: '刪除日誌 {nums} 條',
                deleteLogErr: '刪除日誌失敗',
                logTableRows: '日誌表總記錄數',
                logTableSize: '日誌表佔用空間',
                unit: {
                    lines: '行',
                    mb: '兆字節',
                },
                deleteLog: '刪除日誌',
                exportCsv: '導出CSV',
                deleteLogHint: '這將執行刪除前一周前所有的日誌，刪除方式為硬刪除，所有被刪除的日誌將無法恢復，建議先保存備份。',
                table: {
                    id: '#',
                    method: '請求方式',
                    path: '請求路徑',
                    code: '狀態碼',
                    clientIp: '客戶端IP',
                    responseTime: '處理時間',
                    requestAt: '請求時間',
                }
            }
        },
        systemConfig: {
            title: '系統設置',
            common: {
                success: '更新配置成功',
                err: '更新配置失敗 ',
            },
            site: {
                common: {
                    title: '站点',
                },
                appName: {
                    title: '站点名称',
                    shallow: '用于显示需要站点名称的地方。',
                    placeholder: '站点名称'
                },
                appSubName: {
                    title: '站点副标题',
                    shallow: '一般显示在主要标题的下面。',
                    placeholder: '副标题'
                },
                appDescription: {
                    title: '站点描述',
                    shallow: '用于显示需要站点描述的地方。',
                    placeholder: '站点描述'
                },
                appUrl: {
                    title: '站点网址',
                    shallow: '当前网站最新网址，将会在邮件等需要用于网址处体现。',
                    placeholder: '站点网址'
                },
                forceHttps: {
                    title: '强制 HTTPS',
                    shallow: '当站点没有使用 HTTPS，CDN 或反代开启强制 HTTPS 时需要开启。'
                },
                logoUrl: {
                    title: 'LOGO',
                    shallow: '用于显示需要 LOGO 的地方。',
                    placeholder: 'logo 的 url 地址'
                },
                subscribeUrl: {
                    title: '订阅 URL',
                    shallow: '用于订阅所使用，留空则为站点 URL。如需多个订阅 URL 随机获取请使用逗号进行分割。',
                    placeholder: '订阅 url'
                },
                tosUrl: {
                    title: '用户条款(TOS)URL',
                    shallow: '用于跳转到用户条款(TOS)',
                    placeholder: '用户条款地址'
                },
                stopRegister: {
                    title: '停止新用户注册',
                    shallow: '开启后任何人都将无法进行注册。'
                },
                inviteRequire: {
                    title: '强制邀请',
                    shallow: '开启后当新用户注册时必需填写邀请码。'
                },
                trialSubscribe: {
                    title: '注册试用',
                    shallow: '选择需要试用的订阅，如果没有选项请先前往订阅管理添加。'
                },
                trialTime: {
                    title: '试用时间(小时)',
                    shallow: '新用户注册时订阅试用时间。'
                },
                currency: {
                    title: '货币单位',
                    shallow: '仅用于展示使用，更改后系统中所有的货币单位都将发生变更。',
                    placeholder: 'CNY'
                },
                currencySymbol: {
                    title: '货币符号',
                    shallow: '仅用于展示使用，更改后系统中所有的货币单位都将发生变更。',
                    placeholder: '¥'
                }
            },
            security: {
                common: {
                    title: '安全设置'
                },
                emailVerify: {
                    title: '邮箱验证',
                    description: '开启后将会强制要求用户进行邮箱验证。',
                },
                gmailAlias: {
                    title: '禁止使用Gmail多别名',
                    description: '开启后Gmail多别名将无法注册。',
                },
                safeMode: {
                    title: '安全模式',
                    description: '开启后除了站点URL以外的绑定本站点的域名访问都将会被403。',
                },
                adminPath: {
                    title: '后台路径',
                    description: '后台管理路径，修改后将会改变原有的admin路径。',
                    placeholder: 'https://x.com/logo.jpeg',
                },
                emailWhitelist: {
                    title: '邮箱后缀白名单',
                    description: '开启后在名单中的邮箱后缀才允许进行注册。',
                },
                recaptcha: {
                    title: '防机器人',
                    description: '开启后将会启用hCaptcha防止机器人。',
                },
                hCaptchaSiteKey: {
                    title: 'hCaptcha SiteKey',
                    description: '该SiteKey用于请求hCaptcha服务器来标识网站编号',
                    placeholder: 'a3ca066c-0ea0-42fe-bcd2-23f4ab48d528',
                },
                ipRegisterLimit: {
                    title: 'IP注册限制',
                    description: '开启后如果IP注册账户达到规则要求将会被限制注册，请注意IP判断可能因为CDN或前置代理导致问题。',
                },
                registerTimes: {
                    title: '次数',
                    description: '达到注册次数后开启惩罚。',
                    placeholder: '请输入',
                },
                lockTime: {
                    title: '惩罚时间(分钟)',
                    description: '需要等待惩罚时间过后才可以再次注册。',
                    placeholder: '请输入',
                },
            },
            frontend: {
                common: {
                    title: '个性化',
                },
                themePropColor: {
                    default: '默认',
                    darkBlueDay: '深蓝色',
                    milkGreenDay: '奶绿色',
                    bambooGreen: '若竹',
                    mistyPine: '雾松',
                    glacierBlue: '冰川蓝',
                    grayTheme: '灰色'
                },
                sidebarStyle: {
                    title: '边栏风格',
                    shallow: '设置侧边栏的颜色。',
                },
                headerStyle: {
                    title: '头部风格',
                    shallow: '设置顶部的颜色。',
                },
                themeColor: {
                    title: '主题色',
                    shallow: '设置整个网页的主题色。',
                },
                background: {
                    title: '背景',
                    shallow: '将会在后台登录页面进行展示。',
                    placeholder: 'https://x.com/logo.jpeg',
                },
            },
            inviteAndRebate: {
                common: {
                    title: '支付和返利',
                },
                inviteRebateEnable: {
                    title: '启用用户返利',
                    description: '如果启用则在被邀请的用户充值时，按照下面设置的充值比例给予用户返利。',
                },
                inviteRebateRate: {
                    title: '返利比例',
                    description: '设置返利的金额比例。',
                    placeholder: '请输入返利比例',
                },
                discountInfo: {
                    title: '优惠信息',
                    description: '设置优惠信息，它将会展示在充值页面顶部。',
                    placeholder: '设置优惠信息',
                },
                inviteInfo: {
                    title: '邀请信息',
                    description: '设置邀请信息，它将会展示在用户邀请页面，用于显示返利比例。',
                    placeholder: '设置返利信息',
                },
            },
            welcome: {
                common: {
                    title: '首页信息',
                },
                homeDescription: {
                    title: '首页描述',
                    description: '设置首页的简要描述内容。',
                    placeholder: '请输入首页描述内容',
                },
                whyChooseUs: {
                    title: '为什么选择我们',
                    description: '设置关于为什么选择我们的描述。',
                    placeholder: '请输入详细描述',
                },
                bilibiliLink: {
                    title: 'Bilibili 官方链接',
                    description: '设置 Bilibili 官方账号的链接地址。',
                    placeholder: 'https://space.bilibili.com/xxxx',
                },
                youtubeLink: {
                    title: 'YouTube 官方链接',
                    description: '设置 YouTube 官方账号的链接地址。',
                    placeholder: 'https://youtube.com/channel/xxxx',
                },
                instagramLink: {
                    title: 'Instagram 官方链接',
                    description: '设置 Instagram 官方账号的链接地址。',
                    placeholder: 'https://instagram.com/xxxx',
                },
                wechatLink: {
                    title: '微信公众账号链接',
                    description: '设置微信公众账号的链接地址。',
                    placeholder: '请输入微信公众链接',
                },
                filingNumber: {
                    title: '备案号',
                    description: '设置站点的备案号。',
                    placeholder: '如：粤ICP备12345678号',
                },
                pageSuffix: {
                    title: '站点后缀',
                    description: '设置站点名称后缀，用于标题显示。',
                    placeholder: '如：- 你的站点名称',
                },
            },
            sendMail: {
                common: {
                    title: '邮件设置',
                    warning: '如果您更改了本页面的配置，需要对队列服务及逆行重启。另外本页配置优先级高于.env中的邮件配置。',
                    sendTestMailTitle: '发送测试邮件',
                    sendTestMailShallow: '邮件将会发送到当前登录管理员的邮箱',
                    sendTestMailTo: '发送测试邮件到',
                    sending: '發送郵件中',
                    success: '成功',
                    receiverAddr: '收信地址',
                    senderHost: '發信服務器',
                    senderPort: '發信端口',
                    senderEncrypt: '發信加密方式',
                    senderUsername: '發信用戶名',
                    sendErr: '發送郵件失敗 ',
                },
                smtpServerAddress: {
                    title: 'SMTP 服务器地址',
                    shallow: '由邮件服务商提供的服务地址',
                    placeholder: '请输入服务器地址',
                },
                smtpServerPort: {
                    title: 'SMTP 服务端口',
                    shallow: '常见的端口有 25, 465, 587',
                    placeholder: '请输入端口号',
                },
                smtpEncryption: {
                    title: 'SMTP 加密方式',
                    shallow: '465 端口加密方式一般为 SSL，587 端口加密方式一般为 TLS',
                    placeholder: '请输入加密方式',
                },
                smtpAccount: {
                    title: 'SMTP 账号',
                    shallow: '由邮件服务商提供的账号',
                    placeholder: '请输入账号',
                },
                smtpPassword: {
                    title: 'SMTP 密码',
                    shallow: '由邮件服务商提供的密码',
                    placeholder: '请输入密码',
                },
                senderAddress: {
                    title: '发件地址',
                    shallow: '由邮件服务商提供的发件地址',
                    placeholder: '请输入发件地址',
                },
                emailTemplate: {
                    title: '邮件模版',
                    shallow: '你可以在文档查看如何自定义邮件模板',
                    placeholder: '请选择邮件模板',
                },
            },
            notice: {
                common: {
                    title: '通知设置'
                },
                displayName: {
                    title: '显示名称',
                    shallow: '仅用于前端页面显示',
                    placeholder: '通用通知接口1',
                },
                barkEndpoint: {
                    title: 'Bark 接入点',
                    shallow: 'Bark 服务器后端 API 地址',
                    placeholder: 'https://<ip>:<port>/<secret-key>',
                },
                barkGroup: {
                    title: 'Bark 群组',
                    shallow: '客户端显示的群组名称',
                    placeholder: 'web',
                },
            },
            appDownload: {
                common: {
                    title: 'APP',
                    hint: '用于自有客户端(APP)的版本管理及更新',
                },
                enabled: {
                    title: '開放下載',
                    shallow: '如果啟用則允許用戶訪問下載頁面',
                },
                windows: {
                    title: 'Windows',
                    shallow: 'Windows 端版本号及下载地址',
                    placeholder: 'https://xxxx.com/xxx.exe',
                },
                macos: {
                    title: 'macOS',
                    shallow: 'macOS 端版本号及下载地址',
                    placeholder: 'https://xxxx.com/xxx.dmg',
                },
                linux: {
                    title: 'Linux',
                    shallow: 'Linux 端版本号及下载地址',
                    placeholder: 'https://xxxx.com/xxx.deb',
                },
                android: {
                    title: 'Android',
                    shallow: 'Android 端版本号及下载地址',
                    placeholder: 'https://xxxx.com/xxx.apk',
                },
                ios: {
                    title: 'IOS',
                    shallow: 'IOS 端版本号及下载地址',
                    placeholder: 'https://xxxx.com/xxx.ipk',
                },
            },
        },
        payConfig: {
            title: '支付設置',
            description: '在這裡可以管理所有支持的付款方式，目前僅支持支付寶支付，但是您也可以先行配置其他支付方式，如沒有符合您的支付流程，可以等待項目進一步完善後在更新日誌中查看是否支持。',
            attention: {
                title: '注意事項',
                point1: '务必先配置支付方式的信息再进行启用，这真的很重要。',
                point2: '修改付款方式配置时，如果显示为"---"则代表该选项已经被设置且非空，如果需要重新修改直接输入新值保存即可。',
                // point3: '由于配置信息有三级缓存，如遇到失败时，可以在重启Redis*和order微服务*后再试。',
                // point3: '目前仅支持支付宝支付，但是您也可以先配置微信和ApplePay，等待项目进一步完善后可以在更新日志中查看是否支持。',
            },
            common: {
                detail: '{method} 配置',
                fillAttention: '為確保安全，不顯示詳細信息，重新填寫以創建或覆蓋已有配置。',
                discountAmount: '優惠金額（可以在系统设置的 “支付和返利” 中设置用户提示信息）',
                saveConfigBtnHint: '保存',
                cancelBtnHint: '取消',
                saveSuccess: '配置保存成功',
                alterSuccess: '配置修改成功',
                discountPlaceholder: '优惠金额（充值金额大于优惠金额时应用优惠）',
                saveOrAlterFailure: '未知错误 ',
            },
            alipay: {
                title: '支付寶',
                config: {
                    appId: '應用Id',
                    appPublicKeyCertContent: '應用公鑰證書內容',
                    appPrivateKey: '應用私鑰',
                    alipayRootCert: '支付寶根證書',
                    alipayPublicKeyCertContent: '支付寶公鑰證書內容',
                }
            },
            wechat: {
                title: '微信支付',
                config: {
                    mchId: '商戶Id',
                    mchCertSerial: '商戶證書序列號',
                    apiV3Key: 'API v3 Key',
                    privateKey: '私鑰',
                }
            },
            apple: {
                title: 'Apple Pay',
                config: {
                    issuerId: 'Issuer Id',
                    bundleId: 'Bundle Id',
                    privateKeyId: 'Private Key Id',
                    privateKeyContent: '私鑰內容',
                }
            },
            addPaymentMethod: '添加支付方式',
            enableBtnHint: '啟用',
            disableBtnHint: '禁用',
            setupPaymentMethod: '配置',
        },
        themeConfig: {
            title: '主題配置',
            using: '使用，',
            setAsCurrent: ''
        },
        groupMgr: {
            title: '權限組管理',
            description: '權限組是用來標示不同的訂閱等級，您可以將同等級別但是額度等有細微區別的訂閱計畫歸納在一個權限組中方便管理。',
            common: {
                addNewGroup: '新建權限組',
                alterGroupName: '修改權限組名稱',
                addConfirmed: '確認添加',
                alterConfirmed: '確認修改',
                cancel: '取消',
                addSuccess: '添加成功',
                addFailure: '添加失敗',
                alterSuccess: '修改權限組成功',
                alterFailure: '修改權限組失敗',
                delSuccess: '刪除成功',
                delFailure: '刪除失敗',
            },
            groupId: '權限組ID',
            groupName: '權限組名稱',
            groupPlaceholder: '輸入權限組名稱',
            userCount: '用戶數量',
            planCount: '訂閱數量',
            operate: '操作',
            editBtnHint: '編輯',
            deleteBtnHint: '刪除',
        },
        docMgr: {
            title: '知識庫管理',
            description: '在這裡您可以編寫給用戶的說明文檔，它可以是軟件的設計說明書、使用教程或注意事項等，其文檔的內容支持Markdown格式。',
            addDoc: '添加新文檔',
            addSuccess: '添加新文檔成功',
            addFailure: '添加文檔失敗',
            titleNotEmpty: '文檔的標題不能為空',
            table: {
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
            form: {
                add: '添加文檔',
                edit: '編輯文檔',
                cancel: '取消',
                confirm: '確認',
                addBtn: '添加',
                editBtn: '修改',
                title: {
                    title: '文檔標題',
                    placeholder: '輸入文檔的標題'
                },
                sort: {
                    title: '排序',
                    placeholder: '輸入文檔的排序級別 越高的數值代表優先級越高',
                },
                category: {
                    title: '分類',
                    placeholder: '文檔將按照該字段進行分類展示',
                },
                lang: {
                    title: '文檔語言',
                    placeholder: '選擇文檔語言',
                }
            }

        },
        planMgr: {
            title: '訂閱管理',
            description: '在這裡可以添加新的訂閱計畫、修改已有訂閱計畫的描述、價格、餘糧、其所属的权限组等。',
            addNewPlan: '添加新訂閱',
            table: {
                id: '#',
                isSale: '啟用銷售',
                isRenew: '允許續費',
                sort: '排序等級',
                group: '所屬權限組',
                name: '名稱',
                nums: '數量',
                residue: '餘量',
                operate: '操作',
                operateBtn: {
                    update: '修改',
                    delete: '刪除',
                }
            },
            mention: {
                common: {
                    success: '成功',
                    failure: '失敗',
                },
                addSuccess: '添加新訂閱成功',
                addFailure: '添加新訂閱失敗',
                updateSuccess: '更新訂閱成功',
                updateFailure: '更新訂閱失敗',
                deleteSuccess: '計畫刪除成功',
                deleteFailure: '計畫刪除失敗'
                // TVNTG-VFJQ3-FQXFP-DVCP6-D3VJ8
            },
            form: {
                title: '添加訂閱',
                items: {
                    name: {
                        title: '訂閱名稱',
                        placeholder: '輸入訂閱計畫的名稱',
                    },
                    describe: {
                        title: '訂閱描述',
                        placeholder: '輸入訂閱的描述（支持Markdown）'
                    },
                    isSale: {
                        title: '啟用銷售',
                    },
                    isRenew: {
                        title: '啟用續費',
                    },
                    group: {
                        title: '所屬權限組',
                        placeholder: '選擇所屬權限組',
                    },
                    capacityLimit: {
                        title: '最大允许用户数目',
                        placeholder: '最大允许用户数目',
                    },
                    sort: {
                        title: '排序',
                        placeholder: '用于前端排序',
                    },
                    periodPlaceholder: {
                        month: '输入月付价格',
                        quarter: '输入季付价格',
                        halfYear: '输入半年付价格',
                        year: '输入年付价格',
                    },
                }
            }
        },
        couponMgr: {
            title: '優惠券管理',
            description: '在這裡您可以為一些特定的節日等添加一些優惠券，它允許用戶在下單的時候使用並按照您設置的比例進行抵折優惠。',
            // addNewCoupon: '添加新優惠券',
            fetchErr: '獲取數據失敗請重試',
            fetchPlanKvFailure: '獲取訂閱計畫列表失敗',
            table: {
                id: '#',
                enabled: '是否啟用',
                name: '名稱',
                code: '券碼',
                residue: '剩餘數量',
                startTime: '啟用時間',
                endTime: '結束時間',
                actions: '操作',
                edit: '編輯',
                delete: '刪除'
            },
            modal: {
                newCoupon: '添加新優惠券',
                editCoupon: '修改優惠券信息',
                confirmAdd: '確認添加',
                confirmEdit: '確認修改',
                emptyNotAllow: '該項目是必填的',
                cancel: '取消',
                name: {
                    title: '優惠券名稱',
                    placeholder: '請輸入優惠券名稱',
                },
                code: {
                    title: '券碼',
                    placeholder: '自定義優惠券碼（留空隨即生成）',
                },
                percentOff: {
                    title: '優惠信息（百分比）',
                    placeholder: '輸入優惠百分比（如-12%OFF）',
                },
                activeRange: {
                    title: '優惠券可使用的有效期範圍',
                },
                capacity: {
                    title: '優惠券最大使用次數',
                    placeholder: '限制最大使用限制（為空則不限制）',
                },
                perUserLimit: {
                    title: '每個用戶可使用優惠券的次數',
                    placeholder: '限制每個用戶可使用的次數（為空則不限制）'
                },
                planLimit: {
                    title: '指定訂閱計畫',
                    placeholder: '限制指定訂閱計畫可以使用優惠（為空則不限制）'
                }
            }

        },
        orderMgr: {
            title: '訂單管理',
            description: '在這裡您可以檢視所有訂閱計畫的訂單，篩選不同用戶的訂單、手動對用戶的訂單處理通過等。',
            table: {
                id: '#',
                orderId: '訂單號',
                email: '用戶郵箱',
                status: {
                    title: '類型',
                    t1: '新購',
                    t2: '續費',
                    t3: '編輯',
                },
                planName: '計畫名稱',
                period: '週期',
                group: '權限組',
                amount: '實付金額',
                price: '原始價格',
                isSuccess: {
                    title: '訂單狀態',
                    cancelOrder: '取消訂單',
                    passOrder: '通過訂單',
                },
                createdAt: '訂單創建時間',
                action: {
                    title: '操作',
                    showDetail: '顯示細節'
                }
            },
            search: '查詢訂單',
            resetSearch: '重置查詢',
            searchModal: {
                email: {
                    title: '用戶郵箱',
                    placeholder: '輸入用戶郵箱（模糊搜索）',
                },
                sort: {
                    title: '排序算法',
                    placeholder: '選擇排序算法',
                    ASC: '升序',
                    DESC: '降序',
                }
            },
            tradeWaiting: '未支付',
            tradeFailure: '交易失敗',
            tradeSuccess: '成功',

        },
        userMgr: {
            userManager: "用戶管理",
            description: '你可以在這裡管理所有的用戶，包括員工和管理員，授予或取消管理權限、設定用戶餘額、重置密碼、手動添加新用戶等操作。',
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
            deleted: '注销',
            nullContent: "NULL",
            balance: "餘額",
            orderCount: "訂單數量",
            planCount: "計劃數",
            actions: "操作",
            updateSuccess: "更新成功",
            addUserSuccess: "添加新用戶成功",
            unknownError: "未知錯誤",
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
        },
        activation: {
            activateLog: '激活紀錄',
            description: '您可以檢視所有售出密鑰的具體激活情況，查看客户端的识别码、激活时间等。',
            click2getKey: '點擊以獲取密鑰內容',
            createdAt: '創建時間',
            turn2keyPage: '轉到密鑰',
            showKey: '顯示密鑰',
            email: '郵箱',
            key: '密鑰',
            search: '搜尋',
            filter: '篩選',
            filterAll: '全部',
            filterActive: '僅有效',
            sortAlgorithm: '排序算法',
            sortASC: '升序',
            sortDESC: '降序',
        },
        keysMgr: {
            keyMgr: '密鑰管理',
            enableKey: '啟用密鑰',
            disableKey: '禁用密鑰',
        },
        noticeMgr: {
            title: '公告管理',
            description: '在這裡可以對公告進行管理，啟用的公告將展示在用戶首頁的輪播圖中，可以設置的公告有優惠活動、節日通知、注意事項等。',
            addNotice: '添加公告',
            table: {
                id: '#',
                show: '是否顯示',
                title: '標題',
                createdAt: '創建時間',
                action: {
                    title: '操作',
                    edit: '編輯',
                    delete: '刪除'
                }
            }
        },
        // docMgr: {
        //
        // },
        adminTicket: {
            ticketMgr: '工單管理',
            description: '您可以在此處檢視到所有用戶提交的工單，工單有三個緊急程度，推薦您從緊急工單開始處理。',
            pendingTicket: '待處理工單',
            finishedTicket: '已完成工單',
        },
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
        microsoft: '以Microsoft帳戶繼續',
        apple: '以Apple帳戶繼續',
        google: '以Google帳戶繼續',
        backHomePage: '回到首頁',
        form: {
            emailRequire: '郵箱地址是必填的',
            passwordRequire: '你還沒有輸入密碼',
        },
        authPass: '驗證通過',
        loginFailure: '登錄失敗',
        checkForm: '請檢查表單',
        if2FaEnabledHint: '如果您啟用了兩步驗證（非必填）',
        reqErr: '請求出現錯誤請稍後再試',
        accountLocked: '您的帳戶可能被註銷或封禁，暫時無法繼續使用，如你仍然認為這是一個錯誤，請聯繫我們的技術支持。'
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
        unknownErr: '未知錯誤',
        verifyCodeErr: '驗證碼錯誤',
        verifyCodeExpireErr: '验证码错误或已過期，請重試或獲取新驗證碼。',
        thisMailAlreadyExist: '該郵箱已經被註冊',
        pageConfigFetchFailure: '配置獲取失敗請刷新重試',
        stopRegisterHint: '很抱歉，目前註冊功能已暫停。如閣下有需要，請稍後再試或联系我們的支持團隊獲取更多信息。感謝您的理解與支持。',
        passwordComplexRequirePart1: '* 密碼需要符合',
        passwordComplexRequirePart2: '複雜度要求',
        passwordComplexHint1: '1. 需要使用大寫字母、小寫字母、數字、特殊符號中的三種及以上',
        passwordComplexHint2: '2. 長度再10位以上',
        form: {
            checkForm: '請檢查表單內容',
            emailFormatErr: '郵箱地址格式不正確',
            gmailLimitErr: '不允許使用Gmail多別名',
            gmailDotNotAllowed: 'Gmail地址中不允許有"."',
            gmailPartLowerForced: 'Gmail地址的本地部分必須全部小寫',
            googlemailNotAllowed: '不支持使用googlemail地址',
            verifyCodeRequire: '請輸入驗證碼',
            verifyCodeFormatErr: '驗證碼格式不正確',
            passwordRequire: '請輸入密碼',
            passwordLengthRequire: '密碼長度必須大於或等於10位',
            passwordComplexRequire: '密碼必須包含大寫字母、小寫字母、數字、特殊符號中的至少三種',
            passwordAgainNotMatch: '兩次輸入的密碼不匹配',
            passwordAgainRequire: '請輸入確認密碼',
            inviteCodeRequire: '邀請碼是必填的',
        },
        hCaptcha: {
            passed: '人機驗證通過',
            expired: '已過期請重試',
            challengeExpired: '驗證超時',
            err: '其他錯誤',
        },
        allRightsReserved: '版權所有',
        securityAndLaws: '該網站受hCaptcha保護和驗證，請遵守當地法律。',
    },
    userSummary: {
        title: '儀表盤',
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
        appDown: {
            title: '應用下載',
            content: '下載我們不同平台的應用程序以獲得更好的體驗',
        },
        support: {
            title: '遇到問題',
            content: '遇到問題可以通過工單與我們的人機溝通',
        },
        haveTicket: '您有 {count} 條待處理的工單',
        toCheckTicket: '去查看',
        showAllKeys: '查看所有密鑰',
    },
    userDocument: {
        title: '使用文檔',
        description: '您可以在這裡查閱不同的文檔，包括但不限於網站的使用方法、注意事項、操作流程等，如果您發現文章中有錯誤，請提交工單。',
        searchPlaceholder: '請輸入要搜索的內容（模糊搜索）',
        searchBtn: '搜索',
        noContentTitle: '無結果',
        noContentTitleHint: '没有符合您搜索结果或语言的文档，嘗試換一個關鍵詞吧。',
    },
    newPurchase: {
        title: '購買訂閱',
        description: '您可以在這裡選擇最適合您的訂閱計畫，如果您的餘額不足請先充值，訂單將為您保留5分鐘。',
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
        title: '個人中心',
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
        faAuth: '兩步驗證2FA',
        faAuthHint: '兩步驗證是一種安全機制，增加了登入帳戶的保護層。用戶在輸入密碼後，還需完成第二步身份驗證，如輸入發送到手機的驗證碼、使用身份驗證應用程序生成的動態碼，或通過指紋等生物特徵確認。',
        faAuthStatus: '兩步驗證狀態',
        faEnabled: '已啟用',
        faNotEnabled: '未啟用',
        setup2Fa: '設置兩步驗證',
        disable2Fa: '取消兩步驗證',
        unknownErr: '未知錯誤',
        disable2FaCancelled: '已取消',
        closed2FaHint: '已經關閉兩步驗證，如有需要請重新添加。',
        setup2FaModal: {
            followStep: '根據提示在您的驗證器上加入',
            step1Title: '按照以下步驟以啟2FA驗證',
            step1Context1: '1. 您的移動設備上需要有一個通用的驗證器',
            step1Context2: '2. 點擊驗證器上的Scan按鈕來掃描此處的QR碼',
            step1Context3: '3. 該QR Code中包含有您的驗證信息和唯一密鑰，請妥善保存',
            step2Context1: '為了確保您的驗證器能夠正常使用，我們需要進行測試。',
            test2Fa: '測試',
            cancel: '取消',
        },
        deleteMyAccountModal: {
            title: '註銷帳戶',
            contentLine1: '帳號註銷是一個不可逆的操作。一旦您確認刪除，您將永久失去該帳號的訪問權限，這意味著您將無法再次登入，且與此帳號相關的所有數據，包括但不限於您的個人資訊、歷史記錄、收藏內容、購買記錄等，都將無法再訪問。',
            contentLine2: '如果您在我們的平台上有正在進行的業務，例如未完成的訂單、正在參與的活動、訂閱服務等，這些都將隨著帳號刪除而終止或取消，可能會給您帶來相應的損失。同時，您與其他用戶之間通過本平台建立的聯繫、互動資訊等也都將不復存在。',
            contentLine3: '請再次確認您的決定，如果您還有任何疑慮或問題，歡迎聯繫我們的客服，我們將竭誠為您解答。若您仍然希望刪除帳號，請點擊「確認刪除」按鈕。',
            inputHint1: '輸入 "',
            inputHint2: '" 以繼續。',
            confirmDelete: '確認刪除',
        },
        failure: '遇到錯誤',
        notLatestHint: '個人信息更新失敗，當前的數據可能不是最新的。',
        resetPassword: {
            previousPasswordRequire: '請輸入舊密碼',
            previousPasswordVerifiedFailure: '舊密碼驗證失敗',
            passwordRequire: '請輸入密碼',
            passwordConflict: '新密碼不能與先前的密碼相同',
            passwordLengthRequire: '密碼長度必須大於或等於10位',
            passwordComplexRequire: '密碼必須包含大寫字母、小寫字母、數字、特殊符號中的至少三種',
            passwordAgainNotMatch: '兩次輸入的密碼不匹配',
            passwordAgainRequire: '請輸入確認密碼',
            updatePasswordFailure: '更新密碼出錯',
        },
        secondary: {
            title: '基礎資料',
            card: {
                avatar: '用戶頭像',
                name: '用戶名',
                lastLoginAt: '上一次登錄時間',
                accountStatus: '帳戶狀態',
            },
            modify: {
                uploadIconHint: '上傳',
                alterAvatar: '新增/修改頭像',
                alterShallow: '* 點擊頭像或用戶名以添加或修改（設置後均不允許清除）',
                alterName: '修改用戶名',
                input: {
                    label: '用户名',
                    placeholder: '輸入用戶名',
                    spaceIsNotAllowed: '用戶名驗證不通過',
                    require: {
                        // p1: '只允許使用英文和數字',
                        p1: '不允許純數字且數字開頭',
                        p2: '不允許有空格',
                        p3: '長度大於三位',
                    }
                }
            },
            mention: {
                alterNameSuccess: '修改用戶名成功',
                alterNameErr: '修改用戶名失敗請稍後再試',
                newNameIsNotValid: '用戶名不合法',
                click2SetName: '點擊以設置用戶名',
                fetchAvatarErr: '獲取頭像數據失敗請稍後再試',
                alterAvatarErr: '修改頭像失敗請稍後再試',
                success: '成功',
                alterAvatarSuccess: '修改頭像成功，',
                uploadImageHint: '您可以上傳圖像作為您的新頭像',
                imageRequire: {
                    title: '注意',
                    p1: '您可以上傳 *.jpg(jpeg), *.png, *.webp 等主流格式。',
                    p2: '圖像的長寬比為1:1（正方形），如果為其他比例，將會被居中裁切為正方形，多餘的部分將會被刪除。',
                    p3: '您上傳的圖像的大小將會被設置為160px。',
                },
                click2Upload: '點擊或拖動文件到該區域來上傳',
                uploadWarn: '*請不要上傳敏感數據，比如您的銀行卡、信用卡、密碼和安全碼等。'
            }
        }
    },
    // tickets: {
    //     userTickets: '歷史工單',
    //     addTicket: '新建工單',
    // },
    userKeys: {
        myKeys: '我的密鑰',
        description: '您可以檢視您的所有密鑰的活動情況、到期日期等，如需要為密鑰設置備注，請前往激活紀錄頁面。',
        noKeys: '您還沒有有效的購買紀錄',
        keyDetail: '密鑰細節',
        keyId: '密鑰ID',
        orderId: '链接订单ID',
        clientId: '激活客戶端ID',
        active: '在有效期內',
        inActive: '已過期',
        valid: '密鑰狀態正常',
        invalid: '密鑰已被禁用',
        isUsed: '已激活使用',
        noUsed: '還未使用',
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
        description: '您所有的訂單將在此處展示，如果您有未支付的訂單將會展示在頂部，您可以點擊繼續支付或取消訂單，已經完成的訂單您可以在此處檢視訂單細節。',
        orderId: '#',
        planName: '訂閱名稱',
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
        },
        orderStatusTags: {
            success: '成功',
            cancelled: '失敗',
            notPay: '未支付'
        },
        orderCancelled: '訂單已取消',
        unknownErr: '未知錯誤',

    },

    userTopUp: {
        topUp: '帳戶充值',
        description: '您可以在這裡進行帳戶充值，支持自定義充值金額，您也可以關注下方是否有優惠信息展示，使用提及的支付方式以獲得優惠。',
        chooseTopUpAmount: '選擇充值金額',
        quickSelect: '快速選擇',
        customAmount: '自定義金額',
        maxAmount: '最大金額: 10,000,000',
        amountInputPlaceHolder: '輸入要充值的金額',
        otherAmount: '其他金額',
        payMethod: '支付方式',
        wechat: '微信支付',
        alipay: '支付寶',
        apple: 'Apple Pay',
        yourAmount: '您的金額',
        discount: '優惠',
        accountBalance: '帳戶餘額',
        balanceResult: '餘額合計',
        commitTopUp: '提交',
        payMethodNotAllow: '支付方式不可用請選擇其他',
        topUpIssueOccur: '充值遇到問題？',
        payIssueOccur: '支付遇到問題？',
        chatWithUs: '聯繫客服',
        pay: '支付',
        qrCodeScannedSuccess: 'QR碼掃描成功',
        orClickToApp: '或點擊條轉到App繼續',
        topUpSuccess: '充值成功',
        thankU: '感謝您的支持'
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
        faCodeManage: '邀請碼管理',
        email: '郵箱',
        createdAt: '註冊時間',
        createFaCodeFailure: '创建失败',
        linkCopiedSuccess: '链接复制成功',
        generateFaCode: '生成邀請碼',
        flushFaCode: '刷新邀請碼',
        faCode: '邀請碼',
        noFaCode: '你還沒有邀請碼，請先生成。',
        faLink: '邀請連結',
        generateFaCodePlease: '請先生成邀請碼',
        usersMyInvited: '我邀請的用戶',
        generateCodeConfirm: '確認生成/刷新',
        generateCodeHint: '請注意，邀請碼創建後不可關閉。',
        positiveClick: '確認',
        negativeClick: '取消',
    },
    userTickets: {
        description: '如果您遇到使用上的問題，請在此處提交工單，我們的技術支持和客服看到後將會進行回覆並在下方表格處標出具體的回覆時間，您可以點擊查看工單和我們交流。',
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
        checkForm: '請檢查表單是否完整',
        cancel: '取消',
        submit: '提交',
        commitNewTicketSuccess: '提交新的工單成功',
        commitNewTicketFailure: '提交新的工單錯誤',
        noReply: '還沒有回覆',
        noTickets: '您還沒有提交過工單',
        ticketCloseSuccess: '工單關閉成功',
        ticketCloseFailure: '工單關閉失敗',
        chatDialog: {
            input: {
                finished: '該工單已經結束',
                inputHere: '輸入要發送的消息',
            },
            send: '發送',
            missingToken: '不可以創建Websocket連接，因為缺少Token。',
            msgEmptyNotAllowed: '消息的內容不可以為空',
            accessNotPermit: '非法訪問',
            sendSuccess: '發送消息成功',
        }
    },
    userActivation: {
        activateLog: '激活紀錄',
        description: '在此處您可以查看您的激活紀錄並設置解除綁定設備，也可以對每一個密鑰和激活紀錄設置備注信息。',
        id: '#',
        orderId: '訂單編號',
        orderStatus: '訂單',
        createdAt: '创建时间',
        operate: '操作',
        userId: '用戶Id',
        email: '郵箱',
        keyId: '密钥Id',
        isBind: '是否激活',
        active: '有效',
        inactive: '無效',
        requestAt: '請求時間',
        clientVersion: '客戶端',
        osType: '操作系統',
        remark: '備注',
        noRemark: '無備注',
        showDetail: '查看詳情',
        actions: '操作',
        details: '細節',
        keyContent: '密鑰內容',
        keyGeneratedAt: '密鑰生成時間',
        activateRequestAt: '激活請求時間',
        useIssueOccur: '使用遇到問題？',
        chatWithUs: '聯繫我們',
        cancelBind: '取消綁定',
        alterRemark: '修改備注',
        commitRemark: '提交',
        updateSuccess: '更新成功',
    },
    userAppDownload: {
        title: 'APP下載',
        description: '',
        common: {
            title: '下載我們的應用程序',
            shallow2: '為不同的客戶端獲取我們的應用程序',
            shallow: '使用我們的應用程序您可以更便捷地訪問我們的服務，免去每次在瀏覽器繁瑣的操作；您可在在文檔中找到詳盡的安裝需求和使用教程，包括運行環境、報錯解決等，如果您遇到了其他問題請聯繫我們的技術支持。',
            noDownload: '很抱歉暂时还没有提供下载，请您稍后再试，如有问题请提交工单以联系我们的支持服务。'
        },
        suffix: {
            p1: '*macOS平台的應用程序請使用macOS14及以上並配合Apple晶片（M1或更高）。',
            p2: '该App所展示的信息和使用请遵守当地的使用法规，是否允许使用该App也应该取决于您公司的IT规范。',
        },
        card: {
            common: {
                welcome: '歡迎來到',
            },
            mobile: {
                designFor: '為移動端設計',
                designShallow: '您可以在這裡獲取我們的移動端應用程序',
                iosDownloadShallow: '下載 IOS 客戶端',
                androidDownloadShallow: '下載 Android 客戶端',
            },
            desktop: {
                designFor: '為桌面端設計',
                designShallow: '您可以在這裡獲取我們的桌面動端應用程序',
                windowsDownloadShallow: '下載 Windows 客戶端',
                osxDownloadShallow: '下載 macOS 客戶端',
                linuxDownloadShallow: '下載 Linux 客戶端',
            },
        },
        downloadDisabledHint: '很抱歉，App下载暂时不可用或被管理员禁用，如您有需要请在工单处联系我们的技术已获取支持。',
        windows: {
            title: 'Windows NT',
            shallow: '该客户端适用于NT内核的Windows操作系统，请查看文档页面以获取兼容性支持。',
        },
        osx: {
            title: 'macOS (OSX)',
            shallow: '该客户端适用于Darwin内核的macOS(OSX)操作系统，请查看文档页面以获取兼容性支持。',
        },
        linux: {
            title: 'Linux',
            shallow: '该客户端适用于Linux内核的各类发行版，由于发行版系列不同请查看文档页面以获取兼容性支持。',
        },
        android: {
            title: 'Android',
            shallow: '该客户端适用于搭载了Google Android操作系统的移动设备，请查看文档页面以获取兼容性支持。',
        },
        ios: {
            title: 'IOS',
            shallow: '该客户端适用于搭载了Apple IOS操作系统的移动设备，请查看文档页面以获取兼容性支持。',
        },
    },
    welcome: {
        A: {
            aboutUs: '關於我們',
            pricing: '定價',
            login: '登錄',
            register: '註冊帳號',
            welcomeTo: '歡迎來到',
            welcomeToSub: '穿过县境上长长的隧道，便是雪国。夜空下，大地一片莹白，火车在信号所前停下来。”在这里川端康成用几近吝啬的简洁文字，拉开了《雪国》的序幕。',
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

    pagination: {
        perPage10: '10 條數據/頁',
        perPage20: '20 條數據/頁',
        perPage50: '50 條數據/頁',
        perPage100: '100 條數據/頁',
    },
    modal: {
        cancel: '取消',
        confirm: '确认',
    },
    week: {
        monday: '週一',
        tuesday: '週二',
        wednesday: '週三',
        thursday: '週四',
        friday: '週五',
        saturday: '週六',
        sunday: '週日'
    },
    period: {
        month: '月付',
        quarter: '季付',
        halfYear: '半年付',
        year: '年付'
    },
    operate: {
        cancel: '取消',
        confirm: '确认',
        update: '更新',
        add: '添加'
    }
}