<?php

declare(strict_types=1);

namespace Konorlevich\Leetcode\IsomorphicStrings;

/**
 * Given two strings s and t, determine if they are isomorphic.
 *
 * Two strings s and t are isomorphic if the characters in s can be replaced to get t.
 *
 * All occurrences of a character must be replaced with another character while preserving the order of characters.
 * No two characters may map to the same character, but a character may map to itself.
 *
 * 1 <= s.length <= 5 * 104
 * t.length == s.length
 * s and t consist of any valid ascii character.
 */
class Solution
{

    /**
     */
    function isIsomorphic(string $s, string $t): bool
    {
        if (strlen($s) != strlen($t)) {
            return false;
        }
        if ($s === $t) {
            return true;
        }
        $replaces = [];
        $replacesReversed = [];

        for ($i = 0; $i < strlen($s); $i++) {
            if (
                array_key_exists($t[$i], $replacesReversed)
                && $s[$i] !== $replacesReversed[$t[$i]]
            ) {
                return false;
            }
            if (
                array_key_exists($s[$i], $replaces) &&
                $t[$i] !== $replaces[$s[$i]]
            ){
                return false;
            }
            $replaces[$s[$i]] = $t[$i];
            $replacesReversed[$t[$i]] = $s[$i];
        }
        return true;
    }
}