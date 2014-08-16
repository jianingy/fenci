/*
 * filename   : const.go
 * created at : 2014-08-15 23:28:09
 * author     : Jianing Yang <jianingy.yang@gmail.com>
 */

package constants

const (
    KEY_TOTAL = "#TOTAL"             // 所有单词的数量综合
    KEY_MAXLENGTH = "#MAXLENGTH"     // 最长单词的中文字数
    KEY_NUMTERMS = "#NUMTERMS"       // 不同单词的个数
)

const (
    _ = iota
    TST_BEGIN
    TST_ALPHA
    TST_NUMBER
    TST_CHINESE_NUMBER
    TST_PUNCTUATION
    TST_WHITESPACE
    TST_HANZI
)


const (
    TOKEN_CHINESE_NUMBER = "零一二三四五六七八九十百千万亿"
)
