function counter(amount) {
  let sum = 0;
  for (let i = 0; i < amount; i++) {
    sum++;
  }
  return sum;
}

function fibonacci(n) {
  if (n <= 1) return n;
  let prev = 0,
    curr = 1;
  for (let i = 2; i <= n; i++) {
    [prev, curr] = [curr, prev + curr];
  }
  return curr;
}

function quickSort(arr) {
  if (arr.length <= 1) return arr;

  const pivot = arr[arr.length - 1];
  const left = [];
  const right = [];

  for (let i = 0; i < arr.length - 1; i++) {
    if (arr[i] < pivot) {
      left.push(arr[i]);
    } else {
      right.push(arr[i]);
    }
  }

  return [...quickSort(left), pivot, ...quickSort(right)];
}

function bubbleSort(arr) {
  const result = [...arr];
  const n = result.length;
  for (let i = 0; i < n - 1; i++) {
    for (let j = 0; j < n - i - 1; j++) {
      if (result[j] > result[j + 1]) {
        [result[j], result[j + 1]] = [result[j + 1], result[j]];
      }
    }
  }
  return result;
}

function generateRandomArray(size) {
  return Array.from({ length: size }, () => Math.floor(Math.random() * 10000));
}

function printArraySample(arr, name) {
  if (arr.length <= 10) {
    return `${name}: ${JSON.stringify(arr)}`;
  }
  return `${name}: [${arr.slice(0, 5)}, ..., ${arr.slice(-5)}]`;
}

function getMilliseconds(hrtime) {
  return (hrtime[0] * 1000 + hrtime[1] / 1000000).toFixed(6);
}

function padRight(str, length) {
  return str.padEnd(length);
}

console.log("NodeJS Benchmark Results:");
console.log("------------------------");

// Counter benchmark
const counterStart = process.hrtime();
const sum = counter(100000);
const counterTime = getMilliseconds(process.hrtime(counterStart));
console.log(
  `${padRight(`1. Counter: ${sum}`, 70)} ${counterTime.padStart(10)} ms`
);

// Fibonacci benchmark
const fibStart = process.hrtime();
const fibResult = fibonacci(100);
const fibTime = getMilliseconds(process.hrtime(fibStart));
console.log(
  `${padRight(`2. Fibonacci: ${fibResult}`, 70)} ${fibTime.padStart(10)} ms`
);

// QuickSort benchmark
const quickSortArr = generateRandomArray(1000);
console.log("3. Quicksort:");
console.log(printArraySample(quickSortArr, "   - Input"));

const quickSortStart = process.hrtime();
const sortedQuick = quickSort(quickSortArr);
const quickSortTime = getMilliseconds(process.hrtime(quickSortStart));

console.log(
  `${padRight(
    printArraySample(sortedQuick, "   - Output"),
    70
  )} ${quickSortTime.padStart(10)} ms`
);

// BubbleSort benchmark
const bubbleSortArr = generateRandomArray(1000);
console.log("4. Bubblesort:");
console.log(printArraySample(bubbleSortArr, "   - Input"));

const bubbleSortStart = process.hrtime();
const sortedBubble = bubbleSort(bubbleSortArr);
const bubbleSortTime = getMilliseconds(process.hrtime(bubbleSortStart));

console.log(
  `${padRight(
    printArraySample(sortedBubble, "   - Output"),
    70
  )} ${bubbleSortTime.padStart(10)} ms`
);
