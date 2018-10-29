/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */

struct ListNode {
     int val;
     ListNode *next;
     ListNode(int x) : val(x), next(nullptr) {}
 };

class Solution {
public:
    void deleteNode(ListNode* node) {
        if (node == nullptr) {
            return;
        }

        if (node->next == nullptr) {
            return;
        }

        int val = node->next->val;
        ListNode *tmp_ptr = node->next->next;

        node->val = val;
        delete(node->next);
        node->next = tmp_ptr;
    }
};
