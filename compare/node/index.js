const counter = require("./counter");
const fibonacci = require("./fibonacci");
const quickSort = require("./quicksort");
const bubbleSort = require("./bubblesort");
const {
  generateRandomArray,
  printArraySample,
  padRight,
  getMilliseconds,
} = require("./util");

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
