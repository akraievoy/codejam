package org.akraievoy.codejam;

import org.testng.annotations.Test;

import static org.testng.Assert.*;

public class MainTest {

  @Test
  public void testSillyCase0() throws Exception {
    final Main main = new Main(">>> simple input here <<<");
    assertEquals(main.call(), ">>> expected result <<<");
  }

  @Test
  public void testSillyCase1() throws Exception {
    final Main main = new Main(">>> simple input here <<<");
    assertEquals(main.call(), ">>> expected result <<<");
  }

  @Test
  public void testSillyCase2() throws Exception {
    final Main main = new Main(">>> simple input here <<<");
    assertEquals(main.call(), ">>> expected result <<<");
  }

  @Test
  public void testSillyCase3() throws Exception {
    final Main main = new Main(">>> simple input here <<<");
    assertEquals(main.call(), ">>> expected result <<<");
  }

  @Test
  public void testSillyCase4() throws Exception {
    final Main main = new Main(">>> simple input here <<<");
    assertEquals(main.call(), ">>> expected result <<<");
  }

}