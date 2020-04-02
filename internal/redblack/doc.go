package redblack

/*
It is a special BST (Binary Search Tree), where:
+ Each node is either BLACK or RED
+ Root is BLACK
+ All leaves (nil children) are BLACK
+ Every RED node has two BLACK children
+ Every path from root to leaf has the same number of BLACK nodes, the number also called "black-height"

It is not a completed balanced BST, but it guarantees an almost-balanced BST:
+ It sacrifices balanced-hight for a O(log(n)) Insert/Delete/Search
+ Any un-almost-balanced situation caused by Insert/Delete can be fix in maximum 3 times rotation

eg: JDK.TreeMap, JDK.TreeSet
*/
