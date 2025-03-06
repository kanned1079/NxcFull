export default {
    "commonHeader": {
        "menuTxt": "메뉴",
        "userData": "사용자 정보",
        "editUserData": "사용자 정보 편집",
        "logout": "로그 아웃하십시오"
    },
    "commonAside": {
        "admin": {
            "dashboard": "식탁보",
            "queueMonit": "서비스 측 모니터링",
            "settings": "설정",
            "systemConfig": "시스템 설정",
            "paymentConfig": "지불 설정",
            "themeConfig": "테마 설정",
            "server": "섬기는 사람",
            "privilege": "권한 관리",
            "finance": "재원",
            "subscription": "가입 관리",
            "coupon": "쿠폰 관리",
            "order": "주문 관리",
            "activate": "레코드를 활성화하십시오",
            "key": "키 관리",
            "user": "사용자",
            "userMgr": "사용자 관리",
            "notice": "발표 관리",
            "ticket": "작업 주문 관리",
            "doc": "지식 데이터베이스 관리"
        },
        "user": {
            "dashboard": "식탁보",
            "document": "문서를 사용하십시오",
            "app": "앱 다운로드",
            "subscription": "구독하다",
            "purchase": "구독 구매",
            "surplus": "내 열쇠",
            "fiance": "재원",
            "topUp": "위로",
            "myOrder": "내 주문",
            "myInvite": "내 초대",
            "user": "사용자",
            "profile": "개인 센터",
            "support": "내 작업 순서",
            "activateLog": "레코드를 활성화하십시오"
        }
    },
    "adminViews": {
        "common": {
            "fetchDataSuccess": "데이터 수집이 성공적으로 수집됩니다",
            "fetchDataFailure": "데이터가 실패하면 나중에 다시 시도하십시오",
            "addSuccess": "성공적으로 추가되었습니다",
            "addFailure": "추가가 실패하면 나중에 시도하십시오",
            "updateSuccess": "수정이 성공적이었습니다",
            "updateFailure": "수정이 실패하면 나중에 다시 시도하십시오",
            "deleteSuccess": "성공적으로 삭제",
            "deleteFailure": "삭제가 실패한 경우 다시 시도하십시오",
            "confirm": "확인하다",
            "cancel": "취소",
            "success": "운영이 성공적입니다",
            "failure": "작동 실패",
            "NotAllowed": "불법 형식",
            "checkForm": "양식을 확인하십시오",
            "unknownErr": "알 수없는 오류",
            "dialog": { "delete": "삭제하고 싶다는 것을 알고 있습니까?" },
            "yes": "예",
            "no": "아니요"
        },
        "login": {
            "secureCard": {
                "title": "안전 검사",
                "securePath": "안전한 길",
                "hint": "시스템 보안을 보장하려면 관리자 로그인 페이지를 입력하기 전에 보안 경로를 입력해야합니다. 아래 입력 상자에 보안 경로를 입력하십시오. 보안 경로를 성공적으로 확인하면 후속 빠른 로그인을 위해 저장을 선택할 수 있습니다.",
                "placeholder": "안전한 길을 입력하십시오",
                "checkBtn": "조사하다",
                "rememberPath": "안전한 길을 기억하십시오"
            },
            "card": {
                "back": "홈페이지로 돌아갑니다",
                "toAdminCenter": "관리 센터에 로그인하십시오",
                "emailAddr": "관리자 이메일",
                "emailInputArea": {
                    "title": "관리자 이메일",
                    "placeholder": "이메일 주소"
                },
                "passwordInputArea": {
                    "title": "관리자 비밀번호",
                    "placeholder": "비밀번호"
                },
                "login": "로그인",
                "forgetPassword": "비밀번호를 잊어 버리십시오",
                "formNotPassed": "양식 검증이 실패합니다"
            },
            "mention": {
                "title": "힌트",
                "description": "이 페이지는 관리자가 관리자가 액세스 할 수있는 관리자 페이지입니다."
            },
            "message": {
                "passwordErr": "잘못된 비밀번호",
                "adminNotExist": "관리자는 존재하지 않습니다",
                "noPrivilege": "허가없이 액세스하십시오",
                "authPassed": "검증이 통과되었습니다",
                "authFailure": "확인이 실패했습니다",
                "otherErr": "다른 오류",
                "pathCheckPassed": "보안 경로 점검이 통과되었습니다",
                "pathCheckFailure": "잘못된 보안 경로",
                "rememberSecureMention": "백엔드 관리의 보안을 보장하려면 개인 컴퓨터가 아닌 경우 확인하지 마십시오."
            }
        },
        "summary": {
            "cockpit": "식탁보",
            "systemConfig": "시스템 설정",
            "paymentConfig": "지불 설정",
            "planMgr": "가입 관리",
            "userMgr": "사용자 관리",
            "orderMgr": "주문 관리",
            "keyMgr": "키 관리",
            "incomeText": "어제의 소득 / 월 소득",
            "pendingTicket": "{nums} 주문이 처리됩니다",
            "toHandle": "가서 확인하십시오",
            "apiAccessCard": "일주일 이내에 API 인터페이스 방문 수",
            "apiAccessCardHint": "이 데이터는 현재 API 액세스를 이해하기위한 것이며 서버 성능을 나타내지 않습니다.",
            "incomeWeek": "일주일 이내에 소득 금액",
            "incomeCardHint": "다음은 일주일에 걸친 소득 차트가 있습니다. 캐시가 지워지면 부정확 한 디스플레이가 발생합니다.",
            "core": "핵심",
            "reqErr": "오류가 발생했습니다",
            "reqErrHint": "개요 정보를 얻는 동안 오류가 발생 하여이 요청을 완료 할 수 없으므로 차트를 다시 표시 할 수 없습니다.",
            "userCard": {
                "title": "사용자 개요",
                "allRegisteredUsers": "일반 등록 사용자",
                "activeUsers": "라이브 사용자",
                "inactiveUsers": "비활성 사용자",
                "blockedUsers": "금지 또는 상환"
            },
            "general": {
                "title": "일반적으로",
                "localTime": "브라우저 시간",
                "osType": "운영 체제 유형",
                "appName": "응용 프로그램 이름",
                "appUrl": "응용 프로그램 URL",
                "currency": "동전 유닛",
                "allowRegister": "등록 허용"
            },
            "system": {
                "title": "시스템 구성",
                "axiosAddr": "HTTP 백엔드 주소",
                "wsAddr": "WebSocket 백엔드 주소",
                "serverTime": "서버 시간",
                "uptime": "운영 시간",
                "gatewayStatus": "API 네트워크 상태",
                "dbStatus": "데이터베이스 상태",
                "redisStatus": "Redis 상태",
                "serverOsType": "서버 운영 체제 유형",
                "serverOsArch": "서버 운영 체제 아키텍처",
                "runMode": "실행 모드",
                "cpuNums": "네트워크 서버의 CPU 코어 수",
                "numCgoCall": "쓰레기 수집 수",
                "time": "2 차",
                "paymentMethods": "결제 방법을 활성화합니다",
                "runOK": "정상적으로 실행하십시오",
                "runErr": "비정상적인 행동",
                "checkServer": "백엔드 서버의 환경 구성을 확인하십시오",
                "stopRegisterHint": "새로운 사용자 등록을 비활성화 한 것 같습니다",
                "toSetting": "설정으로 돌아갑니다"
            }
        },
        "queueMonit": {
            "title": "서비스 측 모니터링",
            "headerHint": "피크 네트워크 중에 자주 문의하면 네트워크 성능 및 데이터베이스 처리량에 영향을 미칩니다.",
            "latency": {
                "title": "서버 지연",
                "retry": "다시 테스트하십시오",
                "hint": "*여기서의 요청은 클라이언트가 서버에 요청을 보낸 후 클라이언트에 대한 성공적인 응답에 필요한 시간을 나타냅니다.",
                "unit": "밀리 초",
                "level": {
                    "l1": {
                        "title": "훌륭한",
                        "description": "이것은 매우 좋은 네트워크 상황이며 지연되지 않을 것입니다."
                    },
                    "l2": {
                        "title": "정상",
                        "description": "이것은 대부분 네트워크 상황이며, 사용자는 거의 지연이 느껴지지 않는다고 생각합니다."
                    },
                    "l3": {
                        "title": "가난한",
                        "description": "이 경우 사용자는 약간의 지연이나 지연을 느낄 수 있습니다."
                    },
                    "l4": {
                        "title": "차이점",
                        "description": "명백한 지연을 느끼고 사용자 경험에 영향을 줄 수 있습니다."
                    },
                    "l5": {
                        "title": "아주 나쁜",
                        "description": "명백한 지연이 있고 로딩 속도가 느리거나 새로 고칠 수 없으므로 사용자 상호 작용 및 경험에 심각한 영향을 미칩니다."
                    },
                    "offline": {
                        "title": "서버 꺼짐",
                        "description": "서버 또는 핸들 요청 오류에 연결할 수 없으면 구성이 올바른지 확인하십시오."
                    }
                }
            },
            "api": {
                "title": "지난 7 일 동안 API 요청 상태",
                "items": {
                    "ok": { "title": "성공 (상태)", "unit": "2 차" },
                    "notFound": { "title": "상태 경로 오류 (404)", "unit": "2 차" },
                    "unAuthorized": { "title": "무단 액세스 (401)", "unit": "2 차" },
                    "login2reg": { "title": "로그인 / 등록", "unit": "2 차" }
                }
            },
            "log": {
                "title": "저널 기록",
                "deleteLogMsg": "저널 삭제 {nums}",
                "deleteLogErr": "저널 삭제가 실패했습니다",
                "logTableRows": "총 일기 기록",
                "logTableSize": "블로그 테이블은 공간을 차지합니다",
                "unit": { "lines": "좋아요", "mb": "메가 문자" },
                "deleteLog": "저널을 삭제하십시오",
                "exportCsv": "수출 CSV",
                "deleteLogHint": "이것은 일주일 전에 모든 저널을 제거합니다.",
                "warn": {
                    "title": "저널을 삭제하고 싶다는 것을 알고 있습니까?",
                    "description": "목록은 즉시 삭제 되며이 작업을 취소 할 수 없습니다."
                },
                "export": {
                    "title": "CSV 파일을 내 보냅니다",
                    "description": "이것은 다음 테이블을 CSV 파일로 내보내고 다운로드를 열 수있는 권한이 없으면 내보내기를 클릭하십시오."
                },
                "table": {
                    "id": "틀",
                    "method": "요청 방법",
                    "path": "경로를 요청하십시오",
                    "code": "상태 코드",
                    "clientIp": "클라이언트 IP",
                    "responseTime": "처리 시간",
                    "requestAt": "요청 시간"
                }
            }
        },
        "systemConfig": {
            "title": "시스템 설정",
            "common": {
                "success": "구성이 성공적으로 업데이트되었습니다",
                "err": "업데이트 구성에 실패했습니다"
            },
            "site": {
                "common": { "title": "대지" },
                "appName": {
                    "title": "사이트 이름",
                    "shallow": "사이트 이름이 필요한 위치를 표시하는 데 사용됩니다.",
                    "placeholder": "사이트 이름"
                },
                "appSubName": {
                    "title": "사이트 자막",
                    "shallow": "일반적으로 주 제목 아래에 표시됩니다.",
                    "placeholder": "부제"
                },
                "appDescription": {
                    "title": "사이트 설명",
                    "shallow": "사이트가 필요한 위치를 표시하는 데 사용됩니다.",
                    "placeholder": "사이트 설명"
                },
                "appUrl": {
                    "title": "사이트 URL",
                    "shallow": "최신 웹 사이트 주소는 필요한 경우 URL에 표시됩니다.",
                    "placeholder": "사이트 URL"
                },
                "forceHttps": {
                    "title": "힘 HTTPS",
                    "shallow": "사이트가 HTTP, CDN 또는 반세대를 사용하여 강제 HTTP를 활성화하지 않는 경우."
                },
                "logoUrl": {
                    "title": "심벌 마크",
                    "shallow": "로고가 필요한 위치를 표시하는 데 사용됩니다.",
                    "placeholder": "로고의 URL 주소"
                },
                "subscribeUrl": {
                    "title": "URL 가입",
                    "shallow": "구독에 사용하려면 사이트 URL로 비워 두십시오. 무작위로 가져 오려면 여러 구독 URL이 필요한 경우 쉼표를 사용하여 분할하십시오.",
                    "placeholder": "URL 가입"
                },
                "tosUrl": {
                    "title": "사용자 (TOS) URL의 약관",
                    "shallow": "사용자 (TOS)의 조건으로 이동하는 데 사용됩니다.",
                    "placeholder": "사용자 이용 약관 주소"
                },
                "stopRegister": {
                    "title": "새로운 사용자 등록을 중단하십시오",
                    "shallow": "켜진 후에 아무도 등록 할 수 없습니다."
                },
                "inviteRequire": {
                    "title": "강제 초대",
                    "shallow": "열면 새 사용자에 등록 할 때 초대장 코드를 작성해야합니다."
                },
                "trialSubscribe": {
                    "title": "시험에 등록하십시오",
                    "shallow": "옵션이없는 경우 구독을 선택하십시오. 구독 관리로 이동하여 먼저 추가하십시오."
                },
                "trialTime": {
                    "title": "시험 시간 (시간)",
                    "shallow": "신규 사용자가 가입 할 때 시험 시간을 구독하십시오."
                },
                "currency": {
                    "title": "통화 장치",
                    "shallow": "디스플레이 사용에만 사용되며 변경 후 시스템의 모든 통화 장치가 변경됩니다.",
                    "placeholder": "CNY"
                },
                "currencySymbol": {
                    "title": "통화 기호",
                    "shallow": "디스플레이 사용에만 사용되며 변경 후 시스템의 모든 통화 장치가 변경됩니다.",
                    "placeholder": "¥"
                }
            },
            "security": {
                "common": { "title": "보안 설정" },
                "emailVerify": {
                    "title": "이메일 확인",
                    "description": "열면 사용자는 이메일 확인을 수행해야합니다."
                },
                "gmailAlias": {
                    "title": "Gmail 다중 별칭을 금지합니다",
                    "description": "열면 여러 Gmail 별칭이 등록되지 않습니다."
                },
                "safeMode": {
                    "title": "안전 모드",
                    "description": "열린 후 사이트 URL을 제외한이 사이트에 대한 도메인 이름 액세스는 403이됩니다."
                },
                "adminPath": {
                    "title": "배경 경로",
                    "description": "배경 관리 경로는 수정 후 원래 관리자 경로를 변경합니다.",
                    "placeholder": "https://x.com/logo.jpeg"
                },
                "emailWhitelist": {
                    "title": "이메일 접미사 화이트리스트",
                    "description": "등록은 목록에 사서함 접미사를 열린 후에 만 ​​허용됩니다."
                },
                "recaptcha": {
                    "title": "안티 로봇",
                    "description": "활성화되면 로봇을 방지 할 수 있도록 HCAPTCHA가 활성화됩니다."
                },
                "hCaptchaSiteKey": {
                    "title": "hcaptcha sitekey",
                    "description": "SiteKey는 HCAPTCHA 서버에 웹 사이트 번호를 식별하도록 요청하는 데 사용됩니다.",
                    "placeholder": "A3CA066C-0-42FE-BCD2-23F4B48D528"
                },
                "ipRegisterLimit": {
                    "title": "IP 등록 제한",
                    "description": "개방 후 IP 등록 계정이 규칙 요구 사항을 충족하면 등록이 제한됩니다."
                },
                "registerTimes": {
                    "title": "빈도",
                    "description": "등록 수에 도달 한 후에 처벌이 활성화됩니다.",
                    "placeholder": "입력하십시오"
                },
                "lockTime": {
                    "title": "처벌 시간 (분)",
                    "description": "다시 등록하기 전에 페널티 타임이 통과 될 때까지 기다려야합니다.",
                    "placeholder": "입력하십시오"
                }
            },
            "frontend": {
                "common": { "title": "개인화" },
                "themePropColor": {
                    "default": "기본",
                    "darkBlueDay": "진한 파란색",
                    "milkGreenDay": "우유 녹색",
                    "bambooGreen": "Ruozhu",
                    "mistyPine": "안개 소나무",
                    "glacierBlue": "빙하 블루",
                    "grayTheme": "회색"
                },
                "sidebarStyle": {
                    "title": "사이드 바 스타일",
                    "shallow": "사이드 바의 색상을 설정하십시오."
                },
                "headerStyle": {
                    "title": "헤드 스타일",
                    "shallow": "상단에 색상을 설정하십시오."
                },
                "themeColor": {
                    "title": "테마 색상",
                    "shallow": "전체 웹 페이지의 테마 색상을 설정하십시오."
                },
                "background": {
                    "title": "배경",
                    "shallow": "배경 로그인 페이지에 표시됩니다.",
                    "placeholder": "https://x.com/logo.jpeg"
                }
            },
            "inviteAndRebate": {
                "common": { "title": "지불 및 리베이트" },
                "inviteRebateEnable": {
                    "title": "사용자 리베이트를 활성화합니다",
                    "description": "활성화 된 경우, 초대 된 사용자가 재충전하면 사용자는 아래에 설정된 재충전 비율에 따라 리베이트됩니다."
                },
                "inviteRebateRate": {
                    "title": "리베이트 비율",
                    "description": "리베이트 금액 비율을 설정하십시오.",
                    "placeholder": "리베이트 비율을 입력하십시오"
                },
                "discountInfo": {
                    "title": "할인 정보",
                    "description": "제안 정보를 설정하면 상단 업 페이지 상단에 표시됩니다.",
                    "placeholder": "할인 정보를 설정합니다"
                },
                "inviteInfo": {
                    "title": "초대 정보",
                    "description": "초대장 정보를 설정하면 리베이트 비율을 표시하기 위해 사용자 초대장 페이지에 표시됩니다.",
                    "placeholder": "리베이트 정보를 설정하십시오"
                }
            },
            "welcome": {
                "common": { "title": "홈페이지 정보" },
                "homeDescription": {
                    "title": "홈 페이지 설명",
                    "description": "홈페이지에 대한 간단한 설명을 설정하십시오.",
                    "placeholder": "홈페이지 설명 내용을 입력하십시오"
                },
                "whyChooseUs": {
                    "title": "왜 우리를 선택하십시오",
                    "description": "우리가 선택한 이유에 대한 설명을 설정하십시오.",
                    "placeholder": "자세한 설명을 입력하십시오"
                },
                "bilibiliLink": {
                    "title": "빌리 빌리 공식 링크",
                    "description": "Bilibili 공식 계정의 링크 주소를 설정하십시오.",
                    "placeholder": "https://space.bilibili.com/xxxx"
                },
                "youtubeLink": {
                    "title": "YouTube 공식 링크",
                    "description": "공식 YouTube 계정의 링크 주소를 설정하십시오.",
                    "placeholder": "https://youtube.com/channel/xxxxx"
                },
                "instagramLink": {
                    "title": "Instagram 공식 링크",
                    "description": "공식 Instagram 계정의 링크 주소를 설정하십시오.",
                    "placeholder": "https://instagram.com/xxxx"
                },
                "wechatLink": {
                    "title": "WeChat 공개 계정 링크",
                    "description": "WeChat 공개 계정의 링크 주소를 설정하십시오.",
                    "placeholder": "WeChat 공개 링크를 입력하십시오"
                },
                "filingNumber": {
                    "title": "등록 번호",
                    "description": "사이트의 등록 번호를 설정하십시오.",
                    "placeholder": "예를 들어 : Guangdong ICP No. 12345678"
                },
                "pageSuffix": {
                    "title": "사이트 접미사",
                    "description": "제목 디스플레이의 사이트 이름 접미사를 설정하십시오.",
                    "placeholder": "예를 들어 :- 사이트 이름"
                }
            },
            "sendMail": {
                "common": {
                    "title": "메일 설정",
                    "warning": "이 페이지의 구성을 변경하면 큐 서비스와 역 방향을 다시 시작해야합니다. 또한이 페이지의 구성 우선 순위는 .env의 메일 구성보다 높습니다.",
                    "sendTestMailTitle": "테스트 이메일 보내기",
                    "sendTestMailShallow": "이메일은 현재 로그인 관리자의 이메일 주소로 전송됩니다.",
                    "sendTestMailTo": "테스트 이메일을 보내십시오",
                    "sending": "이메일 보내기",
                    "success": "성공",
                    "receiverAddr": "받은 주소",
                    "senderHost": "배달 서버",
                    "senderPort": "포트를 보내십시오",
                    "senderEncrypt": "메시지 전송을위한 암호화 방법",
                    "senderUsername": "신용 이름",
                    "sendErr": "이메일 실패를 보냅니다"
                },
                "smtpServerAddress": {
                    "title": "SMTP 서버 주소",
                    "shallow": "이메일 서비스 제공 업체가 제공하는 서비스 주소",
                    "placeholder": "서버 주소를 입력하십시오"
                },
                "smtpServerPort": {
                    "title": "SMTP 서비스 포트",
                    "shallow": "공통 포트는 25, 465, 587입니다",
                    "placeholder": "포트 번호를 입력하십시오"
                },
                "smtpEncryption": {
                    "title": "SMTP 암호화 방법",
                    "shallow": "465 포트 암호화 방법은 일반적으로 SSL이고 587 포트 암호화 방법은 일반적으로 TLS입니다.",
                    "placeholder": "암호화 방법을 입력하십시오"
                },
                "smtpAccount": {
                    "title": "SMTP 계정",
                    "shallow": "이메일 서비스 제공 업체가 제공하는 계정",
                    "placeholder": "계정 번호를 입력하십시오"
                },
                "smtpPassword": {
                    "title": "SMTP 비밀번호",
                    "shallow": "이메일 서비스 제공 업체가 제공하는 비밀번호",
                    "placeholder": "비밀번호를 입력하십시오"
                },
                "senderAddress": {
                    "title": "주소를 보내십시오",
                    "shallow": "메일 서비스 제공 업체가 제공 한 전송 주소",
                    "placeholder": "보내는 주소를 입력하십시오"
                },
                "emailTemplate": {
                    "title": "이메일 템플릿",
                    "shallow": "문서에서 이메일 템플릿을 사용자 정의하는 방법을 볼 수 있습니다.",
                    "placeholder": "이메일 템플릿을 선택하십시오"
                }
            },
            "notice": {
                "common": { "title": "알림 설정" },
                "displayName": {
                    "title": "표시 이름",
                    "shallow": "프론트 엔드 페이지 디스플레이 용",
                    "placeholder": "일반 알림 인터페이스 1"
                },
                "barkEndpoint": {
                    "title": "껍질 액세스 포인트",
                    "shallow": "Bark Server 백엔드 API 주소",
                    "placeholder": "https : // <ip> : <port>/<secret-key>"
                },
                "barkGroup": {
                    "title": "껍질 그룹",
                    "shallow": "클라이언트가 표시하는 그룹 이름",
                    "placeholder": "편물"
                }
            },
            "appDownload": {
                "common": {
                    "title": "앱",
                    "hint": "자체 클라이언트 (앱)의 버전 관리 및 업데이트 용"
                },
                "enabled": {
                    "title": "부하를 열고 떨어 뜨립니다",
                    "shallow": "활성화 된 경우 사용자가 다운로드 페이지에 액세스 할 수 있습니다"
                },
                "windows": {
                    "title": "창",
                    "shallow": "Windows 버전 번호 및 다운로드 주소",
                    "placeholder": "https://xxxx.com/xxx.exe"
                },
                "macos": {
                    "title": "마코스",
                    "shallow": "MacOS 버전 번호 및 다운로드 주소",
                    "placeholder": "https://xxxx.com/xxx.dmg"
                },
                "linux": {
                    "title": "리눅스",
                    "shallow": "Linux 버전 번호 및 다운로드 주소",
                    "placeholder": "https://xxxx.com/xxx.deb"
                },
                "android": {
                    "title": "기계적 인조 인간",
                    "shallow": "안드로이드 버전 번호 및 다운로드 주소",
                    "placeholder": "https://xxxx.com/xxx.apk"
                },
                "ios": {
                    "title": "iOS",
                    "shallow": "iOS 버전 번호 및 다운로드 주소",
                    "placeholder": "https://xxxx.com/xxx.ipk"
                }
            }
        },
        "payConfig": {
            "title": "지불 설정",
            "description": "지원되는 모든 결제 방법은 현재 관리 할 수 ​​있지만 지불 가방 지불만으로도 다른 결제 방법을 충족하지 않으면 프로젝트가 더 개선 될 때까지 기다릴 수 있습니다.",
            "attention": {
                "title": "주목해야 할 것",
                "point1": "결제 방법 정보를 활성화하기 전에 구성하는 것이 정말 중요합니다.",
                "point2": "결제 메소드 구성을 수정하면 \"---\"로 표시되면 옵션이 설정되었으며 다시 수정 해야하는 경우 새 값을 입력해야합니다."
            },
            "common": {
                "detail": "{method} 구성",
                "fillAttention": "안전을 보장하기 위해 자세한 정보가 표시되지 않으므로 기존 구성을 작성하거나 덮어 쓰도록 리필하십시오.",
                "discountAmount": "할인 금액 (사용자 프롬프트 정보는 시스템에서 설정 한 \"지불 및 리베이트\"에서 설정할 수 있음)",
                "saveConfigBtnHint": "구하다",
                "cancelBtnHint": "취소",
                "saveSuccess": "구성이 성공적으로 저장되었습니다",
                "alterSuccess": "구성 수정이 성공적이었습니다",
                "discountPlaceholder": "할인 금액 (재충전 금액이 할인 금액보다 큰 경우)",
                "saveOrAlterFailure": "알 수없는 오류"
            },
            "alipay": {
                "title": "지불",
                "config": {
                    "appId": "응용 프로그램 ID",
                    "appPublicKeyCertContent": "응용 프로그램 키 인증서 컨텐츠",
                    "appPrivateKey": "애플리케이션 개인 정보 보호 키",
                    "alipayRootCert": "지불 가방 인증서",
                    "alipayPublicKeyCertContent": "지불 보물 증명서 내용"
                }
            },
            "wechat": {
                "title": "wechat 지불",
                "config": {
                    "mchId": "가맹점 ID",
                    "mchCertSerial": "판매자 인증서 일련 번호",
                    "apiV3Key": "API V3 키",
                    "privateKey": "개인 정보 보호 키"
                }
            },
            "apple": {
                "title": "Apple Pay",
                "config": {
                    "issuerId": "발행자 ID",
                    "bundleId": "번들 ID",
                    "privateKeyId": "개인 키 ID",
                    "privateKeyContent": "개인 정보 보호 내용"
                }
            },
            "addPaymentMethod": "결제 방법을 추가하십시오",
            "enableBtnHint": "시작",
            "disableBtnHint": "장애가 있는",
            "setupPaymentMethod": "구성"
        },
        "themeConfig": {
            "title": "주제 구성",
            "using": "사용,",
            "setAsCurrent": ""
        },
        "groupMgr": {
            "title": "권한 관리",
            "description": "올바른 그룹은 다른 가입 수준을 표시하는 데 사용되며 쉽게 관리하기 위해 올바른 그룹에서 동일한 레벨이지만 다른 크기를 구독 할 수 있습니다.",
            "common": {
                "addNewGroup": "새로운 권한 그룹",
                "alterGroupName": "권한 그룹의 이름을 수정하십시오",
                "addConfirmed": "ADD를 확인하십시오",
                "alterConfirmed": "수정을 확인하십시오",
                "cancel": "취소",
                "addSuccess": "성공적으로 추가되었습니다",
                "addFailure": "실패를 추가하십시오",
                "alterSuccess": "수정 권한 그룹이 성공했습니다",
                "alterFailure": "수정 권한 그룹에 실패했습니다",
                "delSuccess": "성공적으로 삭제",
                "delFailure": "삭제가 실패했습니다",
                "delMention": "목록은 즉시 삭제 되며이 작업을 취소 할 수 없습니다. 구독 계획의 관련 그룹은주의해서 삭제해야합니다.",
                "delNotAllowed": "이 올바른 그룹에는 관련 리소스가 있으며 삭제할 수 없습니다."
            },
            "groupId": "권리 그룹 ID",
            "groupName": "허가 그룹의 이름",
            "groupPlaceholder": "허가 그룹 이름을 입력합니다",
            "userCount": "사용자 수",
            "planCount": "구독 수량",
            "operate": "작동하다",
            "editBtnHint": "편집하다",
            "deleteBtnHint": "삭제"
        },
        "docMgr": {
            "title": "지식 데이터베이스 관리",
            "description": "여기에서는 사용자에게 설명 파일을 작성할 수 있습니다. 소프트웨어 디자인 매뉴얼, 튜토리얼 또는 예방 조치 등이 될 수 있습니다.",
            "addDoc": "새 문서를 추가하십시오",
            "addSuccess": "새 문서를 성공적으로 추가합니다",
            "addFailure": "파일 추가에 실패했습니다",
            "titleNotEmpty": "문서의 제목은 비어있을 수 없습니다",
            "table": {
                "docId": "틀",
                "isShow": "표시 여부",
                "sortAs": "종류",
                "lang": "언어",
                "category": "범주",
                "title": "제목",
                "createdAt": "창조 시간",
                "updatedAt": "업데이트 시간",
                "operate": "작동하다",
                "edit": "편집하다",
                "delete": "삭제"
            },
            "form": {
                "add": "문서를 추가하십시오",
                "edit": "문서를 편집하십시오",
                "cancel": "취소",
                "confirm": "확인하다",
                "addBtn": "추가",
                "editBtn": "개정하다",
                "title": {
                    "title": "파일 제목",
                    "placeholder": "파일 제목을 입력하십시오"
                },
                "sort": {
                    "title": "종류",
                    "placeholder": "입력 파일의 순서 수준이 높을수록 값이 높을수록 우선 순위가 높습니다."
                },
                "category": {
                    "title": "범주",
                    "placeholder": "문서는이 필드에 따라 카테고리로 표시됩니다."
                },
                "lang": {
                    "title": "문서 언어",
                    "placeholder": "문서 언어를 선택하십시오"
                }
            }
        },
        "planMgr": {
            "title": "가입 관리",
            "description": "여기에서는 새 구독 계획을 추가하고 구독 계획에 이미 포함 된 설명, 가격, 잔액, 권한 그룹을 수정할 수 있습니다.",
            "addNewPlan": "새 구독을 추가하십시오",
            "table": {
                "id": "틀",
                "isSale": "판매 활성화",
                "isRenew": "연속 요금을 허용하십시오",
                "sort": "정렬 레벨",
                "group": "소위 권리 그룹",
                "name": "이름",
                "nums": "수량",
                "residue": "양",
                "operate": "작동하다",
                "operateBtn": { "update": "개정하다", "delete": "삭제" }
            },
            "mention": {
                "common": {
                    "success": "성공",
                    "failure": "실패하다",
                    "delMention": "구독 계획이 이미 판매중인 경우 삭제하십시오."
                }
            },
            "form": {
                "title": "구독을 추가하십시오",
                "items": {
                    "name": {
                        "title": "구독 이름",
                        "placeholder": "가입 계획의 이름을 입력하십시오"
                    },
                    "describe": {
                        "title": "구독 설명",
                        "placeholder": "구독에 대한 설명을 입력하십시오 (지원 Markdown)"
                    },
                    "isSale": { "title": "판매 활성화" },
                    "isRenew": { "title": "지속적인 수수료를 활성화합니다" },
                    "group": {
                        "title": "소위 권리 그룹",
                        "placeholder": "정당한 그룹을 선택하십시오"
                    },
                    "capacityLimit": {
                        "title": "허용 사용자의 최대 수",
                        "placeholder": "허용 사용자의 최대 수"
                    },
                    "planResidue": {
                        "title": "나머지 금액",
                        "placeholder": "나머지 금액"
                    },
                    "sort": { "title": "종류", "placeholder": "프론트 엔드 분류" },
                    "periodPlaceholder": {
                        "month": "월별 지불 가격을 입력하십시오",
                        "quarter": "분기 별 유료 가격을 입력하십시오",
                        "halfYear": "반년 지불 가격을 입력하십시오",
                        "year": "연간 지불 가격을 입력하십시오"
                    }
                }
            }
        },
        "couponMgr": {
            "title": "쿠폰 관리",
            "description": "여기에는 특정 축제 등에 대한 쿠폰을 추가 할 수 있으며,이 쿠폰은 주문할 때 사용자가 사용할 수 있으며 설정 한 비율에 따라 할인을 제공 할 수 있습니다.",
            "fetchErr": "데이터가 실패하면 다시",
            "fetchPlanKvFailurese": "구독 계획 목록이 실패했습니다",
            "table": {
                "id": "틀",
                "enabled": "활성화 여부",
                "name": "이름",
                "code": "쿠폰",
                "percentOff": "할인 정보",
                "capacity": "총 수량",
                "residue": "나머지 금액",
                "startTime": "시간을 활성화하십시오",
                "endTime": "종료 시간",
                "actions": "작동하다",
                "edit": "편집하다",
                "delete": "삭제"
            },
            "modal": {
                "newCoupon": "새 쿠폰을 추가하십시오",
                "editCoupon": "쿠폰 정보를 수정하십시오",
                "confirmAdd": "ADD를 확인하십시오",
                "confirmEdit": "수정을 확인하십시오",
                "emptyNotAllow": "이 항목이 필요합니다",
                "delMention": "쿠폰은 항목이 삭제 된 직후에 유효하지 않으며이 작업을 인출 할 수 없습니다.",
                "cancel": "취소",
                "name": {
                    "title": "쿠폰 이름",
                    "placeholder": "쿠폰 이름을 입력하십시오"
                },
                "code": {
                    "title": "쿠폰",
                    "placeholder": "쿠폰 코드를 사용자 정의하십시오 (비워두고 가능한 한 빨리 생성)"
                },
                "percentOff": {
                    "title": "정보 제공 (백분율)",
                    "placeholder": "할인 비율을 입력하십시오 (예 : -12%할인)"
                },
                "activeRange": { "title": "이용 가능한 쿠폰의 면제 범위" },
                "capacity": {
                    "title": "사용 된 최대 쿠폰 수",
                    "placeholder": "최대 사용량 제한 제한 (비어있는 경우 제한 없음)"
                },
                "residue": {
                    "title": "사용 된 나머지 쿠폰의 수",
                    "placeholder": "나머지 쿠폰 사용 횟수를 설정"
                },
                "perUserLimit": {
                    "title": "각 사용자가 쿠폰을 사용할 수있는 횟수",
                    "placeholder": "각 사용자가 사용할 수있는 횟수를 제한합니다 (공간 제한 없음)"
                },
                "planLimit": {
                    "title": "구독 계획을 지정하십시오",
                    "placeholder": "제안을 사용하기위한 지정된 가입 계획의 제한 (공간 제한 없음)"
                }
            }
        },
        "orderMgr": {
            "title": "주문 관리",
            "description": "여기서는 모든 구독 계획 주문, 다른 사용자의 필터 주문, 사용자의 주문을 수동으로 처리하는 등을 볼 수 있습니다.",
            "table": {
                "id": "틀",
                "orderId": "주문 번호",
                "email": "사용자 사서함",
                "status": {
                    "title": "유형",
                    "t1": "새로 구입했습니다",
                    "t2": "계속해서 수수료",
                    "t3": "편집하다"
                },
                "planName": "계획 이름",
                "period": "주기",
                "group": "권리 그룹",
                "amount": "실제 지불 금액",
                "price": "원래 가격",
                "isSuccess": {
                    "title": "주문 상태",
                    "cancelOrder": "주문을 취소하십시오",
                    "passOrder": "순서로"
                },
                "createdAt": "주문 생성 시간",
                "action": {
                    "title": "작동하다",
                    "showDetail": "세부 사항을 보여줍니다"
                }
            },
            "search": "문의 명령",
            "resetSearch": "쿼리 재설정",
            "failureReason": "실패의 원인",
            "couponId": "쿠폰 ID",
            "couponName": "쿠폰 이름",
            "noEntry": "없이",
            "orderDetail": "주문 세부 사항",
            "searchModal": {
                "email": {
                    "title": "사용자 사서함",
                    "placeholder": "사용자 이메일 입력 (퍼지 검색)"
                },
                "sort": {
                    "title": "정렬 알고리즘",
                    "placeholder": "정렬 알고리즘을 선택하십시오",
                    "ASC": "오름차순 순서",
                    "DESC": "하강 순서"
                }
            },
            "tradeWaiting": "지불되지 않습니다",
            "tradeFailure": "거래에 실패했습니다",
            "tradeSuccess": "성공"
        },
        "userMgr": {
            "userManager": "사용자 관리",
            "description": "직원 및 관리자를 포함한 모든 사용자를 관리하고 관리 권한을 부여하거나 취소하고 사용자 잔액 설정, 비밀번호 재설정, 수동으로 새 사용자 추가 등을 관리 할 수 ​​있습니다.",
            "enterEmail": "이메일을 입력하십시오",
            "enterValidEmail": "올바른 이메일 형식을 입력하십시오",
            "enterPassword": "비밀번호를 입력하십시오",
            "passwordMinLength": "비밀번호 길이는 6 자리 미만일 수 없습니다",
            "passwordMaxLength": "비밀번호 길이는 20 자리를 초과 할 수 없습니다",
            "passwordStrength": "비밀번호에는 문자, 숫자 및 특수 문자가 포함되어야합니다",
            "validationSuccess": "검증이 통과되었습니다",
            "validationFailed": "양식 확인 실패, 입력 항목을 확인하십시오",
            "editProfile": "정보 편집",
            "banUser": "로그인을 비활성화합니다",
            "unbanUser": "사용자를 차단 해제하십시오",
            "noRecord": "기록되지 않았습니다",
            "normal": "정상",
            "banned": "금지",
            "deleted": "로그 아웃하십시오",
            "nullContent": "널",
            "balance": "양",
            "orderCount": "주문 수량",
            "planCount": "계획",
            "actions": "작동하다",
            "updateSuccess": "성공적으로 업데이트하십시오",
            "addUserSuccess": "새로운 사용자 추가 성공",
            "unknownError": "알 수없는 오류",
            "email": "이메일",
            "registerDate": "등록일",
            "isAdmin": "관리자",
            "isStaff": "직원",
            "accountStatus": "계정 상태",
            "inviteCode": "초대장 코드",
            "query": "질문",
            "reset": "다시 놓기",
            "addNewUser": "추가 사용자",
            "searchUser": "문의 사용자",
            "cancel": "취소",
            "submit": "제출하다",
            "userEmail": "사용자 사서함",
            "inputUserEmail": "사용자 이메일 입력 (퍼지 검색)",
            "inputEmail": "이메일을 입력하십시오",
            "password": "비밀번호",
            "inputPassword": "비밀번호를 입력하십시오",
            "editUser": "사용자 편집",
            "no": "아니요",
            "yes": "예",
            "mention": {
                "title": "계정을 비활성화 하시겠습니까?",
                "content": "사용자가 차단되면 사용자는 {AppName}에 로그인 할 수 없으며 이와 관련된 모든 리소스를 사용할 수 없습니다."
            }
        },
        "activation": {
            "activateLog": "레코드를 활성화하십시오",
            "description": "판매 된 모든 키의 특정 활성화 상태를보고 클라이언트의 식별 코드, 활성화 시간 등을 볼 수 있습니다.",
            "click2getKey": "주요 콘텐츠를 얻으려면 클릭하십시오",
            "createdAt": "창조 시간",
            "turn2keyPage": "키로 돌아갑니다",
            "showKey": "열쇠를 보여주십시오",
            "email": "이메일",
            "key": "열쇠",
            "search": "찾다",
            "filter": "상영",
            "filterAll": "모두",
            "filterActive": "효과적입니다",
            "sortAlgorithm": "정렬 알고리즘",
            "sortASC": "오름차순 순서",
            "sortDESC": "하강 순서"
        },
        "keysMgr": {
            "keyMgr": "키 관리",
            "description": "생성 된 모든 주요 내용 및 사용 상태, 유효성 기간 등을 확인할 수 있습니다. 키를 비활성화 할 수도 있습니다.",
            "enableKey": "키를 활성화하십시오",
            "disableKey": "키를 비활성화합니다",
            "table": {
                "email": "이메일 주소",
                "key": "주요 내용",
                "isValid": "유효합니까?",
                "isUsed": "사용 여부",
                "isExpired": "만료되는지 여부",
                "active": "활동적인",
                "inactive": "만료되었습니다",
                "yes": "효율적인",
                "no": "만료되었습니다",
                "ok": "정상",
                "blocked": "상태를 비활성화합니다",
                "unUsed": "사용되지 않습니다",
                "used": "사용된",
                "showDetail": "세부 사항을 보여줍니다",
                "blockKey": "키를 비활성화합니다",
                "unblockKey": "받지 못한"
            },
            "searchModal": {
                "searchMethod": "문의 방법",
                "byId": "ID의 쿼리",
                "byContent": "키를 통한 문의 (퍼지)",
                "keyId": "키 ID",
                "idPlaceholder": "여기에 키의 ID를 입력하십시오",
                "keyContent": "예어",
                "contentPlaceholder": "여기에 키의 일부를 입력하십시오"
            },
            "mention": {
                "blockOk": "키 비활성화 성공적 ID : {id}",
                "title": "키를 비활성화 하시겠습니까?",
                "content": "사용자 경험을 보장하려면 키가 비활성화되면 키의 내용을 다시 확인하십시오."
            },
            "detailModal": {
                "title": "주요 세부 사항",
                "userId": "사용자 ID",
                "planName": "구독 계획 이름",
                "expiredAt": "만료 날짜",
                "keyGeneratedAt": "주요 생성 날짜",
                "clientId": "클라이언트 ID"
            }
        },
        "noticeMgr": {
            "title": "발표 관리",
            "description": "공지 사항은 여기에서 관리 할 수있는 공지 사항에 표시됩니다.",
            "addNotice": "공지 사항을 추가하십시오",
            "table": {
                "id": "틀",
                "show": "표시 여부",
                "title": "제목",
                "createdAt": "창조 시간",
                "action": { "title": "작동하다", "edit": "편집하다", "delete": "삭제" }
            },
            "mention": {
                "title": "삭제하고 싶다는 것을 알고 있습니까?",
                "content": "발표는 즉시 삭제 되며이 조치를 철회 할 수 없습니다."
            },
            "modal": {
                "addNew": "새로운 발표",
                "title": {
                    "title": "발표 제목",
                    "placeholder": "휠 캐스트에서 큰 타이틀로 표시하십시오"
                },
                "content": {
                    "title": "발표 내용",
                    "placeholder": "발표를 작성하는 주요 내용"
                },
                "tag": {
                    "title": "발표 태그",
                    "placeholder": "공지 사항에 대한 태그를 입력하십시오"
                },
                "img": {
                    "title": "배경 이미지 URL",
                    "placeholder": "채우지 않으면 기본 배경을 사용하십시오"
                }
            }
        },
        "adminTicket": {
            "ticketMgr": "작업 주문 관리",
            "description": "여기에서 제출 된 모든 사용자는 세 가지 긴급한 작업 주문이 있으며, 비상 업무 주문에서 처리를 시작하는 것이 좋습니다.",
            "pendingTicket": "처리 할 작업 주문",
            "finishedTicket": "완료된 작업 순서",
            "type": {
                "pendingDescription": "라이브 작업 순서, 이것은 작업 순서가 완료된 것으로 확인되면 직장이 닫히지 않으면 항상 여기에 배치됩니다.",
                "finishedDescription": "완료된 작업 주문, 여기에서 볼 수 있습니다."
            },
            "chooseOneNecessary": "하나 이상의 항목을 선택해야합니다",
            "mention": {
                "title": "티켓을 닫고 싶습니까?",
                "content": "주문은 완료된 주문으로 보관되어 컨텐츠를 다시 볼 수 있지만 주문에 응답 할 수 없습니다."
            }
        }
    },
    "userLogin": {
        "loginToContinue": "계속하려면 로그인하십시오",
        "email": "이메일 주소",
        "password": "비밀번호",
        "haveNoAccount": "아직 계정이 없습니까?",
        "login": "로그인",
        "reg": "지금 등록하십시오",
        "otherMethods": "또는 다른 방법을 사용하여 계속하십시오",
        "github": "Github 계정을 계속하십시오",
        "microsoft": "Microsoft 계정을 계속하십시오",
        "apple": "Apple 계정을 계속하십시오",
        "google": "Google 계정을 계속하십시오",
        "backHomePage": "홈페이지로 돌아갑니다",
        "form": {
            "emailRequire": "이메일 주소가 필요합니다",
            "passwordRequire": "아직 비밀번호를 입력하지 않았습니다"
        },
        "authPass": "검증이 통과되었습니다",
        "loginFailure": "로그인이 실패했습니다",
        "checkForm": "양식을 확인하십시오",
        "if2FaEnabledHint": "2 단계 확인을 활성화 한 경우 (필요하지 않음)",
        "reqErr": "오류가 있으면 나중에 다시 시도하십시오",
        "accountLocked": "귀하의 계정은 등록되거나 금지 될 수 있으며 현재로서는 오류라고 생각되면 기술 지원에 문의하십시오.",
        "tokenNotExist": "토큰을 제공하십시오"
    },
    "userRegister": {
        "backHomePage": "홈페이지로 돌아갑니다",
        "newAccount": "새 계정을 준비하십시오",
        "email": "이메일 주소",
        "verifyCode": "이메일 확인 코드",
        "invalidEmailFormat": "이메일 형식은 불법입니다",
        "sendVerifyCode": "이메일 확인 코드를 보내십시오",
        "sendSuccess": "이메일이 성공적으로 전송되었습니다. 이메일을 확인하십시오.",
        "pwd": "비밀번호",
        "pwdAgain": "비밀번호를 확인하십시오",
        "inviteCode": "초대장 코드 (선택 사항)",
        "agreement": "나는 읽고 동의했다",
        "terminalUserAgreement": "사용자 용어",
        "reg": "등록하다",
        "infoIncomplete": "불완전한 정보",
        "pwdIncorrect": "비밀번호가 일관되지 않습니다",
        "regSuccess": "등록이 성공적이며 로그인으로 돌아갑니다",
        "regFailure": "등록이 실패했습니다",
        "success": "성공",
        "failure": "실패하다",
        "unknownErr": "알 수없는 오류",
        "verifyCodeErr": "확인 코드 오류",
        "verifyCodeExpireErr": "확인 코드가 잘못되었거나 만료되었습니다.",
        "thisMailAlreadyExist": "이메일이 등록되었습니다",
        "pageConfigFetchFailure": "구성 획득이 실패하면 새로 고쳐서 다시 시도하십시오",
        "stopRegisterTitle": "등록이 중단되었습니다",
        "stopRegisterHint": "죄송합니다. 등록 기능이 중단되었습니다. 필요한 경우 나중에 다시 시도하거나 지원 팀에 문의하여 자세한 내용은 문의하십시오. 이해와 지원에 감사드립니다.",
        "passwordComplexRequirePart1": "* 비밀번호는 준수해야합니다",
        "passwordComplexRequirePart2": "복잡성 요구 사항",
        "passwordComplexHint1": "1. 3 가지 이상의 큰 글자, 작은 글자, 숫자, 특별 기호가 필요합니다.",
        "passwordComplexHint2": "2. 길이는 10 자리 이상입니다",
        "form": {
            "checkForm": "양식의 내용을 확인하십시오",
            "emailFormatErr": "이메일 주소 형식이 잘못되었습니다",
            "gmailLimitErr": "Gmail 여러 이름이 허용되지 않습니다",
            "gmailDotNotAllowed": "\".\"",
            "gmailPartLowerForced": "Gmail 주소의 로컬 부분은 모든 독자 여야합니다.",
            "googlemailNotAllowed": "Googlemail 주소는 지원되지 않습니다",
            "verifyCodeRequire": "확인 코드를 입력하십시오",
            "verifyCodeFormatErr": "확인 코드 형식이 올바르지 않습니다",
            "passwordRequire": "비밀번호를 입력하십시오",
            "passwordLengthRequire": "비밀번호 길이는 10 자리보다 크거나 같아야합니다.",
            "passwordComplexRequire": "비밀번호에는 적어도 세 가지 유형의 큰 글자, 작은 글자, 숫자 및 특수 기호가 포함되어야합니다.",
            "passwordAgainNotMatch": "암호는 두 번 불일치를 입력했습니다",
            "passwordAgainRequire": "확인 비밀번호를 입력하십시오",
            "inviteCodeRequire": "초대장 코드가 필요합니다"
        },
        "hCaptcha": {
            "passed": "인간-기계 검증이 통과되었습니다",
            "expired": "만료 된 후 다시 시도하십시오",
            "challengeExpired": "검증은 시간을 초과합니다",
            "err": "다른 오류"
        },
        "allRightsReserved": "모든 권리 보유",
        "securityAndLaws": "이 웹 사이트는 Hcaptcha에 의해 보호되고 확인되며 현지 법률의 적용을받습니다."
    },
    "userSummary": {
        "title": "악기",
        "myPlan": "내 구독",
        "shortcut": "바로 가기",
        "timeLeft": "구독은 유효하며 {msg}에서 만료됩니다.",
        "toPurchase": "구독 구매",
        "tutorial": {
            "title": "튜토리얼을 봅니다",
            "content": "{name} 사용 방법 배우기"
        },
        "checkKey": {
            "title": "키를 봅니다",
            "content": "사용하기 위해 클라이언트에 키를 빠르게 소개합니다"
        },
        "renewPlan": {
            "title": "가입을 계속하십시오",
            "content": "현재 구독에 대한 지속적인 결제를하십시오"
        },
        "appDown": {
            "title": "응용 프로그램 다운로드",
            "content": "더 나은 물리적 경험을 위해 다른 플랫폼에서 응용 프로그램을 다운로드하십시오."
        },
        "support": {
            "title": "문제가 발생합니다",
            "content": "문제가 발생하면 작업 순서를 통해 사람들과 의사 소통 할 수 있습니다."
        },
        "haveTicket": "처리 할 {count} 순서가 있습니다",
        "toCheckTicket": "가서 확인하십시오",
        "showAllKeys": "모든 키를보십시오"
    },
    "userDocument": {
        "title": "문서를 사용하십시오",
        "description": "웹 사이트 사용 방법, 예방 조치, 운영 절차 등을 포함하여 여기에서 다른 문서를 확인할 수 있습니다. 기사에 오류가있는 경우 작업 주문을 제출하십시오.",
        "searchPlaceholder": "검색하려는 콘텐츠를 입력하십시오 (퍼지 검색)",
        "searchBtn": "찾다",
        "noContentTitle": "결과가 없습니다",
        "noContentTitleHint": "검색 결과 나 언어와 일치하는 문서는 없으므로 키워드로 변경하십시오."
    },
    "newPurchase": {
        "title": "구독 구매",
        "description": "당신은 여기에서 가장 적합한 구독 계획을 선택할 수 있습니다.",
        "headerPlaceholder": "자신에게 가장 적합한 계획을 선택하십시오",
        "purchaseBtn": "주문하다",
        "noLeft": "남은 양이 충분하지 않습니다",
        "monthPay": "월별 결제",
        "moreMethod": "더 많은 옵션"
    },
    "newSettlement": {
        "err": "실수",
        "monthPay": "월별 결제",
        "quarterPay": "분기 별 지불",
        "halfYearPay": "반년 지불",
        "yearPay": "연간 지불",
        "payCycle": "지불주기",
        "couponPlaceholder": "쿠폰이 있습니까?",
        "verifyBtn": "확인",
        "orderTotalTitle": "총 주문 금액",
        "total": "총",
        "submitOrder": "주문을 제출하십시오",
        "coupon": "쿠폰",
        "notify": {
            "passTitle": "검증이 통과되었습니다",
            "couponVerified": "쿠폰 유효",
            "couponInvalid": "쿠폰이 유효하지 않습니다",
            "couponIsNull": "입력 된 쿠폰 코드는 비어있을 수 없습니다"
        }
    },
    "userProfile": {
        "title": "개인 센터",
        "myWallet": "내 돈 가방",
        "walletSub": "계정 잔액 (소비에만 사용)",
        "alertPwd": "키를 수정하십시오",
        "oldPwd": "오래된 열쇠",
        "oldPwdSub": "이전 비밀번호를 입력하십시오",
        "newPwd": "새로운 키",
        "newPwdSub": "새 키를 입력하십시오",
        "newPwdAgain": "키를 확인하십시오",
        "newPwdAgainSub": "다시 새 키를 입력하십시오",
        "saveBtn": "구하다",
        "notify": "알림",
        "enableNotify": "만료 통지를 활성화합니다",
        "deleteAccount": "등록 계정",
        "deleteAccountSub": "귀하의 계정은 삭제 된 것으로 태그됩니다. 서비스 재사용이 필요한 경우 다시 등록하십시오.",
        "deleteBtn": "내 계정을 등록하십시오",
        "oldPwdVerifiedFailure": "오래된 비밀번호 확인에 실패했습니다",
        "alertFailure": "키를 수정하지 못했습니다",
        "alertSuccess": "수정이 성공적이었습니다",
        "alertSuccessSub": "새 비밀번호로 로그인하십시오",
        "errorPwdFormat": "비밀번호 형식 오류",
        "pwdNotMatch": "비밀번호 입력은 일관성이 없습니다",
        "oldPwdNotNull": "이전 비밀번호는 비어있을 수 없습니다",
        "toTopUp": "재충전으로 가십시오",
        "deleteMyTitle": "경고하다",
        "deleteMyContent": "계정을 삭제 하시겠습니까? 서비스를 사용해야하는 경우 다시 등록하십시오.",
        "deleteMyPositiveText": "삭제를 확인하십시오",
        "deleteMyNegativeText": "취소",
        "deletedSuccessMsg": "계정이 삭제되었으며 나중에 시간이 걸릴 것입니다.",
        "deleteErrOccur": "오류가 발생했습니다",
        "faAuth": "2 단계 검증 2FA",
        "faAuthHint": "2 단계 검증은 계정에 로그인하기위한 보호 수준을 추가하는 보안 메커니즘입니다. 비밀번호를 입력 한 후 사용자는 전화로 전송 된 검증 코드를 입력하거나 Identity Verification Application을 사용하여 동적 코드를 생성하거나 지문과 같은 생물학적 특성을 통해 확인하는 등의 ID 검증의 두 번째 단계를 완료해야합니다.",
        "faAuthStatus": "2 단계 확인 상태",
        "faEnabled": "다시 시작",
        "faNotEnabled": "활성화되지 않았습니다",
        "setup2Fa": "2 단계 검증 설정",
        "disable2Fa": "2 단계 확인을 취소하십시오",
        "unknownErr": "알 수없는 오류",
        "disable2FaCancelled": "취소",
        "closed2FaHint": "2 단계 확인이 닫혔습니다. 필요한 경우 다시 추가하십시오.",
        "setup2FaModal": {
            "followStep": "프롬프트에 따라 검증 장치에 추가하십시오.",
            "step1Title": "아래 단계를 따라 2FA 확인을 활성화하십시오",
            "step1Context1": "1. 모바일 장치에 보편적 인 검증 장치가 있어야합니다.",
            "step1Context2": "2. 확인 장치의 스캔 버튼을 클릭하여 여기에서 QR 코드를 스캔합니다.",
            "step1Context3": "3.이 QR 코드에는 검증 정보와 고유 키가 포함되어 있습니다.",
            "step2Context1": "검증 장치를 정상적으로 사용할 수 있으려면 테스트를 수행해야합니다.",
            "test2Fa": "시험",
            "cancel": "취소"
        },
        "deleteMyAccountModal": {
            "title": "등록 계정",
            "contentLine1": "회계는 돌이킬 수없는 운영입니다. 삭제되었는지 확인하면 계정에 대한 액세스에 대한 액세스가 영구적으로 손실되므로 다시 로그인 할 수 없으며 개인 정보, 기록, 수집 콘텐츠, 구매 기록 등을 포함 하여이 계정과 관련된 모든 데이터는 다시 액세스 할 수 없습니다.",
            "contentLine2": "미완성 주문, 참여중인 활동, 구독 서비스 등과 같은 플랫폼에서 지속적인 비즈니스가있는 경우 계정 제거로 종료되거나 취소되므로 해당 손실이 발생할 수 있습니다. 동시에이 플랫폼을 통해 귀하와 다른 사용자와 다른 사용자에게 연락처 및 상호 작용 정보가 더 이상 존재하지 않습니다.",
            "contentLine3": "결정을 다시 확인하십시오. 질문이나 질문이 있으시면 고객 서비스에 문의하십시오. 진심으로 답변하겠습니다. 여전히 계정을 삭제하려면 확인 버튼 확인 버튼을 클릭하십시오.",
            "inputHint1": "입력하다 \"",
            "inputHint2": "\" 계속하다.",
            "confirmDelete": "삭제를 확인하십시오"
        },
        "failure": "오류가 발생했습니다",
        "notLatestHint": "개인 정보 업데이트가 실패했으며 현재 데이터가 최신 데이터가 아닐 수 있습니다.",
        "resetPassword": {
            "previousPasswordRequire": "이전 비밀번호를 입력하십시오",
            "previousPasswordVerifiedFailure": "오래된 비밀번호 확인에 실패했습니다",
            "passwordRequire": "비밀번호를 입력하십시오",
            "passwordConflict": "새 비밀번호는 이전 비밀번호와 동일 할 수 없습니다.",
            "passwordLengthRequire": "비밀번호 길이는 10 자리보다 크거나 같아야합니다.",
            "passwordComplexRequire": "비밀번호에는 적어도 세 가지 유형의 큰 글자, 작은 글자, 숫자 및 특수 기호가 포함되어야합니다.",
            "passwordAgainNotMatch": "암호는 두 번 불일치를 입력했습니다",
            "passwordAgainRequire": "확인 비밀번호를 입력하십시오",
            "updatePasswordFailure": "비밀번호를 업데이트 할 때 오류가 발생했습니다"
        },
        "secondary": {
            "title": "기본 정보",
            "card": {
                "avatar": "사용자 헤드 이미지",
                "name": "사용자 이름",
                "lastLoginAt": "마지막 로그인 시간",
                "accountStatus": "계정 상태"
            },
            "modify": {
                "uploadIconHint": "업로드",
                "alterAvatar": "헤드 이미지를 추가/수정합니다",
                "alterShallow": "* 추가 또는 수정하려면 헤드 이미지 또는 사용자 이름을 클릭하십시오 (설정 후 청소할 수 없습니다)",
                "alterName": "사용자 이름을 수정하십시오",
                "input": {
                    "label": "사용자 이름",
                    "placeholder": "사용자 이름을 입력하십시오",
                    "spaceIsNotAllowed": "사용자 이름 확인이 전달되지 않습니다",
                    "require": {
                        "p1": "순수한 숫자는 허용되지 않으며 숫자는 열립니다",
                        "p2": "공간이 허용되지 않습니다",
                        "p3": "3보다 길다"
                    }
                }
            },
            "mention": {
                "alterNameSuccess": "사용자 이름 수정이 성공적이었습니다",
                "alterNameErr": "사용자 이름이 실패하면 나중에 다시 시도하십시오",
                "newNameIsNotValid": "불법 사용자 이름",
                "click2SetName": "사용자 이름을 설정하려면 클릭하십시오",
                "fetchAvatarErr": "헤드 이미지 데이터가 실패하면 나중에 다시 시도하십시오.",
                "alterAvatarErr": "헤드 이미지가 실패하면 나중에 다시 시도하십시오",
                "success": "성공",
                "alterAvatarSuccess": "헤드 이미지는 성공적으로 수정되었습니다.",
                "uploadImageHint": "새 헤드 라인으로 이미지를 업로드 할 수 있습니다",
                "imageRequire": {
                    "title": "알아채다",
                    "p1": "*.jpg (jpeg), *.png, *.webp 등과 같은 주류 형식을 업로드 할 수 있습니다.",
                    "p2": "이미지의 길이 비율은 1 : 1 (정사각형)이며, 비율이 다른 경우 중앙의 정사각형으로 절단되고 초과는 삭제됩니다.",
                    "p3": "업로드 한 이미지의 크기는 160px로 설정됩니다."
                },
                "click2Upload": "파일을 클릭하거나 영역으로 드래그하여 업로드하십시오.",
                "uploadWarn": "*은행 카드, 신용 카드, 비밀번호 및 보안 코드와 같은 민감한 데이터를 업로드하지 마십시오."
            }
        }
    },
    "userKeys": {
        "myKeys": "내 열쇠",
        "description": "키에 대한 메모를 설정 해야하는 경우 이벤트 상태, 만료 날짜 등을 볼 수 있습니다.",
        "noKeys": "아직 유효한 구매 기록이 없습니다",
        "keyDetail": "주요 세부 사항",
        "keyId": "키 ID",
        "orderId": "링크 주문 ID",
        "clientId": "클라이언트 ID를 활성화합니다",
        "active": "유효 기간 동안",
        "inActive": "만료되었습니다",
        "valid": "핵심 조건은 정상입니다",
        "invalid": "키가 비활성화되었습니다",
        "isUsed": "사용하도록 활성화",
        "noUsed": "아직 사용되지 않습니다",
        "releaseData": "주요 생성 날짜",
        "expirationData": "만료 날짜",
        "none": "없이",
        "authorizationFor": "주요 인증",
        "hoverClickMention": "클릭 보드에 복사하려면 클릭하십시오",
        "copiedSuccessMessage": "키 사본이 성공적으로 복사되었습니다.",
        "copyFailure": "복사 실패",
        "hoverCopiedSuccessMention": "성공적으로 복사하십시오"
    },
    "userOrders": {
        "myOrders": "내 주문",
        "description": "모든 주문이 여기에 표시되며, 무급 주문이 있으면 지불 계속을 클릭하거나 주문을 취소 할 수 있으며 주문이 완료된 후 주문 세부 정보를 볼 수 있습니다.",
        "orderId": "틀",
        "planName": "구독 이름",
        "planCycle": "가입주기",
        "orderPrice": "주문 금액",
        "orderStatus": "주문 상태",
        "createdAt": "창조 시간",
        "operate": "작동하다",
        "showDetail": "주문 세부 사항",
        "cancelOrder": "주문을 취소하십시오",
        "canceled": "취소",
        "period": {
            "monthPay": "월별 결제",
            "quarterPay": "분기 별 지불",
            "halfYearPay": "반년 지불",
            "yearPay": "연간 지불"
        },
        "orderStatusTags": {
            "success": "성공",
            "cancelled": "실패하다",
            "notPay": "지불되지 않습니다"
        },
        "orderCancelled": "주문 취소",
        "unknownErr": "알 수없는 오류"
    },
    "userTopUp": {
        "topUp": "계정 재충전",
        "description": "여기에서 계정을 재충전하고 사용자 정의 재충전 금액을 지원할 수 있으며 아래에 표시되는 할인 정보가 있는지 여부에주의를 기울이고 언급 된 지불 방법을 사용하여 할인을받을 수도 있습니다.",
        "chooseTopUpAmount": "재충전 금액을 선택하십시오",
        "quickSelect": "빠르게 선택하십시오",
        "customAmount": "맞춤형 금액",
        "maxAmount": "최대 금액 : 10,000,000",
        "amountInputPlaceHolder": "재충전 할 금액을 입력하십시오",
        "otherAmount": "다른 금액",
        "payMethod": "지불 방법",
        "wechat": "wechat 지불",
        "alipay": "지불",
        "apple": "Apple Pay",
        "yourAmount": "금액",
        "discount": "제안",
        "accountBalance": "계정 잔액",
        "balanceResult": "균형의 양",
        "commitTopUp": "제출하다",
        "payMethodNotAllow": "결제 방법을 사용할 수 없습니다. 다른 사람을 선택하십시오",
        "topUpIssueOccur": "재충전 문제가 있습니까?",
        "payIssueOccur": "지불에 어려움이 있습니까?",
        "chatWithUs": "고객 서비스에 문의하십시오",
        "pay": "지불하다",
        "qrCodeScannedSuccess": "QR 코드 스캔이 성공적으로 스캔됩니다",
        "orClickToApp": "또는 막대를 클릭하고 앱으로 돌리고 계속",
        "topUpSuccess": "성공적으로 재충전",
        "thankU": "지원해 주셔서 감사합니다"
    },
    "userConfirmOrder": {
        "switchPlan": "구독하려면 전환하십시오",
        "cancelOrder": "주문을 취소하십시오",
        "yourPrice": "당신의 가격",
        "couponOffset": "쿠폰 할인",
        "price": "가격",
        "submit": "제출하다",
        "monthPay": "월별 결제",
        "quarterPay": "분기 별 지불",
        "halfYearPay": "반년 지불",
        "yearPay": "연간 지불",
        "goodInfo": "제품 정보",
        "cycleOrType": "사이클/유형",
        "orderNumber": "주문 번호",
        "createdAt": "생성 날짜",
        "orderExpired": "주문이 초과되었습니다",
        "paySuccessfully": "지불이 성공적이었습니다. 귀하의 지원에 감사드립니다.",
        "balanceNotEnough": "잔액이 충분하지 않으면 먼저 주문이 5 분 동안 예약되어 있습니다.",
        "orderErrorOccur": "주문하는 동안 오류가 발생했습니다",
        "orderCancelled": "주문 취소"
    },
    "paymentResultParts": {
        "goodInfoView": { "goodInfo": "제품 정보" },
        "orderInfoView": { "orderInfo": "주문 정보" }
    },
    "orderPartUniversal": {
        "period": {
            "monthPay": "월별 결제",
            "quarterPay": "분기 별 지불",
            "halfYearPay": "반년 지불",
            "yearPay": "연간 지불"
        },
        "orderDataHex": {
            "goodInfo": "제품 정보",
            "orderInfo": "주문 정보",
            "cycleOrType": "사이클/유형",
            "orderNumber": "주문 번호",
            "createdAt": "생성 날짜",
            "amount": "지불 금액",
            "paidAt": "지불 시간"
        }
    },
    "orderDetail": {
        "finished": "완전한",
        "finishedAndSuccessDescription": "주문은 성공적으로 지불되고 개설되었습니다",
        "useManual": "튜토리얼을 봅니다",
        "payPending": "아직 지불하지 않았습니다",
        "pendingDescription": "주문은 정기적으로 보관되며 아래 버튼을 클릭하여 지불을 계속할 수 있습니다.",
        "toPay": "지불하러 가십시오",
        "outDate": "만료되었습니다",
        "outDateDescription": "주문을 취소하거나 지정된 시간 내에 지불을 완료하지 않았으므로 주문이 취소되었으며 구독을 다시 선택할 수 있습니다.",
        "chooseNewPlan": "새 구독 계획을 선택하십시오"
    },
    "userInvite": {
        "myInvite": "내 초대",
        "unit": "사람들의 수",
        "inviteCodeMgr": "초대장 코드",
        "generateInviteCode": "임의의 초대장 코드를 생성합니다",
        "faCodeManage": "초대장 코드 관리",
        "email": "이메일",
        "createdAt": "등록 시간",
        "createFaCodeFailure": "창조가 실패했습니다",
        "linkCopiedSuccess": "링크가 성공적으로 복사되었습니다",
        "generateFaCode": "초대장 코드를 생성하십시오",
        "flushFaCode": "초대장 코드를 새로 고치십시오",
        "faCode": "초대장 코드",
        "noFaCode": "아직 초대장 코드가 없습니다.",
        "faLink": "초대장 링크",
        "generateFaCodePlease": "초대장 코드를 작성하십시오",
        "usersMyInvited": "내가 초대하는 사용자",
        "generateCodeConfirm": "생성/새로 고침을 확인합니다",
        "generateCodeHint": "생성 후 초대장 코드를 닫을 수 없습니다.",
        "positiveClick": "확인하다",
        "negativeClick": "취소"
    },
    "userTickets": {
        "description": "사용 문제가 발생하면 기술 지원 및 고객 서비스가 응답하고 아래 양식에 응답 할 수 있습니다.",
        "ticketId": "틀",
        "ticketSubject": "작업 주문 테마",
        "ticketUrgency": "작업 주문 등급",
        "ticketContent": "작업 주문 내용",
        "ticketUrgencyLevel": { "low": "낮은", "med": "가운데", "high": "높은" },
        "ticketStatus": "작업 주문 상태",
        "ticketCreatedAt": "창조 시간",
        "lastResponse": "마지막 답변",
        "operate": "작동하다",
        "checkTicket": "작업 순서를보십시오",
        "closeTicket": "작업 순서를 닫으십시오",
        "userTickets": "역사가 작동합니다",
        "addTicket": "새로운 건축 주문",
        "ticketActive": "닫히지 않습니다",
        "ticketInActive": "닫은",
        "form": {
            "ticketSubjectDescription": "작업 주문 주제를 입력하십시오",
            "ticketUrgencyDescription": "작업 순서의 긴급 성을 선택하십시오",
            "ticketBody": "해결하려는 문제를 가능한 한 포괄적으로 입력하십시오.",
            "ticketNotFinished": "전체 작업 순서에 대한 정보를 사용해보십시오"
        },
        "checkForm": "양식이 완료되었는지 확인하십시오",
        "cancel": "취소",
        "submit": "제출하다",
        "commitNewTicketSuccess": "새로운 작업 명령을 성공적으로 제출하십시오",
        "commitNewTicketFailure": "새 작업 주문 오류를 제출하십시오",
        "noReply": "아직 답장이 없습니다",
        "noTickets": "아직 작업 명령을 제출하지 않았습니다",
        "ticketCloseSuccess": "작업 명령이 성공적으로 폐쇄되었습니다",
        "ticketCloseFailure": "작업 순서 폐쇄가 실패했습니다",
        "chatDialog": {
            "input": {
                "finished": "작업 순서가 끝났습니다",
                "inputHere": "전송 될 메시지를 입력하십시오"
            },
            "send": "보내다",
            "missingToken": "토큰이 누락되어 WebSocket 연결을 만들 수 없습니다.",
            "msgEmptyNotAllowed": "메시지의 내용은 비어있을 수 없습니다",
            "accessNotPermit": "불법 방문",
            "sendSuccess": "메시지를 성공적으로 보냅니다"
        }
    },
    "userActivation": {
        "activateLog": "레코드를 활성화하십시오",
        "description": "여기에서 활성화 레코드를보고 결정되지 않은 장치를 설정하거나 각 키 및 활성화 레코드에 대한 메모 정보를 설정할 수 있습니다.",
        "id": "틀",
        "orderId": "주문 번호",
        "orderStatus": "주문하다",
        "createdAt": "창조 시간",
        "operate": "작동하다",
        "userId": "사용자 ID",
        "email": "이메일",
        "keyId": "키 ID",
        "isBind": "활성화 여부",
        "active": "효율적인",
        "inactive": "효과적인",
        "requestAt": "요청 시간",
        "clientVersion": "클라이언트 측",
        "osType": "운영 체제",
        "remark": "메모",
        "noRemark": "메모가 없습니다",
        "showDetail": "자세한 내용을보십시오",
        "actions": "작동하다",
        "details": "세부",
        "keyContent": "주요 내용",
        "keyGeneratedAt": "주요 생성 시간",
        "activateRequestAt": "요청 시간을 활성화하십시오",
        "useIssueOccur": "그것을 사용하는 데 어려움이 있습니까?",
        "chatWithUs": "저희에게 연락하십시오",
        "cancelBind": "법령을 취소하십시오",
        "alterRemark": "메모를 수정하십시오",
        "commitRemark": "제출하다",
        "updateSuccess": "성공적으로 업데이트하십시오",
        "setRemark": "여기에서 메모 정보를 설정하십시오"
    },
    "userAppDownload": {
        "title": "앱 다운로드",
        "description": "",
        "common": {
            "title": "우리 앱을 다운로드하십시오",
            "shallow2": "다른 고객을위한 응용 프로그램을 받으십시오",
            "shallow": "우리의 응용 프로그램을 사용하면 서비스에보다 쉽게 ​​액세스 할 수있어 매번 브라우저의 복잡한 작동을 제거 할 수 있습니다. 다른 문제에 직면하는 경우 문서에서 자세한 설치 요구 사항 및 튜토리얼을 찾을 수 있습니다.",
            "noDownload": "아직 다운로드 할 수 없습니다. 나중에 다시 시도해보십시오."
        },
        "suffix": {
            "p1": "*MACOS 플랫폼의 응용 프로그램은 MacOS14 이상을 사용하여 Apple Chips (M1 이상)와 협력하십시오.",
            "p2": "이 앱에서 표시되는 정보 및 사용은 앱이 허용되는지 여부에 따라 회사의 IT 사양에 따라야합니다."
        },
        "card": {
            "common": { "welcome": "환영" },
            "mobile": {
                "designFor": "모바일 터미널을위한 설계",
                "designShallow": "여기에서 모바일 앱을 얻을 수 있습니다",
                "iosDownloadShallow": "iOS 클라이언트를 다운로드하십시오",
                "androidDownloadShallow": "Android 클라이언트를 다운로드하십시오"
            },
            "desktop": {
                "designFor": "데스크탑 디자인",
                "designShallow": "여기에서 데스크탑 모바일 애플리케이션을 얻을 수 있습니다",
                "windowsDownloadShallow": "Windows 클라이언트를 다운로드하십시오",
                "osxDownloadShallow": "MacOS 클라이언트를 다운로드하십시오",
                "linuxDownloadShallow": "Linux 클라이언트를 다운로드하십시오"
            }
        },
        "downloadDisabledHint": "죄송합니다. 관리자가 일시적으로 사용할 수 없거나 장애가 있습니다.",
        "windows": {
            "title": "Windows NT",
            "shallow": "이 클라이언트는 NT 커널이있는 Windows 운영 체제에 적합합니다. 호환성 지원은 문서 페이지를 확인하십시오."
        },
        "osx": {
            "title": "마코스 (OSX)",
            "shallow": "이 클라이언트는 Darwin 커널의 MACOS (OSX) 운영 체제에 적합합니다. 호환성 지원은 문서 페이지를 확인하십시오."
        },
        "linux": {
            "title": "리눅스",
            "shallow": "이 클라이언트는 다른 배포 시리즈로 인해 Linux 커널의 다양한 배포에 적합합니다."
        },
        "android": {
            "title": "기계적 인조 인간",
            "shallow": "이 클라이언트는 Google Android 운영 체제가 장착 된 모바일 장치에 적합합니다. 호환성 지원은 문서 페이지를 확인하십시오."
        },
        "ios": {
            "title": "iOS",
            "shallow": "이 클라이언트는 Apple iOS 운영 체제가 장착 된 모바일 장치에 적합합니다. 호환성 지원은 문서 페이지를 확인하십시오."
        }
    },
    "welcome": {
        "A": {
            "aboutUs": "우리에 대해",
            "pricing": "주문 가격",
            "login": "로그인하십시오",
            "register": "등록 계정",
            "welcomeTo": "환영",
            "welcomeToSub": "카운티의 긴 터널을 통과하는 것은 눈 국가입니다. 밤하늘 아래에서 지구는 흰색이었고 기차는 신호 스테이션 앞에서 멈췄습니다. \"여기, Kawabata Yasunari는 매우 인색한 간결한 텍스트로 눈 국가를 시작했습니다.",
            "startingUse": "시작하세요",
            "whyUs": "왜 우리를 선택하십시오",
            "whyUsSub": "\"현의 긴 터널을 통과 할 때, 그것은 눈의 나라입니다. 밤하늘 아래에서 지구는 흰색이며, 기차는 신호 스테이션 앞에서 멈 춥니 다.\"",
            "browseSafe": "안전하게 찾아보십시오",
            "browseSafeSub": "우수한 방화벽 필터 시스템은 온라인 물고기 및 악의적 인 웹 사이트를 효과적으로 방지 할 수 있습니다.",
            "encrypt": "엔드 투 엔드 암호화",
            "encryptSub": "이중 방향 SSL 및 엔드 투 엔드 암호화 개인 정보 보호 및 보안을 보호하십시오. 서버조차도 정보를 읽을 수 없습니다.",
            "mgr": "효율적인 관리",
            "mgrSub": "한 사용자 인터페이스는 완전하고 풍부한 관리 기능으로 모든 키를 관리하며 사정 문제를 구독하는 것에 대해 걱정할 필요가 없습니다.",
            "fast": "편리하고 빠릅니다",
            "fastSub": "WebApps 또는 타사 소프트웨어에 포함 된 완전한 API 파일 제공",
            "fastLink": "빠른 링크",
            "subscribeUs": "우리를 따르십시오",
            "filingsCode": "파일 번호 {code}"
        }
    },
    "pagination": {
        "perPage10": "10 데이터/페이지",
        "perPage20": "20 데이터/페이지",
        "perPage50": "50 데이터/페이지",
        "perPage100": "100 데이터/페이지"
    },
    "modal": { "cancel": "취소", "confirm": "확인하다" },
    "week": {
        "monday": "월요일",
        "tuesday": "화요일",
        "wednesday": "수요일",
        "thursday": "목요일",
        "friday": "금요일",
        "saturday": "토요일",
        "sunday": "일요일"
    },
    "period": {
        "month": "월별 결제",
        "quarter": "분기 별 지불",
        "halfYear": "반년 지불",
        "year": "연간 지불"
    },
    "operate": {
        "cancel": "취소",
        "confirm": "확인하다",
        "update": "고쳐 쓰다",
        "add": "추가"
    },
    "notFound": {
        "title": "404 페이지가 존재하지 않습니다",
        "description": "요청한 페이지를 찾을 수 없으며 삭제되거나 오류와 연결되었을 수 있습니다. 이것이 오류라고 생각되면 저희에게 연락 할 작업 주문을 제출하십시오.",
        "p1": "{sec} s 이후 홈페이지로 돌아가고 브라우저가 응답하지 않으면 아래 버튼을 클릭하십시오.",
        "manualBack": "홈페이지로 돌아갑니다"
    },
    "forbidden": {
        "title": "403 권리 없음",
        "description": "이 페이지에 액세스 할 수있는 충분한 권한이 없을 수도 있습니다. 이것이 오류라고 생각되면 저희에게 연락 할 작업 주문을 제출하십시오."
    }
}
