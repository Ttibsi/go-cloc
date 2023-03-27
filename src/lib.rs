use std::collections::HashMap;
use std::error::Error;
use std::fmt::Write;
use std::fs::File;
use std::io::{BufRead, BufReader};
use std::{fs, path::Path};

fn check_for_files(dir: &str) -> Result<HashMap<String, i32>, Box<dyn Error>> {
    let mut results: HashMap<String, i32> = HashMap::new();

    for entry_result in fs::read_dir(dir)? {
        let entry_path = entry_result?.path();
        //
        //TODO: Ignore what's in .gitignore
        if entry_path.starts_with("./.git") {
            continue;
        }
        if entry_path.starts_with("./target") {
            continue;
        }

        if entry_path.is_dir() {
            // Handle directory
            results.extend(check_for_files(entry_path.to_str().unwrap()));
        } else if entry_path.is_file() {
            // Handle file
            let ext = Path::new(entry_path.to_str().unwrap())
                .extension()
                .and_then(|ext| ext.to_str())
                .unwrap_or("")
                .to_string();

            if results.get(&ext).is_some() {
                results.insert(
                    ext.clone(),
                    (BufReader::new(File::open(entry_path)?).lines().count()
                        + results.get(&ext.clone()).unwrap().to_owned())
                    .try_into()
                    .unwrap(),
                );
            } else {
                results.insert(
                    ext.clone(),
                    BufReader::new(File::open(entry_path)).lines().count(),
                );
            }
        } else {
            // Handle other types of entries, like symbolic links
            println!("Found other type of entry: {:?}", entry_path);
        }
    }

    return Ok(results);
}

pub fn scan(args: Vec<String>) -> HashMap<String, i32> {
    let mut results: HashMap<String, i32> = HashMap::new();

    for arg in args {
        let ret = check_for_files(&arg);
        results.extend(ret);
    }

    return results;
}

pub fn print(res: HashMap<String, i32>) {
    let mut output = String::new();
    let mut len = 0;

    for (k, v) in res.iter() {
        if len < k.len() {
            len = k.len();
        }

        write!(&mut output, "| {} | {} |\n", k, v.to_string()).unwrap();
    }

    println!("+{}+", "-".repeat(len + 7));
    println!("{}", output);
    println!("+{}+", "-".repeat(len + 7));

    // println!("{:#?}", res)
}
