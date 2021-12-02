const measurements = require('./data');

function part1() {
  let inc = 0;
  for (let i = 0; i + 1 < measurements.length; i++) {
    if (measurements[i] < measurements[i + 1]) inc++;
  }
  return inc;
}

function part2() {
  let inc = 0;
  for (let i = 0; i + 3 < measurements.length; i++) {
    const prev = measurements.slice(i, i + 3).reduce((acc, curr) => acc + curr);
    const after = measurements.slice(i + 1, i + 4).reduce((acc, curr) => acc + curr);
    if (after > prev) inc++;
  }
  return inc;
}

console.log(part1());
console.log(part2());
