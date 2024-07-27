// ## ---- Reverse words ---- ##
// Complete the function that accepts a string parameter, and reverses each word in the string. All spaces in the string should be retained.

// ### Example
// "This is an example!" ==> "sihT si na !elpmaxe"
// "double  spaces"      ==> "elbuod  secaps"

// ---- RECURSIVE VARIATION -------//
function reverseStr(str:string):string {
  if (str === "") {
    return ""
  } else {
    return reverseStr(str.substring(1)) + str.charAt(0)
  }
}

function reverseWord(str:string,callback:Function):string {
  const arr:Array<string> = str.split(" ")
  const newArr:Array<string> = arr.map(item => callback(item)) 
  return newArr.join(" ")
}

// Using recursive and try using dependency injection
// ------------------------------------------------------- //
// ---------- CLEAN CODE VARIATION ------------

function reverseString(str:string):string {
  return str.split(" ").map(function(item:string):string {
    return item.split("").reverse().join()
  }).join(" ")
}
