import {type GlobalThemeOverrides} from "naive-ui"

export const milkGreenDayLight: GlobalThemeOverrides = {
    common: {
        primaryColor: '#33a79e',
        primaryColorPressed: '#33a79e',
        bodyColor: '#eff2f7',
        modalColor: '#eff2f7',
        tableColor: '#eff2f7',
        primaryColorHover: '#33a79e',
        borderRadius: '6px',
    },
    Button: {
        color: '#33a79e',
        colorFocusPrimary: '#29a08c',
        colorHoverPrimary: '#29a08c',
        colorPressedPrimary: '#29a08c',
        textColorPrimary: '#fff',
        textColorHoverPrimary: '#fff',
        textColorPressedPrimary: '#fff',
        textColorFocusPrimary: '#fff',
        colorTextPrimary: '#33a79e',
        colorTextHoverPrimary: '#29a08c',
        textColorTextHoverPrimary: '#29a08c',
        textColorPrimaryDisabled: '#fff',
    },
    Input: {
        borderHover: '1px solid #33a79e',
        borderFocus: '1px solid #33a79e',
    },
    Card: {
        colorEmbedded: '#fff',  // 比bodyColor略深的颜色
        borderRadius: '6px',
    },
    Switch: {
        railColorActive: '#33a79e',
    },
    DataTable: {
        borderRadius: '6px !important',
    },
    Notification: {
        color: 'rgba(255, 255, 255, 1)',
        bordered: '1px solid #33a79e',
    },
    Dropdown: {
        color: '#fff',
        borderFocus: '1px solid #33a79e',
        borderHover: '1px solid #33a79e',
    },
    Select: {
        borderFocus: '1px solid #33a79e',
        borderHover: '1px solid #33a79e',
    },
    Carousel: {
        borderRadius: '6px',
    },
    Tabs: {
        tabBorderRadius: '6px',
    },
}

export const milkGreenDayNight: GlobalThemeOverrides = {
    common: {
        primaryColor: '#006f61',
        primaryColorPressed: '#006f61',
        bodyColor: '#2d2f2f',
        modalColor: '#2d2f2f',
        tableColorStriped: 'rgba(40, 41, 41, 1)',
        tableColor: 'rgba(40, 41, 41, 1)',
        tableHeaderColor: 'rgba(40, 41, 41, 1)',
        tableColorHover: 'rgba(40, 41, 41, 0.5)',
        primaryColorHover: '#006f61',
        borderRadius: '6px',
    },
    Button: {
        color: '#006f61',
        colorFocusPrimary: '#2c857a',
        colorHoverPrimary: '#2c857a',
        colorPressedPrimary: '#2c857a',
        textColorPrimary: '#fff',
        textColorHoverPrimary: '#fff',
        textColorPressedPrimary: '#fff',
        textColorFocusPrimary: '#fff',
        textColorError: '#ffe3e3',
        textColorHoverError: '#ffe3e3',
        textColorPressedError: '#ffe3e3',
        textColorFocusError: '#ffe3e3',
        textColorTextPrimary: '#fff',
        textColorTextHoverPrimary: '#2c857a',
    },
    Input: {
        borderHover: '1px solid #006f61',
        borderFocus: '1px solid #006f61',
    },
    Card: {
        colorEmbedded: 'rgba(40, 41, 41, 1)',
        borderRadius: '6px',
    },
    Switch: {
        railColorActive: '#006f61',
    },
    DataTable: {
        thColor: 'rgba(40, 41, 41, 1)',
        tdColor: 'rgba(40, 41, 41, 1)',
        borderRadius: '6px !important',
    },
    Notification: {
        color: 'rgba(40, 45, 47, 1)',
    },
    Menu: {
        itemTextColorActive: '#33a79e',
        itemTextColorActiveHover: '#008784',
        itemIconColorActive: '#33a79e',
        itemIconColorActiveHover: '#008784',
        itemTextColorChildActive: '#33a79e',
        itemTextColorChildActiveHover: '#008784',
        arrowColorChildActive: '#33a79e',
        arrowColorChildActiveHover: '#008784',
    },
    Dropdown: {
        color: '#252525',
    },
    Carousel: {
        borderRadius: '6px',
    },
    Tabs: {
        tabBorderRadius: '6px',
    },
}