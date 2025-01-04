import {create} from "zustand";

import {useThemeSlice, type ThemeState} from "./slices/themeSlice";

type StoreState = ThemeState

const useStore = create<StoreState>((...set) => ({
    ...useThemeSlice(...set),
}));

export default useStore