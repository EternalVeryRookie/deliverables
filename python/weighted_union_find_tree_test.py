import random
import time
import unittest
from pathlib import Path
from dataclasses import dataclass
import json

from weighted_union_find_tree import WeightedUnionFindTree

@dataclass(frozen=True)
class Testcase:
  n: int
  q: int
  queries: list[list[int]]
  want: list[int|None]

  @staticmethod
  def from_json(filepath:str) -> 'Testcase':
    with open(filepath) as f:
      obj = json.load(f)

    return Testcase(n = obj["n"], q = obj["q"], queries = obj["queries"], want = obj["want"])


class WeightedUnionTestCase(unittest.TestCase):
  def test_union(self):
    testcase_filepath = Path(__file__).parent / "weighted_union_testcase.json"
    testcase = Testcase.from_json(testcase_filepath)
    wuft = WeightedUnionFindTree(range(testcase.n))
    actual = self.__apply_queries(testcase.queries, wuft)
    self.assertListEqual(actual, testcase.want)
  
  def test_performance(self):
    data_size = 100000
    query_size = 200000
    uft = WeightedUnionFindTree(range(data_size))
    def gen_random_query(size: int):
      query_type = random.randint(0, 1)
      x = random.randint(0, size-2)
      y = random.randint(x+1, size-1)
      if query_type == 0:
        return [0, x, y, 0]
      else:
        return [1, x, y]

    queries = map(lambda x: gen_random_query(data_size), range(query_size))
    start_time = time.process_time()
    self.__apply_queries(queries, uft)
    end_time = time.process_time()
    #200000個のクエリーを2秒以内で完了させることを期待
    self.assertLess(end_time - start_time, 2)

  def __apply_queries(self, queries:list[list[int]], wuft: WeightedUnionFindTree) -> list[int|None]:
    actual: list[int|None] = []
    for query in queries:
      x = query[1]
      y = query[2]
      if query[0] == 0:
        z = query[3]
        wuft.union(x, y, z)
      else:
        delta = wuft.diff(y, x)
        actual.append(delta)
        
    return actual


if __name__ == "__main__":
  unittest.main()