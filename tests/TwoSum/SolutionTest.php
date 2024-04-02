<?php

declare(strict_types=1);

namespace TwoSum;

use Konorlevich\Leetcode\TwoSum\Solution;
use PHPUnit\Framework\Attributes\CoversClass;
use PHPUnit\Framework\Attributes\DataProvider;
use PHPUnit\Framework\TestCase;

#[CoversClass(Solution::class)] class SolutionTest extends TestCase
{
    static function dataProvider(): array
    {
        return [
            "empty" => [
                "nums" => [],
                "target" => 0,
                "want" => [],
            ],
            "[2,7,11,15]" => [
                "nums" => [2, 7, 11, 15],
                "target" => 9,
                "want" => [0, 1],
            ],
            "[3,2,4]" => [
                "nums" => [3, 2, 4],
                "target" => 6,
                "want" => [1, 2],
            ],
            "[3,3]" => [
                "nums" => [3, 3],
                "target" => 6,
                "want" => [0, 1],
            ],

            //Example 1:
            //
            //Input: nums = [2,7,11,15], target = 9
            //Output: [0,1]
            //Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
            //Example 2:
            //
            //Input: nums = [3,2,4], target = 6
            //Output: [1,2]
            //Example 3:
            //
            //Input: nums = [3,3], target = 6
            //Output: [0,1]
        ];
    }

    /**
     * @param Integer[] $nums
     * @param Integer $target
     * @param Integer[] $want
     * @return void
     */
    #[DataProvider('dataProvider')] public function testTwoSum(array $nums, int $target, array $want)
    {
        $this->assertEquals($want, (new Solution())->twoSum($nums, $target));
    }
}
