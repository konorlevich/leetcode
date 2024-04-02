<?php

declare(strict_types=1);

namespace Konorlevich\Leetcode\AddTwoNumbers;

use Konorlevich\Leetcode\common\ListNode\ListNode;

/**
 * Definition for a singly-linked list.
 * class ListNode {
 *     public $val = 0;
 *     public $next = null;
 *     function __construct($val = 0, $next = null) {
 *         $this->val = $val;
 *         $this->next = $next;
 *     }
 * }
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
        if ($l3->val > 10) {
            $l3->next = new ListNode(1);
            $l3->val -= 10;
        }
        $l3->next = $this->addTwoNumbers($l3->next, $this->addTwoNumbers($l1->next, $l2->next));
        return $l3;
    }
}