import {ref} from "vue"
import useAppInfosStore from "@/stores/useAppInfosStore";

interface DataSize {
    page: number,
    pageSize: number
}

const useTablePagination = () => {
    const appInfoStore = useAppInfosStore()

    const dataSize = ref<DataSize>({
        page: appInfoStore.defaultTablePagination.page,
        pageSize: appInfoStore.defaultTablePagination.size,
    })
    const pageCount = ref<number>(0)

    return [dataSize, pageCount]
}

export default useTablePagination