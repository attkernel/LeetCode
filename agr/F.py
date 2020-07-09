'''
数组运算
Description

给你一个 n 个数的数组 a_1 ,a_2, ..., a_na (无序)一次操作可以把两个数替换成这两个数的和。

比如数组 [2, 1, 4][2,1,4] 可以变成: [3, 4], [1, 6], [2, 5][3,4],[1,6],[2,5] 。

请问，在不限次合并操作之后，数组中最多能有多少个数可以被 33 整除。


Input
第一行包含一个整数 t(1<=t<=1000)t(1≤t≤1000) ,表示有 t 个样例每个样例的第一行有一个整数 n(1<=n<=100)n(1≤n≤100) 表示数组中数的个数, 第二行有 n 个整数 a_1 ,a_2, ..., a_n (1 <=a_i<=1e9)为数组元素。


Output
每个样例输出一行包含 m 表示该数组中在操作之后最多能有 m 个数可以被 3 整除。


Sample Input 1 

2
5
3 1 2 3 1
7
1 1 1 1 1 2 2

Sample Output 1

3
3

Hint

[3, 1, 2, 3, 1] -> [3, 3, 3, 1]
[1,1,1,1,1,2,2]−>[1,1,1,1,2,3]−>[1,1,1,3,3]−>[2,1,3,3]−>[3,3,3]
'''
#!/usr/bin/python
# -*- coding: UTF-8 -*-

def F(src):
    p1=0
    p2=0
    sum=0
    for idx, value in enumerate(src):
        sum=sum+value
    for idx, value in enumerate(src):
        if value%3==sum%3 and sum%3!=0:
            del src[idx]
            break
    while p2<len(src):
        if src[p2]%3==0:
            src[p1],src[p2]=src[p2],src[p1]
            p1=p1+1
        p2=p2+1
    p2=len(src)-1
    end=len(src)-1
    while p1<=p2:
        if (src[p1]+src[p2])%3==0:
            p1=p1+1
            src[p2],src[end]=src[end],src[p2]
            end=end-1
            p2=end
            continue
        p2=p2-1
    return p1+(end-p1+1)/3
        

if __name__=="__main__":
    getLine = raw_input()
    while True:
        try:
            getLine = raw_input()
            getLine = raw_input()
            cus=getLine.strip().split(" ")
            cus=map(int, cus)
            print F(cus)
        except EOFError:
            break