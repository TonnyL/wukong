#[macro_use]
extern crate prettytable;

use std::borrow::Cow;

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
    println!("Hello, world!");
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
    table.add_row(row!["Rank", "Full Name", "Description", "Language", "Stars(Total/Period)"]);
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
            Cell::new(&format!("{}/{}", repository.stars, repository.current_period_stars)),
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