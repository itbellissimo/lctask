https://leetcode.com/problems/maximum-ice-cream-bars/description/

**1833. Maximum Ice Cream Bars**

```markdown
It is a sweltering summer day, and a boy wants to buy some ice cream bars.

At the store, there are n ice cream bars. You are given an array costs of length n, where costs[i] is the price of the ith ice cream bar in coins. The boy initially has coins coins to spend, and he wants to buy as many ice cream bars as possible. 

Note: The boy can buy the ice cream bars in any order.

Return the maximum number of ice cream bars the boy can buy with coins coins.
```
!!!!!!!!!!!!!!!!**You must solve the problem by counting sort.**!!!!!!!!!!!!!!!


```markdown
Example 1:

Input: costs = [1,3,2,4,1], coins = 7
Output: 4
Explanation: The boy can buy ice cream bars at indices 0,1,2,4 for a total price of 1 + 3 + 2 + 1 = 7.
Example 2:

Input: costs = [10,6,8,7,7,8], coins = 5
Output: 0
Explanation: The boy cannot afford any of the ice cream bars.
Example 3:

Input: costs = [1,6,3,1,2,5], coins = 20
Output: 6
Explanation: The boy can buy all the ice cream bars for a total price of 1 + 6 + 3 + 1 + 2 + 5 = 18.
 

Constraints:

costs.length == n
1 <= n <= 105
1 <= costs[i] <= 105
1 <= coins <= 108
```


**Solution**

https://leetcode.com/problems/maximum-ice-cream-bars/solutions/2885725/maximum-ice-cream-bars/?orderBy=most_votes
See Approach 2: Counting Sort (Greedy)


A comparison-based sorting method (like heapsort, mergesort, etc.) takes $(n \log n)$ time. However, using counting sort, we can access the elements in sorted order in linear time.

Counting sort is a sorting technique that is based on the keys between specific ranges. We store each element's frequency in an array and thus using this new array we can access all elements in sorted order.
As the input array's element's range is not very large, we can use counting sort here. If you are new to counting sort, then we recommend you take a look at it in our Sorting Explore Card.