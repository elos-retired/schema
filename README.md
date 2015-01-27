Schema [![GoDoc](https://godoc.org/github.com/elos/schema?status.svg)](https://godoc.org/github.com/elos/schema)
-------

Schema is a semi-high-level package that defines the framework for defining an object relation layer for go system applications. Note: this is likely not your traditional concept of an "ORM" persay. It is a much more light-weight.

Essentially a schema is a relationship map with validity checking which exposes the Link and Unlink functions to link two "models" of data.

As such, schema also introduces the concept of a Linkable model. It extends the notion of a [data](https://github.com/elos/data) Record.
