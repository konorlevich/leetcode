<?php

declare(strict_types=1);

namespace Konorlevich\Leetcode\AddTwoNumbers;

use Konorlevich\Leetcode\common\ListNode\ListNode;

/**
 * You are given two non-empty linked lists representing two non-negative integers.
 * The digits are stored in reverse order, and each of their nodes contains a single digit.
 * Add the two numbers and return the sum as a linked list.
 *
 * You may assume the two numbers do not contain any leading zero, except the number 0 itself.
 *
 * It is guaranteed that the list represents a number that does not have leading zeros.
 */
class Solution
{

    /**
     * @param ?ListNode $l1
     * @param ?ListNode $l2
     * @return ?ListNode
     */
    function addTwoNumbers(?ListNode $l1, ?ListNode $l2): ?ListNode
    {
        if ($l1 === null) {
            return $l2;
        }
        if ($l2 === null) {
            return $l1;
        }
        $l3 = new ListNode($l1->val + $l2->val);
        if ($l3->val >= 10) {
            $l3->next = new ListNode(1);
            $l3->val -= 10;
        }
        $l3->next = $this->addTwoNumbers($l3->next, $this->addTwoNumbers($l1->next, $l2->next));
        return $l3;
    }
}