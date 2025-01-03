import {create} from "zustand";

import createCounterStore from "./modules/counter";
import createChannelStore from "./modules/channel";

const useStore = create((...a) => {
    return {
        ...createCounterStore(...a),
        ...createChannelStore(...a),
    }
})

export default useStore