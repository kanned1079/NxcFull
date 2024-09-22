export default {
    commonHeader: {
        menuTxt: 'Menu',
        userData: 'User Data',
        editUserData: 'Edit User Data',
        logout: 'Logout',
    },
    commonAside: {
        admin: {},
        user: {
            dashboard: 'Dashboard',
            document: 'Manual',
            subscribe: 'Subscribe',
            purchase: 'Purchase',
            surplus: 'Balance',
            fiance: 'Finance',
            myOrder: 'Order',
            myInvite: 'Invite',
            user: 'User',
            profile: 'Profile',
            support: 'Work Order',
            activateLog: 'Activation Log'
        }
    },
    userSummary: {
        myPlan: 'My Subscription',
        shortcut: 'Shortcut',
        timeLeft: 'Subscription is valid and will expire in {msg}.',
        toPurchase: 'Purchase Subscription',
        tutorial: {
            title: 'View Tutorial',
            content: 'Learn how to use {name}',
        },
        checkKey: {
            title: 'View Key',
            content: 'Quickly import the key into the corresponding client for use',
        },
        renewPlan: {
            title: 'Renew Subscription',
            content: 'Renew your current subscription',
        },
        support: {
            title: 'Encountered a problem',
            content: 'If you encounter a problem, you can communicate with our staff through work orders.',
        }
    },
    userDocument: {
        searchPlaceholder: 'Please enter the content to be searched (fuzzy search)',
        searchBtn: 'Search',
    },
    newPurchase: {
        headerPlaceholder: 'Choose the plan that suits you best',
        purchaseBtn: 'Purchase',
        monthPay: 'Monthly payment',
        moreMethod: 'More options',
    },
    newSettlement: {
        err: 'Error',
        monthPay: 'Monthly payment',
        quarterPay: 'Quarterly payment',
        halfYearPay: 'Half-year payment',
        yearPay: 'Annual payment',
        payCycle: 'Payment cycle',
        couponPlaceholder: 'Have a coupon?',
        verifyBtn: 'Verify',
        orderTotalTitle: 'Total order amount',
        total: 'Total',
        submitOrder: 'Submit order',
        coupon: 'Coupon',
        notify: {
            passTitle: 'Verification passed',
            couponVerified: 'Coupon is valid',
            couponInvalid: "Coupon is invalid",
            couponIsNull: 'The entered coupon code cannot be empty',
        }
    },
    // pass
    userProfile: {
        myWallet: 'My Wallet',
        walletSub: 'Account balance (only for consumption)',
        alertPwd: 'Change key',
        oldPwd: 'Old key',
        oldPwdSub: 'Please enter the old password',
        newPwd: 'New key',
        newPwdSub: 'Please enter the new key',
        newPwdAgain: 'Confirm key',
        newPwdAgainSub: 'Please enter the new key again',
        saveBtn: 'Save',
        notify: 'Notification',
        enableNotify: 'Enable expiration notification',
        deleteAccount: 'Cancel account',
        deleteAccountSub: 'Your account will be marked as deleted. If you need to use our services again, please re-register.',
        deleteBtn: 'Cancel my account',
        oldPwdVerifiedFailure: 'Old password verification failed',
        alertFailure: 'Key change failed',
        alertSuccess: 'Modified successfully',
        alertSuccessSub: 'Please log in with the new password.',
        errorPwdFormat: 'Password format error',
        pwdNotMatch: 'The two passwords do not match',
        oldPwdNotNull: 'The old password cannot be empty',
    },
};