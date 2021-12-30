from typing import Sequence, Tuple, TypeVar, Generic, Union

T = TypeVar("T")

class _Node(Generic[T]):
  def __init__(self, value: T) -> None:
    self.__value = value
    self.weight: Union[int, None] = None
    self.__height = 1
    self.__parent: Union['_Node[T]', None] = None

  #target - self = delta
  def union(self, target: '_Node[T]', delta: int):
    self_root = self.root()
    target_root = target.root()
    if self_root[0].is_equal(target_root[0]):
      return

    new_weight = delta + target_root[1] - self_root[1]
    if self_root[0].height <= target_root[0].height:
      self_root[0].__parent = target_root[0]
      self_root[0].weight = new_weight
      if self_root[0].height == target_root[0].height:
        target_root[0].__height += 1
    else:
      target_root[0].__parent = self_root[0]
      target_root[0].weight = -new_weight
      if self_root[0].height == target_root[0].height:
        self_root[0].__height += 1

  def is_equal(self, target: '_Node[T]') -> bool:
    return target.__value == self.__value

  def is_root(self):
    return self.__parent is None

  def total_weight_to_root(self):
    if self.is_root():
      return 0

    return self.wieght + self.__parent.total_weight_to_root()

  def root(self) -> Tuple['_Node[T]', int]:
    root = self
    total_weight = 0
    while not root.is_root():
      total_weight += root.weight
      root = root.__parent

    return (root, total_weight)

  @property
  def height(self) -> int:
    return self.__height


class WeightedUnionFindTree(Generic[T]):
  def __init__(self, values: Sequence[T]):
    self.__node_set = list(map(lambda x: _Node(x), values))

  def union(self, xi: int, yi: int, weight: int):
    if weight < 0:
      self.union(yi, xi, -weight)

    self.__node_set[xi].union(self.__node_set[yi], weight)

  # xi - yi
  def diff(self, xi: int, yi: int) -> Union[int, None]:
    xi_root = self.__node_set[xi].root()
    yi_root = self.__node_set[yi].root()

    if not xi_root[0].is_equal(yi_root[0]):
      return None

    return yi_root[1] - xi_root[1]


if __name__ == "__main__":
  import sys
  input = sys.stdin.readline
  n_q = input().split()
  n = int(n_q[0])
  q = int(n_q[1])

  wuft = WeightedUnionFindTree(range(n))
  output = ""
  for _ in range(q):
    line = input().split()
    x, y = int(line[1]), int(line[2])
    if line[0] == "0":
      wuft.union(x, y, int(line[3]))
    else:
      delta = wuft.diff(y, x)
      if delta is None:
        output += "?\n"
      else:
        output += f"{delta}\n"

  print(output, end="")