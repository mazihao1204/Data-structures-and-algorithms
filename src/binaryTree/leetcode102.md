leetcode 102题

二叉树的层序遍历



#### 递归

##### js代码

~~~javascript
const levelOrder = function(root){
    if(root === null){
        return []
    }
    let res = []
    const dfs = function(node,level){
    	if(res.length === level){
            res.push([])
        }
        res[level].push(node.val)
        if(node.left){
            dfs(node.left,level+1)
        }
        if(node.right){
            dfs(node.right,level+1)
        }
	}
    dfs(root,0)
    return res
}

~~~



#### BFS

##### Java代码

~~~java
public List<List<Integer>> levelOrder(TreeNode root) {
        Queue<TreeNode> queue = new LinkedList<TreeNode>();
        List<List<Integer>> wrapList = new LinkedList<List<Integer>>();
        if(root == null) {
            return wrapList;
        }
        queue.offer(root);
        while (!queue.isEmpty()){
            int levelNum = queue.size();
            List<Integer> subList = new LinkedList<Integer>();
            for (int i = 0;i<levelNum;i++){
                TreeNode node = queue.poll();
                subList.add(node.val);
                if(node.left != null){
                    queue.add(node.left);
                }
                if(node.right != null) {
                    queue.add(node.right);
                }
            }
            wrapList.add(subList);
        }
        return wrapList;
    }
~~~

