package com.devandrew._2281;

public class _2281 {

    public static long MOD = 1000000007;

    class MinSegmentTree {
        private int[] tree;
        private int n;

        public MinSegmentTree(int[] arr, int n) {
            int height = (int)Math.ceil(Math.log(n) / Math.log(2));
            int size = 2 * (int)Math.pow(2, height) - 1;
            this.tree = new int[size];
            this.n = n;
            this.constructTree(arr, 0, n - 1, 0);
        }

        public int queryMin(int l, int r) {
            return queryMin(l, r, 0, n - 1, 0);
        }

        private int queryMin(int l, int r, int tl, int tr, int i) {
            if (tl >= l && tr <= r) {
                return tree[i];
            }

            if (tr < l || tl > r) {
                return Integer.MAX_VALUE;
            }

            int m = (tl + tr) / 2;
            return Math.min(queryMin(l, r, tl, m, 2 * i + 1),
                    queryMin(l, r, m + 1, tr, 2 * i + 2));
        }

        private int constructTree(int[] arr, int l, int r, int i) {
            if (l == r) {
                tree[i] = arr[l];
                return tree[i];
            }

            int m = (l + r) / 2;
            tree[i] = Math.min(constructTree(arr, l, m, 2 * i + 1), constructTree(arr, m + 1, r, 2 * i + 2));
            return tree[i];
        }
    }

    class SumOfSubarraysCalculator {
        private int[] strength;

        private long[] prefix;

        public SumOfSubarraysCalculator(int[] strength) {
            this.strength = strength;
            this.prefix = new long[strength.length + 1];

            for (int i = 0; i < strength.length; i++) {
                this.prefix[i + 1] = strength[i] + this.prefix[i];
            }
        }

        public long sum(int l, int r) {
            return this.prefix[r + 1] - this.prefix[l];
        }

        public long calculate(int l, int r) {
            if (l == r) {
                return strength[l];
            }

            int m = (l + r) / 2;

            long leftResult = calculate(l, m) % MOD;
            long rightResult = calculate(m + 1, r) % MOD;

            int leftLen = m - l + 1;
            int rightLen = r - m;

            long sum = 0;
            for (int i = l, curr = rightLen; i <= m; i++, curr += rightLen) {
                sum += ((long) curr * strength[i]) % MOD;
            }

            for (int i = r, curr = leftLen; i > m; i--, curr += leftLen) {
                sum += ((long) curr * strength[i]) % MOD;
            }

            return (leftResult + rightResult + sum) % MOD;
        }
    }

    public int totalStrength(int[] strength) {
        MinSegmentTree tree = new MinSegmentTree(strength, strength.length);
        SumOfSubarraysCalculator calculator = new SumOfSubarraysCalculator(strength);

        long ans = 0;
        for (int i = 0; i < strength.length; i++) {
            for (int j = i; j < strength.length; j++) {
                int min = tree.queryMin(i, j);
                long sum = calculator.sum(i, j);

                ans = (ans + (min * sum)) % MOD;
            }
        }

        return (int)ans;
    }
}
