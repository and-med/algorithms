const DIGITS = 4;
const POWER = 5;

let f = (d, pos) => d * Math.pow(10, pos) - Math.pow(d, POWER);

console.log(
  [``.padEnd(22)]
    .concat(
      Array(10)
        .fill()
        .map((_, pos) => pos.toString().padEnd(9))
    )
    .join('')
);
console.log(
  Array(DIGITS)
    .fill()
    .map((_, pos) =>
      [`i * ${Math.pow(10, pos)} - i ^ ${POWER}`.padEnd(22)]
        .concat(
          Array(10)
            .fill()
            .map((_, d) => f(d, pos).toString().padEnd(9))
        )
        .join('')
    )
    .join('\n')
);
