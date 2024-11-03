package com.devandrew.helpers;

import com.google.gson.Gson;
import com.google.gson.reflect.TypeToken;

import java.io.File;
import java.io.FileNotFoundException;
import java.io.FileWriter;
import java.io.IOException;
import java.lang.reflect.Type;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.nio.file.StandardOpenOption;
import java.util.Scanner;
import java.util.function.Consumer;
import java.util.function.Function;

public class Parser {
    private final String currentProgram;
    private final Gson gson = new Gson();

    public Parser(String programKey) {
        this.currentProgram = programKey;
    }

    public int[] parseArray(String json) {
        Type type = new TypeToken<int[]>() {
        }.getType();

        return gson.fromJson(json, type);
    }

    public Integer[] parseNullableArray(String json) {
        Type type = new TypeToken<Integer[]>() {
        }.getType();

        return gson.fromJson(json, type);
    }

    public String[] parseStringArray(String json) {
        Type type = new TypeToken<String[]>() {
        }.getType();

        return gson.fromJson(json, type);
    }

    public char[][] parseCharArrays(String json) {
        Type type = new TypeToken<char[][]>() {
        }.getType();

        return gson.fromJson(json, type);
    }

    public int[][] parseArrays(String json) {
        Type type = new TypeToken<int[][]>() {
        }.getType();

        return gson.fromJson(json, type);
    }

    public void withInput(Consumer<Scanner> consumer) {
        withInput(consumer, String.format("./inputs/%s.txt", currentProgram));
    }

    public void withInput(Consumer<Scanner> consumer, String path) {
        File file = new File(path);
        try (Scanner scanner = new Scanner(file)) {
            consumer.accept(scanner);
        } catch (FileNotFoundException ex) {
            System.out.printf("Error occurred opening file: %s%n", ex.getMessage());
        }
    }

    public <T> void withInputAndOutputProblemAdjusted(Function<Scanner, T> operation) {
        withInputAndOutput(operation,
                String.format("./src/main/java/com/devandrew/_%s/input.txt", currentProgram),
                String.format("./src/main/java/com/devandrew/_%s/output.txt", currentProgram)
        );
    }

    public <T> void withInputAndOutput(Function<Scanner, T> operation) {
        withInputAndOutput(operation,
                String.format("./inputs/%s.txt", currentProgram),
                String.format("./outputs/%s.txt", currentProgram)
        );
    }

    public <T> void withInputAndOutput(Function<Scanner, T> operation, String inputPath, String outputPath) {
        this.createOutputFile(outputPath);

        withInput(scanner -> {
            while (scanner.hasNext()) {
                this.appendToOutput(operation.apply(scanner), outputPath);
            }
        }, inputPath);
    }

    public void createOutputFile(String path) {
        try {
            File output = new File(path);
            if (output.exists()) {
                output.delete();
            }

            output.createNewFile();
        } catch (IOException ex) {
            System.out.printf("Error occurred creating file: %s%n", ex.getMessage());
        }
    }

    public void createOutputFile() {
        this.createOutputFile(String.format("./outputs/%s.txt", currentProgram));
    }

    public <T> void appendToOutput(T sol) {
        this.appendToOutput(sol, String.format("./outputs/%s.txt", currentProgram));
    }

    public <T> void appendToOutput(T sol, String path) {
        try {
            String output = String.format("%s\n", gson.toJson(sol));

            Files.write(Paths.get(path), output.getBytes(), StandardOpenOption.APPEND);
        } catch (IOException ex) {
            System.out.printf("Error occurred writing to file: %s%n", ex.getMessage());
        }
    }

    public <T> void outputSolution(T sol) {
        try {
            File output = new File(String.format("./outputs/%s.txt", currentProgram));
            output.createNewFile();
            FileWriter writer = new FileWriter(output);
            writer.append(gson.toJson(sol));
            writer.close();
        } catch (IOException ex) {
            System.out.printf("Error occurred writing to file: %s%n", ex.getMessage());
        }
    }
}
