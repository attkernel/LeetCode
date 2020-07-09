'''
monster 要喝水
Description

公司终于安好饮水机了，monster 迫不及待要去接水，但是他到那里才发现前面已经有n个同事了。他数了数，饮水机一共有m个接水口。所有的同事严格按照先来后到去接水（m个接水口同时工作，哪个水龙头有空人们就去哪里，如果 n<m，那么就只有n个接水口工作）。每个人都有一个接水的时间，当一个人接完水后，另一个人马上去接，不会浪费时间。monster 着急要开会，所以他想知道什么时候才能轮到他。


Input
第一行两个整数n和m，表示 monster 前面有n个人，饮水机有m个接水口。n,m<1100。第二行n个整数，表示每个同学的接水时间。


Output
一行，一个数，表示轮到 monster 接水的时间


Sample Input 1 

4 2
1 1 1 1

Sample Output 1

3
'''
#!/usr/bin/python
# -*- coding: UTF-8 -*-

def F(n,m,times):
    if m>n:
        return 0
    if m==n:
        min=999999999
        for idx, value in enumerate(times):
            if value<min:
                min=value
        return min
    dic={}
    for idx, value in enumerate(times):
        dic[idx+1]=value
    nums=m+1
    count=0
    while nums<=n+1:
        for i in range(1,m+1):
            dic[i]=dic[i]-1
            if dic[i]==0:
                if dic.has_key(nums):
                    dic[i]=dic[nums]
                    nums=nums+1
                else:
                    return count+2
        count=count+1
    return count+1

if __name__=="__main__":
    while True:
        try:
            getLine = raw_input()
            cus=getLine.strip().split(" ")
            cus=map(int, cus)
            getLine = raw_input()
            sup=getLine.strip().split(" ")
            sup=map(int, sup)
            print F(cus[0],cus[1],sup)
        except EOFError:
            break