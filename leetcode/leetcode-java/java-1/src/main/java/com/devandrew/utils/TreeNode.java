package com.devandrew.utils;

public class TreeNode {
    public int val;
    public TreeNode left;
    public TreeNode right;
    public TreeNode() {}
    public TreeNode(int val) { this.val = val; }
    public TreeNode(int val, TreeNode left, TreeNode right) {
        this.val = val;
        this.left = left;
        this.right = right;
    }

    public static TreeNode fromArray(Integer[] arr, int idx) {
        if (idx < arr.length) {
            TreeNode left = fromArray(arr, idx * 2 + 1);
            TreeNode right = fromArray(arr, idx * 2 + 2);
            return arr[idx] == null ? null : new TreeNode(arr[idx], left, right);
        }
        return null;
    }
}
