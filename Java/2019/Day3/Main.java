import java.util.*;
import java.io.*;
import java.lang.Math;
// The Principal Question
// What is the Manhattan distance from the central port to the closest intersection?

public class Main {
    public static void main(String[] args) {
        File infile = new File("input.in");

        // Taxicab distance between two points
        try {
            Scanner in = new Scanner(infile);

            String wire1 = in.nextLine();
            String wire2 = in.nextLine();

            String[] temp1 = wire1.split(",");
            String[] temp2 = wire2.split(",");

            int xN = temp1.length;
            int yN = temp2.length;

            int[] x = new int[xN];
            int[] y = new int[yN];

            for (int i = 0; i < xN; i++) {
                temp1[i] = temp1[i].replaceAll("[\\D]", "");
                temp2[i] = temp2[i].replaceAll("[\\D]", "");
                x[i] = Integer.parseInt(temp1[i]);
                y[i] = Integer.parseInt(temp2[i]);
            }

            // Manhattan Distance Formula:
            // \x1 - x2\ + \y1 - y2\

            System.out.println(xN);
            System.out.println(yN);
            in.close();
        } catch (Exception e) {
            System.out.println("Error Occured. Try Again.");
            e.printStackTrace();
        }
    }

    public static void printArr(int[] arr) {
        for (int i = 0; i < arr.length; i++) {
            System.out.printf("%d ", arr[i]);
        }
    }

}
