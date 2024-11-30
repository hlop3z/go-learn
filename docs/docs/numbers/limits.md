# Numbers Limits

---

## Integer

The range for an `int64` in Go is based on its size (64 bits) and the two's complement representation used for signed integers:

```bash
9223372036854775807
```

### **Range of `int64`**

- **Minimum value**: `-9,223,372,036,854,775,808` (i.e., `-(2^63)`)
- **Maximum value**: `9,223,372,036,854,775,807` (i.e., `2^63 - 1`)

### **Key Points**

- `int64` is a **signed 64-bit integer**.
- It can represent approximately **±9.2 quintillion**.

This makes `int64` ideal for large numerical computations or cases requiring very large integer ranges. If you don't need negative values, you can use `uint64`, which has a range of `0` to `18,446,744,073,709,551,615` (i.e., `2^64 - 1`).

---

## Float

The range for a `float64` in Go (as defined by the IEEE 754 standard for 64-bit floating-point numbers) is:

### Min

```bash
0.000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000005
```

### Max

```bash
179769313486231570814527423731704356798070567525844996598917476803157260780028538760589558632766878171540458953514382464234321326889464182768467546703537516986049910576551282076245490090389328944075868508455133942304583236903222948165808559332123348274797826204144723168738177180919299881250404026184124858368.000000
```

### **Positive Range**

- **Smallest positive non-zero value (denormalized)**: `5e-324`
- **Largest positive value**: `1.7976931348623157e+308`

### **Negative Range**

- **Largest negative value**: `-1.7976931348623157e+308`
- **Smallest negative non-zero value (denormalized)**: `-5e-324`

### **Special Values**

- **Zero**: `0` (positive and negative)
- **Infinity**: `+Inf` and `-Inf`
- **Not a Number (NaN)**: Represents undefined or unrepresentable values (e.g., `0/0`).

### **Precision**

- A `float64` provides approximately **15–17 decimal digits** of precision.

This makes `float64` suitable for a wide range of scientific and engineering applications requiring high precision and a large range of values.

---

## Safe Numbers (JavaScript and Business Logic)

| Title               | Number                    | Notes                                                              |
| ------------------- | ------------------------- | ------------------------------------------------------------------ |
| Minimum             | **`0.00001`**             | (100 Thousandth) Minimum safe decimal precision for user input     |
| Maximum             | **`1000000000`**          | (1 Billion) Maximum safe integer for business logic                |
| Safe Numbers        | **`1000000000.00001`**    | (1 Billion + 100 Thousandth) Decimals => `SQL_DECIMAL(15, 5)`      |
| Storage Max Integer | **`1000000000000000000`** | (1 Quintillion) Maximum value supported by Database/GoLang storage |

```go
// ------------------------------------------------------------------------
// Safe Numbers (JavaScript and Business Logic)
// ------------------------------------------------------------------------

const (
 // MinSafeDecimal represents the smallest safe decimal value for user input.
 MinSafeDecimal float64 = 0.00001 // (100 Thousandth)

 // MaxSafeInteger represents the largest safe integer value for business logic.
 MaxSafeInteger int64 = 1000000000 // (1 Billion)

 // MaxSafeNumbers represents the largest safe decimal value for precision.
 MaxSafeNumbers float64 = 1000000000.00001 // (1 Billion + 100 Thousandth)

 // StorageMaxInteger represents the maximum value supported by database/GoLang storage.
 StorageMaxInteger int64 = 1000000000000000000 // (1 Quintillion)
)
```
