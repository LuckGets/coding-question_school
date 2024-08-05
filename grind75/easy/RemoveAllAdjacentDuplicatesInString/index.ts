function removeDuplicates(s: string): string {
  const stack: Array<string> = [];
  for (let i = 0; i < s.length; i++) {
    if (stack[stack.length - 1] === s.charAt(i) && i !== 0) {
      stack.pop();
    } else {
      stack.push(s.charAt(i));
    }
  }
  return stack.join("");
}
