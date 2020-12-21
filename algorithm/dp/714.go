package main

// 714. 买卖股票的最佳时机含手续费
// 给定一个整数数组 prices，其中第 i 个元素代表了第 i 天的股票价格 ；非负整数 fee 代表了交易股票的手续费用。

// 你可以无限次地完成交易，但是你每笔交易都需要付手续费。如果你已经购买了一个股票，在卖出它之前你就不能再继续购买股票了。

// 返回获得利润的最大值。

// 注意：这里的一笔交易指买入持有并卖出股票的整个过程，每笔交易你只需要为支付一次手续费。

// 示例 1:

// 输入: prices = [1, 3, 2, 8, 4, 9], fee = 2
// 输出: 8
// 解释: 能够达到的最大利润:
// 在此处买入 prices[0] = 1
// 在此处卖出 prices[3] = 8
// 在此处买入 prices[4] = 4
// 在此处卖出 prices[5] = 9
// 总利润: ((8 - 1) - 2) + ((9 - 4) - 2) = 8.
// 注意:

// 0 < prices.length <= 50000.
// 0 < prices[i] < 50000.
// 0 <= fee < 50000.
// 通过次数59,397提交次数85,259

//提示  贪心算法 动态规划

//贪心
//每次购买的时候都加上利息费用
func maxProfit(prices []int, fee int) int {
	n := len(prices)
	buy := prices[0] + fee
	profit := 0
	for i := 1; i < n; i++ {
		if prices[i]+fee < buy {
			buy = prices[i] + fee
		} else if prices[i] > buy {
			profit += prices[i] - buy
			buy = prices[i]
		}
	}
	return profit
}

// 动态规划方程
//状态转移方程 dp[i]>dp[i-1]+2

// 定义二维数组 dp[n][2]dp[n][2]：

// dp[i][0] 表示第 ii 天不持有可获得的最大利润；
// dp[i][1] 表示第 ii 天持有可获得的最大利润（注意是第 ii 天持有，而不是第 ii 天买入）。
// 定义状态转移方程：

// 不持有：dp[i][0] = max(dp[i - 1][0], dp[i - 1][1] + prices[i] - fee)

// 对于今天不持有，可以从两个状态转移过来：1. 昨天也不持有；2. 昨天持有，今天卖出。两者取较大值。

// 持有：dp[i][1] = max(dp[i - 1][1], dp[i - 1][0] - prices[i])

// 对于今天持有，可以从两个状态转移过来：1. 昨天也持有；2. 昨天不持有，今天买入。两者取较大值。

func maxProfitDp(prices []int, fee int) int {
	n := len(prices)
	dp := make([][2]int, n)
	dp[0][1] = -prices[0]
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i]-fee)
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[n-1][0]
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}