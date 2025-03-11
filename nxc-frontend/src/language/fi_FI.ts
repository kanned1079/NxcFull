export default {
    "commonHeader": {
        "menuTxt": "valikko",
        "userData": "Käyttäjätiedot",
        "editUserData": "Muokkaa käyttäjätietoja",
        "logout": "Kirjautua sisään"
    },
    "commonAside": {
        "admin": {
            "dashboard": "Pöytälauta",
            "queueMonit": "Huoltopuolen seuranta",
            "settings": "Asetukset",
            "systemConfig": "Järjestelmäasetukset",
            "paymentConfig": "Maksuasetukset",
            "themeConfig": "Teema -asetukset",
            "server": "Palvelin",
            "privilege": "Viranomaisen hallinta",
            "finance": "Rahoitus",
            "subscription": "Tilauksen hallinta",
            "coupon": "Kuponginhallinta",
            "order": "Tilauksen hallinta",
            "activate": "Aktivoi tietue",
            "key": "Avainhallinta",
            "user": "Käyttäjä",
            "userMgr": "Käyttäjän hallinta",
            "notice": "Ilmoituksen hallinta",
            "ticket": "Työjärjestyksen hallinta",
            "doc": "Tietotietokannan hallinta"
        },
        "user": {
            "dashboard": "Pöytälauta",
            "document": "Käytä asiakirjaa",
            "app": "Sovelluslataus",
            "subscription": "Tilata",
            "purchase": "Ostaa tilaus",
            "surplus": "Minun avain",
            "fiance": "Rahoitus",
            "topUp": "täydentää",
            "myOrder": "Minun tilaukseni",
            "myInvite": "Kutsuni",
            "user": "Käyttäjä",
            "profile": "Henkilökohtainen keskus",
            "support": "Työjärjestykseni",
            "activateLog": "Aktivoi tietue"
        }
    },
    "adminViews": {
        "common": {
            "fetchDataSuccess": "Tietojen hankkiminen onnistuneesti",
            "fetchDataFailure": "Jos tiedot epäonnistuvat, yritä uudelleen myöhemmin",
            "addSuccess": "Lisätty onnistuneesti",
            "addFailure": "Jos lisäys epäonnistuu, kokeile myöhemmin",
            "updateSuccess": "Muutos oli onnistunut",
            "updateFailure": "Jos muutos epäonnistuu, yritä uudelleen myöhemmin",
            "deleteSuccess": "Poisto onnistuneesti",
            "deleteFailure": "Yritä uudelleen, jos poisto epäonnistui",
            "confirm": "Vahvistaa",
            "cancel": "Peruuttaa",
            "success": "Toiminta on onnistunut",
            "failure": "Toiminta epäonnistui",
            "NotAllowed": "Laiton muoto",
            "checkForm": "Tarkista lomake",
            "unknownErr": "Tuntematon virhe",
            "dialog": { "delete": "Tiedätkö, että haluat poistaa sen?" },
            "yes": "kyllä",
            "no": "ei"
        },
        "login": {
            "secureCard": {
                "title": "Turvatarkastus",
                "securePath": "Turvallinen polku",
                "hint": "Järjestelmän suojauksen varmistamiseksi sinun on syötettävä suojapolku ennen kuin voit kirjoittaa järjestelmänvalvojan kirjautumissivulle. Syötä tietoturvapolku alla olevaan syöttöruutuun. Suojauspolun onnistuneesti voit valita tallentaa seuraavaa nopeaa kirjautumista varten.",
                "placeholder": "Anna turvallinen polku",
                "checkBtn": "tutkia",
                "rememberPath": "Muista turvallinen polku"
            },
            "card": {
                "back": "Palaa kotisivulle",
                "toAdminCenter": "Kirjaudu sisään järjestelmänvalvojan keskukseen",
                "emailAddr": "Järjestelmänvalvojan sähköposti",
                "emailInputArea": {
                    "title": "Järjestelmänvalvojan sähköposti",
                    "placeholder": "Sähköpostiosoite"
                },
                "passwordInputArea": {
                    "title": "Järjestelmänvalvojan salasana",
                    "placeholder": "salasana"
                },
                "login": "sisäänkirjautuminen",
                "forgetPassword": "unohda salasana",
                "formNotPassed": "Lomakkeen varmennus epäonnistuu"
            },
            "mention": {
                "title": "vihje",
                "description": "Tämä sivu on järjestelmänvalvojan sivu, johon pääsee vain järjestelmänvalvojat."
            },
            "message": {
                "passwordErr": "Väärä salasana",
                "adminNotExist": "Järjestelmänvalvojaa ei ole olemassa",
                "noPrivilege": "Pääsy ilman lupaa",
                "authPassed": "Vahvistus ohi",
                "authFailure": "Vahvistus epäonnistui",
                "otherErr": "Muut virheet",
                "pathCheckPassed": "Turvapolun tarkistus kulki",
                "pathCheckFailure": "Virheellinen turvallisuuspolku",
                "rememberSecureMention": "Älä tarkista, että tämä ei ole yksityinen tietokoneesi varmistaaksesi taustaohjelman hallinnan turvallisuuden."
            }
        },
        "summary": {
            "cockpit": "Pöytälauta",
            "systemConfig": "Järjestelmäasetukset",
            "paymentConfig": "Maksuasetukset",
            "planMgr": "Tilauksen hallinta",
            "userMgr": "Käyttäjän hallinta",
            "orderMgr": "Tilauksen hallinta",
            "keyMgr": "Avainhallinta",
            "incomeText": "Eilinen tulo / kuukausitulot",
            "pendingTicket": "Sinulla on käsiteltävä {nums} tilauksia",
            "toHandle": "Mene ja tarkista se",
            "apiAccessCard": "API -käyttöliittymän lukumäärä viikossa viikossa",
            "apiAccessCardHint": "Nämä tiedot on tarkoitettu vain ymmärtämään nykyisen sovellusliittymän käyttöoikeus, eivätkä edusta palvelimen suorituskykyäsi.",
            "incomeWeek": "Tulomäärä viikon kuluessa",
            "incomeCardHint": "Tässä on tulosmäärät yli viikon, mikä aiheuttaa epätarkkoja näytön, jos välimuisti selvitetään.",
            "core": "ydin",
            "reqErr": "Kohtasi virheen",
            "reqErrHint": "Yleiskatsaustiedot saatiin virhe, joka aiheutti tämän pyynnön kyvyttömyyden loppuun, joten kaaviota ei voida näyttää.",
            "userCard": {
                "title": "Käyttäjän yleiskatsaus",
                "allRegisteredUsers": "Yleinen rekisteröinnin käyttäjä",
                "activeUsers": "Live -käyttäjät",
                "inactiveUsers": "Passiiviset käyttäjät",
                "blockedUsers": "Kielto tai korvaus"
            },
            "general": {
                "title": "yleensä",
                "localTime": "Selaimen aika",
                "osType": "Käyttöjärjestelmätyyppi",
                "appName": "Sovelluksen nimi",
                "appUrl": "Sovellus -URL -osoite",
                "currency": "Kolikkoyksikkö",
                "allowRegister": "Salli rekisteröinti"
            },
            "system": {
                "title": "Järjestelmän kokoonpano",
                "axiosAddr": "HTTP -taustaosoite",
                "wsAddr": "WebSocket -taustaosoite",
                "serverTime": "Palvelinaika",
                "uptime": "Operatiivinen aika",
                "gatewayStatus": "API -verkon tila",
                "dbStatus": "Tietokannan tila",
                "redisStatus": "Redis -tila",
                "serverOsType": "Palvelimen käyttöjärjestelmätyyppi",
                "serverOsArch": "Palvelimen käyttöjärjestelmän arkkitehtuuri",
                "runMode": "Käyttötila",
                "cpuNums": "Verkkopalvelimien prosessorin ytimien lukumäärä",
                "numCgoCall": "Jätehuoltokokoelmien lukumäärä",
                "time": "Toisen luokan",
                "paymentMethods": "Ota maksutapa käyttöön",
                "runOK": "Juoksua normaalisti",
                "runErr": "Epänormaali käyttäytyminen",
                "checkServer": "Tarkista taustapalvelimen ympäristömääritykset",
                "stopRegisterHint": "Näyttää siltä, ​​että olet poistanut käytöstä uuden käyttäjän rekisteröinnin",
                "toSetting": "Kääntyä asetuksiin"
            }
        },
        "queueMonit": {
            "title": "Huoltopuolen seuranta",
            "headerHint": "Älä pysy tällä sivulla pitkään.",
            "latency": {
                "title": "Palvelinviive",
                "retry": "Testata uudelleen",
                "hint": "*Pyynnöstä viitataan tässä ajassa, joka vaaditaan onnistuneelle vastaukselle asiakkaalle sen jälkeen, kun asiakas lähettää pyynnön palvelimelle.",
                "unit": "millisekunti",
                "level": {
                    "l1": {
                        "title": "Erinomainen",
                        "description": "Tämä on erittäin hyvä verkkotilanne, ja tuskin tuntuu viivästyneeltä."
                    },
                    "l2": {
                        "title": "normaali",
                        "description": "Tämä on verkkotilanne suurimman osan ajasta, ja käyttäjät melkein tuntevat, että he eivät tunne viivästyneitä."
                    },
                    "l3": {
                        "title": "Huono",
                        "description": "Tässä tapauksessa käyttäjä voi tuntea pienen viiveen tai viivästyksen."
                    },
                    "l4": {
                        "title": "Ero",
                        "description": "Voit tuntea ilmeisen viiveen ja vaikuttaa käyttökokemukseen."
                    },
                    "l5": {
                        "title": "Erittäin huono",
                        "description": "Viive on ilmeinen ja lastausnopeus on hitaampi tai edes kykene päivittämään, mikä vaikuttaa vakavasti käyttäjän vuorovaikutukseen ja kokemukseen."
                    },
                    "offline": {
                        "title": "Laittaa",
                        "description": "Tarkista, onko kokoonpano oikeassa palvelimessa tai käsittelemään pyyntövirheitä."
                    }
                }
            },
            "api": {
                "title": "API -pyynnön tila viimeisen 7 päivän aikana",
                "items": {
                    "ok": { "title": "Menestys (StatusOk)", "unit": "Toisen luokan" },
                    "notFound": {
                        "title": "Tilapolkuvirhe (404)",
                        "unit": "Toisen luokan"
                    },
                    "unAuthorized": {
                        "title": "Luvaton pääsy (401)",
                        "unit": "Toisen luokan"
                    },
                    "login2reg": {
                        "title": "Kirjaudu sisään / rekisteröidy",
                        "unit": "Toisen luokan"
                    }
                }
            },
            "log": {
                "title": "Päiväkirjarekisteri",
                "deleteLogMsg": "Poista päiväkirja {nums}",
                "deleteLogErr": "Poista päiväkirja epäonnistui",
                "logTableRows": "Päiväkirjarekisterit",
                "logTableSize": "Blogipöydät miehittävät tilaa",
                "unit": { "lines": "Hyvä", "mb": "Megakarakki" },
                "deleteLog": "Poista päiväkirja",
                "exportCsv": "Vienti CSV",
                "deleteLogHint": "Tämä toteuttaa kaikkien lehtien poistaminen viikko sitten.",
                "warn": {
                    "title": "Tiedätkö, että haluat poistaa päiväkirjasi?",
                    "description": "Luettelo poistetaan välittömästi, etkä voi peruuttaa tätä operaatiota."
                },
                "export": {
                    "title": "Vie CSV -tiedosto",
                    "description": "Tämä vie seuraavan taulukon CSV -tiedostoksi ja lataa sen paikallisesti. Jos sinulla ei ole lupaa avata latausta, lataus voi epäonnistua."
                },
                "table": {
                    "id": "Hio",
                    "method": "Pyyntömenetelmä",
                    "path": "Pyydä reittiä",
                    "code": "Tilakoodi",
                    "clientIp": "Asiakkaan IP",
                    "responseTime": "Käsittelyaika",
                    "requestAt": "Pyyntöaika"
                }
            }
        },
        "systemConfig": {
            "title": "Järjestelmäasetukset",
            "common": {
                "success": "Päivitetty kokoonpano onnistuneesti",
                "err": "Päivitysmääritys epäonnistui"
            },
            "site": {
                "common": { "title": "Paikka" },
                "appName": {
                    "title": "Sivuston nimi",
                    "shallow": "Käytetään näyttämään, missä sivuston nimi vaaditaan.",
                    "placeholder": "Sivuston nimi"
                },
                "appSubName": {
                    "title": "Sivuston tekstitys",
                    "shallow": "Yleensä pääotsikon alla.",
                    "placeholder": "tekstitys"
                },
                "appDescription": {
                    "title": "Sivuston kuvaus",
                    "shallow": "Käytetään näyttämään, missä sivusto vaaditaan.",
                    "placeholder": "Sivuston kuvaus"
                },
                "appUrl": {
                    "title": "Sivuston URL",
                    "shallow": "Viimeisin verkkosivustoosoite näkyy URL -osoitteessa, jos tarvitset sitä.",
                    "placeholder": "Sivuston URL"
                },
                "forceHttps": {
                    "title": "Force Https",
                    "shallow": "Kun sivusto ei käytä HTTP: tä, CDN: tä tai sukupolven vastaista pakotettujen HTTP: ien mahdollistamiseksi."
                },
                "logoUrl": {
                    "title": "LOGO",
                    "shallow": "Käytetään näyttämään, missä logo vaaditaan.",
                    "placeholder": "Logon URL -osoite"
                },
                "subscribeUrl": {
                    "title": "Tilaa URL -osoite",
                    "shallow": "Käytetään tilaamiseen, jätä tyhjä sivuston URL -osoitteeksi. Jos tarvitset useita tilaus -URL -osoitteita satunnaisesti hakeaksesi, jaa pilkkuja.",
                    "placeholder": "Tilaa URL -osoite"
                },
                "tosUrl": {
                    "title": "Käyttäjän ehdot (TOS)",
                    "shallow": "Käytetään hyppäämään käyttäjän termeihin (TOS)",
                    "placeholder": "Käyttäjän ehdot Osoite"
                },
                "stopRegister": {
                    "title": "Lopeta uusi käyttäjän rekisteröinti",
                    "shallow": "Kukaan ei voi rekisteröidä sen jälkeen, kun se on kytketty päälle."
                },
                "inviteRequire": {
                    "title": "Pakkokutsu",
                    "shallow": "Avaamisen jälkeen sinun on täytettävä kutsukoodi rekisteröidyt uuteen käyttäjälle."
                },
                "trialSubscribe": {
                    "title": "Rekisteröidy oikeudenkäyntiin",
                    "shallow": "Valitse Tilaus, jonka haluat kokeilla."
                },
                "trialTime": {
                    "title": "Koeaika (tuntia)",
                    "shallow": "Tilaa kokeiluaika, kun uudet käyttäjät kirjautuvat sisään."
                },
                "currency": {
                    "title": "Valuuttayksikkö",
                    "shallow": "Sitä käytetään vain näytön käyttöön, ja muutoksen jälkeen kaikki järjestelmän valuuttayksiköt muuttuvat.",
                    "placeholder": "CNY"
                },
                "currencySymbol": {
                    "title": "Valuutan symboli",
                    "shallow": "Sitä käytetään vain näytön käyttöön, ja muutoksen jälkeen kaikki järjestelmän valuuttayksiköt muuttuvat.",
                    "placeholder": "¥"
                }
            },
            "security": {
                "common": { "title": "Suojausasetukset" },
                "emailVerify": {
                    "title": "Sähköpostivahvistus",
                    "description": "Avaamisen jälkeen käyttäjän on suoritettava sähköpostin varmennus."
                },
                "gmailAlias": {
                    "title": "Kielletä Gmail -useita aliaksia",
                    "description": "Avaamisen jälkeen useita Gmail -aliasia ei rekisteröidä."
                },
                "safeMode": {
                    "title": "Turvatila",
                    "description": "Avaamisen jälkeen verkkotunnuksen pääsy tähän sivustoon paitsi sivuston URL -osoite on 403."
                },
                "adminPath": {
                    "title": "Taustatie",
                    "description": "Taustahallintapolku muuttaa alkuperäistä järjestelmänvalvojan polkua muutoksen jälkeen.",
                    "placeholder": "https://x.com/logo.jpeg"
                },
                "emailWhitelist": {
                    "title": "Sähköposti jälkiliitteen sallintaluettelo",
                    "description": "Rekisteröinti on sallittua vasta sen jälkeen, kun se on avannut postilaatikon jälkiliitteen luettelossa."
                },
                "recaptcha": {
                    "title": "Anti-robotti",
                    "description": "Kun HCaptcha on käytössä, se otetaan käyttöön robotin estämiseksi."
                },
                "hCaptchaSiteKey": {
                    "title": "HCaptcha SiteKey",
                    "description": "SiteKey: tä käytetään pyytämään HCaptcha -palvelinta verkkosivuston numeron tunnistamiseen",
                    "placeholder": "A3CA066C-0EA0-42FE-BCD2-23F4AB48D528"
                },
                "ipRegisterLimit": {
                    "title": "IP -rekisteröintirajoitukset",
                    "description": "Avaamisen jälkeen, jos IP-rekisteröintitili täyttää sääntöjen vaatimukset, rekisteröinti on rajoitettu."
                },
                "registerTimes": {
                    "title": "taajuus",
                    "description": "Rangaistus aktivoituu rekisteröintien lukumäärän saavuttamisen jälkeen.",
                    "placeholder": "Ole hyvä ja kirjoita"
                },
                "lockTime": {
                    "title": "Rangaistusaika (minuutteja)",
                    "description": "Sinun on odotettava, kunnes rangaistusaika on kulunut ennen kuin voit rekisteröityä uudelleen.",
                    "placeholder": "Ole hyvä ja kirjoita"
                }
            },
            "frontend": {
                "common": { "title": "Persoonallisuus" },
                "themePropColor": {
                    "default": "laiminlyönti",
                    "darkBlueDay": "Tummansininen",
                    "milkGreenDay": "Maitovihreä",
                    "bambooGreen": "Ruzhu",
                    "mistyPine": "Mänty",
                    "glacierBlue": "Jäätikön sininen",
                    "grayTheme": "harmaa"
                },
                "sidebarStyle": {
                    "title": "Sivupalkkityyli",
                    "shallow": "Aseta sivupalkin väri."
                },
                "headerStyle": {
                    "title": "Päätyyli",
                    "shallow": "Aseta väri yläosaan."
                },
                "themeColor": {
                    "title": "Teemaväri",
                    "shallow": "Aseta koko verkkosivun teeman väri."
                },
                "background": {
                    "title": "tausta",
                    "shallow": "Se näytetään taustan kirjautumissivulla.",
                    "placeholder": "https://x.com/logo.jpeg"
                }
            },
            "inviteAndRebate": {
                "common": { "title": "Maksu ja alennus" },
                "inviteRebateEnable": {
                    "title": "Ota käyttäjän alennukset käyttöön",
                    "description": "Jos kutsuttu käyttäjä latautuu käytössä, käyttäjä palautetaan alla olevan ladattavan suhteen mukaisesti."
                },
                "inviteRebateRate": {
                    "title": "Alennussuhde",
                    "description": "Aseta alennuksen määrän suhde.",
                    "placeholder": "Anna alennussuhde"
                },
                "discountInfo": {
                    "title": "Alennustiedot",
                    "description": "Aseta tarjoustiedot ja ne näytetään ylhäällä olevan sivun yläosassa.",
                    "placeholder": "Aseta alennustiedot"
                },
                "inviteInfo": {
                    "title": "Kutsutiedot",
                    "description": "Aseta kutsutiedot, ne näytetään käyttäjän kutsusivulla alennussuhteen näyttämiseksi.",
                    "placeholder": "Aseta alennustiedot"
                }
            },
            "welcome": {
                "common": { "title": "Kotisivun tiedot" },
                "homeDescription": {
                    "title": "Kotisivun kuvaus",
                    "description": "Aseta lyhyt kuvaus kotisivusta.",
                    "placeholder": "Kirjoita kotisivun kuvaus sisältö"
                },
                "whyChooseUs": {
                    "title": "Miksi valita meidät",
                    "description": "Määritä kuvauksemme siitä, miksi valitsimme.",
                    "placeholder": "Anna yksityiskohtainen kuvaus"
                },
                "bilibiliLink": {
                    "title": "Bilibili -virallinen linkki",
                    "description": "Aseta Bilibili -virallisen tilin linkin osoite.",
                    "placeholder": "https://space.bilibili.com/xxxx"
                },
                "youtubeLink": {
                    "title": "YouTube -virallinen linkki",
                    "description": "Aseta virallisen YouTube -tilin linkin osoite.",
                    "placeholder": "https://youtube.com/channel/xxxxx"
                },
                "instagramLink": {
                    "title": "Instagram -virallinen linkki",
                    "description": "Aseta virallisen Instagram -tilisi linkin osoite.",
                    "placeholder": "https://instagram.com/xxxx"
                },
                "wechatLink": {
                    "title": "WeChat Public Account Link",
                    "description": "Aseta WeChat -julkisen tilin linkin osoite.",
                    "placeholder": "Anna WeChat -julkinen linkki"
                },
                "filingNumber": {
                    "title": "Rekisterinumero",
                    "description": "Aseta sivuston rekisteröintinumero.",
                    "placeholder": "Esimerkiksi: Guangdong ICP nro 12345678"
                },
                "pageSuffix": {
                    "title": "Sivuston jälki",
                    "description": "Aseta sivuston nimi.",
                    "placeholder": "Esimerkiksi:- Sivustosi nimi"
                }
            },
            "sendMail": {
                "common": {
                    "title": "Postiasetukset",
                    "warning": "Jos muutat tämän sivun kokoonpanoa, sinun on käynnistettävä jonopalvelu ja käänteinen suunta. Lisäksi tämän sivun kokoonpanoprioriteetti on korkeampi kuin .Env -postin kokoonpano.",
                    "sendTestMailTitle": "Lähetä testiviesti",
                    "sendTestMailShallow": "Sähköposti lähetetään nykyisen kirjautuneen järjestelmänvalvojan sähköpostiosoitteeseen",
                    "sendTestMailTo": "Lähetä testiposti osoitteeseen",
                    "sending": "Lähettää sähköposti",
                    "success": "menestys",
                    "receiverAddr": "Vastaanotettu osoite",
                    "senderHost": "Toimituspalvelin",
                    "senderPort": "Lähettää satama",
                    "senderEncrypt": "Salausmenetelmä viestien lähettämiseen",
                    "senderUsername": "Luottotunnus",
                    "sendErr": "Lähetä sähköpostin epäonnistuminen"
                },
                "smtpServerAddress": {
                    "title": "SMTP -palvelimen osoite",
                    "shallow": "Sähköpostipalveluntarjoajan toimittama palveluosoite",
                    "placeholder": "Anna palvelimen osoite"
                },
                "smtpServerPort": {
                    "title": "SMTP -palveluportti",
                    "shallow": "Yleiset satamat ovat 25, 465, 587",
                    "placeholder": "Anna portinumero"
                },
                "smtpEncryption": {
                    "title": "SMTP -salausmenetelmä",
                    "shallow": "465 Portin salausmenetelmä on yleensä SSL, ja 587 portin salausmenetelmä on yleensä TLS.",
                    "placeholder": "Anna salausmenetelmä"
                },
                "smtpAccount": {
                    "title": "SMTP -tili",
                    "shallow": "Sähköpostipalveluntarjoajan toimittama tili",
                    "placeholder": "Anna tilinumero"
                },
                "smtpPassword": {
                    "title": "SMTP -salasana",
                    "shallow": "Sähköpostipalveluntarjoaja toimittaa salasana",
                    "placeholder": "Anna salasanasi"
                },
                "senderAddress": {
                    "title": "Lähettää osoite",
                    "shallow": "Postipalveluntarjoajan toimittama lähetysosoite",
                    "placeholder": "Anna lähetysosoite"
                },
                "emailTemplate": {
                    "title": "Sähköpostipohja",
                    "shallow": "Voit tarkastella, miten asiakirjan sähköpostimallia voidaan mukauttaa",
                    "placeholder": "Valitse sähköpostimalli"
                }
            },
            "notice": {
                "common": { "title": "Ilmoitusasetukset" },
                "displayName": {
                    "title": "näytönimi",
                    "shallow": "Vain etusivun näytölle",
                    "placeholder": "Yleinen ilmoitusrajapinta 1"
                },
                "barkEndpoint": {
                    "title": "Kuoren tukiasema",
                    "shallow": "Bark Server Bacend API -osoite",
                    "placeholder": "https: // <IP>: <port>/<cript-Key>"
                },
                "barkGroup": {
                    "title": "Haukaryhmä",
                    "shallow": "Ryhmän nimi, jonka asiakkaan näyttää",
                    "placeholder": "verkko"
                }
            },
            "appDownload": {
                "common": {
                    "title": "Soveltaa",
                    "hint": "Omien asiakkaiden versionhallinta- ja päivitys (sovellukset)"
                },
                "enabled": {
                    "title": "Avaa ja pudota kuorma",
                    "shallow": "Jos käytössä on, anna käyttäjien käyttää lataussivua"
                },
                "windows": {
                    "title": "Ikkunat",
                    "shallow": "Windows -versionumero ja lataa osoite",
                    "placeholder": "https://xxxx.com/xxx.exe"
                },
                "macos": {
                    "title": "macos",
                    "shallow": "MacOS -versionumero ja lataa osoite",
                    "placeholder": "https://xxxx.com/xxx.dmg"
                },
                "linux": {
                    "title": "Linux",
                    "shallow": "Linux -versionumero ja lataa osoite",
                    "placeholder": "https://xxxx.com/xxx.deb"
                },
                "android": {
                    "title": "Androidi",
                    "shallow": "Android -versionumero ja lataa osoite",
                    "placeholder": "https://xxxx.com/xxx.apk"
                },
                "ios": {
                    "title": "IOS",
                    "shallow": "IOS -versionumero ja lataa osoite",
                    "placeholder": "https://xxxx.com/xxx.ipk"
                }
            }
        },
        "payConfig": {
            "title": "Maksuasetukset",
            "description": "Kaikkia tuettuja maksutapoja voidaan hallita. Tällä hetkellä vain maksutapa on tuettu, mutta voit ensin määrittää muut maksutavat.",
            "attention": {
                "title": "Asiat huomataan",
                "point1": "On todella tärkeää määrittää maksutapatiedot ennen sen käyttöönottoa.",
                "point2": "Kun muokkaat maksutavan kokoonpanoa, jos se näytetään \"---\", se tarkoittaa, että vaihtoehto on asetettu eikä ole tyhjä."
            },
            "common": {
                "detail": "{menetelmä} kokoonpano",
                "fillAttention": "Turvallisuuden varmistamiseksi ei näytetä yksityiskohtaisia ​​tietoja, jotka täyttyvät olemassa olevien kokoonpanojen luomiseksi tai korvaamiseksi.",
                "discountAmount": "Alennussumma (käyttäjän kehotetiedot voidaan asettaa järjestelmän asettamassa \"maksu ja alennus\")",
                "saveConfigBtnHint": "tallentaa",
                "cancelBtnHint": "Peruuttaa",
                "saveSuccess": "Kokoonpano tallennetaan onnistuneesti",
                "alterSuccess": "Konfiguraatiomuutos oli onnistunut",
                "discountPlaceholder": "Alennusmäärä (jos latausmäärä on suurempi kuin alennusmäärä)",
                "saveOrAlterFailure": "Tuntematon virhe"
            },
            "alipay": {
                "title": "Maksu",
                "config": {
                    "appId": "Sovellustunnus",
                    "appPublicKeyCertContent": "Sovellusavaimen varmenteen sisältö",
                    "appPrivateKey": "Sovelluksen tietosuoja -avain",
                    "alipayRootCert": "Maksupuskikirja",
                    "alipayPublicKeyCertContent": "Maksun aartetodistuksen sisältö"
                }
            },
            "wechat": {
                "title": "WeChat -maksu",
                "config": {
                    "mchId": "Kauppiastunnus",
                    "mchCertSerial": "Kauppiasvarmenteen sarjanumero",
                    "apiV3Key": "API V3 -näppäin",
                    "privateKey": "Tietosuoja -avain"
                }
            },
            "apple": {
                "title": "Apple Pay",
                "config": {
                    "issuerId": "Liikkeeseenlaskijan tunnus",
                    "bundleId": "Nipputunnus",
                    "privateKeyId": "Yksityinen avaintunnus",
                    "privateKeyContent": "Yksityisyydensisällön sisältö"
                }
            },
            "addPaymentMethod": "Lisää maksutapa",
            "enableBtnHint": "Alkaa",
            "disableBtnHint": "Vammainen",
            "setupPaymentMethod": "Kokoonpano"
        },
        "themeConfig": {
            "title": "Aiheen kokoonpano",
            "using": "käyttää,",
            "setAsCurrent": ""
        },
        "groupMgr": {
            "title": "Viranomaisen hallinta",
            "description": "Oikeaa ryhmää käytetään erilaisten tilaustasojen merkitsemiseen, ja voit tilata saman tason, mutta erikokoiset oikeat ryhmät helpon hallinnan saavuttamiseksi.",
            "common": {
                "addNewGroup": "Uusi luparyhmä",
                "alterGroupName": "Muokkaa luparyhmän nimeä",
                "addConfirmed": "Vahvista lisää",
                "alterConfirmed": "Vahvista muutos",
                "cancel": "Peruuttaa",
                "addSuccess": "Lisätty onnistuneesti",
                "addFailure": "Lisätä epäonnistumisia",
                "alterSuccess": "Muutosluparyhmä oli onnistunut",
                "alterFailure": "Muutosluparyhmä epäonnistui",
                "delSuccess": "Poisto onnistuneesti",
                "delFailure": "Poisto epäonnistui",
                "delMention": "Luettelo poistetaan välittömästi, etkä voi peruuttaa tätä operaatiota. Tilaussuunnitelmien asiaankuuluvat ryhmät olisi poistettava varoen.",
                "delNotAllowed": "Tällä oikealla ryhmällä on liittyviä resursseja, eikä sitä voida poistaa."
            },
            "groupId": "Oikeusryhmän henkilöllisyystodistus",
            "groupName": "Luvan nimi",
            "groupPlaceholder": "Syötä luparyhmän nimi",
            "userCount": "Käyttäjien lukumäärä",
            "planCount": "Tilausmäärä",
            "operate": "toimia",
            "editBtnHint": "Muokata",
            "deleteBtnHint": "poistaa"
        },
        "docMgr": {
            "title": "Tietotietokannan hallinta",
            "description": "Täällä voit kirjoittaa kuvaustiedoston käyttäjällesi.",
            "addDoc": "Lisää uusi asiakirja",
            "addSuccess": "Uuden asiakirjan lisääminen onnistuneesti",
            "addFailure": "Tiedoston lisääminen epäonnistui",
            "titleNotEmpty": "Asiakirjan otsikko ei voi olla tyhjä",
            "table": {
                "docId": "Hio",
                "isShow": "Näyttääkö",
                "sortAs": "Järjestellä",
                "lang": "Kieli",
                "category": "Luokka",
                "title": "Otsikko",
                "createdAt": "Luomisaika",
                "updatedAt": "Päivitysaika",
                "operate": "toimia",
                "edit": "Muokata",
                "delete": "poistaa"
            },
            "form": {
                "add": "Lisää asiakirja",
                "edit": "Muokkaa asiakirjaa",
                "cancel": "Peruuttaa",
                "confirm": "Vahvistaa",
                "addBtn": "Lisätä jtk",
                "editBtn": "Tarkistaa",
                "title": {
                    "title": "Tiedoston otsikko",
                    "placeholder": "Kirjoita tiedoston otsikko"
                },
                "sort": {
                    "title": "Järjestellä",
                    "placeholder": "Tulotiedoston tilaus on korkeampi."
                },
                "category": {
                    "title": "Luokka",
                    "placeholder": "Asiakirja näytetään tämän kentän mukaan luokissa"
                },
                "lang": {
                    "title": "Asiakirjakieli",
                    "placeholder": "Valitse asiakirjan kieli"
                }
            }
        },
        "planMgr": {
            "title": "Tilauksen hallinta",
            "description": "Täällä voit lisätä uusia tilaussuunnitelmia, muuttaa kuvausta, hintaa, saldoa, luvan ryhmää, johon tilaussuunnitelmat ovat jo mukana jne.",
            "addNewPlan": "Lisää uusi tilaus",
            "table": {
                "id": "Hio",
                "isSale": "Ota myynti käyttöön",
                "isRenew": "Salli jatkuva maksu",
                "sort": "Lajitella",
                "group": "Ns. Oikeusryhmä",
                "name": "Nimi",
                "nums": "määrä",
                "residue": "Määrä",
                "operate": "toimia",
                "operateBtn": { "update": "Tarkistaa", "delete": "poistaa" }
            },
            "mention": {
                "common": {
                    "success": "menestys",
                    "failure": "Epäonnistuminen",
                    "delMention": "Jos tilaussuunnitelma on jo myynnissä, poista se varovaisesti."
                }
            },
            "form": {
                "title": "Lisää tilaus",
                "items": {
                    "name": {
                        "title": "Tilauksen nimi",
                        "placeholder": "Kirjoita tilaussuunnitelman nimi"
                    },
                    "describe": {
                        "title": "Tilauskuvaus",
                        "placeholder": "Kirjoita tilauksen kuvaus (tukikerros)"
                    },
                    "isSale": { "title": "Ota myynti käyttöön" },
                    "isRenew": { "title": "Ota jatkamismaksu käyttöön" },
                    "group": {
                        "title": "Ns. Oikeusryhmä",
                        "placeholder": "Valitse oikea ryhmä"
                    },
                    "capacityLimit": {
                        "title": "Enimmäismäärä sallittuja käyttäjiä",
                        "placeholder": "Enimmäismäärä sallittuja käyttäjiä"
                    },
                    "planResidue": {
                        "title": "Jäljellä oleva määrä",
                        "placeholder": "Jäljellä oleva määrä"
                    },
                    "sort": {
                        "title": "Järjestellä",
                        "placeholder": "Etuosan lajittelu"
                    },
                    "periodPlaceholder": {
                        "month": "Syötä kuukausimaksuhinta",
                        "quarter": "Syötä neljännesvuosittain maksettu hinta",
                        "halfYear": "Syötä puolivuotinen maksuhinta",
                        "year": "Syötä vuotuinen maksuhinta"
                    }
                }
            }
        },
        "couponMgr": {
            "title": "Kuponginhallinta",
            "description": "Täällä voit lisätä joitain kuponkeja tietyille festivaaleille jne., Joiden avulla käyttäjät voivat käyttää niitä tilauksia ja tarjota alennuksia asettamasi osuuden mukaan.",
            "fetchErr": "Uudelleen, jos tiedot epäonnistuvat",
            "fetchPlanKvFailurese": "Hanki tilaussuunnitelmaluettelo epäonnistui",
            "table": {
                "id": "Hio",
                "enabled": "Onko mahdollista",
                "name": "Nimi",
                "code": "Kuponki",
                "percentOff": "Alennustiedot",
                "capacity": "Kokonaismäärä",
                "residue": "Jäljellä oleva määrä",
                "startTime": "Ottaa käyttöön aika",
                "endTime": "Päätysaika",
                "actions": "toimia",
                "edit": "Muokata",
                "delete": "poistaa"
            },
            "modal": {
                "newCoupon": "Lisää uusi kuponki",
                "editCoupon": "Muokkaa kuponkitietoja",
                "confirmAdd": "Vahvista lisää",
                "confirmEdit": "Vahvista muutos",
                "emptyNotAllow": "Tätä kohdetta vaaditaan",
                "delMention": "Kuponki on virheellinen heti, kun merkintä on poistettu, etkä voi peruuttaa tätä operaatiota.",
                "cancel": "Peruuttaa",
                "name": {
                    "title": "Kupongin nimi",
                    "placeholder": "Kirjoita kupongin nimi"
                },
                "code": {
                    "title": "Kuponki",
                    "placeholder": "Mukauta kuponkikoodia (jätä se tyhjäksi ja luo se mahdollisimman pian)"
                },
                "percentOff": {
                    "title": "Tarjoa tietoja (prosenttiosuus)",
                    "placeholder": "Syötä alennusprosentti (kuten -12%alennus)"
                },
                "activeRange": {
                    "title": "Saatavilla olevien kuponkejen poikkeusvalikoima"
                },
                "capacity": {
                    "title": "Käytettyjen kuponkien enimmäismäärä",
                    "placeholder": "Rajoita enimmäiskäyttöraja (ei rajaa, jos tyhjä)"
                },
                "residue": {
                    "title": "Käytettyjen jäljellä olevien kuponkien lukumäärä",
                    "placeholder": "Aseta kuinka monta kertaa kuponkia käytetään jäljellä"
                },
                "perUserLimit": {
                    "title": "Kuinka monta kertaa kukin käyttäjä voi käyttää kuponkia",
                    "placeholder": "Rajoita kuinka monta kertaa käyttäjä voi käyttää (ei rajaa tilaa)"
                },
                "planLimit": {
                    "title": "Määritä tilaussuunnitelma",
                    "placeholder": "Määritetyn tilaussuunnitelman rajoitukset tarjouksen käyttämiseksi (ei rajaa tilaa varten)"
                }
            }
        },
        "orderMgr": {
            "title": "Tilauksen hallinta",
            "description": "Täällä voit tarkastella kaikkia tilauksen suunnittelijoita tilauksia, suodattimien tilauksia eri käyttäjille, käsitellä käyttäjien tilauksia manuaalisesti ja paljon muuta.",
            "table": {
                "id": "Hio",
                "orderId": "Tilausnumero",
                "email": "Käyttäjän postilaatikko",
                "status": {
                    "title": "Tyyppi",
                    "t1": "Äskettäin ostettu",
                    "t2": "Jatkaa maksua",
                    "t3": "Muokata"
                },
                "planName": "Suunnitella",
                "period": "Kiertää",
                "group": "Oikeusryhmä",
                "amount": "Todellinen maksu",
                "price": "Alkuperäinen hinta",
                "isSuccess": {
                    "title": "Tilaustila",
                    "cancelOrder": "Peruuta tilaus",
                    "passOrder": "Tilapäisesti"
                },
                "createdAt": "Tilan luomisaika",
                "action": { "title": "toimia", "showDetail": "Näyttää yksityiskohdat" }
            },
            "search": "Tiedustelumääräys",
            "resetSearch": "Palauttaa kysely",
            "failureReason": "Epäonnistumisen syy",
            "couponId": "Kuponkitunnus",
            "couponName": "Kupongin nimi",
            "noEntry": "-a",
            "orderDetail": "Tilaustiedot",
            "searchModal": {
                "email": {
                    "title": "Käyttäjän postilaatikko",
                    "placeholder": "Kirjoita käyttäjän sähköposti (sumea haku)"
                },
                "sort": {
                    "title": "Lajittelualgoritmi",
                    "placeholder": "Valitse lajittelualgoritmi",
                    "ASC": "Nouseva järjestys",
                    "DESC": "laskeva järjestys"
                }
            },
            "tradeWaiting": "Ei maksettu",
            "tradeFailure": "Kauppa epäonnistui",
            "tradeSuccess": "menestys"
        },
        "userMgr": {
            "userManager": "Käyttäjän hallinta",
            "description": "Voit hallita kaikkia täällä olevia käyttäjiä, mukaan lukien työntekijät ja järjestelmänvalvojat, myöntää tai peruuttaa hallintaoikeudet, asettaa käyttäjän saldot, palauttaa salasanat, lisätä uusia käyttäjiä manuaalisesti ja paljon muuta.",
            "enterEmail": "Anna sähköpostiosoitteesi",
            "enterValidEmail": "Anna oikea sähköpostimuoto",
            "enterPassword": "Anna salasanasi",
            "passwordMinLength": "Salasanan pituus ei voi olla alle 6 numeroa",
            "passwordMaxLength": "Salasanan pituus ei voi ylittää 20 numeroa",
            "passwordStrength": "Salasanojen on oltava kirjaimia, numeroita ja erikoismerkkejä",
            "validationSuccess": "Vahvistus ohi",
            "validationFailed": "Lomakkeen varmennus epäonnistui, tarkista syöttökohde",
            "editProfile": "Muokkaa tietoja",
            "banUser": "Poista kirjautuminen käytöstä",
            "unbanUser": "Poista käyttäjä",
            "noRecord": "Tallentamaton",
            "normal": "normaali",
            "banned": "Kielletty",
            "deleted": "Kirjautua sisään",
            "nullContent": "Tyhjä",
            "balance": "Määrä",
            "orderCount": "Tilausmäärä",
            "planCount": "Suunnitelmat",
            "actions": "toimia",
            "updateSuccess": "Päivitä onnistuneesti",
            "addUserSuccess": "Uuden käyttäjän lisääminen onnistuneesti",
            "unknownError": "Tuntematon virhe",
            "email": "Sähköposti",
            "registerDate": "Rekisteröintipäivä",
            "isAdmin": "Ylläpitäjä",
            "isStaff": "henkilökunta",
            "accountStatus": "Tilin tila",
            "inviteCode": "Kutsukoodi",
            "query": "Kysely",
            "reset": "Nollata",
            "addNewUser": "Lisätty käyttäjä",
            "searchUser": "Kyselyn käyttäjä",
            "cancel": "Peruuttaa",
            "submit": "lähettää",
            "userEmail": "Käyttäjän postilaatikko",
            "inputUserEmail": "Kirjoita käyttäjän sähköposti (sumea haku)",
            "inputEmail": "Kirjoittaa sähköposti",
            "password": "salasana",
            "inputPassword": "Syötä salasana",
            "editUser": "Muokkaa käyttäjää",
            "no": "ei",
            "yes": "kyllä",
            "mention": {
                "title": "Haluatko varmasti poistaa tilisi käytöstä?",
                "content": "Jos käyttäjä on estetty, käyttäjä ei voi kirjautua sisään {AppName} ja kaikki siihen liittyvät resurssit eivät ole käytettävissä."
            }
        },
        "activation": {
            "activateLog": "Aktivoi tietue",
            "description": "Voit tarkastella kaikkien myytyjen avaimien erityistä aktivointitilaa, tarkastella asiakkaan tunnistuskoodia, aktivointiaikaa jne.",
            "click2getKey": "Napsauta saadaksesi avainsisällön",
            "createdAt": "Luomisaika",
            "turn2keyPage": "Kääntyä avaimeen",
            "showKey": "Näytä avain",
            "email": "Sähköposti",
            "key": "Avain",
            "search": "Haku",
            "filter": "Seulonta",
            "filterAll": "kaikki",
            "filterActive": "Vain tehokas",
            "sortAlgorithm": "Lajittelualgoritmi",
            "sortASC": "Nouseva järjestys",
            "sortDESC": "laskeva järjestys"
        },
        "keysMgr": {
            "keyMgr": "Avainhallinta",
            "description": "Voit tarkistaa kaikki luodut avainsisällöt ja niiden käyttötilan, voimassaolon ajan jne. Voit myös poistaa avaimen käytöstä.",
            "enableKey": "Ota avain käyttöön",
            "disableKey": "Poista avaimet käytöstä",
            "table": {
                "email": "Sähköpostiosoite",
                "key": "Avainsisältö",
                "isValid": "Onko se voimassa vai ei",
                "isUsed": "Käytetäänkö sitä",
                "isExpired": "Onko se vanhentunut",
                "active": "aktiivinen",
                "inactive": "Vanhentunut",
                "yes": "tehokas",
                "no": "Vanhentunut",
                "ok": "normaali",
                "blocked": "Poista asema käytöstä",
                "unUsed": "Ei käytetty",
                "used": "Käytetty",
                "showDetail": "Näyttää yksityiskohdat",
                "blockKey": "Poista avaimet käytöstä",
                "unblockKey": "Kiilto-"
            },
            "searchModal": {
                "searchMethod": "Tutkimusmenetelmä",
                "byId": "Kysely henkilötodistuksella",
                "byContent": "Kysely näppäinten kautta (sumea)",
                "keyId": "Avaintunnus",
                "idPlaceholder": "Kirjoita avaimen tunnus täältä",
                "keyContent": "Avainsana",
                "contentPlaceholder": "Kirjoita osa avaimesta tähän"
            },
            "mention": {
                "blockOk": "Avaimen poistaminen onnistuneesti ID: {id}",
                "title": "Haluatko varmasti poistaa avaimen käytöstä?",
                "content": "Varmista käyttäjäkokemuksen varmistamalla avaimen sisältö uudelleen. Kun avain on poistettu käytöstä, sitä ei enää käytetä asiakkaan puolella."
            },
            "detailModal": {
                "title": "Tärkeimmät yksityiskohdat",
                "userId": "Käyttäjätunnus",
                "planName": "Tilaussuunnitelman nimi",
                "expiredAt": "Voimassaolopäivä",
                "keyGeneratedAt": "Tärkein sukupolven päivämäärä",
                "clientId": "Asiakastunnus"
            }
        },
        "noticeMgr": {
            "title": "Ilmoituksen hallinta",
            "description": "Ilmoituksia voi hallita täällä.",
            "addNotice": "Lisää ilmoitus",
            "table": {
                "id": "Hio",
                "show": "Näyttääkö",
                "title": "Otsikko",
                "createdAt": "Luomisaika",
                "action": { "title": "toimia", "edit": "Muokata", "delete": "poistaa" }
            },
            "mention": {
                "title": "Tiedätkö, että haluat poistaa sen?",
                "content": "Ilmoitus poistetaan välittömästi, etkä voi peruuttaa tätä kannetta."
            },
            "modal": {
                "addNew": "Uusi ilmoitus",
                "title": {
                    "title": "Ilmoituksen otsikko",
                    "placeholder": "Näytä suurena otsikkona pyörän valettuna"
                },
                "content": {
                    "title": "Ilmoitussisältö",
                    "placeholder": "Ilmoituksen kirjoittamisen päähenkilö"
                },
                "tag": {
                    "title": "Ilmoitustunnisteet",
                    "placeholder": "Kirjoita ilmoituksen tunniste"
                },
                "img": {
                    "title": "Taustakuvan URL -osoite",
                    "placeholder": "Käytä oletustausta, jos et täytä"
                }
            }
        },
        "adminTicket": {
            "ticketMgr": "Työjärjestyksen hallinta",
            "description": "Voit tarkastella kaikkia käyttäjiä, jotka ovat lähettäneet työtilauksia.",
            "pendingTicket": "Työmääräykset käsitellään",
            "finishedTicket": "Suoritettu työtilaus",
            "type": {
                "pendingDescription": "Live -työjärjestys, tämä on työjärjestys, joka sinun tulee käsitellä. Jos työjärjestys vahvistetaan valmistuvan, sulje se.",
                "finishedDescription": "Suoritetut työtilaukset, voit tarkastella niitä täällä."
            },
            "chooseOneNecessary": "Sinun tulisi valita vähintään yksi kohde",
            "mention": {
                "title": "Haluatko varmasti sulkea lipun?",
                "content": "Tilaus arkistoidaan täytetylle tilaukselle sen sulkemisen jälkeen."
            }
        }
    },
    "userLogin": {
        "loginToContinue": "Kirjaudu sisään jatkaaksesi",
        "email": "Sähköpostiosoite",
        "password": "salasana",
        "haveNoAccount": "Eikö sinulla ole vielä tiliäsi?",
        "login": "sisäänkirjautuminen",
        "reg": "Ilmoittautua nyt",
        "otherMethods": "Tai käytä muita menetelmiä jatkaaksesi",
        "github": "Jatka GitHub -tilillä",
        "microsoft": "Jatka Microsoft -tiliä",
        "apple": "Jatka Apple -tiliä",
        "google": "Jatka Google -tiliä",
        "backHomePage": "Takaisin kotisivulle",
        "form": {
            "emailRequire": "Sähköpostiosoite vaaditaan",
            "passwordRequire": "Et ole vielä syöttänyt salasanasi"
        },
        "authPass": "Vahvistus ohi",
        "loginFailure": "Kirjautuminen epäonnistui",
        "checkForm": "Tarkista lomake",
        "if2FaEnabledHint": "Jos olet ottanut käyttöön kaksivaiheisen varmennuksen (ei vaadita)",
        "reqErr": "Yritä myöhemmin uudelleen, jos virhe on",
        "accountLocked": "Tilisi voidaan rekisteröidä tai kieltää, eikä sitä voida käyttää tällä hetkellä.",
        "tokenNotExist": "Anna tunnusmerkki"
    },
    "userRegister": {
        "backHomePage": "Takaisin kotisivulle",
        "newAccount": "Valmista uusi tilisi",
        "email": "Sähköpostiosoite",
        "verifyCode": "Sähköpostivahvistuskoodi",
        "invalidEmailFormat": "Sähköpostimuoto on laiton",
        "sendVerifyCode": "Lähetä sähköpostin vahvistuskoodi",
        "sendSuccess": "Sähköposti on lähetetty onnistuneesti, tarkista sähköposti.",
        "pwd": "salasana",
        "pwdAgain": "Vahvista salasana",
        "inviteCode": "Kutsikoodi (valinnainen)",
        "agreement": "Olen lukenut ja suostunut",
        "terminalUserAgreement": "Käyttäjän ehdot",
        "reg": "rekisteröidä",
        "infoIncomplete": "Epätäydelliset tiedot",
        "pwdIncorrect": "Salasana epäjohdonmukainen",
        "regSuccess": "Rekisteröinti on onnistunut ja palaa kirjautumiseen",
        "regFailure": "Rekisteröinti epäonnistui",
        "success": "menestys",
        "failure": "Epäonnistuminen",
        "unknownErr": "Tuntematon virhe",
        "verifyCodeErr": "Varmennuskoodivirhe",
        "verifyCodeExpireErr": "Vahvistuskoodi on väärä tai on vanhentunut.",
        "thisMailAlreadyExist": "Sähköposti on rekisteröity",
        "pageConfigFetchFailure": "Jos kokoonpanon hankinta epäonnistuu, päivitä ja yritä uudelleen",
        "stopRegisterTitle": "Rekisteröinti on lopetettu",
        "stopRegisterHint": "Valitettavasti rekisteröintitoiminto on keskeytetty. Jos tarvitset sitä, yritä myöhemmin uudelleen tai ota yhteyttä tukitiimiin saadaksesi lisätietoja. Kiitos ymmärryksestäsi ja tuestasi.",
        "passwordComplexRequirePart1": "* Salasanojen on oltava vaatimusten mukaisia",
        "passwordComplexRequirePart2": "Monimutkaisuusvaatimukset",
        "passwordComplexHint1": "1. Vaaditaan kolme tai enemmän suuria kirjaimia, pieniä kirjaimia, numeroita, erityisiä symboleja.",
        "passwordComplexHint2": "2. Pituus on yli 10 numeroa",
        "form": {
            "checkForm": "Tarkista lomakkeen sisältö",
            "emailFormatErr": "Sähköpostiosoitteen muoto on väärä",
            "gmailLimitErr": "Gmail -useita nimiä ei sallita",
            "gmailDotNotAllowed": "\".\"",
            "gmailPartLowerForced": "Gmail -osoitteen paikallisen osan on oltava kaikki lukijat",
            "googlemailNotAllowed": "GoogleMail -osoitetta ei tueta",
            "verifyCodeRequire": "Anna varmennuskoodi",
            "verifyCodeFormatErr": "Vahvistuskoodimuoto on väärä",
            "passwordRequire": "Anna salasanasi",
            "passwordLengthRequire": "Salasanan pituuden on oltava suurempi tai yhtä suuri kuin 10 numeroa",
            "passwordComplexRequire": "Salasanassa on oltava vähintään kolme tyyppiä suuria kirjaimia, pieniä kirjaimia, numeroita ja erityisiä symboleja.",
            "passwordAgainNotMatch": "Salasanat syötetään kahdesti epäsuhta",
            "passwordAgainRequire": "Anna vahvista salasana",
            "inviteCodeRequire": "Kutsukoodi vaaditaan"
        },
        "hCaptcha": {
            "passed": "Ihmisen ja koneen todentaminen ohi",
            "expired": "Yritä uudelleen vanhentuneen jälkeen",
            "challengeExpired": "Vahvistus ylittää ajan",
            "err": "Muut virheet"
        },
        "allRightsReserved": "Kaikki oikeudet pidätetään",
        "securityAndLaws": "HCaptcha on suojattu ja vahvistettu verkkosivustolla, ja se on paikallisten lakien soveltamisessa."
    },
    "userSummary": {
        "title": "Instrumentti",
        "myPlan": "Tilaukseni",
        "shortcut": "Pikakuvakkeet",
        "timeLeft": "Tilaus on kelvollinen ja vanhenee {msg}.",
        "toPurchase": "Ostaa tilaus",
        "tutorial": {
            "title": "Katso opetusohjelma",
            "content": "Opi käyttämään {nime}"
        },
        "checkKey": {
            "title": "Näytä avaimet",
            "content": "Esittele nopeasti avaimet asiakkaalle käytettäväksi"
        },
        "renewPlan": {
            "title": "Jatka tilausta",
            "content": "Suorita jatkuva maksu nykyisestä tilauksestasi"
        },
        "appDown": {
            "title": "Sovelluksen lataus",
            "content": "Lataa sovelluksemme eri alustoilta paremman fyysisen kokemuksen saavuttamiseksi"
        },
        "support": {
            "title": "Ongelman kanssa",
            "content": "Jos kohtaat ongelmia, voit kommunikoida ihmisten kanssa työjärjestyksen kautta."
        },
        "haveTicket": "Sinulla on {kreivi} käsitettävä käsitettävä",
        "toCheckTicket": "Mene ja tarkista se",
        "showAllKeys": "Näytä kaikki avaimet"
    },
    "userDocument": {
        "title": "Käytä asiakirjaa",
        "description": "Voit tarkistaa täällä erilaisia ​​asiakirjoja, mukaan lukien, mutta rajoittumatta verkkosivuston käyttömenetelmiin, varotoimenpiteisiin, toimintamenettelyihin jne. Jos löydät artikkelista virheitä, lähetä työtilaus.",
        "searchPlaceholder": "Kirjoita hakujen sisältö (sumea haku)",
        "searchBtn": "haku",
        "noContentTitle": "Ei tuloksia",
        "noContentTitleHint": "Ei ole asiakirjaa, joka vastaa hakutuloksia tai kieltä, yritä vaihtaa avainsanaan."
    },
    "newPurchase": {
        "title": "Ostaa tilaus",
        "description": "Voit valita täältä parhaiten sopivan tilaussuunnitelman. Jos saldo on riittävä, lataa ensin tilaus.",
        "headerPlaceholder": "Valitse sinulle sopiva suunnitelma",
        "purchaseBtn": "Tilata",
        "noLeft": "Riittämätön määrä jäljellä olevaa",
        "monthPay": "Kuukausimaksu",
        "moreMethod": "Lisää vaihtoehtoja"
    },
    "newSettlement": {
        "err": "Virhe",
        "monthPay": "Kuukausimaksu",
        "quarterPay": "Neljännesvuosittainen maksu",
        "halfYearPay": "Puolen vuoden maksu",
        "yearPay": "Vuotuinen maksu",
        "payCycle": "Maksusykli",
        "couponPlaceholder": "Onko sinulla kuponki?",
        "verifyBtn": "Varmennus",
        "orderTotalTitle": "Tilauksen kokonaismäärä",
        "total": "Kokonais",
        "submitOrder": "Antaa tilaus",
        "coupon": "Kupongit",
        "notify": {
            "passTitle": "Vahvistus ohi",
            "couponVerified": "Kuponki voimassa oleva",
            "couponInvalid": "Kuponki on virheellinen",
            "couponIsNull": "Syötetty kuponkikoodi ei voi olla tyhjä"
        }
    },
    "userProfile": {
        "title": "Henkilökohtainen keskus",
        "myWallet": "Rahalaukku",
        "walletSub": "Tilin saldo (käytetään vain kulutukseen)",
        "alertPwd": "Muokkaa avainta",
        "oldPwd": "Vanha avain",
        "oldPwdSub": "Anna vanha salasana",
        "newPwd": "Uusi avain",
        "newPwdSub": "Anna uusi avain",
        "newPwdAgain": "Vahvista avain",
        "newPwdAgainSub": "Anna uusi avain uudelleen",
        "saveBtn": "tallentaa",
        "notify": "ilmoittaa",
        "enableNotify": "Ota voimassa oleva ilmoitus",
        "deleteAccount": "Rekisteröidä tili",
        "deleteAccountSub": "Tilisi merkitään poistetuksi, jos sinun on käytettävä palveluita uudelleen, rekisteröidy uudelleen",
        "deleteBtn": "Rekisteröi tilini",
        "oldPwdVerifiedFailure": "Vanha salasanan tarkistus epäonnistui",
        "alertFailure": "Avaimen muokkaaminen epäonnistui",
        "alertSuccess": "Muutos oli onnistunut",
        "alertSuccessSub": "Kirjaudu sisään uudella salasanallasi",
        "errorPwdFormat": "Salasanamuotovirhe",
        "pwdNotMatch": "Salasanan tulot ovat epäjohdonmukaisia",
        "oldPwdNotNull": "Vanha salasana ei voi olla tyhjä",
        "toTopUp": "Mennä lataamaan",
        "deleteMyTitle": "varoittaa",
        "deleteMyContent": "Oletko varma, että poistat tilisi? Jos sinun on käytettävä palvelua, rekisteröi uudelleen.",
        "deleteMyPositiveText": "Vahvista deleetio",
        "deleteMyNegativeText": "Peruuttaa",
        "deletedSuccessMsg": "Tili on poistettu ja aika on aika myöhemmin.",
        "deleteErrOccur": "Kohtattu virhe",
        "faAuth": "Kaksivaiheinen varmennus 2FA",
        "faAuthHint": "Kaksivaiheinen varmennus on tietoturvamekanismi, joka lisää suojaustasoja kirjautumiseen tilillesi. Salasanan syöttämisen jälkeen käyttäjän on suoritettava identiteettivahvistuksen toinen vaihe, kuten puhelimeen lähetetyn varmennuskoodin syöttäminen, identiteettivahvistussovelluksen avulla dynaamisen koodin luomiseksi tai sen vahvistaminen biologisten ominaisuuksien, kuten sormenjälkien, avulla.",
        "faAuthStatus": "Kaksivaiheinen varmennustila",
        "faEnabled": "Uudelleenkäytetty",
        "faNotEnabled": "Ei aktivoitu",
        "setup2Fa": "Kaksivaiheinen varmennus",
        "disable2Fa": "Peruuta kaksivaiheinen varmennus",
        "unknownErr": "Tuntematon virhe",
        "disable2FaCancelled": "Peruutettu",
        "closed2FaHint": "Kaksivaiheinen varmennus on suljettu, lisää se tarvittaessa uudelleen.",
        "setup2FaModal": {
            "followStep": "Lisää se tarkistuslaitteeseesi kehotusten mukaisesti",
            "step1Title": "Noudata alla olevia vaiheita ottaaksesi 2FA: n varmennuksen",
            "step1Context1": "1. Sinulla on oltava yleinen varmennuslaite mobiililaitteessasi",
            "step1Context2": "2.",
            "step1Context3": "3. Tämä QR -koodi sisältää varmennustiedot ja ainutlaatuiset avaimet.",
            "step2Context1": "Jotta varmistetaan, että varmennuslaitetta voidaan käyttää normaalisti, meidän on suoritettava testi.",
            "test2Fa": "Testata",
            "cancel": "Peruuttaa"
        },
        "deleteMyAccountModal": {
            "title": "Rekisteröidä tili",
            "contentLine1": "Kirjanpito on peruuttamaton toimenpide. Kun olet vahvistanut, että se on poistettu, menetät pysyvästi pääsyn tilille, mikä tarkoittaa, että et voi kirjautua sisään uudelleen ja kaikki tähän tilille liittyvät tiedot, mukaan lukien, mutta rajoittumatta henkilökohtaisiin tietoihisi, historiaan, keräyssisältöön, ostotietueisiin jne., Ei voi käyttää sitä uudelleen.",
            "contentLine2": "Jos sinulla on jatkuvaa liiketoimintaa alustallamme, kuten keskeneräiset tilaukset, osallistuminen, tilauspalvelut jne., Ne lopetetaan tai peruutetaan tilisi poistamisella, mikä saattaa tuoda sinulle vastaavat tappiot. Samanaikaisesti yhteystiedot ja vuorovaikutustiedot sinua ja muita käyttäjiä tämän alustan kautta ei enää ole.",
            "contentLine3": "Vahvista päätöksesi uudelleen. Jos sinulla on kysymyksiä tai kysymyksiä, ota yhteyttä asiakaspalveluun ja vastaamme siihen vilpittömästi. Jos haluat silti poistaa tilisi, napsauta Vahvista Poista -painiketta.",
            "inputHint1": "Tulla \"",
            "inputHint2": "\"Jatka.",
            "confirmDelete": "Vahvista poista"
        },
        "failure": "Kohtasi virheen",
        "notLatestHint": "Henkilötietojen päivitys epäonnistui, ja nykyiset tiedot eivät välttämättä ole viimeisin.",
        "resetPassword": {
            "previousPasswordRequire": "Anna vanha salasana",
            "previousPasswordVerifiedFailure": "Vanha salasanan tarkistus epäonnistui",
            "passwordRequire": "Anna salasanasi",
            "passwordConflict": "Uusi salasana ei voi olla sama kuin edellinen salasana",
            "passwordLengthRequire": "Salasanan pituuden on oltava suurempi tai yhtä suuri kuin 10 numeroa",
            "passwordComplexRequire": "Salasanassa on oltava vähintään kolme tyyppiä suuria kirjaimia, pieniä kirjaimia, numeroita ja erityisiä symboleja.",
            "passwordAgainNotMatch": "Salasanat syötetään kahdesti epäsuhta",
            "passwordAgainRequire": "Anna vahvista salasana",
            "updatePasswordFailure": "Salasanan päivittämisessä tapahtui virhe"
        },
        "secondary": {
            "title": "Perustiedot",
            "card": {
                "avatar": "Käyttäjän pään kuva",
                "name": "Käyttäjänimi",
                "lastLoginAt": "Viimeinen kirjautumisaika",
                "accountStatus": "Tilin tila"
            },
            "modify": {
                "uploadIconHint": "Ladata",
                "alterAvatar": "Lisää/muokkaa pääkuvia",
                "alterShallow": "* Napsauta pääkuvaa tai käyttäjänimi lisätäksesi tai muokataksesi (tyhjentäminen ei ole sallittua asetuksen jälkeen)",
                "alterName": "Muokkaa käyttäjänimeä",
                "input": {
                    "label": "käyttäjänimi",
                    "placeholder": "Kirjoita käyttäjänimi",
                    "spaceIsNotAllowed": "Käyttäjätunnuksen varmennus ei ohita",
                    "require": {
                        "p1": "Puhdas numerot eivät ole sallittuja ja numerot avataan",
                        "p2": "Tilaa ei sallittu",
                        "p3": "Yli kolme"
                    }
                }
            },
            "mention": {
                "alterNameSuccess": "Käyttäjätunnuksen muutos oli onnistunut",
                "alterNameErr": "Jos käyttäjänimi epäonnistuu, yritä myöhemmin uudelleen",
                "newNameIsNotValid": "Laiton käyttäjänimi",
                "click2SetName": "Napsauta Aseta käyttäjätunnus",
                "fetchAvatarErr": "Jos pääkuvat tiedot epäonnistuvat, yritä uudelleen myöhemmin",
                "alterAvatarErr": "Jos pääkuva epäonnistuu, yritä uudelleen myöhemmin",
                "success": "menestys",
                "alterAvatarSuccess": "修改頭像成功，",
                "uploadImageHint": "Voit ladata kuvia uutena otsikkosi",
                "imageRequire": {
                    "title": "Huomautus",
                    "p1": "Voit ladata valtamuodot, kuten *.jpg (jpeg), *.png, *.webp jne.",
                    "p2": "Kuvan pituussuhde on 1: 1 (neliö), jos se on erilainen suhde, se leikataan keskustan neliöksi ja ylimääräinen poistetaan.",
                    "p3": "Lähettämäsi kuvan koko asetetaan arvoon 160px."
                },
                "click2Upload": "Napsauta tai vedä tiedostoa alueelle ladataksesi",
                "uploadWarn": "*Älä lataa arkaluontoisia tietoja, kuten pankkikorttia, luottokorttia, salasanaa ja turvakoodia."
            }
        }
    },
    "userKeys": {
        "myKeys": "Minun avain",
        "description": "Voit tarkastella tapahtuman tilaa, vanhenemispäivää jne. Kaikille avaimille.",
        "noKeys": "Sinulla ei ole vielä voimassa olevaa ostotietuetta",
        "keyDetail": "Tärkeimmät yksityiskohdat",
        "keyId": "Avaintunnus",
        "orderId": "Linkkitilaustunnus",
        "clientId": "Aktivoi asiakastunnus",
        "active": "Pätevyyden aikana",
        "inActive": "Vanhentunut",
        "valid": "Avain ehto on normaali",
        "invalid": "Avain on poistettu käytöstä",
        "isUsed": "Aktivoitu",
        "noUsed": "Ei vielä käytetty",
        "releaseData": "Tärkein sukupolven päivämäärä",
        "expirationData": "Voimassaolopäivä",
        "none": "-a",
        "authorizationFor": "Avainlupa",
        "hoverClickMention": "Napsauta kopioidaksesi leikepöydälle",
        "copiedSuccessMessage": "Avainkopio on kopioitu onnistuneesti.",
        "copyFailure": "Kopiointi epäonnistuminen",
        "hoverCopiedSuccessMention": "Kopioida onnistuneesti"
    },
    "userOrders": {
        "myOrders": "Minun tilaukseni",
        "description": "Kaikki tilauksesi näytetään täällä, jos sinulla on maksamattomia tilauksia, voit napsauttaa jatkaa maksua tai peruuttaa tilauksen ja voit tarkastella tilauksen yksityiskohtia täytettyjen tilausten jälkeen.",
        "orderId": "Hio",
        "planName": "Tilauksen nimi",
        "planCycle": "Tilausjakso",
        "orderPrice": "Tilausmäärä",
        "orderStatus": "Tilaustila",
        "createdAt": "Luomisaika",
        "operate": "toimia",
        "showDetail": "Tilaustiedot",
        "cancelOrder": "Peruuta tilaus",
        "canceled": "Peruutettu",
        "period": {
            "monthPay": "Kuukausimaksu",
            "quarterPay": "Neljännesvuosittainen maksu",
            "halfYearPay": "Puolen vuoden maksu",
            "yearPay": "Vuotuinen maksu"
        },
        "orderStatusTags": {
            "success": "menestys",
            "cancelled": "Epäonnistuminen",
            "notPay": "Ei maksettu"
        },
        "orderCancelled": "Tilaus peruutettu",
        "unknownErr": "Tuntematon virhe"
    },
    "userTopUp": {
        "topUp": "Tili ladata",
        "description": "Voit ladata tilisi täältä, tukea räätälöityjä latausmääriä, ja voit myös kiinnittää huomiota siihen, onko alla näkyviä alennustietoja ja käyttää mainitulla maksutapa alennuksilla.",
        "chooseTopUpAmount": "Valitse latausmäärä",
        "quickSelect": "Nopeasti valita",
        "customAmount": "Räätälöity määrä",
        "maxAmount": "Enimmäismäärä: 10 000 000",
        "amountInputPlaceHolder": "Syötä ladattava määrä",
        "otherAmount": "Muut määrät",
        "payMethod": "Maksutapa",
        "wechat": "WeChat -maksu",
        "alipay": "Maksu",
        "apple": "Apple Pay",
        "yourAmount": "Sinun summasi",
        "discount": "Tarjoukset",
        "accountBalance": "Tilin saldo",
        "balanceResult": "Tasapainon määrä",
        "commitTopUp": "lähettää",
        "payMethodNotAllow": "Maksutapa ei ole käytettävissä.",
        "topUpIssueOccur": "Onko sinulla ongelmia lataus?",
        "payIssueOccur": "Onko sinulla ongelmia maksun kanssa?",
        "chatWithUs": "Ota yhteyttä asiakaspalveluun",
        "pay": "Maksaa",
        "qrCodeScannedSuccess": "QR -koodin skannaus onnistuneesti",
        "orClickToApp": "Tai napsauta palkkia ja käänny sovellukseen ja jatka",
        "topUpSuccess": "Lataus onnistuneesti",
        "thankU": "Kiitos tuestasi"
    },
    "userConfirmOrder": {
        "switchPlan": "Vaihtaa",
        "cancelOrder": "Peruuta tilaus",
        "yourPrice": "Sinun hinta",
        "couponOffset": "Kuponki -alennus",
        "price": "Hinta",
        "submit": "lähettää",
        "monthPay": "Kuukausimaksu",
        "quarterPay": "Neljännesvuosittainen maksu",
        "halfYearPay": "Puolen vuoden maksu",
        "yearPay": "Vuotuinen maksu",
        "goodInfo": "Tuotetiedot",
        "cycleOrType": "Sykli/tyyppi",
        "orderNumber": "Tilausnumero",
        "createdAt": "Luomispäivä",
        "orderExpired": "Tilaus on ylittänyt",
        "paySuccessfully": "Maksu oli onnistunut, kiitos tuestasi.",
        "balanceNotEnough": "Jos saldo on riittämätön, lataa ensin tilaus.",
        "orderErrorOccur": "Tilauksen aikana tapahtui virhe",
        "orderCancelled": "Tilaus peruutettu"
    },
    "paymentResultParts": {
        "goodInfoView": { "goodInfo": "Tuotetiedot" },
        "orderInfoView": { "orderInfo": "Tilata tiedot" }
    },
    "orderPartUniversal": {
        "period": {
            "monthPay": "Kuukausimaksu",
            "quarterPay": "Neljännesvuosittainen maksu",
            "halfYearPay": "Puolen vuoden maksu",
            "yearPay": "Vuotuinen maksu"
        },
        "orderDataHex": {
            "goodInfo": "Tuotetiedot",
            "orderInfo": "Tilata tiedot",
            "cycleOrType": "Sykli/tyyppi",
            "orderNumber": "Tilausnumero",
            "createdAt": "Luomispäivä",
            "amount": "Maksu",
            "paidAt": "Maksuaika"
        }
    },
    "orderDetail": {
        "finished": "Valmis",
        "finishedAndSuccessDescription": "Tilaus on maksettu ja avattu",
        "useManual": "Katso opetusohjelma",
        "payPending": "Ei vielä maksettu",
        "pendingDescription": "訂單暫時保留，可以點擊下面的按鈕以繼續支付。",
        "toPay": "Mennä maksamaan",
        "outDate": "Ovat vanhentuneet",
        "outDateDescription": "Koska olet peruuttanut tilauksen tai et suorittanut maksua määritettyyn aikaan, tilaus on peruutettu ja voit valita tilauksesi uudelleen.",
        "chooseNewPlan": "Valitse uusi tilaussuunnitelma"
    },
    "userInvite": {
        "myInvite": "Kutsuni",
        "description": "Jos järjestelmänvalvoja on ottanut käyttöön kutsualennukset, voit nähdä erityiset alennussäännöt täältä. Kun olet luonut kutsulinkin, jaa se muiden kanssa rekisteröityäksesi ja saat alennuksen.",
        "unit": "Ihmisten lukumäärä",
        "inviteCodeMgr": "Kutsumasi",
        "generateInviteCode": "Luo satunnaisia ​​kutsukoodeja",
        "faCodeManage": "Kutsukäyttö",
        "email": "Sähköposti",
        "createdAt": "Rekisteröintiaika",
        "createFaCodeFailure": "Luominen epäonnistui",
        "linkCopiedSuccess": "Linkki kopioitu onnistuneesti",
        "generateFaCode": "Luo kutsukoodi",
        "flushFaCode": "Päivitä kutsukoodi",
        "faCode": "Kutsukoodi",
        "noFaCode": "Sinulla ei ole vielä kutsukäkkiä, tee se.",
        "faLink": "Kutsulinkki",
        "generateFaCodePlease": "Luo kutsukoodi",
        "usersMyInvited": "Kutsun käyttäjät",
        "generateCodeConfirm": "Vahvista Luo/virkistä",
        "generateCodeHint": "Huomaa, että kutsukoodia ei voida sulkea luomisen jälkeen.",
        "positiveClick": "Vahvistaa",
        "negativeClick": "Peruuttaa"
    },
    "userTickets": {
        "description": "Jos kohtaat käytön ongelmat, lähetä työtilaus täällä.",
        "ticketId": "Hio",
        "ticketSubject": "Työtilausteema",
        "ticketUrgency": "Työtilausluokka",
        "ticketContent": "Työtilauksen sisältö",
        "ticketUrgencyLevel": {
            "low": "Matala",
            "med": "keskimmäinen",
            "high": "korkea"
        },
        "ticketStatus": "Työjärjestyksen tila",
        "ticketCreatedAt": "Luomisaika",
        "lastResponse": "Viimeinen vastaus",
        "operate": "toimia",
        "checkTicket": "Näytä työtilaus",
        "closeTicket": "Sulkea työjärjestys",
        "userTickets": "Historia toimii",
        "addTicket": "Uusi rakennustilaus",
        "ticketActive": "Ei suljettu",
        "ticketInActive": "Suljettu",
        "form": {
            "ticketSubjectDescription": "Kirjoita työtilausaihe",
            "ticketUrgencyDescription": "Valitse työjärjestyksen kiireellisyys",
            "ticketBody": "Syötä ratkaistava ongelma mahdollisimman kattavasti.",
            "ticketNotFinished": "Kokeile täydellisen työtilauksen tietoja"
        },
        "checkForm": "Tarkista, onko lomake valmis",
        "cancel": "Peruuttaa",
        "submit": "lähettää",
        "commitNewTicketSuccess": "Lähetä uusi työtilaus onnistuneesti",
        "commitNewTicketFailure": "Lähetä uusi työtilausvirhe",
        "noReply": "Ei vielä vastausta",
        "noTickets": "Et ole vielä lähettänyt työtilausta",
        "ticketCloseSuccess": "Työtilaus suljettiin onnistuneesti",
        "ticketCloseFailure": "Työjärjestyksen sulkeminen epäonnistui",
        "chatDialog": {
            "input": {
                "finished": "Työmääräys on tehty",
                "inputHere": "Kirjoita lähetettävä viesti"
            },
            "send": "Lähettää",
            "missingToken": "WebSocket -yhteyttä ei voida luoda, koska rahakkeet puuttuvat.",
            "msgEmptyNotAllowed": "Viestin sisältö ei voi olla tyhjä",
            "accessNotPermit": "Laiton vierailu",
            "sendSuccess": "Lähetä viesti onnistuneesti"
        }
    },
    "userActivation": {
        "activateLog": "Aktivoi tietue",
        "description": "Täällä voit tarkastella aktivointitietueesi ja asettaa määrittelemättömän laitteen tai määrittää kunkin avaimen ja aktivointitietueen muistiinpanot.",
        "id": "Hio",
        "orderId": "Tilausnumero",
        "orderStatus": "Tilata",
        "createdAt": "Luomisaika",
        "operate": "toimia",
        "userId": "Käyttäjätunnus",
        "email": "Sähköposti",
        "keyId": "Avaintunnus",
        "isBind": "Aktivoida",
        "active": "tehokas",
        "inactive": "Tehoton",
        "requestAt": "Pyyntöaika",
        "clientVersion": "Asiakaspuoli",
        "osType": "Käyttöjärjestelmä",
        "remark": "Huomautus",
        "noRemark": "Ei muistiinpano",
        "showDetail": "Katso yksityiskohdat",
        "actions": "toimia",
        "details": "Yksityiskohdat",
        "keyContent": "Avainsisältö",
        "keyGeneratedAt": "Tärkein sukupolven aika",
        "activateRequestAt": "Aktivoi pyyntöaika",
        "useIssueOccur": "Onko sinulla vaikeuksia käyttää sitä?",
        "chatWithUs": "Ota yhteyttä",
        "cancelBind": "Peruuta asetus",
        "alterRemark": "Muokkaa muistiinpanoja",
        "commitRemark": "lähettää",
        "updateSuccess": "Päivitä onnistuneesti",
        "setRemark": "Määritä muistiinpanot täältä"
    },
    "userAppDownload": {
        "title": "Sovelluslataus",
        "description": "",
        "common": {
            "title": "Lataa sovelluksemme",
            "shallow2": "Hanki sovelluksemme eri asiakkaille",
            "shallow": "Sovelluksemme avulla voit käyttää palveluitamme helpommin, poistamalla selaimen monimutkaisen toiminnan joka kerta;",
            "noDownload": "Olemme pahoillamme, että latausta ei ole vielä saatavilla. Yritä myöhemmin uudelleen."
        },
        "suffix": {
            "p1": "*Käytä MacOS -alustalla sovelluksia MacOS14: llä ja yllä ja tee yhteistyötä Apple -sirujen kanssa (M1 tai korkeampi).",
            "p2": "Tämän sovelluksen esittämien tietojen ja käytön tulisi noudattaa paikallisia käyttömääräyksiä."
        },
        "card": {
            "common": { "welcome": "Tervetuloa" },
            "mobile": {
                "designFor": "Suunnittelu mobiililaitteelle",
                "designShallow": "Voit saada mobiilisovelluksemme täältä",
                "iosDownloadShallow": "Lataa iOS -asiakas",
                "androidDownloadShallow": "Lataa Android -asiakas"
            },
            "desktop": {
                "designFor": "Työpöydän suunnittelu",
                "designShallow": "Voit saada työpöydän mobiilisovelluksen täältä",
                "windowsDownloadShallow": "Lataa Windows -asiakas",
                "osxDownloadShallow": "Lataa MacOS -asiakas",
                "linuxDownloadShallow": "Lataa Linux -asiakas"
            }
        },
        "downloadDisabledHint": "Valitettavasti järjestelmänvalvoja ei ole väliaikaisesti käytettävissä tai vammaisia. Jos tarvitset sitä, ota yhteyttä tekniikkaan.",
        "windows": {
            "title": "Windows NT",
            "shallow": "Tämä asiakas sopii Windows -käyttöjärjestelmille, joissa on NT -ytimet, tarkista dokumentointisivu yhteensopivuustuki."
        },
        "osx": {
            "title": "MacOS (OSX)",
            "shallow": "Tämä asiakas sopii Darwin -ytimen MacOS (OSX) -käyttöjärjestelmään, tarkista dokumentaatiosivu yhteensopivuustuki."
        },
        "linux": {
            "title": "Linux",
            "shallow": "Tämä asiakas sopii Linux -ytimen erilaisiin jakeluihin. Eri jakelusarjojen vuoksi tarkista dokumentointisivu yhteensopivuuden tueksi."
        },
        "android": {
            "title": "Androidi",
            "shallow": "Tämä asiakas sopii mobiililaitteille, jotka on varustettu Google Android -käyttöjärjestelmällä, tarkista asiakirja -sivulta yhteensopivuustuki."
        },
        "ios": {
            "title": "IOS",
            "shallow": "Tämä asiakas sopii mobiililaitteisiin, jotka on varustettu Apple IOS -käyttöjärjestelmällä, tarkista dokumentointisivu yhteensopivuustuki."
        }
    },
    "welcome": {
        "A": {
            "aboutUs": "Meistä",
            "pricing": "Tilaushinta",
            "login": "Kirjautua sisään",
            "register": "Rekisteröidä tili",
            "welcomeTo": "Tervetuloa",
            "welcomeToSub": "Pitkän tunnelin läpi läänissä on lumimaassa. Yötaivaan alla maa oli valkoinen ja juna pysähtyi signaaliaseman edessä. \"Kawabata Yasunari aloitti täällä lumimaasta erittäin kivellä ytimekäs teksti.",
            "startingUse": "Aloittaa",
            "whyUs": "Miksi valita meidät",
            "whyUsSub": "\"Prefektuurin pitkän tunnelin läpi se on lumimaa. Yötaivaan alla maa on valkoinen ja juna pysähtyy signaaliaseman edessä.\"",
            "browseSafe": "Selata",
            "browseSafeSub": "Erinomainen palomuurisuodatinjärjestelmä voi tehokkaasti estää online -kaloja ja haitallisia verkkosivustoja",
            "encrypt": "Päähän -salaus",
            "encryptSub": "Kaksisuuntainen SSL ja päähän -salaus suojaavat yksityisyyttäsi ja turvallisuuttasi, jopa palvelimet eivät voi lukea tietojasi",
            "mgr": "Tehokas hallinta",
            "mgrSub": "Yksi käyttöliittymä hallitsee kaikkia avaimia täydellisillä ja rikkailla hallintatoiminnoilla, eikä sinun tarvitse huolehtia siemensyöksyn ongelmien tilaamisesta",
            "fast": "Kätevä ja nopea",
            "fastSub": "Anna täydelliset API-tiedostot WebAppeille tai upotettu kolmansien osapuolien ohjelmistoihin",
            "fastLink": "Pikayhteys",
            "subscribeUs": "Seurata meitä",
            "filingsCode": "Tiedostonumero {koodi}"
        }
    },
    "pagination": {
        "perPage10": "10 data/sivu",
        "perPage20": "20 data/sivu",
        "perPage50": "50 data/sivu",
        "perPage100": "100 data/sivu"
    },
    "modal": { "cancel": "Peruuttaa", "confirm": "vahvistaa" },
    "week": {
        "monday": "maanantai",
        "tuesday": "tiistai",
        "wednesday": "keskiviikko",
        "thursday": "torstai",
        "friday": "perjantai",
        "saturday": "Lauantai",
        "sunday": "sunnuntai"
    },
    "period": {
        "month": "Kuukausimaksu",
        "quarter": "Neljännesvuosittainen maksu",
        "halfYear": "Puolen vuoden maksu",
        "year": "Vuotuinen maksu"
    },
    "operate": {
        "cancel": "Peruuttaa",
        "confirm": "vahvistaa",
        "update": "uudistaa",
        "add": "Lisätä jtk"
    },
    "notFound": {
        "title": "404 Sivua ei ole olemassa",
        "description": "Emme löydä pyytämääsi sivua, se on ehkä poistettu tai linkitetty virheisiin. Jos luulet tämän olevan virhe, lähetä työtilaus ottaaksesi yhteyttä.",
        "p1": "Se palaa kotisivulle {s.} S: n jälkeen, ja jos selaimesi ei vastaa, napsauta alla olevaa painiketta.",
        "manualBack": "Palaa kotisivulle"
    },
    "forbidden": {
        "title": "403 Ei oikeuksia",
        "description": "Sinulla ei ehkä ole riittävää lupaa käyttää tätä sivua. Jos luulet tämän olevan virhe, lähetä työtilaus ottaaksesi yhteyttä."
    }
}
