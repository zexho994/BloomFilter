# BloomFilter
BloomFilter by Go , go 实现的布隆过滤器

![](https://tva1.sinaimg.cn/large/008i3skNgy1gqbdqavvkkj30c60au74i.jpg)


## 设计

![](https://imgconvert.csdnimg.cn/aHR0cHM6Ly91c2VyLWdvbGQtY2RuLnhpdHUuaW8vMjAxOS8xMC8yOC8xNmUxMTJmYmQwMzFmZTcx?x-oss-process=image/format,png)

### Bit数组长度设计

Bit数组的宽度取决于两个因素：

1. 预估数据量 m
2. 允许误判率 k

数据量越多，误判率要求低，要求的数组长度就越大

### Hash函数设计

对于输入的数据，使用Hash(s)获取n个bit位。

- 对于Add(s),将n个bit位存储到Bit数组中。
- 对于Get(s),查询n个bit位是否都存在，从而判断数据是否存在。

数学证明： http://pages.cs.wisc.edu/~cao/papers/summary-cache/node8.html

### 测试
```golang
func TestBloomFilter_StringExist(t *testing.T) {
	total := 1000000
  fp := 0.05
	bf := NewBloomFilter(uint(total), fp)
	for i := 0; i < total; i++ {
		bf.addString(strconv.Itoa(i))
	}

	fail := 0
	for i := 0; i < total; i++ {
		if !bf.exist(strconv.Itoa(i)) {
			fail++
		}
	}
	fmt.Println("fail count = ", fail)

	fail = 0
	for i := total; i < total+10000; i++ {
		if bf.exist(strconv.Itoa(i)) {
			fail++
		}
	}
	fmt.Println("fail count = ", fail)
}
```
