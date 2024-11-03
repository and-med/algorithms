package com.devandrew._124;

import com.devandrew.utils.TreeNode;

public class _124_OofN {
    private int max;

    public int maxPathSum(TreeNode root) {
        this.max = Integer.MIN_VALUE;
        maxPathEndingAt(root);
        return this.max;
    }

    public int maxPathEndingAt(TreeNode root) {
        if (root == null) {
            return 0;
        }

        int leftMax = Math.max(maxPathEndingAt(root.left), 0);
        int rightMax = Math.max(maxPathEndingAt(root.right), 0);

        int ans = Math.max(root.val + leftMax, root.val + rightMax);
        this.max = Math.max(this.max, Math.max(ans, root.val + leftMax + rightMax));

        return ans;
    }
}
