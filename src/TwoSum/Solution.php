<?php

declare(strict_types=1);

namespace Konorlevich\Leetcode\TwoSum;

/**
 * Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
 *
 * You may assume that each input would have exactly one solution, and you may not use the same element twice.
 *
 * You can return the answer in any order.
 *
 * Only one valid answer exists.
 */
class Solution
{

    /**
     * @param Integer[] $nums
     * @param Integer $target
     * @return Integer[]
     */
    function twoSum(array $nums, int $target): array
    {
        foreach ($nums as $k => $v) {
            foreach (array_slice($nums, $k + 1, null, true) as $k1 => $v1) {
                if (($v + $v1) === $target) {
                    return [$k, $k1];
                }
            }
        }
        return [];
    }
}