leetcode 第104题

给定一个二叉树，找出其最大深度。



#### 递归解法：

##### java

~~~ java
public int maxDepth(Treenode root){
    if(root == null){
        return 0;
    }else{
        int left = maxDepth(root.left);
        int right = maxDepth(root.right);
        return Math.max(left,right) + 1;
    }
}
~~~

