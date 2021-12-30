from typing import Sequence, TypeVar, Generic

T = TypeVar("T")

class _Node(Generic[T]):
  def __init__(self, value:T) -> None:
    self.__value = value
    self.__parent: _Node[T] | None = None
    self.__height = 1

  def is_root(self) -> bool:
    return self.__parent is None

  def union(self, target: '_Node[T]') -> None:
    self = self.root
    target = target.root
    if self.value == target.value:
      return

    if target.height > self.height:
      tmp = self
      self = target
      target = tmp

    target.__parent = self
    if target.height == self.height:
      self.__height += 1

  @property
  def height(self) -> int:
    return self.__height

  @property
  def value(self) -> T:
    return self.__value

  @property
  def root(self) -> '_Node[T]':
    if self.is_root():
      return self
    self.__parent = self.__parent.root
    return self.__parent


class UnionFindTree(Generic[T]):
  def __init__(self, values: Sequence[T]) -> None:
    self.__trees = list(map(lambda x: _Node(x), values))

  def union(self, xi: int, yi: int) -> None:
    self.__trees[xi].union(self.__trees[yi])

  def is_same(self, xi: int, yi: int) -> bool:
    return self.__trees[xi].root.value == self.__trees[yi].root.value
    

if __name__ == "__main__":
  import sys
  input = sys.stdin.readline
  n_q = input().split()
  n = int(n_q[0])
  q = int(n_q[1])

  uft = UnionFindTree(range(n))
  output = ""
  for _ in range(q):
    line = input().split()
    x, y = int(line[1]), int(line[2])
    if line[0] == "0":
      uft.union(x, y)
    else:
      if uft.is_same(x, y):
        output += "1\n"
      else:
        output += "0\n"

  print(output, end="")