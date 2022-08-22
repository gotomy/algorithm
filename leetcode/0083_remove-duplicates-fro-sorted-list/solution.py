from operator import ne
from typing import Optional

from typing import Optional
from typing import List

class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

def node_print(head):
    if head is None:
        return
    cur = head
    while cur:
        print(cur.val)
        cur = cur.next

def create(nodes: List[int]) -> ListNode:
    if len(nodes) == 0:
        return None
    head = ListNode(nodes[0])
    cur = head
    for val in nodes[1:]:
        cur.next = ListNode(val)
        cur = cur.next
    return head



class Solution:
    def deleteDuplicates(self, head: Optional[ListNode]) -> Optional[ListNode]:
        cur = head

        while cur:
            right = cur.next
            if (right is None):
                break
            if (cur.val == right.val):
                right = right.next
                cur.next = right
            else:
                cur = right
            
        node_print(head)
        return head


head = create([])
solution = Solution()
solution.deleteDuplicates(head)
