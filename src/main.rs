mod lib;
use std::collections::HashMap;
use std::env;

fn main() {
    let res = lib::scan(env::args().skip(1).collect());
    lib::print(res);
}
