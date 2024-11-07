# 2.Add Two Numbers

You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

### Example

Example 1.
![Example1 picture](https://assets.leetcode.com/uploads/2020/10/02/addtwonumber1.jpg)

> **Input**: l1 = [2,4,3], l2 = [5,6,4] > **Output**: [7,0,8] > **Explanation**: 342 + 465 = 807.

Example 2.

> **Input**:l1 = [0], l2 = [0] > **Output**: [0]

Example 3.

> **Input**:l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]

> **Output**: [8,9,9,9,0,0,0,1]

Constraints

-       The number of nodes in each linked list is in the range [1, 100].
-       0 <= Node.val <= 9
-       It is guaranteed that the list represents a number that does not have leading zeros.

### Solution

#### Explanation

REF : (Rahul Varma from leetcode)[https://leetcode.com/problems/add-two-numbers/solutions/3675747/beats-100-c-java-python-beginner-friendly/]

##### Intuition:

The Intuition is to iterate through two linked lists representing non-negative integers in reverse order, starting from the least significant digit. It performs digit-wise addition along with a carry value and constructs a new linked list to represent the sum. The process continues until both input lists and the carry value are exhausted. The resulting linked list represents the sum of the input numbers in the correct order.

##### Explanation:

1. Create a placeholder node called dummyHead with a value of 0. This node will hold the resulting linked list.
2. Initialize a pointer called tail and set it to dummyHead. This pointer will keep track of the last node in the result list.
3. Initialize a variable called carry to 0. This variable will store the carry value during addition.
4. Start a loop that continues until there are no more digits in both input lists (l1 and l2) and there is no remaining carry value.
5. Inside the loop:

- Check if there is a digit in the current node of l1. If it exists, assign its value to a variable called digit1. Otherwise, set digit1 to 0.
- Check if there is a digit in the current node of l2. If it exists, assign its value to a variable called digit2. Otherwise, set digit2 to 0.
- Add the current digits from l1 and l2, along with the carry value from the previous iteration, and store the sum in a variable called sum.
- Calculate the unit digit of sum by taking the modulus (%) of sum by 10. This digit will be placed in a new node for the result.
- Update the carry variable by dividing sum by 10 and taking the integer division (/) part. This gives us the carry value for the next iteration.
- Create a new node with the calculated digit as its value.
- Attach the new node to the tail node of the result list.
- Move the tail pointer to the newly added node.
- Move to the next nodes in both l1 and l2, if they exist. If either list is exhausted, set the corresponding pointer to nullptr.

6. After the loop, obtain the actual result list by skipping the dummyHead node.
7. Return the resulting list.

##### While loop

###### GOlang

```go
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	carry := 0
	current := dummy

	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum = l2.Val
			l2 = l2.Next
		}

		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
		carry = sum / 10
	}
	return dummy.Next
}
```

###### typescript

```typescript
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
  let carry: number = 0;
  const dummy: ListNode = new ListNode();
  let current: ListNode = dummy;

  do {
    let sum: number = carry;

    if (l1 !== null) {
      sum += l1.val;
      l1 = l1.next;
    }
    if (l2 !== null) {
      sum += l2.val;
      l2 = l2.next;
    }

    carry = Math.trunc(sum / 10);
    current.next = new ListNode(Math.trunc(sum % 10));
    current = current.next;
  } while (l1 !== null || l2 !== null || carry > 0);
  return dummy.next;
}
```

##### Recursion Solution

###### GOlang

```go
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return sumTwoListNode(l1, l2, 0)
}

func sumTwoListNode(l1 *ListNode, l2 *ListNode, carry int) *ListNode {
	if l1 == nil && l2 == nil && carry == 0 {
		return nil
	}

	sum := carry
	current := &ListNode{}

	if l1 != nil {
		sum += l1.Val
		l1 = l1.Next
	}

	if l2 != nil {
		sum += l2.Val
		l2 = l2.Next
	}

	current.Val = sum % 10
	current.Next = sumTwoListNode(l1, l2, sum/10)
	return current
}
```

###### Typescript

```typescript
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
}unction addTwoNumbers(
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
```

### BigO : O(n)
