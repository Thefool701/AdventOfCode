import java.util.Scanner;
import java.io.*;
import java.lang.Math;

public class Src {
    public static void main(String[] args) {
        int mass = 0;
        int totalFuel = 0;

        try {
            File infile = new File("input.in");
            Scanner in = new Scanner(infile);

            while (in.hasNext()) {
                int tempMass = 0;
                mass = in.nextInt();

                int temp = fuelCounter(mass);
                totalFuel += temp;

                while (temp > 0) {
                    int foo = fuelCounter(temp);
                    if (foo < 0) {
                        foo = 0;
                    }
                    tempMass += foo;
                    temp = foo;
                }
                totalFuel += tempMass;
            }

            System.out.println(totalFuel);

        } catch (Exception e) {
            System.out.println("Error Occured.");
            e.printStackTrace();
        }

    }

    public static int fuelCounter(int mass) {
        int result = (int) Math.floor(mass / 3) - 2;

        return result;
    }

}
