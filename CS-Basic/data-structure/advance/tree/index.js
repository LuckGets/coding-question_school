// var NodeTree = /** @class */ (function () {
//     function NodeTree(val) {
//         this.val = val;
//         this.left = null;
//         this.right = null;
//     }
//     return NodeTree;
// }());
// var a = new NodeTree("a");
// var b = new NodeTree("b");
// var c = new NodeTree("c");
// var d = new NodeTree("d");
// var e = new NodeTree("e");
// var f = new NodeTree("f");
// a.left = b;
// a.right = c;
// b.left = d;
// b.right = e;
// c.left = f;
// var depthFirstValue = function (root) {
//     var result = [];
//     var stack = [root];
//     while (stack.length > 0) {
//         var curr = stack.pop();
//         result.push(curr.val);
//         if (curr.right)
//             stack.push(curr.right);
//         if (curr.left)
//             stack.push(curr.left);
//     }
//     console.log(result);
// };
// depthFirstValue(a);
