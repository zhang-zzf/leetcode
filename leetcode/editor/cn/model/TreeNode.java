package model;

import com.alibaba.fastjson.JSON;

import java.util.LinkedList;
import java.util.List;
import java.util.Queue;

public class TreeNode {

    public int val;
    public TreeNode left;
    public TreeNode right;

    public static TreeNode decode(String nodesStr) {
        List<Integer> nodes = JSON.parseArray(nodesStr, Integer.class);
        if (nodes.isEmpty()) {
            return null;
        }
        TreeNode root = new TreeNode(nodes.remove(0));
        Queue<TreeNode> queue = new LinkedList<>();
        queue.add(root);
        while (!queue.isEmpty()) {
            TreeNode r = queue.poll();
            if (r == null) {
                continue;
            }
            boolean leftFilled = false;
            for (int i = 0; i < Math.min(2, nodes.size()); i++) {
                Integer val = nodes.remove(0);
                TreeNode node = val == null ? null : new TreeNode(val);
                if (!leftFilled) {
                    r.left = node;
                    leftFilled = true;
                } else {
                    r.right = node;
                }
                queue.offer(node);
            }
        }
        return root;
    }

    public TreeNode(int x) {
        val = x;
    }

    public TreeNode(int x, TreeNode left, TreeNode right) {
        this.val = x;
        this.left = left;
        this.right = right;
    }

}
