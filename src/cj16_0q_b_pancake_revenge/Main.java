import java.io.*;
import java.util.ArrayList;
import java.util.List;
import java.util.Optional;
import java.util.Scanner;
import java.util.function.BiFunction;

public class Main implements BiFunction<Scanner, PrintWriter, Optional<String>> {

  public Optional<String> apply(Scanner in, PrintWriter out) {
    final String noAdjancedDupes = in.next().replaceAll("\\+\\++", "+").replaceAll("\\-\\-+", "-");
    final String noRedundantTrailingHappies = noAdjancedDupes.replaceAll("\\++$", "");
    return Optional.of(String.valueOf(noRedundantTrailingHappies.length()));
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
