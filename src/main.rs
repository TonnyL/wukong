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
