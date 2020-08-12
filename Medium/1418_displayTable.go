/*
给你一个数组 orders，表示客户在餐厅中完成的订单，确切地说， orders[i]=[customerNamei,tableNumberi,foodItemi] ，其中 customerNamei 是客户的姓名，tableNumberi 是客户所在餐桌的桌号，而 foodItemi 是客户点的餐品名称。

请你返回该餐厅的 点菜展示表 。在这张表中，表中第一行为标题，其第一列为餐桌桌号 “Table” ，后面每一列都是按字母顺序排列的餐品名称。接下来每一行中的项则表示每张餐桌订购的相应餐品数量，第一列应当填对应的桌号，后面依次填写下单的餐品数量。

注意：客户姓名不是点菜展示表的一部分。此外，表中的数据行应该按餐桌桌号升序排列。



示例 1：

输入：orders = [["David","3","Ceviche"],["Corina","10","Beef Burrito"],["David","3","Fried Chicken"],["Carla","5","Water"],["Carla","5","Ceviche"],["Rous","3","Ceviche"]]
输出：[["Table","Beef Burrito","Ceviche","Fried Chicken","Water"],["3","0","2","1","0"],["5","0","1","0","1"],["10","1","0","0","0"]]
解释：
点菜展示表如下所示：
Table,Beef Burrito,Ceviche,Fried Chicken,Water
3    ,0           ,2      ,1            ,0
5    ,0           ,1      ,0            ,1
10   ,1           ,0      ,0            ,0
对于餐桌 3：David 点了 "Ceviche" 和 "Fried Chicken"，而 Rous 点了 "Ceviche"
而餐桌 5：Carla 点了 "Water" 和 "Ceviche"
餐桌 10：Corina 点了 "Beef Burrito"
示例 2：

输入：orders = [["James","12","Fried Chicken"],["Ratesh","12","Fried Chicken"],["Amadeus","12","Fried Chicken"],["Adam","1","Canadian Waffles"],["Brianna","1","Canadian Waffles"]]
输出：[["Table","Canadian Waffles","Fried Chicken"],["1","2","0"],["12","0","3"]]
解释：
对于餐桌 1：Adam 和 Brianna 都点了 "Canadian Waffles"
而餐桌 12：James, Ratesh 和 Amadeus 都点了 "Fried Chicken"
示例 3：

输入：orders = [["Laura","2","Bean Burrito"],["Jhon","2","Beef Burrito"],["Melissa","2","Soda"]]
输出：[["Table","Bean Burrito","Beef Burrito","Soda"],["2","1","1","1"]]


提示：

1 <= orders.length <= 5 * 10^4
orders[i].length == 3
1 <= customerNamei.length, foodItemi.length <= 20
customerNamei 和 foodItemi 由大小写英文字母及空格字符 ' ' 组成。
tableNumberi 是 1 到 500 范围内的整数。

*/
package main

import (
	"sort"
	"strconv"
)

func displayTable(orders [][]string) [][]string {
	if orders == nil || len(orders) == 0 {
		return [][]string{}
	}
	tmp := make([]string, 0)
	//1、首先获取所有的菜名
	hash1 := make(map[string]bool, 0)
	for i := 0; i < len(orders); i++ {
		hash1[orders[i][2]] = true
	}
	for key, _ := range hash1 {
		tmp = append(tmp, key)
	}
	//将菜名按照字典序进行排序
	sort.Strings(tmp)
	tmp = append([]string{"Table"}, tmp...)
	//2、获取桌号和点菜数
	hash2 := make(map[string]map[string]int, 0)
	for i := 0; i < len(orders); i++ {
		if _, ok := hash2[orders[i][1]]; !ok {
			hash2[orders[i][1]] = make(map[string]int, 0)
		}
		hash2[orders[i][1]][orders[i][2]]++
	}
	//保存结果
	res := make([][]string, len(hash2))
	for i := 0; i < len(hash2); i++ {
		res[i] = make([]string, len(tmp))
	}
	//3、依次填入数据
	//先填入座位标号
	i := 0
	for key, _ := range hash2 {
		res[i][0] = key
		i++
	}
	//填入具体的订购菜品数量
	for i := 0; i < len(hash2); i++ {
		for j := 1; j < len(tmp); j++ {
			val := hash2[res[i][0]][tmp[j]]
			res[i][j] = strconv.Itoa(val)
		}
	}
	//4、进行排序
	sort.Slice(res, func(i, j int) bool {
		val1, _ := strconv.Atoi(res[i][0])
		val2, _ := strconv.Atoi(res[j][0])
		return val1 < val2
	})
	res = append([][]string{tmp}, res...)
	return res
}

func main() {}
