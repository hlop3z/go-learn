# Python | JavaScript | SQL | Go | JSON | GraphQL

| **Field/Type**   | **Python**                 | **JavaScript**                | **SQL**                            | **Go**                              | **JSON**                                           | **GraphQL**                                 |
| ---------------- | -------------------------- | ----------------------------- | ---------------------------------- | ----------------------------------- | -------------------------------------------------- | ------------------------------------------- |
| **Integer**      | `int`                      | `Number`                      | `INTEGER`                          | `int`, `int32`, `int64`             | `number`                                           | `Int`                                       |
| **Float**        | `float`                    | `Number`                      | `FLOAT`, `DOUBLE`                  | `float32`, `float64`                | `number`                                           | `Float`                                     |
| **String**       | `str`                      | `String`                      | `VARCHAR`, `TEXT`                  | `string`                            | `string`                                           | `String`                                    |
| **Boolean**      | `bool`                     | `Boolean`                     | `BOOLEAN`                          | `bool`                              | `true`, `false`                                    | `Boolean`                                   |
| **Array/List**   | `list`, `tuple`            | `Array`                       | `ARRAY`                            | `[]type` (slice)                    | `[]type` (array)                                   | `[Type]`                                    |
| **Object/Dict**  | `dict`                     | `Object`                      | `JSON`, `TEXT` (via serialization) | `map[string]Type` (map)             | `object`                                           | `Object`                                    |
| **Null**         | `None`                     | `null`                        | `NULL`                             | `nil`                               | `null`                                             | `null`                                      |
| **Date/Time**    | `datetime`, `date`         | `Date`, `DateTime`            | `DATE`, `TIMESTAMP`                | `time.Time`                         | `string` (ISO 8601)                                | `DateTime`                                  |
| **UUID**         | `uuid.UUID`                | `UUID` (via library)          | `UUID`                             | `uuid.UUID`                         | `string` (UUID string)                             | `ID` (commonly used for unique identifiers) |
| **Binary**       | `bytes`                    | `Buffer`                      | `BLOB`, `BYTEA`                    | `[]byte`                            | `string` (Base64)                                  | (No specific type, usually as string)       |
| **Float or Int** | `Union[int, float]`        | `Number`                      | `INTEGER` or `FLOAT`               | `float64` or `int`                  | `number`                                           | `Float` (can be int or float)               |
| **Any**          | `Any` (`object`, `Union`)  | `any` (via `any` or `Object`) | (depends on DB type)               | `interface{}`                       | `any` (using `interface{}` for complex or unknown) | `Any` (usually as `Scalar` type)            |
| **Enum**         | `Enum` (via `enum` module) | Enum (via objects)            | `ENUM`                             | `enum` (custom or string constants) | `string` (set of allowed values)                   | `enum` (with specific allowed values)       |

## **Explanations**

- **Python**: Commonly uses `int`, `str`, `float`, `bool`, and `dict` for data types. For JSON-like structures, `dict` is used, and `list` or `tuple` is used for arrays.
- **JavaScript**: Primarily uses `Number` for integers and floats, `String` for strings, and `Object` for JSON objects. JavaScript uses `Array` for lists.
- **SQL**: SQL has its own standard types like `INTEGER`, `VARCHAR`, and `BOOLEAN`. For date/time, SQL uses `DATE`, `TIME`, and `TIMESTAMP`.
- **Go**: Go uses `int`, `float64`, `string`, `bool`, and `map` for maps. Arrays in Go are slices (`[]Type`). Go does not directly have an equivalent for `None` but uses `nil` for uninitialized pointers or slices.
- **JSON**: In JSON, types are dynamic and loosely typed, such as `number` (for both integers and floats), `string`, `boolean`, and `object` (equivalent to a JSON object). Arrays in JSON are represented as a list of items.
- **GraphQL**: GraphQL uses types like `Int`, `Float`, `String`, `Boolean`, and `ID` (usually for unique identifiers). It allows complex types like `Object` and `Enum` to define allowable values or structures. The `Any` type is often represented as `Scalar` in GraphQL.

### **Special Notes**

- **UUIDs**: Both Python and Go have specialized libraries to handle UUIDs. In JSON and GraphQL, UUIDs are typically serialized as strings.
- **Date/Time**: All the systems handle date and time, but the format can differ. Most systems (Go, Python, GraphQL, JSON) use ISO 8601 string representations for date and time.
- **Null Values**: The concept of `null` is the same across most systems, though in Go, it’s represented as `nil` for references. SQL uses `NULL`, and GraphQL uses `null`.
