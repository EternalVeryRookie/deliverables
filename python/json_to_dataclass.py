
from types import GenericAlias, UnionType
from typing import Literal, NewType, Tuple, Type, TypeVar, Union, Final
import json

T = TypeVar("T")

class UnMatchTypeError(Exception):
  "annotationの型と実際の値の型が異なる場合に発生する例外"
  def __init__(self, cls: type, member_name: str, want_type: type, actual_type: type):
    message = f"{cls}.{member_name}に{actual_type}の値を代入できません。 型指定は{want_type}です。"
    super().__init__(message)


def _to_list(l: list, cls: Type[T]) -> list[T]:
  def convert_list_elem_to_cls(elem):
    if type(elem) == dict:
      return to_cls(elem, cls)
    elif type(elem) == list:
      return _to_list(elem, cls)
    else:
      return elem

  return list(map(convert_list_elem_to_cls, l))


def _convert(value: None|bool|int|float|str|list|dict, annotation: Type[T]) -> Tuple[T, bool]:
  generic_type = None
  if isinstance(annotation, GenericAlias):
    generic_type = annotation
    annotation = annotation.__origin__

  if isinstance(value, annotation):
  # listに対して型パラメーターが設定されている場合は、リストの各要素を指定された型になるように変換する
    if isinstance(value, list) and (generic_type is not None):
      return _to_list(value, generic_type.__args__[0]), True
    else:
      return value, True
  elif isinstance(value, dict):
    return to_cls(value, annotation), True
  else:
    return None, False


def to_cls(dic: dict, cls: Type[T]) -> T:
  if not hasattr(cls, "__annotations__"):
    return cls(**dic)

  args = {}
  for key in dic:
    # 型アノテーションがない場合は空文字になる模様
    if cls.__annotations__[key] == "":
      args[key] = dic[key]
      continue

    member_type = cls.__annotations__[key]
    # 型アノテーションを文字列で指定した場合は__annotations__[key]にその文字列が文字列型としてそのまま設定されるため
    # その文字列を評価して型を受け取る
    if isinstance(member_type, str):
      member_type = eval(member_type)

    if isinstance(member_type, type):
      args[key], is_success = _convert(dic[key], member_type)
      if not is_success:
        raise UnMatchTypeError(cls, key, member_type, type(dic[key]))
    else:
      # UnionやLiteral、Finalなどはannotationに記述できるがtypeのサブタイプではないためこちらの制御ブロックに入る
      member_type_origin = member_type
      if hasattr(member_type_origin, "__origin__"):
        member_type_origin = member_type_origin.__origin__

      if member_type_origin == Literal:
        if not(dic[key] in member_type.__args__):
          raise UnMatchTypeError(cls, key, member_type, type(dic[key]))

        args[key] = dic[key]
      elif member_type_origin == Union or member_type_origin == UnionType:
        # プリミティブのみサポートしてクラス、リテラル、辞書、リストなどはサポート対象外にする
        support_union_types = [type(None), bool, int, float, str, bytes]
        for arg in member_type.__args__:
          if not(arg in support_union_types):
            raise Exception(f"Union型でサポートされている型は{support_union_types}のみです。{arg}はサポート対象外です。")

        if isinstance(dic[key], member_type):
          args[key] = dic[key]
        else:
          raise Exception(f"{cls}.{key}に代入可能な型は{member_type}です。{type(dic[key])}は代入できません。")
      elif member_type_origin == Final:
        args[key], is_success = _convert(dic[key], member_type.__args__[0])
        if not is_success:
          raise UnMatchTypeError(cls, key, member_type, type(dic[key]))
      elif isinstance(member_type_origin, NewType):
        value, is_success = _convert(dic[key], member_type_origin.__supertype__)
        if not is_success:
          raise UnMatchTypeError(cls, key, member_type, type(dic[key]))
        args[key] = member_type_origin(value)
      else:
        raise Exception(f"サポート対象外の型です。{key}: {member_type}")


  return cls(**args)


def convert_to_dataclass(json_str: str, c: Type[T]) -> T:
  dic = json.loads(json_str)
  if isinstance(dic, list):
    raise Exception("リストはサポート対象外")

  return to_cls(dic, c)


def convert_to_dataclass_from_json_file(filepath: str, c: Type[T]) -> T:
  with open(filepath, "r") as f:
    dic = json.load(f)

  if isinstance(dic, list):
    raise Exception("リストはサポート対象外")

  return to_cls(dic, c)
