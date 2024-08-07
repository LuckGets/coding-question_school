function isPalindrome(s: string): boolean {
  // Check if string is " " or not, if it's then return true.
  if (s.trim() === "") {
    return true;
  }
  // Regex to check if the character is a string or num.
  const regex: RegExp = /^[A-Za-z0-9]+$/;

  // Declare stack which to push char in.
  const stack: Array<string> = [];

  // Loop to put char in stack
  for (let i = 0; i < s.length; i++) {
    // If special char, continue
    if (!regex.test(s[i])) {
      continue;
    }
    // push char to stack
    const strToPush = s[i].toLowerCase();
    stack.push(strToPush);
  }
  // Reverse the string in stack
  const reverStack = [...stack].reverse();

  // Test if it's palindrome or not.
  return stack.join("") === reverStack.join("") ? true : false;
}

function isPalindromeWithTwoPointer(s: string): boolean {
  // Check if string is " " or not, if it's then return true.
  if (s.trim() === "") {
    return true;
  }

  // Regex to check if the character is a string or num.
  const regex: RegExp = /^[A-Za-z0-9]+$/;

  // Declare two pointer which is first and last.
  let first: number = 0;
  let last: number = s.length - 1;

  // loop for two pointer
  while (first < last) {
    const currFirst = s.charAt(first);
    const currLast = s.charAt(last);

    if (!regex.test(currFirst)) {
      // If special char, continue
      first++;
    } else if (!regex.test(currLast)) {
      // If special char, continue
      last--;
    } else if (currFirst.toLowerCase() === currLast.toLowerCase()) {
      // If same char, continue
      first++;
      last--;
    } else {
      // If it's not the same char, then it's not palindrome. Return fasle.
      return false;
    }
  }
  // If the loop finished, then there is no err in loop. So it's palindrome
  return true;
}
