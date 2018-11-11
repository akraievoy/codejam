import java.io.BufferedWriter;
import java.io.OutputStreamWriter;
import java.io.PrintWriter;
import java.util.Arrays;
import java.util.Scanner;

public class Solution {

  public final static String VOWELS = "euioa";
  public final static String CONSONANTS = "qwrtpsdfghjklzxcvbnm";

  public static void main(String[] args) {
    final Scanner in = new Scanner(System.in);
    final int n = in.nextInt();

    long firstConsonantCount = CONSONANTS.length();
    for (int i = 1; i < n; i++) {
      firstConsonantCount *= (i % 2 == 0 ? CONSONANTS : VOWELS).length();
    }

    final PrintWriter out = new PrintWriter(new BufferedWriter(new OutputStreamWriter(System.out), 4 * 1024 * 1024));

    final char[] password = new char[n];
    Arrays.fill(password, 'y');
    final long[] memo = new long[n];
    Arrays.fill(memo, -1L);

    for (long i = 0; i < firstConsonantCount; i++) {
      long v = i;
      for (int pos = n - 1; pos >= 0 && memo[pos] != v; pos--) {
        final String chars = pos % 2 == 0 ? CONSONANTS : VOWELS;
        password[pos] =  chars.charAt((int) (v % chars.length()));
        memo[pos] = v;
        v = v / chars.length();
      }
      out.print(password);
      out.println();
    }

    long firstVowelCount = VOWELS.length();
    for (int i = 1; i < n; i++) {
      firstVowelCount *= (i % 2 == 0 ? VOWELS : CONSONANTS).length();
    }

    Arrays.fill(password, 'y');
    Arrays.fill(memo, -1L);
    for (long i = 0; i < firstVowelCount; i++) {
      long v = i;
      for (int pos = n - 1; pos >= 0 && memo[pos] != v; pos--) {
        final String chars = pos % 2 == 0 ? VOWELS : CONSONANTS;
        password[pos] = chars.charAt((int) (v % chars.length()));
        memo[pos] = v;
        v = v / chars.length();
      }
      out.print(password);
      out.println();
    }

    out.flush();
    out.close();
  }
}
