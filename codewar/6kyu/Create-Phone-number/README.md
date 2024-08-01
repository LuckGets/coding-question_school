# Create Phone Number

Write a function that accepts an array of 10 integers (between 0 and 9), that returns a string of those numbers in the form of a phone number.

### Example

```
createPhoneNumber([1, 2, 3, 4, 5, 6, 7, 8, 9, 0]) // => returns "(123) 456-7890"
```

The returned format must be correct in order to complete this challenge.

Don't forget the space after the closing parentheses!

### Solution

```javascript
// My solution
function createPhoneNumber<T>(number: T[]): string {
  const firstThree: T[] = number.splice(0, 3);
  const secondThree: T[] = number.splice(0, 3);
  return `(${firstThree.join("")}) ${secondThree.join("")}-${number.join("")}`;
}
```

```javascript
// Pro solution
function cleverCreatePhone(number: number[]): string {
  let format: string = "(xxx) xxx-xxxx";

  for (let i = 0; i < number.length; i++) {
    format = format.replace("x", number[i].toString());
  }
  return format;
}
```

> The Pro solution make a terrible performance because it loop 2 time
> 1st loop is simple for loop
> 2nd loop is replace which loop the first string itfound
> which make BigO of this solution is n^2
> but it's actually fine for this solution because the maximum length of the number array is 10. So theworst or BigO time of this solution is 100

```typescript
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
```
