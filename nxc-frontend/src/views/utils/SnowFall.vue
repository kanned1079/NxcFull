<script lang="ts" setup>
import {onMounted, onUnmounted, ref, watch} from 'vue';
import * as confetti from 'canvas-confetti';

// Props
const props = defineProps({
  show: {
    type: Boolean,
    default: false,
  },
});

// Refs
const canvasRef = ref<HTMLCanvasElement | null>(null);
let confettiInstance: confetti.CreateTypes | null = null;

// 初始化 confetti 实例
const initConfetti = () => {
  if (canvasRef.value) {
    confettiInstance = confetti.create(canvasRef.value, {
      resize: true,
      useWorker: true,
    });
  }
};

// 雪花飘落动画
const fireSnowfall = () => {
  if (!confettiInstance) return;

  const duration = 3 * 1000; // 动画时长
  const animationEnd = Date.now() + duration;
  let skew = 1;

  const randomInRange = (min: number, max: number) => Math.random() * (max - min) + min;

  const frame = () => {
    const timeLeft = animationEnd - Date.now();
    const ticks = Math.max(100, 500 * (timeLeft / duration));
    skew = Math.max(0.8, skew - 0.001);

    confettiInstance!({
      zIndex: 1200,
      particleCount: 6, // 每帧生成一个粒子
      startVelocity: 0, // 初始速度
      ticks: ticks, // 粒子存活时间
      origin: {
        x: Math.random(), // 随机横向位置
        y: Math.random() * skew - 0.2, // 垂直起始位置偏移
      },
      colors: ['#ffffff'], // 雪花颜色
      shapes: ['circle'], // 雪花形状
      gravity: randomInRange(0.4, 1), // 随机重力
      scalar: randomInRange(0.2, 1.2), // 随机大小
      drift: randomInRange(-0.4, 0.4), // 横向漂移
    });

    if (timeLeft > 0) {
      requestAnimationFrame(frame);
    }
  };

  frame();
};

// 监听 `show` 属性，触发动画
watch(
    () => props.show,
    (newValue) => {
      if (newValue) {
        fireSnowfall();
      }
    },
);

// 生命周期钩子
onMounted(() => {
  initConfetti();
});

onUnmounted(() => {
  confettiInstance = null;
});
</script>

<script lang="ts">
export default {
  name: 'SnowFall',
};
</script>

<template>
  <canvas ref="canvasRef" class="confetti-canvas"></canvas>
</template>

<style scoped>
.confetti-canvas {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  z-index: 9999;
  pointer-events: none;
}
</style>