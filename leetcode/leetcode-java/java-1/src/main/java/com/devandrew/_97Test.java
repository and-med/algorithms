package com.devandrew;

import com.devandrew.helpers.Parser;

public class _97Test {
    public static void run() {
        Parser parser = new Parser("97");

        parser.withInput((scanner) -> {
            parser.createOutputFile();

            while (scanner.hasNext()) {
                String s1 = scanner.nextLine();
                String s2 = scanner.nextLine();
                String s3 = scanner.nextLine();

                boolean res = new _97InterleavingStrings().isInterleave(s1, s2, s3);
                parser.appendToOutput(res);
            }
        });
    }
}
