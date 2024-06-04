import gleam/dict
import gleam/dynamic
import gleam/function
import gleam/int
import gleam/io
import gleam/iterator
import gleam/json
import gleam/list
import gleam/result
import gleam/string
import stdin.{stdin}

pub fn main() {
  stdin()
  |> iterator.flat_map(fn(it) {
    json.decode(it, dynamic.dict(dynamic.string, dynamic.dynamic))
    |> result.lazy_unwrap(dict.new)
    |> dict.keys
    |> iterator.from_list
  })
  |> iterator.group(function.identity)
  |> dict.map_values(fn(_, value) { list.length(value) })
  |> dict.to_list
  |> list.sort(fn(a, b) { int.compare(a.1, b.1) })
  |> list.each(fn(it) {
    [it.0, " = ", int.to_string(it.1)]
    |> string.concat
    |> io.println
  })
}
