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
   out/build -in data/people-daily.txt
   out/fenci -text 北冰洋汽水真好喝


问题
----

因为是简易实现所以有如下问题：

1. 没法处理标点符号
2. 没法处理数字
3. 没法处理人名地名
4. 没法加入用户自定义字典


动机
----

玩玩...
