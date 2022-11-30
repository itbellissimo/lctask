https://leetcode.com/problems/number-of-provinces/description/

There are n cities. Some of them are connected, while some are not. 
If city a is connected directly with city b, and city b is connected directly with city c, 
then city a is connected indirectly with city c.

A province is a group of directly or indirectly connected cities and no other cities outside of the group.

You are given an n x n matrix isConnected where isConnected[i][j] = 1 
if the ith city and the jth city are directly connected, 
and isConnected[i][j] = 0 otherwise.

Return the total number of provinces.



**Keywords**

DFS or Depth First Search

**Links:**

https://en.wikipedia.org/wiki/Depth-first_search
https://www.programiz.com/dsa/graph-dfs


There are two realizations. Classical - with recursion.

And second one - without recursion. 
The easiest way is to store on the stack don't visited neighboring nodes. This stack is used to future visits.
The second way without recursion is save on stack nodes indexes and it's neighboring node indexes.
Third one is realize graph on pointers

**Depth First Search Algorithm**
```
A standard DFS implementation puts each vertex of the graph into one of two categories:

Visited
Not Visited
The purpose of the algorithm is to mark each vertex as visited while avoiding cycles.

The DFS algorithm works as follows:

Start by putting any one of the graph's vertices on top of a stack.
Take the top item of the stack and add it to the visited list.
Create a list of that vertex's adjacent nodes. Add the ones which aren't in the visited list to the top of the stack.
Keep repeating steps 2 and 3 until the stack is empty.
```


**Golang:**
Three ways of taking input

1. fmt.Scanln(&input)

2. reader.ReadString()

3. scanner.Scan()

Here we recommend using bufio.NewScanner


