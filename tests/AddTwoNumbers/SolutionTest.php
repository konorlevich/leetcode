<?php

namespace AddTwoNumbers;

use Konorlevich\Leetcode\AddTwoNumbers\Solution;
use Konorlevich\Leetcode\common\ListNode\ListNode;
use PHPUnit\Framework\Attributes\CoversClass;
use PHPUnit\Framework\Attributes\DataProvider;
use PHPUnit\Framework\Attributes\UsesClass;
use PHPUnit\Framework\TestCase;

#[CoversClass(Solution::class)]
#[UsesClass(ListNode::class)]
class SolutionTest extends TestCase
{

    /**
     * @return array
     */
     public static function dataProvider(): array
    {
        return [
            'empty' => [
                'l1' => null,
                'l2' => null,
                'expected' => null,
            ],
            '1+1' => [
                'l1' => ListNode::fromList([1]),
                'l2' => ListNode::fromList([1]),
                'expected' => ListNode::fromList([2]),
            ],
            '9+9' => [
                'l1' => ListNode::fromList([9]),
                'l2' => ListNode::fromList([9]),
                'expected' => ListNode::fromList([8,1]),
            ],
            '123+12' => [
                'l1' => ListNode::fromList([1,2,3]),
                'l2' => ListNode::fromList([1,2]),
                'expected' => ListNode::fromList([2,4,3]),
            ],
        ];
    }

    /**
     * @param ListNode|null $l1
     * @param ListNode|null $l2
     * @param ListNode|null $expected
     * @return void
     */
    #[DataProvider('dataProvider')]
    public function testAddTwoNumbers(?ListNode $l1, ?ListNode $l2, ?ListNode $expected)
    {
        $this->assertEquals($expected, (new Solution())->addTwoNumbers($l1, $l2));
    }
}
