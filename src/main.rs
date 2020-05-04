#[macro_use]
extern crate clap;
#[macro_use]
extern crate prettytable;

use std::borrow::Cow;
use std::error::Error;

use clap::{App, AppSettings, Arg};
use prettytable::{Cell, Row, Table};
use serde::Deserialize;

#[derive(Deserialize)]
struct Repository {
    author: String,
    name: String,
    description: String,
    language: Option<String>,
    stars: i32,
    #[serde(rename = "currentPeriodStars")]
    current_period_stars: i32,
}

#[derive(Deserialize)]
struct Developer {
    username: String,
    name: String,
    repo: Repo,
}

#[derive(Deserialize)]
struct Repo {
    name: String,
    description: String,
}

#[derive(Deserialize)]
struct Language {
    #[serde(rename = "urlParam")]
    url_param: String,
    name: String,
}

fn main() {
    let matches = App::new("wukong")
        .setting(AppSettings::SubcommandRequiredElseHelp)
        .about("A command-line tool for browsing GitHub trending written by Rust.")
        .version(crate_version!())
        .author(crate_authors!())
        .after_help("https://github.com/TonnyL/wukong/")
        .subcommand(
            App::new("repos")
                .about("See the developers that the GitHub community is most excited about")
                .aliases(&["r", "repositories", "repository"])
                .args(&[
                    Arg::with_name("lang")
                        .short('l')
                        .long("lang")
                        .takes_value(true)
                        .required(false)
                        .about("filter by programming language"),
                    Arg::with_name("period")
                        .short('p')
                        .long("period")
                        .takes_value(true)
                        .required(false)
                        .about("filter by time period"),
                    Arg::with_name("spoken_lang")
                        .short('s')
                        .long("spoken_language")
                        .takes_value(true)
                        .required(false)
                        .about("filter by spoken language"),
                ]),
        )
        .subcommand(
            App::new("devs")
                .about("These are the developers building the hot tools today")
                .aliases(&["d", "developers", "developer"])
                .args(&[
                    Arg::with_name("lang")
                        .short('l')
                        .long("language")
                        .takes_value(true)
                        .required(false)
                        .about("filter by programming language"),
                    Arg::with_name("period")
                        .short('p')
                        .long("period")
                        .takes_value(true)
                        .required(false)
                        .about("filter by time period"),
                ]),
        )
        .subcommand(
            App::new("langs")
                .about("List all the available programming language options")
                .aliases(&["l", "languages", "language"]),
        )
        .subcommand(
            App::new("spoken_langs")
                .about("List all the available spoken language options")
                .aliases(&["sl", "spoken-languages", "spoken-language", "spoken-lang"]),
        )
        .get_matches();

    match matches.subcommand() {
        ("langs", _) => {
            let result = list_languages();
            match result {
                Ok(value) => {
                    show_table_of_languages(value);
                }
                Err(e) => print_err_msg(e),
            }
        }
        ("spoken_langs", _) => {
            let result = list_spoken_language_codes();
            match result {
                Ok(value) => {
                    show_table_of_languages(value);
                }
                Err(e) => print_err_msg(e),
            }
        }
        ("repos", Some(sub_m)) => {
            let result = list_repositories(
                sub_m.value_of("lang").unwrap_or("").to_string(),
                sub_m.value_of("period").unwrap_or("daily").to_string(),
                sub_m.value_of("spoken_lang").unwrap_or("").to_string(),
            );
            match result {
                Ok(value) => {
                    show_table_of_repositories(value);
                }
                Err(e) => print_err_msg(e),
            }
        }
        ("devs", Some(sub_m)) => {
            let result = list_developers(
                sub_m.value_of("lang").unwrap_or("").to_string(),
                sub_m.value_of("period").unwrap_or("daily").to_string(),
            );

            match result {
                Ok(value) => {
                    show_table_of_developers(value);
                }
                Err(e) => print_err_msg(e),
            }
        }
        _ => {}
    }
}

/// Display the languages data as a table.
fn show_table_of_languages(languages: Vec<Language>) {
    let mut table = Table::new();
    table.add_row(row!["name", "value"]);
    for language in languages.iter() {
        table.add_row(Row::new(vec![
            Cell::new(&language.name),
            Cell::new(&language.url_param),
        ]));
    }
    println!("\n");
    table.printstd();
    println!();
}

