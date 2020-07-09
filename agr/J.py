'''
base64 与 36 进制
Description

在一次渗透特别行动中，由于一个特别的隧道只能传递 0-9A-Z 这 36 个字符（注意：全部大写），聪明的你决定发明 36 进制来解决数据传输问题。

在 36 进制表达方案中，0 仍然代表 0，9 也依然代表 9，A 代表 10，B 代表 11，依次类推，Z 代表 35。

与常见 10 进制书写方法类似（如：123 这个数字，1代表百位，是最大的一位），你发明的 36 进制编码也是按照最大的位放在最左边来输出的。

题目任务：给定一个 base64 字符串，将这个 base64 所编码的原始内容转换成 36 进制并输出。

提示：首先需要将 base64 字符串解码得到原始内容，然后将原始内容转换成 36 进制的最短表达。

提示2：仔细观察样例输入输出，确保弄明白编码的方式。


Input
一个字符串，确保是符合 base64 格式的字符串。


Output
编码为 36 进制的最终结果。


Sample Input 1 

YQ==
Sample Output 1

2P
Sample Input 2 

eW8=
Sample Output 2

NZJ
Sample Input 3 

cGFzc3dvcmQgaXMgQWRtaW4xMjM0NTY=
Sample Output 3

3N4I4HGY6IPFI151WM1CU9FGSRJ8QLQSL92U
'''
#!/usr/bin/python
# -*- coding: UTF-8 -*-

import base64
import sys

Base36 = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

def Convert36(base64Str):
    data=base64.b64decode(base64Str)
    res = ['%02X' % ord(i) for i in data]
    return NumTo36Hex(int("0x"+"".join(res), 16))

def NumTo36Hex(num):
    if num==0:
        return "0"
    str=""
    while num>0:
        tmp=num%36
        str=Base36[tmp]+str
        num=num/36
    return str


if __name__=="__main__":
    while True:
        try:
            getLine = raw_input()
            print Convert36(getLine)
        except EOFError:
            break