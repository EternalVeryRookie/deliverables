import unittest
from dataclasses import dataclass
import json
from pathlib import Path
import time
import random

from union_find_tree import UnionFindTree

@dataclass
class TestCase:
  n: int
  q: int
  queries: list[list[int]]
  want: list[bool]

  @staticmethod
  def from_json(filepath: str):
    with open(filepath) as f:
      obj = json.load(f)

    return TestCase(obj["n"], obj["q"], obj["queries"], obj["want"])

class UnionTestCase(unittest.TestCase):
  def test_union(self):
    testcase_filepath = Path(__file__).parent / "union_testcase.json"
    testcase = TestCase.from_json(testcase_filepath)
    uft = UnionFindTree(list(range(testcase.n)))
    actual = self.__apply_query(testcase.queries, uft)

    self.assertListEqual(testcase.want, actual)

  def test_performance(self):
    data_size = 10000
    query_size = 100000
    uft = UnionFindTree(list(range(data_size)))
    def gen_random_query(size: int):
      tmp = [random.randint(0, 1)]
      tmp.extend(random.sample(range(size), 2))
      return tmp

    queries = map(lambda x: gen_random_query(data_size), range(query_size))
    start_time = time.process_time()
    self.__apply_query(queries, uft)
    end_time = time.process_time()
    #100000個のクエリーを2秒以内で完了させることを期待
    self.assertLess(end_time - start_time, 2)

  def __apply_query(self, queries: list[list[int]], uft: UnionFindTree) -> list[bool]:
    actual = []
    i = 0
    for query in queries:
      i += 1
      if query[0] == 0:
        uft.union(query[1], query[2])
      else:
        is_same = uft.is_same(query[1], query[2])
        actual.append(is_same)

    return actual

if __name__ == "__main__":
  unittest.main()