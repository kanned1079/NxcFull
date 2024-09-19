import {marked} from 'marked';
import DOMPurify from 'dompurify';

export const convertMd = (mdText: string) => {
    let formatted = marked.parse(mdText);

    return DOMPurify.sanitize(formatted);
}