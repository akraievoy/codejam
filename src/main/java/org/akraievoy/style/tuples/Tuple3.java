package org.akraievoy.style.tuples;

public class Tuple3<E0, E1, E2> {
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
