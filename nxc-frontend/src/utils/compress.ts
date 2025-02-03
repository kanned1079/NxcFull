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
 * 裁剪并调整图像大小的函数。
 *
 * 该函数会从指定的 `HTMLImageElement` 中裁剪出一部分区域，并将其缩放到目标大小后返回为 `Blob` 格式的图片数据。
 *
 * @param {HTMLImageElement} image - 要裁剪的图像对象。
 * @param {number} x - 裁剪区域的左上角 X 坐标（相对于原始图像）。
 * @param {number} y - 裁剪区域的左上角 Y 坐标（相对于原始图像）。
 * @param {number} cropSize - 裁剪区域的宽度和高度（正方形）。
 * @param {number} outputSize - 输出图像的最终宽度和高度（缩放后，仍为正方形）。
 * @returns {Promise<Blob>} 返回一个 `Promise`，解析为包含裁剪和缩放后图像数据的 `Blob`。
 *
 * @throws {Error} 当无法获取 `CanvasRenderingContext2D` 时，会触发错误。
 * @throws {Error} 当 `canvas.toBlob` 失败时，会触发错误。
 *
 * @example
 * const image = new Image();
 * image.src = 'example.jpg';
 * image.onload = async () => {
 *     try {
 *         const croppedBlob = await cropAndResizeImage(image, 50, 50, 100, 150);
 *         console.log('裁剪后的图片 Blob:', croppedBlob);
 *     } catch (error) {
 *         console.error('裁剪失败:', error);
 *     }
 * };
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

        if (ctx) {  // 确保创建的元素存在
            ctx.drawImage(image, x, y, cropSize, cropSize, 0, 0, outputSize, outputSize);
            canvas.toBlob(blob => {
                if (blob) resolve(blob);
                else reject(new Error('Canvas create Blob failure.'));
            }, 'image/jpeg', 0.9);
        } else {
            reject(new Error('cannot get context of Canvas.'));
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

        console.log('compressed size:', (compressedFile.size / 1024).toFixed(2) + ' KB');
        return compressedFile;
    } catch (error) {
        console.error('resize and compress:', error);
        throw new Error('img process failure, please try again');
    }
};