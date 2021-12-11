use witgen::witgen;

use juniper::{
    graphql_object, EmptyMutation, EmptySubscription, FieldError, GraphQLEnum, RootNode, Variables,
};

#[derive(Clone, Copy, Debug)]
struct Context;
impl juniper::Context for Context {}

#[derive(Clone, Copy, Debug, GraphQLEnum)]
enum UserKind {
    Admin,
    User,
    Guest,
}

#[derive(Clone, Debug)]
struct User {
    id: i32,
    kind: UserKind,
    name: String,
}

#[graphql_object(context = Context)]
impl User {
    fn id(&self) -> i32 {
        self.id
    }

    fn kind(&self) -> UserKind {
        self.kind
    }

    fn name(&self) -> &str {
        &self.name
    }

    async fn friends(&self) -> Vec<User> {
        vec![]
    }
}

#[derive(Clone, Copy, Debug)]
struct Query;

#[graphql_object(context = Context)]
impl Query {
    async fn users() -> Vec<User> {
        vec![User {
            id: 1,
            kind: UserKind::Admin,
            name: "user1".into(),
        }]
    }

    // /// Fetch a URL and return the response body text.
    // async fn request(url: String) -> Result<String, FieldError> {
    //     Ok(reqwest::get(&url).await?.text().await?)
    // }
}

type Schema = RootNode<'static, Query, EmptyMutation<Context>, EmptySubscription<Context>>;

fn schema() -> Schema {
    Schema::new(
        Query,
        EmptyMutation::<Context>::new(),
        EmptySubscription::<Context>::new(),
    )
}
fn main() {}

#[witgen]
async fn run(query: String) -> String {
    // Create a context object.
    let ctx = Context;

    let m = EmptyMutation::new();
    let s = EmptySubscription::new();

    let v = Variables::new();

    let sch = Schema::new(Query, m, s);
    // Run the executor.
    let res = juniper::execute(
        &query, // "query { favoriteEpisode }",
        None,
        &sch,
        &v,
        &ctx,
    )
    .await;
    let (value, errors) = res.unwrap();

    (value.to_string())
    // Ensure the value matches.
}
