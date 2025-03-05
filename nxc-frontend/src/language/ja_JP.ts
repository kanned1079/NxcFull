export default {
    "commonHeader": {
    "menuTxt": "メニュー",
        "userData": "ユーザー情報",
        "editUserData": "ユーザー情報を編集します",
        "logout": "ログアウト"
},
    "commonAside": {
    "admin": {
        "dashboard": "テーブルボード",
            "queueMonit": "サービス側の監視",
            "settings": "設定",
            "systemConfig": "システム設定",
            "paymentConfig": "支払い設定",
            "themeConfig": "テーマ設定",
            "server": "サーバ",
            "privilege": "権限管理",
            "finance": "ファイナンス",
            "subscription": "サブスクリプション管理",
            "coupon": "クーポン管理",
            "order": "注文管理",
            "activate": "レコードをアクティブにします",
            "key": "キー管理",
            "user": "ユーザー",
            "userMgr": "ユーザー管理",
            "notice": "発表管理",
            "ticket": "ワークオーダー管理",
            "doc": "知識データベース管理"
    },
    "user": {
        "dashboard": "テーブルボード",
            "document": "ドキュメントを使用します",
            "app": "アプリのダウンロード",
            "subscription": "購読する",
            "purchase": "サブスクリプションを購入します",
            "surplus": "私の鍵",
            "fiance": "ファイナンス",
            "topUp": "補充します",
            "myOrder": "私の注文",
            "myInvite": "私の招待",
            "user": "ユーザー",
            "profile": "パーソナルセンター",
            "support": "私の作業注文",
            "activateLog": "レコードをアクティブにします"
    }
},
    "adminViews": {
    "common": {
        "fetchDataSuccess": "データの取得に成功しました",
            "fetchDataFailure": "データが失敗した場合は、後でもう一度やり直してください",
            "addSuccess": "正常に追加されました",
            "addFailure": "追加が失敗した場合は、後で試してください",
            "updateSuccess": "変更は成功しました",
            "updateFailure": "変更が失敗した場合は、後でもう一度やり直してください",
            "deleteSuccess": "削除に正常に",
            "deleteFailure": "削除が失敗した場合は、再試行してください",
            "confirm": "確認する",
            "cancel": "キャンセル",
            "success": "操作は成功しました",
            "failure": "操作に失敗しました",
            "NotAllowed": "違法な形式",
            "checkForm": "フォームを確認してください",
            "unknownErr": "不明なエラー",
            "dialog": { "delete": "あなたはそれを削除したいことを知っていますか？" },
        "yes": "はい",
            "no": "いいえ"
    },
    "login": {
        "secureCard": {
            "title": "安全検査",
                "securePath": "安全な道",
                "hint": "システムセキュリティを確保するには、管理者ログインページを入力する前にセキュリティパスを入力する必要があります。 以下の入力ボックスにセキュリティパスを入力してください。",
                "placeholder": "安全な道を入力してください",
                "checkBtn": "診る",
                "rememberPath": "安全な道を覚えておいてください"
        },
        "card": {
            "back": "ホームページに戻ります",
                "toAdminCenter": "管理センターにログインします",
                "emailAddr": "管理者のメール",
                "emailInputArea": {
                "title": "管理者のメール",
                    "placeholder": "電子メールアドレス"
            },
            "passwordInputArea": {
                "title": "管理者パスワード",
                    "placeholder": "パスワード"
            },
            "login": "ログイン",
                "forgetPassword": "パスワードを忘れてください",
                "formNotPassed": "フォーム検証に失敗します"
        },
        "mention": {
            "title": "ヒント",
                "description": "このページは、管理者がアクセスできる場合のみ、管理者の許可がない場合は、以下のリンクをクリックしてホームページに戻ります。"
        },
        "message": {
            "passwordErr": "パスワードが正しくありません",
                "adminNotExist": "管理者は存在しません",
                "noPrivilege": "許可なくアクセスします",
                "authPassed": "検証が合格しました",
                "authFailure": "検証に失敗しました",
                "otherErr": "その他のエラー",
                "pathCheckPassed": "セキュリティパスチェックが通過しました",
                "pathCheckFailure": "間違ったセキュリティパス",
                "rememberSecureMention": "バックエンド管理のセキュリティを確保するために、これがあなたのプライベートコンピューターではないかどうかを確認しないでください。"
        }
    },
    "summary": {
        "cockpit": "テーブルボード",
            "systemConfig": "システム設定",
            "paymentConfig": "支払い設定",
            "planMgr": "サブスクリプション管理",
            "userMgr": "ユーザー管理",
            "orderMgr": "注文管理",
            "keyMgr": "キー管理",
            "incomeText": "昨日の収入 /月収",
            "pendingTicket": "処理する{nums}注文があります",
            "toHandle": "行ってチェックしてください",
            "apiAccessCard": "1週間以内にAPIインターフェイス訪問の数",
            "apiAccessCardHint": "このデータは、現在のAPIアクセスを理解するためだけのものであり、サーバーのパフォーマンスを表していません。",
            "incomeWeek": "1週間以内に収入額",
            "incomeCardHint": "これは1週間にわたる収入額のチャートです。これにより、キャッシュがクリアされた場合に不正確な表示が発生します。",
            "core": "コア",
            "reqErr": "エラーが発生しました",
            "reqErrHint": "概要情報の取得中にエラーが発生したため、リクエストが完了できないため、後で再試行できません。",
            "userCard": {
            "title": "ユーザーの概要",
                "allRegisteredUsers": "一般的な登録ユーザー",
                "activeUsers": "ライブユーザー",
                "inactiveUsers": "非アクティブユーザー",
                "blockedUsers": "禁止または払い戻し"
        },
        "general": {
            "title": "一般的に",
                "localTime": "ブラウザ時間",
                "osType": "オペレーティングシステムの種類",
                "appName": "アプリケーション名",
                "appUrl": "アプリケーションURL",
                "currency": "コインユニット",
                "allowRegister": "登録を許可します"
        },
        "system": {
            "title": "システム構成",
                "axiosAddr": "HTTPバックエンドアドレス",
                "wsAddr": "WebSocketバックエンドアドレス",
                "serverTime": "サーバー時間",
                "uptime": "運用時間",
                "gatewayStatus": "APIネットワークステータス",
                "dbStatus": "データベースステータス",
                "redisStatus": "Redisステータス",
                "serverOsType": "サーバーオペレーティングシステムの種類",
                "serverOsArch": "サーバーオペレーティングシステムアーキテクチャ",
                "runMode": "実行モード",
                "cpuNums": "ネットワークサーバーのCPUコアの数",
                "numCgoCall": "ごみ収集の数",
                "time": "二流",
                "paymentMethods": "支払い方法を有効にします",
                "runOK": "正常に実行します",
                "runErr": "異常な行動",
                "checkServer": "バックエンドサーバーの環境構成を確認してください",
                "stopRegisterHint": "新しいユーザー登録を無効にしているようです",
                "toSetting": "設定に頼ります"
        }
    },
    "queueMonit": {
        "title": "サービス側の監視",
            "headerHint": "ピークネットワーク中に頻繁にこのページに留まらないでください。",
            "latency": {
            "title": "サーバーの遅延",
                "retry": "もう一度テストします",
                "hint": "*ここでのリクエストとは、クライアントがサーバーにリクエストを送信した後のクライアントへの応答を成功させるために必要な時間を指します。",
                "unit": "ミリ秒",
                "level": {
                "l1": {
                    "title": "素晴らしい",
                        "description": "これは非常に良いネットワークの状況であり、あなたはほとんど遅れを感じることはありません。"
                },
                "l2": {
                    "title": "普通",
                        "description": "これはほとんどの場合ネットワークの状況であり、ユーザーはほとんど遅れを感じないと感じています。"
                },
                "l3": {
                    "title": "貧しい",
                        "description": "この場合、ユーザーはわずかな遅延または遅延を感じることがあります。"
                },
                "l4": {
                    "title": "違い",
                        "description": "明らかな遅延を感じ、ユーザーエクスペリエンスに影響を与えることができます。"
                },
                "l5": {
                    "title": "非常に悪い",
                        "description": "明らかな遅延があり、負荷速度が遅くなるか、更新することができないため、ユーザーの相互作用と経験に深刻な影響を与えます。"
                },
                "offline": {
                    "title": "サーバーオフ",
                        "description": "サーバーに接続したり、リクエストエラーを処理したりすることができません。構成が正しいかどうかを確認してください。"
                }
            }
        },
        "api": {
            "title": "過去7日間のAPI要求ステータス",
                "items": {
                "ok": { "title": "成功（statusok）", "unit": "二流" },
                "notFound": {
                    "title": "ステータスパスエラー（404）",
                        "unit": "二流"
                },
                "unAuthorized": { "title": "不正アクセス（401）", "unit": "二流" },
                "login2reg": { "title": "ログイン /登録", "unit": "二流" }
            }
        },
        "log": {
            "title": "ジャーナルレコード",
                "deleteLogMsg": "ジャーナルを削除{nums}",
                "deleteLogErr": "削除ジャーナルに失敗しました",
                "logTableRows": "総日記記録",
                "logTableSize": "ブログテーブルはスペースを占有します",
                "unit": { "lines": "わかりました", "mb": "メガキャラクター" },
            "deleteLog": "ジャーナルを削除します",
                "exportCsv": "CSVをエクスポートします",
                "deleteLogHint": "これにより、1週間前にすべてのジャーナルが削除されます。",
                "warn": {
                "title": "あなたはあなたがあなたの日記を削除したいことを知っていますか？",
                    "description": "リストはすぐに削除され、この操作を取り消すことはできません。"
            },
            "export": {
                "title": "CSVファイルをエクスポートします",
                    "description": "これにより、CSVファイルとしてエクスポートし、ダウンロードを開く許可がない場合は、[確認]ボタンをクリックします。"
            },
            "table": {
                "id": "＃",
                    "method": "リクエスト方法",
                    "path": "ルートを求めてください",
                    "code": "ステータスコード",
                    "clientIp": "クライアントIP",
                    "responseTime": "処理時間",
                    "requestAt": "リクエスト時間"
            }
        }
    },
    "systemConfig": {
        "title": "システム設定",
            "common": {
            "success": "構成を正常に更新しました",
                "err": "更新構成に失敗しました"
        },
        "site": {
            "common": { "title": "サイト" },
            "appName": {
                "title": "サイト名",
                    "shallow": "サイト名が必要な場所を表示するために使用されます。",
                    "placeholder": "サイト名"
            },
            "appSubName": {
                "title": "サイトサブタイトル",
                    "shallow": "通常、メインタイトルの下に表示されます。",
                    "placeholder": "字幕"
            },
            "appDescription": {
                "title": "サイトの説明",
                    "shallow": "サイトが必要な場所を表示するために使用されます。",
                    "placeholder": "サイトの説明"
            },
            "appUrl": {
                "title": "サイトURL",
                    "shallow": "最新のWebサイトアドレスは、必要な場合にURLに表示されます。",
                    "placeholder": "サイトURL"
            },
            "forceHttps": {
                "title": "強制https",
                    "shallow": "サイトがHTTPS、CDN、または抗ジェネレーションを使用して強制HTTPSを有効にしない場合。"
            },
            "logoUrl": {
                "title": "ロゴ",
                    "shallow": "ロゴが必要な場所を表示するために使用されます。",
                    "placeholder": "ロゴのURLアドレス"
            },
            "subscribeUrl": {
                "title": "サブスクライブURL",
                    "shallow": "サブスクリプションに使用して、サイトURLとして空白のままにします。 ランダムにフェッチするために複数のサブスクリプションURLが必要な場合は、コンマを使用してそれらを分割してください。",
                    "placeholder": "サブスクライブURL"
            },
            "tosUrl": {
                "title": "ユーザー（TOS）URLの用語",
                    "shallow": "ユーザー（TOS）の条件にジャンプするために使用されます",
                    "placeholder": "ユーザーの利用規約アドレス"
            },
            "stopRegister": {
                "title": "新しいユーザー登録を停止します",
                    "shallow": "オンになってから登録することはできません。"
            },
            "inviteRequire": {
                "title": "強制招待状",
                    "shallow": "開いた後、新しいユーザーに登録するときは招待コードを入力する必要があります。"
            },
            "trialSubscribe": {
                "title": "裁判に登録します",
                    "shallow": "試してみたいサブスクリプションを選択してください。"
            },
            "trialTime": {
                "title": "試用時間（時間）",
                    "shallow": "新しいユーザーがサインアップしたときの試用時間を購読してください。"
            },
            "currency": {
                "title": "通貨ユニット",
                    "shallow": "ディスプレイの使用にのみ使用され、変更後、システム内のすべての通貨ユニットが変更されます。",
                    "placeholder": "人民元"
            },
            "currencySymbol": {
                "title": "通貨シンボル",
                    "shallow": "ディスプレイの使用にのみ使用され、変更後、システム内のすべての通貨ユニットが変更されます。",
                    "placeholder": "¥"
            }
        },
        "security": {
            "common": { "title": "セキュリティ設定" },
            "emailVerify": {
                "title": "電子メールの確認",
                    "description": "開いた後、ユーザーは電子メールの確認を実行する必要があります。"
            },
            "gmailAlias": {
                "title": "Gmail複数のエイリアスを禁止します",
                    "description": "開いた後、複数のGmailエイリアスは登録されません。"
            },
            "safeMode": {
                "title": "セーフモード",
                    "description": "開いた後、サイトURLを除くこのサイトへのドメイン名アクセスは403になります。"
            },
            "adminPath": {
                "title": "背景パス",
                    "description": "バックグラウンド管理パスは、変更後に元の管理パスを変更します。",
                    "placeholder": "https://x.com/logo.jpeg"
            },
            "emailWhitelist": {
                "title": "電子メール接尾辞ホワイトリスト",
                    "description": "登録は、リストにメールボックスの接尾辞を開いた後にのみ許可されます。"
            },
            "recaptcha": {
                "title": "反ロボット",
                    "description": "有効にすると、ロボットを防ぐためにHcaptchaが有効になります。"
            },
            "hCaptchaSiteKey": {
                "title": "hcaptcha sitekey",
                    "description": "SiteKeyは、hcaptchaサーバーにWebサイト番号を識別するために要求するために使用されます",
                    "placeholder": "A3CA066C-0EA0-42FE-BCD2-23F4AB48D528"
            },
            "ipRegisterLimit": {
                "title": "IP登録制限",
                    "description": "開設後、IP登録口座が規則要件を満たしている場合、登録は制限されていることに注意してください。"
            },
            "registerTimes": {
                "title": "頻度",
                    "description": "登録数に達した後、罰はアクティブになります。",
                    "placeholder": "入力してください"
            },
            "lockTime": {
                "title": "罰の時間（分）",
                    "description": "ペナルティ時間が経過するまで待つ必要があります。",
                    "placeholder": "入力してください"
            }
        },
        "frontend": {
            "common": { "title": "パーソナライズ" },
            "themePropColor": {
                "default": "デフォルト",
                    "darkBlueDay": "濃い青",
                    "milkGreenDay": "ミルクグリーン",
                    "bambooGreen": "ruozhu",
                    "mistyPine": "フォグパイン",
                    "glacierBlue": "氷河青",
                    "grayTheme": "グレー"
            },
            "sidebarStyle": {
                "title": "サイドバースタイル",
                    "shallow": "サイドバーの色を設定します。"
            },
            "headerStyle": {
                "title": "ヘッドスタイル",
                    "shallow": "上部に色を設定します。"
            },
            "themeColor": {
                "title": "テーマの色",
                    "shallow": "Webページ全体のテーマ色を設定します。"
            },
            "background": {
                "title": "背景",
                    "shallow": "バックグラウンドログインページに表示されます。",
                    "placeholder": "https://x.com/logo.jpeg"
            }
        },
        "inviteAndRebate": {
            "common": { "title": "支払いとリベート" },
            "inviteRebateEnable": {
                "title": "ユーザーリベートを有効にします",
                    "description": "有効にすると、招待されたユーザーが充電すると、以下の充電比率に従ってユーザーがリベートされます。"
            },
            "inviteRebateRate": {
                "title": "リベート比",
                    "description": "リベート量比を設定します。",
                    "placeholder": "リベート比率を入力してください"
            },
            "discountInfo": {
                "title": "割引情報",
                    "description": "オファー情報を設定すると、トップアップページの上部に表示されます。",
                    "placeholder": "割引情報を設定します"
            },
            "inviteInfo": {
                "title": "招待情報",
                    "description": "招待情報を設定すると、リベート比を表示するためにユーザー招待状ページに表示されます。",
                    "placeholder": "リベート情報を設定します"
            }
        },
        "welcome": {
            "common": { "title": "ホームページ情報" },
            "homeDescription": {
                "title": "ホームページの説明",
                    "description": "ホームページの簡単な説明を設定します。",
                    "placeholder": "ホームページの説明コンテンツを入力してください"
            },
            "whyChooseUs": {
                "title": "なぜ私たちを選ぶのか",
                    "description": "選択した理由についての説明を設定します。",
                    "placeholder": "詳細な説明を入力してください"
            },
            "bilibiliLink": {
                "title": "Bilibiliの公式リンク",
                    "description": "Bilibiliの公式アカウントのリンクアドレスを設定します。",
                    "placeholder": "https://space.bilibili.com/xxxx"
            },
            "youtubeLink": {
                "title": "YouTubeの公式リンク",
                    "description": "公式YouTubeアカウントのリンクアドレスを設定します。",
                    "placeholder": "https://youtube.com/channel/xxxxx"
            },
            "instagramLink": {
                "title": "Instagramの公式リンク",
                    "description": "公式Instagramアカウントのリンクアドレスを設定します。",
                    "placeholder": "https://instagram.com/xxxx"
            },
            "wechatLink": {
                "title": "WeChatパブリックアカウントリンク",
                    "description": "WeChatパブリックアカウントのリンクアドレスを設定します。",
                    "placeholder": "WeChatパブリックリンクを入力してください"
            },
            "filingNumber": {
                "title": "登録番号",
                    "description": "サイトの登録番号を設定します。",
                    "placeholder": "例：Guangdong ICP No. 12345678"
            },
            "pageSuffix": {
                "title": "サイト接尾辞",
                    "description": "タイトル表示用のサイト名接尾辞を設定します。",
                    "placeholder": "例： - サイト名"
            }
        },
        "sendMail": {
            "common": {
                "title": "メール設定",
                    "warning": "このページの構成を変更する場合、キューサービスと逆方向を再起動する必要があります。 さらに、このページの構成優先度は、.ENVのメール構成よりも高くなっています。",
                    "sendTestMailTitle": "テストメールを送信します",
                    "sendTestMailShallow": "電子メールは、現在のログイン管理者のメールアドレスに送信されます",
                    "sendTestMailTo": "にテストメールを送信します",
                    "sending": "メールを送信します",
                    "success": "成功",
                    "receiverAddr": "受け取った住所",
                    "senderHost": "配信サーバー",
                    "senderPort": "ポートを送信します",
                    "senderEncrypt": "メッセージを送信するための暗号化方法",
                    "senderUsername": "クレジット名",
                    "sendErr": "電子メールの失敗を送信します"
            },
            "smtpServerAddress": {
                "title": "SMTPサーバーアドレス",
                    "shallow": "電子メールサービスプロバイダーが提供するサービスアドレス",
                    "placeholder": "サーバーアドレスを入力してください"
            },
            "smtpServerPort": {
                "title": "SMTPサービスポート",
                    "shallow": "一般的なポートは25、465、587です",
                    "placeholder": "ポート番号を入力してください"
            },
            "smtpEncryption": {
                "title": "SMTP暗号化方法",
                    "shallow": "465ポート暗号化法は一般にSSLであり、587ポート暗号化方法は一般にTLSです。",
                    "placeholder": "暗号化方法を入力してください"
            },
            "smtpAccount": {
                "title": "SMTPアカウント",
                    "shallow": "電子メールサービスプロバイダーが提供するアカウント",
                    "placeholder": "アカウント番号を入力してください"
            },
            "smtpPassword": {
                "title": "SMTPパスワード",
                    "shallow": "電子メールサービスプロバイダーが提供するパスワード",
                    "placeholder": "パスワードを入力してください"
            },
            "senderAddress": {
                "title": "住所を送信します",
                    "shallow": "メールサービスプロバイダーが提供する送信アドレス",
                    "placeholder": "送信アドレスを入力してください"
            },
            "emailTemplate": {
                "title": "電子メールテンプレート",
                    "shallow": "ドキュメントで電子メールテンプレートをカスタマイズする方法を表示できます",
                    "placeholder": "電子メールテンプレートを選択してください"
            }
        },
        "notice": {
            "common": { "title": "通知設定" },
            "displayName": {
                "title": "表示名",
                    "shallow": "フロントエンドページの表示のみ",
                    "placeholder": "一般通知インターフェイス1"
            },
            "barkEndpoint": {
                "title": "樹皮アクセスポイント",
                    "shallow": "樹皮サーバーバックエンドAPIアドレス",
                    "placeholder": "https：// <ip>：<port>/<Secret-Key>"
            },
            "barkGroup": {
                "title": "樹皮グループ",
                    "shallow": "クライアントによって表示されるグループ名",
                    "placeholder": "ウェブ"
            }
        },
        "appDownload": {
            "common": {
                "title": "アプリ",
                    "hint": "自分のクライアント（アプリ）のバージョン管理と更新用"
            },
            "enabled": {
                "title": "負荷を開いてドロップします",
                    "shallow": "有効にする場合は、ユーザーがダウンロードページにアクセスできるようにします"
            },
            "windows": {
                "title": "Windows",
                    "shallow": "Windowsバージョン番号とアドレスをダウンロードします",
                    "placeholder": "https://xxxx.com/xxx.exe"
            },
            "macos": {
                "title": "macos",
                    "shallow": "MacOSバージョン番号とダウンロードアドレス",
                    "placeholder": "https://xxxx.com/xxx.dmg"
            },
            "linux": {
                "title": "Linux",
                    "shallow": "Linuxバージョン番号とアドレスをダウンロードします",
                    "placeholder": "https://xxxx.com/xxx.deb"
            },
            "android": {
                "title": "アンドロイド",
                    "shallow": "Androidバージョン番号とダウンロードアドレス",
                    "placeholder": "https://xxxx.com/xxx.apk"
            },
            "ios": {
                "title": "iOS",
                    "shallow": "iOSバージョン番号とダウンロードアドレス",
                    "placeholder": "https://xxxx.com/xxx.ipk"
            }
        }
    },
    "payConfig": {
        "title": "支払い設定",
            "description": "サポートされているすべての支払い方法は、現在、支払いバッグの支払いのみがサポートされていますが、最初に他の支払い方法を構成することもできます。",
            "attention": {
            "title": "注意すべきこと",
                "point1": "有効にする前に、支払い方法情報を構成することが本当に重要です。",
                "point2": "支払い方法を変更する場合、「---」として表示される場合、オプションが設定されていて、再度変更する必要がない場合は、保存するために新しい値を入力します。"
        },
        "common": {
            "detail": "{method}構成",
                "fillAttention": "安全性を確保するために、既存の構成を作成または上書きするための補充の詳細情報は表示されません。",
                "discountAmount": "割引額（ユーザープロンプト情報は、システムによって設定された「支払いとリベート」で設定できます）",
                "saveConfigBtnHint": "保存",
                "cancelBtnHint": "キャンセル",
                "saveSuccess": "構成は正常に保存されました",
                "alterSuccess": "構成の変更が成功しました",
                "discountPlaceholder": "割引額（充電額が割引額よりも大きい場合）",
                "saveOrAlterFailure": "不明なエラー"
        },
        "alipay": {
            "title": "支払い",
                "config": {
                "appId": "アプリケーションID",
                    "appPublicKeyCertContent": "アプリケーションキー証明書のコンテンツ",
                    "appPrivateKey": "アプリケーションプライバシーキー",
                    "alipayRootCert": "支払いバッグ証明書帳",
                    "alipayPublicKeyCertContent": "支払いバッグキー証明書のコンテンツ"
            }
        },
        "wechat": {
            "title": "wechatの支払い",
                "config": {
                "mchId": "マーチャントID",
                    "mchCertSerial": "マーチャント証明書のシリアル番号",
                    "apiV3Key": "API V3キー",
                    "privateKey": "プライバシーキー"
            }
        },
        "apple": {
            "title": "Apple Pay",
                "config": {
                "issuerId": "発行者ID",
                    "bundleId": "バンドルID",
                    "privateKeyId": "秘密鍵ID",
                    "privateKeyContent": "プライバシーコンテンツ"
            }
        },
        "addPaymentMethod": "支払い方法を追加します",
            "enableBtnHint": "始める",
            "disableBtnHint": "無効",
            "setupPaymentMethod": "構成"
    },
    "themeConfig": {
        "title": "トピック構成",
            "using": "使用、",
            "setAsCurrent": ""
    },
    "groupMgr": {
        "title": "権限管理",
            "description": "適切なグループは、さまざまなサブスクリプションレベルをマークするために使用されます。同じレベルで、適切なグループで異なるサイズをサブスクライブするために、簡単に管理することができます。",
            "common": {
            "addNewGroup": "新しい許可グループ",
                "alterGroupName": "許可グループの名前を変更します",
                "addConfirmed": "追加することを確認します",
                "alterConfirmed": "変更を確認します",
                "cancel": "キャンセル",
                "addSuccess": "正常に追加されました",
                "addFailure": "障害を追加します",
                "alterSuccess": "変更許可グループは成功しました",
                "alterFailure": "変更許可グループは失敗しました",
                "delSuccess": "削除に正常に",
                "delFailure": "削除が失敗しました",
                "delMention": "リストはすぐに削除され、この操作を取り消すことはできません。 サブスクリプションプランに関連するグループは、注意して削除する必要があります。",
                "delNotAllowed": "この適切なグループには関連するリソースがあり、削除することはできません。"
        },
        "groupId": "権利グループID",
            "groupName": "許可グループの名前",
            "groupPlaceholder": "許可グループ名を入力します",
            "userCount": "ユーザー数",
            "planCount": "サブスクリプション数",
            "operate": "動作します",
            "editBtnHint": "編集",
            "deleteBtnHint": "消去"
    },
    "docMgr": {
        "title": "知識データベース管理",
            "description": "ここでは、ユーザーに説明ファイルを書くことができます。これは、ソフトウェアデザインマニュアル、チュートリアルまたは予防策などです。",
            "addDoc": "新しいドキュメントを追加します",
            "addSuccess": "新しいドキュメントを正常に追加します",
            "addFailure": "ファイルの追加に失敗しました",
            "titleNotEmpty": "ドキュメントのタイトルを空にすることはできません",
            "table": {
            "docId": "＃",
                "isShow": "表示するかどうか",
                "sortAs": "選別",
                "lang": "言語",
                "category": "カテゴリ",
                "title": "タイトル",
                "createdAt": "作成時間",
                "updatedAt": "時間を更新します",
                "operate": "動作します",
                "edit": "編集",
                "delete": "消去"
        },
        "form": {
            "add": "ドキュメントを追加します",
                "edit": "ドキュメントを編集します",
                "cancel": "キャンセル",
                "confirm": "確認する",
                "addBtn": "に追加",
                "editBtn": "改訂",
                "title": {
                "title": "ファイルタイトル",
                    "placeholder": "ファイルのタイトルを入力します"
            },
            "sort": {
                "title": "選別",
                    "placeholder": "入力ファイルの順序レベルは高くなるほど、優先度が高くなります。"
            },
            "category": {
                "title": "カテゴリ",
                    "placeholder": "ドキュメントは、このフィールドに従ってカテゴリに表示されます"
            },
            "lang": {
                "title": "ドキュメント言語",
                    "placeholder": "ドキュメント言語を選択します"
            }
        }
    },
    "planMgr": {
        "title": "サブスクリプション管理",
            "description": "ここでは、新しいサブスクリプションプランを追加し、サブスクリプションプランが既に含まれている説明、価格、バランス、許可グループを変更するなどします。",
            "addNewPlan": "新しいサブスクリプションを追加します",
            "table": {
            "id": "＃",
                "isSale": "販売を有効にします",
                "isRenew": "継続的な料金を許可します",
                "sort": "ソートレベル",
                "group": "いわゆる権利グループ",
                "name": "名前",
                "nums": "量",
                "residue": "額",
                "operate": "動作します",
                "operateBtn": { "update": "改訂", "delete": "消去" }
        },
        "mention": {
            "common": {
                "success": "成功",
                    "failure": "失敗",
                    "delMention": "サブスクリプションプランがすでに販売されている場合は、削除するように注意してください。"
            }
        },
        "form": {
            "title": "サブスクリプションを追加します",
                "items": {
                "name": {
                    "title": "サブスクリプション名",
                        "placeholder": "サブスクリプションプランの名前を入力します"
                },
                "describe": {
                    "title": "サブスクリプションの説明",
                        "placeholder": "サブスクリプションの説明を入力します（サポートマークダウン）"
                },
                "isSale": { "title": "販売を有効にします" },
                "isRenew": { "title": "継続料金を有効にします" },
                "group": {
                    "title": "いわゆる権利グループ",
                        "placeholder": "正当なグループを選択します"
                },
                "capacityLimit": {
                    "title": "許可されたユーザーの最大数",
                        "placeholder": "許可されたユーザーの最大数"
                },
                "planResidue": { "title": "残りの金額", "placeholder": "残りの金額" },
                "sort": {
                    "title": "選別",
                        "placeholder": "フロントエンドのソート用"
                },
                "periodPlaceholder": {
                    "month": "毎月の支払い価格を入力します",
                        "quarter": "四半期ごとの有料価格を入力します",
                        "halfYear": "半年の支払い価格を入力します",
                        "year": "年間支払い価格を入力します"
                }
            }
        }
    },
    "couponMgr": {
        "title": "クーポン管理",
            "description": "ここでは、特定のフェスティバルなどにいくつかのクーポンを追加できます。これにより、ユーザーは注文を行うときに使用できるようにし、設定した割合に応じて割引を提供できます。",
            "fetchErr": "データが失敗した場合は再試行します",
            "fetchPlanKvFailurese": "サブスクリプションプランリストを取得しました",
            "table": {
            "id": "＃",
                "enabled": "有効にするかどうか",
                "name": "名前",
                "code": "クーポン",
                "percentOff": "割引情報",
                "capacity": "総量",
                "residue": "残りの金額",
                "startTime": "時間を有効にします",
                "endTime": "終了時間",
                "actions": "動作します",
                "edit": "編集",
                "delete": "消去"
        },
        "modal": {
            "newCoupon": "新しいクーポンを追加します",
                "editCoupon": "クーポン情報を変更します",
                "confirmAdd": "追加することを確認します",
                "confirmEdit": "変更を確認します",
                "emptyNotAllow": "このアイテムが必要です",
                "delMention": "クーポンは、エントリが削除された直後に無効になり、この操作を引き出すことはできません。",
                "cancel": "キャンセル",
                "name": {
                "title": "クーポン名",
                    "placeholder": "クーポン名を入力してください"
            },
            "code": {
                "title": "クーポン",
                    "placeholder": "クーポンコードをカスタマイズします（空白のままにして、できるだけ早く生成します）"
            },
            "percentOff": {
                "title": "情報を提供する（パーセンテージ）",
                    "placeholder": "割引率を入力します（-12％オフなど）"
            },
            "activeRange": { "title": "利用可能なクーポンの免除範囲" },
            "capacity": {
                "title": "使用されるクーポンの最大数",
                    "placeholder": "最大使用制限の制限（空の場合は制限なし）"
            },
            "residue": {
                "title": "使用される残りのクーポンの数",
                    "placeholder": "クーポンが使用されている回数を設定する"
            },
            "perUserLimit": {
                "title": "各ユーザーがクーポンを使用できる回数",
                    "placeholder": "各ユーザーが使用できる回数を制限します（スペースに制限なし）"
            },
            "planLimit": {
                "title": "サブスクリプションプランを指定します",
                    "placeholder": "指定されたサブスクリプションプランの制限オファーを使用する（スペースに制限なし）"
            }
        }
    },
    "orderMgr": {
        "title": "注文管理",
            "description": "ここでは、すべてのサブスクリプションプランの注文を表示したり、さまざまなユーザーの注文をフィルタリングしたり、ユーザーの注文を手動で処理するなどを表示できます。",
            "table": {
            "id": "＃",
                "orderId": "注文番号",
                "email": "ユーザーメールボックス",
                "status": {
                "title": "タイプ",
                    "t1": "新しく購入",
                    "t2": "継続料金",
                    "t3": "編集"
            },
            "planName": "計画名",
                "period": "サイクル",
                "group": "権利グループ",
                "amount": "実際の支払い額",
                "price": "元の価格",
                "isSuccess": {
                "title": "ステータスを注文します",
                    "cancelOrder": "注文をキャンセルします",
                    "passOrder": "注文によって"
            },
            "createdAt": "作成時間を注文します",
                "action": { "title": "動作します", "showDetail": "詳細を表示します" }
        },
        "search": "お問い合わせ注文",
            "resetSearch": "クエリをリセットします",
            "failureReason": "失敗の原因",
            "couponId": "クーポンID",
            "couponName": "クーポン名",
            "noEntry": "それなし",
            "orderDetail": "詳細を注文します",
            "searchModal": {
            "email": {
                "title": "ユーザーメールボックス",
                    "placeholder": "ユーザーメール（ファジー検索）を入力します"
            },
            "sort": {
                "title": "ソートアルゴリズム",
                    "placeholder": "ソートアルゴリズムを選択します",
                    "ASC": "昇順",
                    "DESC": "降順"
            }
        },
        "tradeWaiting": "支払われていません",
            "tradeFailure": "トランザクションは失敗しました",
            "tradeSuccess": "成功"
    },
    "userMgr": {
        "userManager": "ユーザー管理",
            "description": "従業員や管理者など、すべてのユーザーを管理したり、管理権を付与またはキャンセルしたり、ユーザーバランスを設定したり、パスワードをリセットしたり、新しいユーザーを手動で追加したりすることができます。",
            "enterEmail": "メールを入力してください",
            "enterValidEmail": "正しい電子メール形式を入力してください",
            "enterPassword": "パスワードを入力してください",
            "passwordMinLength": "パスワードの長さは6桁未満にすることはできません",
            "passwordMaxLength": "パスワードの長さは20桁を超えることはできません",
            "passwordStrength": "パスワードには、文字、数字、特殊文字が含まれている必要があります",
            "validationSuccess": "検証が合格しました",
            "validationFailed": "フォーム検証が失敗しました。入力アイテムを確認してください",
            "editProfile": "情報を編集します",
            "banUser": "ログインを無効にします",
            "unbanUser": "ユーザーのブロックを解除します",
            "noRecord": "録音されていません",
            "normal": "普通",
            "banned": "禁止",
            "deleted": "ログアウト",
            "nullContent": "ヌル",
            "balance": "額",
            "orderCount": "数量を注文します",
            "planCount": "予定",
            "actions": "動作します",
            "updateSuccess": "正常に更新します",
            "addUserSuccess": "新しいユーザーを正常に追加します",
            "unknownError": "不明なエラー",
            "email": "メール",
            "registerDate": "登録日",
            "isAdmin": "管理者",
            "isStaff": "スタッフ",
            "accountStatus": "アカウントステータス",
            "inviteCode": "招待コード",
            "query": "クエリ",
            "reset": "リセット",
            "addNewUser": "追加されたユーザー",
            "searchUser": "お問い合わせユーザー",
            "cancel": "キャンセル",
            "submit": "提出する",
            "userEmail": "ユーザーメールボックス",
            "inputUserEmail": "ユーザーメール（ファジー検索）を入力します",
            "inputEmail": "電子メールを入力してください",
            "password": "パスワード",
            "inputPassword": "パスワードを入力します",
            "editUser": "ユーザーを編集します",
            "no": "いいえ",
            "yes": "はい",
            "mention": {
            "title": "アカウントを無効にしたいですか？",
                "content": "ユーザーがブロックされている場合、ユーザーは{appname}にログインできず、関連するすべてのリソースが利用できなくなります。"
        }
    },
    "activation": {
        "activateLog": "レコードをアクティブにします",
            "description": "販売されているすべてのキーの特定のアクティベーションステータスを表示したり、クライアントの識別コード、アクティベーション時間などを表示できます。",
            "click2getKey": "クリックしてキーコンテンツを取得します",
            "createdAt": "作成時間",
            "turn2keyPage": "キーに目を向けます",
            "showKey": "キーを表示します",
            "email": "メール",
            "key": "鍵",
            "search": "検索",
            "filter": "スクリーニング",
            "filterAll": "全て",
            "filterActive": "効果的です",
            "sortAlgorithm": "ソートアルゴリズム",
            "sortASC": "昇順",
            "sortDESC": "降順"
    },
    "keysMgr": {
        "keyMgr": "キー管理",
            "description": "生成されたすべてのキーコンテンツとその使用状況、有効期間などを確認できます。キーを無効にすることもできます。",
            "enableKey": "キーを有効にします",
            "disableKey": "キーを無効にします",
            "table": {
            "email": "電子メールアドレス",
                "key": "キーコンテンツ",
                "isValid": "有効かどうか",
                "isUsed": "それを使用するかどうか",
                "isExpired": "有効期限が切れるかどうか",
                "active": "アクティブ",
                "inactive": "期限切れ",
                "yes": "効率的",
                "no": "期限切れ",
                "ok": "普通",
                "blocked": "ステータスを無効にします",
                "unUsed": "使用されていません",
                "used": "使用済み",
                "showDetail": "詳細を表示します",
                "blockKey": "キーを無効にします",
                "unblockKey": "禁止されていません"
        },
        "searchModal": {
            "searchMethod": "お問い合わせ方法",
                "byId": "IDによるクエリ",
                "byContent": "キーを介した問い合わせ（ファジー）",
                "keyId": "キーID",
                "idPlaceholder": "ここにキーのIDを入力してください",
                "keyContent": "キーワード",
                "contentPlaceholder": "ここにキーの一部を入力してください"
        },
        "mention": {
            "blockOk": "キーを無効にして正常にid：{id}",
                "title": "キーを無効にしたいですか？",
                "content": "ユーザーエクスペリエンスを確保するために、キーのコンテンツを再度確認してください。"
        },
        "detailModal": {
            "title": "重要な詳細",
                "userId": "ユーザーID",
                "planName": "サブスクリプションプラン名",
                "expiredAt": "有効期限",
                "keyGeneratedAt": "キー生成日",
                "clientId": "クライアントID"
        }
    },
    "noticeMgr": {
        "title": "発表管理",
            "description": "ここで発表は、ユーザーのホームページにあるホイールバローの画像に表示されます。",
            "addNotice": "発表を追加します",
            "table": {
            "id": "＃",
                "show": "表示するかどうか",
                "title": "タイトル",
                "createdAt": "作成時間",
                "action": { "title": "動作します", "edit": "編集", "delete": "消去" }
        },
        "mention": {
            "title": "あなたはそれを削除したいことを知っていますか？",
                "content": "発表はすぐに削除され、このアクションを撤回することはできません。"
        },
        "modal": {
            "addNew": "新しい発表",
                "title": {
                "title": "発表タイトル",
                    "placeholder": "ホイールキャストの大きなタイトルとして表示します"
            },
            "content": {
                "title": "アナウンスコンテンツ",
                    "placeholder": "発表を書くことの主な内容"
            },
            "tag": {
                "title": "アナウンスタグ",
                    "placeholder": "アナウンスのタグを入力してください"
            },
            "img": {
                "title": "背景画像URL",
                    "placeholder": "入力しない場合は、デフォルトの背景を使用してください"
            }
        }
    },
    "adminTicket": {
        "ticketMgr": "ワークオーダー管理",
            "description": "ここでは、すべてのユーザーが提出された作業注文を表示できます。",
            "pendingTicket": "処理される作業注文",
            "finishedTicket": "完了した作業注文",
            "type": {
            "pendingDescription": "ライブオーダー、これは最初に処理する必要があります。",
                "finishedDescription": "完了した作業注文、ここで表示できます。"
        },
        "chooseOneNecessary": "少なくとも1つのアイテムを選択する必要があります",
            "mention": {
            "title": "チケットを閉じたいですか？",
                "content": "注文は、閉じた後、完成した注文にアーカイブされますが、注文に返信することはできません。"
        }
    }
},
    "userLogin": {
    "loginToContinue": "ログインして続行します",
        "email": "電子メールアドレス",
        "password": "パスワード",
        "haveNoAccount": "まだアカウントを持っていませんか？",
        "login": "ログイン",
        "reg": "今すぐ登録してください",
        "otherMethods": "または、他の方法を使用して続行します",
        "github": "Githubアカウントを続行します",
        "microsoft": "Microsoftアカウントを続行します",
        "apple": "Appleアカウントを続行します",
        "google": "Googleアカウントを続行します",
        "backHomePage": "ホームページに戻ります",
        "form": {
        "emailRequire": "メールアドレスが必要です",
            "passwordRequire": "まだパスワードを入力していません"
    },
    "authPass": "検証が合格しました",
        "loginFailure": "ログインに失敗しました",
        "checkForm": "フォームを確認してください",
        "if2FaEnabledHint": "2段階の検証を有効にした場合（不要）",
        "reqErr": "エラーがある場合は、後でもう一度試してください",
        "accountLocked": "アカウントは登録または禁止されている可能性があり、これがまだエラーだと思われる場合は、テクニカルサポートにお問い合わせください。",
        "tokenNotExist": "トークンを提供してください"
},
    "userRegister": {
    "backHomePage": "ホームページに戻ります",
        "newAccount": "新しいアカウントを準備します",
        "email": "電子メールアドレス",
        "verifyCode": "電子メール検証コード",
        "invalidEmailFormat": "電子メール形式は違法です",
        "sendVerifyCode": "電子メール検証コードを送信します",
        "sendSuccess": "メールは正常に送信されました。メールを確認してください。",
        "pwd": "パスワード",
        "pwdAgain": "パスワードを認証する",
        "inviteCode": "招待コード（オプション）",
        "agreement": "私は読んで同意しました",
        "terminalUserAgreement": "ユーザー用語",
        "reg": "登録する",
        "infoIncomplete": "不完全な情報",
        "pwdIncorrect": "一貫性のないパスワード",
        "regSuccess": "登録は成功し、ログインに戻ります",
        "regFailure": "登録に失敗しました",
        "success": "成功",
        "failure": "失敗",
        "unknownErr": "不明なエラー",
        "verifyCodeErr": "検証コードエラー",
        "verifyCodeExpireErr": "検証コードが正しくないか、もう一度試してみてください。",
        "thisMailAlreadyExist": "メールは登録されています",
        "pageConfigFetchFailure": "構成の取得が失敗した場合は、更新して再試行してください",
        "stopRegisterTitle": "登録は停止しました",
        "stopRegisterHint": "申し訳ありませんが、登録機能は停止されています。 必要な場合は、後でもう一度お試しするか、サポートチームにお問い合わせください。 ご理解とご支援ありがとうございます。",
        "passwordComplexRequirePart1": "*パスワードをコンプライアンスする必要があります",
        "passwordComplexRequirePart2": "複雑さの要件",
        "passwordComplexHint1": "1. 3種類以上の大きな文字、小さな文字、数字、特別なシンボルが必要です。",
        "passwordComplexHint2": "2。長さは10桁以上です",
        "form": {
        "checkForm": "フォームのコンテンツを確認してください",
            "emailFormatErr": "メールアドレス形式は正しくありません",
            "gmailLimitErr": "Gmail複数の名前は許可されていません",
            "gmailDotNotAllowed": "「。」",
            "gmailPartLowerForced": "Gmailアドレスのローカル部分はすべての読者でなければなりません",
            "googlemailNotAllowed": "GoogleMailアドレスはサポートされていません",
            "verifyCodeRequire": "確認コードを入力してください",
            "verifyCodeFormatErr": "検証コード形式は正しくありません",
            "passwordRequire": "パスワードを入力してください",
            "passwordLengthRequire": "パスワードの長さは10桁以上でなければなりません",
            "passwordComplexRequire": "パスワードには、少なくとも3種類の大きな文字、小さな文字、数字、特別なシンボルが含まれている必要があります。",
            "passwordAgainNotMatch": "パスワードは2回の不一致を入力しました",
            "passwordAgainRequire": "確認パスワードを入力してください",
            "inviteCodeRequire": "招待コードが必要です"
    },
    "hCaptcha": {
        "passed": "ヒューマンマシン検証が合格しました",
            "expired": "有効期限が切れたらもう一度やり直してください",
            "challengeExpired": "確認は時間を超えています",
            "err": "その他のエラー"
    },
    "allRightsReserved": "無断転載を禁じます",
        "securityAndLaws": "このウェブサイトは、Hcaptchaによって保護および検証されており、現地の法律の対象となっています。"
},
    "userSummary": {
    "title": "楽器",
        "myPlan": "私のサブスクリプション",
        "shortcut": "ショートカット",
        "timeLeft": "サブスクリプションは有効であり、{MSG}で期限切れになります。",
        "toPurchase": "サブスクリプションを購入します",
        "tutorial": {
        "title": "チュートリアルをご覧ください",
            "content": "{name}の使用方法を学ぶ"
    },
    "checkKey": {
        "title": "キーを表示します",
            "content": "使用するためにクライアントにキーをすばやく紹介します"
    },
    "renewPlan": {
        "title": "サブスクリプションを続けます",
            "content": "現在のサブスクリプションに対して継続的な支払いを行います"
    },
    "appDown": {
        "title": "アプリケーションのダウンロード",
            "content": "さまざまなプラットフォームからアプリケーションをダウンロードして、より良い身体的体験をする"
    },
    "support": {
        "title": "問題に遭遇しました",
            "content": "問題が発生した場合は、作業順序を通じて私たちの人々と通信できます。"
    },
    "haveTicket": "処理される{count}順序があります",
        "toCheckTicket": "行ってチェックしてください",
        "showAllKeys": "すべてのキーを表示します"
},
    "userDocument": {
    "title": "ドキュメントを使用します",
        "description": "ここでは、ウェブサイトの使用方法、予防策、操作手順などを含むがこれらに限定されないさまざまなドキュメントを確認できます。記事にエラーが見つかった場合は、作業注文を送信してください。",
        "searchPlaceholder": "検索するコンテンツを入力してください（ファジー検索）",
        "searchBtn": "検索",
        "noContentTitle": "結果はありません",
        "noContentTitleHint": "検索結果や言語に一致するドキュメントはありません。キーワードに変更してみてください。"
},
    "newPurchase": {
    "title": "サブスクリプションを購入します",
        "description": "ここで最適なサブスクリプションプランを選択できます。",
        "headerPlaceholder": "あなたに最適な計画を選択してください",
        "purchaseBtn": "注文",
        "noLeft": "残りの量が不十分です",
        "monthPay": "毎月の支払い",
        "moreMethod": "その他のオプション"
},
    "newSettlement": {
    "err": "間違い",
        "monthPay": "毎月の支払い",
        "quarterPay": "四半期支払い",
        "halfYearPay": "半年の支払い",
        "yearPay": "年間支払い",
        "payCycle": "支払いサイクル",
        "couponPlaceholder": "クーポンをお持ちですか？",
        "verifyBtn": "検証",
        "orderTotalTitle": "合計注文額",
        "total": "合計",
        "submitOrder": "注文を送信します",
        "coupon": "クーポン",
        "notify": {
        "passTitle": "検証が合格しました",
            "couponVerified": "有効なクーポン",
            "couponInvalid": "クーポンは無効です",
            "couponIsNull": "入力されたクーポンコードは空にすることはできません"
    }
},
    "userProfile": {
    "title": "パーソナルセンター",
        "myWallet": "私のお金のバッグ",
        "walletSub": "アカウント残高（消費にのみ使用）",
        "alertPwd": "キーを変更します",
        "oldPwd": "古い鍵",
        "oldPwdSub": "古いパスワードを入力してください",
        "newPwd": "新しいキー",
        "newPwdSub": "新しいキーを入力してください",
        "newPwdAgain": "キーを確認します",
        "newPwdAgainSub": "もう一度新しいキーを入力してください",
        "saveBtn": "保存",
        "notify": "通知します",
        "enableNotify": "有効期限を有効にします",
        "deleteAccount": "アカウントを登録します",
        "deleteAccountSub": "アカウントは削除されたとタグ付けされます。サービスを再利用する必要がある場合は、再登録してください",
        "deleteBtn": "アカウントを登録してください",
        "oldPwdVerifiedFailure": "古いパスワード検証に失敗しました",
        "alertFailure": "キーの変更に失敗しました",
        "alertSuccess": "変更は成功しました",
        "alertSuccessSub": "新しいパスワードでログインしてください",
        "errorPwdFormat": "パスワード形式エラー",
        "pwdNotMatch": "パスワード入力は一貫性がありません",
        "oldPwdNotNull": "古いパスワードを空にすることはできません",
        "toTopUp": "リチャージに移動します",
        "deleteMyTitle": "警告",
        "deleteMyContent": "アカウントを削除しますか？ サービスを使用する必要がある場合は、再登録してください。",
        "deleteMyPositiveText": "削除を確認します",
        "deleteMyNegativeText": "キャンセル",
        "deletedSuccessMsg": "アカウントは削除されており、期間が経ちます。",
        "deleteErrOccur": "遭遇したエラー",
        "faAuth": "2段階検証2FA",
        "faAuthHint": "2段階の検証は、アカウントにログインするための保護レベルを追加するセキュリティメカニズムです。 パスワードを入力した後、ユーザーは、電話に送信された検証コードを入力したり、ID検証アプリケーションを使用して動的コードを生成したり、指紋などの生物学的特性を通じて確認したりするなど、身元確認の2番目のステップを完了する必要があります。",
        "faAuthStatus": "2段階検証ステータス",
        "faEnabled": "再起動",
        "faNotEnabled": "アクティブ化されていません",
        "setup2Fa": "2段階の検証を設定します",
        "disable2Fa": "2段階の検証をキャンセルします",
        "unknownErr": "不明なエラー",
        "disable2FaCancelled": "キャンセル",
        "closed2FaHint": "2段階の検証は閉じられています。必要に応じてもう一度追加してください。",
        "setup2FaModal": {
        "followStep": "プロンプトに従って検証装置に追加します",
            "step1Title": "2FA検証を有効にするには、以下の手順に従ってください",
            "step1Context1": "1.モバイルデバイスにユニバーサル検証デバイスが必要です",
            "step1Context2": "2。検証デバイスのスキャンボタンをクリックして、こちらのQRコードをスキャンします",
            "step1Context3": "3.このQRコードには、検証情報と一意のキーが含まれています。",
            "step2Context1": "検証デバイスを正常に使用できるようにするには、テストを実施する必要があります。",
            "test2Fa": "テスト",
            "cancel": "キャンセル"
    },
    "deleteMyAccountModal": {
        "title": "アカウントを登録します",
            "contentLine1": "会計は不可逆的な操作です。 削除されていることを確認すると、アカウントへのアクセスが永久にアクセスできます。つまり、個人情報、履歴、収集コンテンツ、購入レコードなどを含むがこれらに限定されない、このアカウントに関連するすべてのデータを再度ログインすることはできません。",
            "contentLine2": "未完成の注文、参加中のアクティビティ、サブスクリプションサービスなど、プラットフォームで継続的なビジネスがある場合、これらはアカウントの削除で終了またはキャンセルされ、対応する損失をもたらす可能性があります。 同時に、このプラットフォームを介したお客様と他のユーザーの連絡先と相互作用情報は存在しなくなります。",
            "contentLine3": "あなたの決定をもう一度確認してください。 それでもアカウントを削除する場合は、[削除]ボタンをクリックします。",
            "inputHint1": "入力 \"",
            "inputHint2": "\" 続く。",
            "confirmDelete": "削除を確認します"
    },
    "failure": "エラーが発生しました",
        "notLatestHint": "個人情報の更新が失敗し、現在のデータが最新のものではない可能性があります。",
        "resetPassword": {
        "previousPasswordRequire": "古いパスワードを入力してください",
            "previousPasswordVerifiedFailure": "古いパスワード検証に失敗しました",
            "passwordRequire": "パスワードを入力してください",
            "passwordConflict": "新しいパスワードは以前のパスワードと同じではありません",
            "passwordLengthRequire": "パスワードの長さは10桁以上でなければなりません",
            "passwordComplexRequire": "パスワードには、少なくとも3種類の大きな文字、小さな文字、数字、特別なシンボルが含まれている必要があります。",
            "passwordAgainNotMatch": "パスワードは2回の不一致を入力しました",
            "passwordAgainRequire": "確認パスワードを入力してください",
            "updatePasswordFailure": "パスワードを更新するときにエラーが発生しました"
    },
    "secondary": {
        "title": "基本情報",
            "card": {
            "avatar": "ユーザーヘッド画像",
                "name": "ユーザー名",
                "lastLoginAt": "最後のログイン時間",
                "accountStatus": "アカウントステータス"
        },
        "modify": {
            "uploadIconHint": "アップロード",
                "alterAvatar": "ヘッド画像を追加/変更します",
                "alterShallow": "*ヘッド画像またはユーザー名をクリックして追加または変更する（設定後にクリアは許可されていません）",
                "alterName": "ユーザー名を変更します",
                "input": {
                "label": "ユーザー名",
                    "placeholder": "ユーザー名を入力します",
                    "spaceIsNotAllowed": "ユーザー名の確認は渡されません",
                    "require": {
                    "p1": "純粋な数字は許可されておらず、数字が開かれています",
                        "p2": "スペースは許可されていません",
                        "p3": "3より長い"
                }
            }
        },
        "mention": {
            "alterNameSuccess": "ユーザー名の変更が成功しました",
                "alterNameErr": "ユーザー名が失敗した場合は、後でもう一度やり直してください",
                "newNameIsNotValid": "違法なユーザー名",
                "click2SetName": "クリックしてユーザー名を設定します",
                "fetchAvatarErr": "ヘッド画像データが失敗した場合は、後でもう一度やり直してください",
                "alterAvatarErr": "ヘッド画像が失敗した場合は、後でもう一度やり直してください",
                "success": "成功",
                "alterAvatarSuccess": "ヘッド画像は正常に変更されました。",
                "uploadImageHint": "画像を新しい見出しとしてアップロードできます",
                "imageRequire": {
                "title": "知らせ",
                    "p1": "*.jpg（jpeg）、 *.png、 *.webpなどの主流形式をアップロードできます。",
                    "p2": "画像の幅比は1：1（正方形）で、異なる比率の場合、中央の正方形にカットされ、過剰が削除されます。",
                    "p3": "アップロードする画像のサイズは160pxに設定されます。"
            },
            "click2Upload": "ファイルをクリックまたはドラッグしてエリアにドラッグしてアップロードします",
                "uploadWarn": "*銀行カード、クレジットカード、パスワード、セキュリティコードなど、機密データをアップロードしないでください。"
        }
    }
},
    "userKeys": {
    "myKeys": "私の鍵",
        "description": "すべてのキーについて、イベントステータス、有効期限などを表示できます。",
        "noKeys": "まだ有効な購入記録はありません",
        "keyDetail": "重要な詳細",
        "keyId": "キーID",
        "orderId": "リンク注文ID",
        "clientId": "クライアントIDをアクティブ化します",
        "active": "有効期間中",
        "inActive": "期限切れ",
        "valid": "重要な状態は正常です",
        "invalid": "キーは無効になっています",
        "isUsed": "使用するためにアクティブ化されています",
        "noUsed": "まだ使用されていません",
        "releaseData": "キー生成日",
        "expirationData": "有効期限",
        "none": "それなし",
        "authorizationFor": "主要な承認",
        "hoverClickMention": "クリックしてクリップボードにコピーします",
        "copiedSuccessMessage": "キーコピーは正常にコピーされました。",
        "copyFailure": "失敗のコピー",
        "hoverCopiedSuccessMention": "正常にコピーされました"
},
    "userOrders": {
    "myOrders": "私の注文",
        "description": "すべての注文がここに表示されます。未払いの注文がある場合は、支払いを続行するか、注文をキャンセルし、注文が完了した後に注文の詳細を表示できます。",
        "orderId": "＃",
        "planName": "サブスクリプション名",
        "planCycle": "サブスクリプションサイクル",
        "orderPrice": "注文金額",
        "orderStatus": "ステータスを注文します",
        "createdAt": "作成時間",
        "operate": "動作します",
        "showDetail": "詳細を注文します",
        "cancelOrder": "注文をキャンセルします",
        "canceled": "キャンセル",
        "period": {
        "monthPay": "毎月の支払い",
            "quarterPay": "四半期支払い",
            "halfYearPay": "半年の支払い",
            "yearPay": "年間支払い"
    },
    "orderStatusTags": {
        "success": "成功",
            "cancelled": "失敗",
            "notPay": "支払われていません"
    },
    "orderCancelled": "注文キャンセル",
        "unknownErr": "不明なエラー"
},
    "userTopUp": {
    "topUp": "アカウント充電",
        "description": "ここでアカウントを充電し、カスタムリチャージ額をサポートすることもできます。また、以下に表示されている割引情報があるかどうかに注意して、上記の支払い方法を使用して割引を受けることもできます。",
        "chooseTopUpAmount": "充電額を選択します",
        "quickSelect": "すぐに選択します",
        "customAmount": "カスタム量",
        "maxAmount": "最大額：10,000,000",
        "amountInputPlaceHolder": "充電される金額を入力します",
        "otherAmount": "その他の金額",
        "payMethod": "支払方法",
        "wechat": "wechatの支払い",
        "alipay": "支払い",
        "apple": "Apple Pay",
        "yourAmount": "あなたの金額",
        "discount": "申し出",
        "accountBalance": "アカウント残高",
        "balanceResult": "バランスの量",
        "commitTopUp": "提出する",
        "payMethodNotAllow": "支払い方法は利用できません",
        "topUpIssueOccur": "充電に問題がありますか？",
        "payIssueOccur": "支払いに問題がありますか？",
        "chatWithUs": "カスタマーサービスに連絡してください",
        "pay": "支払う",
        "qrCodeScannedSuccess": "QRコードスキャンは正常にスキャンします",
        "orClickToApp": "または、バーをクリックしてアプリに移動して続行します",
        "topUpSuccess": "正常に充電します",
        "thankU": "ご支援ありがとうございます"
},
    "userConfirmOrder": {
    "switchPlan": "サブスクライブに切り替えます",
        "cancelOrder": "注文をキャンセルします",
        "yourPrice": "あなたの価格",
        "couponOffset": "クーポン割引",
        "price": "価格",
        "submit": "提出する",
        "monthPay": "毎月の支払い",
        "quarterPay": "四半期支払い",
        "halfYearPay": "半年の支払い",
        "yearPay": "年間支払い",
        "goodInfo": "製品情報",
        "cycleOrType": "サイクル/タイプ",
        "orderNumber": "注文番号",
        "createdAt": "作成日",
        "orderExpired": "注文が超えています",
        "paySuccessfully": "支払いは成功しました、あなたのサポートに感謝します。",
        "balanceNotEnough": "バランスが不十分な場合は、最初に充電してください。",
        "orderErrorOccur": "注文中にエラーが発生しました",
        "orderCancelled": "注文キャンセル"
},
    "paymentResultParts": {
    "goodInfoView": { "goodInfo": "製品情報" },
    "orderInfoView": { "orderInfo": "情報を注文します" }
},
    "orderPartUniversal": {
    "period": {
        "monthPay": "毎月の支払い",
            "quarterPay": "四半期支払い",
            "halfYearPay": "半年の支払い",
            "yearPay": "年間支払い"
    },
    "orderDataHex": {
        "goodInfo": "製品情報",
            "orderInfo": "情報を注文します",
            "cycleOrType": "サイクル/タイプ",
            "orderNumber": "注文番号",
            "createdAt": "作成日",
            "amount": "支払い額",
            "paidAt": "支払い時間"
    }
},
    "orderDetail": {
    "finished": "完了しました",
        "finishedAndSuccessDescription": "注文は正常に支払われ、開設されました",
        "useManual": "チュートリアルをご覧ください",
        "payPending": "まだ支払われていません",
        "pendingDescription": "注文は定期的に保持され、下のボタンをクリックして支払いを続けることができます。",
        "toPay": "支払いに行きます",
        "outDate": "有効期限が切れています",
        "outDateDescription": "注文をキャンセルするか、指定された時間内に支払いを完了しなかったため、注文はキャンセルされ、サブスクリプションを再選択できます。",
        "chooseNewPlan": "新しいサブスクリプションプランを選択してください"
},
    "userInvite": {
    "myInvite": "私の招待",
        "unit": "人数",
        "inviteCodeMgr": "あなたの招待コード",
        "generateInviteCode": "ランダムな招待コードを生成します",
        "faCodeManage": "招待コード管理",
        "email": "メール",
        "createdAt": "登録時間",
        "createFaCodeFailure": "作成に失敗しました",
        "linkCopiedSuccess": "リンクは正常にコピーされました",
        "generateFaCode": "招待コードを生成します",
        "flushFaCode": "招待コードを更新します",
        "faCode": "招待コード",
        "noFaCode": "招待コードはまだありません。",
        "faLink": "招待状リンク",
        "generateFaCodePlease": "招待コードを作成してください",
        "usersMyInvited": "私が招待するユーザー",
        "generateCodeConfirm": "生成/更新を確認します",
        "generateCodeHint": "作成後は招待状コードを閉じることはできないことに注意してください。",
        "positiveClick": "確認する",
        "negativeClick": "キャンセル"
},
    "userTickets": {
    "description": "使用に関する問題が発生した場合は、こちらのテクニカルサポートとカスタマーサービスを送信して、以下のフォームで特定の返信時間をマークしてください。",
        "ticketId": "＃",
        "ticketSubject": "作業注文のテーマ",
        "ticketUrgency": "ワークオーダーグレード",
        "ticketContent": "ワークオーダーコンテンツ",
        "ticketUrgencyLevel": { "low": "低い", "med": "真ん中", "high": "高い" },
    "ticketStatus": "ワークオーダーステータス",
        "ticketCreatedAt": "作成時間",
        "lastResponse": "最後の返信",
        "operate": "動作します",
        "checkTicket": "作業注文を表示します",
        "closeTicket": "作業指示を閉じます",
        "userTickets": "歴史は機能します",
        "addTicket": "新しい建設注文",
        "ticketActive": "閉じていない",
        "ticketInActive": "閉じた",
        "form": {
        "ticketSubjectDescription": "作業注文のトピックを入力してください",
            "ticketUrgencyDescription": "作業注文の緊急性を選択してください",
            "ticketBody": "できるだけ包括的に解決したい問題を入力してください。",
            "ticketNotFinished": "完全な作業注文に関する情報をお試しください"
    },
    "checkForm": "フォームが完了しているかどうかを確認してください",
        "cancel": "キャンセル",
        "submit": "提出する",
        "commitNewTicketSuccess": "新しい作業注文を正常に送信します",
        "commitNewTicketFailure": "新しい作業注文エラーを送信します",
        "noReply": "まだ返信はありません",
        "noTickets": "まだ作業指示書を提出していません",
        "ticketCloseSuccess": "作業注文は正常に終了しました",
        "ticketCloseFailure": "作業注文の閉鎖は失敗しました",
        "chatDialog": {
        "input": {
            "finished": "作業順序が終了しました",
                "inputHere": "送信するメッセージを入力します"
        },
        "send": "送信",
            "missingToken": "トークンが欠落しているため、WebSocket接続は作成できません。",
            "msgEmptyNotAllowed": "メッセージのコンテンツを空にすることはできません",
            "accessNotPermit": "違法訪問",
            "sendSuccess": "メッセージを正常に送信します"
    }
},
    "userActivation": {
    "activateLog": "レコードをアクティブにします",
        "description": "ここでは、アクティベーションレコードを表示して、未定のデバイスを設定するか、各キーとアクティベーションレコードのメモ情報を設定できます。",
        "id": "＃",
        "orderId": "注文番号",
        "orderStatus": "注文",
        "createdAt": "作成時間",
        "operate": "動作します",
        "userId": "ユーザーID",
        "email": "メール",
        "keyId": "キーID",
        "isBind": "アクティブ化するかどうか",
        "active": "効率的",
        "inactive": "効果がない",
        "requestAt": "リクエスト時間",
        "clientVersion": "クライアント側",
        "osType": "オペレーティング·システム",
        "remark": "注記",
        "noRemark": "ノートなし",
        "showDetail": "詳細を表示します",
        "actions": "動作します",
        "details": "詳細",
        "keyContent": "キーコンテンツ",
        "keyGeneratedAt": "キー生成時間",
        "activateRequestAt": "リクエスト時間をアクティブにします",
        "useIssueOccur": "それを使用するのに問題がありますか？",
        "chatWithUs": "お問い合わせ",
        "cancelBind": "法令をキャンセルします",
        "alterRemark": "メモを変更します",
        "commitRemark": "提出する",
        "updateSuccess": "正常に更新します",
        "setRemark": "ここでメモ情報をセットアップします"
},
    "userAppDownload": {
    "title": "アプリのダウンロード",
        "description": "",
        "common": {
        "title": "アプリをダウンロードしてください",
            "shallow2": "さまざまなクライアント向けにアプリケーションを入手してください",
            "shallow": "アプリケーションを使用して、サービスに簡単にアクセスでき、毎回ブラウザの複雑な操作を排除できます。",
            "noDownload": "ダウンロードはまだありません。後で再試行してください。"
    },
    "suffix": {
        "p1": "*MacOSプラットフォームのアプリケーションについては、Macos14以上を使用して、Appleチップ（M1以降）と協力してください。",
            "p2": "このアプリによって表示される情報と使用は、アプリが許可されているかどうかも、会社のIT仕様に依存する必要があります。"
    },
    "card": {
        "common": { "welcome": "いらっしゃいませ" },
        "mobile": {
            "designFor": "モバイルターミナルのデザイン",
                "designShallow": "ここでモバイルアプリを入手できます",
                "iosDownloadShallow": "iOSクライアントをダウンロードします",
                "androidDownloadShallow": "Androidクライアントをダウンロードします"
        },
        "desktop": {
            "designFor": "デスクトップ用のデザイン",
                "designShallow": "ここでデスクトップモバイルアプリケーションを入手できます",
                "windowsDownloadShallow": "Windowsクライアントをダウンロードします",
                "osxDownloadShallow": "MacOSクライアントをダウンロードします",
                "linuxDownloadShallow": "Linuxクライアントをダウンロードします"
        }
    },
    "downloadDisabledHint": "申し訳ありませんが、アプリのダウンロードは、管理者が必要とする場合は、テクノロジーにお問い合わせください。",
        "windows": {
        "title": "Windows NT",
            "shallow": "このクライアントは、NTカーネルを備えたWindowsオペレーティングシステムに適しています。互換性のサポートについては、ドキュメントページを確認してください。"
    },
    "osx": {
        "title": "MacOS（OSX）",
            "shallow": "このクライアントは、ダーウィンカーネルのMACOS（OSX）オペレーティングシステムに適しています。互換性サポートについては、ドキュメントページを確認してください。"
    },
    "linux": {
        "title": "Linux",
            "shallow": "このクライアントは、Linuxカーネルのさまざまな分布に適しています。"
    },
    "android": {
        "title": "アンドロイド",
            "shallow": "このクライアントは、Google Androidオペレーティングシステムを備えたモバイルデバイスに適しています。互換性のサポートについては、ドキュメントページを確認してください。"
    },
    "ios": {
        "title": "iOS",
            "shallow": "このクライアントは、Apple iOSオペレーティングシステムを備えたモバイルデバイスに適しています。互換性のサポートについては、ドキュメントページを確認してください。"
    }
},
    "welcome": {
    "A": {
        "aboutUs": "私たちについて",
            "pricing": "注文価格",
            "login": "ログイン",
            "register": "アカウントを登録します",
            "welcomeTo": "いらっしゃいませ",
            "welcomeToSub": "郡の長いトンネルを通過するのは雪の国です。 夜空の下では、地球は白く、列車は信号ステーションの前で止まりました。 「ここで、川谷ヤスナリは非常にケチな簡潔なテキストで雪の国をキックオフしました。",
            "startingUse": "始めましょう",
            "whyUs": "なぜ私たちを選ぶのか",
            "whyUsSub": "「県の長いトンネルを通過すると、雪の国です。夜空の下では、地球は白く、信号ステーションの前で列車が止まります。",
            "browseSafe": "安全に閲覧します",
            "browseSafeSub": "優れたファイアウォールフィルターシステムは、オンラインの魚や悪意のあるウェブサイトを効果的に防ぐことができます",
            "encrypt": "エンドツーエンド暗号化",
            "encryptSub": "二重ウェイSSLとエンドツーエンドの暗号化は、プライバシーとセキュリティを保護し、サーバーでさえあなたの情報を読み取ることができません",
            "mgr": "効率的な管理",
            "mgrSub": "1つのユーザーインターフェイスは、完全かつリッチな管理機能を備えたすべてのキーを管理し、射精の問題を購読することを心配する必要はありません",
            "fast": "便利で速い",
            "fastSub": "WebAppsまたは埋め込まれたサードパーティソフトウェアに完全なAPIファイルを提供する",
            "fastLink": "クイックリンク",
            "subscribeUs": "私たちに従ってください",
            "filingsCode": "ファイル番号{code}"
    }
},
    "pagination": {
    "perPage10": "10データ/ページ",
        "perPage20": "20データ/ページ",
        "perPage50": "50データ/ページ",
        "perPage100": "100データ/ページ"
},
    "modal": { "cancel": "キャンセル", "confirm": "確認する" },
    "week": {
    "monday": "月曜日",
        "tuesday": "火曜日",
        "wednesday": "水曜日",
        "thursday": "木曜日",
        "friday": "金曜日",
        "saturday": "土曜日",
        "sunday": "日曜日"
},
    "period": {
    "month": "毎月の支払い",
        "quarter": "四半期支払い",
        "halfYear": "半年の支払い",
        "year": "年間支払い"
},
    "operate": {
    "cancel": "キャンセル",
        "confirm": "確認する",
        "update": "更新します",
        "add": "に追加"
},
    "notFound": {
    "title": "404ページは存在しません",
        "description": "要求したページを見つけることができません。削除されているか、エラーでリンクされている可能性があります。 これがエラーだと思われる場合は、お問い合わせの作業注文を送信してください。",
        "p1": "{sec}の後にホームページに戻り、ブラウザが応答しない場合は、下のボタンをクリックします。",
        "manualBack": "ホームページに戻ります"
},
    "forbidden": {
    "title": "403権利なし",
        "description": "このページにアクセスするのに十分な許可がない場合があります。 これがエラーだと思われる場合は、お問い合わせの作業注文を送信してください。"
}
}
