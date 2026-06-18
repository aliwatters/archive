# Learn Rust

12/24/22 - https://doc.rust-lang.org/book/

## Chapter 01 Notes; 12/24

- `println!("string");` -- `name!()` run a macro,`name()` run a function.
- `rustc` -- compiler
- `cargo` -- is the tool chain
  - `cargo new project_name` -- new dir and scaffolding in `project_name` -- note it's a git init'd repo
  - `cargo build` -- makes a binary in `target` dir
  - `cargo run` -- builds and runs
  - `cargo check` -- doesn't run or build, just syntax check
  - `cargo build --release` -- makes a prod version

Typical workflow;

```
$ git clone example.org/someproject
$ cd someproject
$ cargo build
```

Note: `rm -r chapter-01/hello_cargo/.git` so that I have a single repo for this.

## Chapter 02 Notes; 12/24

```
let apples = 5; // immutable
let mut bananas = 5; // mutable
```

Immutable by default, mut is mutable.

- `.read_line(&mut guess)` -- & indicates a reference
- `.expect("msg')` -- echos `msg` on error
- `{var}` -- var interpolation

```
let x = 5;
let y = 10;

println!("x = {x} and y + 2 = {}", y + 2);
```

Dependencies default to carat, "0.x.y" is "^0.x.y" (in `Cargo.toml`), run `cargo build` to bring new deps in.
