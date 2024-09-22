export default {
    commonHeader: {
        menuTxt: 'メニュー',
        userData: 'ユーザーデータ',
        editUserData: 'ユーザーデータを編集する',
        logout: 'ログアウト',
    },
    commonAside: {
        admin: {},
        user: {
            dashboard: ' ダッシュボード ',
            document: ' マニュアル ',
            subscribe: ' 購読 ',
            purchase: ' 購入 ',
            surplus: ' 残高 ',
            fiance: ' 財務 ',
            myOrder: ' 注文 ',
            myInvite: ' 招待 ',
            user: ' ユーザー ',
            profile: ' プロフィール ',
            support: ' 依頼 ',
            activateLog: ' アクティベーションログ '
        }
    },
    userSummary: {
        myPlan: '私の購読',
        shortcut: 'ショートカット',
        timeLeft: '購読有効、{msg}で期限切れになります。',
        toPurchase: '購読を購入する',
        tutorial: {
            title: 'チュートリアルを見る',
            content: '{name}の使い方を学ぶ',
        },
        checkKey: {
            title: 'キーを確認する',
            content: 'キーを対応するクライアントに迅速にインポートして使用する',
        },
        renewPlan: {
            title: '購読を更新する',
            content: '現在の購読を更新する',
        },
        support: {
            title: '問題が発生しました',
            content: '問題が発生した場合はチケットで私たちのサポートチームとコミュニケーションできます。',
        }
    },
    userDocument: {
        searchPlaceholder: '検索する内容を入力してください（曖昧検索）',
        searchBtn: '検索',
    },
    newPurchase: {
        headerPlaceholder: 'あなたに最適なプランを選択してください',
        purchaseBtn: '購入',
        monthPay: '月払い',
        moreMethod: 'もっと選択',
    },
    newSettlement: {
        err: 'エラー',
        monthPay: '月払い',
        quarterPay: '四半期払い',
        halfYearPay: '半年払い',
        yearPay: '年払い',
        payCycle: '支払いサイクル',
        couponPlaceholder: 'クーポンあり？',
        verifyBtn: '検証',
        orderTotalTitle: '注文総額',
        total: '合計',
        submitOrder: '注文を提出する',
        coupon: 'クーポン',
        notify: {
            passTitle: '検証通過',
            couponVerified: 'クーポン有効',
            couponInvalid: "クーポン無効",
            couponIsNull: '入力したクーポンコードは空ではないようにしてください。',
        }
    },
    // pass
    userProfile: {
        myWallet: '私の財布',
        walletSub: 'アカウント残高（消費専用）',
        alertPwd: 'キーを変更する',
        oldPwd: '古いキー',
        oldPwdSub: '古いパスワードを入力してください',
        newPwd: '新しいキー',
        newPwdSub: '新しいキーを入力してください',
        newPwdAgain: 'キーを確認する',
        newPwdAgainSub: '新しいキーをもう一度入力してください',
        saveBtn: '保存',
        notify: '通知',
        enableNotify: '期限切れ通知を有効にする',
        deleteAccount: 'アカウントを削除する',
        deleteAccountSub: 'あなたのアカウントは削除される予定です。もう一度私たちのサービスを使用する場合は、再登録してください。',
        deleteBtn: '私のアカウントを削除する',
        oldPwdVerifiedFailure: '古いパスワードの検証に失敗しました',
        alertFailure: 'キーの変更に失敗しました',
        alertSuccess: '変更成功',
        alertSuccessSub: '新しいパスワードでログインしてください。',
        errorPwdFormat: 'パスワードの形式エラー',
        pwdNotMatch: '二つのパスワードが一致しません',
        oldPwdNotNull: '古いパスワードは空ではないようにしてください。',
    },
};