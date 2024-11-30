function fibonacci(n) {
  if (n <= 1) {
    return BigInt(n);
  }
  let prev = BigInt(0);
  let curr = BigInt(1);
  for (let i = 2; i <= n; i++) {
    const temp = curr;
    curr = prev + curr;
    prev = temp;
  }
  return curr;
}

module.exports = fibonacci;
