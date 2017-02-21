import java.util.Comparator;
import java.util.NoSuchElementException;
import java.util.Scanner;

public class Solution {
  public static class IntHeap {
    private final int leafValue;
    private final Comparator<Integer> order;

    private int[] elems;
    private int size = 0;

    protected IntHeap(int leafValue, Comparator<Integer> order, int capacity) {
      this.leafValue = leafValue;
      this.order = order;
      this.elems = new int[capacity];
    }

    public static IntHeap max(int capacity) {
      return new IntHeap(Integer.MIN_VALUE, Comparator.reverseOrder(), capacity);
    }

    public static IntHeap min(int capacity) {
      return new IntHeap(Integer.MAX_VALUE, Comparator.naturalOrder(), capacity);
    }

    public void add(int elem) {
      if (size == elems.length) {
        size *= 2;
        final int[] elemsPrev = elems;
        elems = new int[elemsPrev.length * 2];
        System.arraycopy(elemsPrev, 0, elems, 0, elemsPrev.length);
      }

      int pos = size;
      int parent = parentFor(pos);
      elems[size] = elem;
      size++;

      while (pos > 0 && order.compare(elems[parent], elems[pos]) > 0) {
        swap(pos, parent);
        pos = parent;
        parent = parentFor(pos);
      }
    }

    public int peek() {
      if (size == 0) {
        throw new NoSuchElementException("¯\\_(ツ)_/¯");
      }
      return elems[0];
    }

    public int removeTop() {
      if (size == 0) {
        throw new NoSuchElementException("¯\\_(ツ)_/¯");
      }
      final int res = elems[0];
      elems[0] = elems[size - 1];
      size--;

      int pos = 0;
      int posChildL = pos * 2 + 1;
      int posChildR = pos * 2 + 2;
      int elemChildL = posChildL < size ? elems[posChildL] : leafValue;
      int elemChildR = posChildR < size ? elems[posChildR] : leafValue;
      int childCompare = order.compare(elemChildL, elemChildR);
      int posTopChild = childCompare < 0 ? posChildL : posChildR;
      int elemTopChild = childCompare < 0 ? elemChildL : elemChildR;

      while (order.compare(elems[pos], elemTopChild) > 0) {
        swap(pos, posTopChild);

        pos = posTopChild;
        posChildL = pos * 2 + 1;
        posChildR = pos * 2 + 2;
        elemChildL = posChildL < size ? elems[posChildL] : leafValue;
        elemChildR = posChildR < size ? elems[posChildR] : leafValue;
        childCompare = order.compare(elemChildL, elemChildR);
        posTopChild = childCompare < 0 ? posChildL : posChildR;
        elemTopChild = childCompare < 0 ? elemChildL : elemChildR;
      }

      return res;
    }

    public int swapToWith(int value) {
      final int res = removeTop();
      add(value);
      return res;
    }

    public int getSize() {
      return size;
    }

    protected void swap(int posA, int posB) {
      int temp = elems[posA];
      elems[posA] = elems[posB];
      elems[posB] = temp;
    }

    protected int parentFor(int pos) {
      return (pos - 1) / 2;
    }

  }

  public static void main(String[] args) {
    final Scanner in = new Scanner(System.in);
    final int n = in.nextInt();

    final IntHeap lower = IntHeap.max(n);
    final IntHeap upper = IntHeap.min(n);

    for (int a_i = 0; a_i < n; a_i++) {
      int nextElem = in.nextInt();
      if (lower.getSize() > 0 && lower.peek() > nextElem) {
        nextElem = lower.swapToWith(nextElem);
      }
      if (upper.getSize() > 0 && upper.peek() < nextElem) {
        nextElem = upper.swapToWith(nextElem);
      }

      if (a_i % 2 == 0) {
        System.out.printf("%.1f\n", (0.0 + nextElem));
        lower.add(nextElem);
      } else {
        upper.add(nextElem);
        System.out.printf("%.1f\n", (0.0 + lower.peek() + upper.peek()) / 2);
      }
    }
  }
}
