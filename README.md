# Wukong

[![Build status](https://github.com/TonnyL/Wukong/workflows/Build/badge.svg)](https://github.com/TonnyL/Wukong/actions?query=workflow%3ABuild)

A command-line tool for browsing GitHub trending written by Rust.

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
% wukong repos --lang x --period y --spoken_lang z
% wukong repos -l x -p y -s z
```

Parameters:
+ `lang`: **optional**, default to `all` which stands for all the languages, see [all the options](#list-all-programming-language-options).
+ `period`: **optional**, default to `daily`, possible values: `daily`, `weekly` and `monthly`.
+ `spoken_lang`: **optional**, list trending repositories of certain spoken languages (e.g English, Chinese), see [all the options](#list-all-spoken-language-options).

```shell script
% wukong repos --lang rust --period daily

+------+---------------------+--------------------------+----------+---------------------+
| Rank | Full Name           | Description              | Language | Stars(Total/Period) |
+------+---------------------+--------------------------+----------+---------------------+
| 1    | diesel-rs/diesel    | A safe, extensible ORM a | Rust     | 5368/5              |
|      |                     | nd Query Builder for Rus |          |                     |
|      |                     | t                        |          |                     |
+------+---------------------+--------------------------+----------+---------------------+
| 2    | 996icu/996.ICU      | Repo for counting stars  | Rust     | 249376/24           |
|      |                     | and contributing. Press  |          |                     |
|      |                     | F to pay respect to glor |          |                     |
|      |                     | ious developers.         |          |                     |
+------+---------------------+--------------------------+----------+---------------------+
| 3    | Rust-SDL2/rust-sdl2 | SDL2 bindings for Rust   | Rust     | 1271/1              |
+------+---------------------+--------------------------+----------+---------------------+
| ...  | ...                 | ...                      | ...      | ...                 |
+------+---------------------+--------------------------+----------+---------------------+
```

### Find trending developers
```shell script
% wukong devs --lang x --period y
% wukong devs -l x -p y
```

Parameters:
+ `lang`: **optional**, default to `all` which stands for all the languages, see [all the options](#list-all-programming-language-options).
+ `period`: **optional**, default to `daily`, possible values: `daily`, `weekly` and `monthly`.

```shell script
% wukong devs

+------+---------------------------------+----------------+--------------------------+
| Rank | Name                            | Repo Name      | Description              |
+------+---------------------------------+----------------+--------------------------+
| 1    | Tim Paine(timkpaine)            | algo-coin      | Python library for algor |
|      |                                 |                | ithmic trading cryptocur |
|      |                                 |                | rencies across multiple  |
|      |                                 |                | exchanges                |
+------+---------------------------------+----------------+--------------------------+
| 2    | Kyle Mathews(KyleAMathews)      | typography.js  | A powerful toolkit for b |
|      |                                 |                | uilding websites with be |
|      |                                 |                | autiful design           |
+------+---------------------------------+----------------+--------------------------+
| 3    | XhmikosR(XhmikosR)              | notepad2-mod   | LOOKING FOR DEVELOPERS - |
|      |                                 |                |  Notepad2-mod, a Notepad |
|      |                                 |                | 2 fork, a fast and light |
|      |                                 |                | -weight Notepad-like tex |
|      |                                 |                | t editor with syntax hig |
|      |                                 |                | hlighting                |
+------+---------------------------------+----------------+--------------------------+
| 4    | Forbes Lindesay(ForbesLindesay) | redux-optimist | Optimistically apply act |
|      |                                 |                | ions that can be later c |
|      |                                 |                | ommited or reverted.     |
+------+---------------------------------+----------------+--------------------------+
| 5    | ...                             | ...            | ...                      |
+------+---------------------------------+----------------+--------------------------+
```

### List all programming language options
```shell script
% wukong langs

+-------------------+-------------------+
| name              | value             |
+-------------------+-------------------+
| 1C Enterprise     | 1c-enterprise     |
+-------------------+-------------------+
| ABAP              | abap              |
+-------------------+-------------------+
| ABNF              | abnf              |
+-------------------+-------------------+   
| ...               | ...               |  
+-------------------+-------------------+  
```

### List all spoken language options
```shell script
% wukong spoken_langs

+-----------------+-------+
| name            | value |
+-----------------+-------+
| Abkhazian       | ab    |
+-----------------+-------+
| Afar            | aa    |
+-----------------+-------+
| Afrikaans       | af    |
+-----------------+-------+
| Akan            | ak    |
+-----------------+-------+
| ...             | ...   |
+-----------------+-------+
```

## License
Wukong is under an MIT license. See the [LICENSE](LICENSE) for more information.