import imageCompression from 'browser-image-compression';

/**
 * 加载图片为 HTMLImageElement
 */
const loadImage = (file: File): Promise<HTMLImageElement> => {
    return new Promise((resolve, reject) => {
        const img = new Image();
        img.src = URL.createObjectURL(file);
        img.onload = () => resolve(img);
        img.onerror = reject;
    });
};

/**
 * 在 Canvas 上裁剪并调整大小
 */
const cropAndResizeImage = (
    image: HTMLImageElement,
    x: number, y: number,
    cropSize: number, outputSize: number
): Promise<Blob> => {
    return new Promise((resolve, reject) => {
        const canvas = document.createElement('canvas');
        canvas.width = outputSize;
        canvas.height = outputSize;
        const ctx = canvas.getContext('2d');

        if (ctx) {
            ctx.drawImage(image, x, y, cropSize, cropSize, 0, 0, outputSize, outputSize);
            canvas.toBlob(blob => {
                if (blob) resolve(blob);
                else reject(new Error('Canvas 生成 Blob 失败'));
            }, 'image/jpeg', 0.9);
        } else {
            reject(new Error('无法获取 Canvas 上下文'));
        }
    });
};

/**
 * 压缩并裁剪用户头像
 * @param file - 用户上传的图片文件
 * @param outputSize - 输出尺寸（宽高）
 * @returns {Promise<File>} - 返回裁剪压缩后的图片文件
 */
export const compressAvatar = async (file: File, outputSize = 150): Promise<File> => {
    try {
        // 读取图片为 HTMLImageElement
        const image = await loadImage(file);

        // 计算裁剪区域（居中裁剪）
        const cropSize = Math.min(image.width, image.height);
        const cropX = (image.width - cropSize) / 2;
        const cropY = (image.height - cropSize) / 2;

        // 裁剪并调整大小
        const croppedBlob = await cropAndResizeImage(image, cropX, cropY, cropSize, outputSize);

        // 将 Blob 转换为 File 类型
        const croppedFile = new File([croppedBlob], file.name, { type: file.type });

        // 压缩图片
        const compressedFile = await imageCompression(croppedFile, {
            maxSizeMB: 0.1, // 目标最大大小 100KB
            maxWidthOrHeight: outputSize,
            useWebWorker: true, // 使用 Web Worker 提高性能
            initialQuality: 0.8, // 初始质量
        });

        console.log('压缩后的图片大小:', (compressedFile.size / 1024).toFixed(2) + ' KB');
        return compressedFile;
    } catch (error) {
        console.error('图片裁剪和压缩失败:', error);
        throw new Error('图片处理失败，请重试');
    }
};