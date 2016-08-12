//为什么java可以处理大数


import java.math.BigInteger;
import java.util.Scanner;
 
public class Main {
    public static void main(String[] args){
        Scanner scan = new Scanner(System.in);
        int T, i;
        T = scan.nextInt();
        for (i = 1; i <= T; ++i){
            BigInteger a = scan.nextBigInteger();
            BigInteger b = scan.nextBigInteger();
            BigInteger c = a.add(b);
            System.out.println("Case " + i + ":");
            System.out.println(a + " + " + b + " = " + c);
            if (i != T) System.out.println();
 
        }
    }
}
 
