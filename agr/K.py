'''
事故应急
Description

作为一家业内技术领先的公司，C 公司一直努力为企业客户提供高质量的产品和服务。

然而天有不测风云，由于某位程序员同学的一个 typo，C 公司核心产品在最新版中引入了一个 bug，于是在新版本对外发布的第二天早上，负责技术支持的 P 同学发现手机上收到了很多家客户的反(吐)馈(槽)，在快速了解清楚原因之后，P 同学决定马上派出技术支持同学们到各个客户现场处理问题。由于每家客户的环境复杂程度不同，以及工作经验原因每位技术支持同学处理问题的能力也不同，P 同学冷静的分析了一下目前的状态：

一共有 N 家客户遇到了问题，第 i 家客户需要能力值达到p[i]的技术支持同学才能解决问题；

一共有 K 位技术支持同学，第 i 位技术支持同学的能力值为w[i]；

由于问题比较复杂，每位技术支持同学当天只能处理一家客户的问题；

每家客户最多只能派一位技术支持同学

P 同学希望当天能处理尽可能多家客户，聪明的你能告诉他当天最多有几家客户的问题能得到处理吗？


Input
第一行输入一个整数 T，代表数据组数，对于每组数据：

第一行输入两个整数 N (1 ≤ N ≤ 10) 和 K (1 ≤ K ≤ 10)，表示出问题的客户数量和技术支持同学的人数；

第二行输入输入 N 个整数，第 i 个整数表示第 i 家客户对能力值的需求；

第三行输入 K 个整数，第 i 个整数表示第 i 位技术支持同学的能力值。


Output
对于每组输入，输出一个整数 R，表示当天最多有 R 家客户的问题能得到处理。


Sample Input 1 

2
3 3
5 7 9
6 8 10
3 5
5 7 9
3 3 5 3 3
Sample Output 1

3

'''
#!/usr/bin/python
# -*- coding: UTF-8 -*-

import sys

def quick_sort(array, l, r):
    if l < r:
        q = partition(array, l, r)
        quick_sort(array, l, q-1)
        quick_sort(array, q + 1, r)
 
def partition(array, l, r):
    x = array[r]
    i = l - 1
    for j in range(l, r):
        if array[j] <= x:
            i += 1
            array[i], array[j] = array[j], array[i]
    array[i + 1], array[r] = array[r], array[i+1]
    return ( i+1 ) 

def F(cus,sup):
    dicsup={}
    count=0
    quick_sort(cus,0,len(cus)-1)
    quick_sort(sup,0,len(sup)-1)
    for icus, valcus in enumerate(cus):
        for isup, valsup in enumerate(sup):
            if  valsup>=valcus:
                del sup[isup]
                count=count+1
                break
    return count

if __name__=="__main__":
    getLine = raw_input()
    while True:
        try:
            getLine = raw_input()
            getLine = raw_input()
            cus=getLine.strip().split(" ")
            cus=map(int, cus)
            getLine = raw_input()
            sup=getLine.strip().split(" ")
            sup=map(int, sup)
            print F(cus,sup)
        except EOFError:
            break
