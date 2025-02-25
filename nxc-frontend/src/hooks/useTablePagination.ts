import {ref, type Ref} from "vue"
import useAppInfosStore from "@/stores/useAppInfosStore";

type DataSize = {
    page: number,
    pageSize: number
}

type UseAppInfosStore = [
    dataSize: Ref<DataSize>,
    dataSize: Ref<number>,
    setDataSize: (page: number, pageSize: number) => void,
    setPageCount: (count: number) => void,
]

const useTablePagination: () => UseAppInfosStore =  () => {
    const appInfoStore = useAppInfosStore()

    const dataSize = ref<DataSize>({
        page: appInfoStore.defaultTablePagination.page,
        pageSize: appInfoStore.defaultTablePagination.size,
    })
    const pageCount = ref<number>(0)

    const setDataSize = (page: number, size: number) => {
        dataSize.value.page = page
        dataSize.value.pageSize = size
    }

    const setPageCount = (count: number) => {
        pageCount.value = count
    }

    return [dataSize, pageCount, setDataSize, setPageCount]
}

export default useTablePagination