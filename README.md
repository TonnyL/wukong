# Wukong

![Build status](https://github.com/TonnyL/Wukong/workflows/Go/badge.svg)

A command-line tool for browsing GitHub trending repositories and developers written by Go.

> The Monkey King, known as Sun Wukong in Chinese, is a legendary figure best known as one of the main characters in the 16th-century Chinese novel Journey to the West (西游记/西遊記) and many later stories and adaptations.
> 
>https://en.wikipedia.org/wiki/Monkey_King

## Installation
Install the **Wukong** binary via **Homebrew**:
```shell script
% brew install TonnyL/tap/Wukong
``` 
Or download [archives of precompiled binaries](https://github.com/TonnyL/Wukong/releases).

## Usage
### Find trending repositories
```shell script
% wukong repo -lang x -period y
```

Parameters:
+ `lang`: **optional**, default to an empty string("") which stands for all the languages, see [all the options](#list-all-language-options).
+ `period` **optional**, default to `daily`, possible values: `daily`, `weekly` and `monthly`.

```shell script
% wukong repo -lang go -period daily

+------+----------------------------+--------------------------+----------+---------------------+-------------------------------------------------------------------+
| RANK |            NAME            |       DESCRIPTION        | LANGUAGE | STARS(TOTAL/PERIOD) |                                URL                                |
+------+----------------------------+--------------------------+----------+---------------------+-------------------------------------------------------------------+
|    1 | OpenDiablo2                | An open source re-implem | Go       | 2626/1625           | https://github.com/OpenDiablo2/OpenDiablo2                        |
|      |                            | entation of Diablo 2     |          |                     |                                                                   |
+------+----------------------------+--------------------------+----------+---------------------+-------------------------------------------------------------------+
|    2 | validator                  | 💯Go Struct and Field    | Go       | 4154/149            | https://github.com/go-playground/validator                        |
|      |                            | validation, including Cr |          |                     |                                                                   |
|      |                            | oss Field, Cross Struct, |          |                     |                                                                   |
|      |                            |  Map, Slice and Array di |          |                     |                                                                   |
|      |                            | ving                     |          |                     |                                                                   |
+------+----------------------------+--------------------------+----------+---------------------+-------------------------------------------------------------------+
|    4 | grpc-go                    | The Go language implemen | Go       | 10019/58            | https://github.com/grpc/grpc-go                                   |
|      |                            | tation of gRPC. HTTP/2 b |          |                     |                                                                   |
|      |                            | ased RPC                 |          |                     |                                                                   |
+------+----------------------------+--------------------------+----------+---------------------+-------------------------------------------------------------------+
|    . | ...                        | ...                      | ...      |  ...                | ...                                                               |
+------+----------------------------+--------------------------+----------+---------------------+-------------------------------------------------------------------+

```

### Find trending developers
```shell script
% wukong dev -lang x -period y
```

Parameters:
+ `lang`: **optional**, default to an empty string("") which stands for all the languages, see [all the options](#list-all-language-options).
+ `period` **optional**, default to `daily`, possible values: `daily`, `weekly` and `monthly`.

```shell script
% wukong dev

+------+------------------------------+--------------------------+----------------------------------+
| RANK |             NAME             |  REPO NAME/DESCRIPTION   |               URL                |
+------+------------------------------+--------------------------+----------------------------------+
|    1 | Alon Zakai(kripken)          | sql.js - SQLite compiled | https://github.com/kripken       |
|      |                              |  to JavaScript through E |                                  |
|      |                              | mscripten                |                                  |
+------+------------------------------+--------------------------+----------------------------------+
|    2 | Klaus Post(klauspost)        | compress - Optimized com | https://github.com/klauspost     |
|      |                              | pression packages        |                                  |
+------+------------------------------+--------------------------+----------------------------------+
|    3 | siddontang(siddontang)       | ledisdb - a high perform | https://github.com/siddontang    |
|      |                              | ance NoSQL powered by Go |                                  |
+------+------------------------------+--------------------------+----------------------------------+
|    . | ...                          | ...                      | ...                              |
+------+------------------------------+--------------------------+----------------------------------+
```

### List all language options
```shell script
% wukong lang
```

```shell script
% wukong lang

+--------------------------------+--------------------------------+
|               ID               |              NAME              |
+--------------------------------+--------------------------------+
|                                | All languages                  |
+--------------------------------+--------------------------------+
| 1c-enterprise                  | 1C Enterprise                  |
+--------------------------------+--------------------------------+
| abap                           | ABAP                           |
+--------------------------------+--------------------------------+
| ...                            | ...                            |
+--------------------------------+--------------------------------+
```

## License
Wukong is under an MIT license. See the [LICENSE](LICENSE) for more information.