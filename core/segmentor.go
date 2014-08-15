/*
 * filename   : build.go
 * created at : 2014-08-15 12:03:10
 * author     : Jianing Yang <jianingy.yang@gmail.com>
 */

package core

import (
    "strconv"
    "io"
    "math"

	"github.com/torbit/cdb"

    . "github.com/jianingy/fenci/constants"
)

type Segmentor struct {
	db *cdb.Cdb
}

// 平滑函数，处理未曾见过的单词的概率
func estimate(v, n, t int) float64 {
    lambda := 0.5
    // 借助 Log 做除法，防止小数太多失真
    return math.Log(float64(v) + lambda) - math.Log(float64(n) + float64(t) * lambda)
}

func NewSegmentor(dbfile string) (*Segmentor, error) {
	db, err := cdb.Open(dbfile)
	if err != nil {
		return nil, err
	}
	// TODO: When to close thie file?
	return &Segmentor{db}, nil
}

func (seg *Segmentor) GetInt(text []byte) (int, error) {
	if str, err := seg.db.Bytes(text); err == nil {
        val, err := strconv.Atoi(string(str))
        return val, err
    } else {
        return 0, err
    }
}

func (seg *Segmentor) DoSentence(text string) ([]string, error) {
    var total, maxlen, numwords int
    var err error

    runes := []rune(text)
    length := len(runes)

    // 获取一些字典里的统计信息
    if total, err = seg.GetInt([]byte(KEY_TOTAL)); err != nil{
        return nil, err
    }
    if maxlen, err = seg.GetInt([]byte(KEY_MAXLENGTH)); err != nil {
        return nil, err
    }
    if numwords, err = seg.GetInt([]byte(KEY_NUMTERMS)); err != nil {
        return nil, err
    }

    // 初始化 DP 需要用的数据结构
    score := make([]float64, length + 1)
    refer := make([]int, length + 1)
    score[0] = math.MaxFloat64


    // 开始 DP 找概率最到的组合
    for i := 0; i < length; i++ {
        boundary := length
        if maxlen + i < length {
            boundary = i + maxlen
        }

        recorded := false
        for j := i + 1; j <= boundary; j++ {
            ch := string(runes[i:j])
            count, err := seg.GetInt([]byte(ch))
            if err == io.EOF { continue } // CDB的特点，找不到词就返回 EOF
            // 计算该词出现的概率。字典里我们只存了一个单词出现的次数
            occur := estimate(count, total, numwords)
            if score[j] < score[i] + occur {
                recorded = true
                score[j] = score[i] + occur
                refer[j] = i  // 记录最大概率句子的回朔位置
            }
        }

        // 处理第一次见到的词的概率
        if !recorded {
            occur := estimate(0, total, numwords)
            if score[i + 1] < score[i] + occur {
                score[i + 1] = score[i] + occur;
                refer[i + 1] = i;
            }
        }
    }

    var tmp, result []string

    for i := length; i > 0; i = refer[i] {
        tmp = append(tmp, string(runes[refer[i]:i]))
    }

    for i := 0; i < len(tmp); i++ {
        result = append(result, tmp[len(tmp) - i - 1])
    }

    return result, nil
}
