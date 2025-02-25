import {useDialog} from "naive-ui"
import {useI18n} from "vue-i18n";

type UseConfirmDialog = {
    makeSuccessDialog: (title: string, content: string, effectClick: () => void) => void;
}

const makeSuccessDialog = (title: string, content: string, effectClick: () => void) => {
    const {t} = useI18n()
    useDialog().success({
        title: t('adminViews.common.dialog.delete'),
        content: t('adminViews.planMgr.mention.common.delMention'),
        positiveText: t('adminViews.common.confirm'),
        negativeText: t('adminViews.common.cancel'),
        showIcon: false,
        actionStyle: {
            marginTop: '20px',
        },
        contentStyle: {
            marginTop: '20px',
        },
        onPositiveClick: () => effectClick(),
    })
}

const useConfirmDialog: () => UseConfirmDialog = () => {
    return {
        makeSuccessDialog,
    }
}

export default useConfirmDialog