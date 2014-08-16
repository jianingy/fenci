简易中文分词
============

特点
----

1. 字典采用 D.J.B 的 CDB: http://cr.yp.to/cdb.html
2. 算法用动态规划计算 UNIGRAM 单句最大词频率
3. 程序很短....


使用
----

    git clone http://github.com/jianingy/fenci
    cd fenci
    make
    mkdir data
    wget -O data/people-daily.txt.gz https://nlpbamboo.googlecode.com/files/people-daily.txt.gz
    gunzip data/people-daily.txt.gz
    out/build -in data/people-daily.txt -user data/user.txt
    out/fenci -text 北冰洋汽水真好喝


关于用户自定义字典
------------------

用户自定义的词请以行分割方式放在一个文件里，构建字典的时候通过“-user”参数传入。
对于自定义词的词频都采用语料库里单词出现的平均值计算。

问题
----

因为是简易实现所以有如下问题：

1. 没法处理人名地名

动机
----

玩玩...
