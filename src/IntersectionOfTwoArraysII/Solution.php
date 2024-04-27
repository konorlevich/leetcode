<?php

declare(strict_types=1);

namespace Konorlevich\Leetcode\IntersectionOfTwoArraysII;

/**
 * Given two integer arrays nums1 and nums2, return an array of their intersection.
 * Each element in the result must appear as many times as it shows in both arrays,
 * and you may return the result in any order.
 *
 * Constraints:
 *
 * - 1 <= nums1.length, nums2.length <= 1000
 * - 0 <= nums1[i], nums2[i] <= 1000
 *
 *
 * Follow up:
 *
 * - What if the given array is already sorted? How would you optimize your algorithm?
 * - What if nums1's size is small compared to nums2's size? Which algorithm is better?
 * - What if elements of nums2 are stored on disk,
 * and the memory is limited such that you cannot load all elements into the memory at once?
 */
class Solution
{

    /**
     *
     */
    const array METHODS = [
        'bruteForce',
        'sortedArrays',
    ];

    /**
     * The easiest solution with O(n^2)
     *
     * @param Integer[] $nums1
     * @param Integer[] $nums2
     * @return Integer[]
     */
    public function bruteForce(array $nums1, array $nums2): array
    {
        $result = [];
        for ($i = 0; $i < count($nums1);) {
            for ($j = 0; $j < count($nums2);) {
                if ($nums1[$i] !== $nums2[$j]) {
                    $j++;
                    continue;
                }
                $result[] = $nums1[$i];
                unset($nums1[$i]);
                $nums1 = array_values($nums1);

                unset($nums2[$j]);
                $nums2 = array_values($nums2);
                continue 2;
            }
            $i++;
        }
        return $result;
    }

    /**
     * We sort both arrays, then we iterate over them using index.
     *
     * @param Integer[] $nums1
     * @param Integer[] $nums2
     * @return Integer[]
     */
    public function sortedArrays(array $nums1, array $nums2): array
    {
        sort($nums1);
        sort($nums2);
        $id1 = 0;
        $id2 = 0;
        $result = [];
        while ($id1 < count($nums1) && $id2 < count($nums2)) {
            if ($nums1[$id1] === $nums2[$id2]) {
                $result[] = $nums1[$id1];
                $id1++;
                $id2++;
                continue;
            }
            if ($nums1[$id1]<$nums2[$id2]) {
                $id1++;
                continue;
            }
            $id2++;
        }
        return $result;
    }
}