<?php

declare(strict_types=1);

namespace IsomorphicStrings;

use Konorlevich\Leetcode\IsomorphicStrings\Solution;
use PHPUnit\Framework\Attributes\CoversClass;
use PHPUnit\Framework\Attributes\DataProvider;
use PHPUnit\Framework\TestCase;

#[CoversClass(Solution::class)]
class SolutionTest extends TestCase
{

    public static function dataProvider(): array
    {
        return [
            'empty' => [
                's' => '',
                't' => '',
                'expected' => true,
            ],
            'a | ' => [
                's' => 'a',
                't' => '',
                'expected' => false,
            ],
            'egg | add' => [
                's' => 'egg',
                't' => 'add',
                'expected' => true,
            ],
            'foo | bar' => [
                's' => 'foo',
                't' => 'bar',
                'expected' => false,
            ],
            'paper | title' => [
                's' => 'paper',
                't' => 'title',
                'expected' => true,
            ],
            'badc | baba' => [
                's' => 'badc',
                't' => 'baba',
                'expected' => false,
            ],
        ];
    }

    /**
     * @param string $s
     * @param string $t
     * @param bool $expected
     * @return void
     */
    #[DataProvider('dataProvider')]
    public function testIsIsomorphic(string $s, string $t, bool $expected): void
    {
        $this->assertEquals($expected, (new Solution)->isIsomorphic($s, $t));
    }
}
