<?php

declare(strict_types=1);

namespace IntersectionOfTwoArraysII;

use Konorlevich\Leetcode\IntersectionOfTwoArraysII\Solution;
use PHPUnit\Framework\Attributes\CoversClass;
use PHPUnit\Framework\Attributes\DataProvider;
use PHPUnit\Framework\TestCase;

#[CoversClass(Solution::class)]
class SolutionTest extends TestCase
{
    /**
     * @return array
     */
    public static function dataProvider(): array
    {
        return [
            'empty' => [
                'nums1' => [],
                'nums2' => [],
                'expected' => [],
            ],
            'nums1 = [1,2,2,1], nums2 = [2,2]' => [
                'nums1' => [1,2,2,1],
                'nums2' => [2,2],
                'expected' => [2,2],
            ],
            'nums1 = [4,9,5], nums2 = [9,4,9,8,4]' => [
                'nums1' => [4,9,5],
                'nums2' => [9,4,9,8,4],
                'expected' => [4,9],
            ],
        ];
    }

    #[DataProvider('dataProvider')]
    public function testIntersect(array $nums1,array $nums2,array $expected)
    {
        foreach (Solution::METHODS as $method){
            $this->assertEquals($expected, (new Solution())->$method($nums1, $nums2), "$method failed");
        }
    }
}
