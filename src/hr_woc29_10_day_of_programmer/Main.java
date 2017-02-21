import java.util.Scanner;

public class Main {

  public static void main(String[] args) {
    final Scanner in = new Scanner(System.in);
    final int y = in.nextInt();
    final String res;
    if (y > 1918) {
      if (y % 4 == 0 && y % 100 != 0 || y % 400 == 0) {
        res = "12.09";
      } else {
        res = "13.09";
      }
    } else if (y < 1918) {
      if (y % 4 == 0) {
        res = "12.09";
      } else {
        res = "13.09";
      }
    } else {
      res = "31.08";
    }
    System.out.println(res + "." + y);
  }
}
