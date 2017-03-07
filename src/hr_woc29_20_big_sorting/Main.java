import java.util.Arrays;
import java.util.Comparator;
import java.util.Scanner;

public class Main {
  public static final Comparator<String> elemLenOrder =
    (e1, e2) -> Integer.compare(e1.length(), e2.length());

  private static String commonPrefix(String s1, String s2) {
    int res = 0;
    int lim = Math.min(s1.length(), s2.length());
    while (res < lim && s1.charAt(res) == s2.charAt(res)) {
      res++;
    }
    return s1.substring(0, res);
  }

  private static void sortSameLen(
    String[] unsorted,
    int len, int prefixLen,
    int rangeMin, int rangeMax
  ) {
    if (rangeMin + 1 >= rangeMax || len == prefixLen) {
      return;
    }
    String nextPrefix =
      commonPrefix(
        unsorted[rangeMin].substring(prefixLen),
        unsorted[rangeMin+1].substring(prefixLen)
      );
    for (int prefixI = rangeMin + 2; prefixI < rangeMax; prefixI++) {
      nextPrefix = commonPrefix(nextPrefix, unsorted[prefixI].substring(prefixLen));
    }

    final int bucketingPos = prefixLen + nextPrefix.length();

    if (bucketingPos >= len) {
      return;
    }

    Arrays.sort(
      unsorted,
      rangeMin,
      rangeMax,
      (o1, o2) -> Character.compare(o1.charAt(bucketingPos), o2.charAt(bucketingPos))
    );

    char bucketChar = unsorted[rangeMin].charAt(bucketingPos);
    int nextRangeMin = rangeMin;
    for (int i = rangeMin + 1; i < rangeMax; i++) {
      if (unsorted[i].charAt(bucketingPos) != bucketChar) {
        sortSameLen(unsorted, len, bucketingPos + 1, nextRangeMin, i);
        nextRangeMin = i;
        bucketChar = unsorted[i].charAt(bucketingPos);
      }
    }
    sortSameLen(unsorted, len, bucketingPos + 1, nextRangeMin, rangeMax);
  }

  public static void main(String[] args) {
    final Scanner in = new Scanner(System.in);
    final int n = in.nextInt();
    final String[] unsorted = new String[n];
    for (int unsorted_i = 0; unsorted_i < n; unsorted_i++) {
      unsorted[unsorted_i] = in.next();
    }

    mainer(unsorted);

    for (int i = 0; i < n; i++) {
      System.out.println(unsorted[i]);
    }
  }

  public static void mainT(String[] args) {
    System.out.println(
      Arrays.toString(
        mainer(
          new String[] {
            "12312321231",
            "12312312312321231",
            "12312321232",
            "12312312312321232",
            "12312321233",
            "12312312312321233",
            "12312321232",
            "12312312312321232",
            "12312321231",
            "12312312312321231",
            "12312921231",
            "12312912312321231",
            "12312921232",
            "12312912312321232",
            "12312921233",
            "12312912312321233",
            "12312921232",
            "12312912312321232",
            "12312921231",
            "12312912312321231"
          }
        )
      )
    );
  }

  private static String[] mainer(String[] unsorted) {
    Arrays.sort(unsorted, elemLenOrder);

    int rangeMin = 0;
    int len = unsorted[0].length();
    for (int i = 0; i < unsorted.length; i++) {
      if (unsorted[i].length() != len) {
        sortSameLen(unsorted, len, 0, rangeMin, i);
        len = unsorted[i].length();
        rangeMin = i;
      }
    }

    sortSameLen(unsorted, len, 0, rangeMin, unsorted.length);

    return unsorted;
  }
}
