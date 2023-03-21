use std::env;

fn main() {
    let res = cl::scan(env::args().skip(1).collect());
    cl::print(res);
}
