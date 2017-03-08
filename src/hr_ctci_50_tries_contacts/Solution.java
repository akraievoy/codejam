import java.io.FileInputStream;
import java.io.FileOutputStream;
import java.io.IOException;
import java.io.PrintStream;
import java.util.ArrayList;
import java.util.Collections;
import java.util.Scanner;

public class Solution {
  public static class Trie {
    private int weight = 0;
    private int nestedWeight = 0;
    private ArrayList<String> paths = new ArrayList<>(3);
    private ArrayList<Trie> tries = new ArrayList<>(3);

    public int getWeight() {
      return weight;
    }

    public int getNestedWeight() {
      return nestedWeight;
    }

    private String internIfShort(String z) {
      if  (z.length() > 3) {
        return z;
      } else {
        return z.intern();
      }
    }

    public Trie add(String path) {
      if (path.isEmpty()) {
        weight += 1;
        return this;
      }

      nestedWeight += 1;

      final int pathSearch = Collections.binarySearch(paths, path.substring(0, 1));
      if (pathSearch >= 0 && paths.get(pathSearch).equals(path)) {
        tries.get(pathSearch).add("");
        return this;
      }

      final int pathInsertion = pathSearch < 0 ? -pathSearch - 1 : pathSearch;
      if (pathInsertion >= paths.size() || paths.get(pathInsertion).charAt(0) != path.charAt(0)) {
        paths.add(pathInsertion, path);
        tries.add(pathInsertion, new Trie().add(""));
        return this;
      }

      final String pathSplit = paths.get(pathInsertion);
      final int prefixLenMax = Math.min(pathSplit.length(), path.length());
      int prefixLen = 0;
      while (prefixLen < prefixLenMax && pathSplit.charAt(prefixLen) == path.charAt(prefixLen)) {
        prefixLen += 1;
      }

      final String pathSplitPrefix = internIfShort(path.substring(0, prefixLen));
      final String pathSplitTail = internIfShort(pathSplit.substring(prefixLen));
      final String pathTail = internIfShort(path.substring(prefixLen));

      if (pathSplitTail.length() > 0) {
        final Trie wrappedTrie = tries.get(pathInsertion);
        final Trie wrapperTrie = new Trie();
        wrapperTrie.paths.add(pathSplitTail);
        wrapperTrie.tries.add(wrappedTrie);
        wrapperTrie.nestedWeight = wrappedTrie.nestedWeight + wrappedTrie.weight;

        paths.set(pathInsertion, pathSplitPrefix);

        tries.set(pathInsertion, wrapperTrie.add(pathTail));
      } else {
        tries.set(pathInsertion, tries.get(pathInsertion).add(pathTail));
      }

      return this;
    }

    public Trie view(String path) {
      if (path.isEmpty()) {
        return this;
      }

      final int pathSearch = Collections.binarySearch(paths, path.substring(0,1));
      if (pathSearch >= 0) {
        return tries.get(pathSearch).view(path.substring(1));
      }

      final int pathInsertion = -pathSearch - 1;
      if (pathInsertion == paths.size() || paths.get(pathInsertion).charAt(0) != path.charAt(0)) {
        return null;
      }

      final String pathSplit = paths.get(pathInsertion);
      final int prefixLenMax = Math.min(pathSplit.length(), path.length());
      int prefixLen = 0;
      while (prefixLen < prefixLenMax && pathSplit.charAt(prefixLen) == path.charAt(prefixLen)) {
        prefixLen += 1;
      }

      if (prefixLen == prefixLenMax) {
        return tries.get(pathInsertion).view(path.substring(prefixLen));
      }
      return null;
    }

    public int weightForPrefix(String prefix) {
      final Trie viewByPartial = view(prefix);
      if (viewByPartial == null) {
        return 0;
      }

      return
        viewByPartial.getNestedWeight() +
          viewByPartial.getWeight();
    }
  }

  public static void mainTT(String[] args) {
    final Trie root = new Trie();
    root.add("abababab");
    root.add("ababab");
    root.add("abab");
    root.add("ab");
    root.add("a");
    root.add("aba");
    root.add("ababa");
    root.add("abababa");
    root.add("ababababa");
    root.add("cdcdcdcd");
    root.add("cdcdcd");
    root.add("cdcd");
    root.add("cd");
    root.add("c");
    root.add("cdc");
    root.add("cdcdc");
    root.add("cdcdcdc");
    root.add("cdcdcdcdc");
    root.add("abababab");
    root.add("ababab");
    root.add("abab");
    root.add("ab");
    root.add("a");
    root.add("aba");
    root.add("ababa");
    root.add("abababa");
    root.add("ababababa");
    root.add("cdcdcdcd");
    root.add("cdcdcd");
    root.add("cdcd");
    root.add("cd");
    root.add("c");
    root.add("cdc");
    root.add("cdcdc");
    root.add("cdcdcdc");
    root.add("cdcdcdcdc");
    System.out.println(root.weightForPrefix("a"));
    System.out.println(root.weightForPrefix("ab"));
    System.out.println(root.weightForPrefix("aba"));
    System.out.println(root.weightForPrefix("abab"));
    System.out.println(root.weightForPrefix("ababa"));
    System.out.println(root.weightForPrefix("ababab"));
    System.out.println(root.weightForPrefix("abababa"));
    System.out.println(root.weightForPrefix("abababab"));
    System.out.println(root.weightForPrefix("ababababa"));
    System.out.println(root.weightForPrefix("ababababac"));
    System.out.println(root.weightForPrefix("ababac"));
    System.out.println(root.weightForPrefix("c"));
    System.out.println(root.weightForPrefix("cd"));
    System.out.println(root.weightForPrefix("cdc"));
    System.out.println(root.weightForPrefix("cdcd"));
    System.out.println(root.weightForPrefix("cdcdc"));
    System.out.println(root.weightForPrefix("cdcdcd"));
    System.out.println(root.weightForPrefix("cdcdcdc"));
    System.out.println(root.weightForPrefix("cdcdcdcd"));
    System.out.println(root.weightForPrefix("cdcdcdcdc"));
  }

  public static void mainT(String[] args) throws IOException {
    try (
      FileInputStream fileIn = new FileInputStream("inout/input02.txt");
      PrintStream fileOut = new PrintStream(new FileOutputStream("inout/output02.actual.txt"))
    ) {
      System.setIn(fileIn);
      System.setOut(fileOut);
      main(args);
    }
  }

  public static void main(String[] args) {
    final Trie root = new Trie();
    final Scanner in = new Scanner(System.in);
    final int n = in.nextInt();
    for (int a0 = 0; a0 < n; a0++) {
      final String op = in.next();
      final String contact = in.next();

      if ("add".equals(op)) {
        root.add(contact);
      } else {
        System.out.println(root.weightForPrefix(contact));
      }

    }
  }
}
