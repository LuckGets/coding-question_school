"use strict";
function findShort(str) {
    let length = 0;
    str.split(" ").forEach(item => {
        if (item.length < length) {
            console.log("item",item.length)
            length = item.length;
        }
    });
    console.log("final length", length)
    return length;
}

findShort("hello there mfk")