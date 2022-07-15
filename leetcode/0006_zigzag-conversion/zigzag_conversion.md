# Z字形变换
将一个给定字符串s根据给定的行数numRows，以从上往下、从左到右进行Z字形反唇排列。

比较输入字符串为“PAYPALISHIRING"行数为3时，排列如下：
```
P   A   H   N
A P L S I I G
Y   I   R
```
之后，你的输出需要从左从右逐行读取，产生出一个新的字符串，比如："PAHNAPLSIIGYIR"。

示例1：
```
输入：s = "PAYPALISHIRING", numRows = 3
输出："PAHNAPLSIIGYIR"
```


示例2：
```
输入：s = "PAYPALISHIRING", numRows = 4
输出："PINALSIGYAHRPI"
解释：
P     I    N
A   L S  I G
Y A   H R
P     I
```