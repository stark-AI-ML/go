const input =
  "123 34 78 97 109 101 34 58 34 65 108 105 99 101 34 44 34 66 111 100 121 34 58 34 72 101 108 108 111 34 44 34 84 105 109 101 34 58 49 50 57 52 55 48 54 51 57 53 56 56 49 53 52 55 48 48 48 125";

const buffer = Buffer.from(input.split(" ").map(Number));

const json = JSON.parse(buffer.toString("utf8"));

console.log(json);
