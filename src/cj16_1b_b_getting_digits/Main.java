import java.io.*;
import java.util.ArrayList;
import java.util.List;
import java.util.Optional;
import java.util.Scanner;
import java.util.function.BiFunction;

public class Main implements BiFunction<Scanner, PrintWriter, Optional<String>> {

  public static List<Tuple3<Integer, Character, String>> codes() {
    final ArrayList<Tuple3<Integer, Character, String>> codes = new ArrayList<>();

    codes.add(new Tuple3<>(0, 'Z', "ZERO"));
    codes.add(new Tuple3<>(2, 'W', "TWO"));
    codes.add(new Tuple3<>(4, 'U', "FOUR"));
    codes.add(new Tuple3<>(6, 'X', "SIX"));
    codes.add(new Tuple3<>(8, 'G', "EIGHT"));
    codes.add(new Tuple3<>(5, 'F', "FIVE"));
    codes.add(new Tuple3<>(7, 'V', "SEVEN"));
    codes.add(new Tuple3<>(3, 'R', "THREE"));
    codes.add(new Tuple3<>(9, 'I', "NINE"));
    codes.add(new Tuple3<>(1, 'N', "ONE"));

    return codes;
  }

  public Optional<String> apply(Scanner in, PrintWriter out) {
    final String input = in.next();
    final int[] chars = new int[26];
    for (char ch : input.toCharArray()) {
      chars[ch-'A'] ++;
    }

    final int[] digits = new int[10];
    final List<Tuple3<Integer, Character, String>> codes = codes();
    for (Tuple3<Integer, Character, String> code : codes) {
      for (int i = chars[code._1-'A']; i > 0; i--) {
        digits[code._0] ++;
        for (char ch : code._2.toCharArray()) {
          chars[ch - 'A'] --;
        }
      }
    }

    final StringBuilder result = new StringBuilder();
    for (int digit = 0; digit < digits.length; digit++) {
      int count = digits[digit];
      for (int i = 0; i < count; i++) {
        result.append(digit);
      }
    }

    return Optional.of(result.toString());
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

  public static class Tuple3<E0, E1, E2> {
    public final E0 _0;
    public final E1 _1;
    public final E2 _2;

    public Tuple3(E0 _0, E1 _1, E2 _2) {
      this._0 = _0;
      this._1 = _1;
      this._2 = _2;
    }

    @Override public boolean equals(Object o) {
      if (this == o) return true;
      if (o == null || getClass() != o.getClass()) return false;

      Tuple3<?, ?, ?> tuple3 = (Tuple3<?, ?, ?>) o;

      return _0.equals(tuple3._0) && _1.equals(tuple3._1) && _2.equals(tuple3._2);
    }

    @Override public int hashCode() {
      int result = _0.hashCode();
      result = 31 * result + _1.hashCode();
      result = 31 * result + _2.hashCode();
      return result;
    }

    @Override public String toString() {
      return "(" + _0 + "," + _1 + "," + _2  + ')';
    }
  }
}
