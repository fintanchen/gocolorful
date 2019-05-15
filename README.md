# gocolorful

_A tiny colorful output library_

This library provider a simple wrap of log output with level.

## Color Definitions

The true color you can select by run the `256-colors.go` in the examples directory.

ANSI-SEQUENCE map:

**8 Colors:**

|  color  |     code     |
| :-----: | :----------: |
|  Black  | \\u001b\[30m |
|   Red   | \\u001b\[31m |
|  Green  | \\u001b\[32m |
|  Yellow | \\u001b\[33m |
|   Blue  | \\u001b\[34m |
| Magenta | \\u001b\[35m |
|   Cyan  | \\u001b\[36m |
|  White  | \\u001b\[37m |
|  Reset  |  \\u001b\[0m |

**Bright Colors:**

|      color     |      code      |
| :------------: | :------------: |
|  Bright Black  | \\u001b\[30;1m |
|   Bright Red   | \\u001b\[31;1m |
|  Bright Green  | \\u001b\[32;1m |
|  Bright Yellow | \\u001b\[33;1m |
|   Bright Blue  | \\u001b\[34;1m |
| Bright Magenta | \\u001b\[35;1m |
|   Bright Cyan  | \\u001b\[36;1m |
|  Bright White  | \\u001b\[37;1m |

## Examples

Show in `examples` folder. It shows how to use and become more useful by wrapper the standard library.

-   Hello World
-   256-colors.go
-   Wrapper `log` (WIP)
-   Wrapper `fmt` (WIP)
