"use strict";
// ## ---- Reverse words ---- ##
// Complete the function that accepts a string parameter, and reverses each word in the string. All spaces in the string should be retained.
// ### Example
// "This is an example!" ==> "sihT si na !elpmaxe"
// "double  spaces"      ==> "elbuod  secaps"
// ---- RECURSIVE VARIATION -------//
function reverseStr(str) {
    if (str === "") {
        return "";
    }
    else {
        return reverseStr(str.substring(1)) + str.charAt(0);
    }
}
function reverseWord(str, callback) {
    const arr = str.split(" ");
    const newArr = arr.map(item => callback(item));
    return newArr.join(" ");
}
// Using recursive and try using dependency injection
// ------------------------------------------------------- //
// ---------- CLEAN CODE VARIATION ------------
function reverseString(str) {
    return str.split(" ").map(function (item) {
        return item.split("").reverse().join();
    }).join(" ");
}
