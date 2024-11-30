# Operators

---

## Python

### Comparison Operators

- `lt` (Less Than `<`)
- `lte` (Less Than or Equal To `<=`)
- `gt` (Greater Than `>`)
- `gte` (Greater Than or Equal To `>=`)
- `eq` (Equal `==`)
- `ne` (Not Equal `!=`)

### Logical Operators

- `and` (Logical AND `and`)
- `or` (Logical OR `or`)
- `not` (Logical NOT `not`)

### Membership and Identity Operators

- `in` (Membership `in`)
- `is` (Identity `is`)

### Special Operators (for frameworks or advanced use)

- `contains` (Check containment)
- `startswith` (String starts with)
- `endswith` (String ends with)

---

---

## SQL

### **Comparison Operators**

- `=` : Equal
- `<>` or `!=` : Not Equal
- `<` : Less Than
- `<=` : Less Than or Equal To
- `>` : Greater Than
- `>=` : Greater Than or Equal To

---

### **Logical Operators**

- `AND` : Logical AND
- `OR` : Logical OR
- `NOT` : Logical NOT

---

### **Arithmetic Operators**

- `+` : Addition
- `-` : Subtraction
- `*` : Multiplication
- `/` : Division
- `%` : Modulus (used for remainders, supported in some SQL dialects)

---

### **String Operators**

- `||` : Concatenate strings (SQL standard, used in PostgreSQL)
- `+` : Concatenate strings (used in SQL Server)
- `CONCAT` : Concatenate strings (MySQL, some others)

---

### **Pattern Matching Operators**

- `LIKE` : Search for a pattern (e.g., `WHERE column LIKE 'A%'`)
- `NOT LIKE` : Search for values that don't match a pattern
- `ILIKE` : Case-insensitive `LIKE` (PostgreSQL)

---

### **Set Operators**

- `IN` : Match a value in a list (e.g., `WHERE column IN ('A', 'B')`)
- `NOT IN` : Match a value not in a list
- `EXISTS` : Check if a subquery returns rows
- `NOT EXISTS` : Check if a subquery returns no rows

---

### **NULL Check Operators**

- `IS NULL` : Check if a value is `NULL`
- `IS NOT NULL` : Check if a value is not `NULL`

---

### **Comparison with Subqueries**

- `ANY` : Compare to any value returned by a subquery (e.g., `WHERE column > ANY (subquery)`)
- `ALL` : Compare to all values returned by a subquery (e.g., `WHERE column > ALL (subquery)`)

---

### **Joins and Relationships**

- `=` : Equality in join conditions
- `ON` : Specifies join conditions
- `USING` : Specifies join columns in certain SQL dialects

---

### **Specialized Operators**

- `BETWEEN` : Specify a range (inclusive, e.g., `WHERE column BETWEEN 10 AND 20`)
- `NOT BETWEEN` : Specify a range exclusion
- `DISTINCT` : Eliminate duplicates (used with `SELECT`)
- `UNION` : Combine results from multiple queries, removing duplicates
- `UNION ALL` : Combine results from multiple queries, including duplicates
- `INTERSECT` : Return rows common to multiple queries
- `EXCEPT` or `MINUS` : Return rows from one query not in another (depends on SQL dialect)

---

### **Full-Text Search Operators**

(Supported in certain SQL dialects)

- `MATCH` : Full-text search (e.g., `MATCH(column) AGAINST ('text')` in MySQL)
- `CONTAINS` : Full-text search in SQL Server
- `TO_TSVECTOR` and `TO_TSQUERY` : Full-text search in PostgreSQL

---

### **Bitwise Operators**

(Supported in some SQL dialects)

- `&` : Bitwise AND
- `|` : Bitwise OR
- `^` : Bitwise XOR
- `~` : Bitwise NOT
- `<<` : Left shift
- `>>` : Right shift

---

### **Aggregate Operators**

(Used with aggregate functions)

- `GROUP BY` : Group rows
- `HAVING` : Filter grouped rows
