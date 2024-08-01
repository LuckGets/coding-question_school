// My solution
function createPhoneNumber<T>(number: T[]): string {
  const firstThree: T[] = number.splice(0, 3);
  const secondThree: T[] = number.splice(0, 3);
  return `(${firstThree.join("")}) ${secondThree.join("")}-${number.join("")}`;
}

// Pro solution
function cleverCreatePhone(number: number[]): string {
  let format: string = "(xxx) xxx-xxxx";

  for (let i = 0; i < number.length; i++) {
    format = format.replace("x", number[i].toString());
  }
  return format;
}
// But this solution make a terrible performance because it loop 2 time
//  First loop is simple for loop
// Second loop is replace which loop the first string it found
//  which make BigO of this solution is n^2
// but it's actually fine for this solution because the maximum length of the number array is 10. So the worst or BigO time of this solution is 100

// My solution but cleaner
function cleanerCreatePhoneNumber(numbers: number[]): string {
  const numbersStr: string = numbers.join("");
  return (
    "(" +
    numbersStr.substring(0, 3) +
    ") " +
    numbersStr.substring(3, 6) +
    "-" +
    numbersStr.substring(6)
  );
}
