## Table of Contents
  - [LeetCode 热门 100 题](#LeetCode-%E7%83%AD%E9%97%A8-100-%E9%A2%98)
    - [0. 概述](#0-%E6%A6%82%E8%BF%B0)
    - [1. 两数之和](#1-%E4%B8%A4%E6%95%B0%E4%B9%8B%E5%92%8C)
    - [2. 两数相加](#2-%E4%B8%A4%E6%95%B0%E7%9B%B8%E5%8A%A0)
    - [3. 无重复字符的最长子串](#3-%E6%97%A0%E9%87%8D%E5%A4%8D%E5%AD%97%E7%AC%A6%E7%9A%84%E6%9C%80%E9%95%BF%E5%AD%90%E4%B8%B2)
    - [4. 寻找两个正序数组的中位数](#4-%E5%AF%BB%E6%89%BE%E4%B8%A4%E4%B8%AA%E6%AD%A3%E5%BA%8F%E6%95%B0%E7%BB%84%E7%9A%84%E4%B8%AD%E4%BD%8D%E6%95%B0)
    - [5. 最长回文子串](#5-%E6%9C%80%E9%95%BF%E5%9B%9E%E6%96%87%E5%AD%90%E4%B8%B2)
    - [10. 正则表达式匹配](#10-%E6%AD%A3%E5%88%99%E8%A1%A8%E8%BE%BE%E5%BC%8F%E5%8C%B9%E9%85%8D)
    - [11. 盛最多水的容器](#11-%E7%9B%9B%E6%9C%80%E5%A4%9A%E6%B0%B4%E7%9A%84%E5%AE%B9%E5%99%A8)
    - [15. 三数之和](#15-%E4%B8%89%E6%95%B0%E4%B9%8B%E5%92%8C)
    - [17. 电话号码的字母组合](#17-%E7%94%B5%E8%AF%9D%E5%8F%B7%E7%A0%81%E7%9A%84%E5%AD%97%E6%AF%8D%E7%BB%84%E5%90%88)
    - [19. 删除链表的倒数第 N 个结点](#19-%E5%88%A0%E9%99%A4%E9%93%BE%E8%A1%A8%E7%9A%84%E5%80%92%E6%95%B0%E7%AC%AC-N-%E4%B8%AA%E7%BB%93%E7%82%B9)
    - [20. 有效的括号](#20-%E6%9C%89%E6%95%88%E7%9A%84%E6%8B%AC%E5%8F%B7)
    - [21. 合并两个有序链表](#21-%E5%90%88%E5%B9%B6%E4%B8%A4%E4%B8%AA%E6%9C%89%E5%BA%8F%E9%93%BE%E8%A1%A8)
    - [22. 括号生成](#22-%E6%8B%AC%E5%8F%B7%E7%94%9F%E6%88%90)
    - [23. 合并K个升序链表](#23-%E5%90%88%E5%B9%B6K%E4%B8%AA%E5%8D%87%E5%BA%8F%E9%93%BE%E8%A1%A8)
    - [31. 下一个排列](#31-%E4%B8%8B%E4%B8%80%E4%B8%AA%E6%8E%92%E5%88%97)
    - [32. 最长有效括号](#32-%E6%9C%80%E9%95%BF%E6%9C%89%E6%95%88%E6%8B%AC%E5%8F%B7)
    - [33. 搜索旋转排序数组](#33-%E6%90%9C%E7%B4%A2%E6%97%8B%E8%BD%AC%E6%8E%92%E5%BA%8F%E6%95%B0%E7%BB%84)
    - [34. 在有序数组中查找值为 x 的区间](#34-%E5%9C%A8%E6%9C%89%E5%BA%8F%E6%95%B0%E7%BB%84%E4%B8%AD%E6%9F%A5%E6%89%BE%E5%80%BC%E4%B8%BA-x-%E7%9A%84%E5%8C%BA%E9%97%B4)
    - [39. 组合总和](#39-%E7%BB%84%E5%90%88%E6%80%BB%E5%92%8C)
    - [42. 接雨水](#42-%E6%8E%A5%E9%9B%A8%E6%B0%B4)
    - [46. 全排列](#46-%E5%85%A8%E6%8E%92%E5%88%97)
    - [48. 旋转图像](#48-%E6%97%8B%E8%BD%AC%E5%9B%BE%E5%83%8F)
    - [49. 字母异位词分组](#49-%E5%AD%97%E6%AF%8D%E5%BC%82%E4%BD%8D%E8%AF%8D%E5%88%86%E7%BB%84)
    - [53. 最大子数组和](#53-%E6%9C%80%E5%A4%A7%E5%AD%90%E6%95%B0%E7%BB%84%E5%92%8C)
    - [55. 跳跃游戏](#55-%E8%B7%B3%E8%B7%83%E6%B8%B8%E6%88%8F)
    - [56. 合并区间](#56-%E5%90%88%E5%B9%B6%E5%8C%BA%E9%97%B4)
    - [62. 不同路径](#62-%E4%B8%8D%E5%90%8C%E8%B7%AF%E5%BE%84)
    - [64. 最小路径和](#64-%E6%9C%80%E5%B0%8F%E8%B7%AF%E5%BE%84%E5%92%8C)
    - [70. 爬楼梯](#70-%E7%88%AC%E6%A5%BC%E6%A2%AF)
    - [72. 编辑距离](#72-%E7%BC%96%E8%BE%91%E8%B7%9D%E7%A6%BB)
    - [75. 颜色分类](#75-%E9%A2%9C%E8%89%B2%E5%88%86%E7%B1%BB)
    - [76. 最小覆盖子串](#76-%E6%9C%80%E5%B0%8F%E8%A6%86%E7%9B%96%E5%AD%90%E4%B8%B2)
    - [78. 子集](#78-%E5%AD%90%E9%9B%86)
    - [79. 单词搜索](#79-%E5%8D%95%E8%AF%8D%E6%90%9C%E7%B4%A2)
    - [84. 柱状图中最大的矩形](#84-%E6%9F%B1%E7%8A%B6%E5%9B%BE%E4%B8%AD%E6%9C%80%E5%A4%A7%E7%9A%84%E7%9F%A9%E5%BD%A2)
    - [85. 最大矩形](#85-%E6%9C%80%E5%A4%A7%E7%9F%A9%E5%BD%A2)
    - [94. 二叉树的中序遍历](#94-%E4%BA%8C%E5%8F%89%E6%A0%91%E7%9A%84%E4%B8%AD%E5%BA%8F%E9%81%8D%E5%8E%86)
    - [96. 不同的二叉搜索树](#96-%E4%B8%8D%E5%90%8C%E7%9A%84%E4%BA%8C%E5%8F%89%E6%90%9C%E7%B4%A2%E6%A0%91)
    - [98. 验证二叉搜索树](#98-%E9%AA%8C%E8%AF%81%E4%BA%8C%E5%8F%89%E6%90%9C%E7%B4%A2%E6%A0%91)
    - [101. 对称二叉树](#101-%E5%AF%B9%E7%A7%B0%E4%BA%8C%E5%8F%89%E6%A0%91)
    - [102. 二叉树的层序遍历](#102-%E4%BA%8C%E5%8F%89%E6%A0%91%E7%9A%84%E5%B1%82%E5%BA%8F%E9%81%8D%E5%8E%86)
    - [104. 二叉树的最大深度](#104-%E4%BA%8C%E5%8F%89%E6%A0%91%E7%9A%84%E6%9C%80%E5%A4%A7%E6%B7%B1%E5%BA%A6)
    - [105. 从前序与中序遍历序列构造二叉树](#105-%E4%BB%8E%E5%89%8D%E5%BA%8F%E4%B8%8E%E4%B8%AD%E5%BA%8F%E9%81%8D%E5%8E%86%E5%BA%8F%E5%88%97%E6%9E%84%E9%80%A0%E4%BA%8C%E5%8F%89%E6%A0%91)
    - [114. 二叉树展开为链表](#114-%E4%BA%8C%E5%8F%89%E6%A0%91%E5%B1%95%E5%BC%80%E4%B8%BA%E9%93%BE%E8%A1%A8)
    - [121. 买卖股票的最佳时机](#121-%E4%B9%B0%E5%8D%96%E8%82%A1%E7%A5%A8%E7%9A%84%E6%9C%80%E4%BD%B3%E6%97%B6%E6%9C%BA)
    - [124. 二叉树中的最大路径和](#124-%E4%BA%8C%E5%8F%89%E6%A0%91%E4%B8%AD%E7%9A%84%E6%9C%80%E5%A4%A7%E8%B7%AF%E5%BE%84%E5%92%8C)
    - [128. 最长连续序列](#128-%E6%9C%80%E9%95%BF%E8%BF%9E%E7%BB%AD%E5%BA%8F%E5%88%97)
    - [136. 只出现一次的数字](#136-%E5%8F%AA%E5%87%BA%E7%8E%B0%E4%B8%80%E6%AC%A1%E7%9A%84%E6%95%B0%E5%AD%97)
    - [139. 单词拆分](#139-%E5%8D%95%E8%AF%8D%E6%8B%86%E5%88%86)
    - [141. 环形链表](#141-%E7%8E%AF%E5%BD%A2%E9%93%BE%E8%A1%A8)
    - [142. 环形链表 II](#142-%E7%8E%AF%E5%BD%A2%E9%93%BE%E8%A1%A8-II)
    - [146. LRU 缓存](#146-LRU-%E7%BC%93%E5%AD%98)
    - [148. 排序链表](#148-%E6%8E%92%E5%BA%8F%E9%93%BE%E8%A1%A8)
    - [152. 乘积最大子数组](#152-%E4%B9%98%E7%A7%AF%E6%9C%80%E5%A4%A7%E5%AD%90%E6%95%B0%E7%BB%84)
    - [155. 最小栈](#155-%E6%9C%80%E5%B0%8F%E6%A0%88)
    - [160. 相交链表](#160-%E7%9B%B8%E4%BA%A4%E9%93%BE%E8%A1%A8)
    - [169. 多数元素](#169-%E5%A4%9A%E6%95%B0%E5%85%83%E7%B4%A0)
    - [198. 打家劫舍](#198-%E6%89%93%E5%AE%B6%E5%8A%AB%E8%88%8D)
    - [200. 岛屿数量](#200-%E5%B2%9B%E5%B1%BF%E6%95%B0%E9%87%8F)
    - [206. 反转链表](#206-%E5%8F%8D%E8%BD%AC%E9%93%BE%E8%A1%A8)
    - [207. 课程表](#207-%E8%AF%BE%E7%A8%8B%E8%A1%A8)
    - [208. 实现 Trie (前缀树)](#208-%E5%AE%9E%E7%8E%B0-Trie-%E5%89%8D%E7%BC%80%E6%A0%91)
    - [215. 数组中的第K个最大元素](#215-%E6%95%B0%E7%BB%84%E4%B8%AD%E7%9A%84%E7%AC%ACK%E4%B8%AA%E6%9C%80%E5%A4%A7%E5%85%83%E7%B4%A0)
    - [221. 最大正方形](#221-%E6%9C%80%E5%A4%A7%E6%AD%A3%E6%96%B9%E5%BD%A2)
    - [226. 翻转二叉树](#226-%E7%BF%BB%E8%BD%AC%E4%BA%8C%E5%8F%89%E6%A0%91)
    - [234. 回文链表](#234-%E5%9B%9E%E6%96%87%E9%93%BE%E8%A1%A8)
    - [235. 二叉搜索树的最近公共祖先](#235-%E4%BA%8C%E5%8F%89%E6%90%9C%E7%B4%A2%E6%A0%91%E7%9A%84%E6%9C%80%E8%BF%91%E5%85%AC%E5%85%B1%E7%A5%96%E5%85%88)
    - [236. 二叉树的最近公共祖先](#236-%E4%BA%8C%E5%8F%89%E6%A0%91%E7%9A%84%E6%9C%80%E8%BF%91%E5%85%AC%E5%85%B1%E7%A5%96%E5%85%88)
    - [238. 除自身以外数组的乘积](#238-%E9%99%A4%E8%87%AA%E8%BA%AB%E4%BB%A5%E5%A4%96%E6%95%B0%E7%BB%84%E7%9A%84%E4%B9%98%E7%A7%AF)
    - [239. 滑动窗口最大值](#239-%E6%BB%91%E5%8A%A8%E7%AA%97%E5%8F%A3%E6%9C%80%E5%A4%A7%E5%80%BC)
    - [240. 搜索二维矩阵 II](#240-%E6%90%9C%E7%B4%A2%E4%BA%8C%E7%BB%B4%E7%9F%A9%E9%98%B5-II)
    - [279. 完全平方数](#279-%E5%AE%8C%E5%85%A8%E5%B9%B3%E6%96%B9%E6%95%B0)
    - [283. 移动零](#283-%E7%A7%BB%E5%8A%A8%E9%9B%B6)
    - [287. 寻找重复数](#287-%E5%AF%BB%E6%89%BE%E9%87%8D%E5%A4%8D%E6%95%B0)
    - [300. 最长递增子序列](#300-%E6%9C%80%E9%95%BF%E9%80%92%E5%A2%9E%E5%AD%90%E5%BA%8F%E5%88%97)
    - [309. 最佳买卖股票时机含冷冻期](#309-%E6%9C%80%E4%BD%B3%E4%B9%B0%E5%8D%96%E8%82%A1%E7%A5%A8%E6%97%B6%E6%9C%BA%E5%90%AB%E5%86%B7%E5%86%BB%E6%9C%9F)
    - [322. 零钱兑换](#322-%E9%9B%B6%E9%92%B1%E5%85%91%E6%8D%A2)
    - [337. 打家劫舍 III](#337-%E6%89%93%E5%AE%B6%E5%8A%AB%E8%88%8D-III)
    - [338. 比特位计数](#338-%E6%AF%94%E7%89%B9%E4%BD%8D%E8%AE%A1%E6%95%B0)
    - [347. 前 K 个高频元素](#347-%E5%89%8D-K-%E4%B8%AA%E9%AB%98%E9%A2%91%E5%85%83%E7%B4%A0)
    - [394. 字符串解码](#394-%E5%AD%97%E7%AC%A6%E4%B8%B2%E8%A7%A3%E7%A0%81)
    - [416. 分割等和子集](#416-%E5%88%86%E5%89%B2%E7%AD%89%E5%92%8C%E5%AD%90%E9%9B%86)
    - [437. 路径总和 III](#437-%E8%B7%AF%E5%BE%84%E6%80%BB%E5%92%8C-III)
    - [438. 找到字符串中所有字母异位词](#438-%E6%89%BE%E5%88%B0%E5%AD%97%E7%AC%A6%E4%B8%B2%E4%B8%AD%E6%89%80%E6%9C%89%E5%AD%97%E6%AF%8D%E5%BC%82%E4%BD%8D%E8%AF%8D)
    - [448. 找到所有数组中消失的数字](#448-%E6%89%BE%E5%88%B0%E6%89%80%E6%9C%89%E6%95%B0%E7%BB%84%E4%B8%AD%E6%B6%88%E5%A4%B1%E7%9A%84%E6%95%B0%E5%AD%97)
    - [461. 汉明距离](#461-%E6%B1%89%E6%98%8E%E8%B7%9D%E7%A6%BB)
    - [494. 目标和](#494-%E7%9B%AE%E6%A0%87%E5%92%8C)
    - [538. 把二叉搜索树转换为累加树](#538-%E6%8A%8A%E4%BA%8C%E5%8F%89%E6%90%9C%E7%B4%A2%E6%A0%91%E8%BD%AC%E6%8D%A2%E4%B8%BA%E7%B4%AF%E5%8A%A0%E6%A0%91)
    - [543. 二叉树的直径](#543-%E4%BA%8C%E5%8F%89%E6%A0%91%E7%9A%84%E7%9B%B4%E5%BE%84)
    - [560. 和为 K 的子数组](#560-%E5%92%8C%E4%B8%BA-K-%E7%9A%84%E5%AD%90%E6%95%B0%E7%BB%84)
    - [581. 最短无序连续子数组](#581-%E6%9C%80%E7%9F%AD%E6%97%A0%E5%BA%8F%E8%BF%9E%E7%BB%AD%E5%AD%90%E6%95%B0%E7%BB%84)
    - [617. 合并二叉树](#617-%E5%90%88%E5%B9%B6%E4%BA%8C%E5%8F%89%E6%A0%91)
    - [647. 回文子串](#647-%E5%9B%9E%E6%96%87%E5%AD%90%E4%B8%B2)
    - [739. 每日温度](#739-%E6%AF%8F%E6%97%A5%E6%B8%A9%E5%BA%A6)

## LeetCode 热门 100 题

### 0. 概述

#### ➤ [LeetCode 热题 HOT 100](https://leetcode.cn/problem-list/2cktkvj/)

#### ➤ 只是刷一遍容易忘,  最好记录一下解题思路

#### ➤ 可直接上 YouTube 搜 「 LeetCode 题号 」,  推荐 NeetCode

### [1. 两数之和](https://leetcode.cn/problems/two-sum/)

1. 双指针双重循环: 对于每一个数都往后遍历一下,  看有没有匹配的
2. 只使用一次循环, 用 map 记录已经看过的数据,  在 map 中查找有没有匹配的数据

### [2. 两数相加](https://leetcode.cn/problems/add-two-numbers/)

```go
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var carry int // 进位
    var head, tail *ListNode

    for l1 != nil || l2 != nil || carry != 0 { // 遍历两个长度不等的链表
        v1, v2 := 0, 0                            // 设置默认值、能简化后续代码
        if l1 != nil {                         // 检查 nil、指针往前移动
            v1 = l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            v2 = l2.Val
            l2 = l2.Next
        }

        sum := v1 + v2 + carry
        carry = sum / 10
        value := sum % 10
        n := &ListNode{value, nil}

        if tail == nil {                      // 循环中初始化首节点
            head = n
            tail = n
        } else {
            tail.Next = n
            tail = n
        }
    }

    return head
}
```

### [3. 无重复字符的最长子串](https://leetcode.cn/problems/longest-substring-without-repeating-characters/)

1. 逐个测试以 s[0]、s[1]、s[2]、... 开头的最长子串,  最终肯定能找到那个最长子串
2. 使用滑动窗口,  相比上面的解法,  能跳过许多无意义的循环

```go
// 假设输入为 #23456#89
func lengthOfLongestSubstring(s string) int {
    result := 0
    start := 0
    seen := make(map[byte]int)

    for j := 0; j < len(s); j++ {
        c := s[j]
        x, ok := seen[c]
        if ok && x >= start { // 此时 seen[c] 和 j 是一对重复字符的索引
            start = x + 1     // 让 start 跳到左边的重复字符的下一个位置
                              // 此时应该把 [0, start-1] 的字符视为不存在
        }

        seen[c] = j
        result = max(result, j-start+1) // [start, j] 中没有出现重复字符
    }

    return result
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```

### [4. 寻找两个正序数组的中位数](https://leetcode.cn/problems/median-of-two-sorted-arrays/)

1. 两个有序数组归并一下
2. 不存整个数组,  只存两个中位数,  空间复杂度能降到 O(1)
3. 核心思路如下:

```go
// 假设要在这两个数组中查找第 k=4 个元素
A: [1, 2, 3]
B: [4, 5, 6]

// 分别取前 k/2 = 2 个元素
A: [1, 2]
B: [4, 5]

// 然后比较 2 和 5
// (1) 因为 2 < 5,  所以 2 不可能是第 k=4 个元素,  因为就算 5 之前的数全都小于 2, 比 2 小的数最多也只有 2 个
// (2) 既然 2 不是第 4 个元素,  2 前面的 1 就更不可能是,  所以排除掉 [1, 2]、共排除了 k/2 个元素

// (3) 若将两个数组合并成一个有序数组,  [1, 2] 必定位于第 4 个元素的左侧
//     这时候把 [1, 2] 抽掉,  原本第 k=4 个元素,  会变成 [3] 和 [4, 5, 6] 中第 k-2 = 2 个元素
// (4) 整个问题转变成在 [3] 和 [4, 5, 6] 两个有序数组中查找第 k=2 个元素
```

```go
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    mid1 := (len(nums1) + len(nums2) - 1) / 2   // 左中位数的索引
    mid2 := (len(nums1) + len(nums2) - 0) / 2   // 右中位数的索引
    mid1Value := getKth(nums1, nums2, mid1 + 1) // 索引加一就是第 k g数
    mid2Value := getKth(nums1, nums2, mid2 + 1)
    return (float64(mid1Value) + float64(mid2Value)) / 2.0
}

// 在两个有序数组中查找第 k 个元素
func getKth(nums1, nums2 []int, k int) int {
    var index1, index2 int
    len1, len2 := len(nums1), len(nums2)
    
    for {
        // 若有一个数组被砍光, 这时候通过索引计算返回第 k 个的元素
        if index1 == len1 {
            return nums2[index2 + k - 1]
        }
        if index2 == len2 {
            return nums1[index1 + k - 1]
        }
        // 在剩下的东西中查找第 1 个元素很简单
        if k == 1 {
            return min(nums1[index1], nums2[index2])
        }

        // 分别取前 k/2 个数
        end1 := min(index1 + k/2, len1) - 1 // 若越界则取最后一个元素
        end2 := min(index2 + k/2, len2) - 1

        // 比较两组数的最大值, 然后砍掉其中一组
        if nums1[end1] <= nums2[end2] {
            k -= end1 - index1 + 1
            index1 = end1 + 1
        } else {
            k -= end2 - index2 + 1
            index2 = end2 + 1
        }
    }
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
```

### [5. 最长回文子串](https://leetcode.cn/problems/longest-palindromic-substring/)

#### ➤ 用动态规划

```go
1. 暴力解           // 就是双重循环遍历所有子串、再判断每个子串是不是回文串、时间复杂度为 O(n^3)
2. 重叠子问题       // 暴力解有很多重叠子问题,  为了算 [0,4] 是否回文串、就要算 [1,3] 子问题,  循环时又会再算一遍
3. 定义 dp 表       // 应该记录这些子问题的解,  定义 dp[i][j] 表示字符串的 [i,j] 区间是否为回文串
4. 寻找递推关系     // dp[i][j] = dp[i+1][j-1] && str[i] == str[j],  关系成立的前提是 [i, j] 的长度 >= 3
4. 寻找递推关系     // 若 [i,j] 的长度为 1/2 则需要单独处理
5. 怎么填 dp 表     // 每个格子依赖它左下方的格子,  可以沿对角线的平行线填,  也可以一列一列的填

func longestPalindrome(s string) string {
    size := len(s)
    dp := make([][]bool, size)
    for i := range dp {
        dp[i] = make([]bool, size)
    }

    var max, start, end int

    // 按子串的长度 L=1,2,3,... 进行迭代,  也就是沿着对角线填表
    for L := 1; L <= size; L++ {
        for i := 0; i <= size-L; i++ {

            j := i + L - 1
            if L >= 3 {
                dp[i][j] = dp[i+1][j-1] && s[i] == s[j]
            } else if L == 2 {
                dp[i][j] = s[i] == s[j]
            } else if L == 1 {
                dp[i][j] = true
            }

            // 出现回文子串则维护最大值
            if dp[i][j] && L > max {
                max = L
                start = i
                end = j
            }
        }
    }

    return s[start : end+1]
}
```

#### ➤ 推荐用中心扩散,  从 i 点往左和往右扩散,  还有一种扩散方式是 i 往左、i+1 往右

```go
func longestPalindrome(s string) string {
    var res string
    for i := 0; i < len(s); i++ {   // 已知输入是 ascii 字符串
        s1 := expand(i, i, s)       // 第一种扩散方式
        s2 := expand(i, i+1, s)     // 第二种扩散方式
        res = max(res, s1, s2)
    }
    return res
}

func expand(left, right int, s string) string {
    for left >= 0 && right < len(s) && s[left] == s[right] {
        left--
        right++
    }
    return s[left+1 : right] // golang 中 "abc"[3:3] 会返回空字符串,  不会报错
}

func max(a, b, c string) string {
    res := a
    if len(res) < len(b) {
        res = b
    }
    if len(res) < len(c) {
        res = c
    }
    return res
}
```

### [10. 正则表达式匹配](https://leetcode.cn/problems/regular-expression-matching/)

#### ➤ [这个视频题解讲的很清晰](https://www.youtube.com/watch?v=HAA8mgxlov8)

```go
// 这是个暴力递归解法,  还没加 memoization, 用 i, j 表示当前要进行匹配的一对字符
func isMatch(s string, p string) bool {
    // cache := make(map[[2]int]bool) // cache 也能用 map 不一定要用二维数组
    var dfs func(i, j int) bool // 用闭包简化参数传递
    dfs = func(i, j int) bool {
        // 如果 i, j 都越界, 说明恰好匹配
        // 如果 i 不越且 j 越界则不匹配
        // 如果 i 越界且 j 不越界,  还是有匹配的可能,  比如 a*b* 能匹配 a
        if i >= len(s) && j >= len(p) {
            return true  
        }
        if j >= len(p) {
            return false
        }

        // 注意 i 有可能越界,  比较两个字符、或者检查 . 符号
        match := i < len(s) && (s[i] == p[j] || p[j] == '.')

        // 如果有 * 号要处理 * 号
        if j+1 < len(p) && p[j+1] == '*' {
            return dfs(i, j+2) || // 选择不用星号, 例如跳过 a*b 中的 a* 来到 b
                (match && dfs(i+1, j)) // 选择使用星号,  那么 a* 产生一个 a,  结果取决于 match
        }

        // 没有 * 号就很简单,  若匹配则继续比较下一对字符
        return match && dfs(i+1, j+1)
    }

    return dfs(0, 0)
}
```

### [11. 盛最多水的容器](https://leetcode.cn/problems/container-with-most-water/)

1. 使用两个指针分别指向最左和最右端, 计算一下当前的面积,  然后移动高度较短的指针
2. 为什么这样做一定能找到最大值呢?
   1. 假设有 [1, 2, 3, 4, 5],  左右指针分别指向 1 和 5
   2. 那么以 1 作为左边的情况有,  [1,5]、[1,4]、[1,3]、[1,2]
   3. 我们计算 [1,5] 的面积后让左指针来到 2 就好,  因为以 1 作为左边时,  不可能有其他情况比 [1,5] 的面积大
   4. 总之以 1 作为左边的最大值已经找到了,  我们成功缩小了问题的规模,  以此类推

```go
func maxArea(height []int) int {
    i := 0
    j := len(height) - 1
    res := 0

    for i < j {
        // j-i 是底边的长度,  高受限于较矮一边
        res = max(res, min(height[i], height[j])*(j-i))
        if height[i] < height[j] {
            i++
        } else {
            j--
        }
    }
    return res
}
```

### [15. 三数之和](https://leetcode.cn/problems/3sum/)

0. #### ➤ [题解](https://www.youtube.com/watch?v=jzZsG8n2R9A)

1. 对数组排序
2. 用 nums[i] 表示三数之和中的第一个数, 让 i 遍历数组, 并跳过重复的数
3. 用 l 和 r 两个指针分别指向两端,  nums[i+1, len(nums)-1],  然后根据 sum 大了还是小了移动 l 或 r

```go
func threeSum(nums []int) [][]int {
    sort.Ints(nums)
    res := make([][]int, 0)

    for i := range nums {
        if i > 0 && nums[i] == nums[i-1] {
            continue // 让 i 移动到下一个不重复的数字
        }

        l, r := i+1, len(nums)-1
        for l < r {
            sum := nums[i] + nums[l] + nums[r]
            switch {
            case sum > 0:
                r-- // 大了移动右指针, 让和变小
            case sum < 0:
                l++ // 小了移动左指针, 让数变大
            default:
                res = append(res, []int{nums[i], nums[l], nums[r]})
                l++
                for l < r && nums[l] == nums[l-1] {
                    l++ // 让左指针移动到下一个不重复的数字,  循环要注意越界
                }
            }
        }
    }
    return res
}
```

### [17. 电话号码的字母组合](https://leetcode.cn/problems/letter-combinations-of-a-phone-number/)

#### ➤ [回溯算法/Backtracking](https://www.youtube.com/watch?v=gBC_Fd8EE8A) 中的回溯是什么含义?

#### ➤ [题解](https://www.youtube.com/watch?v=0snEunUacZY)

```go
func letterCombinations(digits string) []string {
    if len(digits) == 0 {
        return []string{}
    }

    m := map[byte]string{
        '2': "abc", '3': "def", '4': "ghi", '5': "jkl",
        '6': "mno", '7': "pqrs", '8': "tuv", '9': "wxyz",
    }
    // 用闭包简化参数传递是个好技巧, i 表示当前要处理哪个数字, path 记录走过的路径
    res := make([]string, 0)
    var backtrack func(i int, path string)
    backtrack = func(i int, path string) {
        if len(path) == len(digits) {
            res = append(res, path)
            return
        }
        for _, letter := range m[digits[i]] {
            backtrack(i+1, path+string(letter))
        }
    }

    backtrack(0, "")
    return res
}
```

### [19. 删除链表的倒数第 N 个结点](https://leetcode.cn/problems/remove-nth-node-from-end-of-list/)

#### ➤ [题解](https://youtu.be/XVuQxVej6y8),  让快指针先走 n 步,  然后快慢指针同步前进,  他们的间隔距离不变

```go
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    // 为了删除倒数第 n 个节点,  需要它前面的那个节点
    // left 本来从 head 开始,  现在从 dummy 开始,  所以最终会滞后 1 个节点
    dummy := &ListNode{0, head}
    left, right := dummy, head

    // right 最终指向 nil 表示链表遍历完毕
    // 如果 right 只先走 1 步,  right 到结尾时 left 指向倒数第 1 个节点
    for i := 0; i < n; i++ {
        right = right.Next
    }

    for right != nil {
        left = left.Next
        right = right.Next
    }

    // 删除节点
    left.Next = left.Next.Next
    return dummy.Next
}
```

### [20. 有效的括号](https://leetcode.cn/problems/valid-parentheses/)

#### ➤ [题解](https://www.youtube.com/watch?v=WTzjTskDFMg)

1. 遇到 ( { [ 则压栈
2. 如果遇到 ) 那么栈的顶端必须是 ( 否则就不合法

```go
type Stack struct{ vals []rune }
func (s *Stack) Len() int      { return len(s.vals) }
func (s *Stack) Push(val rune) { s.vals = append(s.vals, val) }

func (s *Stack) Pop() (rune, bool) {
    if len(s.vals) == 0 {
        var zero rune
        return zero, false
    }
    top := s.vals[len(s.vals)-1]
    s.vals = s.vals[:len(s.vals)-1]
    return top, true
}

func isValid(s string) bool {
    var stack Stack
    m := map[rune]rune{
        ')': '(',
        ']': '[',
        '}': '{',
    }

    for _, c := range s {
        if val, ok := m[c]; ok { // 如果是 ) ] } 那么检查 top 是否为匹配的括号
            top, _ := stack.Pop()
            if top != val {
                return false
            }
        } else { // 否则遇到了 ( [ { 那么压栈
            stack.Push(c)
        }
    }

    return stack.Len() == 0 // 如果栈刚好空了,  说明括号全都能匹配
}
```

### [21. 合并两个有序链表](https://leetcode.cn/problems/merge-two-sorted-lists/)

#### ➤ 用递归的空间复杂度为 O(m+n),  用迭代的空间复杂度为 O(1)

```go
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    dummy := &ListNode{0, nil}
    tail := dummy

    for list1 != nil && list2 != nil {
        if list1.Val < list2.Val {
            tail.Next = list1
            list1 = list1.Next
        } else {
            tail.Next = list2
            list2 = list2.Next
        }
        tail = tail.Next
    }

    // 肯定有一个链表还没遍历完
    if list1 != nil {
        tail.Next = list1
    } else {
        tail.Next = list2
    }

    return dummy.Next
}
```

### [22. 括号生成](https://leetcode.cn/problems/generate-parentheses/)

#### ➤ [题解、回溯法](https://www.youtube.com/watch?v=s9fokUqJ76A)

1. 如果 `(` 多于 `)` 那么还有可能生成有效值,  反之不可能
2. 只要 open < n 就能继续添加 open parenthesis `(`
3. 只有 close < open 时才能添加 close parenthesis `)`

```go
func generateParenthesis(n int) []string {
    res := make([]string, 0)
    stack := make([]byte, 0, 2*n)
    var backtrack func(open, close int)

    backtrack = func(open, close int) {
        if open == n && open == close {
            res = append(res, string(stack))
            return
        }
        if open < n {
            stack = append(stack, '(')   // 先尝试加左括号
            backtrack(open+1, close)     // open 数加一
            stack = stack[:len(stack)-1] // 回来后要清理 stack
        }
        if close < open {
            stack = append(stack, ')')   // 再尝试加右括号
            backtrack(open, close+1)
            stack = stack[:len(stack)-1]
        }

    }

    backtrack(0, 0)
    return res
}
```

### [23. 合并K个升序链表](https://leetcode.cn/problems/merge-k-sorted-lists/)

#### ➤ [题解](https://www.youtube.com/watch?v=q5a5OiGbT6Q)、其实就是归并排序,  两个两个的归并、逐渐生成更大的有序链表

```go
func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists) == 0 {
        return nil
    }

    for len(lists) > 1 {                    // lists 表示待合并的一堆链表
        mergedLists := make([]*ListNode, 0) // 记录两两合并后的链表
        for i := 0; i < len(lists); i += 2 {
            list1 := lists[i]
            var list2 *ListNode             // 最后一组凑不出一对,  list2 为 nil
            if i+1 < len(lists) {
                list2 = lists[i+1]
            }
            mergedLists = append(mergedLists, merge(list1, list2))
        }
        lists = mergedLists // 很巧妙,  再来一次
    }

    return lists[0]
}

func merge(list1, list2 *ListNode) *ListNode {
    if list2 == nil {
        return list1
    }

    dummy := &ListNode{0, nil}
    tail := dummy

    for list1 != nil && list2 != nil {
        if list1.Val < list2.Val {
            tail.Next = list1
            list1 = list1.Next
        } else {
            tail.Next = list2
            list2 = list2.Next
        }
        tail = tail.Next
    }

    if list1 != nil {
        tail.Next = list1
    } else {
        tail.Next = list2
    }

    return dummy.Next
}
```

### [31. 下一个排列](https://leetcode.cn/problems/next-permutation/)

#### ➤ [这是题解](https://www.youtube.com/watch?v=quAS1iydq7U)

#### ➤ 假设输入为  2 3 5 4 1, 从后往前找到 3,  再找到第一个比 3 大的数 4,  然后交换两个数,  再把 5 3 1 逆序

#### ➤ [参考代码](https://leetcode.cn/problems/next-permutation/solution/xia-yi-ge-pai-lie-by-leetcode-solution/)

```go
func nextPermutation(nums []int) {
    n := len(nums)

    i := n - 2
    for i >= 0 && nums[i] >= nums[i+1] {
        i-- // 从后往前找、找到第一对严格升序的数: i 和 i+1
    }

    if i >= 0 { // 如果 i 不越界,  就从后往前找,  找到第一个严格大于 nums[i] 的数
        j := n - 1
        for nums[j] <= nums[i] {
            j-- // j 不可能越界,  因为 i 都找到了, 那么一定有比 i 大的 j
        }
        nums[i], nums[j] = nums[j], nums[i] // 交换两个数
    }
    reverse(nums[i+1:]) // 逆序 i 后面的数
}

func reverse(nums []int) {
    // n/2 是右中位数,  因为要小于右中位数,  i 最多到左中位数
    n := len(nums)
    for i := 0; i < n/2; i++ {
        nums[i], nums[n-1-i] = nums[n-1-i], nums[i]
    }
}
```

### [32. 最长有效括号](https://leetcode.cn/problems/longest-valid-parentheses/)

#### ➤ [题解](https://www.bilibili.com/video/BV13g41157hK?t=598.1&p=22)

1. 不难想到用 [0, i] 的答案推出 [0, i+1] 的答案
2. 那么 dp[i] 应该如何定义呢?  
   1. 若定义 dp[i] 表示 [0, i] 中有一段合法子串且长度为 dp[i],  那么找不到任何联系,  无法推出 dp[i+1] 的值
   2. 若定义 dp[i] 表示 [0, i] 中存在一段合法子串,  且子串刚好以 i 结尾,  且长度为 dp[i]
   3. 那么比较一下 i+1 和 i+1 - dp[i] - 1 这两个字符,  似乎就能得到 dp[i+1] 的值

```go
func longestValidParentheses(s string) int {
    dp := make([]int, len(s))
    max := 0
    for i := 1; i < len(s); i++ {        // i 从 1 开始,  下面的 i-1 就不会越界
        if s[i] == ')' {                 // 如果遇到左括号,  那么肯定是 0,  也就不用处理了
            pre := i - dp[i-1] - 1       // dp[i-1] 决定了要往前跳多少个字符
            if pre < 0 {                 // 比如 ()) 就会让 pre 越界
                continue
            }
            if s[pre] == '(' {           // 如果能配对,  至少是 dp[i-1] + 2
                dp[i] = dp[i-1] + 2      
                if pre-1 >= 0 {          // 再往前看一个, 把它也接上
                    dp[i] += dp[pre-1]   
                }
            }
            if dp[i] > max {
                max = dp[i]
            }
        }
    }
    return max
}
```

### [33. 搜索旋转排序数组](https://leetcode.cn/problems/search-in-rotated-sorted-array/)

#### ➤ [题解](https://youtu.be/U8XENwh8Oy8)

1. 旋转后产生两个递增序列,  左边的递增序列的所有值都更大
2. mid 要么位于左边的递增序列, 要么位于右边的递增序列
3. 当 nums[l] <= nums[mid] 时,  mid 位于左边的递增序列

```go
func search(nums []int, target int) int {
    l, r := 0, len(nums)-1
    for l <= r {
        mid := (l + r) / 2
        if target == nums[mid] {
            return mid
        }

        if nums[l] <= nums[mid] { // mid 位于左边的递增序列
            // 如果要找一个更大的数,  只能往右半部分找
            // 或者想找一个更小的,  并且 target 比 nums[l] 还小那么也得去右半部分
            if target > nums[mid] || target < nums[l] {
                l = mid + 1
            } else {
                r = mid - 1
            }

        } else { // mid 位于右边的递增序列
            // 如果要找一个更小的数,  只能去左边
            // 或者想找一个更大的,  且 target 比 nums[r] 都大,  那么也要去左边
            if target < nums[mid] || target > nums[r] {
                r = mid - 1
            } else {
                l = mid + 1
            }
        }
    }
    return -1
}
```

### [34. 在有序数组中查找值为 x 的区间](https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/)

#### ➤ [题解](https://www.youtube.com/watch?v=4sQL7R5ySUU)

1. 如果用朴素的二分搜索找到头,  再滑动一下找到尾,  那么时间复杂度是 O(n) 而不是 O(logn)
2. 改造一下二分搜索, 让它能找到头和尾,  然后执行两次二分搜索

```go
func searchRange(nums []int, target int) []int {
    left := binarySearch(nums, target, true)
    right := binarySearch(nums, target, false)
    return []int{left, right}
}

func binarySearch(nums []int, target int, leftBias bool) int {
    l, r := 0, len(nums)-1
    i := -1
    for l <= r {
        m := (l + r) / 2
        if target < nums[m] {
            r = m - 1
        } else if target > nums[m] {
            l = m + 1
        } else {
            i = m // 一般的二分找到就直接 return 了,  但这里还得跑
            if leftBias {
                r = m - 1
            } else {
                l = m + 1
            }
        }
    }

    return i
}
```

### [39. 组合总和](https://leetcode.cn/problems/combination-sum/)

#### ➤ [题解](https://youtu.be/GBKI9VSKdGg)

1. 观察这棵树能发现有 []、[2]、[2 2]、[2 2 2]、... 等 2 的个数分别为 0/1/2/3/... 的子树
2. 在 [2] 这颗子树下,  2 的个数已经固定为 1 不能再用 2 所以接下来考虑要不要使用 3

```go
func combinationSum(candidates []int, target int) [][]int {
    res := make([][]int, 0)
    path := make([]int, 0)
    var dfs func(i, total int)

    dfs = func(i, total int) {
        if total == target {
            res = append(res, copySlice(path)) // 要复制一下, 别添加同一个 slice
            return
        }
        if i >= len(candidates) || total > target { // 没有选择了, 或者已经超过了
            return
        }
        path = append(path, candidates[i])
        dfs(i, total+candidates[i]) // 使用 candidates[i]
        path = path[:len(path)-1]   // 回溯后清理路径

        dfs(i+1, total) // 不用 candidates[i], 后续子树也不准用它,  所以 i+1
    }

    dfs(0, 0)
    return res
}

func copySlice(s []int) []int {
    s2 := make([]int, len(s))
    copy(s2, s)
    return s2
}
```

### [42. 接雨水](https://leetcode.cn/problems/trapping-rain-water/)

#### ➤ [题解](https://www.youtube.com/watch?v=ZI2z5pq0TqA)

1. 位置 i 能接的雨水数是 min(maxLeft[i], maxRight[i]) - height[i]
2. 可以做两次遍历, 把 maxLeft 和 maxRight 两个数组算出来
3. 空间复杂度还能优化到 O(1),  如果 maxL <= maxR 那么移动 L 指针,  并且 L 位置能接 maxL - height[L]
4. 当 maxL <= maxR 时,  L 位置的接水量只和 L、maxL 相关,  maxR 究竟是什么不重要

```go
func trap(height []int) int {
    res := 0
    l, r := 0, len(height)-1
    leftMax, rightMax := height[l], height[r]

    for l < r {
        if leftMax <= rightMax {
            l += 1                            // 移动左指针, 右边的最大值是什么无所谓
            leftMax = max(leftMax, height[l]) // 在这行更新 leftMax
            res += leftMax - height[l]        // 这行就不可能出现负数了
        } else {
            r -= 1
            rightMax = max(rightMax, height[r])
            res += rightMax - height[r]
        }
    }

    return res
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```

### [46. 全排列](https://leetcode.cn/problems/permutations/)

```go
func permute(nums []int) [][]int {
    res := make([][]int, 0)
    var backtrack func(start int)

    backtrack = func(cur int) {
        if cur == len(nums) {
            res = append(res, copySlice(nums))
            return
        }

        for i := cur; i < len(nums); i++ {
            nums[cur], nums[i] = nums[i], nums[cur] // [i, n-1] 是可选的东西,  选择哪个,  就把他放到 cur 位置
            backtrack(cur + 1)                      // cur+1 让剩下的选择变少
            nums[cur], nums[i] = nums[i], nums[cur] // 撤销选择
        }
    }

    backtrack(0)
    return res
}

func copySlice(s []int) []int {
    s2 := make([]int, len(s))
    copy(s2, s)
    return s2
}
```

### [48. 旋转图像](https://leetcode.cn/problems/rotate-image/)

#### ➤ [题解](https://www.youtube.com/watch?v=fMSJSS7eO1w)、从外到内一圈一圈的旋转、用 l/r/t/b 四个指针表示当前要旋转的圈

```go
func rotate(matrix [][]int) {
    l, r := 0, len(matrix)-1 // 用 left right top bottom 表示一个圈,  其中 top/bottom 表示行号
    for l < r {
        top, bottom := l, r // 矩阵中每一圈的 top 和 left 总是相等
        for i := 0; i < r-top; i++ {
            temp := matrix[top][l+i]               // 暂存左上角的值,  把位置空出来, 左上角往右移动所以 l + i
            matrix[top][l+i] = matrix[bottom-i][l] // 用左下角填充左上角, 左下角往上移动,  所以 bottom - i
            matrix[bottom-i][l] = matrix[bottom][r-i]
            matrix[bottom][r-i] = matrix[top+i][r]
            matrix[top+i][r] = temp
        }
        l += 1
        r -= 1
    }
}
```

### [49. 字母异位词分组](https://leetcode.cn/problems/group-anagrams/)

#### ➤ 统计一下字符串中的 a/b/c/... 各出现了多少次,  用统计结果作为 key,  进行分组

```go
func groupAnagrams(strs []string) [][]string {
    res := make(map[[26]int][]string)

    for _, str := range strs {
        var count [26]int
        for i := range str {
            count[str[i]-'a']++
        }
        res[count] = append(res[count], str)
    }

    values := make([][]string, 0, len(res)) // make slice 的时候第二个参数记得填 0
    for _, v := range res {
        values = append(values, v)
    }
    return values
}
```

### [53. 最大子数组和](https://leetcode.cn/problems/maximum-subarray/)

#### ➤ [题解](https://www.youtube.com/watch?v=5WZl3MMT0Eg)

1. 像滑动窗口那样、遇到一个和为负数的前缀就砍掉
2. 用动态规划,  dp[i] 表示以 i 作为结尾的最大和, dp[i] = max(dp[i-1]+nums[i], nums[i]),  如果接上更大就接上

```go
func maxSubArray(nums []int) int {
    maxSub := nums[0]  // 不要用 0,  因为最大值可能是负数
    curSum := 0

    for _, n := range nums {
        if curSum < 0 {
            curSum = 0 // 遇到和为负数的前缀,  就砍掉这个前缀
        }
        curSum += n
        maxSub = max(maxSub, curSum)
    }
    return maxSub
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```

### [55. 跳跃游戏](https://leetcode.cn/problems/jump-game/)

#### ➤ [题解](https://youtu.be/Yan0cv2cLy8)

1. 暴力解就是, 从索引 0 开始,  逐个尝试自己的选择,  比如跳 1/2/3 步
2. 从暴力解可以观察到,  有很多的重复子问题,  通过不同的路径跳到同一个位置,  可以加上 memoization 
3. 使用贪心,  从后往前看,  不断尝试把 goal 往前移动,  只要能从起点跳到新的 goal 就一定存在路径

```go
func canJump(nums []int) bool {
    n := len(nums)
    goal := n - 1
    for i := n - 1; i >= 0; i-- {
        if i+nums[i] >= goal {
            goal = i // 如果 i 位置能跳到 goal,  则 goal 变成 i
        }
    }
    return goal == 0
}
```

### [56. 合并区间](https://leetcode.cn/problems/merge-intervals/)

#### ➤ [题解](https://www.youtube.com/watch?v=44H3cEC2fFM)

1. intervals 按起点排序
2. 遍历 intervals 并判断是否与上一个重叠, 若重叠则合并, 合并时注意 [1,5]、[2,4] 这种情况要取 max(5,4)

```go
func merge(intervals [][]int) [][]int {
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })

    // res[len(res)-1] 是上一个元素,  把上个元素初始化成首元素,  方便循环中直接合并
    res := [][]int{{intervals[0][0], intervals[0][1]}}

    for _, i := range intervals {
        if last := res[len(res)-1]; i[0] <= last[1] { // 当前元素的头 <= 上个元素的尾
            last[1] = max(last[1], i[1])              // 则进行合并, 存在 [1,5]、[2,4] 这种情况所以取 max
        } else {
            res = append(res, []int{i[0], i[1]})      // 复制 slice 而不是直接用 i,  避免共享数据被修改
        }
    }
    return res
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```

### [62. 不同路径](https://leetcode.cn/problems/unique-paths/)

#### ➤ [题解](https://www.youtube.com/watch?v=IlEsdxuD4lY)

1. 当前点的路径数 = 右边的路径数 + 下面的路径数
2. 从下往上、从右往左一行一行填表,  因为只依赖右边和下面,  所以空间能优化成一行

```go
func uniquePaths(m int, n int) int {

    row := make([]int, n) // 用 row 表示最后一行,  全都是 1
    for i := range row {
        row[i] = 1
    }

    // 最后一行和最后一列都是 1,  所以从倒数第二行和倒数第二列开始
    for r := m - 2; r >= 0; r-- {
        for c := n - 2; c >= 0; c-- {
            // row[c+1] 是右边的值,  row[c] 是下面的值, 可以把结果存到 newRow,  但原地更新就好
            row[c] = row[c+1] + row[c]
        }
    }
    return row[0]
}
```

### [64. 最小路径和](https://leetcode.cn/problems/minimum-path-sum/)

#### ➤ [题解](https://www.youtube.com/watch?v=pGMsrvt0fpk)

1. 当前点的最短路径 = 当前点的值 + min(右边的最短路径, 下面的最短路径)

```go
func minPathSum(grid [][]int) int {
    m, n := len(grid), len(grid[0])

    row := make([]int, n)                    // 用 row 表示最后一行, 然后进行初始化
    row[n-1] = grid[m-1][n-1]
    for i := n - 2; i >= 0; i-- {            // 从倒数第二列开始
        row[i] = row[i+1] + grid[m-1][i]     // 右边的最短路径 + 当前的长度
    }

    for r := m - 2; r >= 0; r-- {            // 从倒数第二行和倒数第二列开始
        row[n-1] = grid[r][n-1] + row[n-1]   // 循环前把最后一列的值维护好
        for c := n - 2; c >= 0; c-- {
            // row[c+1] 是右边的值,  row[c] 是下面的值, 可以把结果存到 newRow,  但原地更新就好
            row[c] = grid[r][c] + min(row[c+1], row[c])
        }
    }
    return row[0]
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
```

### [70. 爬楼梯](https://leetcode.cn/problems/climbing-stairs/)

#### ➤ [题解](https://www.youtube.com/watch?v=Y0lT9Fck7qI)

1. 可以通过不同的路径跳到同一个台阶,  然后遇到相同的子问题....
2. 定义 dp[i] 表示从 i 台阶开始,  跳到 n 台阶有多少种跳法
3. dp[i] = dp[i+1] + dp[i+2]  // 从 i 可以跳到 i+1/i+2,  总的跳法就是两种情况之和

```go
func climbStairs(n int) int {
    one, two := 1, 1
    for i := n - 2; i >= 0; i-- { // 利用 n = 2/3 的特例,  得知 i 从 n-2 开始
        temp := one
        one = one + two // 只需要最近的两项
        two = temp      // 新的 two 是之前 one 的值
    }
    return one
}
```

### [72. 编辑距离](https://leetcode.cn/problems/edit-distance/)

#### ➤ [题解](https://www.youtube.com/watch?v=XYi2-LPrwm4)

1. 对两个字符串逐字符进行比较,  如果匹配则 i++, j++,  如果不匹配则对三种编辑方式都试一遍,  然后会产生三种子问题
2. 所有的子问题全都在这样一张 dp 表里面: 定义 `dp[i][j]` 表示后缀 `word1[i:]` 到后缀 `word2[j:]` 的编辑距离
3. 当发现 `word1[i]` 和 `word2[j]` 不匹配时,  若选择替换字符,  那么就会来到 `dp[i+1][j+1]` 这个子问题
4. 若选择添加字符,  就会来到 `dp[i][j+1]` 这个子问题,  总而言之 `dp[i][j]` 依赖右边、下面、以及右下角这三个子问题

```go
func minDistance(word1 string, word2 string) int {
    // dp 表格要额外增加一行一列,  表示从空串到另一个子串的最短编辑距离
    m, n := len(word1), len(word2)
    dp := make([][]int, m+1)
    for i := range dp {
        dp[i] = make([]int, n+1)
    }

    // 初始化最后一行和最后一列的值
    for i := 0; i < n+1; i++ {
        dp[m][i] = len(word2) - i
    }
    for i := 0; i < m+1; i++ {
        dp[i][n] = len(word1) - i
    }

    // 从倒数第二行填到第一行,  从倒数第二列开始,  从右往左填表
    for i := m - 1; i >= 0; i-- {
        for j := n - 1; j >= 0; j-- {
            if word1[i] == word2[j] {
                // 两个字符匹配, 那么当前子问题的解等于右下角子问题的解
                dp[i][j] = dp[i+1][j+1]
            } else {
                // 否则三种编辑方式都试一下,  取一种最小的,  因为使用了编辑所以要加一
                dp[i][j] = 1 + min(dp[i+1][j], dp[i][j+1], dp[i+1][j+1])
            }
        }
    }
    return dp[0][0]
}

func min(a, b, c int) int {
    res := a
    if b < res {
        res = b
    }
    if c < res {
        res = c
    }
    return res
}
```

[1143. 最长公共子序列](https://leetcode.cn/problems/longest-common-subsequence/)

```go
func longestCommonSubsequence(text1 string, text2 string) int {
    m, n := len(text1), len(text2) // 定义 dp[i][j] 表示后缀 text1[i:] 和后缀 text2[j:] 的最长公共子序列
    dp := make([][]int, m+1)       // 额外添加一行一列, 因为依赖右边和下面的值
    for i := range dp {
        dp[i] = make([]int, n+1)
    }

    for i := m - 1; i >= 0; i-- { // 从倒数第二行和倒数第二列开始
        for j := n - 1; j >= 0; j-- {
            if text1[i] == text2[j] {
                dp[i][j] = 1 + dp[i+1][j+1] // 若匹配, 那么序列长度 +1 并取决于右下角子问题的解
            } else {
                dp[i][j] = max(dp[i+1][j], dp[i][j+1]) // 否则从两个子问题的解中选一个更长的
            }
        }
    }

    return dp[0][0]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```

### [75. 颜色分类](https://leetcode.cn/problems/sort-colors/)

#### ➤ [题解](https://www.youtube.com/watch?v=4xbWSRZHqac)

1. 可以用桶排序,  因为只会出现 0/1/2 三个数字,  统计一下每个数字的出现次数,  然后生成答案
2. 使用快速排序的 partition 方法,  
3. 注意把一个数发送到大于区域后, 可能换过来一个小于 pivot 的数,  所以此时 i 不能移动

```go
func sortColors(nums []int) {
    l, r := 0, len(nums)-1
    i, pivot := 0, 1

    swap := func(i, j int) {
        nums[i], nums[j] = nums[j], nums[i]
    }

    for i <= r { // i 和 r 相等时,  r 还没被处理,  s
        switch {
        case nums[i] < pivot:
            swap(i, l)
            l++
            i++
        case nums[i] > pivot:
            swap(i, r) // 换过来一个未知数, 所以 i 保持不动
            r--
        case nums[i] == pivot:
            i++
        }
    }
}
```

### [76. 最小覆盖子串](https://leetcode.cn/problems/minimum-window-substring/)

#### ➤ [题解](https://www.youtube.com/watch?v=jSto0O4AJbM)

1.  使用滑动窗口,  窗口里面维护两种数据,  ①当前满足了几个条件 ②当前窗口内各个字符的计数
2. 如果 have < need 那么窗口的右指针移动一步,  纳入新的字符, 并更新 ① 和 ②
3. 如果 have == need 那么窗口的左指针移动一步,  删掉字符,  尝试找到更短的覆盖子串

```go
func minWindow(s string, t string) string {
    // 统计 t 中各字符的出现次数
    count := make(map[byte]int)
    for i := range t {
        count[t[i]]++
    }

    res := []int{0, -1}
    resLen := math.MaxInt
    window := make(map[byte]int) // 用于统计窗口中各字符的出现次数
    have, need := 0, len(count)  // 需要满足 need 个条件,  当前满足 0 个

    for l, r := 0, 0; r < len(s); r++ {
        // 添加了一个字符,  更新 window 的统计数据
        c := s[r]
        window[c]++

        // 若 window[c] 首次达到了 count[c] 则满足的条件数加一
        // 另外如果 c 是不相关字符,  那么 count[c] 是 0 而 window[c] 大于 0
        if window[c] == count[c] {
            have++
        }

        for have == need {
            // 找到了一个覆盖子串,  那么更新最小的覆盖子串
            if r-l+1 < resLen {
                resLen = r - l + 1
                res[0], res[1] = l, r
            }
            // 然后移动窗口的左指针,  尝试找更短的覆盖子串,  移动前要更新数据
            // 另外如果 s[l] 是不相关字符, 那么 count[s[l]] 是 0 而 window[s[l]] 不可能小于 0
            window[s[l]]--
            if window[s[l]] < count[s[l]] {
                have--
            }
            l++
        }
    }

    return s[res[0] : res[1]+1]
}
```

### [78. 子集](https://leetcode.cn/problems/subsets/)

#### ➤ [题解](https://www.youtube.com/watch?v=Vn2v6ajA7U0)、从 nums[0] 开始,  决定用或者不用 nums[i],  不同的选择将导致不同的状态,  然后继续做选择,  叶子节点就是解

```go
func subsets(nums []int) [][]int {
    path := make([]int, 0, len(nums))
    res := make([][]int, 0)
    var backtrack func(i int)

    backtrack = func(i int) {
        if i == len(nums) {
            res = append(res, copySlice(path))
            return
        }
        path = append(path, nums[i]) // 第一种选择,  使用 nums[i]
        backtrack(i + 1)
        path = path[:len(path)-1]    // 回溯后进行清理
        backtrack(i + 1)             // 第二种选择,  不使用 nums[i]
    }

    backtrack(0)
    return res
}

func copySlice(s []int) []int {
    s2 := make([]int, len(s))
    copy(s2, s)
    return s2
}
```

### [79. 单词搜索](https://leetcode.cn/problems/word-search/)

#### ➤ [题解](https://www.youtube.com/watch?v=pfiQ_PS1g8E)

1. 在每个格子都试一下深度优先搜索,  用哈希表记录用过的点来避免重复使用同一个点

```go
func exist(board [][]byte, word string) bool {
    m, n := len(board), len(board[0])
    path := make(map[[2]int]bool) // 记录路径经过了哪些点, 避免重复使用一个点
    var dfs func(r, c, i int) bool

    dfs = func(r, c, i int) bool {
        if i == len(word) {
            return true // 没有需要匹配的字符了
        }
        if r < 0 || c < 0 || r >= m || c >= n {
            return false // 处理 (r,c) 越界
        }
        if board[r][c] != word[i] || path[[2]int{r, c}] {
            return false // 字符不匹配,  或者当前点曾经用过了
        }

        path[[2]int{r, c}] = true  // 如果字符能匹配,  那么当前点计入路径,  再往四个方向都试一下
        res := dfs(r+1, c, i+1) || dfs(r, c+1, i+1) || dfs(r-1, c, i+1) || dfs(r, c-1, i+1)
        path[[2]int{r, c}] = false // 返回前进行清理
        return res
    }

    for r := range board {
        for c := range board[r] {
            if dfs(r, c, 0) {
                return true
            }
        }
    }

    return false // 时间复杂度是 m * n * 4^len(word)
}
```

### [84. 柱状图中最大的矩形](https://leetcode.cn/problems/largest-rectangle-in-histogram/)

#### ➤ [题解](https://www.youtube.com/watch?v=zx5Sw9130L0)

1. 维护一个栈,  栈中的元素都有继续延伸底边的可能
2. 若遇到 `栈顶元素 > heights[i]` 那么栈顶元素就丧失了继续延伸底边的可能,  这时候可以计算面积并弹栈
3. 若 `新的栈顶也 > heights[i]` 那么新的栈顶也会丧失继续延伸底边的可能,  这时候就能计算面积并弹栈
4. 若 `新的栈顶 <= heights[i]` 那么它还能继续延伸, 所以把 `heigts[i]` 压栈
5. 总之栈中元素的高度是递增的,  这种栈叫做单调栈
6. 压栈时, 起点能往前延伸,  比如 `[3 5 6] 4` 在弹出 5、6 并压入 4 时, 4 的起始索引记为 1
7. 弹栈时要计算面积, 并且底边能往后延伸,  比如 `[3 4 5] 2`,  计算 3/4 的面积时底边能延伸到 5

```go
func largestRectangleArea(heights []int) int {
    type area struct {
        start  int
        height int
    }
    maxArea := 0
    stack := make([]area, 0)

    for i, h := range heights {
        start := i
        for len(stack) > 0 && stack[len(stack)-1].height > h {
            top := stack[len(stack)-1]                       // 栈顶元素的底边不可能往后延伸了, 所以能计算面积并弹栈
            stack = stack[:len(stack)-1]                     // 只要栈非空、并且栈顶高于 heigts[i] 就一直弹栈
            maxArea = max(maxArea, top.height*(i-top.start)) // 计算面积
            start = top.start                                // 弹栈时 heights[i] 的起点可以往前延伸到栈顶元素的起点
        }

        stack = append(stack, area{start, h}) // 栈顶 <= heights[i] 那么把它压栈
    }

    for _, item := range stack {
        maxArea = max(maxArea, item.height*(len(heights)-item.start)) // 栈中剩下元素的底边都能延伸到末尾
    }

    return maxArea
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```

### [85. 最大矩形](https://leetcode.cn/problems/maximal-rectangle/)

#### ➤ [题解](https://leetcode.cn/problems/maximal-rectangle/solution/xiang-xi-tong-su-de-si-lu-fen-xi-duo-jie-fa-by-1-8/)、依次考虑前 1/2/3/... 行并统计一下每一列的高度、把问题转化成 84 题、然后调用上一题的函数

```go
func maximalRectangle(matrix [][]byte) int {
    m, n := len(matrix), len(matrix[0])
    res := 0

    heights := make([]int, n)
    for r := 0; r < m; r++ {
        for c := 0; c < n ; c++ {
            if matrix[r][c] == '0' { // 坑,  格子中是 '0' 而不是 0
                heights[c] = 0
            } else {
                heights[c]++
            }
        }
        res = max(res, largestRectangleArea(heights))
    }
    return res
}
```

### [94. 二叉树的中序遍历](https://leetcode.cn/problems/binary-tree-inorder-traversal/)

#### ➤ [题解](https://www.youtube.com/watch?v=g_S5WuasWUE)

1. 尝试掌握递归和迭代两种方法

```go
func inorderTraversal(root *TreeNode) []int {
    res := make([]int, 0)
    var inorder func(root *TreeNode)
    inorder = func(root *TreeNode) {
        if root == nil {
            return
        }
        inorder(root.Left)
        res = append(res, root.Val) // 回到此位置时就是中序遍历
        inorder(root.Right)
    }
    inorder(root)
    return res
}


// 前序和中序的迭代写法一样,  只是使用 val 的时机不同
// 后续的迭代写法是往右的前序遍历,  再逆序一下
func inorderTraversal(root *TreeNode) []int {
    res := make([]int, 0)
    stack := make([]*TreeNode, 0)
    cur := root

    for cur != nil || len(stack) > 0 {

        for cur != nil {               // 尝试一直往左走
            stack = append(stack, cur) // 首次来到这个节点, 进行压栈
            cur = cur.Left             // 然后往左走
        }

        cur = stack[len(stack)-1]    // 弹栈
        stack = stack[:len(stack)-1] //
        res = append(res, cur.Val)   // 回来时就是中序遍历
        cur = cur.Right              // 然后往右走
    }
    return res
}
```

### [96. 不同的二叉搜索树](https://leetcode.cn/problems/unique-binary-search-trees/)

#### ➤ [题解](https://www.youtube.com/watch?v=Ox0TenN3Zpg)

1. 分情况讨论,  如果 1/2/3/4/5 分别作为根节点,  总共有多少种搜索二叉树?  
   如果 2 作为根节点,  那么 1 一定在左子树,  3,4,5 一定在右子树
2. 那么 3,4,5 又能构成多少种搜索二叉树?  所以这是一个递归的问题
3. 3,4,5 节点的值不重要,  3,4,5 共三个节点,  这决定了能构成多少个搜索二叉树
4. numTree[3] = (numTree[0] * numTree[2]) + (numTree[1] * numTree[1]) + (numTree[2] * numTree[0])

```go
func numTrees(n int) int {
    dp := make([]int, n+1)
    dp[0] = 1
    dp[1] = 1
    // 根据这个特例写代码就很方便:
    // dp[3] = dp[0]*dp[2] + dp[1]*dp[1] + dp[2]*dp[0]
    for i := 2; i <= n; i++ {
        total := 0
        for j := 0; j < i; j++ {
            total += dp[j] * dp[i-j-1] // 两索引之和为 i-1
        }
        dp[i] = total
    }
    return dp[n]
}
```

### [98. 验证二叉搜索树](https://leetcode.cn/problems/validate-binary-search-tree/)

#### ➤ [题解](https://www.youtube.com/watch?v=s6ATEkipzow)

1. 检查中序遍历是否单调递增,  就能判断是否搜索二叉树
2. 递归解法,  以 5 为根, 往左走那么左边元素都必须小于 5,  往右走那么右边元素都必须大于 5

```go
// 递归解法
func isValidBST(root *TreeNode) bool {
    var check func(root *TreeNode, left, right int) bool
    check = func(node *TreeNode, left, right int) bool {
        if node == nil {
            return true  // nil 视为搜索二叉树
        }
        if !(node.Val > left && node.Val < right) {
            return false // 检查当前节点是否满足它的限制条件
        }
        return check(node.Left, left, node.Val) && // 往左走只需修改上界
            check(node.Right, node.Val, right)     // 往右走只需修改下界
    }

    return check(root, math.MinInt, math.MaxInt)
}

// 中序遍历解法
func isValidBST(root *TreeNode) bool {
    stack := make([]*TreeNode, 0)
    cur := root
    res := []int{math.MinInt, math.MinInt}
    
    for cur != nil || len(stack) > 0 {
        for cur != nil {
            stack = append(stack, cur)
            cur = cur.Left
        }

        cur = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        if res[len(res)-1] >= cur.Val {
            return false // 如果中序遍历的上一个元素 >= 当前元素
        }

        res[0], res[1] = res[1], cur.Val
        cur = cur.Right
    }
    return true
}
```

### [101. 对称二叉树](https://leetcode.cn/problems/symmetric-tree/)

```go
func isSymmetric(root *TreeNode) bool {
    if root == nil {
        return true
    }

    // 用 check 检查 p、q 两颗子树是否对称
    var check func(p, q *TreeNode) bool
    check = func(p, q *TreeNode) bool {
        if p == nil && q == nil {
            return true // 两个空子树是对称的
        }
        if p == nil || q == nil || p.Val != q.Val {
            return false // 仅一个子树为空,  或者都不为空但值不相等
        }

        // 现在能确定 p、q 子树的根节点是对称的
        // 然后递归地检查 p.Left 和 q.Right 这一对子树是否对称
        return check(p.Left, q.Right) && check(p.Right, q.Left)
    }

    return check(root.Left, root.Right)
}
```

### [102. 二叉树的层序遍历](https://leetcode.cn/problems/binary-tree-level-order-traversal/)

#### ➤ [题解](https://www.youtube.com/watch?v=6ZnyEApgFYg)

1. 用一个队列存储每一层的节点
1. 其实就是 BFS 广度优先遍历

 ```go
 func levelOrder(root *TreeNode) [][]int {
     q := []*TreeNode{root}
     res := make([][]int, 0)
 
     for len(q) > 0 {
         qLen := len(q)                          // 当前层有多少个节点
         level := make([]int, 0, qLen)
 
         for i := 0; i < qLen; i++ {             // 把一层的节点全都出队
             node := q[0]
             q = q[1:]
             if node != nil {                    // 遇到 nil 节点则跳过
                 level = append(level, node.Val)
                 q = append(q, node.Left)        // 左节点入队
                 q = append(q, node.Right)       // 右节点入队
             }
         }
         if len(level) > 0 {                     // 可能整层都是 nil 节点
             res = append(res, level)
         }
     }
     return res
 }
 ```

### [104. 二叉树的最大深度](https://leetcode.cn/problems/maximum-depth-of-binary-tree/) 

```go
func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0 // 空子树的高度为 0,  当前子树的高度是 1 + 较高的那颗子树
    }
    return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```

### [105. 从前序与中序遍历序列构造二叉树](https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/)

#### ➤ [题解](https://www.youtube.com/watch?v=ihj4IQGZ2zc)

1. 先序: [①, ②, ③, ④]  中序: [③, ②, ①, ④]
2. 首先 `①` 是根节点,  然后看看中序遍历, `①` 的左侧只有 2 个元素,  所以 `[②, ③]` 属于左子树
3. `[②, ③]` 是左子树的先序遍历,  `[③, ②]` 是左子树的中序遍历,  所以可用递归解决子问题

```go
func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 {
        return nil // 一个元素都没有,  所以返回空子树
    }

    root := &TreeNode{preorder[0], nil, nil}
    mid := indexOf(inorder, root.Val)
    root.Left = buildTree(preorder[1:1+mid], inorder[:mid])   // [1:1+mid] 表示从 1 开始, 长度为 mid 个
    root.Right = buildTree(preorder[1+mid:], inorder[mid+1:]) // [1+mid:] 表示剩下的内容
    return root
}

func indexOf(s []int, target int) int {
    for i := range s {
        if s[i] == target {
            return i
        }
    }
    return -1
}
```

### [114. 二叉树展开为链表](https://leetcode.cn/problems/flatten-binary-tree-to-linked-list/)

#### ➤ [题解](https://www.youtube.com/watch?v=rKnD7rLT0lI)

1. 把左子树展开成链表,  若展开完毕,  再把左子树插入到右子树
2. 子树展开成链表后,  返回链表的尾节点

```go
func flatten(root *TreeNode) {
    var flat func(root *TreeNode) *TreeNode
    
    flat = func(root *TreeNode) *TreeNode {
        // 这么递归下去肯定遇到 root == nil, 这就是 base case
        if root == nil {
            return nil
        }

        // 把左子树展开,  把右子树也展开,  然后和 root 拼起来
        leftTail := flat(root.Left)
        rightTail := flat(root.Right)

        // 若左子树不为空,  那么要进行拼接
        if root.Left != nil {
            leftTail.Right = root.Right
            root.Right = root.Left
            root.Left = nil
        }

        if rightTail != nil { // 只要右子树不为空就返回 rightTail
            return rightTail  // 此处不能用 root.Right != nil 做判断,  因为 root.Right 已经被改了
        }
        if leftTail != nil {  // 右子树为空且左子树不为空,  则返回 leftTail
            return leftTail
        }
        return root           // 左右子树都为空, 则返回 root
    }

    flat(root)
}
```

### [121. 买卖股票的最佳时机](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/)

#### ➤ [题解](https://www.youtube.com/watch?v=1pkOgXD63yU)

1. 让 L 指向当前找到的最低点,  然后移动 R 并计算利润
2. 若 R 移动到了新的最低点,  则移动 L 到 R,  然后继续移动 R 并计算利润
3. 总之 L 会遍历所有的低点,  答案的低点一定会被 L 覆盖,  答案的高点也会被 R 扫描到

```go
func maxProfit(prices []int) int {
    l, r := 0, 1
    maxP := 0

    for r < len(prices) {
        if prices[r] > prices[l] { // 有钱挣
            if profit := prices[r] - prices[l]; profit > maxP { // 挣更多
                maxP = profit
            }
        } else { // 亏钱,  找到了比 l 更低的低点
            l = r
        }
        r++
    }
    return maxP
}
```

### [124. 二叉树中的最大路径和](https://leetcode.cn/problems/binary-tree-maximum-path-sum/)

#### ➤ [题解](https://www.youtube.com/watch?v=Hr5cWUld4vU)

1. 节点值 + max(0, 左子树的返回值) + max(0, 右子树的返回值) = 潜在的最大值
2. 节点值 + max(0, max(左子树的返回值, 右子树的返回值)) = 节点的返回值
3. 整个计算过程自底向上,  每个子树都会向上层贡献一个值,  让上层决定要不要选用

```go
func maxPathSum(root *TreeNode) int {
    res := root.Val // 不要初始化成 0,  万一整颗树都是负数呢
    var dfs func(root *TreeNode) int
    
    dfs = func(root *TreeNode) int {
        if root == nil {
            return 0                              // nil 节点只能向上层贡献一个 0
        }
        leftMax := max(0, dfs(root.Left))         // 若子树返回负值,  则丢弃这个值
        rightMax := max(0, dfs(root.Right))

        res = max(res, root.Val+leftMax+rightMax) // 把左右子树连起来, 可能产生最大值
        return root.Val + max(leftMax, rightMax)  // 选择较大的子树,  方便向上层贡献更大的值
    }
    dfs(root)
    return res
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```

### [128. 最长连续序列](https://leetcode.cn/problems/longest-consecutive-sequence/)

#### ➤ [题解](https://www.youtube.com/watch?v=P6RZZMu_maU)

1. 把数组做成集合
2. 遍历数组,  寻找递增序列的起点,  比如处理 100 时,  因为 99 也在集合里,  所以 100 不是起点
3. 找到起点 99 后, 轮询 100、101、102、... 是否存在

 ```go
func longestConsecutive(nums []int) int {
    set := make(map[int]bool, len(nums))
    for _, n := range nums {
        set[n] = true
    }

    maxL := 0
    for num := range set { // 遍历去重后的集合比数组更快一点
        if !set[num-1] {   // 找到了连续序列的起点
            length := 1
            for set[num+length] {
                length++
            }
            if length > maxL {
                maxL = length
            }
        }
    }
    return maxL
}
 ```

### [136. 只出现一次的数字](https://leetcode.cn/problems/single-number/)

#### ➤ [题解](https://www.youtube.com/watch?v=qMPX1AOa83k)

1. 把数字转成二进制形式,  全部数字做异或就能得到答案
2. 因为把这些二进制数对齐,  单看某一列,  其中的 0 可以忽略,  因为 `n xor 0 = n`
3. 并且其中 1 的出现次数必然是偶数个、也会被消去

```go
func singleNumber(nums []int) int {
    res := 0 // 因为 n ^ 0 = n
    for _, n := range nums {
        res ^= n
    }
    return res
}
```

### [139. 单词拆分](https://leetcode.cn/problems/word-break/)

#### ➤ [题解](https://www.youtube.com/watch?v=Sx9NNgInc3A)

```go
func wordBreak(s string, wordDict []string) bool {
    n := len(s)
    dp := make([]bool, n+1) // dp[i] 表示 s[i:] 是否能被 word break
    dp[n] = true            // base case 是空字符串

    for i := n - 1; i >= 0; i-- {
        for _, word := range wordDict {
            wordL := len(word)
            if i+wordL <= n && s[i:i+wordL] == word { // 取和 word 一样长度的前缀匹配一下
                dp[i] = dp[i+wordL]                   // 若匹配那么 dp[i] 的答案取决于 dp[i+wordL]
            }
            if dp[i] {
                break // 比如 abcd 和 ab、cd、abcd 有两种匹配方式,  找到一种就行了
            }
        }
    }

    return dp[0]
}
```

### [141. 环形链表](https://leetcode.cn/problems/linked-list-cycle/)

#### ➤ [题解](https://www.youtube.com/watch?v=gBTe7lFR3vc)

1. 用 set 记录已经见过的节点,  遇到重复节点说明有环
2. 用快慢指针,  如果是无环链表,  快慢指针不可能相遇,  如果是有环链表,  快指针会在环内追上慢指针

```go
func hasCycle(head *ListNode) bool {
    slow, fast := head, head
    for fast != nil && fast.Next != nil { // 检查 fast 是否越界
        fast = fast.Next.Next
        slow = slow.Next
        if fast == slow {
            return true
        }
    }
    return false
}
```

###  [142. 环形链表 II](https://leetcode.cn/problems/linked-list-cycle-ii/)

1. 简单的办法是使用 set 记录见过的节点,  若遇到重复节点,  那就是入环节点
2. 使用快慢指针,  相遇时快指针比慢指针多走 n 圈:  `fast = slow + n圈`,  又因为 `fast = 2*slow` 所以 `slow = n圈`
3. 总而言之,  `慢指针从头`走 slow 步会来到相遇点,  `慢指针从相遇点`走 slow 步会回到相遇点 (因为 slow = n圈)
4. 若有两个慢指针, 一个从头、另一个从相遇点都走 slow 步,  那他们会在相遇点相遇,  
   又因为速度是相同的,  所以在相遇点的前一个点, 他们也会相遇, 而首个相遇点就是入环点

```go
// (1) 使用 set 记录见过的节点
func detectCycle(head *ListNode) *ListNode {
    node := head
    seen := make(map[*ListNode]bool)
    for node != nil {
        if seen[node] {
            return node
        }
        seen[node] = true
        node = node.Next
    }
    return nil
}

// (2) 使用快慢指针
func detectCycle(head *ListNode) *ListNode {
    fast, slow := head, head
    for fast != nil && fast.Next != nil {
        fast = fast.Next.Next
        slow = slow.Next
        if fast == slow { // 有环找环
            fast = head   // fast 从头开始,  再次相遇就是入环点
            for fast != slow {
                fast = fast.Next
                slow = slow.Next
            }
            return fast
        }
    }

    return nil // 无环返回 nil
}
```

### [146. LRU 缓存](https://leetcode.cn/problems/lru-cache/)

#### ➤ [题解](https://www.youtube.com/watch?v=7ABFKPK2hD4)

```go
type node struct {
    key, value int
    prev, next *node
}

type LRUCache struct {
    capacity    int
    left, right *node
    data        map[int]*node
}

func Constructor(capacity int) LRUCache {
    left, right := &node{}, &node{}
    left.next = right // left.prev 用不到
    right.prev = left // right.next 用不到
    data := make(map[int]*node)
    return LRUCache{capacity, left, right, data}
}

func (c *LRUCache) Get(key int) int {
    if n, ok := c.data[key]; ok {
        c.remove(n) // 从链表移除
        c.insert(n) // 插入到 right 前面
        return n.value
    }
    return -1
}

func (c *LRUCache) Put(key int, value int) {
    // key 存在则从链表移除那个节点
    if old, ok := c.data[key]; ok {
        c.remove(old)
    }

    // 存储新节点并插入到 right 前面
    n := &node{key: key, value: value}
    c.data[key] = n
    c.insert(n)

    // 可能超长,  那么移除 lru 节点
    if len(c.data) > c.capacity {
        lru := c.left.next
        delete(c.data, lru.key)
        c.remove(lru)
    }
}

func (c *LRUCache) remove(n *node) {
    prev, next := n.prev, n.next // 记下 prev 和 next
    prev.next = next
    next.prev = prev
}

func (c *LRUCache) insert(n *node) {
    prev, next := c.right.prev, c.right // 记下 prev 和 next
    prev.next = n                       // 连好两个 next 指针
    n.next = next
    next.prev = n
    n.prev = prev
}
```

### [148. 排序链表](https://leetcode.cn/problems/sort-list/)

#### ➤ [题解](https://leetcode.cn/problems/sort-list/?favorite=2cktkvj)

```go
func sortList(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    return mergeSort(head)
}

func mergeSort(head *ListNode) *ListNode {
    if head.Next == nil {
        return head                // 只有一个节点
    }

    // fast 从 head.Next 开始,  可以用 2/3/4 节点来思考边界情况
    mid, fast := head, head.Next
    for fast != nil && fast.Next != nil {
        fast = fast.Next.Next
        mid = mid.Next
    }

    right := mergeSort(mid.Next)
    mid.Next = nil                 // 拆成两个链表
    left := mergeSort(head)
    return merge(left, right)
}

func merge(head1, head2 *ListNode) *ListNode {
    dummy := &ListNode{}
    tail := dummy
    for head1 != nil && head2 != nil {
        if head1.Val < head2.Val { // 其实 cpu 计算 < 和 <= 的性能是一样的
            tail.Next = head1
            head1 = head1.Next
        } else {
            tail.Next = head2
            head2 = head2.Next
        }
        tail = tail.Next
    }
    if head1 != nil {
        tail.Next = head1
    }
    if head2 != nil {
        tail.Next = head2
    }
    return dummy.Next
}
```

### [152. 乘积最大子数组](https://leetcode.cn/problems/maximum-product-subarray/)

1. 若只遇到一个负数,  比如 [2 3 -4 5 6 7],  currMax 会在 -4 之后从 5 重新开始累积
2. 若只遇到两个负数,  比如 [2 3 -4 5 6 7 -2],  currMax 会从 -2 和 currMin 的乘积中产生
3. 若遇到 3 个负数,  其实和只遇到一个负数的情况一样,  前两个负负得正了

```go
func maxProduct(nums []int) int {
    res := nums[0]
    curMin, curMax := 1, 1
    for _, n := range nums {
        tmp := curMax
        curMax = max(n*curMax, n*curMin, n) // 若 n 与 curMax/curMin 的符号相反则返回 n, 符号相同则可能产生最大值
        curMin = min(n*tmp, n*curMin, n)
        res = max(res, curMax)
    }
    return res
}

func max(nums ...int) int {
    res := nums[0]
    for _, n := range nums {
        if n > res {
            res = n
        }
    }
    return res
}

func min(nums ...int) int {
    res := nums[0]
    for _, n := range nums {
        if n < res {
            res = n
        }
    }
    return res
}
```

### [155. 最小栈](https://leetcode.cn/problems/min-stack/)

#### ➤ [题解](https://www.youtube.com/watch?v=qkLl7nAwDPo)

1. 使用一个 minStack 记录到目前为止的最小值,  所以 `最小值` 和 `次最小值` 都会被记录, 弹栈后次最小值升级为最小值
2. 两个栈的高度是同步的,  压栈时用 minStack 记录到目前为止的最小值,  弹栈时两个栈同步弹栈

```go
type MinStack struct {
    stack    []int
    minStack []int
}

func Constructor() MinStack {
    return MinStack{}
}

func (s *MinStack) Push(val int) {
    s.stack = append(s.stack, val)
    if len(s.minStack) == 0 {
        s.minStack = append(s.minStack, val)
        return // 注意检查栈为空
    }
    curMin := s.minStack[len(s.minStack)-1]
    if val < curMin {
        curMin = val
    }
    s.minStack = append(s.minStack, curMin)
}

func (s *MinStack) Pop() {
    s.stack = s.stack[:len(s.stack)-1]
    s.minStack = s.minStack[:len(s.minStack)-1]
}

func (s *MinStack) Top() int {
    return s.stack[len(s.stack)-1]
}

func (s *MinStack) GetMin() int {
    return s.minStack[len(s.minStack)-1]
}
```

### [160. 相交链表](https://leetcode.cn/problems/intersection-of-two-linked-lists/)

#### ➤ [题解](https://www.youtube.com/watch?v=D0X0BONOQhI)

1. 容易想到的办法是,  先把 a 链表的全部节点放入 set,  然后 b 链表遍历的过程中判断是否遇到相交点
2. 让长链表先走 n 步,  然后大家同步走,  首个相交点就是答案
3. 走完各自的路后, 我把你的路走一遍, 你把我的走一遍,  我们最后会停在相同的终点  
   因为终点是相遇点, 速度又相同, 所以终点的前一个点也是相遇点,  而我们首次相遇时就是链表的相交点

```go
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    n1, n2 := headA, headB
    if n1 == nil || n2 == nil {
        return nil // 确保 n1 和 n2 都不为 nil
    }

    for n1 != n2 {
        if n1 != nil {
            n1 = n1.Next
        } else {
            n1 = headB // 走到 nil 了就换一条路,  若无交点,  最终 n1 == n2 == nil
        }

        if n2 != nil {
            n2 = n2.Next
        } else {
            n2 = headA // 走到 nil 了就换一条路,  若有交点,  那么首次相遇就是链表交点
        }
    }
    return n1
}

func getIntersectionNode2(headA, headB *ListNode) *ListNode {
    // 找出两链表的长度
    lenA, lenB := 0, 0
    for n := headA; n != nil; n = n.Next {
        lenA++
    }
    for n := headB; n != nil; n = n.Next {
        lenB++
    }
    // 希望 headA 比 headB 短, 若不满足则交换
    if lenB < lenA {
        lenA, lenB = lenB, lenA
        headA, headB = headB, headA
    }
    // 链表 B 更长,  让他先走一段距离,  然后相遇时就是相交点
    for i := 0; i < lenB-lenA; i++ {
        headB = headB.Next
    }

    for headA != nil && headB != nil {
        if headA == headB && headA != nil { // 注意先判断相交再移动指针
            return headA
        }
        headA = headA.Next
        headB = headB.Next
    }
    return nil
}
```

### [169. 多数元素](https://leetcode.cn/problems/majority-element/)

#### ➤ [题解](https://www.youtube.com/watch?v=7pnhv842keE)

用哈希表记录所见数字的出现次数,  若发现了出现次数更多的数字,  则进行更新

使用 Boyer–Moore majority vote algorithm,  注意众数的出现次数 > 一半,  比如 [2 2 4 4] 则不存在众数  
以 [2 2 1 1 1] 为例,  遇见相同数字 count++, 遇见不同数字 count--,  当 count 为零时换一个数,  众数最终会活下来  
假设喜欢露琪亚、妮莉艾露、井上织姬的人分别为 3、1、1,  哪怕所有非众数联合起来,  也打不过众数

```go
// 哈希表
func majorityElement(nums []int) int {
    count := make(map[int]int, len(nums))
    maxCount, res := 0, nums[0]
    for _, n := range nums {
        count[n]++
        if count[n] > maxCount {
            res = n
            maxCount = count[n]
        }
    }
    return res
}

// 用于求众数的 Boyer–Moore majority vote algorithm
func majorityElement(nums []int) int {
    count, res := 0, nums[0]
    for _, n := range nums {
        if count == 0 {      // count 为 0 所以之前的数据不存在众数
            count++
            res = n
        } else if n == res { // 遇到相同的数 count++
            count++
        } else {
            count--          // 遇到不同的数 count--
        }
    }
    return res
}
```

### [198. 打家劫舍](https://leetcode.cn/problems/house-robber/)

#### ➤ [题解](https://www.youtube.com/watch?v=73r3KWiEvyk)

1. 问题是 `[0:n]` 能挣多少钱,  
2. 若打劫 `nums[0]`, 那么剩下的问题是 `[2:n]` 能挣多少钱, 无论 `[2:n]` 的最优解长啥样,  它都逃不出我的选择范围
3. 若不打劫 `nums[0]` 那么剩下的问题是 `[1:n]` 能挣多少钱,  在这两种选择中最优的选择就是问题的解
4. 定义 `dp[i]` 表示 `[i:n]` 的解, 则有 `dp[i] = max(nums[i] + dp[i+2], dp[i+1])`
5. 总之利用后两个状态就能推出当前状态

```go
func rob(nums []int) int {
    res := 0
    a, b := 0, 0                             // 用 a, b 追踪最后两个状态
    for i := len(nums) - 1; i >= 0; i-- {
        res = max(nums[i]+b, a)
        a, b = res, a                        // a, b 向前移动
    }
    return res
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```

### [200. 岛屿数量](https://leetcode.cn/problems/number-of-islands/)

#### ➤ [题解](https://www.youtube.com/watch?v=pV2kpPD66nE)

1. 在一个点执行 bfs 就能把整座岛标记为已见过,  可以在每个点都跑一下 bfs
2. 这道题主要考察图的遍历,  需要用 set 记录已经见过的节点,  避免重复访问

```go
func numIslands(grid [][]byte) int {
    m, n := len(grid), len(grid[0])
    seen := make(map[point]bool)
    var bfs func(r, c int)

    bfs = func(r, c int) {
        p := point{r, c}
        seen[p] = true // 入队前标记节点已访问
        q := []point{p}

        for len(q) > 0 {
            p, q = q[0], q[1:]
            for _, ap := range adjacent(p) {
                if ap.row >= 0 && ap.row < m && ap.col >= 0 && ap.col < n && // 若邻接点合法才添加到队列, ①检查越界
                    grid[ap.row][ap.col] == '1' && !seen[ap] { // ②检查是否为陆地 ③检查是否访问过
                    seen[ap] = true // 入队前标记为已访问,  否则会重复添加节点,  造成死循环
                    q = append(q, ap)
                }
            }
        }
    }

    count := 0
    for r := 0; r < m; r++ {
        for c := 0; c < n; c++ {
            if grid[r][c] == '1' && !seen[point{r, c}] {
                bfs(r, c) // 标记整座岛为已访问
                count++   // 发现了一座新岛
            }
        }
    }
    return count
}

type point struct {
    row int
    col int
}

func adjacent(p point) []point {
    return []point{
        {p.row + 1, p.col},
        {p.row - 1, p.col},
        {p.row, p.col + 1},
        {p.row, p.col - 1},
    }
}
```

### [206. 反转链表](https://leetcode.cn/problems/reverse-linked-list/)

#### ➤ [题解](https://www.youtube.com/watch?v=G0_I-ZF0S38)

```go
func reverseList(head *ListNode) *ListNode {
    var cur = head
    var prev *ListNode

    for cur != nil {     // cur 最终停在 nil
        next := cur.Next // 暂存 next
        cur.Next = prev  // 完成 cur 的反转
        prev = cur       // prev 移动到 cur
        cur = next       // cur 移动到 next
    }

    return prev
}
```

### [207. 课程表](https://leetcode.cn/problems/course-schedule/)

#### ➤ [题解](https://www.youtube.com/watch?v=EgI5nU9etnU)

1. 先遍历所有边生成邻接表,  然后判断有向图是否存在环,  
2. 每个顶点执行一下 dfs,  用 visit 避免重复访问,  在 dfs 的过程中记录 path,  若跑回路径上的某点, 说明存在环

> A 课程依赖 B, 而 B 课程依赖 C, 请给出合适的学习顺序  
> 用 `A -> B` 表示 A 依赖 B,  这些依赖关系连起来就是一个有向图  
> 先判断有向图是否存在环,  若存在环则表示循环依赖,  此问题无解  
> 若不存在环,  `无环有向图` 的 `后序遍历` 就是学习顺序,  这个方法叫做拓扑排序

```go
func canFinish(numCourses int, prerequisites [][]int) bool {
    
    adjacent := make(map[int][]int) // 遍历每条边, 生成邻接表
    for _, edge := range prerequisites {
        adjacent[edge[0]] = append(adjacent[edge[0]], edge[1])
    }

    visit := make(map[int]bool)    // 记录处理完的顶点
    path := make(map[int]bool)     // 记录当前路径上的顶点
    var dfs func(v int) bool

    dfs = func(v int) bool {
        if path[v] {
            return true            // 访问的点在路径上已存在,  所以有环
        }

        path[v] = true             // 把顶点添加到路径
        for _, av := range adjacent[v] {
            if !visit[av] {        // 过滤访问过的邻接点
                if dfs(av) {       // 在访问完邻接点后, 发现有环
                    return true    // 那么直接返回有环, 不必继续递归
                }
            }
        }
        path[v] = false // 返回前清除路径
        visit[v] = true // 等邻接点都访问完了,  这个点也就访问完了
        return false    // 没有找到环
    }

    for v := range adjacent {      // 在每个顶点尝试 dfs 寻找环
        if dfs(v) {
            return false
        }
    }
    return true
}
```

### [208. 实现 Trie (前缀树)](https://leetcode.cn/problems/implement-trie-prefix-tree/)

#### ➤ [题解](https://www.youtube.com/watch?v=oobqoCJlHA0)

1. Trie（发音类似 "try"）或者说 前缀树 是一种树形数据结构
2. 这一数据结构有相当多的应用情景，例如自动补完和拼写检查
3. 因为只有 26 个字母,  所以可以用 array,  但用 map 更方便

```go
type TrieNode struct {
    endOfWord bool
    children  map[rune]*TrieNode
}

type Trie struct {
    root *TrieNode
}

func Constructor() Trie {
    // 注意初始化节点时, 节点中的 map 也要初始化,  写一个 nil map 会 panic
    return Trie{&TrieNode{false, make(map[rune]*TrieNode)}}
}

func (t *Trie) Insert(word string) {
    cur := t.root
    for _, c := range word {
        if cur.children[c] == nil { // 有 'a' 这个节点吗? 没有就初始化一下
            cur.children[c] = &TrieNode{false, make(map[rune]*TrieNode)}
        }
        cur = cur.children[c]
    }
    cur.endOfWord = true
}

func (t *Trie) Search(word string) bool {
    cur := t.root
    for _, c := range word {
        if cur.children[c] == nil { // 发现没有 'a' 这个节点就返回 false
            return false
        }
        cur = cur.children[c]
    }
    return cur.endOfWord
}

func (t *Trie) StartsWith(prefix string) bool {
    cur := t.root
    for _, c := range prefix {
        if cur.children[c] == nil { // 发现没有 'a' 这个节点就返回 false
            return false
        }
        cur = cur.children[c]
    }
    return true
}
```

### [215. 数组中的第K个最大元素](https://leetcode.cn/problems/kth-largest-element-in-an-array/)

#### ➤ [题解](https://www.youtube.com/watch?v=XEmy13g1Qxc&t=2s)

1. 排序后返回第 k 大的元素,  时间复杂度是 nlogn
2. 建堆, 然后弹出 k 次,  n + klogn
3. quick select,  平均 O(n),  最差 O(n^2),  随机选一个 piviot 进行 partition,  然后去左边或右边找 kth

```go
func findKthLargest(nums []int, k int) int {
    k = len(nums) - k // 第 k 大就是排序后倒数第 k 个, 把问题改成寻找索引为 k 的数,  后面会好处理一些
    var quickSelect func(l, r int) int

    quickSelect = func(l, r int) int {
        pivot := nums[r] // 取 [l, r] 区间的最后一个元素作为 pivot
        p := l           // p 表示小于等于区域的下一个位置
        for i := l; i < r; i++ {
            if nums[i] <= pivot {
                nums[p], nums[i] = nums[i], nums[p] // 往小于等于区域发货
                p++
            }
            // 这个 partition 只分成两个区域 <= 和 >,  有的 partition 会分成 < 和 = 和 > 三个区域
            // 遇到大于 pivot 的数则啥也不干,  只有 i++,  所以 p 最终停在大于区域
        }

        nums[p], nums[r] = nums[r], nums[p] // 交换后 p 就是 pivot 这个数在排序数组中的索引
        if k < p {
            return quickSelect(l, p-1)
        } else if k > p {
            return quickSelect(p+1, r)
        } else {
            return nums[p]
        }
    }

    return quickSelect(0, len(nums)-1)
}
```

### [221. 最大正方形](https://leetcode.cn/problems/maximal-square/)

#### ➤ [题解](https://www.youtube.com/watch?v=6X7Ha2PrDmM)

1. 定义 `dp[r][c]` 为以 (r,c) 为左上角的最大正方形的边长
2. 如果 `dp[r][c]` 是 1 并且它的右边/下面/右下角都是 1,  那么 `dp[r][c]` 是 2
3. 如果 `dp[r][c]` 是 1 并且它的右边/下面/右下角都是 2,  那么 `dp[r][c]` 是 3
4. 总之 `dp[r][c] = 1 + min(dp[r+1][c], dp[r][c+1], dp[r+1][c+1])` // 如果 `matrix[r][c]` 是 1 的话

```go
func maximalSquare(matrix [][]byte) int {
    // 额外增加一行一列,  dp 表的初始值全都是 0
    m, n := len(matrix), len(matrix[0])
    dp := make([][]int, m+1)
    for i := range dp {
        dp[i] = make([]int, n+1)
    }

    // 从最后一行到第一行,  从右往左填表
    res := 0
    for r := m - 1; r >= 0; r-- {
        for c := n - 1; c >= 0; c-- {
            if matrix[r][c] == '0' {
                dp[r][c] = 0
            } else {
                dp[r][c] = 1 + min(dp[r+1][c], dp[r][c+1], dp[r+1][c+1])
            }
            if dp[r][c] > res {
                res = dp[r][c]
            }
        }
    }
    return res * res
}

func min(nums ...int) int {
    res := nums[0]
    for _, n := range nums {
        if n < res {
            res = n
        }
    }
    return res
}
```

### [226. 翻转二叉树](https://leetcode.cn/problems/invert-binary-tree/)

```go
func invertTree(root *TreeNode) *TreeNode {
    var invert func(root *TreeNode)
    invert = func(root *TreeNode) {
        // base case
        if root == nil {
            return
        }
        // 在先序遍历时,  把左右子树互换,  然后进入左右子树继续互换
        root.Left, root.Right = root.Right, root.Left
        invert(root.Left)
        invert(root.Right)
    }
    invert(root)
    return root
}
```

### [234. 回文链表](https://leetcode.cn/problems/palindrome-linked-list/)

#### ➤ [题解](https://www.youtube.com/watch?v=yOzXms1J6Nk)

1. 把链表转成数组,  问题就简单了, 需要 O(n) 的额外空间
2. 把右半链表反转,  这么做只需 O(1) 的额外空间

```go
func isPalindrome(head *ListNode) bool {
    // fast 和 slow 同一起点的话,  slow 停在中间节点 or 靠右的中间节点
    fast, slow := head, head
    for fast != nil && fast.Next != nil {
        fast = fast.Next.Next
        slow = slow.Next
    }

    // 反转右半部分, 然后遍历右半部分并逐一比对
    rightHead := reverse(slow)
    left, right := head, rightHead

    res := true
    for right != nil {
        if left.Val != right.Val {
            res = false
            break
        }
        left = left.Next
        right = right.Next
    }

    reverse(rightHead) // 把链表还原会比较好,  避免副作用
    return res
}

func reverse(head *ListNode) *ListNode {
    var prev, cur *ListNode = nil, head
    for cur != nil {
        next := cur.Next // 暂存 next 节点
        cur.Next = prev  // 反转
        prev = cur       // 移动 prev 和 cur
        cur = next
    }
    return prev
}
```

### [235. 二叉搜索树的最近公共祖先](https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-search-tree/)

#### ➤ [题解](https://www.youtube.com/watch?v=gs2LMfuOR9k)

1. 从 `root` 出发, 若 `p` 和 `q` 分别属于不同的子树,  那么 `root` 就是最近公共祖先
2. 如果 `p` 和 `q` 都属于左子树,  那么 `root.left` 就是比 `root` 更近的公共祖先
3. 如果 `p` 或 `q` 刚好等于 `root`,  那么最近公共祖先就是 `root`, 比如 [1 2 3] 中 1 和 2 的最近公共祖先是 1

```go
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    cur := root
    for cur != nil {
        if p.Val < cur.Val && q.Val < cur.Val {
            cur = cur.Left
        } else if p.Val > cur.Val && q.Val > cur.Val {
            cur = cur.Right
        } else {
            return cur
        }
    }
    return nil
}
```

### [236. 二叉树的最近公共祖先](https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/)

1. 参照题目中的示例图 1,  总共有两种情况
2. 若寻找 7 和 4,  那么会返回 2 然后一直把 2 往上传
3. 若寻找 2 和 7,  那么会返回 2 然后一直把 2 往上传

```go
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    if root == nil {
        return nil // dfs 的 base case
    }
    if root.Val == p.Val || root.Val == q.Val {
        return root // 找到了 p、q 子树则返回它
    }

    left := lowestCommonAncestor(root.Left, p, q)
    right := lowestCommonAncestor(root.Right, p, q)

    // 某节点的左子树包含 p 且右子树包含 q,  那么它就是最近公共祖先
    if left != nil && right != nil {
        return root
    }
    // 把找到的 p、q 子树往上层传递
    if left != nil {
        return left
    }
    if right != nil {
        return right
    }
    return nil
}
```

### [238. 除自身以外数组的乘积](https://leetcode.cn/problems/product-of-array-except-self/)

#### ➤ [题解](https://www.youtube.com/watch?v=bNvIQI2wAjk)

1. 两次遍历,  把前缀数组和后缀数组的乘积全都算出来,  然后把前缀和后缀组合起来就好
2. 先把前缀数组算出来,  然后从后往前遍历,  在遍历过程中也能知道后缀是什么
3. 然后再把前缀数组当成输出数组来用,  这么做对比上一种方法则不需要额外的存储空间

```go
func productExceptSelf(nums []int) []int {
    prefix := make([]int, len(nums)) // 定义 prefix[i] 为 i 这个位置的前缀的乘积
    prefix[0] = 1                    // 虽然 0 这个位置没有前缀,  不妨把它设成 1

    // 从第二项开始
    for i := 1; i < len(prefix); i++ {
        prefix[i] = prefix[i-1] * nums[i-1]
    }

    // 用 i 的前缀乘以 i 的后缀
    suffix := 1
    for i := len(nums) - 1; i >= 0; i-- {
        prefix[i] = prefix[i] * suffix
        suffix *= nums[i]
    }

    return prefix
}
```

### [239. 滑动窗口最大值](https://leetcode.cn/problems/sliding-window-maximum/)

#### ➤ [题解](https://www.youtube.com/watch?v=DfljaUwZsOk)

```go
func maxSlidingWindow(nums []int, k int) []int {
    l, r := 0, 0
    q := make([]int, 0) // 一个单调递减的双端队列, 存的是索引
    res := make([]int, 0)

    for r < len(nums) {
        // 只要比栈顶大就一直弹栈, 如果 <= 栈顶则压栈
        for len(q) > 0 && nums[r] > nums[q[len(q)-1]] {
            q = q[:len(q)-1]
        }
        q = append(q, r)

        if r >= k-1 {
            res = append(res, nums[q[0]]) // 在移动窗口前,  把最大值加到结果里
            l++                           // 等窗口长度达到 k 才移动左指针
        }
        r++

        if l > q[0] {                     // 移动左指针会让最大值过期
            q = q[1:]
        }

    }
    return res
}
```

### [240. 搜索二维矩阵 II](https://leetcode.cn/problems/search-a-2d-matrix-ii/)

```go
func searchMatrix(matrix [][]int, target int) bool {
    // 从右上角开始,  往左走变小,  往右走变大
    m, n := len(matrix), len(matrix[0])
    r, c := 0, n-1

    for r < m && c >= 0 {
        if target < matrix[r][c] {
            c-- // 需要更小的数,  只能往左走,  因为下面的数都更大,  这样就排除了一列
        } else if target > matrix[r][c] {
            r++
        } else {
            return true
        }
    }
    return false
}
```

### [279. 完全平方数](https://leetcode.cn/problems/perfect-squares/)

#### ➤ [题解](https://www.youtube.com/watch?v=HLZLwjzIVGo)

1. 先画决策树,  分别选择 1/4/9/...,  能发现有重叠子问题,  比如通过不同的路径到达 5 这个状态
2. `dp[i] = 1 + min(dp[i-1], dp[i-4], dp[i-9], ...)`

```go
func numSquares(n int) int {
    dp := make([]int, n+1)
    for i := range dp {
        dp[i] = n // 最多用 n 个 1 来凑 n
    }
    dp[0] = 0

    for i := 1; i <= n; i++ {
        for j := 1; j*j <= i; j++ {
            dp[i] = min(dp[i], 1+dp[i-j*j])
        }
    }
    return dp[n]
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
```

### [283. 移动零](https://leetcode.cn/problems/move-zeroes/)

#### ➤ [题解](https://www.youtube.com/watch?v=aayNRwUN3Do)

1. 可以创建两个数组,  把数组分成两组再合并回一个,  但这么做需要 O(n) 的空间
2. 用 partition

```go
func moveZeroes(nums []int) {
    l := 0
    for i := range nums {
        // 遇到为 0 的数不变,  否则和 l 交换
        if nums[i] != 0 {
            nums[l], nums[i] = nums[i], nums[l]
            l++
        }
    }
}
```

### [287. 寻找重复数](https://leetcode.cn/problems/find-the-duplicate-number/)

#### ➤ [题解](https://www.youtube.com/watch?v=wjYnzkAhcNk)

1. 把数组的格子看成节点,  把格子中的值视作指针,  整个数组就能视作链表,  然后问题变成有环链表寻找入环点

```go
func findDuplicate(nums []int) int {
    fast, slow := 0, 0
    // nums[slow] 就是 slow 这个节点的 next 指针
    for {
        slow = nums[slow]
        fast = nums[fast]
        fast = nums[fast]
        // 和环形链表那道题的做法一样,  寻找入环点
        if fast == slow {
            fast = 0
            for fast != slow {
                slow = nums[slow]
                fast = nums[fast]
            }
            return fast
        }
    }
}
```

### [300. 最长递增子序列](https://leetcode.cn/problems/longest-increasing-subsequence/)

#### ➤ [题解](https://www.youtube.com/watch?v=cjWnW0hdF1Y)

1. 定义 `dp[i]` 表示以 `nums[i]` 作为开头的最长递增子序列
2. `dp[i] = 1 + max(dp[i+1], dp[i+2], dp[i+3], ...)`,  此处 `dp[i+1]` 有效的前提是 `nums[i] < nums[i+1]`

```go
func lengthOfLIS(nums []int) int {
    n := len(nums)
    dp := make([]int, n)
    for i := range dp {
        dp[i] = 1             // 若初始化为 0 会导致错误的结果
    }

    res := 0
    for i := n - 1; i >= 0; i-- {
        for j := 1; i+j < n; j++ {
            if nums[i] < nums[i+j] {
                dp[i] = max(dp[i], 1+dp[i+j])
            }
        }
        res = max(res, dp[i]) // 注意以 nums[0] 开头的最长递增子序列并非答案
    }

    return res
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```

### [309. 最佳买卖股票时机含冷冻期](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-cooldown/)

#### ➤ [题解](https://www.youtube.com/watch?v=I7j0F7AHpb8)

1. 画决策树,  从第一天开始,  状态为利润
2. 有 buy/cooldown 两个选择,  选了 buy 之后有 sell/cooldown 两个选择...

```go
type SubProblem struct {
    i      int
    buying bool
}

func maxProfit(prices []int) int {
    cache := make(map[SubProblem]int)
    var dfs func(i int, buying bool) int

    // 用 dfs(i, buying) 表示子问题能挣多少钱
    dfs = func(i int, buying bool) int {
        sub := SubProblem{i, buying}
        if i >= len(prices) {
            return 0
        }
        if answer, ok := cache[sub]; ok {
            return answer
        }

        if buying {
            // buying 阶段要么 buy,  要么 cooldown
            buy := dfs(i+1, false) - prices[i]
            cooldown := dfs(i+1, true)
            cache[sub] = max(buy, cooldown)
        } else {
            sell := dfs(i+2, true) + prices[i]
            cooldown := dfs(i+1, false)
            cache[sub] = max(sell, cooldown)
        }
        return cache[sub]
    }

    return dfs(0, true)
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```

### [322. 零钱兑换](https://leetcode.cn/problems/coin-change/)

1. 定义 dp[i] 表示凑齐 i 最少要 dp[i] 个硬币
2. `dp[i] = 1 + min(dp[i-coin1], dp[i-coin2], dp[i-coin3], ...)`

```go
func coinChange(coins []int, amount int) int {
    max := amount + 1
    dp := make([]int, amount+1)
    for i := range dp {
        dp[i] = max // 最多为 amount 个 1, 不可能到 max
    }
    dp[0] = 0 // base case

    for i := 1; i <= amount; i++ {
        for j := 0; j < len(coins); j++ {
            if i-coins[j] >= 0 {
                dp[i] = min(dp[i], 1+dp[i-coins[j]]) // 就算 dp[i-coins[j]] 无解也不怕
            }
        }
    }

    if dp[amount] == max {
        return -1
    }
    return dp[amount]
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
```

### [337. 打家劫舍 III](https://leetcode.cn/problems/house-robber-iii/)

#### ➤ [题解](https://www.youtube.com/watch?v=nHR8ytpzz7c)

```go
func rob(root *TreeNode) int {
    var dfs func(root *TreeNode) (int, int)

    dfs = func(root *TreeNode) (int, int) {
        if root == nil {
            return 0, 0
        }

        withL, withoutL := dfs(root.Left) // 返回打劫 root 和不打劫 root 分别能挣多少钱
        withR, withoutR := dfs(root.Right)

        return root.Val + withoutL + withoutR, max(withL, withoutL) + max(withR, withoutR)
    }
    return max(dfs(root))
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```

### [338. 比特位计数](https://leetcode.cn/problems/counting-bits/)

#### ➤ [题解](https://www.youtube.com/watch?v=RyBM56RIWrM)

1. O(nlogn) 解法是不断的除以 2 取余,  若余数是 1 则 1 的个数加一
2. 注意 0~3 和 4~7 除了最高位不同, 剩下部分是重复的: `dp[i] = 1 + dp[i - offset]`

```go
func countBits(n int) []int {
    dp := make([]int, n+1)
    offset := 1
    for i := 1; i <= n; i++ {
        if offset*2 == i {
            offset = i
        }
        dp[i] = 1 + dp[i-offset]
    }
    return dp
}
```

### [347. 前 K 个高频元素](https://leetcode.cn/problems/top-k-frequent-elements/)

#### ➤ [题解](https://www.youtube.com/watch?v=YPTqKIgVk-k)

1. 统计每个数的出现次数,  然后用小根堆存储最大的 k 个元素,  O(nlogk)
2. 统计每个数的出现次数,  然后按出现次数进行分组,  然后依次看有没有出现了 6/5/4/... 次的数

```go
func topKFrequent(nums []int, k int) []int {
    // 统计频次
    count := make(map[int]int)
    for _, n := range nums {
        count[n]++
    }

    // 按频次分组
    freq := make([][]int, len(nums)+1)
    for n, c := range count {
        freq[c] = append(freq[c], n)
    }

    // 依次看, 有没有出现 n/n-1/n-2/... 次的数
    res := make([]int, 0, k)
    for i := len(freq) - 1; i >= 0; i-- {
        for _, n := range freq[i] {
            res = append(res, n)
            if len(res) == k {
                return res
            }
        }
    }
    return nil
}
```

### [394. 字符串解码](https://leetcode.cn/problems/decode-string/)

#### ➤ [题解](https://www.youtube.com/watch?v=qB0zZpBJlh8)

```go
func decodeString(s string) string {
    stack := make([]string, 0)
    for i := range s {
        if s[i] != ']' {
            stack = append(stack, s[i:i+1])
        } else {
            // 遇到 ] 则不停弹栈,  获取一个字符串
            substr := ""
            for stack[len(stack)-1] != "[" {
                substr = stack[len(stack)-1] + substr
                stack = stack[:len(stack)-1]
            }
            stack = stack[:len(stack)-1]

            // 获取 [] 前的数字
            k := ""
            for len(stack) > 0 && stack[len(stack)-1] >= "0" && stack[len(stack)-1] <= "9" {
                k = stack[len(stack)-1] + k
                stack = stack[:len(stack)-1]
            }

            // 重复 k 次然后压栈
            count, _ := strconv.Atoi(k)
            stack = append(stack, strings.Repeat(substr, count))
        }
    }
    return strings.Join(stack, "")
}
```

### [416. 分割等和子集](https://leetcode.cn/problems/partition-equal-subset-sum/)

#### ➤ [题解](https://www.youtube.com/watch?v=IsvocB5BJhw)

```go
func canPartition(nums []int) bool {
    // sum 为奇数则不能分成相等的两份
    sum := 0
    for _, n := range nums {
        sum += n
    }
    if sum%2 != 0 {
        return false
    }

    // 用集合记录当前可以凑出来的数
    target := sum / 2
    set := make(map[int]bool)
    set[0] = true

    for i := 0; i < len(nums); i++ {
        newSet := make(map[int]bool, len(set))
        for n := range set {
            newSet[n] = true
            newSet[n+nums[i]] = true
        }
        set = newSet
        if set[target] {
            return true
        }
    }
    return false
}
```

### [437. 路径总和 III](https://leetcode.cn/problems/path-sum-iii/)

#### ➤ [题解](https://leetcode.cn/problems/path-sum-iii/solution/dui-qian-zhui-he-jie-fa-de-yi-dian-jie-s-dey6/)

```go
func pathSum(root *TreeNode, targetSum int) int {
    res := 0
    prefixSum := map[int]int{0: 1}
    var dfs func(root *TreeNode, cur int)

    dfs = func(root *TreeNode, cur int) {
        if root == nil {
            return
        }
        cur += root.Val
        res += prefixSum[cur-targetSum]
        prefixSum[cur]++
        dfs(root.Left, cur)
        dfs(root.Right, cur)
        prefixSum[cur]-- // 返回了,  当前这个前缀和就不存在了
    }

    dfs(root, 0)
    return res
}
```

### [438. 找到字符串中所有字母异位词](https://leetcode.cn/problems/find-all-anagrams-in-a-string/)

#### ➤ [题解](https://www.youtube.com/watch?v=G8xtZy0fDKg)

1. 维护滑动窗口中各字符的出现次数,  记录当前已满足几个条件

```go
func findAnagrams(s string, p string) []int {
    count := make(map[byte]int)
    for i := range p {
        count[p[i]]++
    }

    res := make([]int, 0)
    window := make(map[byte]int)
    l, r := 0, 0
    have, need := 0, len(count)

    for r < len(s) {
        // 更新窗口内数据
        c := s[r]
        window[c]++
        if window[c] == count[c] {
            have++
        }
        if have == need {
            res = append(res, l)
        }

        // 移动窗口,  并更新窗口内数据
        if r >= len(p)-1 {
            c := s[l]
            if window[c] == count[c] {
                have--
            }
            window[c]--
            l++
        }
        r++
    }

    return res
}
```

### [448. 找到所有数组中消失的数字](https://leetcode.cn/problems/find-all-numbers-disappeared-in-an-array/)

#### ➤ [题解](https://www.youtube.com/watch?v=8i-f24YFWC4)

```go
func findDisappearedNumbers(nums []int) []int {
    // 把 1 映射到索引 0,  把 n 映射到索引 n-1
    // 添加一个负号表示当前索引对应的数已经出现过
    for _, n := range nums {
        n := int(math.Abs(float64(n)))
        nums[n-1] = -int(math.Abs(float64(nums[n-1])))
    }

    res := make([]int, 0)
    for i := range nums {
        if nums[i] > 0 {
            res = append(res, i+1)
        }
    }
    return res
}
```

### [461. 汉明距离](https://leetcode.cn/problems/hamming-distance/)

```go
func hammingDistance(x int, y int) int {
    // return bits.OnesCount(uint(x^y))

    n := x ^ y
    res := 0
    for n != 0 {
        res += n & 1
        n >>= 1 // 右移一位, 统计 1 的个数
    }
    return res
}
```

### [494. 目标和](https://leetcode.cn/problems/target-sum/)

#### ➤ [题解](https://www.youtube.com/watch?v=g0npyaQtAQM)

```go
func findTargetSumWays(nums []int, target int) int {
    cache := make(map[[2]int]int)
    var backtrack func(i, total int) int

    backtrack = func(i, total int) int {
        if i == len(nums) {
            if total == target {
                return 1
            }
            return 0
        }
        sub := [2]int{i, total}
        if answer, ok := cache[sub]; ok {
            return answer
        }
        // 两个选择分别为加上/减去 nums[i]
        cache[sub] = backtrack(i+1, total+nums[i]) + backtrack(i+1, total-nums[i])
        return cache[sub]
    }

    return backtrack(0, 0)
}
```

### [538. 把二叉搜索树转换为累加树](https://leetcode.cn/problems/convert-bst-to-greater-tree/)

#### ➤ [题解](https://www.youtube.com/watch?v=7vVEJwVvAlI)

```go
func convertBST(root *TreeNode) *TreeNode {
    curSum := 0
    var dfs func(root *TreeNode)
    dfs = func(root *TreeNode) {
        if root == nil {
            return
        }
        // 从右向左的中序遍历
        dfs(root.Right)
        curSum += root.Val
        root.Val = curSum
        dfs(root.Left)
    }
    dfs(root)
    return root
}
```

### [543. 二叉树的直径](https://leetcode.cn/problems/diameter-of-binary-tree/)

#### ➤ [题解](https://www.youtube.com/watch?v=bkxqA8Rfv04)

```go
func diameterOfBinaryTree(root *TreeNode) int {
    res := 0
    var dfs func(root *TreeNode) int
    dfs = func(root *TreeNode) int {
        if root == nil {
            return 0
        }
        leftDepth := dfs(root.Left)
        rightDepth := dfs(root.Right)
        res = max(res, leftDepth+rightDepth)
        return 1 + max(leftDepth, rightDepth) // 返回当前子树的最大深度
    }
    dfs(root)
    return res
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```

### [560. 和为 K 的子数组](https://leetcode.cn/problems/subarray-sum-equals-k/)

#### ➤ [题解](https://www.youtube.com/watch?v=fFVZt-6sgyo)

1. 用 prefixSum 记录所有的前缀和,  如果 curSum 为 10 且有两个前缀的和为 3,  那么就有两个子数组和为 7

```go
func subarraySum(nums []int, k int) int {
    res := 0
    curSum := 0
    prefixSum := map[int]int{0: 1}
    for _, n := range nums {
        curSum += n
        res += prefixSum[curSum-k]
        prefixSum[curSum]++
    }
    return res
}
```

### [581. 最短无序连续子数组](https://leetcode.cn/problems/shortest-unsorted-continuous-subarray/)

#### ➤ [题解](https://leetcode.cn/problems/shortest-unsorted-continuous-subarray/solution/si-lu-qing-xi-ming-liao-kan-bu-dong-bu-cun-zai-de-/)

```go
func findUnsortedSubarray(nums []int) int {
    n := len(nums)
    maxN, minN := math.MinInt, math.MaxInt
    l, r := -1, -1
    for i := range nums {
        maxN = max(maxN, nums[i]) // 只要在有序部分就有 nums[i] == maxN
        if nums[i] < maxN {       // 若在无序部分则不断更新右边界
            r = i
        }

        minN = min(minN, nums[n-1-i])
        if nums[n-1-i] > minN {
            l = n - 1 - i
        }
    }
    if r == -1 {
        return 0
    }
    return r - l + 1
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
```

### [617. 合并二叉树](https://leetcode.cn/problems/merge-two-binary-trees/)

#### ➤ [题解](https://www.youtube.com/watch?v=QHH6rIK3dDQ)

```go
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
    if root1 == nil { // 子树和 nil 合并的结果是子树本身
        return root2
    }
    if root2 == nil {
        return root1
    }

    root := &TreeNode{Val: root1.Val + root2.Val}
    root.Left = mergeTrees(root1.Left, root2.Left) // 两者的左子树合并后挂到左子树
    root.Right = mergeTrees(root1.Right, root2.Right)
    return root
}
```

### [647. 回文子串](https://leetcode.cn/problems/palindromic-substrings/)

1. 用中心扩散,  对于每个 i 都从 `(i, i)` 和 `(i, i+1)` 两个中心进行扩散,  就能找到所有回文串

```go
func countSubstrings(s string) int {
    res := 0
    for i := range s {
        res += expand(s, i, i)
        res += expand(s, i, i+1)
    }
    return res
}

func expand(s string, l, r int) int {
    for l >= 0 && r < len(s) && s[l] == s[r] {
        l--
        r++
    }

    length := r - l - 1
    if length%2 == 0 {
        return length / 2
    }
    return length/2 + 1
}
```

### [739. 每日温度](https://leetcode.cn/problems/daily-temperatures/)

#### ➤ [题解](https://www.youtube.com/watch?v=cTBiBSnjO3c)

1. 维护单调递减的栈,  nums[i] <= top 则压栈,  nums[i] > top 则弹栈并计算输出值

```go
func dailyTemperatures(temperatures []int) []int {
    stack := make([][2]int, 0)
    res := make([]int, len(temperatures))

    for i, n := range temperatures {
        // 当前数大于栈顶,  则弹栈并计算输出
        for len(stack) > 0 && n > stack[len(stack)-1][1] {
            top := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            res[top[0]] = i - top[0]
        }
        stack = append(stack, [2]int{i, n})
    }
    return res
}
```
