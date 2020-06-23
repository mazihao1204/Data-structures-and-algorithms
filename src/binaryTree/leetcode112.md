#### leetcode 112 路劲总和

给定一个二叉树和一个目标和，判断该树中是否存在根节点到叶子节点的路劲，这条路劲上所有节点值相加等于目标和。

说明：叶子节点是指没有子节点的节点

----

**递归解法**

~~~java
public boolean hasPathSum(TreeNode root, int sum) {
	if(root == null){
        return false;
    }
    sum -= root.val;
    if((root.left == null) && (root.right == null)){
        return (sum == 0);
    }
    return hasPathSum(root.left,sum) || hasPathSum(root.right,sum);
}
~~~

