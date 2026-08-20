# jsontree

JsonTree implements [fastjson](https://github.com/valyala/fastjson)-like flexible
JSON value parsing using [stdlib json](https://pkg.go.dev/encoding/json) package.
The only reason we need it is to avoid deadling with `any` and to enforce type-
safety while minimising reliance on reflection.

I honestly don't feel like this thing is worth extracting into a separate library
but may do it at some point in the future. If you're reading this and thinking
that you want it in your project, maybe consider using fastjson instead. If that
is undesirable for some reason, consider creating an
[issue](https://github.com/imcrazytwkr/feedhub/issues).
