import java.io.*;
import java.util.ArrayList;
import java.util.List;
import java.util.Optional;
import java.util.Scanner;
import java.util.function.BiFunction;

public class Main implements BiFunction<Scanner, PrintWriter, Optional<String>> {

  public Optional<String> apply(Scanner in, PrintWriter out) {
    return Optional.of(sheepSleepNumber(in.nextLong()));
  }

  String sheepSleepNumber(long seed) {
    if (seed == 0) {
      return "INSOMNIA";
    }

    final byte[] digitsMissing = new byte[]{10};
    final boolean[] digitSeen = new boolean[10];

    mergePresent(seed, digitSeen, digitsMissing);

    long current = seed;
    while (digitsMissing[0] > 0) {
      current = current + seed;
      mergePresent(current, digitSeen, digitsMissing);
    }

    return String.valueOf(current);
  }

  void mergePresent(long num, boolean[] digitsSeen, byte[] digitsMissing) {
    boolean singleZeroCase = true;
    while (singleZeroCase || num > 0) {
      singleZeroCase = false;
      final int digit = (int) num % 10;
      if (!digitsSeen[digit]) {
        digitsSeen[digit] = true;
        digitsMissing[0]--;
      }
      num = num / 10;
    }
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
