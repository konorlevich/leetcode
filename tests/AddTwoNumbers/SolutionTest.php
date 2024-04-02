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
            '0+0' => [
                'l1' => ListNode::fromList([0]),
                'l2' => ListNode::fromList([0]),
                'expected' => ListNode::fromList([0]),
            ],
            '9+9' => [
                'l1' => ListNode::fromList([9]),
                'l2' => ListNode::fromList([9]),
                'expected' => ListNode::fromList([1,8]),
            ],
            '123+12' => [
                'l1' => ListNode::fromList([1,2,3]),
                'l2' => ListNode::fromList([1,2]),
                'expected' => ListNode::fromList([1,3,5]),
            ],
            '243+564' => [
                'l1' => ListNode::fromList([3,4,2]),
                'l2' => ListNode::fromList([4,6,5]),
                'expected' => ListNode::fromList([8,0,7]),
            ],
            '343+564' => [
                'l1' => ListNode::fromList([3,4,3]),
                'l2' => ListNode::fromList([5,6,4]),
                'expected' => ListNode::fromList([9,0,7]),
            ],
            '9999999+9999' => [
                'l1' => ListNode::fromList([9,9,9,9,9,9,9]),
                'l2' => ListNode::fromList([9,9,9,9]),
                'expected' => ListNode::fromList([1,0,0,0,9,9,9,8]),
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
