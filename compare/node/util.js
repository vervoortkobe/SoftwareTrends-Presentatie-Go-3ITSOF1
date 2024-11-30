function generateRandomArray(size) {
  return Array.from({ length: size }, () =>
    Math.floor(Math.random() * 100_000_000)
  );
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

module.exports = {
  generateRandomArray,
  printArraySample,
  getMilliseconds,
  padRight,
};
