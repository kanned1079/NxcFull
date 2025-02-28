import { ref, type Ref } from "vue";

// 定义网络请求状态的类型
export type IsNetRequestedSuccess = 'waiting' | 'finished' | 'err';

// 定义 useNetworkResultHook 函数的返回类型
type UseNetworkResultHook = () => [
    isNetRequestedSuccess: Ref<IsNetRequestedSuccess>,
    setIsNetRequestedSuccess: (status: IsNetRequestedSuccess) => void
];

// 修正后的 useNetworkResultHook 函数
const useNetworkResultHook: UseNetworkResultHook = () => {
    // 使用 ref 创建响应式的网络请求状态
    const isNetRequestedSuccess = ref<IsNetRequestedSuccess>('waiting');
    // 定义更新网络请求状态的函数
    const setIsNetRequestedSuccess = (status: IsNetRequestedSuccess) => {
        console.log('setIsNetRequestedSuccess', status);
        isNetRequestedSuccess.value = status;
    };

    // 返回包含状态和更新函数的数组
    return [isNetRequestedSuccess, setIsNetRequestedSuccess];
};

export default useNetworkResultHook;