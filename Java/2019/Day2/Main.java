import java.util.*;
import java.io.*;
import java.security.*;

public class Main {
    private static final SecureRandom randomNumbers = new SecureRandom();

    public static void main(String[] args) {
        File infile = new File("input.in");
        String line = "";
        int opcode = 0;
        int pos1 = 0;
        int pos2 = 0;
        int inp1 = 0;
        int inp2 = 0;
        int result = 0;
        int des = 0;
        int noun = 0;
        int verb = 0;

        try {
            Scanner in = new Scanner(infile);

            line = in.nextLine();

            // Turn string to string array
            String[] tempLine = line.split(",");
            int n = tempLine.length;
            int[] arr = new int[n];

            // Loop and eheck each and every value if it is a int, and store
            // into an int array
            for (int i = 0; i < n; i++) {
                arr[i] = Integer.parseInt(tempLine[i]);
            }
            // READ THE FULL INSTRUCTIONS NEXT TIME jeez...
            // what the fuck
            noun = 54;
            verb = 85;

            arr[1] = noun;
            arr[2] = verb;

            // Begin looping the int array
            for (int i = 0; opcode != 99; i += 4) {
                opcode = arr[i];
                // Check opcode first
                if (opcode == 1) {
                    // Check positions first
                    pos1 = arr[i + 1];
                    pos2 = arr[i + 2];
                    des = arr[i + 3];

                    // Then get input from positions
                    inp1 = arr[pos1];
                    inp2 = arr[pos2];

                    // Add the two inputs
                    result = add(inp1, inp2);

                    // Then store into the arr des
                    arr[des] = result;
                }

                if (opcode == 2) {
                    // Check positions first
                    pos1 = arr[i + 1];
                    pos2 = arr[i + 2];
                    des = arr[i + 3];

                    // Then get input from positions
                    inp1 = arr[pos1];
                    inp2 = arr[pos2];

                    // Add the two inputs
                    result = multiply(inp1, inp2);

                    // Then store into the arr des
                    arr[des] = result;
                }
            }
            int ans = 100 * noun + verb;
            System.out.printf("Noun: %d%n", noun);
            System.out.printf("Verb: %d%n", verb);
            System.out.printf("Answer: %d%n", ans);
            System.out.printf("Input Result: %d%n", arr[0]);
            // Resetting values
            pos1 = 0;
            pos2 = 0;
            inp1 = 0;
            inp2 = 0;
            result = 0;
            des = 0;
            noun = 0;
            verb = 0;
            in.close();
        } catch (Exception e) {
            System.out.println("Error Occured.");
            e.printStackTrace();
        }

    }

    public static int add(int x, int y) {
        return x + y;
    }

    public static int multiply(int x, int y) {
        return x * y;
    }

}
