import {type GlobalThemeOverrides} from "naive-ui"

export const glacierBlueLight: GlobalThemeOverrides = {
    common: {
        primaryColor: '#6390b9',
        primaryColorPressed: '#6390b9',
        bodyColor: '#eff2f7',
        modalColor: '#eff2f7',
        tableColor: '#eff2f7',
        primaryColorHover: '#6390b9',
        borderRadius: '6px',
    },
    Button: {
        color: '#6390b9',
        colorFocusPrimary: '#5380a3',
        colorHoverPrimary: '#5380a3',
        colorPressedPrimary: '#5380a3',
        textColorPrimary: '#fff',
        textColorHoverPrimary: '#fff',
        textColorPressedPrimary: '#fff',
        textColorFocusPrimary: '#fff',
        colorTextPrimary: '#6390b9',
        colorTextHoverPrimary: '#5380a3',
        textColorTextHoverPrimary: '#5380a3',
        textColorPrimaryDisabled: '#fff',
    },
    Input: {
        borderHover: '1px solid #6390b9',
        borderFocus: '1px solid #6390b9',
    },
    Card: {
        colorEmbedded: '#fff',  // 比bodyColor略深的颜色
        borderRadius: '6px'
    },
    Switch: {
        railColorActive: '#6390b9'
    },
    DataTable: {
        borderRadius: '6px !important'
    },
    Notification: {
        color: 'rgba(255, 255, 255, 1)',
        bordered: '1px solid #6390b9',
    },
    Dropdown: {
        color: '#fff',
        borderFocus: '1px solid #6390b9',
        borderHover: '1px solid #6390b9',
    },
    Select: {
        borderFocus: '1px solid #6390b9',
        borderHover: '1px solid #6390b9',
    },
    Carousel: {
        borderRadius: '6px'
    },
    Tabs: {
        tabBorderRadius: '6px'
    }
}

export const glacierBlueNight: GlobalThemeOverrides = {
    // 深色模式
    common: {
        primaryColor: '#4f7698',
        primaryColorPressed: '#4f7698',
        bodyColor: '#2d2e2f',
        modalColor: '#2d2e2f',
        tableColorStriped: 'rgba(40, 41, 41, 1)',
        tableColor: 'rgba(40, 41, 41, 1)',
        tableHeaderColor: 'rgba(40, 41, 41, 1)',
        tableColorHover: 'rgba(40, 41, 41, 0.5)',
        primaryColorHover: '#4f7698',
        borderRadius: '6px',
        popoverColor: 'rgb(66, 67, 68)',
    },
    Button: {
        color: '#4f7698',
        colorFocusPrimary: '#5d86aa',
        colorHoverPrimary: '#5d86aa',
        colorPressedPrimary: '#5d86aa',
        textColorPrimary: '#fff',
        textColorHoverPrimary: '#fff',
        textColorPressedPrimary: '#fff',
        textColorFocusPrimary: '#fff',
        textColorError: '#ffe3e3',
        textColorHoverError: '#ffe3e3',
        textColorPressedError: '#ffe3e3',
        textColorFocusError: '#ffe3e3',
        textColorTextPrimary: '#fff',
        textColorTextHoverPrimary: '#5d86aa',
    },
    Input: {
        borderHover: '1px solid #4f7698',
        borderFocus: '1px solid #4f7698',
    },
    Card: {
        colorEmbedded: 'rgba(40, 41, 41, 1)',
        borderRadius: '6px'
    },
    Switch: {
        railColorActive: '#4f7698'
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
        itemTextColorActive: '#79a7ca',
        itemTextColorActiveHover: '#6794b2',
        itemIconColorActive: '#79a7ca',
        itemIconColorActiveHover: '#6794b2',
        itemTextColorChildActive: '#79a7ca',
        itemTextColorChildActiveHover: '#6794b2',
        arrowColorChildActive: '#79a7ca',
        arrowColorChildActiveHover: '#6794b2'
    },
    Dropdown: {
        color: '#252525'
    },
    Carousel: {
        borderRadius: '6px'
    },
    Tabs: {
        tabBorderRadius: '6px'
    }
}

