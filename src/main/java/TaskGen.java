import java.io.BufferedWriter;
import java.io.OutputStreamWriter;
import java.io.PrintWriter;
import java.util.Random;

public class TaskGen {
  public static void main(String[] args) {
    final long seed;
    if (args.length > 0 && !"ctm".equals(args[0])) {
      seed = Long.parseLong(args[0]);
    } else {
      seed = System.currentTimeMillis();
    }
    System.err.println("seed = " + seed);
    final Random random = new Random(seed);

    final int n = args.length > 1 ? Integer.parseInt(args[1]) : 1 + random.nextInt(40000);
    final int queries = args.length > 2 ? Integer.parseInt(args[2]) : 1 + random.nextInt(40000);

    final PrintWriter out = new PrintWriter(new BufferedWriter(new OutputStreamWriter(System.out), 4 * 1024 * 1024));
    out.println(n + " " + queries);

    for (int i = 0; i < n; i++) {
      if (i > 0) {
        out.print(" ");
      }
      out.print(random.nextInt(40001));
    }
    out.println();

    for (int q = 0; q < queries; q++) {
      final int modulo = 1 + random.nextInt(40000);
      out.println(
        random.nextInt(n/2+1) + " " +
          (n/2 + random.nextInt((n + 1)/2)) + " "
          + modulo + " " + random.nextInt(modulo)
      );
    }

    out.flush();
    out.close();
  }
}
