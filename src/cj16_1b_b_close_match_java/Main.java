import java.io.*;
import java.util.*;
import java.util.function.BiFunction;

public class Main implements BiFunction<Scanner, PrintWriter, Optional<String>> {
  final char[] min(char[] arr, int from) {
    final char[] res = arr.clone();
    for (int i = from; i < res.length; i++) {
      res[i] = res[i] == '?' ? '0' : res[i];
    }
    return res;
  }

  final char[] max(char[] arr, int from) {
    final char[] res = arr.clone();
    for (int i = from; i < res.length; i++) {
      res[i] = res[i] == '?' ? '9' : res[i];
    }
    return res;
  }

  @Override
  public Optional<String> apply(Scanner in, PrintWriter out) {
    final String[] strings = in.nextLine().trim().split(" +");

    final List<Tuple3<String, String, Long>> options = new ArrayList<>();

    final char[][] arr = {strings[0].toCharArray(), strings[1].toCharArray()};
    for (int pos = 0; pos < arr[0].length; pos++) {
      final char[] arr0 = arr[0].clone();
      final char[] arr1 = arr[1].clone();

      boolean is0eq1 = true;
      boolean is0gt1 = false;
      for (int i = 0; i < pos; i++) {
        if (arr0[i] == arr1[i]) {
          if (arr0[i] == '?') {
            arr0[i] = '0';
            arr1[i] = '0';
          }
        } else if (arr0[i] == '?') {
          arr0[i] = arr1[i];
        } else if (arr1[i] == '?') {
          arr1[i] = arr0[i];
        } else if (is0eq1) {
          is0eq1 = false;
          is0gt1 = arr0[i] > arr1[i];
        }
      }

      if (is0eq1) {
        if (arr0[pos] == arr1[pos]) {
          if (arr0[pos] == '?') {
              arr0[pos] = '0';
              arr1[pos] = '0';
              addOption(options, min(arr0, pos + 1), min(arr1, pos + 1));
              arr0[pos] = '1';
              arr1[pos] = '0';
              addOption(options, min(arr0, pos + 1), max(arr1, pos + 1));
              arr0[pos] = '0';
              arr1[pos] = '1';
              addOption(options, max(arr0, pos + 1), min(arr1, pos + 1));
          }
        } else if (arr0[pos] == '?') {
          arr0[pos] = arr1[pos];
          addOption(options, min(arr0, pos + 1), min(arr1, pos + 1));
          if (arr1[pos] < '9') {
            arr0[pos] = (char) (arr1[pos] + 1);
            addOption(options, min(arr0, pos + 1), max(arr1, pos + 1));
          }
          if (arr1[pos] > '0') {
            arr0[pos] = (char) (arr1[pos] - 1);
            addOption(options, max(arr0, pos + 1), min(arr1, pos + 1));
          }
        } else if (arr1[pos] == '?') {
          arr1[pos] = arr0[pos];
          addOption(options, min(arr0, pos + 1), min(arr1, pos + 1));
          if (arr0[pos] < '9') {
            arr1[pos] = (char) (arr0[pos] + 1);
            addOption(options, max(arr0, pos + 1), min(arr1, pos + 1));
          }
          if (arr0[pos] > '0') {
            arr1[pos] = (char) (arr0[pos] - 1);
            addOption(options, min(arr0, pos + 1), max(arr1, pos + 1));
          }
        } else {
          addOption(options, min(arr0, pos), min(arr1, pos));
        }
      } else if (is0gt1) {
        addOption(options, min(arr0, pos), max(arr1, pos));
      } else {
        addOption(options, max(arr0, pos), min(arr1, pos));
      }
    }

    options.sort((o1, o2) -> {
      final int diffCompared = Long.compare(o1._2, o2._2);
      if (diffCompared != 0) {
        return diffCompared;
      }
      final int cCompared = o1._0.compareTo(o2._0);
      if (cCompared != 0) {
        return cCompared;
      }
      return o1._1.compareTo(o2._1);
    });

    final Tuple3<String, String, Long> res = options.get(0);
    return Optional.of(res._0 + " " + res._1);
  }

  void addOption(List<Tuple3<String, String, Long>> options, char[] cArg, char[] jArg) {
    final String c = String.valueOf(cArg);
    final String j = String.valueOf(jArg);
    options.add(new Tuple3<>(c, j, Math.abs(Long.parseLong(c) - Long.parseLong(j))));
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
