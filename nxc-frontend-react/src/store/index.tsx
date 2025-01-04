import {create} from "zustand";

import {useThemeSlice, type ThemeState} from "./slices/themeSlice";
import {useCountSlice, type CountState} from "./slices/countSlice";

type StoreState = ThemeState & CountState

const useStore = create<StoreState>((...set) => ({
    ...useThemeSlice(...set),
    ...useCountSlice(...set),
}));

export default useStore