package com.devandrew._872;

import java.util.Optional;
import java.util.Stack;

public class _872 {
    public static TreeNode fromArray(Integer[] arr, int idx) {
        if (idx < arr.length) {
            TreeNode left = _872.fromArray(arr, idx * 2 + 1);
            TreeNode right = _872.fromArray(arr, idx * 2 + 2);
            return arr[idx] == null ? null : new TreeNode(arr[idx], left, right);
        }
        return null;
    }

    private static class TreeNode {
        int val;
        TreeNode left;
        TreeNode right;
        public TreeNode() {}
        public TreeNode(int val) { this.val = val; }
        public TreeNode(int val, TreeNode left, TreeNode right) {
            this.val = val;
            this.left = left;
            this.right = right;
        }
    }
    private static class TreeIterator {
        private enum Dir { Left, Right }
        private static class StackNode {
            public TreeNode node;
            public Dir dir;
            public StackNode(TreeNode node, Dir dir) {
                this.node = node;
                this.dir = dir;
            }
        }

        private final Stack<StackNode> stack;
        private TreeNode current;

        public TreeIterator(TreeNode root) {
            stack = new Stack<>();
            current = down(root);
        }

        public boolean hasNext() {
            return current != null;
        }


        public int getNext() {
            int res = this.current.val;
            this.current = null;
            while (!stack.isEmpty() && this.current == null) {
                StackNode intersection = this.up();
                if (intersection != null && intersection.node.right != null) {
                    stack.add(new StackNode(intersection.node, Dir.Right));
                    this.current = down(intersection.node.right);
                }
            }
            return res;
        }

        public StackNode up() {
            StackNode node = stack.pop();
            while (node != null && node.dir == Dir.Right) {
                node = stack.isEmpty() ? null : stack.pop();
            }
            return node;
        }

        public TreeNode down(TreeNode node) {
            while (node != null && (node.left != null || node.right != null)) {
                stack.add(new StackNode(node, node.left != null ? Dir.Left : Dir.Right));
                node = node.left != null ? node.left : node.right;
            }
            return node;
        }
    }

    public boolean leafSimilar(TreeNode root1, TreeNode root2) {
        TreeIterator left = new TreeIterator(root1);
        TreeIterator right = new TreeIterator(root2);

        while (left.hasNext() && right.hasNext()) {
            if (left.getNext() != right.getNext()) {
                return false;
            }
        }

        return left.hasNext() == right.hasNext();
    }
}
