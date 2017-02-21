import java.util.Scanner;
import java.util.Stack;

public class Solution {
  public static class MyQueue<E> {
    private final Stack<E> heads = new Stack<E>();
    private final Stack<E> tails = new Stack<E>();

    private void fetchHeads() {
      if (heads.isEmpty() && !tails.isEmpty()) {
        while (!tails.isEmpty()) {
          heads.push(tails.pop());
        }
      }
    }

    public MyQueue() {

    }

    public void enqueue(E elem) {
      tails.push(elem);
    }

    public E dequeue() {
      fetchHeads();
      return heads.pop();
    }

    public E peek() {
      fetchHeads();
      return heads.peek();
    }
  }

  public static void main(String[] args) {
    MyQueue<Integer> queue = new MyQueue<Integer>();

    Scanner scan = new Scanner(System.in);
    int n = scan.nextInt();

    for (int i = 0; i < n; i++) {
      int operation = scan.nextInt();
      if (operation == 1) { // enqueue
        queue.enqueue(scan.nextInt());
      } else if (operation == 2) { // dequeue
        queue.dequeue();
      } else if (operation == 3) { // print/peek
        System.out.println(queue.peek());
      }
    }
    scan.close();
  }
}
