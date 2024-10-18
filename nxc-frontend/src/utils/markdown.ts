import {marked} from 'marked';
// import DOMPurify from 'dompurify';

export const convertMd = (mdText: string) => {
    // return DOMPurify.sanitize(formatted);
    return marked.parse(mdText)
}