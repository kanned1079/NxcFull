export default {
    "commonHeader": {
        "menuTxt": "菜单",
        "userData": "用户资料",
        "editUserData": "编辑用户资料",
        "logout": "退出登陆"
    },
    "commonAside": {
        "admin": {
            "dashboard": "仪表板",
            "queueMonit": "服务端监控",
            "settings": "设置",
            "systemConfig": "系统设置",
            "paymentConfig": "支付设置",
            "themeConfig": "主题设置",
            "server": "服务器",
            "privilege": "权限组管理",
            "finance": "财务",
            "subscription": "订阅管理",
            "coupon": "优惠券管理",
            "order": "订单管理",
            "activate": "激活纪录",
            "key": "密钥管理",
            "user": "用户",
            "userMgr": "用户管理",
            "notice": "公告管理",
            "ticket": "工单管理",
            "doc": "知识库管理"
        },
        "user": {
            "dashboard": "仪表板",
            "document": "使用文档",
            "app": "APP下载",
            "subscription": "订阅",
            "purchase": "购买订阅",
            "surplus": "我的密钥",
            "fiance": "财务",
            "topUp": "充值",
            "myOrder": "我的订单",
            "myInvite": "我的邀请",
            "user": "用户",
            "profile": "个人中心",
            "support": "我的工单",
            "activateLog": "激活纪录"
        }
    },
    "adminViews": {
        "common": {
            "fetchDataSuccess": "获取数据成功",
            "fetchDataFailure": "获取数据失败请稍后再试",
            "addSuccess": "添加成功",
            "addFailure": "添加失败请稍后尝试",
            "updateSuccess": "修改成功",
            "updateFailure": "修改失败请稍后再试",
            "deleteSuccess": "删除成功",
            "deleteFailure": "删除失败请稍后再试",
            "confirm": "确认",
            "cancel": "取消",
            "success": "操作成功",
            "failure": "操作失败",
            "NotAllowed": "非法格式",
            "checkForm": "请检查表单",
            "unknownErr": "未知错误",
            "dialog": { "delete": "您确认要删除吗" },
            "yes": "是",
            "no": "否"
        },
        "login": {
            "secureCard": {
                "title": "安全检查",
                "securePath": "安全路径",
                "hint": "为保障系统安全，您需输入安全路径方可进入管理员登录页面。 请在下方输入框中输入安全路径，成功验证安全路径后您可选择选择保存，以便后续快捷登录。",
                "placeholder": "请输入安全路径",
                "checkBtn": "检查",
                "rememberPath": "记住安全路径"
            },
            "card": {
                "back": "返回首页",
                "toAdminCenter": "登录到管理中心",
                "emailAddr": "管理员邮箱",
                "emailInputArea": { "title": "管理员邮箱", "placeholder": "邮箱地址" },
                "passwordInputArea": { "title": "管理员密码", "placeholder": "密码" },
                "login": "登入",
                "forgetPassword": "忘记密码",
                "formNotPassed": "表单验证不通过"
            },
            "mention": {
                "title": "提示",
                "description": "此页面是管理员页面，只有管理员才能访问，如果您没有管理员权限或意外来到此处，请点击下面的链接返回主页。"
            },
            "message": {
                "passwordErr": "密码不正确",
                "adminNotExist": "管理员不存在",
                "noPrivilege": "无权限访问",
                "authPassed": "验证通过",
                "authFailure": "验证失败",
                "otherErr": "其他错误",
                "pathCheckPassed": "安全路径检查通过",
                "pathCheckFailure": "安全路径不正确",
                "rememberSecureMention": "为了保证后台管理的安全性，如果这不是您的私人电脑请不要勾选。"
            }
        },
        "summary": {
            "cockpit": "仪表板",
            "systemConfig": "系统设置",
            "paymentConfig": "支付设置",
            "planMgr": "订阅管理",
            "userMgr": "用户管理",
            "orderMgr": "订单管理",
            "keyMgr": "密钥管理",
            "incomeText": "昨日收入 / 当月收入",
            "pendingTicket": "您有 {nums} 条待处理工单",
            "toHandle": "去查看",
            "apiAccessCard": "一周内API接口访问次数",
            "apiAccessCardHint": "此数据仅供您了解当前的API访问情况，并不代表您的服务器性能。",
            "incomeWeek": "一周内收入金额",
            "incomeCardHint": "这里展示了一周内的收入金额图表，如果缓存被清空则会导致显示不准确。",
            "core": "核心",
            "reqErr": "遇到错误",
            "reqErrHint": "在获取概览信息时遇到错误，导致本次请求无法完成，因此图表暂不能显示，请您稍后再试。",
            "userCard": {
                "title": "用户概览",
                "allRegisteredUsers": "总注册用户",
                "activeUsers": "活跃用户",
                "inactiveUsers": "非活跃用户",
                "blockedUsers": "封禁或注销"
            },
            "general": {
                "title": "一般",
                "localTime": "浏览器时间",
                "osType": "操作系统类型",
                "appName": "应用名称",
                "appUrl": "应用网址",
                "currency": "货币单位",
                "allowRegister": "允许注册"
            },
            "system": {
                "title": "系统配置",
                "axiosAddr": "HTTP后端地址",
                "wsAddr": "Websocket后端地址",
                "serverTime": "服务器时间",
                "uptime": "运行时长",
                "gatewayStatus": "API网关状态",
                "dbStatus": "数据库状态",
                "redisStatus": "Redis状态",
                "serverOsType": "服务器操作系统类型",
                "serverOsArch": "服务器操作系统架构",
                "runMode": "运行模式",
                "cpuNums": "网关服务器CPU核心数",
                "numCgoCall": "垃圾回收次数",
                "time": "次",
                "paymentMethods": "启用的支付方式",
                "runOK": "运行正常",
                "runErr": "运行异常",
                "checkServer": "请检查您后端服务器的环境配置",
                "stopRegisterHint": "您似乎禁用了新用户注册",
                "toSetting": "转到设置"
            }
        },
        "queueMonit": {
            "title": "服务端监控",
            "headerHint": "请勿长时间停留在此页面，在网络高峰时段频繁查询将些许影响网关性能和数据库吞吐量。",
            "latency": {
                "title": "服务器延迟",
                "retry": "再次测试",
                "hint": "*这里的请求指的是客户端发送请求到服务器后，服务器给予成功响应到客户端所需要的时间。",
                "unit": "毫秒",
                "level": {
                    "l1": {
                        "title": "优秀",
                        "description": "这是很优秀的网络情况，几乎不会感觉到延迟。"
                    },
                    "l2": {
                        "title": "正常",
                        "description": "这是大多数时候的网络情况，用户几乎感觉不到延迟。"
                    },
                    "l3": {
                        "title": "较差",
                        "description": "在在此情况下用户可能感觉到轻微卡顿或延迟。"
                    },
                    "l4": {
                        "title": "差",
                        "description": "可以感受到明显的延迟，有影响用户体验。"
                    },
                    "l5": {
                        "title": "非常差",
                        "description": "出现明显的延迟并且加载速度变慢甚至无法刷新，严重影响用户互动和体验。"
                    },
                    "offline": {
                        "title": "服务器离线",
                        "description": "无法连接到服务器或处理请求错误，请检查是否配置正确。"
                    }
                }
            },
            "api": {
                "title": "最近7天API请求状态",
                "items": {
                    "ok": { "title": "成功 (StatusOK)", "unit": "次" },
                    "notFound": { "title": "状态路径错误 (404)", "unit": "次" },
                    "unAuthorized": { "title": "未授权访问 (401)", "unit": "次" },
                    "login2reg": { "title": "登录 / 注册", "unit": "次" }
                }
            },
            "log": {
                "title": "日志记录",
                "deleteLogMsg": "删除日志 {nums} 条",
                "deleteLogErr": "删除日志失败",
                "logTableRows": "日志表总记录数",
                "logTableSize": "日志表占用空间",
                "unit": { "lines": "行", "mb": "兆字节" },
                "deleteLog": "删除日志",
                "exportCsv": "导出CSV",
                "deleteLogHint": "这将执行删除前一周前所有的日志，删除方式为硬删除，所有被删除的日志将无法恢复，建议先保存备份。",
                "warn": {
                    "title": "您确认要删除日志吗？",
                    "description": "将立即删除条目，您不能撤销此操作。"
                },
                "export": {
                    "title": "导出csv文件",
                    "description": "这将会导出下面的表格为csv文件并下载到本地，如果您没有开启下载权限，下载有可能会失败，点击确认按钮以执行导出。"
                },
                "table": {
                    "id": "#",
                    "method": "请求方式",
                    "path": "请求路径",
                    "code": "状态码",
                    "clientIp": "客户端IP",
                    "responseTime": "处理时间",
                    "requestAt": "请求时间"
                }
            }
        },
        "systemConfig": {
            "title": "系统设置",
            "common": { "success": "更新配置成功", "err": "更新配置失败" },
            "site": {
                "common": { "title": "站点" },
                "appName": {
                    "title": "站点名称",
                    "shallow": "用于显示需要站点名称的地方。",
                    "placeholder": "站点名称"
                },
                "appSubName": {
                    "title": "站点副标题",
                    "shallow": "一般显示在主要标题的下面。",
                    "placeholder": "副标题"
                },
                "appDescription": {
                    "title": "站点描述",
                    "shallow": "用于显示需要站点描述的地方。",
                    "placeholder": "站点描述"
                },
                "appUrl": {
                    "title": "站点网址",
                    "shallow": "当前网站最新网址，将会在邮件等需要用于网址处体现。",
                    "placeholder": "站点网址"
                },
                "forceHttps": {
                    "title": "强制 HTTPS",
                    "shallow": "当站点没有使用 HTTPS，CDN 或反代开启强制 HTTPS 时需要开启。"
                },
                "logoUrl": {
                    "title": "LOGO",
                    "shallow": "用于显示需要 LOGO 的地方。",
                    "placeholder": "logo 的 url 地址"
                },
                "subscribeUrl": {
                    "title": "订阅 URL",
                    "shallow": "用于订阅所使用，留空则为站点 URL。 如需多个订阅 URL 随机获取请使用逗号进行分割。",
                    "placeholder": "订阅 url"
                },
                "tosUrl": {
                    "title": "用户条款(TOS)URL",
                    "shallow": "用于跳转到用户条款(TOS)",
                    "placeholder": "用户条款地址"
                },
                "stopRegister": {
                    "title": "停止新用户注册",
                    "shallow": "开启后任何人都将无法进行注册。"
                },
                "inviteRequire": {
                    "title": "强制邀请",
                    "shallow": "开启后当新用户注册时必需填写邀请码。"
                },
                "trialSubscribe": {
                    "title": "注册试用",
                    "shallow": "选择需要试用的订阅，如果没有选项请先前往订阅管理添加。"
                },
                "trialTime": {
                    "title": "试用时间(小时)",
                    "shallow": "新用户注册时订阅试用时间。"
                },
                "currency": {
                    "title": "货币单位",
                    "shallow": "仅用于展示使用，更改后系统中所有的货币单位都将发生变更。",
                    "placeholder": "CNY"
                },
                "currencySymbol": {
                    "title": "货币符号",
                    "shallow": "仅用于展示使用，更改后系统中所有的货币单位都将发生变更。",
                    "placeholder": "¥"
                }
            },
            "security": {
                "common": { "title": "安全设置" },
                "emailVerify": {
                    "title": "邮箱验证",
                    "description": "开启后将会强制要求用户进行邮箱验证。"
                },
                "gmailAlias": {
                    "title": "禁止使用Gmail多别名",
                    "description": "开启后Gmail多别名将无法注册。"
                },
                "safeMode": {
                    "title": "安全模式",
                    "description": "开启后除了站点URL以外的绑定本站点的域名访问都将会被403。"
                },
                "adminPath": {
                    "title": "后台路径",
                    "description": "后台管理路径，修改后将会改变原有的admin路径。",
                    "placeholder": "https://x.com/logo.jpeg"
                },
                "emailWhitelist": {
                    "title": "邮箱后缀白名单",
                    "description": "开启后在名单中的邮箱后缀才允许进行注册。"
                },
                "recaptcha": {
                    "title": "防机器人",
                    "description": "开启后将会启用hCaptcha防止机器人。"
                },
                "hCaptchaSiteKey": {
                    "title": "hCaptcha SiteKey",
                    "description": "该SiteKey用于请求hCaptcha服务器来标识网站编号",
                    "placeholder": "a3ca066c-0ea0-42fe-bcd2-23f4ab48d528"
                },
                "ipRegisterLimit": {
                    "title": "IP注册限制",
                    "description": "开启后如果IP注册账户达到规则要求将会被限制注册，请注意IP判断可能因为CDN或前置代理导致问题。"
                },
                "registerTimes": {
                    "title": "次数",
                    "description": "达到注册次数后开启惩罚。",
                    "placeholder": "请输入"
                },
                "lockTime": {
                    "title": "惩罚时间(分钟)",
                    "description": "需要等待惩罚时间过后才可以再次注册。",
                    "placeholder": "请输入"
                }
            },
            "frontend": {
                "common": { "title": "个性化" },
                "themePropColor": {
                    "default": "默认",
                    "darkBlueDay": "深蓝色",
                    "milkGreenDay": "奶绿色",
                    "bambooGreen": "若竹",
                    "mistyPine": "雾松",
                    "glacierBlue": "冰川蓝",
                    "grayTheme": "灰色"
                },
                "sidebarStyle": {
                    "title": "边栏风格",
                    "shallow": "设置侧边栏的颜色。"
                },
                "headerStyle": { "title": "头部风格", "shallow": "设置顶部的颜色。" },
                "themeColor": {
                    "title": "主题色",
                    "shallow": "设置整个网页的主题色。"
                },
                "background": {
                    "title": "背景",
                    "shallow": "将会在后台登录页面进行展示。",
                    "placeholder": "https://x.com/logo.jpeg"
                }
            },
            "inviteAndRebate": {
                "common": { "title": "支付和返利" },
                "inviteRebateEnable": {
                    "title": "启用用户返利",
                    "description": "如果启用则在被邀请的用户充值时，按照下面设置的充值比例给予用户返利。"
                },
                "inviteRebateRate": {
                    "title": "返利比例",
                    "description": "设置返利的金额比例。",
                    "placeholder": "请输入返利比例"
                },
                "discountInfo": {
                    "title": "优惠信息",
                    "description": "设置优惠信息，它将会展示在充值页面顶部。",
                    "placeholder": "设置优惠信息"
                },
                "inviteInfo": {
                    "title": "邀请信息",
                    "description": "设置邀请信息，它将会展示在用户邀请页面，用于显示返利比例。",
                    "placeholder": "设置返利信息"
                }
            },
            "welcome": {
                "common": { "title": "首页信息" },
                "homeDescription": {
                    "title": "首页描述",
                    "description": "设置首页的简要描述内容。",
                    "placeholder": "请输入首页描述内容"
                },
                "whyChooseUs": {
                    "title": "为什么选择我们",
                    "description": "设置关于为什么选择我们的描述。",
                    "placeholder": "请输入详细描述"
                },
                "bilibiliLink": {
                    "title": "Bilibili 官方链接",
                    "description": "设置 Bilibili 官方账号的链接地址。",
                    "placeholder": "https://space.bilibili.com/xxxx"
                },
                "youtubeLink": {
                    "title": "YouTube 官方链接",
                    "description": "设置 YouTube 官方账号的链接地址。",
                    "placeholder": "https://youtube.com/channel/xxxx"
                },
                "instagramLink": {
                    "title": "Instagram 官方链接",
                    "description": "设置 Instagram 官方账号的链接地址。",
                    "placeholder": "https://instagram.com/xxxx"
                },
                "wechatLink": {
                    "title": "微信公众账号链接",
                    "description": "设置微信公众账号的链接地址。",
                    "placeholder": "请输入微信公众链接"
                },
                "filingNumber": {
                    "title": "备案号",
                    "description": "设置站点的备案号。",
                    "placeholder": "如：粤ICP备12345678号"
                },
                "pageSuffix": {
                    "title": "站点后缀",
                    "description": "设置站点名称后缀，用于标题显示。",
                    "placeholder": "如：- 你的站点名称"
                }
            },
            "sendMail": {
                "common": {
                    "title": "邮件设置",
                    "warning": "如果您更改了本页面的配置，需要对队列服务及逆行重启。 另外本页配置优先级高于.env中的邮件配置。",
                    "sendTestMailTitle": "发送测试邮件",
                    "sendTestMailShallow": "邮件将会发送到当前登录管理员的邮箱",
                    "sendTestMailTo": "发送测试邮件到",
                    "sending": "发送邮件中",
                    "success": "成功",
                    "receiverAddr": "收信地址",
                    "senderHost": "发信服务器",
                    "senderPort": "发信端口",
                    "senderEncrypt": "发信加密方式",
                    "senderUsername": "发信用户名",
                    "sendErr": "发送邮件失败"
                },
                "smtpServerAddress": {
                    "title": "SMTP 服务器地址",
                    "shallow": "由邮件服务商提供的服务地址",
                    "placeholder": "请输入服务器地址"
                },
                "smtpServerPort": {
                    "title": "SMTP 服务端口",
                    "shallow": "常见的端口有 25, 465, 587",
                    "placeholder": "请输入端口号"
                },
                "smtpEncryption": {
                    "title": "SMTP 加密方式",
                    "shallow": "465 端口加密方式一般为 SSL，587 端口加密方式一般为 TLS",
                    "placeholder": "请输入加密方式"
                },
                "smtpAccount": {
                    "title": "SMTP 账号",
                    "shallow": "由邮件服务商提供的账号",
                    "placeholder": "请输入账号"
                },
                "smtpPassword": {
                    "title": "SMTP 密码",
                    "shallow": "由邮件服务商提供的密码",
                    "placeholder": "请输入密码"
                },
                "senderAddress": {
                    "title": "发件地址",
                    "shallow": "由邮件服务商提供的发件地址",
                    "placeholder": "请输入发件地址"
                },
                "emailTemplate": {
                    "title": "邮件模版",
                    "shallow": "你可以在文档查看如何自定义邮件模板",
                    "placeholder": "请选择邮件模板"
                }
            },
            "notice": {
                "common": { "title": "通知设置" },
                "displayName": {
                    "title": "显示名称",
                    "shallow": "仅用于前端页面显示",
                    "placeholder": "通用通知接口1"
                },
                "barkEndpoint": {
                    "title": "Bark 接入点",
                    "shallow": "Bark 服务器后端 API 地址",
                    "placeholder": "https://<ip>:<port>/<secret-key>"
                },
                "barkGroup": {
                    "title": "Bark 群组",
                    "shallow": "客户端显示的群组名称",
                    "placeholder": "web"
                }
            },
            "appDownload": {
                "common": {
                    "title": "APP",
                    "hint": "用于自有客户端(APP)的版本管理及更新"
                },
                "enabled": {
                    "title": "开放下载",
                    "shallow": "如果启用则允许用户访问下载页面"
                },
                "windows": {
                    "title": "Windows",
                    "shallow": "Windows 端版本号及下载地址",
                    "placeholder": "https://xxxx.com/xxx.exe"
                },
                "macos": {
                    "title": "macOS",
                    "shallow": "macOS 端版本号及下载地址",
                    "placeholder": "https://xxxx.com/xxx.dmg"
                },
                "linux": {
                    "title": "Linux",
                    "shallow": "Linux 端版本号及下载地址",
                    "placeholder": "https://xxxx.com/xxx.deb"
                },
                "android": {
                    "title": "Android",
                    "shallow": "Android 端版本号及下载地址",
                    "placeholder": "https://xxxx.com/xxx.apk"
                },
                "ios": {
                    "title": "IOS",
                    "shallow": "IOS 端版本号及下载地址",
                    "placeholder": "https://xxxx.com/xxx.ipk"
                }
            }
        },
        "payConfig": {
            "title": "支付设置",
            "description": "在这里可以管理所有支持的付款方式，目前仅支持支付宝支付，但是您也可以先行配置其他支付方式，如没有符合您的支付流程，可以等待项目进一步完善后在更新日志中查看是否支持。",
            "attention": {
                "title": "注意事项",
                "point1": "务必先配置支付方式的信息再进行启用，这真的很重要。",
                "point2": "修改付款方式配置时，如果显示为\"---\"则代表该选项已经被设置且非空，如果需要重新修改直接输入新值保存即可。"
            },
            "common": {
                "detail": "{method} 配置",
                "fillAttention": "为确保安全，不显示详细信息，重新填写以创建或覆盖已有配置。",
                "discountAmount": "优惠金额（可以在系统设置的 \"支付和返利\" 中设置用户提示信息）",
                "saveConfigBtnHint": "保存",
                "cancelBtnHint": "取消",
                "saveSuccess": "配置保存成功",
                "alterSuccess": "配置修改成功",
                "discountPlaceholder": "优惠金额（充值金额大于优惠金额时应用优惠）",
                "saveOrAlterFailure": "未知错误"
            },
            "alipay": {
                "title": "支付宝",
                "config": {
                    "appId": "应用Id",
                    "appPublicKeyCertContent": "应用公钥证书内容",
                    "appPrivateKey": "应用私钥",
                    "alipayRootCert": "支付宝根证书",
                    "alipayPublicKeyCertContent": "支付宝公钥证书内容"
                }
            },
            "wechat": {
                "title": "微信支付",
                "config": {
                    "mchId": "商户Id",
                    "mchCertSerial": "商户证书序列号",
                    "apiV3Key": "API v3 Key",
                    "privateKey": "私钥"
                }
            },
            "apple": {
                "title": "Apple Pay",
                "config": {
                    "issuerId": "Issuer Id",
                    "bundleId": "Bundle Id",
                    "privateKeyId": "Private Key Id",
                    "privateKeyContent": "私钥内容"
                }
            },
            "addPaymentMethod": "添加支付方式",
            "enableBtnHint": "启用",
            "disableBtnHint": "禁用",
            "setupPaymentMethod": "配置"
        },
        "themeConfig": {
            "title": "主题配置",
            "using": "使用，",
            "setAsCurrent": ""
        },
        "groupMgr": {
            "title": "权限组管理",
            "description": "权限组是用来标示不同的订阅等级，您可以将同等级别但是额度等有细微区别的订阅计画归纳在一个权限组中方便管理。",
            "common": {
                "addNewGroup": "新建权限组",
                "alterGroupName": "修改权限组名称",
                "addConfirmed": "确认添加",
                "alterConfirmed": "确认修改",
                "cancel": "取消",
                "addSuccess": "添加成功",
                "addFailure": "添加失败",
                "alterSuccess": "修改权限组成功",
                "alterFailure": "修改权限组失败",
                "delSuccess": "删除成功",
                "delFailure": "删除失败",
                "delMention": "将立即删除条目，您不能撤销此操作。 已经关联订阅计画的权限组谨慎删除。",
                "delNotAllowed": "该权限组有关联资源，不可以删除。"
            },
            "groupId": "权限组ID",
            "groupName": "权限组名称",
            "groupPlaceholder": "输入权限组名称",
            "userCount": "用户数量",
            "planCount": "订阅数量",
            "operate": "操作",
            "editBtnHint": "编辑",
            "deleteBtnHint": "删除"
        },
        "docMgr": {
            "title": "知识库管理",
            "description": "在这里您可以编写给用户的说明文档，它可以是软件的设计说明书、使用教程或注意事项等，其文档的内容支持Markdown格式。",
            "addDoc": "添加新文档",
            "addSuccess": "添加新文档成功",
            "addFailure": "添加文档失败",
            "titleNotEmpty": "文档的标题不能为空",
            "table": {
                "docId": "#",
                "isShow": "是否显示",
                "sortAs": "排序",
                "lang": "语言",
                "category": "分类",
                "title": "标题",
                "createdAt": "创建时间",
                "updatedAt": "更新时间",
                "operate": "操作",
                "edit": "编辑",
                "delete": "删除"
            },
            "form": {
                "add": "添加文档",
                "edit": "编辑文档",
                "cancel": "取消",
                "confirm": "确认",
                "addBtn": "添加",
                "editBtn": "修改",
                "title": { "title": "文档标题", "placeholder": "输入文档的标题" },
                "sort": {
                    "title": "排序",
                    "placeholder": "输入文档的排序级别 越高的数值代表优先级越高"
                },
                "category": {
                    "title": "分类",
                    "placeholder": "文档将按照该字段进行分类展示"
                },
                "lang": { "title": "文档语言", "placeholder": "选择文档语言" }
            }
        },
        "planMgr": {
            "title": "订阅管理",
            "description": "在这里可以添加新的订阅计画、修改已有订阅计画的描述、价格、余粮、其所属的权限组等。",
            "addNewPlan": "添加新订阅",
            "alterPlan": "修改订阅信息",
            "table": {
                "id": "#",
                "isSale": "启用销售",
                "isRenew": "允许续费",
                "sort": "排序等级",
                "group": "所属权限组",
                "name": "名称",
                "nums": "数量",
                "residue": "余量",
                "operate": "操作",
                "operateBtn": { "update": "修改", "delete": "删除" }
            },
            "mention": {
                "common": {
                    "success": "成功",
                    "failure": "失败",
                    "delMention": "如果订阅计画已经处于销售中，谨慎删除。"
                }
            },
            "form": {
                "title": "添加订阅",
                "items": {
                    "name": { "title": "订阅名称", "placeholder": "输入订阅计画的名称" },
                    "describe": {
                        "title": "订阅描述",
                        "placeholder": "输入订阅的描述（支持Markdown）"
                    },
                    "isSale": { "title": "启用销售" },
                    "isRenew": { "title": "启用续费" },
                    "group": { "title": "所属权限组", "placeholder": "选择所属权限组" },
                    "capacityLimit": {
                        "title": "最大允许用户数目",
                        "placeholder": "最大允许用户数目"
                    },
                    "planResidue": {
                        "title": "剩余数量",
                        "placeholder": "目前剩余的数量"
                    },
                    "sort": { "title": "排序", "placeholder": "用于前端排序" },
                    "periodPlaceholder": {
                        "month": "输入月付价格",
                        "quarter": "输入季付价格",
                        "halfYear": "输入半年付价格",
                        "year": "输入年付价格"
                    }
                }
            }
        },
        "couponMgr": {
            "title": "优惠券管理",
            "description": "在这里您可以为一些特定的节日等添加一些优惠券，它允许用户在下单的时候使用并按照您设置的比例进行抵折优惠。",
            "fetchErr": "获取数据失败请重试",
            "fetchPlanKvFailurese": "获取订阅计画列表失败",
            "table": {
                "id": "#",
                "enabled": "是否启用",
                "name": "名称",
                "code": "券码",
                "percentOff": "优惠信息",
                "capacity": "总数量",
                "residue": "剩余数量",
                "startTime": "启用时间",
                "endTime": "结束时间",
                "actions": "操作",
                "edit": "编辑",
                "delete": "删除"
            },
            "modal": {
                "newCoupon": "添加新优惠券",
                "editCoupon": "修改优惠券信息",
                "confirmAdd": "确认添加",
                "confirmEdit": "确认修改",
                "emptyNotAllow": "该项目是必填的",
                "delMention": "条目删除后优惠券将会立即失效，您不能撤销此操作。",
                "cancel": "取消",
                "name": { "title": "优惠券名称", "placeholder": "请输入优惠券名称" },
                "code": {
                    "title": "券码",
                    "placeholder": "自定义优惠券码（留空随即生成）"
                },
                "percentOff": {
                    "title": "优惠信息（百分比）",
                    "placeholder": "输入优惠百分比（如-12%OFF）"
                },
                "activeRange": { "title": "优惠券可使用的有效期范围" },
                "capacity": {
                    "title": "优惠券最大使用次数",
                    "placeholder": "限制最大使用限制（为空则不限制）"
                },
                "residue": {
                    "title": "优惠券剩余使用次数",
                    "placeholder": "设置优惠券剩余使用的次数"
                },
                "perUserLimit": {
                    "title": "每个用户可使用优惠券的次数",
                    "placeholder": "限制每个用户可使用的次数（为空则不限制）"
                },
                "planLimit": {
                    "title": "指定订阅计画",
                    "placeholder": "限制指定订阅计画可以使用优惠（为空则不限制）"
                }
            }
        },
        "orderMgr": {
            "title": "订单管理",
            "description": "在这里您可以检视所有订阅计画的订单，筛选不同用户的订单、手动对用户的订单处理通过等。",
            "table": {
                "id": "#",
                "orderId": "订单号",
                "email": "用户邮箱",
                "status": { "title": "类型", "t1": "新购", "t2": "续费", "t3": "编辑" },
                "planName": "计画名称",
                "period": "周期",
                "group": "权限组",
                "amount": "实付金额",
                "price": "原始价格",
                "isSuccess": {
                    "title": "订单状态",
                    "cancelOrder": "取消订单",
                    "passOrder": "通过订单"
                },
                "createdAt": "订单创建时间",
                "action": { "title": "操作", "showDetail": "显示细节" }
            },
            "search": "查询订单",
            "resetSearch": "重置查询",
            "failureReason": "失败原因",
            "couponId": "优惠券ID",
            "couponName": "优惠券名称",
            "noEntry": "无",
            "orderDetail": "订单详情",
            "searchModal": {
                "email": {
                    "title": "用户邮箱",
                    "placeholder": "输入用户邮箱（模糊搜索）"
                },
                "sort": {
                    "title": "排序算法",
                    "placeholder": "选择排序算法",
                    "ASC": "升序",
                    "DESC": "降序"
                }
            },
            "tradeWaiting": "未支付",
            "tradeFailure": "交易失败",
            "tradeSuccess": "成功"
        },
        "userMgr": {
            "userManager": "用户管理",
            "description": "你可以在这里管理所有的用户，包括员工和管理员，授予或取消管理权限、设定用户余额、重置密码、手动添加新用户等操作。",
            "enterEmail": "请输入邮箱",
            "enterValidEmail": "请输入正确的邮箱格式",
            "enterPassword": "请输入密码",
            "passwordMinLength": "密码长度不能少于 6 位",
            "passwordMaxLength": "密码长度不能超过 20 位",
            "passwordStrength": "密码必须包含字母、数字和特殊字符",
            "validationSuccess": "验证通过",
            "validationFailed": "表单验证失败，请检查输入项",
            "editProfile": "编辑资料",
            "banUser": "封禁用户",
            "unbanUser": "解封用户",
            "noRecord": "无记录",
            "normal": "正常",
            "banned": "封禁",
            "deleted": "注销",
            "nullContent": "NULL",
            "balance": "余额",
            "orderCount": "订单数量",
            "planCount": "计划数",
            "actions": "操作",
            "updateSuccess": "更新成功",
            "addUserSuccess": "添加新用户成功",
            "unknownError": "未知错误",
            "email": "邮箱",
            "registerDate": "注册日期",
            "isAdmin": "管理员",
            "isStaff": "员工",
            "accountStatus": "帐户状态",
            "inviteCode": "邀请码",
            "query": "查询",
            "reset": "重置",
            "addNewUser": "新增用户",
            "searchUser": "查询用户",
            "cancel": "取消",
            "submit": "提交",
            "userEmail": "用户邮箱",
            "inputUserEmail": "输入用户邮箱（模糊搜索）",
            "inputEmail": "输入邮箱",
            "password": "密码",
            "inputPassword": "输入密码",
            "editUser": "编辑用户",
            "no": "否",
            "yes": "是",
            "mention": {
                "title": "您确定要封禁用户吗？",
                "content": "如果封禁该用户，那么该用户将无法登录到{appName}，并且与其相关的所有资源将变为不可用。"
            }
        },
        "activation": {
            "activateLog": "激活纪录",
            "description": "您可以检视所有售出密钥的具体激活情况，查看客户端的识别码、激活时间等。",
            "click2getKey": "点击以获取密钥内容",
            "createdAt": "创建时间",
            "turn2keyPage": "转到密钥",
            "showKey": "显示密钥",
            "email": "邮箱",
            "key": "密钥",
            "search": "搜寻",
            "filter": "筛选",
            "filterAll": "全部",
            "filterActive": "仅有效",
            "sortAlgorithm": "排序算法",
            "sortASC": "升序",
            "sortDESC": "降序"
        },
        "keysMgr": {
            "keyMgr": "密钥管理",
            "description": "您可以查询到所有生成的密钥内容和他们的使用情况、有效期等等，您也可以执行禁用某一个密钥。",
            "enableKey": "启用密钥",
            "disableKey": "禁用密钥",
            "table": {
                "email": "邮箱地址",
                "key": "密钥内容",
                "isValid": "是否有效",
                "isUsed": "是否使用",
                "isExpired": "是否到期",
                "active": "活跃",
                "inactive": "已过期",
                "yes": "有效",
                "no": "过期",
                "ok": "正常",
                "blocked": "禁用状态",
                "unUsed": "未使用",
                "used": "已使用",
                "showDetail": "显示细节",
                "blockKey": "禁用密钥",
                "unblockKey": "解除禁用"
            },
            "searchModal": {
                "searchMethod": "查询方式",
                "byId": "通过ID进行查询",
                "byContent": "通过密钥进行查询（模糊）",
                "keyId": "密钥ID",
                "idPlaceholder": "在这里输入密钥的ID",
                "keyContent": "密钥码",
                "contentPlaceholder": "在这里输入密钥的部分内容"
            },
            "mention": {
                "blockOk": "禁用密钥成功 ID:{id}",
                "title": "您确定要禁用密钥吗",
                "content": "为保证用户体验，请再一次确认密钥内容，密钥一旦禁用后，在客户端将不能再继续使用。"
            },
            "detailModal": {
                "title": "密钥细节",
                "userId": "用户ID",
                "planName": "订阅计画名称",
                "expiredAt": "到期日期",
                "keyGeneratedAt": "密钥生成日期",
                "clientId": "客户端ID"
            }
        },
        "noticeMgr": {
            "title": "公告管理",
            "description": "在这里可以对公告进行管理，启用的公告将展示在用户首页的轮播图中，可以设置的公告有优惠活动、节日通知、注意事项等。",
            "addNotice": "添加公告",
            "table": {
                "id": "#",
                "show": "是否显示",
                "title": "标题",
                "createdAt": "创建时间",
                "action": { "title": "操作", "edit": "编辑", "delete": "删除" }
            },
            "mention": {
                "title": "您确认要删除吗？",
                "content": "将立即删除该公告，您不能撤销此操作。"
            },
            "modal": {
                "addNew": "新建公告",
                "title": {
                    "title": "公告标题",
                    "placeholder": "作为大标题显示在轮播图中"
                },
                "content": { "title": "公告内容", "placeholder": "编写公告的主要内容" },
                "tag": { "title": "公告标签", "placeholder": "输入公告的标签" },
                "img": { "title": "背景图片URL", "placeholder": "不填写则使用默认背景" }
            }
        },
        "adminTicket": {
            "ticketMgr": "工单管理",
            "description": "您可以在此处检视到所有用户提交的工单，工单有三个紧急程度，推荐您从紧急工单开始处理。",
            "pendingTicket": "待处理工单",
            "finishedTicket": "已完成工单",
            "type": {
                "pendingDescription": "活跃的工单，这是您应该先进行处理的工单，如果工单确认已经完成请关闭，如不关闭则该工单将始终置于此处。",
                "finishedDescription": "已经完成的工单，您可以在这里检视它们。"
            },
            "chooseOneNecessary": "您应该至少选择一个项目",
            "mention": {
                "title": "您确定要关闭该工单吗？",
                "content": "该工单在关闭后将会归档至已完成工单中，您可以再次检视他们的内容，但是再也无法回复此工单。"
            }
        }
    },
    "userLogin": {
        "loginToContinue": "登入以继续",
        "email": "邮件地址",
        "password": "密码",
        "haveNoAccount": "还没有您的帐户？",
        "login": "登入",
        "reg": "立即注册",
        "otherMethods": "或使用其他方式继续",
        "github": "以Github帐户继续",
        "microsoft": "以Microsoft帐户继续",
        "apple": "以Apple帐户继续",
        "google": "以Google帐户继续",
        "backHomePage": "回到首页",
        "form": {
            "emailRequire": "邮箱地址是必填的",
            "passwordRequire": "你还没有输入密码"
        },
        "authPass": "验证通过",
        "loginFailure": "登录失败",
        "checkForm": "请检查表单",
        "if2FaEnabledHint": "如果您启用了两步验证（非必填）",
        "reqErr": "请求出现错误请稍后再试",
        "accountLocked": "您的帐户可能被注销或封禁，暂时无法继续使用，如你仍然认为这是一个错误，请联系我们的技术支持。",
        "tokenNotExist": "请提供Token"
    },
    "userRegister": {
        "backHomePage": "回到首页",
        "newAccount": "准备您的新帐户",
        "email": "邮箱地址",
        "verifyCode": "邮箱验证码",
        "invalidEmailFormat": "邮箱格式不合法",
        "sendVerifyCode": "发送邮箱验证码",
        "sendSuccess": "邮件发送成功，请查看邮箱。",
        "pwd": "密码",
        "pwdAgain": "确认密码",
        "inviteCode": "邀请码（选填）",
        "agreement": "我已阅读并同意",
        "terminalUserAgreement": "用户条款",
        "reg": "注册",
        "infoIncomplete": "信息不完整",
        "pwdIncorrect": "密码不一致",
        "regSuccess": "注册成功返回登录",
        "regFailure": "注册失败",
        "success": "成功",
        "failure": "失败",
        "unknownErr": "未知错误",
        "verifyCodeErr": "验证码错误",
        "verifyCodeExpireErr": "验证码错误或已过期，请重试或获取新验证码。",
        "thisMailAlreadyExist": "该邮箱已经被注册",
        "pageConfigFetchFailure": "配置获取失败请刷新重试",
        "stopRegisterTitle": "已停止注册",
        "stopRegisterHint": "很抱歉，目前注册功能已暂停。 如阁下有需要，请稍后再试或联系我们的支持团队获取更多信息。 感谢您的理解与支持。",
        "passwordComplexRequirePart1": "* 密码需要符合",
        "passwordComplexRequirePart2": "复杂度要求",
        "passwordComplexHint1": "1. 需要使用大写字母、小写字母、数字、特殊符号中的三种及以上",
        "passwordComplexHint2": "2. 长度再10位以上",
        "form": {
            "checkForm": "请检查表单内容",
            "emailFormatErr": "邮箱地址格式不正确",
            "gmailLimitErr": "不允许使用Gmail多别名",
            "gmailDotNotAllowed": "Gmail地址中不允许有\".\"",
            "gmailPartLowerForced": "Gmail地址的本地部分必须全部小写",
            "googlemailNotAllowed": "不支持使用googlemail地址",
            "verifyCodeRequire": "请输入验证码",
            "verifyCodeFormatErr": "验证码格式不正确",
            "passwordRequire": "请输入密码",
            "passwordLengthRequire": "密码长度必须大于或等于10位",
            "passwordComplexRequire": "密码必须包含大写字母、小写字母、数字、特殊符号中的至少三种",
            "passwordAgainNotMatch": "两次输入的密码不匹配",
            "passwordAgainRequire": "请输入确认密码",
            "inviteCodeRequire": "邀请码是必填的"
        },
        "hCaptcha": {
            "passed": "人机验证通过",
            "expired": "已过期请重试",
            "challengeExpired": "验证超时",
            "err": "其他错误"
        },
        "allRightsReserved": "版权所有",
        "securityAndLaws": "该网站受hCaptcha保护和验证，请遵守当地法律。"
    },
    "userSummary": {
        "title": "仪表盘",
        "myPlan": "我的订阅",
        "shortcut": "捷径",
        "timeLeft": "订阅有效，将在 {msg} 过期。",
        "toPurchase": "购买订阅",
        "tutorial": { "title": "查看教程", "content": "学习如何使用 {name}" },
        "checkKey": {
            "title": "查看密钥",
            "content": "快速将密钥导入对应客户端进行使用"
        },
        "renewPlan": { "title": "续费订阅", "content": "对您当前的订阅进行续费" },
        "appDown": {
            "title": "应用下载",
            "content": "下载我们不同平台的应用程序以获得更好的体验"
        },
        "support": {
            "title": "遇到问题",
            "content": "遇到问题可以通过工单与我们的人机沟通"
        },
        "haveTicket": "您有 {count} 条待处理的工单",
        "toCheckTicket": "去查看",
        "showAllKeys": "查看所有密钥"
    },
    "userDocument": {
        "title": "使用文档",
        "description": "您可以在这里查阅不同的文档，包括但不限于网站的使用方法、注意事项、操作流程等，如果您发现文章中有错误，请提交工单。",
        "searchPlaceholder": "请输入要搜索的内容（模糊搜索）",
        "searchBtn": "搜索",
        "noContentTitle": "无结果",
        "noContentTitleHint": "没有符合您搜索结果或语言的文档，尝试换一个关键词吧。"
    },
    "newPurchase": {
        "title": "购买订阅",
        "description": "您可以在这里选择最适合您的订阅计画，如果您的余额不足请先充值，订单将为您保留5分钟。",
        "headerPlaceholder": "选择最适合您的计画",
        "purchaseBtn": "订购",
        "noLeft": "剩余数量不足",
        "monthPay": "月付",
        "moreMethod": "更多选择"
    },
    "newSettlement": {
        "err": "错误",
        "monthPay": "月付",
        "quarterPay": "季付",
        "halfYearPay": "半年付",
        "yearPay": "年付",
        "payCycle": "付款周期",
        "couponPlaceholder": "有优惠券 ？",
        "verifyBtn": "验证",
        "orderTotalTitle": "订单总额",
        "total": "总计",
        "submitOrder": "提交订单",
        "coupon": "优惠券",
        "notify": {
            "passTitle": "验证通过",
            "couponVerified": "优惠券有效",
            "couponInvalid": "优惠券无效",
            "couponIsNull": "输入的优惠券码不能为空"
        }
    },
    "userProfile": {
        "title": "个人中心",
        "myWallet": "我的钱包",
        "walletSub": "帐户余额（仅用于消费）",
        "alertPwd": "修改密钥",
        "oldPwd": "旧密钥",
        "oldPwdSub": "请输入旧密码",
        "newPwd": "新密钥",
        "newPwdSub": "请输入新的密钥",
        "newPwdAgain": "确认密钥",
        "newPwdAgainSub": "请再输入一遍新的密钥",
        "saveBtn": "保存",
        "notify": "通知",
        "enableNotify": "启用到期通知",
        "deleteAccount": "注销帐户",
        "deleteAccountSub": "您的帐户将被标记为删除，如果需要重新使用我们的服务，请重新注册",
        "deleteBtn": "注销我的帐户",
        "oldPwdVerifiedFailure": "旧密码验证失败",
        "alertFailure": "密钥修改失败",
        "alertSuccess": "修改成功",
        "alertSuccessSub": "请使用新密码登录",
        "errorPwdFormat": "密码格式错误",
        "pwdNotMatch": "两次密码输入不一致",
        "oldPwdNotNull": "旧密码不能为空",
        "toTopUp": "去充值",
        "deleteMyTitle": "警告",
        "deleteMyContent": "确定删除您的账户吗？ 如需要使用服务请重新注册。",
        "deleteMyPositiveText": "确认删除",
        "deleteMyNegativeText": "取消",
        "deletedSuccessMsg": "账户已删除，后会有期。",
        "deleteErrOccur": "遇到错误",
        "faAuth": "两步验证2FA",
        "faAuthHint": "两步验证是一种安全机制，增加了登入帐户的保护层。 用户在输入密码后，还需完成第二步身份验证，如输入发送到手机的验证码、使用身份验证应用程序生成的动态码，或通过指纹等生物特征确认。",
        "faAuthStatus": "两步验证状态",
        "faEnabled": "已启用",
        "faNotEnabled": "未启用",
        "setup2Fa": "设置两步验证",
        "disable2Fa": "取消两步验证",
        "unknownErr": "未知错误",
        "disable2FaCancelled": "已取消",
        "closed2FaHint": "已经关闭两步验证，如有需要请重新添加。",
        "setup2FaModal": {
            "followStep": "根据提示在您的验证器上加入",
            "step1Title": "按照以下步骤以启2FA验证",
            "step1Context1": "1. 您的移动设备上需要有一个通用的验证器",
            "step1Context2": "2. 点击验证器上的Scan按钮来扫描此处的QR码",
            "step1Context3": "3. 该QR Code中包含有您的验证信息和唯一密钥，请妥善保存",
            "step2Context1": "为了确保您的验证器能够正常使用，我们需要进行测试。",
            "test2Fa": "测试",
            "cancel": "取消"
        },
        "deleteMyAccountModal": {
            "title": "注销帐户",
            "contentLine1": "帐号注销是一个不可逆的操作。 一旦您确认删除，您将永久失去该帐号的访问权限，这意味着您将无法再次登入，且与此帐号相关的所有数据，包括但不限于您的个人资讯、历史记录、收藏内容、购买记录等，都将无法再访问。",
            "contentLine2": "如果您在我们的平台上有正在进行的业务，例如未完成的订单、正在参与的活动、订阅服务等，这些都将随着帐号删除而终止或取消，可能会给您带来相应的损失。 同时，您与其他用户之间通过本平台建立的联系、互动资讯等也都将不复存在。",
            "contentLine3": "请再次确认您的决定，如果您还有任何疑虑或问题，欢迎联系我们的客服，我们将竭诚为您解答。 若您仍然希望删除帐号，请点击「确认删除」按钮。",
            "inputHint1": "输入 \"",
            "inputHint2": "\" 以继续。",
            "confirmDelete": "确认删除"
        },
        "failure": "遇到错误",
        "notLatestHint": "个人信息更新失败，当前的数据可能不是最新的。",
        "resetPassword": {
            "previousPasswordRequire": "请输入旧密码",
            "previousPasswordVerifiedFailure": "旧密码验证失败",
            "passwordRequire": "请输入密码",
            "passwordConflict": "新密码不能与先前的密码相同",
            "passwordLengthRequire": "密码长度必须大于或等于10位",
            "passwordComplexRequire": "密码必须包含大写字母、小写字母、数字、特殊符号中的至少三种",
            "passwordAgainNotMatch": "两次输入的密码不匹配",
            "passwordAgainRequire": "请输入确认密码",
            "updatePasswordFailure": "更新密码出错"
        },
        "secondary": {
            "title": "基础资料",
            "card": {
                "avatar": "用户头像",
                "name": "用户名",
                "lastLoginAt": "上一次登录时间",
                "accountStatus": "帐户状态"
            },
            "modify": {
                "uploadIconHint": "上传",
                "alterAvatar": "新增/修改头像",
                "alterShallow": "* 点击头像或用户名以添加或修改（设置后均不允许清除）",
                "alterName": "修改用户名",
                "input": {
                    "label": "用户名",
                    "placeholder": "输入用户名",
                    "spaceIsNotAllowed": "用户名验证不通过",
                    "require": {
                        "p1": "不允许纯数字且数字开头",
                        "p2": "不允许有空格",
                        "p3": "长度大于三位"
                    }
                }
            },
            "mention": {
                "alterNameSuccess": "修改用户名成功",
                "alterNameErr": "修改用户名失败请稍后再试",
                "newNameIsNotValid": "用户名不合法",
                "click2SetName": "点击以设置用户名",
                "fetchAvatarErr": "获取头像数据失败请稍后再试",
                "alterAvatarErr": "修改头像失败请稍后再试",
                "success": "成功",
                "alterAvatarSuccess": "修改头像成功，",
                "uploadImageHint": "您可以上传图像作为您的新头像",
                "imageRequire": {
                    "title": "注意",
                    "p1": "您可以上传 *.jpg(jpeg), *.png, *.webp 等主流格式。",
                    "p2": "图像的长宽比为1:1（正方形），如果为其他比例，将会被居中裁切为正方形，多余的部分将会被删除。",
                    "p3": "您上传的图像的大小将会被设置为160px。"
                },
                "click2Upload": "点击或拖动文件到该区域来上传",
                "uploadWarn": "*请不要上传敏感数据，比如您的银行卡、信用卡、密码和安全码等。"
            }
        }
    },
    "userKeys": {
        "myKeys": "我的密钥",
        "description": "您可以检视您的所有密钥的活动情况、到期日期等，如需要为密钥设置备注，请前往激活纪录页面。",
        "noKeys": "您还没有有效的购买纪录",
        "keyDetail": "密钥细节",
        "keyId": "密钥ID",
        "orderId": "链接订单ID",
        "clientId": "激活客户端ID",
        "active": "在有效期内",
        "inActive": "已过期",
        "valid": "密钥状态正常",
        "invalid": "密钥已被禁用",
        "isUsed": "已激活使用",
        "noUsed": "还未使用",
        "releaseData": "密钥生成日期",
        "expirationData": "到期日期",
        "none": "无",
        "authorizationFor": "密钥授权给",
        "hoverClickMention": "点击可以复制到剪贴板",
        "copiedSuccessMessage": "密钥复制成功，请参照文档说明继续操作。",
        "copyFailure": "复制失败",
        "hoverCopiedSuccessMention": "复制成功"
    },
    "userOrders": {
        "myOrders": "我的订单",
        "description": "您所有的订单将在此处展示，如果您有未支付的订单将会展示在顶部，您可以点击继续支付或取消订单，已经完成的订单您可以在此处检视订单细节。",
        "orderId": "#",
        "planName": "订阅名称",
        "planCycle": "订阅周期",
        "orderPrice": "订单金额",
        "orderStatus": "订单状态",
        "createdAt": "创建时间",
        "operate": "操作",
        "showDetail": "订单细节",
        "cancelOrder": "取消订单",
        "canceled": "已取消",
        "period": {
            "monthPay": "月付",
            "quarterPay": "季付",
            "halfYearPay": "半年付",
            "yearPay": "年付"
        },
        "orderStatusTags": {
            "success": "成功",
            "cancelled": "失败",
            "notPay": "未支付"
        },
        "orderCancelled": "订单已取消",
        "unknownErr": "未知错误"
    },
    "userTopUp": {
        "topUp": "帐户充值",
        "description": "您可以在这里进行帐户充值，支持自定义充值金额，您也可以关注下方是否有优惠信息展示，使用提及的支付方式以获得优惠。",
        "chooseTopUpAmount": "选择充值金额",
        "quickSelect": "快速选择",
        "customAmount": "自定义金额",
        "maxAmount": "最大金额: 10,000,000",
        "amountInputPlaceHolder": "输入要充值的金额",
        "otherAmount": "其他金额",
        "payMethod": "支付方式",
        "wechat": "微信支付",
        "alipay": "支付宝",
        "apple": "Apple Pay",
        "yourAmount": "您的金额",
        "discount": "优惠",
        "accountBalance": "帐户余额",
        "balanceResult": "余额合计",
        "commitTopUp": "提交",
        "payMethodNotAllow": "支付方式不可用请选择其他",
        "topUpIssueOccur": "充值遇到问题？",
        "payIssueOccur": "支付遇到问题？",
        "chatWithUs": "联系客服",
        "pay": "支付",
        "qrCodeScannedSuccess": "QR码扫描成功",
        "orClickToApp": "或点击条转到App继续",
        "topUpSuccess": "充值成功",
        "thankU": "感谢您的支持"
    },
    "userConfirmOrder": {
        "switchPlan": "切换订阅",
        "cancelOrder": "取消订单",
        "yourPrice": "您的价格",
        "couponOffset": "优惠券抵折",
        "price": "价格",
        "submit": "提交",
        "monthPay": "月付",
        "quarterPay": "季付",
        "halfYearPay": "半年付",
        "yearPay": "年付",
        "goodInfo": "商品信息",
        "cycleOrType": "周期/类型",
        "orderNumber": "订单号",
        "createdAt": "创建日期",
        "orderExpired": "订单已超时",
        "paySuccessfully": "支付成功，感谢您的支持。",
        "balanceNotEnough": "您的余额不足请先充值，该订单保留五分钟。",
        "orderErrorOccur": "订单遇到错误",
        "orderCancelled": "订单已取消"
    },
    "paymentResultParts": {
        "goodInfoView": { "goodInfo": "商品信息" },
        "orderInfoView": { "orderInfo": "订单信息" }
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
            "orderInfo": "订单信息",
            "cycleOrType": "周期/类型",
            "orderNumber": "订单号",
            "createdAt": "创建日期",
            "amount": "支付金额",
            "paidAt": "支付时间"
        }
    },
    "orderDetail": {
        "finished": "已完成",
        "finishedAndSuccessDescription": "订单已成功支付并开通",
        "useManual": "查看使用教程",
        "payPending": "尚未支付",
        "pendingDescription": "订单暂时保留，可以点击下面的按钮以继续支付。",
        "toPay": "去支付",
        "outDate": "已失效",
        "outDateDescription": "由于您取消了订单或未在指定时间内完成支付，因此该订单已被取消，您可以重新选取订阅。",
        "chooseNewPlan": "选择新的订阅计画"
    },
    "userInvite": {
        "myInvite": "我的邀请",
        "description": "如果管理员开启了邀请返利，您可以在这里看到具体的返利规则，生成邀请连结后分享给其他人注册后即可获取到返利。",
        "unit": "人数",
        "inviteCodeMgr": "您的邀请码",
        "generateInviteCode": "生成随机邀请码",
        "faCodeManage": "邀请码管理",
        "email": "邮箱",
        "createdAt": "注册时间",
        "createFaCodeFailure": "创建失败",
        "linkCopiedSuccess": "链接复制成功",
        "generateFaCode": "生成邀请码",
        "flushFaCode": "刷新邀请码",
        "faCode": "邀请码",
        "noFaCode": "你还没有邀请码，请先生成。",
        "faLink": "邀请连结",
        "generateFaCodePlease": "请先生成邀请码",
        "usersMyInvited": "我邀请的用户",
        "generateCodeConfirm": "确认生成/刷新",
        "generateCodeHint": "请注意，邀请码创建后不可关闭。",
        "positiveClick": "确认",
        "negativeClick": "取消"
    },
    "userTickets": {
        "description": "如果您遇到使用上的问题，请在此处提交工单，我们的技术支持和客服看到后将会进行回覆并在下方表格处标出具体的回覆时间，您可以点击查看工单和我们交流。",
        "ticketId": "#",
        "ticketSubject": "工单主题",
        "ticketUrgency": "工单级别",
        "ticketContent": "工单内容",
        "ticketUrgencyLevel": { "low": "低", "med": "中", "high": "高" },
        "ticketStatus": "工单状态",
        "ticketCreatedAt": "创建时间",
        "lastResponse": "最后回覆",
        "operate": "操作",
        "checkTicket": "查看工单",
        "closeTicket": "关闭工单",
        "userTickets": "历史工单",
        "addTicket": "新建工单",
        "ticketActive": "未关闭",
        "ticketInActive": "已关闭",
        "form": {
            "ticketSubjectDescription": "请输入工单主题",
            "ticketUrgencyDescription": "请选择工单紧急程度",
            "ticketBody": "请输入你想要解决的问题，尽量全面。",
            "ticketNotFinished": "请补全工单的信息再试"
        },
        "checkForm": "请检查表单是否完整",
        "cancel": "取消",
        "submit": "提交",
        "commitNewTicketSuccess": "提交新的工单成功",
        "commitNewTicketFailure": "提交新的工单错误",
        "noReply": "还没有回覆",
        "noTickets": "您还没有提交过工单",
        "ticketCloseSuccess": "工单关闭成功",
        "ticketCloseFailure": "工单关闭失败",
        "chatDialog": {
            "input": {
                "finished": "该工单已经结束",
                "inputHere": "输入要发送的消息"
            },
            "send": "发送",
            "missingToken": "不可以创建Websocket连接，因为缺少Token。",
            "msgEmptyNotAllowed": "消息的内容不可以为空",
            "accessNotPermit": "非法访问",
            "sendSuccess": "发送消息成功"
        }
    },
    "userActivation": {
        "activateLog": "激活纪录",
        "description": "在此处您可以查看您的激活纪录并设置解除绑定设备，也可以对每一个密钥和激活纪录设置备注信息。",
        "id": "#",
        "orderId": "订单编号",
        "orderStatus": "订单",
        "createdAt": "创建时间",
        "operate": "操作",
        "userId": "用户Id",
        "email": "邮箱",
        "keyId": "密钥Id",
        "isBind": "是否激活",
        "active": "有效",
        "inactive": "无效",
        "requestAt": "请求时间",
        "clientVersion": "客户端",
        "osType": "操作系统",
        "remark": "备注",
        "noRemark": "无备注",
        "showDetail": "查看详情",
        "actions": "操作",
        "details": "细节",
        "keyContent": "密钥内容",
        "keyGeneratedAt": "密钥生成时间",
        "activateRequestAt": "激活请求时间",
        "useIssueOccur": "使用遇到问题？",
        "chatWithUs": "联系我们",
        "cancelBind": "取消绑定",
        "alterRemark": "修改备注",
        "commitRemark": "提交",
        "updateSuccess": "更新成功",
        "setRemark": "在这里设置备注信息"
    },
    "userAppDownload": {
        "title": "APP下载",
        "description": "",
        "common": {
            "title": "下载我们的应用程序",
            "shallow2": "为不同的客户端获取我们的应用程序",
            "shallow": "使用我们的应用程序您可以更便捷地访问我们的服务，免去每次在浏览器繁琐的操作；您可在在文档中找到详尽的安装需求和使用教程，包括运行环境、报错解决等，如果您遇到了其他问题请联系我们的技术支持。",
            "noDownload": "很抱歉暂时还没有提供下载，请您稍后再试，如有问题请提交工单以联系我们的支持服务。"
        },
        "suffix": {
            "p1": "*macOS平台的应用程序请使用macOS14及以上并配合Apple晶片（M1或更高）。",
            "p2": "该App所展示的信息和使用请遵守当地的使用法规，是否允许使用该App也应该取决于您公司的IT规范。"
        },
        "card": {
            "common": { "welcome": "欢迎来到" },
            "mobile": {
                "designFor": "为移动端设计",
                "designShallow": "您可以在这里获取我们的移动端应用程序",
                "iosDownloadShallow": "下载 IOS 客户端",
                "androidDownloadShallow": "下载 Android 客户端"
            },
            "desktop": {
                "designFor": "为桌面端设计",
                "designShallow": "您可以在这里获取我们的桌面动端应用程序",
                "windowsDownloadShallow": "下载 Windows 客户端",
                "osxDownloadShallow": "下载 macOS 客户端",
                "linuxDownloadShallow": "下载 Linux 客户端"
            }
        },
        "downloadDisabledHint": "很抱歉，App下载暂时不可用或被管理员禁用，如您有需要请在工单处联系我们的技术已获取支持。",
        "windows": {
            "title": "Windows NT",
            "shallow": "该客户端适用于NT内核的Windows操作系统，请查看文档页面以获取兼容性支持。"
        },
        "osx": {
            "title": "macOS (OSX)",
            "shallow": "该客户端适用于Darwin内核的macOS(OSX)操作系统，请查看文档页面以获取兼容性支持。"
        },
        "linux": {
            "title": "Linux",
            "shallow": "该客户端适用于Linux内核的各类发行版，由于发行版系列不同请查看文档页面以获取兼容性支持。"
        },
        "android": {
            "title": "Android",
            "shallow": "该客户端适用于搭载了Google Android操作系统的移动设备，请查看文档页面以获取兼容性支持。"
        },
        "ios": {
            "title": "IOS",
            "shallow": "该客户端适用于搭载了Apple IOS操作系统的移动设备，请查看文档页面以获取兼容性支持。"
        }
    },
    "welcome": {
        "A": {
            "aboutUs": "关于我们",
            "pricing": "定价",
            "login": "登录",
            "register": "注册帐号",
            "welcomeTo": "欢迎来到",
            "welcomeToSub": "穿过县境上长长的隧道，便是雪国。 夜空下，大地一片莹白，火车在信号所前停下来。 \"在这里川端康成用几近吝啬的简洁文字，拉开了《雪国》的序幕。",
            "startingUse": "开始使用",
            "whyUs": "为什么选择我们",
            "whyUsSub": "\"穿过县境上长长的隧道，便是雪国。 夜空下，大地一片莹白，火车在信号所前停下来。 \"在这里川端康成用几近吝啬的简洁文字，拉开了《雪国》的序幕。",
            "browseSafe": "浏览安全",
            "browseSafeSub": "优秀的防火墙过滤系统能有效防御网路钓鱼和恶意网站",
            "encrypt": "端到端加密",
            "encryptSub": "双向 SSL 和端对端加密保护您的隐私安全，即使是服务器也无法读取您的信息",
            "mgr": "高效管理",
            "mgrSub": "一个用户介面管理所有密钥，管理功能完善丰富，无须担心订阅泄露问题",
            "fast": "方便快捷",
            "fastSub": "提供完整的 API 文档供 WebApp 或是嵌入到第三方软件中",
            "fastLink": "快速连结",
            "subscribeUs": "关注我们",
            "filingsCode": "备案号 {code}"
        }
    },
    "pagination": {
        "perPage10": "10 条数据/页",
        "perPage20": "20 条数据/页",
        "perPage50": "50 条数据/页",
        "perPage100": "100 条数据/页"
    },
    "modal": { "cancel": "取消", "confirm": "确认" },
    "week": {
        "monday": "周一",
        "tuesday": "周二",
        "wednesday": "周三",
        "thursday": "周四",
        "friday": "周五",
        "saturday": "周六",
        "sunday": "周日"
    },
    "period": {
        "month": "月付",
        "quarter": "季付",
        "halfYear": "半年付",
        "year": "年付"
    },
    "operate": {
        "cancel": "取消",
        "confirm": "确认",
        "update": "更新",
        "add": "添加"
    },
    "notFound": {
        "title": "404 页面不存在",
        "description": "我们无法找到您请求的页面，它可能已经被删除或连结有误。如果您认为这是一个错误，请提交工单以联系我们。",
        "p1": "将在 {sec}s 后返回主页，如果您的浏览器没有响应，请点击以下按钮。",
        "manualBack": "返回主页"
    },
    "forbidden": {
        "title": "403 无权限",
        "description": "您可能没有足够的权限来访问本页面。如果您认为这是一个错误，请提交工单以联系我们。"
    }
}
