use std::collections::HashMap;
use std::fs::File;
use std::io::{BufRead, BufReader};
use std::{fs, path::Path};

fn get_lines(inp_file: String) -> i32 {
    let file = File::open(inp_file).expect("Unable to open file");
    let reader = BufReader::new(file);
    let mut line_count = 0;

    for _ in reader.lines() {
        line_count += 1;
    }

    return line_count;
}

fn check_for_files(dir: String) -> HashMap<String, i32> {
    let mut results: HashMap<String, i32> = HashMap::new();

    let dir_entries = match fs::read_dir(dir) {
        Ok(entries) => entries,
        Err(e) => {
            println!("Error reading directory: {:?}", e);
            return results;
        }
    };

    for entry_result in dir_entries {
        let entry = match entry_result {
            Ok(e) => e,
            Err(e) => {
                println!("Error reading directory entry: {:?}", e);
                continue;
            }
        };
        let entry_path = entry.path();
        //TODO: Ignore what's in .gitignore
        if entry_path.starts_with("./.git") {
            continue;
        }
        if entry_path.starts_with("./target") {
            continue;
        }

        if entry_path.is_dir() {
            // Handle directory
            results.extend(check_for_files(entry_path.to_str().unwrap().to_string()));
        } else if entry_path.is_file() {
            // Handle file
            let ext = Path::new(entry_path.to_str().unwrap())
                .extension()
                .and_then(|ext| ext.to_str())
                .unwrap_or("")
                .to_string();

            if results.get(&ext) != None {
                results.insert(
                    ext.clone(),
                    get_lines(entry_path.to_str().unwrap().to_string())
                        + results.get(&ext.clone()).unwrap().to_owned(),
                );
            } else {
                results.insert(
                    ext.clone(),
                    get_lines(entry_path.to_str().unwrap().to_string()),
                );
            }
        } else {
            // Handle other types of entries, like symbolic links
            println!("Found other type of entry: {:?}", entry_path);
        }
    }

    return results;
}

pub fn scan(args: Vec<String>) -> HashMap<String, i32> {
    let mut results: HashMap<String, i32> = HashMap::new();

    for arg in args {
        let ret = check_for_files(arg);
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

        output += "| ";
        output += k;
        output += " | ";
        output += &v.to_string();
        output += " |";
        output += "\n"
    }

    println!("+{}+", "-".repeat(len + 7));
    println!("{}", output);
    println!("+{}+", "-".repeat(len + 7));

    // println!("{:#?}", res)
}
