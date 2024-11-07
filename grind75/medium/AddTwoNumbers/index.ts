class ListNode {
  val: number;
  next: ListNode | null;
  constructor(val?: number, next?: ListNode | null) {
    this.val = val === undefined ? 0 : val;
    this.next = next === undefined ? null : next;
  }
}

function addTwoNumbers(
  l1: ListNode | null,
  l2: ListNode | null
): ListNode | null {
  return sumTheTwoNodeList(l1, l2, 0);
}

function sumTheTwoNodeList(
  l1: ListNode | null,
  l2: ListNode | null,
  carry: number
): ListNode | null {
  if (l1 === null && l2 === null && carry === 0) {
    return null;
  }
  let sum = carry;
  const current = new ListNode();
  if (l1 !== null) {
    sum += l1.val;
    l1 = l1.next;
  }

  if (l2 !== null) {
    sum += l2.val;
    l2 = l2.next;
  }
  current.val = Math.trunc(sum % 10);
  current.next = sumTheTwoNodeList(l1, l2, Math.trunc(sum / 10));
  return current;
}

// function addTwoNumbers(l1: ListNode | null, l2: ListNode | null): ListNode | null {
//     let carry : number = 0
//     const dummy :ListNode = new ListNode()
//     let current : ListNode = dummy

//     do {
//         let sum : number = carry

//         if (l1 !== null) {
//             sum += l1.val
//             l1 = l1.next
//         }
//         if (l2 !== null) {
//             sum += l2.val
//             l2 = l2.next
//         }

//         carry = Math.trunc(sum / 10)
//         current.next = new ListNode(Math.trunc(sum % 10))
//         current = current.next
//     } while(l1 !== null || l2 !== null || carry > 0)
//     return dummy.next
// };