/// Display the developers data as a table.
fn show_table_of_developers(developers: Vec<Developer>) {
    let mut table = Table::new();
    table.add_row(row!["Rank", "Name", "Repo Name", "Description"]);
    for (index, developer) in developers.iter().enumerate() {
        let developers_name = if developer.username.is_empty() {
            Cow::Borrowed(&developer.name)
        } else {
            Cow::Owned(format!("{}({})", developer.name, developer.username))
        };
        table.add_row(Row::new(vec![
            Cell::new(&(index as i32 + 1).to_string()),
            Cell::new(&developers_name),
            Cell::new(&developer.repo.name),
            Cell::new(&limit_string_with_break_lines(&developer.repo.description)),
        ]));
    }
    println!("\n");
    table.printstd();
    println!();
}

/// Display the repositories data as a table.
fn show_table_of_repositories(repositories: Vec<Repository>) {
    let mut table = Table::new();
    table.add_row(row![
        "Rank",
        "Full Name",
        "Description",
        "Language",
        "Stars(Total/Period)"
    ]);
    for (index, repository) in repositories.iter().enumerate() {
        let language = match &repository.language {
            Some(v) => v,
            None => "",
        };
        table.add_row(Row::new(vec![
            Cell::new(&(index as i32 + 1).to_string()),
            Cell::new(&format!("{}/{}", repository.author, repository.name)),
            Cell::new(&limit_string_with_break_lines(&repository.description)),
            Cell::new(language),
            Cell::new(&format!(
                "{}/{}",
                repository.stars, repository.current_period_stars
            )),
        ]));
    }
    println!("\n");
    table.printstd();
    println!();
}

fn limit_string_with_break_lines(input: &str) -> String {
    let mut result = String::from("");
    let i = String::from(input);

    let mut length = 0;

    for c in i.chars() {
        result.push(c);

        length += c.len_utf8();
        if length >= 24 {
            result.push_str("\n");

            length = 0;
        }
    }

    result
}

/// Receive an array of all options of certain programming languages(e.g. Rust, Golang).
fn list_languages() -> Result<Vec<Language>, Box<dyn Error>> {
    let resp = reqwest::blocking::get(
        "https://raw.githubusercontent.com/huchenme/github-trending-api/master/src/languages.json",
    )?
    .json::<Vec<Language>>()?;
    Ok(resp)
}

/// Receive an array of all options of certain spoken languages(e.g. Chinese, English)
fn list_spoken_language_codes() -> Result<Vec<Language>, Box<dyn Error>> {
    let resp = reqwest::blocking::get(
        "https://raw.githubusercontent.com/huchenme/github-trending-api/master/src/spoken-languages.json"
    )?.json::<Vec<Language>>()?;
    Ok(resp)
}

/// Receive an array of trending repositories.
///
/// # Arguments
///
/// * `lang` - Optional, list trending repositories of certain programming languages
/// * `period` - Optional, default to `daily`, possible values: `daily`, `weekly` and `monthly`
/// * `spoken_lang_code` - optional, list trending repositories of certain spoken languages (e.g English, Chinese)
///
fn list_repositories(
    lang: String,
    period: String,
    spoken_lang_code: String,
) -> Result<Vec<Repository>, Box<dyn Error>> {
    let resp = reqwest::blocking::get(&format!(
        "https://github-trending-api.now.sh/repositories?language={}&since={}&spoken_language_code={}", 
        lang, period, spoken_lang_code
    ))?
    .json::<Vec<Repository>>()?;
    Ok(resp)
}

/// Receive an array of trending developers.
///
/// # Arguments
///
/// * `lang` - Optional, list trending repositories of certain programming languages
/// * `period` - Optional, default to `daily`, possible values: `daily`, `weekly` and `monthly`
///
fn list_developers(lang: String, period: String) -> Result<Vec<Developer>, Box<dyn Error>> {
    let resp = reqwest::blocking::get(&format!(
        "https://github-trending-api.now.sh/developers?language={}&since={}",
        lang, period
    ))?
    .json::<Vec<Developer>>()?;
    Ok(resp)
}

fn print_err_msg(err: Box<dyn Error>) {
    eprintln!("wukong got an error: {}", err)
}
