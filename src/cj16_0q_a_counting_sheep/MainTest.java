import org.testng.annotations.Test;

import static org.testng.Assert.*;

import java.io.PrintWriter;
import java.io.StringReader;
import java.io.StringWriter;
import java.util.Optional;
import java.util.Scanner;

public class MainTest {

  @Test
  public void testMergePresent(){
    testMergePresentCase(1, (byte) 9, false, true, false, false, false, false, false, false, false, false);
    testMergePresentCase(0, (byte) 9, true, false, false, false, false, false, false, false, false, false);
    testMergePresentCase(13579, (byte) 5, false, true, false, true, false, true, false, true, false, true);
    testMergePresentCase(20468, (byte) 5, true, false, true, false, true, false, true, false, true, false);
  }

  void testMergePresentCase(int number, byte digitsMissingExp, boolean zeroPresent, boolean onePresent, boolean twoPresent, boolean threePresent, boolean fourPresent, boolean fivePresent, boolean sixPresent, boolean sevenPresent, boolean eightPresent, boolean ninePresent) {
    final boolean[] intoDigits = new boolean[10];
    final byte[] digitsMissing = {10};
    new Main().mergePresent(number, intoDigits, digitsMissing);
    assertEquals(intoDigits, new boolean[] {zeroPresent, onePresent, twoPresent, threePresent, fourPresent, fivePresent, sixPresent, sevenPresent, eightPresent, ninePresent});
    assertEquals(digitsMissing, new byte[] {digitsMissingExp});
  }

  @Test
  public void testSillyCase0() throws Exception {
    assertEquals(run("0"), "INSOMNIA");
  }

  @Test
  public void testSillyCase1() throws Exception {
    assertEquals(run("1"), "10");
  }

  @Test
  public void testSillyCase2() throws Exception {
    assertEquals(run("2"), "90");
  }

  @Test
  public void testSillyCase3() throws Exception {
    assertEquals(run("11"), "110");
  }

  @Test
  public void testSillyCase4() throws Exception {
    assertEquals(run("1692"), "5076");
  }

  private static String run(String... lines) {
    final StringWriter inWriter = new StringWriter();
    final PrintWriter inPrintWriter = new PrintWriter(inWriter);
    for (String line : lines) {
      inPrintWriter.println(line);
    }
    inPrintWriter.close();

    final StringWriter stringWriter = new StringWriter();
    final PrintWriter printWriter = new PrintWriter(stringWriter);
    final Optional<String> maybeSingleLine = new Main().apply(
        new Scanner(new StringReader(inWriter.toString())),
        printWriter
    );
    maybeSingleLine.ifPresent(printWriter::print);
    printWriter.close();

    return stringWriter.toString();
  }

}