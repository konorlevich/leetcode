<?php

declare(strict_types=1);

namespace Konorlevich\Leetcode\common\ListNode;

class ListNode
{
    public int $val = 0;
    public ?ListNode $next = null;

    function __construct(int $val = 0, ?ListNode $next = null)
    {
        $this->val = $val;
        $this->next = $next;
    }

    /**
     * @param Integer[] $list
     * @return ?ListNode
     */
    public static function fromList(array $list): ?ListNode
    {
        $node = null;
        for ($i = 0; $i < count($list) ; $i++) {
            $node = new ListNode($list[$i], $node);
        }
        return $node;
    }
}
