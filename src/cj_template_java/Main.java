import java.io.*;
import java.util.Arrays;
import java.util.Optional;
import java.util.Scanner;
import java.util.function.BiFunction;

public class Main implements BiFunction<Scanner, PrintWriter, Optional<String>> {
  @Override
  public Optional<String> apply(Scanner in, PrintWriter out) {
    final int n = in.nextInt();
    long sum = 0;
    for (int i=0; i < n; i++) {
        sum += in.nextLong();
    }

    return Optional.of(String.valueOf(sum));
  }

  public static void main(String[] args) {
    final Main main = new Main();

    try (
        final Scanner in = new Scanner(new BufferedReader(new InputStreamReader(System.in)));
        final PrintWriter out = new PrintWriter(new BufferedWriter(new OutputStreamWriter(System.out)));
    ) {
      final int testCount = in.nextInt();

      for (int index = 0; index < testCount; index++) {
        out.print("Case #" + (index+1)+": ");
        Optional<String> optResult = main.apply(in, out);
        if (optResult.isPresent()) {
          out.println(optResult.get());
        } else {
          out.println();
        }
      }
    }
  }

}
