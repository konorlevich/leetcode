<?php

declare(strict_types=1);

namespace common\ListNode;

use Konorlevich\Leetcode\common\ListNode\ListNode;
use PHPUnit\Framework\Attributes\CoversClass;
use PHPUnit\Framework\Attributes\DataProvider;
use PHPUnit\Framework\TestCase;

#[CoversClass(ListNode::class)]class ListNodeTest extends TestCase
{
    public static function dataProvider(): array
    {
        return [
            'empty' => [
                'list'=> [],
                'expected'=> null,
            ],
            '[1,2,3,4]' => [
                'list'=> [1,2,3,4],
                'expected'=> new ListNode(4,
                    new ListNode(3,
                    new ListNode(2,
                    new ListNode(1)
                    ))),
            ],
            '[4,3,2,1]' => [
                'list'=> [4,3,2,1],
                'expected'=> new ListNode(1,
                    new ListNode(2,
                    new ListNode(3,
                    new ListNode(4)
                    ))),
            ],
        ];
    }

    /**
     * @param array $list
     * @param ListNode|null $expected
     * @return void
     */
    #[DataProvider('dataProvider')] public function testFromList(
        array $list,
        ?ListNode $expected,
    )
    {
        $this->assertEquals($expected, ListNode::fromList($list));
    }
}
