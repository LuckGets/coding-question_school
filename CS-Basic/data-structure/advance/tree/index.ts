class NodeTree {
  val: string;
  left: NodeTree | null;
  right: NodeTree | null;
  constructor(val: string) {
    this.val = val;
    this.left = null;
    this.right = null;
  }
}

const a = new NodeTree("a");
const b = new NodeTree("b");
const c = new NodeTree("c");
const d = new NodeTree("d");
const e = new NodeTree("e");
const f = new NodeTree("f");

a.left = b;
a.right = c;
b.left = d;
b.right = e;
c.left = f;

// const depthFirstValue = (root: NodeTree) => {
//   const result: Array<string> = [];
//   const stack = [root];
//   while (stack.length > 0) {
//     const curr = stack.pop() as NodeTree;
//     result.push(curr.val);
//     if (curr.right) stack.push(curr.right);
//     if (curr.left) stack.push(curr.left);
//   }
//   return result;
// };

// Recursive Mode
const depthFirstValue = (root: NodeTree): Array<string> => {
  if (root === null) return [];
  const leftValue = depthFirstValue(root.left as NodeTree);
  const rightValue = depthFirstValue(root.right as NodeTree);
  return [root.val, ...leftValue, ...rightValue];
};

const breadthFirstValue = (root: NodeTree): Array<string> => {
  if (!root) return [];
  const result: Array<string> = [];
  const queue: Array<NodeTree> = [root];

  while (queue.length > 0) {
    const current = queue.shift() as NodeTree;
    result.push(current.val);

    if (current.left) queue.push(current.left);
    if (current.right) queue.push(current.right);
  }
  return result;
};

depthFirstValue(a);
