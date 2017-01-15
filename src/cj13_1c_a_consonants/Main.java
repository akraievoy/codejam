import java.io.*;
import java.util.Arrays;
import java.util.Optional;
import java.util.Scanner;
import java.util.function.BiFunction;

public class Main implements BiFunction<Scanner, PrintWriter, Optional<String>> {
  private final static char[] vowels = "aeiou".toCharArray();

  @Override
  public Optional<String> apply(Scanner in, PrintWriter out) {
    final String name = in.next();
    final int L = in.nextInt();
    final int n = name.length();

    System.out.println("name = '" + name + "'");
    System.out.println("L = " + L);

    long total = 0;
    int prevStart = -1;
    int sequenceCount = 0;
    for (int pos = 0; pos < L; pos++) {
      if (!isVowel(name.charAt(pos))) {
        sequenceCount++;
      }
      if (pos >= n && !isVowel(name.charAt(pos - n))) {
        sequenceCount--;
      }
      if (sequenceCount == n) {
        final int start = pos - n + 1;
        total += ((long) (L-pos)) * (start - prevStart);
        prevStart = start;
      }
    }

    return Optional.of(String.valueOf(total));
  }

  private static boolean isVowel(char ch) {
    return Arrays.binarySearch(vowels, ch) >= 0;
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
