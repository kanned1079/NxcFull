export default {
    "commonHeader": {
        "menuTxt": "menu",
        "userData": "User information",
        "editUserData": "Edit user information",
        "logout": "Log out"
    },
    "commonAside": {
        "admin": {
            "dashboard": "Tableboard",
            "queueMonit": "Service side monitoring",
            "settings": "Settings",
            "systemConfig": "System Settings",
            "paymentConfig": "Payment Settings",
            "themeConfig": "Theme settings",
            "server": "Server",
            "privilege": "Authority management",
            "finance": "Finance",
            "subscription": "Subscription Management",
            "coupon": "Coupon Management",
            "order": "Order Management",
            "activate": "Activate the record",
            "key": "Key management",
            "user": "User",
            "userMgr": "User Management",
            "notice": "Announcement Management",
            "ticket": "Work order management",
            "doc": "Knowledge database management"
        },
        "user": {
            "dashboard": "Tableboard",
            "document": "Use the document",
            "app": "APP Download",
            "subscription": "Subscribe",
            "purchase": "Buy Subscription",
            "surplus": "My key",
            "fiance": "Finance",
            "topUp": "top up",
            "myOrder": "My Order",
            "myInvite": "My invitation",
            "user": "User",
            "profile": "Personal Center",
            "support": "My work order",
            "activateLog": "Activate the record"
        }
    },
    "adminViews": {
        "common": {
            "fetchDataSuccess": "Acquisition of data successfully",
            "fetchDataFailure": "If the data is failed, please try again later",
            "addSuccess": "Added successfully",
            "addFailure": "If the addition fails, please try later",
            "updateSuccess": "Modification was successful",
            "updateFailure": "If the modification fails, please try again later",
            "deleteSuccess": "Deletion successfully",
            "deleteFailure": "Please try again if the delete failed",
            "confirm": "Confirm",
            "cancel": "Cancel",
            "success": "Operation is successful",
            "failure": "Operation failed",
            "NotAllowed": "Illegal format",
            "checkForm": "Please check the form",
            "unknownErr": "Unknown error",
            "dialog": { "delete": "Do you know you want to delete it?" },
            "yes": "yes",
            "no": "no"
        },
        "login": {
            "secureCard": {
                "title": "Safety inspection",
                "securePath": "Safe path",
                "hint": "To ensure system security, you need to enter the security path before you can enter the administrator login page. Please enter the security path in the input box below. After successfully verifying the security path, you can choose to choose Save for subsequent quick login.",
                "placeholder": "Please enter a safe path",
                "checkBtn": "examine",
                "rememberPath": "Remember the safe path"
            },
            "card": {
                "back": "Return to homepage",
                "toAdminCenter": "Log in to the Admin Center",
                "emailAddr": "Administrator email",
                "emailInputArea": {
                    "title": "Administrator email",
                    "placeholder": "Email address"
                },
                "passwordInputArea": {
                    "title": "Administrator password",
                    "placeholder": "password"
                },
                "login": "login",
                "forgetPassword": "forget the password",
                "formNotPassed": "The form verification fails"
            },
            "mention": {
                "title": "hint",
                "description": "This page is an administrator page, which can only be accessed by administrators. If you do not have administrator permissions or come here accidentally, please click the link below to return to the home page."
            },
            "message": {
                "passwordErr": "Incorrect password",
                "adminNotExist": "The administrator does not exist",
                "noPrivilege": "Access without permission",
                "authPassed": "Verification passed",
                "authFailure": "Verification failed",
                "otherErr": "Other errors",
                "pathCheckPassed": "The security path check passed",
                "pathCheckFailure": "Incorrect security path",
                "rememberSecureMention": "To ensure the security of backend management, please do not check it if this is not your private computer."
            }
        },
        "summary": {
            "cockpit": "Tableboard",
            "systemConfig": "System Settings",
            "paymentConfig": "Payment Settings",
            "planMgr": "Subscription Management",
            "userMgr": "User Management",
            "orderMgr": "Order Management",
            "keyMgr": "Key management",
            "incomeText": "Yesterday's income / Monthly income",
            "pendingTicket": "You have {nums} orders to be processed",
            "toHandle": "Go and check it out",
            "apiAccessCard": "Number of API interface visits within a week",
            "apiAccessCardHint": "This data is only for you to understand current API access and does not represent your server performance.",
            "incomeWeek": "Income amount within one week",
            "incomeCardHint": "Here is a chart of income amounts over a week, which will cause inaccurate display if the cache is cleared.",
            "core": "core",
            "reqErr": "Encountered an error",
            "reqErrHint": "An error was encountered while obtaining the overview information, which caused the request to be unable to complete, so the chart cannot be displayed. Please try again later.",
            "userCard": {
                "title": "User Overview",
                "allRegisteredUsers": "General registration user",
                "activeUsers": "Live users",
                "inactiveUsers": "Inactive users",
                "blockedUsers": "Ban or reimbursement"
            },
            "general": {
                "title": "generally",
                "localTime": "Browser time",
                "osType": "Operating system type",
                "appName": "Application Name",
                "appUrl": "Application URL",
                "currency": "Coin unit",
                "allowRegister": "Allow registration"
            },
            "system": {
                "title": "System configuration",
                "axiosAddr": "HTTP backend address",
                "wsAddr": "Websocket backend address",
                "serverTime": "Server time",
                "uptime": "Operational time",
                "gatewayStatus": "API network status",
                "dbStatus": "Database status",
                "redisStatus": "Redis Status",
                "serverOsType": "Server operating system type",
                "serverOsArch": "Server operating system architecture",
                "runMode": "Running mode",
                "cpuNums": "Number of CPU cores for network servers",
                "numCgoCall": "Number of garbage collections",
                "time": "Second-rate",
                "paymentMethods": "Enable payment method",
                "runOK": "Run normally",
                "runErr": "Abnormal behavior",
                "checkServer": "Please check the environment configuration of your backend server",
                "stopRegisterHint": "You seem to have disabled new user registration",
                "toSetting": "Turn to Settings"
            }
        },
        "queueMonit": {
            "title": "Service side monitoring",
            "headerHint": "Please do not stay on this page for a long time. Frequent inquiries during peak networks will affect network performance and database throughput.",
            "latency": {
                "title": "Server delay",
                "retry": "Test again",
                "hint": "*The request here refers to the time required for successful response to the client after the client sends a request to the server.",
                "unit": "millisecond",
                "level": {
                    "l1": {
                        "title": "Excellent",
                        "description": "This is a very good network situation and you will hardly feel delayed."
                    },
                    "l2": {
                        "title": "normal",
                        "description": "This is the network situation most of the time, and users almost feel that they don’t feel delayed."
                    },
                    "l3": {
                        "title": "Poor",
                        "description": "In this case the user may feel a slight lag or delay."
                    },
                    "l4": {
                        "title": "Difference",
                        "description": "You can feel the obvious delay and affect the user experience."
                    },
                    "l5": {
                        "title": "Very bad",
                        "description": "There is obvious delay and the loading speed is slower or even unable to refresh, which seriously affects user interaction and experience."
                    },
                    "offline": {
                        "title": "Server off",
                        "description": "Unable to connect to the server or handle request errors, please check if the configuration is correct."
                    }
                }
            },
            "api": {
                "title": "API request status in the last 7 days",
                "items": {
                    "ok": { "title": "Success (StatusOK)", "unit": "Second-rate" },
                    "notFound": {
                        "title": "Status Path Error (404)",
                        "unit": "Second-rate"
                    },
                    "unAuthorized": {
                        "title": "Unauthorized access (401)",
                        "unit": "Second-rate"
                    },
                    "login2reg": { "title": "Login / Register", "unit": "Second-rate" }
                }
            },
            "log": {
                "title": "Journal Record",
                "deleteLogMsg": "Delete the journal {nums}",
                "deleteLogErr": "Delete journal failed",
                "logTableRows": "Total diary records",
                "logTableSize": "Blog tables occupy space",
                "unit": { "lines": "OK", "mb": "Mega-character" },
                "deleteLog": "Delete the journal",
                "exportCsv": "Export CSV",
                "deleteLogHint": "This will execute the removal of all journals a week ago. The removal method is hard delete. All deleted journals will not be restored. It is recommended to save the backup first.",
                "warn": {
                    "title": "Do you know you want to delete your journal?",
                    "description": "The list will be deleted immediately and you cannot revoke this operation."
                },
                "export": {
                    "title": "Export the csv file",
                    "description": "This will export the following table as the csv file and download it locally. If you do not have the permission to open the download, the download may fail. Click the confirm button to execute the export."
                },
                "table": {
                    "id": "#",
                    "method": "Request method",
                    "path": "Please request a route",
                    "code": "Status code",
                    "clientIp": "Client IP",
                    "responseTime": "Processing time",
                    "requestAt": "Request time"
                }
            }
        },
        "systemConfig": {
            "title": "System Settings",
            "common": {
                "success": "Updated configuration successfully",
                "err": "Update configuration failed"
            },
            "site": {
                "common": { "title": "Site" },
                "appName": {
                    "title": "Site name",
                    "shallow": "Used to display where the site name is required.",
                    "placeholder": "Site name"
                },
                "appSubName": {
                    "title": "Site subtitle",
                    "shallow": "Generally displayed below the main title.",
                    "placeholder": "subtitle"
                },
                "appDescription": {
                    "title": "Site Description",
                    "shallow": "Used to display where the site is required.",
                    "placeholder": "Site Description"
                },
                "appUrl": {
                    "title": "Site URL",
                    "shallow": "The latest website address will be displayed at the URL if you need it.",
                    "placeholder": "Site URL"
                },
                "forceHttps": {
                    "title": "Force HTTPS",
                    "shallow": "When the site does not use HTTPS, CDN or anti-generation to enable forced HTTPS."
                },
                "logoUrl": {
                    "title": "LOGO",
                    "shallow": "Used to display where the LOGO is required.",
                    "placeholder": "The URL address of the logo"
                },
                "subscribeUrl": {
                    "title": "Subscribe URL",
                    "shallow": "Used for subscription, leave blank as the site URL. If you need multiple subscription URLs to be randomly fetched, please use commas to split them.",
                    "placeholder": "Subscribe url"
                },
                "tosUrl": {
                    "title": "Terms of User (TOS) URL",
                    "shallow": "Used to jump to the Terms of User (TOS)",
                    "placeholder": "User Terms and Conditions Address"
                },
                "stopRegister": {
                    "title": "Stop new user registration",
                    "shallow": "No one will be able to register after it is turned on."
                },
                "inviteRequire": {
                    "title": "Forced invitation",
                    "shallow": "After opening, you must fill in the invitation code when registering with a new user."
                },
                "trialSubscribe": {
                    "title": "Register for trial",
                    "shallow": "Select the subscription you want to try. If there is no option, please go to Subscription Management to add it first."
                },
                "trialTime": {
                    "title": "Trial time (hours)",
                    "shallow": "Subscribe to trial time when new users sign up."
                },
                "currency": {
                    "title": "Currency Unit",
                    "shallow": "It is only used for display use, and after the change, all currency units in the system will change.",
                    "placeholder": "CNY"
                },
                "currencySymbol": {
                    "title": "Currency symbol",
                    "shallow": "It is only used for display use, and after the change, all currency units in the system will change.",
                    "placeholder": "¥"
                }
            },
            "security": {
                "common": { "title": "Security settings" },
                "emailVerify": {
                    "title": "Email Verification",
                    "description": "After opening, the user will be required to perform email verification."
                },
                "gmailAlias": {
                    "title": "Prohibit Gmail multiple alias",
                    "description": "After opening, multiple Gmail alias will not be registered."
                },
                "safeMode": {
                    "title": "Safe Mode",
                    "description": "After opening, the domain name access to this site except the site URL will be 403."
                },
                "adminPath": {
                    "title": "Background path",
                    "description": "The background management path will change the original admin path after modification.",
                    "placeholder": "https://x.com/logo.jpeg"
                },
                "emailWhitelist": {
                    "title": "Email suffix whitelist",
                    "description": "Registration is only allowed after opening the mailbox suffix in the list."
                },
                "recaptcha": {
                    "title": "Anti-robot",
                    "description": "When enabled, hCaptcha will be enabled to prevent the robot."
                },
                "hCaptchaSiteKey": {
                    "title": "hCaptcha SiteKey",
                    "description": "The SiteKey is used to request the hCaptcha server to identify the website number",
                    "placeholder": "a3ca066c-0ea0-42fe-bcd2-23f4ab48d528"
                },
                "ipRegisterLimit": {
                    "title": "IP registration restrictions",
                    "description": "After opening, if the IP registration account meets the rules requirements, the registration will be restricted. Please note that IP judgment may cause problems due to CDN or pre-agent."
                },
                "registerTimes": {
                    "title": "frequency",
                    "description": "Punishment will be activated after the number of registrations is reached.",
                    "placeholder": "Please enter"
                },
                "lockTime": {
                    "title": "Punishment time (minutes)",
                    "description": "You need to wait until the penalty time has passed before you can register again.",
                    "placeholder": "Please enter"
                }
            },
            "frontend": {
                "common": { "title": "Personalization" },
                "themePropColor": {
                    "default": "default",
                    "darkBlueDay": "Dark blue",
                    "milkGreenDay": "Milk green",
                    "bambooGreen": "Ruozhu",
                    "mistyPine": "Fog Pine",
                    "glacierBlue": "Glacier Blue",
                    "grayTheme": "grey"
                },
                "sidebarStyle": {
                    "title": "Sidebar style",
                    "shallow": "Set the color of the sidebar."
                },
                "headerStyle": {
                    "title": "Head style",
                    "shallow": "Set the color at the top."
                },
                "themeColor": {
                    "title": "Theme color",
                    "shallow": "Set the theme color of the entire web page."
                },
                "background": {
                    "title": "background",
                    "shallow": "It will be displayed on the background login page.",
                    "placeholder": "https://x.com/logo.jpeg"
                }
            },
            "inviteAndRebate": {
                "common": { "title": "Payment and rebate" },
                "inviteRebateEnable": {
                    "title": "Enable user rebates",
                    "description": "If enabled, when the invited user recharges, the user will be rebated according to the recharge ratio set below."
                },
                "inviteRebateRate": {
                    "title": "Rebate ratio",
                    "description": "Set the rebate amount ratio.",
                    "placeholder": "Please enter the rebate ratio"
                },
                "discountInfo": {
                    "title": "Discount information",
                    "description": "Set the offer information and it will be displayed at the top of the top-up page.",
                    "placeholder": "Set up discount information"
                },
                "inviteInfo": {
                    "title": "Invitation information",
                    "description": "Set the invitation information, it will be displayed on the user invitation page to display the rebate ratio.",
                    "placeholder": "Set rebate information"
                }
            },
            "welcome": {
                "common": { "title": "Home page information" },
                "homeDescription": {
                    "title": "Home page description",
                    "description": "Set a brief description of the homepage.",
                    "placeholder": "Please enter the homepage description content"
                },
                "whyChooseUs": {
                    "title": "Why choose us",
                    "description": "Set up our description about why we selected.",
                    "placeholder": "Please enter a detailed description"
                },
                "bilibiliLink": {
                    "title": "Bilibili official link",
                    "description": "Set the link address of Bilibili official account.",
                    "placeholder": "https://space.bilibili.com/xxxx"
                },
                "youtubeLink": {
                    "title": "YouTube official link",
                    "description": "Set the link address of the official YouTube account.",
                    "placeholder": "https://youtube.com/channel/xxxxx"
                },
                "instagramLink": {
                    "title": "Instagram official link",
                    "description": "Set the link address of your official Instagram account.",
                    "placeholder": "https://instagram.com/xxxx"
                },
                "wechatLink": {
                    "title": "WeChat public account link",
                    "description": "Set the link address of the WeChat public account.",
                    "placeholder": "Please enter the WeChat public link"
                },
                "filingNumber": {
                    "title": "Registration number",
                    "description": "Set up the registration number of the site.",
                    "placeholder": "For example: Guangdong ICP No. 12345678"
                },
                "pageSuffix": {
                    "title": "Site suffix",
                    "description": "Set the site name suffix for title display.",
                    "placeholder": "For example:- Your site name"
                }
            },
            "sendMail": {
                "common": {
                    "title": "Mail settings",
                    "warning": "If you change the configuration of this page, you need to restart the queue service and reverse direction. In addition, the configuration priority of this page is higher than the mail configuration in .env.",
                    "sendTestMailTitle": "Send test email",
                    "sendTestMailShallow": "The email will be sent to the current logged-in administrator's email address",
                    "sendTestMailTo": "Send a test email to",
                    "sending": "Send email",
                    "success": "success",
                    "receiverAddr": "Received address",
                    "senderHost": "Delivery server",
                    "senderPort": "Send port",
                    "senderEncrypt": "Encryption method for sending messages",
                    "senderUsername": "Credit name",
                    "sendErr": "Send email failure"
                },
                "smtpServerAddress": {
                    "title": "SMTP server address",
                    "shallow": "Service address provided by email service provider",
                    "placeholder": "Please enter the server address"
                },
                "smtpServerPort": {
                    "title": "SMTP service port",
                    "shallow": "Common ports are 25, 465, 587",
                    "placeholder": "Please enter the port number"
                },
                "smtpEncryption": {
                    "title": "SMTP encryption method",
                    "shallow": "The 465 port encryption method is generally SSL, and the 587 port encryption method is generally TLS.",
                    "placeholder": "Please enter the encryption method"
                },
                "smtpAccount": {
                    "title": "SMTP account",
                    "shallow": "Account provided by email service provider",
                    "placeholder": "Please enter your account number"
                },
                "smtpPassword": {
                    "title": "SMTP Password",
                    "shallow": "Password provided by email service provider",
                    "placeholder": "Please enter your password"
                },
                "senderAddress": {
                    "title": "Send address",
                    "shallow": "The sending address provided by the mail service provider",
                    "placeholder": "Please enter the sending address"
                },
                "emailTemplate": {
                    "title": "Email Template",
                    "shallow": "You can view how to customize the email template in the document",
                    "placeholder": "Please select the email template"
                }
            },
            "notice": {
                "common": { "title": "Notification Settings" },
                "displayName": {
                    "title": "display name",
                    "shallow": "For front-end page display only",
                    "placeholder": "General Notification Interface 1"
                },
                "barkEndpoint": {
                    "title": "Bark access point",
                    "shallow": "Bark Server Backend API Address",
                    "placeholder": "https://<ip>:<port>/<secret-key>"
                },
                "barkGroup": {
                    "title": "Bark Groups",
                    "shallow": "The group name displayed by the client",
                    "placeholder": "web"
                }
            },
            "appDownload": {
                "common": {
                    "title": "APP",
                    "hint": "For version management and update of own clients (APPs)"
                },
                "enabled": {
                    "title": "Open and drop the load",
                    "shallow": "If enabled, allow users to access the download page"
                },
                "windows": {
                    "title": "Windows",
                    "shallow": "Windows version number and download address",
                    "placeholder": "https://xxxx.com/xxx.exe"
                },
                "macos": {
                    "title": "macOS",
                    "shallow": "macOS version number and download address",
                    "placeholder": "https://xxxx.com/xxx.dmg"
                },
                "linux": {
                    "title": "Linux",
                    "shallow": "Linux version number and download address",
                    "placeholder": "https://xxxx.com/xxx.deb"
                },
                "android": {
                    "title": "Android",
                    "shallow": "Android version number and download address",
                    "placeholder": "https://xxxx.com/xxx.apk"
                },
                "ios": {
                    "title": "IOS",
                    "shallow": "IOS version number and download address",
                    "placeholder": "https://xxxx.com/xxx.ipk"
                }
            }
        },
        "payConfig": {
            "title": "Payment Settings",
            "description": "All supported payment methods can be managed here. Currently, only payment bag payment is supported, but you can also configure other payment methods first. If it does not meet your payment process, you can wait for the project to be further improved and check whether it is supported in the update log.",
            "attention": {
                "title": "Things to note",
                "point1": "It is really important to configure the payment method information before enabling it.",
                "point2": "When modifying the payment method configuration, if it is displayed as \"---\", it means that the option has been set and is not empty. If you need to modify it again, just enter the new value to save."
            },
            "common": {
                "detail": "{method} configuration",
                "fillAttention": "To ensure safety, no detailed information is displayed, refilling to create or overwrite existing configurations.",
                "discountAmount": "Discount amount (user prompt information can be set in the \"Payment and Rebate\" set by the system)",
                "saveConfigBtnHint": "save",
                "cancelBtnHint": "Cancel",
                "saveSuccess": "Configuration saved successfully",
                "alterSuccess": "Configuration modification was successful",
                "discountPlaceholder": "Discount amount (if the recharge amount is greater than the discount amount)",
                "saveOrAlterFailure": "Unknown error"
            },
            "alipay": {
                "title": "Payment",
                "config": {
                    "appId": "Application Id",
                    "appPublicKeyCertContent": "Application Key Certificate Content",
                    "appPrivateKey": "Application Privacy Key",
                    "alipayRootCert": "Payment Bag Certificate Book",
                    "alipayPublicKeyCertContent": "Payment Bag Key Certificate Content"
                }
            },
            "wechat": {
                "title": "WeChat Payment",
                "config": {
                    "mchId": "Merchant Id",
                    "mchCertSerial": "Merchant Certificate Serial Number",
                    "apiV3Key": "API v3 Key",
                    "privateKey": "Privacy key"
                }
            },
            "apple": {
                "title": "Apple Pay",
                "config": {
                    "issuerId": "Issuer Id",
                    "bundleId": "Bundle Id",
                    "privateKeyId": "Private Key Id",
                    "privateKeyContent": "Privacy content"
                }
            },
            "addPaymentMethod": "Add payment method",
            "enableBtnHint": "Start",
            "disableBtnHint": "Disabled",
            "setupPaymentMethod": "Configuration"
        },
        "themeConfig": {
            "title": "Topic Configuration",
            "using": "use,",
            "setAsCurrent": ""
        },
        "groupMgr": {
            "title": "Authority management",
            "description": "The right group is used to mark different subscription levels, and you can subscribe to the same level but different sizes in a right group for easy management.",
            "common": {
                "addNewGroup": "New permission group",
                "alterGroupName": "Modify the name of the permission group",
                "addConfirmed": "Confirm to add",
                "alterConfirmed": "Confirm the modification",
                "cancel": "Cancel",
                "addSuccess": "Added successfully",
                "addFailure": "Add failure",
                "alterSuccess": "The modification permission group was successful",
                "alterFailure": "The modification permission group failed",
                "delSuccess": "Deletion successfully",
                "delFailure": "Deletion failed",
                "delMention": "The list will be deleted immediately and you cannot revoke this operation. Relevant groups for subscription plans should be deleted with caution.",
                "delNotAllowed": "This right group has related resources and cannot be deleted."
            },
            "groupId": "Rights Group ID",
            "groupName": "Name of permission group",
            "groupPlaceholder": "Enter permission group name",
            "userCount": "Number of users",
            "planCount": "Subscription quantity",
            "operate": "operate",
            "editBtnHint": "Edit",
            "deleteBtnHint": "delete"
        },
        "docMgr": {
            "title": "Knowledge database management",
            "description": "Here you can write a description file to your user. It can be a software design manual, tutorial or precautions, etc. The content of the document supports the Markdown format.",
            "addDoc": "Add a new document",
            "addSuccess": "Adding a new document successfully",
            "addFailure": "Adding file failed",
            "titleNotEmpty": "The title of the document cannot be empty",
            "table": {
                "docId": "#",
                "isShow": "Whether to display",
                "sortAs": "Sort",
                "lang": "Language",
                "category": "Category",
                "title": "Title",
                "createdAt": "Creation time",
                "updatedAt": "Update time",
                "operate": "operate",
                "edit": "Edit",
                "delete": "delete"
            },
            "form": {
                "add": "Add a document",
                "edit": "Edit the document",
                "cancel": "Cancel",
                "confirm": "Confirm",
                "addBtn": "Add to",
                "editBtn": "Revise",
                "title": {
                    "title": "File title",
                    "placeholder": "Enter the title of the file"
                },
                "sort": {
                    "title": "Sort",
                    "placeholder": "The ordering level of the input file is higher. The higher the value represents the higher the priority."
                },
                "category": {
                    "title": "Category",
                    "placeholder": "The document will be displayed in categories according to this field"
                },
                "lang": {
                    "title": "Document language",
                    "placeholder": "Select a document language"
                }
            }
        },
        "planMgr": {
            "title": "Subscription Management",
            "description": "Here you can add new subscription plans, modify the description, price, balance, permission group to which the subscription plans are already included, etc.",
            "addNewPlan": "Add a new subscription",
            "table": {
                "id": "#",
                "isSale": "Enable sales",
                "isRenew": "Allow continuous fee",
                "sort": "Sort level",
                "group": "The so-called rights group",
                "name": "Name",
                "nums": "quantity",
                "residue": "Amount",
                "operate": "operate",
                "operateBtn": { "update": "Revise", "delete": "delete" }
            },
            "mention": {
                "common": {
                    "success": "success",
                    "failure": "Fail",
                    "delMention": "If the subscription plan is already on sale, please be careful to delete it."
                }
            },
            "form": {
                "title": "Add a subscription",
                "items": {
                    "name": {
                        "title": "Subscription name",
                        "placeholder": "Enter the name of the subscription plan"
                    },
                    "describe": {
                        "title": "Subscription Description",
                        "placeholder": "Enter a description for subscription (Support Markdown)"
                    },
                    "isSale": { "title": "Enable sales" },
                    "isRenew": { "title": "Enable Continuing Fee" },
                    "group": {
                        "title": "The so-called rights group",
                        "placeholder": "Select the rightful group"
                    },
                    "capacityLimit": {
                        "title": "Maximum number of allowed users",
                        "placeholder": "Maximum number of allowed users"
                    },
                    "planResidue": {
                        "title": "The remaining amount",
                        "placeholder": "The remaining amount"
                    },
                    "sort": { "title": "Sort", "placeholder": "For front-end sorting" },
                    "periodPlaceholder": {
                        "month": "Enter monthly payment price",
                        "quarter": "Enter the quarterly paid price",
                        "halfYear": "Enter the half year payment price",
                        "year": "Enter annual payment price"
                    }
                }
            }
        },
        "couponMgr": {
            "title": "Coupon Management",
            "description": "Here you can add some coupons for certain festivals, etc., which allow users to use them when placing orders and offer discounts according to the proportion you set.",
            "fetchErr": "Retry if the data is failed",
            "fetchPlanKvFailurese": "Get subscription plan list failed",
            "table": {
                "id": "#",
                "enabled": "Whether to enable",
                "name": "Name",
                "code": "Coupon",
                "percentOff": "Discount information",
                "capacity": "Total quantity",
                "residue": "The remaining amount",
                "startTime": "Enable time",
                "endTime": "End time",
                "actions": "operate",
                "edit": "Edit",
                "delete": "delete"
            },
            "modal": {
                "newCoupon": "Add a new coupon",
                "editCoupon": "Modify coupon information",
                "confirmAdd": "Confirm to add",
                "confirmEdit": "Confirm the modification",
                "emptyNotAllow": "This item is required",
                "delMention": "The coupon will be invalid immediately after the entry is deleted and you cannot withdraw this operation.",
                "cancel": "Cancel",
                "name": {
                    "title": "Coupon name",
                    "placeholder": "Please enter the coupon name"
                },
                "code": {
                    "title": "Coupon",
                    "placeholder": "Customize coupon code (leave it blank and generate it as soon as possible)"
                },
                "percentOff": {
                    "title": "Offer information (percentage)",
                    "placeholder": "Enter the percentage of discount (such as -12%OFF)"
                },
                "activeRange": { "title": "Exemption range of coupons available" },
                "capacity": {
                    "title": "Maximum number of coupons used",
                    "placeholder": "Limit maximum usage limit (no limit if empty)"
                },
                "residue": {
                    "title": "The number of remaining coupons used",
                    "placeholder": "Set the number of times the coupon is used remaining"
                },
                "perUserLimit": {
                    "title": "The number of times each user can use a coupon",
                    "placeholder": "Limit the number of times each user can use (no limit on space)"
                },
                "planLimit": {
                    "title": "Specify a subscription plan",
                    "placeholder": "Restrictions on the specified subscription plan to use the offer (no limit on space)"
                }
            }
        },
        "orderMgr": {
            "title": "Order Management",
            "description": "Here you can view all subscription-planned orders, filter orders for different users, manually handle orders for users, and more.",
            "table": {
                "id": "#",
                "orderId": "Order number",
                "email": "User mailbox",
                "status": {
                    "title": "Type",
                    "t1": "Newly purchased",
                    "t2": "Continue fee",
                    "t3": "Edit"
                },
                "planName": "Plan name",
                "period": "Cycle",
                "group": "Rights Group",
                "amount": "Real amount of payment",
                "price": "Original price",
                "isSuccess": {
                    "title": "Order status",
                    "cancelOrder": "Cancel an order",
                    "passOrder": "By order"
                },
                "createdAt": "Order creation time",
                "action": { "title": "operate", "showDetail": "Show details" }
            },
            "search": "Inquiry orders",
            "resetSearch": "Reset query",
            "failureReason": "Cause of failure",
            "couponId": "Coupon ID",
            "couponName": "Coupon name",
            "noEntry": "without",
            "orderDetail": "Order details",
            "searchModal": {
                "email": {
                    "title": "User mailbox",
                    "placeholder": "Enter user email (fuzzy search)"
                },
                "sort": {
                    "title": "Sorting algorithm",
                    "placeholder": "Select the sorting algorithm",
                    "ASC": "Ascending order",
                    "DESC": "descending order"
                }
            },
            "tradeWaiting": "Not paid",
            "tradeFailure": "Transaction failed",
            "tradeSuccess": "success"
        },
        "userMgr": {
            "userManager": "User Management",
            "description": "You can manage all users here, including employees and administrators, grant or cancel management rights, set user balances, reset passwords, add new users manually, and more.",
            "enterEmail": "Please enter your email",
            "enterValidEmail": "Please enter the correct email format",
            "enterPassword": "Please enter your password",
            "passwordMinLength": "The password length cannot be less than 6 digits",
            "passwordMaxLength": "The password length cannot exceed 20 digits",
            "passwordStrength": "Passwords must contain letters, numbers and special characters",
            "validationSuccess": "Verification passed",
            "validationFailed": "Form verification failed, please check the input item",
            "editProfile": "Edit information",
            "banUser": "Disable the login",
            "unbanUser": "Unblock the user",
            "noRecord": "Unrecorded",
            "normal": "normal",
            "banned": "Banned",
            "deleted": "Log out",
            "nullContent": "NULL",
            "balance": "Amount",
            "orderCount": "Order quantity",
            "planCount": "Plans",
            "actions": "operate",
            "updateSuccess": "Update successfully",
            "addUserSuccess": "Adding new user successfully",
            "unknownError": "Unknown error",
            "email": "Email",
            "registerDate": "Registration date",
            "isAdmin": "Administrator",
            "isStaff": "staff",
            "accountStatus": "Account status",
            "inviteCode": "Invitation code",
            "query": "Query",
            "reset": "Reset",
            "addNewUser": "Added user",
            "searchUser": "Inquiry user",
            "cancel": "Cancel",
            "submit": "submit",
            "userEmail": "User mailbox",
            "inputUserEmail": "Enter user email (fuzzy search)",
            "inputEmail": "Enter email",
            "password": "password",
            "inputPassword": "Enter password",
            "editUser": "Edit user",
            "no": "no",
            "yes": "yes",
            "mention": {
                "title": "Are you sure you want to disable your account?",
                "content": "If the user is blocked, the user will not be able to log in to {appName} and all resources associated with it will become unavailable."
            }
        },
        "activation": {
            "activateLog": "Activate the record",
            "description": "You can view the specific activation status of all sold keys, view the client's identification code, activation time, etc.",
            "click2getKey": "Click to get key content",
            "createdAt": "Creation time",
            "turn2keyPage": "Turn to key",
            "showKey": "Show key",
            "email": "Email",
            "key": "Key",
            "search": "Search",
            "filter": "Screening",
            "filterAll": "all",
            "filterActive": "Only effective",
            "sortAlgorithm": "Sorting algorithm",
            "sortASC": "Ascending order",
            "sortDESC": "descending order"
        },
        "keysMgr": {
            "keyMgr": "Key management",
            "description": "You can check all generated key contents and their usage status, validity period, etc. You can also disable a key.",
            "enableKey": "Enable key",
            "disableKey": "Disable keys",
            "table": {
                "email": "Email address",
                "key": "Key content",
                "isValid": "Is it valid or not",
                "isUsed": "Whether to use it",
                "isExpired": "Whether it expires",
                "active": "active",
                "inactive": "Expired",
                "yes": "efficient",
                "no": "Expired",
                "ok": "normal",
                "blocked": "Disable status",
                "unUsed": "Not used",
                "used": "Used",
                "showDetail": "Show details",
                "blockKey": "Disable keys",
                "unblockKey": "Unbanned"
            },
            "searchModal": {
                "searchMethod": "Inquiry method",
                "byId": "Query by ID",
                "byContent": "Inquiry through keys (fuzzy)",
                "keyId": "Key ID",
                "idPlaceholder": "Enter the ID of the key here",
                "keyContent": "Keyword",
                "contentPlaceholder": "Enter part of the key here"
            },
            "mention": {
                "blockOk": "Disabling key successfully ID:{id}",
                "title": "Are you sure you want to disable the key?",
                "content": "To ensure the user experience, please confirm the content of the key again. Once the key is disabled, it will no longer be used on the client side."
            },
            "detailModal": {
                "title": "Key details",
                "userId": "User ID",
                "planName": "Subscription plan name",
                "expiredAt": "Date of Expiry",
                "keyGeneratedAt": "Key generation date",
                "clientId": "Client ID"
            }
        },
        "noticeMgr": {
            "title": "Announcement Management",
            "description": "Announcements can be managed here. The activated announcements will be displayed in the wheelbarrow picture on the user's homepage. Announcements that can be set include promotions, day notices, precautions, etc.",
            "addNotice": "Add an announcement",
            "table": {
                "id": "#",
                "show": "Whether to display",
                "title": "Title",
                "createdAt": "Creation time",
                "action": { "title": "operate", "edit": "Edit", "delete": "delete" }
            },
            "mention": {
                "title": "Do you know you want to delete it?",
                "content": "The announcement will be deleted immediately and you cannot withdraw this action."
            },
            "modal": {
                "addNew": "New announcement",
                "title": {
                    "title": "Announcement Title",
                    "placeholder": "Show as a big title in the wheel cast"
                },
                "content": {
                    "title": "Announcement content",
                    "placeholder": "The main contents of writing an announcement"
                },
                "tag": {
                    "title": "Announcement Tags",
                    "placeholder": "Enter the tag for the announcement"
                },
                "img": {
                    "title": "Background image URL",
                    "placeholder": "Use the default background if you don't fill in"
                }
            }
        },
        "adminTicket": {
            "ticketMgr": "Work order management",
            "description": "You can view all users submitted work orders here. There are three urgency levels of work orders, and you are recommended to start processing from an emergency work order.",
            "pendingTicket": "Work orders to be processed",
            "finishedTicket": "Completed work order",
            "type": {
                "pendingDescription": "Live work order, this is the work order you should process first. If the work order is confirmed to be completed, please close it. If it is not closed, the work order will always be placed here.",
                "finishedDescription": "Completed work orders, you can view them here."
            },
            "chooseOneNecessary": "You should choose at least one item",
            "mention": {
                "title": "Are you sure you want to close the ticket?",
                "content": "The order will be archived into the completed order after it is closed. You can view their content again, but you can never reply to the order."
            }
        }
    },
    "userLogin": {
        "loginToContinue": "Log in to continue",
        "email": "Email address",
        "password": "password",
        "haveNoAccount": "Don't have your account yet?",
        "login": "login",
        "reg": "Register now",
        "otherMethods": "Or use other methods to continue",
        "github": "Continue with Github account",
        "microsoft": "Continue with Microsoft Account",
        "apple": "Continue with Apple Account",
        "google": "Continue with Google Account",
        "backHomePage": "Back to homepage",
        "form": {
            "emailRequire": "Email address is required",
            "passwordRequire": "You haven't entered your password yet"
        },
        "authPass": "Verification passed",
        "loginFailure": "Login failed",
        "checkForm": "Please check the form",
        "if2FaEnabledHint": "If you enabled two-step verification (not required)",
        "reqErr": "Please try again later if there is an error",
        "accountLocked": "Your account may be registered or banned and cannot be used at the moment. If you still think this is an error, please contact our technical support.",
        "tokenNotExist": "Please provide a token"
    },
    "userRegister": {
        "backHomePage": "Back to homepage",
        "newAccount": "Prepare your new account",
        "email": "Email address",
        "verifyCode": "Email verification code",
        "invalidEmailFormat": "The email format is illegal",
        "sendVerifyCode": "Send email verification code",
        "sendSuccess": "The email has been sent successfully, please check the email.",
        "pwd": "password",
        "pwdAgain": "Confirm password",
        "inviteCode": "Invitation code (optional)",
        "agreement": "I've read and agreed",
        "terminalUserAgreement": "User Terms",
        "reg": "register",
        "infoIncomplete": "Incomplete information",
        "pwdIncorrect": "Password inconsistent",
        "regSuccess": "Registration is successful and returns to login",
        "regFailure": "Registration failed",
        "success": "success",
        "failure": "Fail",
        "unknownErr": "Unknown error",
        "verifyCodeErr": "Verification code error",
        "verifyCodeExpireErr": "The verification code is incorrect or has expired. Please try again or obtain a new verification code.",
        "thisMailAlreadyExist": "The email has been registered",
        "pageConfigFetchFailure": "If configuration acquisition fails, please refresh and try again",
        "stopRegisterTitle": "Registration has been stopped",
        "stopRegisterHint": "Sorry, the registration function has been suspended. If you need it, please try again later or contact our support team for more information. Thank you for your understanding and support.",
        "passwordComplexRequirePart1": "* Passwords need to be in compliance",
        "passwordComplexRequirePart2": "Complexity requirements",
        "passwordComplexHint1": "1. Three or more types of large letters, small letters, numbers, special symbols are required.",
        "passwordComplexHint2": "2. The length is more than 10 digits",
        "form": {
            "checkForm": "Please check the content of the form",
            "emailFormatErr": "The email address format is incorrect",
            "gmailLimitErr": "Gmail multiple names are not allowed",
            "gmailDotNotAllowed": "\".\" is not allowed in the Gmail address.\"",
            "gmailPartLowerForced": "The local part of the Gmail address must be all the readers",
            "googlemailNotAllowed": "Googlemail address is not supported",
            "verifyCodeRequire": "Please enter the verification code",
            "verifyCodeFormatErr": "The verification code format is incorrect",
            "passwordRequire": "Please enter your password",
            "passwordLengthRequire": "The password length must be greater than or equal to 10 digits",
            "passwordComplexRequire": "The password must contain at least three types of large letters, small letters, numbers, and special symbols.",
            "passwordAgainNotMatch": "Passwords entered twice mismatch",
            "passwordAgainRequire": "Please enter the confirm password",
            "inviteCodeRequire": "The invitation code is required"
        },
        "hCaptcha": {
            "passed": "Human-machine verification passed",
            "expired": "Please try again after expired",
            "challengeExpired": "Verification exceeds the time",
            "err": "Other errors"
        },
        "allRightsReserved": "All rights reserved",
        "securityAndLaws": "The website is protected and verified by hCaptcha and is subject to local laws."
    },
    "userSummary": {
        "title": "The instrument",
        "myPlan": "My Subscription",
        "shortcut": "Shortcuts",
        "timeLeft": "The subscription is valid and will expire in {msg}.",
        "toPurchase": "Buy Subscription",
        "tutorial": {
            "title": "View the tutorial",
            "content": "Learn how to use {name}"
        },
        "checkKey": {
            "title": "View keys",
            "content": "Quickly introduce keys to the client for use"
        },
        "renewPlan": {
            "title": "Continue Subscription",
            "content": "Make a continuous payment for your current subscription"
        },
        "appDown": {
            "title": "Application download",
            "content": "Download our applications from different platforms for better physical experience"
        },
        "support": {
            "title": "Encountered with a problem",
            "content": "If you encounter any problems, you can communicate with our people through the work order."
        },
        "haveTicket": "You have the {count} order to be processed",
        "toCheckTicket": "Go and check it out",
        "showAllKeys": "View all keys"
    },
    "userDocument": {
        "title": "Use the document",
        "description": "You can check different documents here, including but not limited to the website usage methods, precautions, operation procedures, etc. If you find any errors in the article, please submit the work order.",
        "searchPlaceholder": "Please enter the content you want to search (fuzzy search)",
        "searchBtn": "search",
        "noContentTitle": "No results",
        "noContentTitleHint": "There is no document that matches your search results or language, try changing to a keyword."
    },
    "newPurchase": {
        "title": "Buy Subscription",
        "description": "You can choose the subscription plan that suits you best here. If your balance is insufficient, please recharge first. The order will be reserved for you for 5 minutes.",
        "headerPlaceholder": "Choose the plan that suits you best",
        "purchaseBtn": "Order",
        "noLeft": "Insufficient amount of remaining",
        "monthPay": "Monthly payment",
        "moreMethod": "More options"
    },
    "newSettlement": {
        "err": "Mistake",
        "monthPay": "Monthly payment",
        "quarterPay": "Quarterly payment",
        "halfYearPay": "Half a year payment",
        "yearPay": "Annual payment",
        "payCycle": "Payment cycle",
        "couponPlaceholder": "Have a coupon?",
        "verifyBtn": "Verification",
        "orderTotalTitle": "Total order amount",
        "total": "Total",
        "submitOrder": "Submit an order",
        "coupon": "Coupons",
        "notify": {
            "passTitle": "Verification passed",
            "couponVerified": "Coupon valid",
            "couponInvalid": "Coupon is invalid",
            "couponIsNull": "The entered coupon code cannot be empty"
        }
    },
    "userProfile": {
        "title": "Personal Center",
        "myWallet": "My money bag",
        "walletSub": "Account balance (used only for consumption)",
        "alertPwd": "Modify key",
        "oldPwd": "Old key",
        "oldPwdSub": "Please enter the old password",
        "newPwd": "New key",
        "newPwdSub": "Please enter a new key",
        "newPwdAgain": "Confirm key",
        "newPwdAgainSub": "Please enter a new key again",
        "saveBtn": "save",
        "notify": "notify",
        "enableNotify": "Enable expiration notice",
        "deleteAccount": "Register Account",
        "deleteAccountSub": "Your account will be tagged as deleted, if you need to reuse our services, please re-register",
        "deleteBtn": "Register my account",
        "oldPwdVerifiedFailure": "Old password verification failed",
        "alertFailure": "Failed to modify key",
        "alertSuccess": "Modification was successful",
        "alertSuccessSub": "Please log in with your new password",
        "errorPwdFormat": "Password format error",
        "pwdNotMatch": "The password inputs are inconsistent",
        "oldPwdNotNull": "The old password cannot be empty",
        "toTopUp": "Go to recharge",
        "deleteMyTitle": "warn",
        "deleteMyContent": "Are you sure to delete your account? If you need to use the service, please re-register.",
        "deleteMyPositiveText": "Confirm deletion",
        "deleteMyNegativeText": "Cancel",
        "deletedSuccessMsg": "The account has been deleted and there will be a period of time later.",
        "deleteErrOccur": "An error encountered",
        "faAuth": "Two-step verification 2FA",
        "faAuthHint": "Two-step verification is a security mechanism that adds protection levels for logging in to your account. After entering the password, the user needs to complete the second step of identity verification, such as entering the verification code sent to the phone, using the identity verification application to generate the dynamic code, or confirming it through biological characteristics such as fingerprints.",
        "faAuthStatus": "Two-step verification status",
        "faEnabled": "Restarted",
        "faNotEnabled": "Not activated",
        "setup2Fa": "Setting up two-step verification",
        "disable2Fa": "Cancel the two-step verification",
        "unknownErr": "Unknown error",
        "disable2FaCancelled": "Canceled",
        "closed2FaHint": "The two-step verification has been closed, please add it again if necessary.",
        "setup2FaModal": {
            "followStep": "Add it to your verification device according to the prompts",
            "step1Title": "Follow the steps below to enable 2FA verification",
            "step1Context1": "1. You need to have a universal verification device on your mobile device",
            "step1Context2": "2. Click the Scan button on the verification device to scan the QR code here",
            "step1Context3": "3. This QR Code contains your verification information and unique keys. Please save it properly.",
            "step2Context1": "In order to ensure that your verification device can be used normally, we need to conduct a test.",
            "test2Fa": "Test",
            "cancel": "Cancel"
        },
        "deleteMyAccountModal": {
            "title": "Register Account",
            "contentLine1": "Accounting is an irreversible operation. Once you confirm that it is deleted, you will permanently lose access to the account, which means you will not be able to log in again and all data related to this account, including but not limited to your personal information, history, collection content, purchase records, etc., will not be able to access it again.",
            "contentLine2": "If you have ongoing business on our platform, such as unfinished orders, activities being participated in, subscription services, etc., these will be terminated or cancelled with your account removal, which may bring you corresponding losses. At the same time, the contacts and interaction information you and other users through this platform will no longer exist.",
            "contentLine3": "Please confirm your decision again. If you have any questions or questions, please contact our customer service and we will sincerely answer it for you. If you still want to delete your account, click the Confirm Delete button.",
            "inputHint1": "Enter \"",
            "inputHint2": "\" Continue.",
            "confirmDelete": "Confirm Delete"
        },
        "failure": "Encountered an error",
        "notLatestHint": "The personal information update failed and the current data may not be the latest.",
        "resetPassword": {
            "previousPasswordRequire": "Please enter the old password",
            "previousPasswordVerifiedFailure": "Old password verification failed",
            "passwordRequire": "Please enter your password",
            "passwordConflict": "The new password cannot be the same as the previous password",
            "passwordLengthRequire": "The password length must be greater than or equal to 10 digits",
            "passwordComplexRequire": "The password must contain at least three types of large letters, small letters, numbers, and special symbols.",
            "passwordAgainNotMatch": "Passwords entered twice mismatch",
            "passwordAgainRequire": "Please enter the confirm password",
            "updatePasswordFailure": "An error occurred when updating password"
        },
        "secondary": {
            "title": "Basic information",
            "card": {
                "avatar": "User head image",
                "name": "Username",
                "lastLoginAt": "Last login time",
                "accountStatus": "Account status"
            },
            "modify": {
                "uploadIconHint": "Upload",
                "alterAvatar": "Add/modify head images",
                "alterShallow": "* Click the head image or user name to add or modify (clearing is not allowed after setting)",
                "alterName": "Modify user name",
                "input": {
                    "label": "username",
                    "placeholder": "Enter user name",
                    "spaceIsNotAllowed": "Username verification does not pass",
                    "require": {
                        "p1": "Pure numbers are not allowed and numbers are opened",
                        "p2": "No space allowed",
                        "p3": "Longer than three"
                    }
                }
            },
            "mention": {
                "alterNameSuccess": "Username modification was successful",
                "alterNameErr": "If the user name fails, please try again later",
                "newNameIsNotValid": "Illegal user name",
                "click2SetName": "Click to set username",
                "fetchAvatarErr": "If the head image data fails, please try again later",
                "alterAvatarErr": "If the head image fails, please try again later",
                "success": "success",
                "alterAvatarSuccess": "The head image was modified successfully.",
                "uploadImageHint": "You can upload images as your new headline",
                "imageRequire": {
                    "title": "Notice",
                    "p1": "You can upload mainstream formats such as *.jpg(jpeg), *.png, *.webp, etc.",
                    "p2": "The image has a width ratio of 1:1 (square), if it is a different ratio, it will be cut into a square in the center, and the excess will be deleted.",
                    "p3": "The size of the image you upload will be set to 160px."
                },
                "click2Upload": "Click or drag the file to the area to upload",
                "uploadWarn": "*Please do not upload sensitive data, such as your bank card, credit card, password and security code."
            }
        }
    },
    "userKeys": {
        "myKeys": "My key",
        "description": "You can view the event status, expiration date, etc. for all your keys. If you need to set up a note for the key, please go to the activation record page.",
        "noKeys": "You have no valid purchase record yet",
        "keyDetail": "Key details",
        "keyId": "Key ID",
        "orderId": "Link Order ID",
        "clientId": "Activate client ID",
        "active": "During the validity period",
        "inActive": "Expired",
        "valid": "The key condition is normal",
        "invalid": "The key has been disabled",
        "isUsed": "Activated to use",
        "noUsed": "Not used yet",
        "releaseData": "Key generation date",
        "expirationData": "Date of Expiry",
        "none": "without",
        "authorizationFor": "Key Authorization",
        "hoverClickMention": "Click to copy to the clipboard",
        "copiedSuccessMessage": "The key copy has been successfully copied. Please refer to the document instructions to continue the operation.",
        "copyFailure": "Copying failure",
        "hoverCopiedSuccessMention": "Successfully copied"
    },
    "userOrders": {
        "myOrders": "My Order",
        "description": "All your orders will be displayed here, if you have unpaid orders, you can click on Continue Payment or cancel the order, and you can view the order details after completed orders.",
        "orderId": "#",
        "planName": "Subscription name",
        "planCycle": "Subscription cycle",
        "orderPrice": "Order amount",
        "orderStatus": "Order status",
        "createdAt": "Creation time",
        "operate": "operate",
        "showDetail": "Order details",
        "cancelOrder": "Cancel an order",
        "canceled": "Canceled",
        "period": {
            "monthPay": "Monthly payment",
            "quarterPay": "Quarterly payment",
            "halfYearPay": "Half a year payment",
            "yearPay": "Annual payment"
        },
        "orderStatusTags": {
            "success": "success",
            "cancelled": "Fail",
            "notPay": "Not paid"
        },
        "orderCancelled": "Order cancelled",
        "unknownErr": "Unknown error"
    },
    "userTopUp": {
        "topUp": "Account recharge",
        "description": "You can recharge your account here, support custom recharge amounts, and you can also pay attention to whether there is any discount information displayed below and use the mentioned payment method to get discounts.",
        "chooseTopUpAmount": "Select the recharge amount",
        "quickSelect": "Quickly choose",
        "customAmount": "Custom amount",
        "maxAmount": "Maximum amount: 10,000,000",
        "amountInputPlaceHolder": "Enter the amount to be recharged",
        "otherAmount": "Other amounts",
        "payMethod": "Payment method",
        "wechat": "WeChat Payment",
        "alipay": "Payment",
        "apple": "Apple Pay",
        "yourAmount": "Your amount",
        "discount": "Offers",
        "accountBalance": "Account balance",
        "balanceResult": "Amount of balance",
        "commitTopUp": "submit",
        "payMethodNotAllow": "Payment method is not available. Please select Other",
        "topUpIssueOccur": "Have trouble recharge?",
        "payIssueOccur": "Have trouble with payment?",
        "chatWithUs": "Contact customer service",
        "pay": "Pay",
        "qrCodeScannedSuccess": "QR code scan successfully",
        "orClickToApp": "Or click the bar and turn to the App and continue",
        "topUpSuccess": "Recharge successfully",
        "thankU": "Thank you for your support"
    },
    "userConfirmOrder": {
        "switchPlan": "Switch to subscribe",
        "cancelOrder": "Cancel an order",
        "yourPrice": "Your price",
        "couponOffset": "Coupon discount",
        "price": "Price",
        "submit": "submit",
        "monthPay": "Monthly payment",
        "quarterPay": "Quarterly payment",
        "halfYearPay": "Half a year payment",
        "yearPay": "Annual payment",
        "goodInfo": "Product Information",
        "cycleOrType": "Cycle/Type",
        "orderNumber": "Order number",
        "createdAt": "Creation date",
        "orderExpired": "Order has exceeded",
        "paySuccessfully": "Payment was successful, thank you for your support.",
        "balanceNotEnough": "If your balance is insufficient, please recharge first. The order is reserved for five minutes.",
        "orderErrorOccur": "An error occurred while ordering",
        "orderCancelled": "Order cancelled"
    },
    "paymentResultParts": {
        "goodInfoView": { "goodInfo": "Product Information" },
        "orderInfoView": { "orderInfo": "Order information" }
    },
    "orderPartUniversal": {
        "period": {
            "monthPay": "Monthly payment",
            "quarterPay": "Quarterly payment",
            "halfYearPay": "Half a year payment",
            "yearPay": "Annual payment"
        },
        "orderDataHex": {
            "goodInfo": "Product Information",
            "orderInfo": "Order information",
            "cycleOrType": "Cycle/Type",
            "orderNumber": "Order number",
            "createdAt": "Creation date",
            "amount": "Payment amount",
            "paidAt": "Payment time"
        }
    },
    "orderDetail": {
        "finished": "Completed",
        "finishedAndSuccessDescription": "Order has been successfully paid and opened",
        "useManual": "View the tutorial",
        "payPending": "Not paid yet",
        "pendingDescription": "Orders are kept at regular intervals and you can click the button below to continue payment.",
        "toPay": "Go to pay",
        "outDate": "Have expired",
        "outDateDescription": "Since you canceled the order or did not complete the payment within the specified time, the order has been cancelled and you can reselect your subscription.",
        "chooseNewPlan": "Choose a new subscription plan"
    },
    "userInvite": {
        "myInvite": "My invitation",
        "description": "If the administrator has enabled invitation rebates, you can see the specific rebate rules here. After generating an invitation link, share it with others to register and you can get the rebate.",
        "unit": "Number of people",
        "inviteCodeMgr": "Your invitation code",
        "generateInviteCode": "Generate random invitation codes",
        "faCodeManage": "Invitation code management",
        "email": "Email",
        "createdAt": "Registration time",
        "createFaCodeFailure": "Creation failed",
        "linkCopiedSuccess": "Link copied successfully",
        "generateFaCode": "Generate an invitation code",
        "flushFaCode": "Refresh the invitation code",
        "faCode": "Invitation code",
        "noFaCode": "You have no invitation code yet, please make it.",
        "faLink": "Invitation link",
        "generateFaCodePlease": "Please create an invitation code",
        "usersMyInvited": "Users I invite",
        "generateCodeConfirm": "Confirm Generate/Refresh",
        "generateCodeHint": "Please note that the invitation code cannot be closed after creation.",
        "positiveClick": "Confirm",
        "negativeClick": "Cancel"
    },
    "userTickets": {
        "description": "If you encounter any problems with use, please submit the work order here. Our technical support and customer service will reply and mark the specific reply time in the form below. You can click to view the work order to communicate with us.",
        "ticketId": "#",
        "ticketSubject": "Work order theme",
        "ticketUrgency": "Work order grade",
        "ticketContent": "Work order content",
        "ticketUrgencyLevel": { "low": "Low", "med": "middle", "high": "high" },
        "ticketStatus": "Work order status",
        "ticketCreatedAt": "Creation time",
        "lastResponse": "Last reply",
        "operate": "operate",
        "checkTicket": "View work order",
        "closeTicket": "Close the work order",
        "userTickets": "History Works",
        "addTicket": "New construction order",
        "ticketActive": "Not closed",
        "ticketInActive": "Closed",
        "form": {
            "ticketSubjectDescription": "Please enter the work order topic",
            "ticketUrgencyDescription": "Please select the urgency of the work order",
            "ticketBody": "Please enter the problem you want to solve, as comprehensively as possible.",
            "ticketNotFinished": "Please try the information on the complete work order"
        },
        "checkForm": "Please check if the form is complete",
        "cancel": "Cancel",
        "submit": "submit",
        "commitNewTicketSuccess": "Submit a new work order successfully",
        "commitNewTicketFailure": "Submit a new work order error",
        "noReply": "No reply yet",
        "noTickets": "You have not submitted a work order yet",
        "ticketCloseSuccess": "Work order closed successfully",
        "ticketCloseFailure": "Work order closing failed",
        "chatDialog": {
            "input": {
                "finished": "The work order has been concluded",
                "inputHere": "Enter the message to be sent"
            },
            "send": "Send",
            "missingToken": "Websocket connection cannot be created because tokens are missing.",
            "msgEmptyNotAllowed": "The content of the message cannot be empty",
            "accessNotPermit": "Illegal visit",
            "sendSuccess": "Send the message successfully"
        }
    },
    "userActivation": {
        "activateLog": "Activate the record",
        "description": "Here you can view your activation record and set the undetermined device, or set up the note information for each key and activation record.",
        "id": "#",
        "orderId": "Order number",
        "orderStatus": "Order",
        "createdAt": "Creation time",
        "operate": "operate",
        "userId": "User Id",
        "email": "Email",
        "keyId": "Key Id",
        "isBind": "Whether to activate",
        "active": "efficient",
        "inactive": "Ineffective",
        "requestAt": "Request time",
        "clientVersion": "Client side",
        "osType": "Operating system",
        "remark": "Note",
        "noRemark": "No note",
        "showDetail": "View details",
        "actions": "operate",
        "details": "Details",
        "keyContent": "Key content",
        "keyGeneratedAt": "Key generation time",
        "activateRequestAt": "Activate request time",
        "useIssueOccur": "Have trouble using it?",
        "chatWithUs": "Contact us",
        "cancelBind": "Cancel the decree",
        "alterRemark": "Modify Notes",
        "commitRemark": "submit",
        "updateSuccess": "Update successfully",
        "setRemark": "Set up the note information here"
    },
    "userAppDownload": {
        "title": "APP Download",
        "description": "",
        "common": {
            "title": "Download our app",
            "shallow2": "Get our applications for different clients",
            "shallow": "Using our application, you can access our services more easily, eliminating the complicated operation of the browser every time; you can find detailed installation requirements and tutorials in the document, including the operation environment, error reporting, etc. If you encounter other problems, please contact our technical support.",
            "noDownload": "We are sorry that there is no download available yet. Please try again later. If you have any questions, please submit a work order to contact our support service."
        },
        "suffix": {
            "p1": "*For applications on macOS platform, please use macOS14 and above and cooperate with Apple chips (M1 or higher).",
            "p2": "The information and use displayed by this app should comply with local usage regulations. Whether the app is allowed should also depend on your company's IT specifications."
        },
        "card": {
            "common": { "welcome": "Welcome" },
            "mobile": {
                "designFor": "Design for mobile terminal",
                "designShallow": "You can get our mobile app here",
                "iosDownloadShallow": "Download IOS Client",
                "androidDownloadShallow": "Download Android Client"
            },
            "desktop": {
                "designFor": "Design for desktop",
                "designShallow": "You can get our desktop mobile application here",
                "windowsDownloadShallow": "Download Windows Client",
                "osxDownloadShallow": "Download macOS client",
                "linuxDownloadShallow": "Download Linux client"
            }
        },
        "downloadDisabledHint": "Sorry, the app download is temporarily unavailable or disabled by the administrator. If you need it, please contact our technology at the ticket office. Support has been obtained.",
        "windows": {
            "title": "Windows NT",
            "shallow": "This client is suitable for Windows operating systems with NT kernels, please check the documentation page for compatibility support."
        },
        "osx": {
            "title": "macOS (OSX)",
            "shallow": "This client is suitable for the macOS (OSX) operating system of the Darwin kernel, please check the documentation page for compatibility support."
        },
        "linux": {
            "title": "Linux",
            "shallow": "This client is suitable for various distributions of the Linux kernel. Due to different distribution series, please check the documentation page for compatibility support."
        },
        "android": {
            "title": "Android",
            "shallow": "This client is suitable for mobile devices equipped with Google Android operating system, please check the documentation page for compatibility support."
        },
        "ios": {
            "title": "IOS",
            "shallow": "This client is suitable for mobile devices equipped with Apple IOS operating system, please check the documentation page for compatibility support."
        }
    },
    "welcome": {
        "A": {
            "aboutUs": "About Us",
            "pricing": "Order price",
            "login": "Log in",
            "register": "Register Account",
            "welcomeTo": "Welcome",
            "welcomeToSub": "Passing through the long tunnel on the county is the Snow Country. Under the night sky, the earth was white and the train stopped in front of the signal station. \"Here, Kawabata Yasunari kicked off the Snow Country with a very stingy concise text.",
            "startingUse": "Get started",
            "whyUs": "Why choose us",
            "whyUsSub": "\"When passing through the long tunnel in the prefecture, it is the Snow Country. Under the night sky, the earth is white, and the train stops in front of the signal station.\" Here, Yasunari Kawabata kicks off the prelude to \"Snow Country\" with almost stingy words.",
            "browseSafe": "Browse safe",
            "browseSafeSub": "The excellent firewall filter system can effectively prevent online fish and malicious websites",
            "encrypt": "End-to-end encryption",
            "encryptSub": "Double-way SSL and end-to-end encryption protect your privacy and security, even servers cannot read your information",
            "mgr": "Efficient management",
            "mgrSub": "One user interface manages all keys, with complete and rich management functions, and no need to worry about subscribing to ejaculate problems",
            "fast": "Convenient and fast",
            "fastSub": "Provide complete API files for WebApps or embedded in third-party software",
            "fastLink": "Quick link",
            "subscribeUs": "Follow us",
            "filingsCode": "File number {code}"
        }
    },
    "pagination": {
        "perPage10": "10 data/page",
        "perPage20": "20 data/page",
        "perPage50": "50 data/page",
        "perPage100": "100 data/page"
    },
    "modal": { "cancel": "Cancel", "confirm": "confirm" },
    "week": {
        "monday": "Monday",
        "tuesday": "Tuesday",
        "wednesday": "Wednesday",
        "thursday": "Thursday",
        "friday": "Friday",
        "saturday": "Saturday",
        "sunday": "Sunday"
    },
    "period": {
        "month": "Monthly payment",
        "quarter": "Quarterly payment",
        "halfYear": "Half a year payment",
        "year": "Annual payment"
    },
    "operate": {
        "cancel": "Cancel",
        "confirm": "confirm",
        "update": "renew",
        "add": "Add to"
    },
    "notFound": {
        "title": "404 The page does not exist",
        "description": "We cannot find the page you requested, it may have been deleted or linked with errors. If you think this is an error, please submit the work order to contact us.",
        "p1": "It will return to the home page after {sec}s, and if your browser does not respond, click the button below.",
        "manualBack": "Return to home page"
    },
    "forbidden": {
        "title": "403 No rights",
        "description": "You may not have sufficient permission to access this page. If you think this is an error, please submit the work order to contact us."
    }
}
