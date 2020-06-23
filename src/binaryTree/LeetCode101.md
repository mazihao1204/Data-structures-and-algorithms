leetcode 101  对称二叉树

给定一个二叉树，检查它是否镜像对称的

递归解法

~~~java
public boolean isSymmetric(TreeNode root){
    return check(root,root)
}

public boolean check(TreeNode p,TreeNode q){
	if(p == null && q == null){
        return true;
    }
    if(p == null || q == null){
        return false;
    }
    return p.val == q.val && check(p.left,q.right) && check(p.right,q.left);
}
~~~

