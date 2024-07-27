
// My variation
function findShort(str:string):number{
  let length:number = str.split(" ")[0].length;
  str.split(" ").forEach(item => {
    if (length < item.length) {
    length = item.length
}
  })
  return length
}

// Cleaner
function findShortCleaner(str:string):number {
  return Math.min(...str.split(" ").map(item => item.length))
}
