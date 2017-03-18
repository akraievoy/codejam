import java.io.BufferedWriter;
import java.io.OutputStreamWriter;
import java.io.PrintWriter;
import java.util.Scanner;

public class SolutionBrute {
  public static void main(String[] args) {
    final Scanner in = new Scanner(System.in);
    final int n = in.nextInt();
    final int queries = in.nextInt();

    final int[] a = new int[n];
    for (int i = 0; i < n; i++) {
      a[i] = in.nextInt();
    }

    final PrintWriter out = new PrintWriter(new BufferedWriter(new OutputStreamWriter(System.out), 4 * 1024 * 1024));
    for (int q = 0; q < queries; q++) {
      final int left = in.nextInt();
      final int right = in.nextInt();
      final int modulo = in.nextInt();
      final int remainder = in.nextInt();

      int count = 0;
      for (int i = left; i <= right; i++) {
        if (a[i] % modulo == remainder) {
          count++;
        }
      }
      out.println(count);
    }

    out.flush();
    out.close();
  }
}
