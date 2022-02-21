# N1QL Query Builder
**Fluent Couchbase N1QL Query Builder for Go**

_The API is currently experimental and may change._

## Statements

### [`SELECT`](https://developer.couchbase.com/documentation/server/current/n1ql/n1ql-language-reference/select-syntax.html)

See the [godoc]() or [the tests](select_test.go) for examples.

## Expressions

http://godoc.org/github.com/wheniwork/n1ql-query-builder/#Expression

## Aspect of gocb-dsl

gocb-dsl has multi-function/feature-set that may help the certain developer to write more
readable and long-life `couchbase queries`. It has simple code-base that any developer can contribute as well.
<br><br>
_gocb-dsl is not ORM but I've future plan that contains gocb-dsl will be used in ORM project_

## TODO

* More functions
    * [x] Aggregate
    * [ ] Array
    * [ ] Case
    * [ ] Collections
    * [ ] Comparison
    * [ ] Conditional
    * [ ] Date
    * [ ] JSON
    * [x] Meta
    * [x] Number
    * [ ] Object
    * [ ] Pattern Matching
    * [x] String
    * [ ] Type
    * [x] Nested Query
* DML statements
* Tests
    * [ ] Unit Test
    * [ ] Integration Test
