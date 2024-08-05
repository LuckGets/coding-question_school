function maxProfit(prices: number[]): number {
  let profit: number = 0;
  let buy: number = prices[0];
  for (let i = 1; i < prices.length; i++) {
    if (prices[i] < buy) {
      buy = prices[i];
    } else if (prices[i] - buy > profit) {
      profit = prices[i] - buy;
    }
  }
  return profit;
}
