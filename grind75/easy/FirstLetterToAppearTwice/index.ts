function repeatChar(word: string): string {
  const arr: number[] = new Array(26).fill(0)
  const defaultChar = "a"
  for (let char of word) {
    const index = char.charCodeAt(0) - defaultChar.charCodeAt(0)
    if (arr[index]) {
      return char
    } else {
      arr[index] = 1
    }
  }
  throw new Error()
}
