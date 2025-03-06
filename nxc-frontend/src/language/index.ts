type LangOption = {
    label: string;
    key: string;
    disabled: boolean;
}

import en_US from "@/language/en_US";
import zh_CN from "@/language/zh_CN";
import ja_JP from "@/language/ja_JP";
import zh_HK from "@/language/zh_HK";
import fr_FR from "@/language/fr_FR";
import fi_FI from "@/language/fi_FI";
import ko_KR from "@/language/ko_KR";
import ru_RU from "@/language/ru_RU";
import de_DE from "@/language/de_DE";

export const langs = {
    en_US,
    zh_CN,
    ja_JP,
    zh_HK,
    fr_FR,
    fi_FI,
    ko_KR,
    ru_RU,
    de_DE,
}

export const langOptions: LangOption[] = [
    {
        label: 'English', // 英文
        key: 'en_US',
        disabled: false
    },
    {
        label: '简体中文',  // 中文简体
        key: 'zh_CN',
        disabled: false
    },
    {
        label: '繁體中文',  // 中文繁体
        key: 'zh_HK',
        disabled: false
    },
    {
        label: '日本語 Japan', // 日语
        key: 'ja_JP',
        disabled: false
    },
    {
        label: '한국인 Koria', // 韩语
        key: 'ko_KR',
        disabled: false
    },
    {
        label: 'Français',  // 法语
        key: 'fr_FR',
        disabled: false
    },
    {
        label: 'suomalainen', // 芬兰语
        key: 'fi_FI',
        disabled: false
    },
    {
        label: 'Русский', // 俄语
        key: 'ru_RU',
        disabled: false
    },
    {
        label: 'Deutsch', // 德语
        key: 'de_DE',
        disabled: false
    },
]