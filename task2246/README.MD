https://leetcode.com/problems/binary-tree-tilt

**2246. Longest Path With Different Adjacent Characterst**

```markdown
You are given a tree (i.e. a connected, undirected graph that has no cycles) rooted at node 0 consisting of n nodes numbered from 0 to n - 1. The tree is represented by a 0-indexed array parent of size n, where parent[i] is the parent of node i. Since node 0 is the root, parent[0] == -1.

You are also given a string s of length n, where s[i] is the character assigned to node i.

Return the length of the longest path in the tree such that no pair of adjacent nodes on the path have the same character assigned to them.

 

Example 1:
```
![img.png](img.png)
```

Input: parent = [-1,0,0,1,1,2], s = "abacbe"
Output: 3
Explanation: The longest path where each two adjacent nodes have different characters in the tree is the path: 0 -> 1 -> 3. The length of this path is 3, so 3 is returned.
It can be proven that there is no longer path that satisfies the conditions. 
Example 2:

```
![img_1.png](img_1.png)
```
Input: parent = [-1,0,0,0], s = "aabc"
Output: 3
Explanation: The longest path where each two adjacent nodes have different characters is the path: 2 -> 0 -> 3. The length of this path is 3, so 3 is returned.
 

Constraints:

n == parent.length == s.length
1 <= n <= 105
0 <= parent[i] <= n - 1 for all i >= 1
parent[0] == -1
parent represents a valid tree.
s consists of only lowercase English letters.
```

**Solution**

```
Step 1. Find two longest path out of n children for each node.
Step 2. Update global max value if longestPath1+longestPath2> max.
Step 3. Return longest path for each child node to parent node using DFS.

If adjacent node letters are not equal set pathlength for children = pathlength+1.
If adjacent node letters are equal set pathlength for children = 0.

```