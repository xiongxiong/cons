An easy to use, key lock tool, lock according to the same key, not lock on different keys.

## Contents

<!-- TOC -->

- [Contents](#contents)
- [Features](#features)
- [Requirements](#requirements)
- [Installation](#installation)
- [Usage](#usage)
- [Example](#example)
    - [Get cons (lock map)](#get-cons-lock-map)
    - [Wait on key using](#wait-on-key-using)
    - [Skip on key using](#skip-on-key-using)
    - [Queue on key using](#queue-on-key-using)
- [Author](#author)
- [License](#license)

<!-- /TOC -->

## Features
- [x] wait on key using, for concurrent usages, only the first usage will execute, others will wait for the first to finish, and then all can continue at the same time, used for waiting for something to happen once and only once, and it's a forerunner condition.
- [x] skip on key using, for concurrent usages, only the first usage will execute, others will be marked as skip, used for something to happen once at the same time.
- [x] queue on key using, for concurrent usages, usages will be queued one by one, until all finished.

## Requirements

- golang 1.10 +

## Installation

- go get -u -v github.com/xiongxiong/cons

## Usage

- If you have only a collection of conditions, you can use the default cons. If you have several collections of conditions which may have same key, you should create different cons for collections to avoid collisions.

## Example

### Get cons (lock map)

```go
var cons = GetCons()
```

### Wait on key using

```go
c := cons.Wait("hello")
defer c.Done()
```

### Skip on key using

```go
c := cons.Skip("hello")
defer c.Done()

if c.Skip {
    println("skip")
    return
}
```

### Queue on key using

```go
c := cons.Queue("hello")
defer c.Done()
```

## Author

xiongxiong, ximengwuheng@163.com

## License

The package cons is available under the MIT license. See the LICENSE file for more info.
