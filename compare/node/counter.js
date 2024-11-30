function counter(amount) {
  let sum = 0;
  for (let i = 0; i < amount; i++) {
    sum++;
  }
  return sum;
}

module.exports = counter;
