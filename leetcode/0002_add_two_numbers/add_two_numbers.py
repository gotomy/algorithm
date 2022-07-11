from operator import ne
from typing import Optional


class ListNode:
    def __init__(self, val = 0, next = None):
       self.val = val
       self.next = next

class Solution:
    def addTwoNumbers(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
        pre = ListNode
        cur = pre
        carry = 0
        while l1 is not None or l2 is not None:
            x = l1.val if l1 else 0
            y = l2.val if l2 else 0
            sum = x + y + carry
            carry = sum // 10
            sum = sum % 10
            cur.next = ListNode(sum)
            cur = cur.next
            if l1 is not None:
                l1 = l1.next
            if l2 is not None:
                l2 = l2.next
        if carry == 1:
            cur.next = ListNode(1)

        return pre.next

p13 = ListNode(3)
p12 = ListNode(4, p13)
p1 = ListNode(2, p12)

p23 = ListNode(4)
p22 = ListNode(6, p23)
p2 = ListNode(5, p22)

solution = Solution()
ret = solution.addTwoNumbers(p1, p2)
while ret:
    print(ret.val)
    ret = ret.next
